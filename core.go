package fxprovider

import (
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/fxprovider/config"
	fxlogger "github.com/ValGoldun/fxprovider/logger"
	"github.com/ValGoldun/logger"

	"go.uber.org/fx"
)

func ProvideApplicationCore[Config any](lifecycle fx.Lifecycle) (*appcontext.AppContext, logger.Logger, Config, error) {
	ctx := appcontext.New()
	var cfg Config

	cfg, err := config.New[Config](ctx)
	if err != nil {
		return nil, logger.Logger{}, cfg, err
	}

	logger, err := fxlogger.New(ctx, lifecycle)
	if err != nil {
		return nil, logger, cfg, err
	}

	return ctx, logger, cfg, nil
}
