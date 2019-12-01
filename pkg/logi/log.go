package logi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	SetLogger(Logger(true))
}

func SetLogger(l *zap.Logger) {
	Log = l
}

func Logger(development bool) *zap.Logger {
	var zapConfig zap.Config

	if development {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
	}
	log, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return log
}
