package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/infra/firebase"
)

func (server *Server) setFirebase() {
	ctx := context.Background()
	firebaseApp, err := firebase.InitFirebaseApp(ctx, server.cfg.Server)
	if err != nil {
		panic(err)
	}

	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			// クライアントから送られてきた JWT 取得
			authHeader := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(authHeader, "Bearer ", "", 1)

			_, err = firebaseApp.VerifyIDToken(ctx, idToken)
			if err != nil {
				// JWT が無効なら Handler に進まず別処理
				fmt.Printf("error verifying ID token: %v\n", err)
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: "error verifying ID token\n",
				}
			}

			return next(c)
		}
	})
}
