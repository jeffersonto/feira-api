package get

import (
	"net/http"
	"strings"

	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/util/commons"
	"github.com/jeffersonto/feira-api/util/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	URLByQuery = "/fairs"
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
		logrus.Tracef("Get FairByQuery Initializing")

		var queryParameters struct {
			Distrito  string `form:"distrito"`
			Regiao5   string `form:"regiao5"`
			NomeFeira string `form:"nomeFeira"`
			Bairro    string `form:"bairro"`
		}

		logrus.Infof("query=%+v", c.Request.URL.Query())
		if err := c.ShouldBindQuery(&queryParameters); err != nil {
			err = exceptions.NewBadRequest(err)
			_ = c.Error(err)
			return
		}

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

		logrus.Tracef("Get FairByQuery Finished")
		c.JSON(http.StatusOK, feira)
	}
}
