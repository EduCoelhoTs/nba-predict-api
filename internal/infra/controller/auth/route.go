package authcontroller

import (
	"net/http"

	_http "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/http"
)

func (c *controller) GetRoutes() _http.Routes {
	return _http.Routes{
		"/auth": {
			{
				Method:      http.MethodPost,
				Path:        "/Login",
				HandlerFunc: c.Login,
			},
		},
	}
}
