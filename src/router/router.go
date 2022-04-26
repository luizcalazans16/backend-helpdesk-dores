package router

import (
	"github.com/gorilla/mux"
)

func generateRoute() *mux.Router {
	r := mux.NewRouter()

	return route.Config(r)
}
