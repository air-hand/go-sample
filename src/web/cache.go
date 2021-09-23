package web

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewCacheClient(config *CacheConnectConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	})
}
