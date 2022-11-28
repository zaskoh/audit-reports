package logger

import (
	"os"

	"github.com/zaskoh/go-starter/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {

	var err error
	var conf zap.Config
	var encoderConfig zapcore.EncoderConfig

	loggerMode := ""
	if config.Base.LogLevel == "production" {
		conf = zap.NewProductionConfig()
		encoderConfig = zap.NewProductionEncoderConfig()
		loggerMode = "production mode"
	} else {
		conf = zap.NewDevelopmentConfig()
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		loggerMode = "development mode"
	}

	encoderConfig.StacktraceKey = "" // hide stacktrace info
	conf.EncoderConfig = encoderConfig
	zapLog, err = conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		os.Exit(4)
	}

	zapLog.Info("Started logger in " + loggerMode)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(message string, fields ...zap.Field) {
	zapLog.Debug(message, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(message string, fields ...zap.Field) {
	zapLog.Warn(message, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields...)
}
