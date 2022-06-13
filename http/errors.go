package http

import (
	"fmt"
)

type ApiError struct {
	HttpStatus int
	Code int
	Message string
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("longbridge openapi error, httpStatus:%d code:%d message:%s", ae.HttpStatus, ae.Code, ae.Message)
}

func NewError(httpStatus int, resp *ApiResponse) error {
	return &ApiError{
		HttpStatus: httpStatus,
		Code: resp.code,
		Message: resp.message,
	}
}