package put

import (
	"net/http"
	"strings"

	"github.com/jeffersonto/feira-api/entity/exceptions"

	"github.com/jeffersonto/feira-api/dto"
	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/util/commons"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	urlUpdateFair = "/feiras/:fairId"
)

type updateFairHandler struct {
	handlers.Handler
}

func NewUpdateHandler(handler handlers.Handler, r *gin.Engine) {
	handle := updateFairHandler{Handler: handler}
	r.PUT(urlUpdateFair, handle.UpdateFair())
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
			_ = c.Error(exceptions.NewBadRequest(err))
			return
		}

		err = handler.Service.UpdateFairByID(fairID, updateFair)
		if err != nil {
			_ = c.Error(err)
			return
		}

		logrus.Tracef("Put UpdateFair Finished")
		c.Status(http.StatusNoContent)
	}
}
