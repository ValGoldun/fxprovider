package fxcontext

import (
	"context"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/fxprovider/fxconfig"
	"github.com/ValGoldun/logger"
)

const (
	env = iota + 1
	log
	conf
)

type AppContext struct {
	context.Context
}

func New() AppContext {
	return AppContext{context.Background()}
}

func (ctx *AppContext) WithEnvironment(env environment.Environment) {
	context.WithValue(ctx.Context, env, env)
}

func (ctx *AppContext) WithLogger(logger logger.Logger) {
	context.WithValue(ctx.Context, log, logger)
}

func (ctx *AppContext) WithApplicationConfig(config fxconfig.Application) {
	context.WithValue(ctx.Context, conf, config)
}
