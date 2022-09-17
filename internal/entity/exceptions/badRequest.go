package exceptions

import (
	"fmt"
	"net/http"
)

type ErrBadRequest struct {
	Status int
	Err    error
}

func (e *ErrBadRequest) Error() string {
	return fmt.Sprintf("bad request, err: %+v", e.Err)
}

func NewBadRequest(err error) error {
	return &ErrBadRequest{
		Status: http.StatusBadRequest,
		Err:    err,
	}
}

func (e *ErrBadRequest) StatusCode() int {
	return e.Status
}
