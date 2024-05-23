package impl

import (
	"github.com/AndreeJait/go-utility/response"
	"github.com/AndreeJait/template-service-go/internal/model/auth"
	"github.com/labstack/echo/v4"
)

func (h *handler) Login(c echo.Context) error {
	//var timeNow = h.timeInternal.Now()
	var param auth.LoginParam
	err := c.Bind(&param)
	if err != nil {
		return err
	}

	resp, err := h.ucAuth.Login(c.Request().Context(), param)
	if err != nil {
		return err
	}

	//since := time.Since(timeNow)
	return response.SuccessOK(c, resp, "authenticated")
}
