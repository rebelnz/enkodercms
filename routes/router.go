package routes

import (
	h "bitbucket.org/enkdr/enkoder/handlers"
	"bitbucket.org/enkdr/enkoder/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.Logger(handler, route.Name)
		if route.Protected {
			handler = middleware.Auth(handler, route.Name)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.NotFoundHandler = http.HandlerFunc(h.NotFound)
	return router
}
