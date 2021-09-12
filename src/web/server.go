package web

import (
	"fmt"
	"net/http"
)

func Serve(portNumber int) {
	config := NewAppConfig()
	db := NewDBClient(config)
	MigrateModels(db)

	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", portNumber),
		Handler: Routes(config, &handler),
	}

	_ = server.ListenAndServe()
}
