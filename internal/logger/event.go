package logger

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

func (l loggerImpl) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logOnStartExecuting(e)
	case *fxevent.OnStartExecuted:
		l.logOnStartExecuted(e)
	case *fxevent.OnStopExecuting:
	case *fxevent.OnStopExecuted:
	case *fxevent.Supplied:
		l.logSupplied(e)
	case *fxevent.Provided:
		l.logProvided(e)
	case *fxevent.Replaced:
	case *fxevent.Decorated:
	case *fxevent.Run:
	case *fxevent.Invoking:
		l.logInvoking(e)
	case *fxevent.Invoked:
		l.logInvoked(e)
	case *fxevent.Stopping:
	case *fxevent.Stopped:
	case *fxevent.RollingBack:
	case *fxevent.RolledBack:
	case *fxevent.Started:
	case *fxevent.LoggerInitialized:
	}
}

func (l loggerImpl) logOnStartExecuting(e *fxevent.OnStartExecuting) {
	flog := l.logger.WithFields(logrus.Fields{
		"caller": e.CallerName,
	})
	if l.logger.Level == logrus.TraceLevel {
		flog = flog.WithFields(logrus.Fields{"callee": e.FunctionName})
	}
	flog.Info("OnStart hook executing")
}

func (l loggerImpl) logOnStartExecuted(e *fxevent.OnStartExecuted) {
	flog := l.logger.WithFields(logrus.Fields{
		"caller": e.CallerName,
	})
	if l.logger.Level == logrus.TraceLevel {
		flog.WithFields(logrus.Fields{"callee": e.FunctionName})
	}
	if e.Err != nil {
		flog.WithFields(logrus.Fields{"error": e.Err}).Error("OnStart hook failed")
	} else {
		flog.WithFields(logrus.Fields{"runtime": e.Runtime.String()}).Info("OnStart hook executed")
	}
}
func (l loggerImpl) logProvided(e *fxevent.Provided) {
	flog := l.logger.WithFields(logrus.Fields{
		"module": moduleName(e.ModuleName),
	})
	if l.logger.Level == logrus.TraceLevel {
		flog = flog.WithFields(logrus.Fields{
			"stacktrace":  e.StackTrace,
			"moduletrace": e.ModuleTrace,
		})
	}
	for _, rtype := range e.OutputTypeNames {
		flog.WithFields(logrus.Fields{
			"constructor": e.ConstructorName,
			"type":        rtype,
			"private":     e.Private,
		}).Info("provided")
	}
	if e.Err != nil {
		flog.WithFields(logrus.Fields{
			"error": e.Err,
		}).Error("error encountered while applying options")
	}
}

func (l loggerImpl) logSupplied(e *fxevent.Supplied) {
	flog := l.logger.WithFields(logrus.Fields{
		"module": moduleName(e.ModuleName),
		"type":   e.TypeName,
	})
	if l.logger.Level == logrus.TraceLevel {
		flog = flog.WithFields(logrus.Fields{
			"stacktrace":  e.StackTrace,
			"moduletrace": e.ModuleTrace,
		})
	}
	if e.Err != nil {
		flog.WithFields(logrus.Fields{
			"error": e.Err,
		}).Error("error encountered while applying options")
	} else {
		flog.Info("supplied")
	}
}

func (l loggerImpl) logInvoking(e *fxevent.Invoking) {
	l.logger.WithFields(logrus.Fields{
		"function": e.FunctionName,
	}).Info("invoking")
}

func (l loggerImpl) logInvoked(e *fxevent.Invoked) {
	if e.Err != nil {
		flog := l.logger.WithFields(logrus.Fields{
			"module":   moduleName(e.ModuleName),
			"error":    e.Err,
			"function": e.FunctionName,
		})
		if l.logger.Level == logrus.TraceLevel {
			flog = flog.WithFields(logrus.Fields{
				"stack": e.Trace,
			})
		}
		flog.Error("invoke failed")
	}
}

func moduleName(moduleName string) string {
	if moduleName == "" {
		return "main"
	}
	return moduleName
}
