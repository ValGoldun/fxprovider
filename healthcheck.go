package fxprovider

import (
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/fxprovider/healthcheck"
)

func AddHealthCheckers(ctx *appcontext.AppContext, checkers ...healthcheck.Checker) {
	for _, checker := range checkers {
		ctx.WithHealthChecker(checker)
	}
}
