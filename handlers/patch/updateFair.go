package patch

import (
	"feira-api/dto"
	"feira-api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	URLUpdateFair = "/fairs/:fairID"
)

type updateFairHandler struct {
	handlers.Handler
}

func NewUpdateFairHandler(handler handlers.Handler, r *gin.Engine) {
	handle := updateFairHandler{Handler: handler}
	r.PATCH(URLUpdateFair, handle.UpdateFair())
}

func (handler *updateFairHandler) UpdateFair() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			newFair dto.Fair
		)
		logrus.Tracef("Patch UpdateFair Initializing")

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

		logrus.Tracef("Patch UpdateFair Finished")
		c.Status(http.StatusNoContent)
	}
}
