package delete

import (
	"net/http"
	"strings"

	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/util/commons"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	urlByID = "/feiras/:fairId"
)

type fairByIDHandler struct {
	handlers.Handler
}

func NewFairByIDyHandler(handler handlers.Handler, r *gin.Engine) {
	handle := fairByIDHandler{Handler: handler}
	r.DELETE(urlByID, handle.FairByID())
}

func (handler *fairByIDHandler) FairByID() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
