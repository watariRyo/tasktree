package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/watariRyo/tasktree/server/usecase"
)

type HandlerImpl struct {
	usecase *usecase.UseCaseImpl
}

type Handler interface {
	Logout(c echo.Context) error
}

var _ Handler = (*HandlerImpl)(nil)

func NewHandler(usecase *usecase.UseCaseImpl) *HandlerImpl {
	return &HandlerImpl{
		usecase: usecase,
	}
}
