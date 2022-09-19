package delete

import (
	"net/http"
	"strings"

	v1 "github.com/jeffersonto/feira-api/internal/handlers/v1"

	"github.com/jeffersonto/feira-api/pkg/commons"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	urlByID = "/feiras/:fairId"
)

type fairByIDHandler struct {
	v1.Handler
}

func NewFairByIDyHandler(handler v1.Handler) {
	handle := fairByIDHandler{Handler: handler}
	handle.RouterGroup.DELETE(urlByID, handle.FairByID)
}

// Feira godoc
// @Summary      Deleta uma feira por ID
// @Description  Deleta uma feira por ID
// @Tags         Feira
// @Accept       json
// @Param        id   path      int  true  "Feira ID"
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /feiras/{id} [delete].
func (handler *fairByIDHandler) FairByID(c *gin.Context) {
	logrus.Tracef("Delete FairByID Initializing")

	fairID, err := commons.ConvertToInt(strings.TrimSpace(c.Param("fairId")))
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = handler.Service.DeleteFairByID(fairID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	logrus.Tracef("Delete FairByID Finished")
	c.Status(http.StatusNoContent)
}
