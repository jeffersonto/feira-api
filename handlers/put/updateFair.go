package put

import (
	"feira-api/dto"
	"feira-api/handlers"
	"feira-api/util/commons"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	URLUpdateFair = "/fairs/:fairId"
)

type updateFairHandler struct {
	handlers.Handler
}

func NewUpdateHandler(handler handlers.Handler, r *gin.Engine) {
	handle := updateFairHandler{Handler: handler}
	r.PUT(URLUpdateFair, handle.UpdateFair())
}

func (handler *updateFairHandler) UpdateFair() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			updateFair dto.Fair
		)
		logrus.Tracef("Put UpdateFair Initializing")

		fairID, err := commons.ConvertToInt(strings.TrimSpace(c.Param("fairId")))
		if err != nil {
			_ = c.Error(err)
			return
		}

		err = c.ShouldBindBodyWith(&updateFair, binding.JSON)
		if err != nil {
			_ = c.Error(err)
			return
		}

		err = handler.FairRepository.Update(fairID, updateFair.ToEntity())
		if err != nil {
			_ = c.Error(err)
			return
		}

		logrus.Tracef("Put UpdateFair Finished")
		c.Status(http.StatusNoContent)
	}
}
