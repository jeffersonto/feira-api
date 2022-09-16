package middleware

import (
	"reflect"

	"github.com/jeffersonto/feira-api/util/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	DefaultStatusCodeError = 500
)

type ResponseError struct {
	Err        *gin.Error
	StatusCode int
	ErrType    string
	Success    bool
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		responseError := captureStatusCodeAndError(c)

		if responseError.Success {
			return
		}

		c.JSON(
			responseError.StatusCode,
			gin.H{
				"error":   responseError.ErrType,
				"message": responseError.Err.Error(),
			},
		)
		c.Abort()
	}
}

func captureStatusCodeAndError(c *gin.Context) ResponseError {
	var responseError ResponseError
	err := c.Errors.Last()

	if err == nil {
		responseError = ResponseError{
			nil,
			200,
			"",
			true,
		}
		return responseError
	}

	logrus.Errorf("error=%+v, request=%+v", err.Err, c.Request)

	var statusCode int
	var errType string
	if ex, ok := err.Err.(exceptions.Error); ok {
		statusCode = ex.StatusCode()
		errType = reflect.TypeOf(err.Err).String()
	} else {
		statusCode = DefaultStatusCodeError
		errType = reflect.TypeOf(err).String()
	}

	responseError = ResponseError{
		err,
		statusCode,
		errType,
		false,
	}

	return responseError
}
