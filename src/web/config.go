package web

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type AppConfig struct {
	IsDebug bool
}

func NewAppConfig() *AppConfig {
	config := AppConfig{
		IsDebug: os.Getenv("DEBUG") != "",
	}
	return &config
}

type DBConnectConfig struct {
	User     string
	Password string
	Host     string
	Port     uint16
	DBName   string
	Timezone string
	Data     map[string]string
}

func (config *DBConnectConfig) DSN() string {
	config.Data["charset"] = "utf8mb4"
	config.Data["parseTime"] = "True"
	config.Data["loc"] = config.Timezone
	extras := make([]string, 0, 256)
	for k, v := range config.Data {
		bytes := make([]byte, 0, 128)
		bytes = append(bytes, k...)
		bytes = append(bytes, "="...)
		bytes = append(bytes, (url.QueryEscape(v))...)
		extras = append(extras, string(bytes))
	}
	extra_string := strings.Join(extras, "&")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.User, config.Password, config.Host, config.Port, config.DBName, extra_string)
}

func NewDBConnectConfigFromEnv() *DBConnectConfig {
	return &DBConnectConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     3306,
		DBName:   os.Getenv("DB_NAME"),
		Timezone: os.Getenv("TZ"),
		Data:     make(map[string]string),
	}
}
