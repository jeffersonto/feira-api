package exceptions

import (
	"net/http"
)

type NoContent struct {
	Status int
}

func (e *NoContent) Error() string {
	return "not found"
}

func NewNoContent() error {
	return &NoContent{
		Status: http.StatusNoContent,
	}
}

func (e *NoContent) StatusCode() int {
	return e.Status
}
