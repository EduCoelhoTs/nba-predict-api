package _http

import (
	"net/http"
	"time"
)

type httpServer struct {
	svr *http.Server
}

type HttpHandler interface {
	RegisterRoutes(routes Routes) http.Handler
}

func NewHttpServer(handler HttpHandler, routes Routes, addr *string) *httpServer {

	h := handler.RegisterRoutes(routes)

	if addr == nil {
		addr = new(string)
		*addr = ":8080"
	}

	return &httpServer{
		svr: &http.Server{
			Addr:              *addr,
			Handler:           h,
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
			ReadHeaderTimeout: 15 * time.Second,
		},
	}
}

func (s *httpServer) Start() error {
	return s.svr.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	return s.svr.Close()
}
