package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewPool(dsn string, logger logrus.FieldLogger) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Logger = &loggerAdapter{logger: logger}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

type loggerAdapter struct {
	logger logrus.FieldLogger
}

func (l *loggerAdapter) Log(_ context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	switch level {
	case pgx.LogLevelTrace:
		l.logger.WithFields(data).Debugf(msg)
	case pgx.LogLevelDebug:
		l.logger.WithFields(data).Debugf(msg)
	case pgx.LogLevelInfo:
		l.logger.WithFields(data).Infof(msg)
	case pgx.LogLevelWarn:
		l.logger.WithFields(data).Warnf(msg)
	case pgx.LogLevelError:
		l.logger.WithFields(data).Errorf(msg)
	case pgx.LogLevelNone:
		l.logger.WithFields(data).Errorf(msg)
	}
}
