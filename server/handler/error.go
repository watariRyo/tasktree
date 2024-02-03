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

	if ae, ok := err.(*errors.APIError); ok {
		status = ae.Status
		code = ae.Code
		msg = ae.Message
		if msg == "" {
			msg = http.StatusText(status)
		}
	} else if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		code = ""
		msg = he.Message.(string)
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
