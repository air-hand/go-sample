package web

import (
	"fmt"
	"net/http"
)

type Server struct {
	PortNumber int
}

func NewServer(portNumber int) *Server {
	return &Server{
		PortNumber: portNumber,
	}
}

func (rcv *Server) Serve() {
	config := NewAppConfig()

	db_config := NewDBConnectConfigFromEnv()
	db := NewDBClient(db_config)
	MigrateModels(db)

	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", rcv.PortNumber),
		Handler: Routes(db_config, &handler),
	}

	_ = server.ListenAndServe()
}
