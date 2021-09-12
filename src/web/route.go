package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(handler *Handler) http.Handler {
	muxer := chi.NewRouter()
	muxer.Use(middleware.Recoverer)
	muxer.Use(middleware.Logger)
	muxer.Get("/", handler.Home)
	muxer.Get("/about", handler.About)
	muxer.Get("/now", handler.NowTime)
	return muxer
}
