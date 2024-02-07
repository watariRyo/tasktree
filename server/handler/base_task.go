package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/errors"
	"github.com/watariRyo/tasktree/server/logger"
)

func (h *HandlerImpl) GetBaseTask(c echo.Context) error {
	// セッションがない場合も正常で返す
	ssidCookie, err := c.Cookie("ssid")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, nil)
	}
	baseTask, err := h.usecase.GetBaseTask(c.Request().Context(), ssidCookie.Value)
	if err != nil {
		logger.Errorf(c.Request().Context(), "usecase err: %s", err.Error())
		return &errors.APIError{
			Status:  http.StatusInternalServerError,
			Code:    "",
			Message: "usecase error",
		}
	}
	return c.JSON(http.StatusOK, baseTask)
}
