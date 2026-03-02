package _chi

import (
	"net/http"

	_http "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/http"
	"github.com/go-chi/chi/v5"
)

type chiHandler struct {
	handler *chi.Mux
}

func NewChiHandler() *chiHandler {
	return &chiHandler{
		handler: chi.NewRouter(),
	}
}

func (h *chiHandler) RegisterRoutes(routes _http.Routes) http.Handler {
	for key, route := range routes {
		h.handler.Route(key, func(rt chi.Router) {
			for _, r := range route {
				rt.With(r.Middlewares...).Method(r.Method, r.Path, r.HandlerFunc)
			}
		})
	}

	return h.handler
}
