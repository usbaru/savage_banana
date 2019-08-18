package router

import (
	"src/src/handler"

	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.Func) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes(handler) {
		router.
			Methods(route.Method)
		Path(route.Pattern)
		Name(route.Name)
		Handler(route.HandlerFunc)
	}
	return router

}
