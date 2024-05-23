package main

import (
	"github.com/AndreeJait/template-service-go/internal/handler/auth"
	authImpl "github.com/AndreeJait/template-service-go/internal/handler/auth/impl"
	"github.com/labstack/echo/v4"
)

type listHandler struct {
	authHandler auth.Handler
}

func (a *Api) initHandler(route *echo.Group) listHandler {
	var list listHandler
	listUc := a.initUseCase()

	var authHandler auth.Handler = authImpl.New(a.log, a.cfg, route, listUc.ucAuth, a.timeLoc)
	list.authHandler = authHandler

	return list
}

func (a *Api) startV1() {
	route := a.e.Group("/api/v1")

	handlers := a.initHandler(route)

	// Register all apis
	handlers.authHandler.RegisterApi()
}
