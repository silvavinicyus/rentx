package routes

import (
	"net/http"
	"rentx/src/framework/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configurate(r *mux.Router) *mux.Router {
	routes := categoriesRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return r
}
