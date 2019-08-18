package router

import (
	"net/http"
	"src/src/handler"
)

type route struct {
	Name        string
	Methd       string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func routes(handler *handler.Funcs) []route {
	return []route{
		{
			Name:        "GetSavage",
			Method:      http.MethodGet,
			Pattern:     "/get",
			HandlerFunc: handler.GetSavage,
		},
	}
}
