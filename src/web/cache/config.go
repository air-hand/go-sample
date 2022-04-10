package cache

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type CacheConnectConfig struct {
	Host string
	Port uint
	DB   uint8
}

func (config *CacheConnectConfig) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		DB:   int(config.DB),
	}
}

func NewCacheConnectConfigFromEnv() *CacheConnectConfig {
	return &CacheConnectConfig{
		Host: os.Getenv("CACHE_HOST"),
		Port: 6379,
		DB:   0,
	}
}
