package service

import (
	"github.com/jeffersonto/feira-api/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/dto"
	"github.com/jeffersonto/feira-api/entity"
)

type FairService interface {
	FindFairByID(ID int64) (entity.Fair, error)
	DeleteFairByID(ID int64) error
	SaveFair(newFair dto.Fair) error
	UpdateFairByID(fairID int64, fairToBeUpdated dto.Fair) error
}

type Fair struct {
	repository fair.FairRepository
}

func NewFairService(repository fair.FairRepository) *Fair {
	return &Fair{repository: repository}
}

func (service *Fair) FindFairByID(ID int64) (entity.Fair, error) {
	return service.repository.GetByID(ID)
}

func (service *Fair) DeleteFairByID(ID int64) error {
	return service.repository.DeleteByID(ID)
}

func (service *Fair) SaveFair(newFair dto.Fair) error {
	return service.repository.Save(newFair.ToEntity())
}

func (service *Fair) UpdateFairByID(fairID int64, fairToBeUpdated dto.Fair) error {
	return service.repository.Update(fairID, fairToBeUpdated.ToEntity())
}
