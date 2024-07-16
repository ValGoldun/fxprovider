package config

import "time"

type Config struct {
	Application struct {
		ServerTimeout time.Duration
		HttpAddress   string
	}
}