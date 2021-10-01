package web

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"local.packages/web/cache"
	"local.packages/web/db"
)

func DatabaseMiddleware(config *db.DBConnectConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			db := db.NewDBClient(config)
			defer db.Close()
			ctx := context.WithValue(r.Context(), "DB", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func SessionMiddleware(config *cache.CacheConnectConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cache_client := cache.NewCacheClient(config)
			store := NewSessionStore(cache_client)
			defer store.Close()
			store.Options(sessions.Options{
				MaxAge: 60 * 60 * 24,
			})
			session, err := store.Get(r, "session-key")
			if err != nil {
				log.Fatal("failed to get session:", err)
			}
			ctx := context.WithValue(r.Context(), "Session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
