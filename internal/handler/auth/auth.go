package auth

import "github.com/labstack/echo/v4"

type Handler interface {
	RegisterApi()

	Login(c echo.Context) error
}
