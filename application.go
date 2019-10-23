package main

import (
	"context"
	"net/http"
	"src/src/handler"
	"src/src/repository"
	"src/src/router"
	"src/src/service"
)

var (
	ctx          = context.Background()
	handlerFuncs *handler.Funcs
)

func init() {
	savageRepository := repository.NewSavage(ctx)
	savageService := service.NewSavageService(savageRepository)
	handlerFuncs = &handler.Funcs{
		Ctx:    ctx,
		Savage: savageService,
	}
}

// Main function with Router
func main() {
	r := router.NewRouter(handlerFuncs)
	http.ListenAndServe(":5000", r)

}
