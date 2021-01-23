package base

import "fmt"

type HttpError struct {
	code    int
	message string
}

func (h HttpError) Error() string {
	return fmt.Sprintf("%d %s", h.code, h.message)
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{code: code, message: message}
}
