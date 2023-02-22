package logger

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"strconv"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zap.Logger
}

func New(level, appEnvironment, appName string) *Logger {
	zapLevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		log.Fatalf("invalid log level: %s", err)
	}

	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.DisableCaller = true
	switch appEnvironment {
	case "prod":
		zapConfig = zap.NewProductionConfig()
	default:
		break
	}

	zapConfig.Level = zapLevel
	logger := zap.Must(zapConfig.Build())

	if appName != "" {
		logger = logger.With(zap.String("app", appName))
	}

	return &Logger{logger: logger}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.logger.Debug(l.msg(message), l.anyArgs(args)...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.logger.Info(l.msg(message), l.anyArgs(args)...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.logger.Warn(l.msg(message), l.anyArgs(args)...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.logger.Error(l.msg(message), l.anyArgs(args)...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.logger.Fatal(l.msg(message), l.anyArgs(args)...)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info(message)
	} else {
		fields := make([]zap.Field, 0, len(args))
		for i, arg := range args {
			fields = append(fields, zap.Any(strconv.Itoa(i), arg))
		}
		l.logger.Info(message, fields...)
	}
}

func (l *Logger) msg(message interface{}) string {
	switch msg := message.(type) {
	case string:
		return msg
	case error:
		return msg.Error()
	default:
		l.Error(fmt.Sprintf("message %v has unknown type %v", message, msg))
		return ""
	}
}

func (l *Logger) anyArgs(args []interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args))
	for i, arg := range args {
		switch a := arg.(type) {
		case error:
			fields = append(fields, zap.Error(a))
		case zap.Field:
			fields = append(fields, a)
		default:
			fields = append(fields, zap.Any(strconv.Itoa(i), a))
		}
	}
	return fields
}
