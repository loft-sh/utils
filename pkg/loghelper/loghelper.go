package loghelper

import (
	"fmt"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Logger interface {
	WithName(name string) Logger
	WithoutName() Logger
	WithValues(keysAndValues ...interface{}) Logger
	V(level int) Logger
	Info(message string, keysAndValues ...interface{})
	Infof(format string, a ...interface{})
	Debug(message string, keysAndValues ...interface{})
	Debugf(format string, a ...interface{})
	Error(err error, message string, keysAndValues ...interface{})
	Errorf(format string, a ...interface{})
}

type logger struct {
	logr.Logger
}

func New(name string) Logger {
	log := &logger{
		ctrl.Log.WithName(name),
	}
	return log
}

func (l *logger) WithName(name string) Logger {
	return &logger{
		l.Logger.WithName(name),
	}
}

func (l *logger) WithoutName() Logger {
	return &logger{
		l.Logger.WithName(""),
	}
}

func (l *logger) WithValues(keysAndValues ...interface{}) Logger {
	return &logger{
		l.Logger.WithValues(keysAndValues...),
	}
}

func (l *logger) V(level int) Logger {
	return &logger{
		l.Logger.V(level),
	}
}

func (l *logger) Info(message string, keysAndValues ...interface{}) {
	l.Logger.Info(message, keysAndValues...)
}

func (l *logger) Infof(format string, a ...interface{}) {
	l.Logger.Info(fmt.Sprintf(format, a...))
}

func (l *logger) Debug(message string, keysAndValues ...interface{}) {
	l.Logger.V(1).Info(message, keysAndValues...)
}

func (l *logger) Debugf(format string, a ...interface{}) {
	l.Logger.V(1).Info(fmt.Sprintf(format, a...))
}

func (l *logger) Error(err error, message string, keysAndValues ...interface{}) {
	l.Logger.Error(err, message, keysAndValues...)
}

func (l *logger) Errorf(format string, a ...interface{}) {
	l.Logger.Error(fmt.Errorf(format, a...), "")
}

func Infof(format string, a ...interface{}) {
	(&logger{ctrl.Log}).Infof(format, a...)
}

func Debugf(format string, a ...interface{}) {
	(&logger{ctrl.Log}).Debugf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	(&logger{ctrl.Log}).Errorf(format, a...)
}
