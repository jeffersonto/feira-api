package put

import (
	"net/http"
	"strings"

	v1 "github.com/jeffersonto/feira-api/internal/handlers/v1"

	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
	"github.com/jeffersonto/feira-api/pkg/commons"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	urlUpdateFair = "/feiras/:fairId"
)

type updateFairHandler struct {
	v1.Handler
}

func NewUpdateHandler(handler v1.Handler) {
	handle := updateFairHandler{Handler: handler}
	handle.RouterGroup.PUT(urlUpdateFair, handle.UpdateFair)
}

func (handler *updateFairHandler) UpdateFair(c *gin.Context) {
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
	c.Status(http.StatusOK)
}
