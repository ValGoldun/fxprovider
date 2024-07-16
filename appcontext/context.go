package appcontext

import (
	"github.com/ValGoldun/fxprovider/config"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/logger"
)

type AppContext struct {
	environment environment.Environment
	logger      logger.Logger
	appConfig   config.Config
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

func (ctx *AppContext) WithApplicationConfig(config config.Config) {
	ctx.appConfig = config
}

func (ctx *AppContext) Environment() environment.Environment {
	return ctx.environment
}

func (ctx *AppContext) Logger() logger.Logger {
	return ctx.logger
}

func (ctx *AppContext) ApplicationConfig() config.Config {
	return ctx.appConfig
}
