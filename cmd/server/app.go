package server

import (
	"net/http"

	"github.com/jeffersonto/feira-api/cmd/server/middleware"
	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	config "github.com/jeffersonto/feira-api/internal/config/db"
	"github.com/jeffersonto/feira-api/internal/handlers"
	"github.com/jeffersonto/feira-api/internal/handlers/delete"
	"github.com/jeffersonto/feira-api/internal/handlers/get"
	"github.com/jeffersonto/feira-api/internal/handlers/post"
	"github.com/jeffersonto/feira-api/internal/handlers/put"
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
