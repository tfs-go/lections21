package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	SomeErr = errors.New("some error")
)

func main() {
	logger := logrus.New()
	if err := someFunc(logger); err != nil {
		switch {
		case errors.Is(err, SomeErr):
			logger.WithError(err).Error("caught some error on someFunc")
		default:
			logger.WithError(err).Error("caught unknown error")
		}
	}
}

func someFunc(logger logrus.FieldLogger) error {
	defer Elapsed(logger, time.Now(), "someFunc()")

	time.Sleep(time.Second)
	if err := doSomething(); err != nil {
		return fmt.Errorf("doSomething error: %w", err)
	}

	return nil
}

func doSomething() error {
	return SomeErr
}

func Elapsed(logger logrus.FieldLogger, startedAt time.Time, msg string, args ...interface{}) {
	logger.WithField("elapsed_ms", elapsedMs(startedAt)).Infof(msg, args...)
}

func DurationToMs(d time.Duration) float64 {
	return float64(d.Nanoseconds()) / 1000.0 / 1000.0
}

func elapsedMs(since time.Time) float64 {
	return DurationToMs(time.Since(since))
}
