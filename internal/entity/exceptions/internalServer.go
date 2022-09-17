package exceptions

import (
	"fmt"
	"net/http"
)

type ErrInternalServer struct {
	Status int
	Err    error
}

func (e *ErrInternalServer) Error() string {
	return fmt.Sprintf("internal server error, err: %+v", e.Err)
}

func NewInternalServer(err error) error {
	return &ErrInternalServer{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func (e *ErrInternalServer) StatusCode() int {
	return e.Status
}
