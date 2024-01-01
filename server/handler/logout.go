package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) Logout(c echo.Context) error {
	return c.JSON(http.StatusNoContent, h.usecase.Logout())
}
