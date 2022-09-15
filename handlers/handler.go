package handlers

import "feira-api/adapters/database/repositories/fair"

type Handler struct {
	FairRepository fair.FairRepository
}

func NewHandler(fairRepository fair.FairRepository) Handler {
	return Handler{FairRepository: fairRepository}
}
