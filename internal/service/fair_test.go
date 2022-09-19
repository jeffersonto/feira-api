package service_test

import (
	"database/sql"
	"errors"
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/entity"
	"github.com/jeffersonto/feira-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repositoryMock struct {
	mock.Mock
}

func TestFindFairByID(t *testing.T) {

	var repository *repositoryMock

	tests := []struct {
		name     string
		input    int64
		warmUP   func(fairID int64)
		expected func(result entity.Fair, err error)
	}{
		{
			name:  "Should successfully query the repository",
			input: 1,
			warmUP: func(fairID int64) {
				repository = new(repositoryMock)
				repository.On("GetByID", fairID).Return(entity.Fair{ID: 1, NomeFeira: "Feira Teste"}, nil)
			},
			expected: func(result entity.Fair, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, result)
				repository.AssertNumberOfCalls(t, "GetByID", 1)
			},
		},
		{
			name:  "Should successfully query the repository, however not find records",
			input: 1,
			warmUP: func(fairID int64) {
				repository = new(repositoryMock)
				repository.On("GetByID", fairID).Return(entity.Fair{}, sql.ErrNoRows)
			},
			expected: func(result entity.Fair, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, entity.Fair{}, result)
				repository.AssertNumberOfCalls(t, "GetByID", 1)
			},
		},
		{
			name:  "Should get a generic error when querying the repository",
			input: 1,
			warmUP: func(fairID int64) {
				repository = new(repositoryMock)
				repository.On("GetByID", fairID).Return(entity.Fair{}, errors.New("internal_server_error"))
			},
			expected: func(result entity.Fair, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, entity.Fair{}, result)
				repository.AssertNumberOfCalls(t, "GetByID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			results, err := service.NewFairService(repository).FindFairByID(tt.input)
			tt.expected(results, err)
		})
	}
}

