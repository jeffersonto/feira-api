package get

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
	handle.RouterGroup.GET(urlByID, handle.FairByID)
}

// Feira godoc
// @Summary      Busca uma feira por ID
// @Description  Busca uma feira por ID
// @Tags         Feira
// @Accept       json
// @Param        id   path      int  true  "Feira ID"
// @Success      200   {object}   entity.Fair
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /feiras/{id} [get].
func (handler *fairByIDHandler) FairByID(c *gin.Context) {
	logrus.Tracef("Get FairByID Initializing")

	fairID, err := commons.ConvertToInt(strings.TrimSpace(c.Param("fairId")))
	if err != nil {
		_ = c.Error(err)
		return
	}

	feira, err := handler.Service.FindFairByID(fairID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	logrus.Tracef("Get FairByID Finished")
	c.JSON(http.StatusOK, feira)
}
