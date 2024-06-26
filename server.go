package fxprovider

import (
	"context"
	"github.com/ValGoldun/logger"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	timeout "github.com/vearne/gin-timeout"
	"go.uber.org/fx"
	"net"
	"net/http"
	"os"
	"time"
)

func ProvideServer(gin *gin.Engine) *http.Server {
	var address = ":8080"

	newAddress, ok := os.LookupEnv("HTTP_ADDRESS")
	if ok {
		address = newAddress
	}

	return &http.Server{Addr: address, Handler: gin}
}

func ProvideServerEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()

	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(handler)

	handler.Use(gin.Recovery())
	handler.Use(timeout.Timeout(timeout.WithTimeout(5*time.Second), timeout.WithDefaultMsg("")))

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
