package get_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/cmd/server/middleware"
	"github.com/jeffersonto/feira-api/internal/entity"
	"github.com/jeffersonto/feira-api/internal/handlers/v1"
	"github.com/jeffersonto/feira-api/internal/handlers/v1/get"
	"github.com/jeffersonto/feira-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFairByID(t *testing.T) {

	var service *serviceMock

	tests := []struct {
		name           string
		pathParameter  string
		expectedFairID int64
		warmUP         func(expectedFairID int64)
		expected       func(result *httptest.ResponseRecorder)
	}{
		{
			name:           "Should successfully get by id and return status code 200",
			pathParameter:  "1",
			expectedFairID: 1,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("FindFairByID", mock.Anything).Return(entity.Fair{ID: 1, NomeFeira: "Feira Teste"}, nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 200, result.Code)
				service.AssertNumberOfCalls(t, "FindFairByID", 1)
			},
		},
		{
			name:           "Should not be able to convert the path parameter and return status code 400",
			pathParameter:  "A",
			expectedFairID: 0,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("FindFairByID", mock.Anything).Return(entity.Fair{}, nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 400, result.Code)
				service.AssertNumberOfCalls(t, "FindFairByID", 0)
			},
		},
		{
			name:           "Should execute the FindFairByID Function, however receive an internal_server_error with status code 500",
			pathParameter:  "1",
			expectedFairID: 1,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("FindFairByID", mock.Anything).Return(entity.Fair{}, fmt.Errorf("internal_server_error"))
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 500, result.Code)
				service.AssertNumberOfCalls(t, "FindFairByID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.expectedFairID)
			router := gin.Default()
			router.Use(middleware.ErrorHandle())
			routerGroupV1 := router.Group("/v1")
			handler := v1.NewHandler(service, routerGroupV1)
			get.NewFairByIDyHandler(handler)
			response := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/v1/feiras/%v", tt.pathParameter), nil)
			router.ServeHTTP(response, req)
			tt.expected(response)
		})
	}
}

type serviceMock struct {
	mock.Mock
	service.FairService
}

func (sm *serviceMock) FindFairByID(id int64) (entity.Fair, error) {
	args := sm.Called(id)
	result, _ := args.Get(0).(entity.Fair)
	return result, args.Error(1)
}
