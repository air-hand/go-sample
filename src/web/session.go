package web

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
)

func NewSessionStore(client *redis.Client) {
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal(err)
		return
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{})
}
