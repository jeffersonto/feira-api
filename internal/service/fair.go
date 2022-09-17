package service

import (
	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	dto2 "github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
)

type FairService interface {
	FindFairByID(id int64) (entity.Fair, error)
	DeleteFairByID(id int64) error
	SaveFair(newFair dto2.Fair) error
	UpdateFairByID(fairID int64, fairToBeUpdated dto2.Fair) error
	FindFairByQuery(filters dto2.QueryParameters) ([]entity.Fair, error)
}

type Fair struct {
	repository fair.FairRepository
}

func NewFairService(repository fair.FairRepository) *Fair {
	return &Fair{repository: repository}
}

func (service *Fair) FindFairByID(id int64) (entity.Fair, error) {
	return service.repository.GetByID(id)
}

func (service *Fair) FindFairByQuery(filters dto2.QueryParameters) ([]entity.Fair, error) {
	fairs, err := service.repository.GetByQueryID(filters.ToEntity())
	if err != nil {
		return fairs, err
	}

	if len(fairs) == 0 {
		return fairs, exceptions.NewNoContent()
	}

	return service.repository.GetByQueryID(filters.ToEntity())
}

func (service *Fair) DeleteFairByID(id int64) error {
	return service.repository.DeleteByID(id)
}

func (service *Fair) SaveFair(newFair dto2.Fair) error {
	return service.repository.Save(newFair.ToEntity())
}

func (service *Fair) UpdateFairByID(fairID int64, fairToBeUpdated dto2.Fair) error {
	return service.repository.Update(fairID, fairToBeUpdated.ToEntity())
}