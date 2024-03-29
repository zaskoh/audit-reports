package main

import (
	"sync"

	"github.com/zaskoh/c4c-reports/c4c"
	"github.com/zaskoh/c4c-reports/config"
	"github.com/zaskoh/c4c-reports/logger"
	"github.com/zaskoh/c4c-reports/sherlock"
	"github.com/zaskoh/discordbooter"
	"go.uber.org/zap"
)

var (
	// WaitGroup for handling clean shutdown
	wg sync.WaitGroup
	// will be set on build
	version = "dev"
)

func main() {
	if err := startup(); err != nil {
		logger.Error(
			"Startup failed",
			zap.Error(err),
		)
	}
}

func startup() error {
	logger.Info("booting c4c-reports - " + version)
	ctx, cancelFunc, cancelChan := config.CreateLaunchContext()
	defer cancelFunc()

	// boot discordbot if active
	if config.DiscordConfig.Active {
		logger.Info("booting discord bot")
		err := discordbooter.Start(ctx, &wg, config.DiscordConfig.Token)
		if err != nil {
			return err
		}
	}

	// boot c4c & sherlock report updater
	c4c.Start()
	sherlock.Start()

	// keep it running until we cancel
	<-cancelChan
	logger.Info("waiting for cleanup before shutdown")

	// use ctx, so you can delete this block here!!!
	//time.Sleep(2 * time.Second)
	ctx.Done()
	// use it, so you can delete this block here!!!

	wg.Wait()
	return nil
}
