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

func (ctx *AppContext) WithEnvironment(environment environment.Environment) {
	context.WithValue(ctx.Context, env, environment)
}

func (ctx *AppContext) WithLogger(logger logger.Logger) {
	context.WithValue(ctx.Context, log, logger)
}

func (ctx *AppContext) WithApplicationConfig(config fxconfig.Application) {
	context.WithValue(ctx.Context, conf, config)
}

func (ctx *AppContext) Environment() environment.Environment {
	return ctx.Value(env).(environment.Environment)
}

func (ctx *AppContext) Logger() logger.Logger {
	return ctx.Value(log).(logger.Logger)
}

func (ctx *AppContext) ApplicationConfig() fxconfig.Application {
	return ctx.Value(conf).(fxconfig.Application)
}
