package exceptions

import (
	"fmt"
	"net/http"
)

type ErrUnknown struct {
	Status int
	Err    error
}

func (e *ErrUnknown) Error() string {
	return fmt.Sprintf("unknown error, err: %+v", e.Err)
}

func NewUnknown(err error) error {
	return &ErrUnknown{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func (e *ErrUnknown) StatusCode() int {
	return e.Status
}
