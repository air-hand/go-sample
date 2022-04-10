package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type MiddlewareFunc = func(http.Handler) http.Handler

func Routes(middlewares []MiddlewareFunc, handler *Handler) http.Handler {
	router := chi.NewRouter()

	for _, m := range middlewares {
		router.Use(m)
	}
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handler.Home)
	router.Get("/about", handler.About)
	router.Get("/now", handler.NowTime)
	router.Get("/db", handler.DBConn)
	router.Get("/users", handler.ListUsers)
	return router
}
