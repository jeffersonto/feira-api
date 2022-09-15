package get

import (
	"feira-api/handlers"
	"feira-api/util/commons"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	URLByID = "/fairs/:fairId"
)

type fairByIDHandler struct {
	handlers.Handler
}

func NewFairByIDyHandler(handler handlers.Handler, r *gin.Engine) {
	handle := fairByIDHandler{Handler: handler}
	r.GET(URLByID, handle.FairByID())
}

func (handler *fairByIDHandler) FairByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Tracef("Get FairByID Initializing")

		fairID, err := commons.ConvertToInt(strings.TrimSpace(c.Param("fairId")))
		if err != nil {
			_ = c.Error(err)
			return
		}

		feira, err := handler.FairRepository.GetByID(fairID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		logrus.Tracef("Get FairByID Finished")
		c.JSON(http.StatusOK, feira)
	}
}
