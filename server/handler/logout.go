package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerImpl) Logout(c echo.Context) error {
	// セッション削除
	// セッションがない場合も正常で返す
	ssidCookie, err := c.Cookie("ssid")
	if err != nil {
		// セッションがなくても
		return c.JSON(http.StatusNoContent, nil)
	}

	return c.JSON(http.StatusNoContent, h.usecase.Logout(ssidCookie.Value))
}
