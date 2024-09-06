package logger

import (
	"context"

	"github.com/is1ab/Arvosana/env"
	"go.uber.org/zap"
)

type loggerKey struct{}

func NewLogger() *zap.SugaredLogger {
	var z *zap.Logger
	if env.DEV {
		z = zap.Must(zap.NewDevelopment())
	} else {
		z = zap.Must(zap.NewProduction())
	}

	return z.Sugar()
}

func Close(logger *zap.SugaredLogger) error {
	return logger.Sync()
}

func WithContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func Ctx(ctx context.Context) *zap.SugaredLogger {
	logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger)
	if !ok {
		panic("no logger context")
	}
	return logger
}
