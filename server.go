package fxprovider

import (
	"context"
	"github.com/ValGoldun/fxprovider/fxcontext"
	"github.com/ValGoldun/logger"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	timeout "github.com/vearne/gin-timeout"
	"go.uber.org/fx"
	"net"
	"net/http"
	"os"
)

func ProvideServer(ctx *fxcontext.AppContext, gin *gin.Engine) *http.Server {
	var address = ":8080"

	if newAddress := ctx.ApplicationConfig().Application.HttpAddress; newAddress != "" {
		address = newAddress
	}

	return &http.Server{Addr: address, Handler: gin}
}

func ProvideServerEngine(ctx *fxcontext.AppContext) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(handler)

	handler.Use(gin.Recovery())
	handler.Use(
		timeout.Timeout(
			timeout.WithTimeout(ctx.ApplicationConfig().Application.ServerTimeout),
			timeout.WithDefaultMsg(""),
			timeout.WithCallBack(
				func(r *http.Request) {
					ctx.Logger().Error("handler timeout", logger.Field{
						Key:   "path",
						Value: r.URL.Path,
					})
				}),
		),
	)

	handler.GET("/state", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	return handler
}

func InvokeHTTPServer(lc fx.Lifecycle, server *http.Server, logger logger.Logger) error {
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
