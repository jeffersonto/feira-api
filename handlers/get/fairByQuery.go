package get

import (
	"net/http"

	"github.com/jeffersonto/feira-api/entity/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/dto"
	"github.com/jeffersonto/feira-api/handlers"
	"github.com/sirupsen/logrus"
)

const (
	URLByQuery = "/feiras"
)

type fairByQueryHandler struct {
	handlers.Handler
}

func NewFairByQueryHandler(handler handlers.Handler, r *gin.Engine) {
	handle := fairByQueryHandler{Handler: handler}
	r.GET(URLByQuery, handle.FairByQuery())
}

func (handler *fairByQueryHandler) FairByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			queryParameters dto.QueryParameters
		)
		logrus.Tracef("Get FairByQuery Initializing")

		logrus.Infof("query=%+v", c.Request.URL.Query())
		if err := c.ShouldBindQuery(&queryParameters); err != nil {
			err = exceptions.NewBadRequest(err)
			_ = c.Error(err)
			return
		}

		feira, err := handler.Service.FindFairByQuery(queryParameters)
		if err != nil {
			_ = c.Error(err)
			return
		}

		logrus.Tracef("Get FairByQuery Finished")
		c.JSON(http.StatusOK, feira)
	}
}
