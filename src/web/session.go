package web

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

func NewSessionStore(client *redis.Client) *redisstore.RedisStore {
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	store.KeyPrefix("session_")
	return store
}
