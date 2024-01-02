package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/domain/model"
)

func (server *Server) setFirebase() {
	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// クライアントから送られてきた JWT 取得
			// authHeader := c.Request().Header.Get("Authorization")
			// idToken := strings.Replace(authHeader, "Bearer ", "", 1)
			jwt, err := c.Cookie("jwt")
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: "no cookie",
				}
			}
			idToken := jwt.Value

			decodeToken, err := server.repo.FirebaseApp.VerifyIDToken(server.ctx, idToken)
			if err != nil {
				// JWT が無効なら Handler に進まず別処理
				fmt.Printf("error verifying ID token: %v\n", err)
				return &echo.HTTPError{
					Code:    http.StatusUnauthorized,
					Message: "error verifying ID token",
				}
			}

			// セッション発行
			_, err = c.Cookie("ssid")
			if err != nil {
				ssid := uuid.NewString()
				duration := decodeToken.Expires - time.Now().Unix()
				err = server.repo.RedisClient.SaveSession(ssid, model.SessionData{
					UserID:    decodeToken.UID,
					IDToken:   idToken,
					ExpiredAt: decodeToken.Expires,
				}, time.Duration(duration)*time.Second)
				if err != nil {
					return &echo.HTTPError{
						Code:    http.StatusUnauthorized,
						Message: "could not get session data.",
					}
				}

				cookie := &http.Cookie{
					Name:   "ssid",
					Value:  ssid,
					MaxAge: int(duration),
					Path:   "/",
				}

				c.SetCookie(cookie)
			}

			return next(c)
		}
	})
}
