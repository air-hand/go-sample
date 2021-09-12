package web

import "os"

type AppConfig struct {
	IsDebug bool
	TZ      string
}

func NewAppConfig() *AppConfig {
	config := AppConfig{
		IsDebug: os.Getenv("DEBUG") != "",
		TZ:      "Asia/Tokyo",
	}
	return &config
}
