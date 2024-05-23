package impl

import (
	"github.com/AndreeJait/go-utility/loggerw"
	"github.com/AndreeJait/go-utility/timew"
	"github.com/AndreeJait/template-service-go/config"
	"github.com/AndreeJait/template-service-go/internal/handler/auth"
	ucAuth "github.com/AndreeJait/template-service-go/internal/usecase/auth"
	"github.com/labstack/echo/v4"
)

type handler struct {
	log          loggerw.Logger
	cfg          *config.Config
	route        *echo.Group
	ucAuth       ucAuth.UseCase
	timeInternal timew.Time
}

func New(log loggerw.Logger, cfg *config.Config, route *echo.Group, ucAuth ucAuth.UseCase, timeInternal timew.Time) auth.Handler {

	return &handler{
		log:          log,
		cfg:          cfg,
		route:        route,
		ucAuth:       ucAuth,
		timeInternal: timeInternal,
	}
}
