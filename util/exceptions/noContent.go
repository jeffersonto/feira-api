package exceptions

import (
	"fmt"
	"net/http"
)

type ErrNoContent struct {
	Status int
	Err    error
}

func (e *ErrNoContent) Error() string {
	return fmt.Sprintf("not found error, err: %+v", e.Err)
}

func NewNoContent(err error) error {
	return &ErrNoContent{
		Status: http.StatusNoContent,
		Err:    err,
	}
}

func (e *ErrNoContent) StatusCode() int {
	return e.Status
}
