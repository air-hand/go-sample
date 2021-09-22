package web

import (
	"context"
	"net/http"
)

func DatabaseMiddleware(config *DBConnectConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			db := NewDBClient(config)
			defer db.Close()
			ctx := context.WithValue(r.Context(), "DB", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
