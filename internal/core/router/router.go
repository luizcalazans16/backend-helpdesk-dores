package router

import (
	"backend-helpdesk-dores/internal/core/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRoute() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
