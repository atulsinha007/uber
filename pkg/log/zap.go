package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	L = getZapLogger()
)

func getZapLogger() *zap.Logger {
	logLevel := zap.DebugLevel
	encodingType := "json"

	logger, err := zap.Config{
		Encoding:         encodingType,
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Development:      false,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			StacktraceKey: "stacktrace",
		},
	}.Build()
	if err != nil {
		log.Fatal("Error in configuring zap")
	}

	return logger
}
