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

	renderer := NewTemplateRenderer(!config.IsDebug)
	handler := Handler{
		renderer: renderer,
	}

	cache_config := NewCacheConnectConfigFromEnv()
	db_config := NewDBConnectConfigFromEnv()

	var middlewares []MiddlewareFunc
	middlewares = append(middlewares, SessionMiddleware(cache_config))
	middlewares = append(middlewares, DatabaseMiddleware(db_config))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", rcv.PortNumber),
		Handler: Routes(middlewares, &handler),
	}

	_ = server.ListenAndServe()
}
