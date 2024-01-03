package errors

import (
	"github.com/watariRyo/tasktree/server/domain/model"
)

type APIError struct {
	Status  int                `json:"status"`
	Code    model.ApiErrorCode `json:"code"`
	Message string             `json:"message"`
}

// errorインタフェースをError()を実装
func (e *APIError) Error() string {
	return string(e.Code) + ": " + e.Message
}

func NewApiError(status int, code model.ApiErrorCode, msg string) *APIError {
	return &APIError{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}
