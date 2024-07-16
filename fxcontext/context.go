package fxcontext

import (
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/fxprovider/fxconfig"
	"github.com/ValGoldun/logger"
)

type AppContext struct {
	environment environment.Environment
	logger      logger.Logger
	appConfig   fxconfig.Config
}

func New() *AppContext {
	return new(AppContext)
}

func (ctx *AppContext) WithEnvironment(environment environment.Environment) {
	ctx.environment = environment
}

func (ctx *AppContext) WithLogger(logger logger.Logger) {
	ctx.logger = logger
}

func (ctx *AppContext) WithApplicationConfig(config fxconfig.Config) {
	ctx.appConfig = config
}

func (ctx *AppContext) Environment() environment.Environment {
	return ctx.environment
}

func (ctx *AppContext) Logger() logger.Logger {
	return ctx.logger
}

func (ctx *AppContext) ApplicationConfig() fxconfig.Config {
	return ctx.appConfig
}
