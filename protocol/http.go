package protocol

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/hello-world/config"
	"github.com/luyasr/hello-world/pkg/ioc"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer() *HttpServer {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	ioc.HttpHandler().RouteRegistry(apiV1)

	return &HttpServer{
		server: &http.Server{
			Addr:    config.C.HttpServer.Addr(),
			Handler: r,
		},
	}
}

func (h *HttpServer) Run() error {
	log.Print("http server running, listen on ", config.C.HttpServer.Addr())
	return h.server.ListenAndServe()
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	log.Print("http server shutdown")
	return h.server.Shutdown(ctx)
}
