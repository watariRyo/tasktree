package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/auth"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/domain/model"
	"github.com/watariRyo/tasktree/server/errors"
)

const TOKEN_EXPIRED_ERR_MSG = "ID token has expired"

func (server *Server) setFirebase() {
	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// クライアントから送られてきた JWT 取得
			// authHeader := c.Request().Header.Get("Authorization")
			// idToken := strings.Replace(authHeader, "Bearer ", "", 1)
			jwt, err := c.Cookie("jwt")
			if err != nil {
				fmt.Println("user not authenticated")
				return &errors.APIError{
					Status: http.StatusUnauthorized,
				}
			}
			idToken := jwt.Value

			decodeToken, err := server.repo.FirebaseApp.VerifyIDToken(server.ctx, idToken)
			if err != nil {
				// JWT が無効なら Handler に進まず別処理
				fmt.Printf("error verifying ID token: %v\n", err)
				if strings.Contains(err.Error(), TOKEN_EXPIRED_ERR_MSG) {
					return &errors.APIError{
						Status:  http.StatusUnauthorized,
						Code:    model.TokenExpired,
						Message: "error verifying ID token",
					}
				} else {
					return &errors.APIError{
						Status:  http.StatusUnauthorized,
						Code:    model.InvalidToken,
						Message: "error verifying ID token",
					}
				}
			}

			// セッション発行
			ssidCookie, err := c.Cookie("ssid")
			if err != nil {
				ssid := uuid.NewString()
				duration, err := server.setRedisSession(ssid, idToken, decodeToken)
				if err != nil {
					return &errors.APIError{
						Status:  http.StatusInternalServerError,
						Code:    model.ServerError,
						Message: "Could not set session data.",
					}
				}

				cookie := &http.Cookie{
					Name:   "ssid",
					Value:  ssid,
					MaxAge: int(duration),
					Path:   "/",
				}

				c.SetCookie(cookie)
			} else {
				ssid := ssidCookie.Value
				sessionData, err := server.repo.RedisClient.GetSession(ssid)
				if err != nil {
					return &errors.APIError{
						Status:  http.StatusUnauthorized,
						Message: "Could not get session data.",
					}
				}
				// idTokenが更新されていた場合、セッションデータを更新
				if sessionData.IDToken != idToken {
					_, err = server.setRedisSession(ssid, idToken, decodeToken)
					if err != nil {
						return &errors.APIError{
							Status:  http.StatusInternalServerError,
							Code:    model.ServerError,
							Message: "Could not set session data.",
						}
					}
				}
			}

			return next(c)
		}
	})
}

func (server *Server) setRedisSession(ssid string, idToken string, decodeToken *auth.Token) (int64, error) {
	duration := decodeToken.Expires - time.Now().Unix()
	user, err := server.repo.UsersRepository.GetUserByUUID(server.ctx, server.repo.DBConnection, decodeToken.UID)

	if err != nil {
		// UserIDで取れない場合は登録
		userExist := server.repo.UsersRepository.Exist(server.ctx, server.repo.DBConnection, decodeToken.UID)
		if !userExist {
			displayDefault := uuid.New().String()
			user, err = server.repo.UsersRepository.Insert(server.ctx, server.repo.DBConnection, &model.User{
				UUID:        decodeToken.UID,
				DisplayName: &displayDefault,
			})
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	err = server.repo.RedisClient.SaveSession(ssid, model.SessionData{
		// UserID:    decodeToken.UID,
		UserID:    user.ID,
		IDToken:   idToken,
		ExpiredAt: decodeToken.Expires,
	}, time.Duration(duration)*time.Second)
	if err != nil {
		return 0, err
	}
	return duration, nil
}