func TestFindFairByQuery(t *testing.T) {

	var (
		repository *repositoryMock

		input = dto.QueryParameters{
			Distrito:  "IGUATEMI",
			Regiao5:   "Leste",
			NomeFeira: "FEIRA TESTE",
			Bairro:    "JD BOA ESPERANCA",
		}

		filter = entity.Filter{
			Distrito:  "IGUATEMI",
			Regiao5:   "Leste",
			NomeFeira: "FEIRA TESTE",
			Bairro:    "JD BOA ESPERANCA",
		}
	)

	tests := []struct {
		name     string
		input    dto.QueryParameters
		filter   entity.Filter
		warmUP   func(filters entity.Filter)
		expected func(result []entity.Fair, err error)
	}{
		{
			name:   "Should successfully query the repository",
			input:  input,
			filter: filter,
			warmUP: func(filters entity.Filter) {
				repository = new(repositoryMock)
				repository.On("GetByQueryID", filters).Return([]entity.Fair{{ID: 1, NomeFeira: "Feira Teste"}}, nil)
			},
			expected: func(result []entity.Fair, err error) {
				assert.Nil(t, err)
				assert.Len(t, result, 1)
				repository.AssertNumberOfCalls(t, "GetByQueryID", 1)
			},
		},
		{
			name:   "Should successfully query the repository, however not find records",
			input:  input,
			filter: filter,
			warmUP: func(filters entity.Filter) {
				repository = new(repositoryMock)
				repository.On("GetByQueryID", filters).Return([]entity.Fair{}, nil)
			},
			expected: func(result []entity.Fair, err error) {
				assert.Error(t, err)
				assert.Len(t, result, 0)
				repository.AssertNumberOfCalls(t, "GetByQueryID", 1)
			},
		},
		{
			name:   "Should get a generic error when querying the repository",
			input:  input,
			filter: filter,
			warmUP: func(filters entity.Filter) {
				repository = new(repositoryMock)
				repository.On("GetByQueryID", filters).Return([]entity.Fair{}, errors.New("internal_server_error"))
			},
			expected: func(result []entity.Fair, err error) {
				assert.NotNil(t, err)
				assert.Len(t, result, 0)
				repository.AssertNumberOfCalls(t, "GetByQueryID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.filter)
			results, err := service.NewFairService(repository).FindFairByQuery(tt.input)
			tt.expected(results, err)
		})
	}
}

func TestDeleteFairByID(t *testing.T) {

	var repository *repositoryMock

	tests := []struct {
		name     string
		input    int64
		warmUP   func(fairID int64)
		expected func(err error)
	}{
		{
			name:  "Should successfully delete the record",
			input: 1,
			warmUP: func(fairID int64) {
				repository = new(repositoryMock)
				repository.On("DeleteByID", fairID).Return(nil)
			},
			expected: func(err error) {
				assert.Nil(t, err)
				repository.AssertNumberOfCalls(t, "DeleteByID", 1)
			},
		},
		{
			name:  "Should get a generic error when querying the repository",
			input: 1,
			warmUP: func(fairID int64) {
				repository = new(repositoryMock)
				repository.On("DeleteByID", fairID).Return(errors.New("internal_server_error"))
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				repository.AssertNumberOfCalls(t, "DeleteByID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			err := service.NewFairService(repository).DeleteFairByID(tt.input)
			tt.expected(err)
		})
	}
}

func TestSaveFair(t *testing.T) {
	var (
		fairInput = dto.Fair{
			ID:        880,
			Longitude: -46450426,
			Latitude:  -23602582,
		}

		fairEntity = entity.Fair{
			ID:        880,
			Longitude: -46450426,
			Latitude:  -23602582,
		}
	)

	var repository *repositoryMock

	tests := []struct {
		name       string
		input      dto.Fair
		fairEntity entity.Fair
		warmUP     func(fairEntity entity.Fair)
		expected   func(err error)
	}{
		{
			name:       "Should successfully save the record",
			input:      fairInput,
			fairEntity: fairEntity,
			warmUP: func(fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("Save", fairEntity).Return(nil)
			},
			expected: func(err error) {
				assert.Nil(t, err)
				repository.AssertNumberOfCalls(t, "Save", 1)
			},
		},
		{
			name:       "Should get a generic error when saving the repository",
			input:      fairInput,
			fairEntity: fairEntity,
			warmUP: func(fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("Save", fairEntity).Return(errors.New("internal_server_error"))
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				repository.AssertNumberOfCalls(t, "Save", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.fairEntity)
			err := service.NewFairService(repository).SaveFair(tt.input)
			tt.expected(err)
		})
	}
}

func TestUpdateFairByID(t *testing.T) {
	var (
		fairInput = dto.Fair{
			ID:        880,
			Longitude: -46450426,
			Latitude:  -23602582,
		}

		fairEntity = entity.Fair{
			ID:        880,
			Longitude: -46450426,
			Latitude:  -23602582,
		}
	)

	type input struct {
		ID   int64
		fair dto.Fair
	}

	var repository *repositoryMock

	tests := []struct {
		name       string
		input      input
		fairEntity entity.Fair
		warmUP     func(ID int64, fairEntity entity.Fair)
		expected   func(err error)
	}{
		{
			name: "Should successfully update the record",
			input: input{
				ID:   1,
				fair: fairInput,
			},
			fairEntity: fairEntity,
			warmUP: func(ID int64, fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("Update", ID, fairEntity).Return(nil)
				repository.On("AlreadyAnID", ID).Return(true, nil)
			},
			expected: func(err error) {
				assert.Nil(t, err)
				repository.AssertNumberOfCalls(t, "Update", 1)
				repository.AssertNumberOfCalls(t, "AlreadyAnID", 1)
			},
		},
		{
			name: "Should not find record to update and return status 204",
			input: input{
				ID:   1,
				fair: fairInput,
			},
			fairEntity: fairEntity,
			warmUP: func(ID int64, fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("AlreadyAnID", ID).Return(false, nil)
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				repository.AssertNumberOfCalls(t, "Update", 0)
				repository.AssertNumberOfCalls(t, "AlreadyAnID", 1)
			},
		},
		{
			name: "Should return a generic error because the search function returned an error",
			input: input{
				ID:   1,
				fair: fairInput,
			},
			fairEntity: fairEntity,
			warmUP: func(ID int64, fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("AlreadyAnID", ID).Return(false, errors.New("internal_server_error"))
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				repository.AssertNumberOfCalls(t, "Update", 0)
				repository.AssertNumberOfCalls(t, "AlreadyAnID", 1)
			},
		},
		{
			name: "Should get a generic error when updating the repository",
			input: input{
				ID:   1,
				fair: fairInput,
			},
			fairEntity: fairEntity,
			warmUP: func(ID int64, fairEntity entity.Fair) {
				repository = new(repositoryMock)
				repository.On("Update", ID, fairEntity).Return(errors.New("internal_server_error"))
				repository.On("AlreadyAnID", ID).Return(true, nil)
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				repository.AssertNumberOfCalls(t, "Update", 1)
				repository.AssertNumberOfCalls(t, "AlreadyAnID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input.ID, tt.fairEntity)
			err := service.NewFairService(repository).UpdateFairByID(tt.input.ID, tt.input.fair)
			tt.expected(err)
		})
	}
}

func (sm *repositoryMock) GetByID(fairID int64) (entity.Fair, error) {
	args := sm.Called(fairID)
	result, _ := args.Get(0).(entity.Fair)
	return result, args.Error(1)
}

func (sm *repositoryMock) GetByQueryID(filters entity.Filter) ([]entity.Fair, error) {
	args := sm.Called(filters)
	result, _ := args.Get(0).([]entity.Fair)
	return result, args.Error(1)
}

func (sm *repositoryMock) DeleteByID(fairID int64) error {
	args := sm.Called(fairID)
	return args.Error(0)
}

func (sm *repositoryMock) Save(fair entity.Fair) error {
	args := sm.Called(fair)
	return args.Error(0)
}

func (sm *repositoryMock) Update(id int64, fair entity.Fair) error {
	args := sm.Called(id, fair)
	return args.Error(0)
}

func (sm *repositoryMock) AlreadyAnID(userID int64) (bool, error) {
	args := sm.Called(userID)
	result, _ := args.Get(0).(bool)
	return result, args.Error(1)
}
