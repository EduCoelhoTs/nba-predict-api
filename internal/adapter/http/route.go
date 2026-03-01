package _http

import "net/http"

type Routes = map[string][]Route

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middlewares []func(h http.Handler) http.Handler
}
