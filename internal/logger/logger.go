package logger

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

type Logger interface {
	LogEvent(fxevent.Event)
	Error(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Logger() *logrus.Logger
	InterceptorLogger() logging.Logger
}

type loggerImpl struct {
	logger *logrus.Logger
}

func NewLogger() Logger {
	l := logrus.New()
	l.Level = logrus.DebugLevel
	return &loggerImpl{
		logger: l,
	}
}

func (l loggerImpl) Error(msg string, data ...interface{}) {
	l.logger.Error(msg)
}

func (l loggerImpl) Info(msg string, data ...interface{}) {
	l.logger.Info(msg)
}

func (l loggerImpl) Debug(msg string, data ...interface{}) {
	l.logger.Debug(msg)
}

func (l loggerImpl) Logger() *logrus.Logger {
	return l.logger
}

func (l loggerImpl) InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make(map[string]any, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		if i.Next() {
			k, v := i.At()
			f[k] = v
		}
		log := l.logger.WithFields(f)

		switch lvl {
		case logging.LevelDebug:
			log.Debug(msg)
		case logging.LevelInfo:
			log.Info(msg)
		case logging.LevelWarn:
			log.Warn(msg)
		case logging.LevelError:
			log.Error(msg)
		default:
			log.Errorf("Attempted to log with unknown level %v", lvl)
		}
	})
}
