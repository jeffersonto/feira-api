package server

import (
	middleware2 "github.com/jeffersonto/feira-api/cmd/server/middleware"
	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/internal/config/db"
	"github.com/jeffersonto/feira-api/internal/handlers"
	delete2 "github.com/jeffersonto/feira-api/internal/handlers/delete"
	get2 "github.com/jeffersonto/feira-api/internal/handlers/get"
	"github.com/jeffersonto/feira-api/internal/handlers/post"
	"github.com/jeffersonto/feira-api/internal/handlers/put"
	"github.com/jeffersonto/feira-api/internal/service"
	"net/http"

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
	router.Use(middleware2.JSONContentType())
	router.Use(middleware2.ErrorHandle())
	router.GET("/ping", health.PingHandler)

	get2.NewFairByQueryHandler(handler, router)
	get2.NewFairByIDyHandler(handler, router)
	delete2.NewFairByIDyHandler(handler, router)
	post.NewFairHandler(handler, router)
	put.NewUpdateHandler(handler, router)

	return router.Run(":" + port)
}

func (h HealthChecker) PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
