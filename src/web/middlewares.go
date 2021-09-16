package web

import (
	"context"
	"net/http"
	"time"
)

func DatabaseMiddleware(config *DBConnectConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			db := NewDBClient(config)
			// TODO: not avoid the cancelFunc
			timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
			ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
