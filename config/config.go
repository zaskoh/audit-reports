package config

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
)

func init() {
	err := loadConfiguration()
	if err != nil {
		os.Exit(1)
	}
}

type loadConfig struct {
	BaseConfig    baseConfig    `yaml:"base"`
	AnotherConfig anotherConfig `yaml:"another"`
}

type baseConfig struct {
	LogLevel string `yaml:"log_level" env:"LOG_LEVEL" env-default:"debug"`
}

type anotherConfig struct {
	BooleanExample bool   `yaml:"boolean_example" env-default:"false"`
	StringExample  string `yaml:"string_example" env:"EXAMPLE_ENV_VAR" env-default:""`
	NotInConfigYml string `yaml:"token" env:"SUPER_SECRET_ENV" env-default:"i am a default value"`
}

// Base contains all the basic configurations
var Base baseConfig

// Add other configs....
var AnotherExample anotherConfig

func loadConfiguration() error {
	var confLoad loadConfig
	if err := cleanenv.ReadConfig("./config.yml", &confLoad); err != nil {
		return err
	}

	Base = confLoad.BaseConfig
	AnotherExample = confLoad.AnotherConfig
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
