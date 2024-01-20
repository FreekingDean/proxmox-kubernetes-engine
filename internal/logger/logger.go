package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type Logger interface {
	LogEvent(fxevent.Event)
	Error(string, ...interface{})
	Info(string, ...interface{})
}

type loggerImpl struct {
	logger *logrus.Logger
}

func NewLogger() Logger {
	return &loggerImpl{
		logger: logrus.New(),
	}
}

func (l loggerImpl) Error(msg string, data ...interface{}) {
	l.logger.Error(msg)
}

func (l loggerImpl) Info(msg string, data ...interface{}) {
	l.logger.Info(msg)
}
