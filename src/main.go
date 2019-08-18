package main

import (
	"context"
	"src/src/handler"
	"src/src/router"
	"src/src/service"
)

var (
	ctx          = context.Background()
	handlerFuncs *handlerFuncs
)

func init() {
	savageRepository := repository.NewSavage(ctx)
	savageService := service.NewSavageService(savageRepository)
	handlerFuncs = &handler.Funcs{
		Ctx:    ctx,
		Savage: savageService,
	}
}

func main() {
	r := router.NewRouter(handlerFuncs)

}
