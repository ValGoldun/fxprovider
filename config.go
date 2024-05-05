package fxprovider

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func ProvideConfig[T any]() (T, error) {
	var cfg T

	var env, ok = os.LookupEnv("APP_ENV")

	if ok {
		viper.SetConfigName(fmt.Sprintf("%s.config", env))
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

	return cfg, nil
}
