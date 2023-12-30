package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) TestMethod(c echo.Context) error {
	return c.JSON(http.StatusOK, h.usecase.TestMethod())
}
