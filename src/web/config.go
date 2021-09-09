package web

import "os"

type AppConfig struct {
	IsDebug bool
}

func NewAppConfig() *AppConfig {
	config := AppConfig{
		IsDebug: os.Getenv("DEBUG") != "",
	}
	return &config
}
