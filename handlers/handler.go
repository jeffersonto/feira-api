package handlers

import "github.com/jeffersonto/feira-api/adapters/database/repositories/fair"

type Handler struct {
	FairRepository fair.FairRepository
}

func NewHandler(fairRepository fair.FairRepository) Handler {
	return Handler{FairRepository: fairRepository}
}
