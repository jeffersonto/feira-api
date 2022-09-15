package post

import (
	"feira-api/dto"
	"feira-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	URLNewFair = "/fairs"
)

type newFairHandler struct {
	handlers.Handler
}

func NewFairHandler(handler handlers.Handler, r *gin.Engine) {
	handle := newFairHandler{Handler: handler}
	r.POST(URLNewFair, handle.NewFair())
}

func (handler *newFairHandler) NewFair() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			newFair dto.Fair
		)
		logrus.Tracef("Post NewFair Initializing")

		err := c.ShouldBindBodyWith(&newFair, binding.JSON)

		if err != nil {
			_ = c.Error(err)
			return
		}

		err = handler.FairRepository.Save(newFair.ToEntity())
		if err != nil {
			_ = c.Error(err)
			return
		}

		logrus.Tracef("Post NewFair Finished")
		c.Status(http.StatusCreated)
	}
}
