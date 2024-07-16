package config

import (
	"fmt"
	"github.com/ValGoldun/fxprovider/appcontext"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/fxprovider/healthcheck"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	Application struct {
		ServerTimeout         time.Duration
		HttpAddress           string
		HealthCheckFailPolicy healthcheck.FailPolicy
	}
}

func New[T any](ctx *appcontext.AppContext) (T, error) {
	var cfg T
	var appConfig Config

	var env, ok = os.LookupEnv("APP_ENV")

	if ok {
		appEnv, err := environment.ParseEnvironment(env)
		if err != nil {
			return cfg, err
		}

		ctx.WithEnvironment(appEnv)

		viper.SetConfigName(fmt.Sprintf("%s.config", appEnv.String()))
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return cfg, err
	}

	ctx.WithApplicationConfig(appConfig)

	return cfg, nil
}
