package fxcontext

import (
	"context"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/logger"
	"time"
)

const (
	env = iota + 1
	log
	serverTimeout
)

type AppContext struct {
	context.Context
}

func New() *AppContext {
	c := AppContext{context.Background()}
	return &c
}

func (ctx *AppContext) WithEnvironment(env environment.Environment) {
	context.WithValue(ctx.Context, env, env)
}

func (ctx *AppContext) WithLogger(logger logger.Logger) {
	context.WithValue(ctx.Context, log, logger)
}

func (ctx *AppContext) WithServerTimeout(timeout time.Duration) {
	context.WithValue(ctx.Context, serverTimeout, timeout)
}
