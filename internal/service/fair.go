package service

import (
	"database/sql"

	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity"
	"github.com/jeffersonto/feira-api/internal/entity/exceptions"
)

type FairService interface {
	FindFairByID(id int64) (entity.Fair, error)
	DeleteFairByID(id int64) error
	SaveFair(newFair dto.Fair) error
	UpdateFairByID(fairID int64, fairToBeUpdated dto.Fair) error
	FindFairByQuery(filters dto.QueryParameters) ([]entity.Fair, error)
}

type Fair struct {
	repository fair.FairRepository
}

func NewFairService(repository fair.FairRepository) *Fair {
	return &Fair{repository: repository}
}

func (service *Fair) FindFairByID(id int64) (entity.Fair, error) {
	fair, err := service.repository.GetByID(id)

	if err == sql.ErrNoRows {
		return fair, exceptions.NewNoContent()
	}

	if err != nil {
		return fair, err
	}

	return fair, nil
}

func (service *Fair) FindFairByQuery(filters dto.QueryParameters) ([]entity.Fair, error) {
	fairs, err := service.repository.GetByQueryID(filters.ToEntity())
	if err != nil {
		return fairs, err
	}

	if len(fairs) == 0 {
		return fairs, exceptions.NewNoContent()
	}

	return fairs, nil
}

func (service *Fair) DeleteFairByID(id int64) error {
	return service.repository.DeleteByID(id)
}

func (service *Fair) SaveFair(newFair dto.Fair) error {
	return service.repository.Save(newFair.ToEntity())
}

func (service *Fair) UpdateFairByID(fairID int64, fairToBeUpdated dto.Fair) error {
	alreadyAnID, err := service.repository.AlreadyAnID(fairID)
	if err != nil {
		return err
	}

	if !alreadyAnID {
		return exceptions.NewNoContent()
	}

	return service.repository.Update(fairID, fairToBeUpdated.ToEntity())
}
