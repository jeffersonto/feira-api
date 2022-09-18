package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/internal/service"
)

type Handler struct {
	Service     service.FairService
	RouterGroup *gin.RouterGroup
}

func NewHandler(service service.FairService, routerGroup *gin.RouterGroup) Handler {
	return Handler{Service: service, RouterGroup: routerGroup}
}
