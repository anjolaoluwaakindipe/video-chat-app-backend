package errs

import "net/http"

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func NewContentNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}

func NewConflictError (message string) *AppError{
	return &AppError{Message: message, Code: http.StatusConflict}
}
