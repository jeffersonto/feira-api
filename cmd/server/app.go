package server

import (
	"net/http"

	v1 "github.com/jeffersonto/feira-api/internal/handlers/v1"
	"github.com/jeffersonto/feira-api/internal/handlers/v1/delete"
	"github.com/jeffersonto/feira-api/internal/handlers/v1/get"
	"github.com/jeffersonto/feira-api/internal/handlers/v1/post"
	"github.com/jeffersonto/feira-api/internal/handlers/v1/put"

	"github.com/jeffersonto/feira-api/cmd/server/middleware"
	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	config "github.com/jeffersonto/feira-api/internal/config/db"
	"github.com/jeffersonto/feira-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthChecker struct{}

func Run(port string) error {
	dbx, err := config.DB()
	if err != nil {
		logrus.Errorf("error running server: %+v", err)
		panic(err)
	}

	fairRepositoryConnection, err := fair.NewRepository(dbx)
	if err != nil {
		logrus.Errorf("error running server: %+v", err)
		panic(err)
	}

	defer fairRepositoryConnection.DB.Close()

	service := service.NewFairService(fairRepositoryConnection)

	health := HealthChecker{}

	router := gin.Default()
	router.Use(middleware.JSONContentType())
	router.Use(middleware.ErrorHandle())
	router.GET("/ping", health.PingHandler)

	routerGroupV1 := router.Group("/v1")

	handler := v1.NewHandler(service, routerGroupV1)

	get.NewFairByQueryHandler(handler)
	get.NewFairByIDyHandler(handler)
	delete.NewFairByIDyHandler(handler)
	post.NewFairHandler(handler)
	put.NewUpdateHandler(handler)

	return router.Run(":" + port)
}

func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
