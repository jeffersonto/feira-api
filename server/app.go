package server

import (
	"feira-api/adapters/database/repositories/fair"
	"feira-api/handlers"
	"feira-api/handlers/delete"
	"feira-api/handlers/get"
	"feira-api/handlers/post"
	"feira-api/handlers/put"
	"feira-api/server/middleware"
	"net/http"

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

	handler := handlers.NewHandler(fairRepositoryConnection)

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
