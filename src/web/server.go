package web

import (
	"net/http"
)

func Serve() {
	config := NewAppConfig()
	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/now", handler.NowTime)
	_ = http.ListenAndServe(":80", nil)
}
