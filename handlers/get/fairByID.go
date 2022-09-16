package get

import (
	"net/http"
	"strings"

	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/util/commons"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	urlByID = "/fairs/:fairId"
)

type fairByIDHandler struct {
	handlers.Handler
}

func NewFairByIDyHandler(handler handlers.Handler, r *gin.Engine) {
	handle := fairByIDHandler{Handler: handler}
	r.GET(urlByID, handle.FairByID())
}

func (handler *fairByIDHandler) FairByID() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
