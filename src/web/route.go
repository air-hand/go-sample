package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(db_config *DBConnectConfig, handler *Handler) http.Handler {
	router := chi.NewRouter()
	router.Use(DatabaseMiddleware(db_config))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handler.Home)
	router.Get("/about", handler.About)
	router.Get("/now", handler.NowTime)
	router.Get("/db", handler.DBConn)
	return router
}
