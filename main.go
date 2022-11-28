package main

import (
	"sync"
	"time"

	"github.com/zaskoh/go-starter/config"
	"github.com/zaskoh/go-starter/logger"
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
	ctx, cancelFunc, cancelChan := config.CreateLaunchContext()
	defer cancelFunc()

	// logic goes here.....
	logger.Debug("Config LogLevel: " + config.Base.LogLevel)
	logger.Error("config not set in config.yml: " + config.AnotherExample.NotInConfigYml)
	logger.Info("Adding config-values can be done in /config/config.go")
	logger.Warn("Version: " + version)
	// logic goes here....

	// keep it running until we cancel
	<-cancelChan
	logger.Info("waiting for cleanup before shutdown")

	// use ctx, so you can delete this block here!!!
	time.Sleep(2 * time.Second)
	ctx.Done()
	// use it, so you can delete this block here!!!

	wg.Wait()
	logger.Info("everything finished shutting down - bye")
	return nil
}
