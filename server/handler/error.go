package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/domain/model"
	"github.com/watariRyo/tasktree/server/errors"
)

func APIErrorHandler(err error, c echo.Context) {
	var (
		status = http.StatusInternalServerError
		code   = model.ServerError
		msg    string
	)

	if he, ok := err.(*errors.APIError); ok {
		status = he.Status
		code = he.Code
		msg = he.Message
		if msg == "" {
			msg = http.StatusText(status)
		}
	} else {
		msg = http.StatusText(status)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(status)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(status, errors.NewApiError(status, code, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
