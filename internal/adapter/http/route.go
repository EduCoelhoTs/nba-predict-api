package http

import "net/http"

type Routes = []Route

type Route struct {
	Method      string
	Path        string
	HandlerFunc func()
	Middlewares []func(h http.Handler) http.Handler
}
