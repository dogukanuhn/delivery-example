package logger

import (
	"github.com/dogukanuhn/delivery-system/cfg"
	"github.com/dogukanuhn/delivery-system/internal/logger/logrus_hook"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ILogger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
}

type Logger struct {
	logger *logrus.Logger
}

func NewInstance() *Logger {

	log := logrus.New()
	return &Logger{logger: log}
}

func WithCollection(collection *mongo.Collection) *Logger {

	log := logrus.New()
	hooker, _ := logrus_hook.NewLogrusMongoHook(cfg.GetDatabase().Collection("logs"))
	log.Hooks.Add(hooker)

	return &Logger{logger: log}
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}
