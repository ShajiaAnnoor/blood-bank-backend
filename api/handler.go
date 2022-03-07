package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"go.uber.org/dig"
)

type params struct {
	dig.In
	Routes   []*routeutils.Route `group:"route"`
	WSRoutes []*routeutils.Route `group:"ws_route"`
}

//Handler provides all the normal and web socket routes
func Handler(p params) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, route := range p.Routes {
		r.Method(route.Method, route.Pattern, route.Handler)
	}

	// a separate function Mount is required to register web socket routes
	for _, route := range p.WSRoutes {
		r.Mount(route.Pattern, route.Handler)
	}

	return r
}
