package myhttp

import (
	"github.com/samber/do"
	"net/http"
)

type RouteProvider interface {
	RegisterRoutes(engine http.Handler)
}

type Router struct {
	g http.Handler
}

func RouterProvider(i *do.Injector) (*Router, error) {
	r := &Router{}

	r.registerMiddlewares()
	r.registerApplicationRoutes()

	return r, nil
}

func (r *Router) Handler() http.Handler {
	return r.g
}

func (r *Router) registerMiddlewares() {

}

func (r *Router) registerApplicationRoutes(providers ...RouteProvider) {
	for _, provider := range providers {
		provider.RegisterRoutes(r.g)
	}
}
