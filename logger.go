package fxprovider

import (
	"context"
	"errors"
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/logger"
	"go.uber.org/fx"
	"syscall"
)

func newLogger(ctx *appcontext.AppContext, lc fx.Lifecycle) (logger.Logger, error) {
	l, err := logger.New(logger.Info)
	if err != nil {
		return logger.Logger{}, err
	}

	ctx.WithLogger(l)

	lc.Append(fx.Hook{OnStop: func(_ context.Context) error {
		err = l.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			return err
		}
		return nil
	}})

	return l, nil
}
