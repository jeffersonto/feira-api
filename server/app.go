package server

import (
	"net/http"

	"github.com/jeffersonto/feira-api/service"

	"github.com/jeffersonto/feira-api/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/handlers/delete"
	"github.com/jeffersonto/feira-api/handlers/get"
	"github.com/jeffersonto/feira-api/handlers/post"
	"github.com/jeffersonto/feira-api/handlers/put"
	"github.com/jeffersonto/feira-api/server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthChecker struct{}

func Run(port string) error {
	fairRepositoryConnection, err := fair.NewRepository()
	if err != nil {
		logrus.Errorf("error running server: %+v", err)
		panic(err)
	}

	defer fairRepositoryConnection.DB.Close()

	service := service.NewFairService(fairRepositoryConnection)
	handler := handlers.NewHandler(service)

	health := HealthChecker{}

	router := gin.Default()
	router.Use(middleware.JSONContentType())
	router.Use(middleware.ErrorHandle())
	router.GET("/ping", health.PingHandler)

	get.NewFairByQueryHandler(handler, router)
	get.NewFairByIDyHandler(handler, router)
	delete.NewFairByIDyHandler(handler, router)
	post.NewFairHandler(handler, router)
	put.NewUpdateHandler(handler, router)

	return router.Run(":" + port)
}

func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
