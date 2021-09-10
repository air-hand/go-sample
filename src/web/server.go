package web

import (
	"fmt"
	"net/http"
)

func Serve(portNumber int) {
	config := NewAppConfig()
	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", portNumber),
		Handler: Routes(&handler),
	}

	_ = server.ListenAndServe()
}
