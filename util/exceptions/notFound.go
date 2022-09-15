package exceptions

import (
	"fmt"
	"net/http"
)

type ErrNotFound struct {
	Status int
	Err    error
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("not found error, err: %+v", e.Err)
}

func NewNotFound(err error) error {
	return &ErrNotFound{
		Status: http.StatusNotFound,
		Err:    err,
	}
}

func (e *ErrNotFound) StatusCode() int {
	return e.Status
}
