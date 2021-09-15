package web

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"local.packages/types"
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
	Data     *types.KeyValueList
}

func (config *DBConnectConfig) DSN() string {
	config.Data.Add(types.KeyValue{Key: "charset", Value: "utf8mb4"})
	config.Data.Add(types.KeyValue{Key: "parseTime", Value: "True"})
	config.Data.Add(types.KeyValue{Key: "loc", Value: config.Timezone})
	config.Data = config.Data.StableSortBy(func(x types.KeyValue, y types.KeyValue) bool {
		return x.Key < y.Key
	}).DistinctBy(func(x types.KeyValue, y types.KeyValue) bool {
		return x.Key == y.Key
	})
	extras := config.Data.MapToString(func(kv types.KeyValue) string {
		bytes := make([]byte, 0, 128)
		bytes = append(bytes, kv.Key...)
		bytes = append(bytes, "="...)
		bytes = append(bytes, (url.QueryEscape(kv.Value))...)
		return string(bytes)
	})
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.User, config.Password, config.Host, config.Port, config.DBName, strings.Join(extras, "&"))
}

func NewDBConnectConfigFromEnv() *DBConnectConfig {
	return &DBConnectConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     3306,
		DBName:   os.Getenv("DB_NAME"),
		Timezone: os.Getenv("TZ"),
		Data:     types.NewKeyValueList(),
	}
}
