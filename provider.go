package fxprovider

import (
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/logger"
	"go.uber.org/fx"
)

func ProvideApplicationCore[Config any](lifecycle fx.Lifecycle) (*appcontext.AppContext, logger.Logger, Config, error) {
	ctx := appcontext.New()
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
