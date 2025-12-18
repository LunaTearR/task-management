package pkg

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppLog interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string, fields ...zap.Field)
	Error(msg interface{}, err error)
}

type appLogsZap struct {
	log *zap.Logger
}

func NewAppLogZap() AppLog {
	var log *zap.Logger

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &appLogsZap{log: log}
}

func (l *appLogsZap) Info(msg string) {
	l.log.Info(msg)
}

func (l *appLogsZap) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *appLogsZap) Warning(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

func (l *appLogsZap) Error(msg interface{}, err error) {
	fields := []zap.Field{}

	if err != nil {
		fields = append(fields, zap.Error(err))
	}

	switch v := msg.(type) {
	case error:
		l.log.Error(v.Error(), fields...)
	case string:
		l.log.Error(v, fields...)
	case nil:
		if err != nil {
			l.log.Error(err.Error(), fields...)
		} else {
			l.log.Error("Unknown error", fields...)
		}
	default:
		l.log.Error(fmt.Sprintf("%v", v), fields...)
	}

}
