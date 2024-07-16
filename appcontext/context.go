package appcontext

import (
	"github.com/ValGoldun/fxprovider/appconfig"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/fxprovider/healthcheck"
	"github.com/ValGoldun/logger"
)

type AppContext struct {
	environment    environment.Environment
	logger         logger.Logger
	appConfig      appconfig.Config
	healthCheckers *healthcheck.Checkers
}

func New() *AppContext {
	return &AppContext{
		healthCheckers: healthcheck.New(),
	}
}

func (ctx *AppContext) WithEnvironment(environment environment.Environment) {
	ctx.environment = environment
}

func (ctx *AppContext) WithLogger(logger logger.Logger) {
	ctx.logger = logger
}

func (ctx *AppContext) WithApplicationConfig(config appconfig.Config) {
	ctx.appConfig = config
}

func (ctx *AppContext) WithHealthChecker(checker healthcheck.Checker) {
	ctx.healthCheckers.Add(checker)
}

func (ctx *AppContext) Environment() environment.Environment {
	return ctx.environment
}

func (ctx *AppContext) Logger() logger.Logger {
	return ctx.logger
}

func (ctx *AppContext) ApplicationConfig() appconfig.Config {
	return ctx.appConfig
}

func (ctx *AppContext) HealthCheckers() *healthcheck.Checkers {
	return ctx.healthCheckers
}
