package post

import (
	"net/http"

	"github.com/jeffersonto/feira-api/util/exceptions"

	"github.com/jeffersonto/feira-api/dto"
	"github.com/jeffersonto/feira-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	urlNewFair = "/fairs"
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
			err = exceptions.NewBadRequest(err)
			_ = c.Error(err)
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
