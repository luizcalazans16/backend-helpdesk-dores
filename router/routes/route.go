package routes

import (
	"backend-helpdesk-dores/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := routeTicket
	routes = append(routes, routeLogin)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.URI,
				middleware.Logger(middleware.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
