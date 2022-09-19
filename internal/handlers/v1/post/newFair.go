package post

import (
	"net/http"

	v1 "github.com/jeffersonto/feira-api/internal/handlers/v1"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
	"github.com/sirupsen/logrus"
)

const (
	urlNewFair = "/feiras"
)

type newFairHandler struct {
	v1.Handler
}

func NewFairHandler(handler v1.Handler) {
	handle := newFairHandler{Handler: handler}
	handle.RouterGroup.POST(urlNewFair, handle.NewFair)
}

func (handler *newFairHandler) NewFair(c *gin.Context) {
	var (
		newFair dto.Fair
	)
	logrus.Tracef("Post NewFair Initializing")

	err := c.ShouldBindBodyWith(&newFair, binding.JSON)

	if err != nil {
		_ = c.Error(exceptions.NewBadRequest(err))
		return
	}

	err = handler.Service.SaveFair(newFair)
	if err != nil {
		_ = c.Error(err)
		return
	}

	logrus.Tracef("Post NewFair Finished")
	c.Status(http.StatusCreated)
}
