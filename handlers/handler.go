package handlers

import (
	"github.com/jeffersonto/feira-api/service"
)

type Handler struct {
	Service service.FairService
}

func NewHandler(Service service.FairService) Handler {
	return Handler{Service: Service}
}
