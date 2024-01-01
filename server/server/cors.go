package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (server *Server) setCors() {
	var corsOrigins []string
	if len(server.cfg.Server.CorsOrigins) != 0 {
		corsOrigins = strings.Split(server.cfg.Server.CorsOrigins, ",")
	}
	server.e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     corsOrigins,
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}),
	)
	// corsで403を返却させる
	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Originヘッダの中身を取得
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			// 許可しているOriginの中でリクエストヘッダのOriginと一致すれば処理継続
			for _, o := range corsOrigins {
				if origin == o {
					return next(c)
				}
			}
			// 一致しない場合は403
			return &echo.HTTPError{
				Code:    http.StatusForbidden,
				Message: "cors forbidden",
			}
		}
	})
}
