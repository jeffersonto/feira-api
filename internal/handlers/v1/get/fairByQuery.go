package get

import (
	"net/http"

	v1 "github.com/jeffersonto/feira-api/internal/handlers/v1"

	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
	"github.com/sirupsen/logrus"
)

const (
	URLByQuery = "/feiras"
)

type fairByQueryHandler struct {
	v1.Handler
}

func NewFairByQueryHandler(handler v1.Handler) {
	handle := fairByQueryHandler{Handler: handler}
	handle.RouterGroup.GET(URLByQuery, handle.FairByQuery)
}

// Feira godoc
// @Summary      Busca feiras Por Query Params
// @Description  Busca feiras Por Query Params
// @Tags         Feira
// @Accept       json
// @Param        distrito    query     string  false  "busca feira por distrito"
// @Param        regiao5     query     string  false  "busca feira por regiao5"
// @Param        nomeFeira   query     string  false  "busca feira por nome da feira"
// @Param        bairro     query     string  false  "busca feira por bairro"
// @Success      200   {object}   []entity.Fair
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /feiras [get].
func (handler *fairByQueryHandler) FairByQuery(c *gin.Context) {
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
