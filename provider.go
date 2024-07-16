package fxprovider

import (
	"github.com/ValGoldun/fxprovider/fxcontext"
	"github.com/ValGoldun/logger"
	"go.uber.org/fx"
)

func ProvideApplicationCore[Config any](lifecycle fx.Lifecycle) (*fxcontext.AppContext, logger.Logger, Config, error) {
	ctx := fxcontext.New()
	var config Config

	config, err := newConfig[Config](ctx)
	if err != nil {
		return nil, logger.Logger{}, config, err
	}

	logger, err := newLogger(ctx, lifecycle)
	if err != nil {
		return nil, logger, config, err
	}

	return ctx, logger, config, nil
}
