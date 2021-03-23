package apierror

import "fmt"

type ApiError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiError(status, code int, message string) error {
	return &ApiError{Status: status, Code: code, Message: message}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Error: %s (%d)", e.Message, e.Code)
}
