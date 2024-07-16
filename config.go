package fxprovider

import (
	"fmt"
	"github.com/ValGoldun/fxprovider/environment"
	"github.com/ValGoldun/fxprovider/fxcontext"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config[T any] struct {
	Base          T
	ServerTimeout time.Duration
}

func ProvideConfig[T any](ctx *fxcontext.AppContext) (T, error) {
	var cfg Config[T]

	var env, ok = os.LookupEnv("APP_ENV")

	if ok {
		appEnv, err := environment.ParseEnvironment(env)
		if err != nil {
			return cfg.Base, err
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
		return cfg.Base, err
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg.Base, err
	}

	ctx.WithServerTimeout(cfg.ServerTimeout)

	return cfg.Base, nil
}
