package fxprovider

import (
	"context"
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/fxprovider/healthcheck"
	"github.com/ValGoldun/logger"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	timeout "github.com/vearne/gin-timeout"
	"go.uber.org/fx"
	"net"
	"net/http"
	"os"
)

func ProvideServerHTTP(appCtx *appcontext.AppContext) (*gin.Engine, *http.Server) {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(handler)

	handler.Use(gin.Recovery())
	handler.Use(
		timeout.Timeout(
			timeout.WithTimeout(appCtx.ApplicationConfig().Application.ServerTimeout),
			timeout.WithDefaultMsg(""),
			timeout.WithCallBack(
				func(r *http.Request) {
					appCtx.Logger().Error("handler timeout", logger.Field{
						Key:   "path",
						Value: r.URL.Path,
					})
				}),
		),
	)

	handler.GET("/state", func(ctx *gin.Context) {
		health := appCtx.HealthCheckers().HealthCheck(appCtx.ApplicationConfig().Application.HealthCheckFailPolicy)

		if !health.IsOK {
			appCtx.Logger().Error("health check failed", logger.Field{
				Key:   "health",
				Value: health.String(),
			})

			ctx.JSON(http.StatusInternalServerError, health)
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusOK, health)
	})

	var address = ":8080"

	if newAddress := appCtx.ApplicationConfig().Application.HttpAddress; newAddress != "" {
		address = newAddress
	}

	return handler, &http.Server{Addr: address, Handler: handler}
}

func InvokeServerHTTP(appCtx *appcontext.AppContext, lc fx.Lifecycle, server *http.Server, logger logger.Logger, checkers ...healthcheck.Checker) error {
	for _, checker := range checkers {
		appCtx.WithHealthChecker(checker)
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	go func() {
		err = server.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return nil

}
