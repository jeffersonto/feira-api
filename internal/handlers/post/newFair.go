package post

import (
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
	"github.com/jeffersonto/feira-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	urlNewFair = "/feiras"
)

type newFairHandler struct {
	handlers.Handler
}

func NewFairHandler(handler handlers.Handler, r *gin.Engine) {
	handle := newFairHandler{Handler: handler}
	r.POST(urlNewFair, handle.NewFair())
}

func (handler *newFairHandler) NewFair() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}