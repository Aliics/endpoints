package endpoints

import (
	"net/http"
)

type HandlerError struct {
	StatusCode int
	Message    string
}

func NewHandlerError(code int, message string) HandlerError {
	return HandlerError{code, message}
}

func BadRequestError(message string) HandlerError {
	return NewHandlerError(http.StatusBadRequest, message)
}

func UnauthorizedErr(message string) HandlerError {
	return NewHandlerError(http.StatusUnauthorized, message)
}

func NotFoundError(message string) HandlerError {
	return NewHandlerError(http.StatusNotFound, message)
}

func (h HandlerError) Error() string {
	return h.Message
}
