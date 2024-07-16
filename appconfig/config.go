package appconfig

import (
	"github.com/ValGoldun/fxprovider/healthcheck"
	"time"
)

type Config struct {
	Application struct {
		ServerTimeout         time.Duration
		HttpAddress           string
		HealthCheckFailPolicy healthcheck.FailPolicy
	}
}
