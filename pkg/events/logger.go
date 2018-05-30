package events

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type logger struct {
	ctx context.Context
}

func (log logger) le() *logrus.Entry {
	ctx := GetContext(log.ctx)
	return logrus.WithFields(logrus.Fields(ctx))
}

func (log logger) Warnf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.le().Warn(err.Error())

	trackError(log.ctx, err, "WARN")
}

func (log logger) Errorf(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	log.le().Error(err.Error())

	trackError(log.ctx, err, "ERROR")
}

func (log logger) Infof(format string, args ...interface{}) {
	log.le().Infof(format, args...)
}

func (log logger) Debugf(format string, args ...interface{}) {
	log.le().Debugf(format, args...)
}

func Log(ctx context.Context) Logger {
	return logger{
		ctx: ctx,
	}
}
