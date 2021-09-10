package web

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func Routes(handler *Handler) http.Handler {
	muxer := pat.New()
	muxer.Get("/", http.HandlerFunc(handler.Home))
	muxer.Get("/about", http.HandlerFunc(handler.About))
	muxer.Get("/now", http.HandlerFunc(handler.NowTime))
	return muxer
}
