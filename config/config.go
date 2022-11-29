package config

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	configPath = flag.String("config", "config.yml", "path to config file")
	_          = flag.Bool("no-discord", false, "use to deactivate the discord logging")
)

func init() {
	flag.Parse()
	err := loadConfiguration()
	if err != nil {
		os.Exit(1)
	}
}

type loadConfig struct {
	BaseConfig    baseConfig    `yaml:"base"`
	DiscordConfig discordConfig `yaml:"discord"`
}

type baseConfig struct {
	LogLevel            string `yaml:"log_level" env:"LOG_LEVEL" env-default:"debug"`
	ReportFile          string `yaml:"report_file" env:"C4C_REPORT_FILE" env-default:"./c4c-reports-backup.json"`
	ReportCheckInterval int    `yaml:"report_check_interval" env-default:"3600"`
}

type discordConfig struct {
	Active  bool
	Token   string `yaml:"token" env:"DISCORD_TOKEN" env-default:""`
	Channel string `yaml:"channel" env:"DISCORD_CHANNEL" env-default:""`
}

// Base contains all the basic configurations
var Base baseConfig

// Add other configs....
var DiscordConfig discordConfig

func loadConfiguration() error {
	var confLoad loadConfig

	if _, err := os.Stat(*configPath); err == nil {
		// if we have a config, load
		if err := cleanenv.ReadConfig(*configPath, &confLoad); err != nil {
			return err
		}
	} else {
		// if config.yml not exists, we just load the env vars
		if err := cleanenv.ReadEnv(&confLoad); err != nil {
			return err
		}
	}

	Base = confLoad.BaseConfig
	DiscordConfig = confLoad.DiscordConfig
	DiscordConfig.Active = !isFlagPassed("no-discord")
	return nil
}

// CreateLaunchContext for handling shutdowns and get a global context
func CreateLaunchContext() (context.Context, func(), chan bool) {
	interruptChan := make(chan os.Signal, 1)
	canceledChanChan := make(chan bool, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	ctx, cancelCtx := context.WithCancel(context.Background())
	go func() {
		defer close(interruptChan)
		<-interruptChan
		cancelCtx()
		canceledChanChan <- true
	}()
	cancel := func() {
		cancelCtx()
		close(canceledChanChan)
	}
	return ctx, cancel, canceledChanChan
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
