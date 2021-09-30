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
	cache_config := NewCacheConnectConfigFromEnv()
	db_config := NewDBConnectConfigFromEnv()

	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", rcv.PortNumber),
		Handler: Routes(db_config, cache_config, &handler),
	}

	_ = server.ListenAndServe()
}
