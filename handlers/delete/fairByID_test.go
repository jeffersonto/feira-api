package delete_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/handlers/delete"
	"github.com/jeffersonto/feira-api/server/middleware"
	"github.com/jeffersonto/feira-api/service"
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
		input          string
		expectedFairID int64
		warmUP         func(expectedFairID int64)
		expected       func(result *httptest.ResponseRecorder)
	}{
		{
			name:           "Should successfully delete and return status code 204",
			input:          "1",
			expectedFairID: 1,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("DeleteFairByID", expectedFairID).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 204, result.Code)
				service.AssertNumberOfCalls(t, "DeleteFairByID", 1)
			},
		},
		{
			name:           "Should not be able to convert the path parameter and return status code 400",
			input:          "A",
			expectedFairID: 0,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("DeleteFairByID", expectedFairID).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 400, result.Code)
				service.AssertNumberOfCalls(t, "DeleteFairByID", 0)
			},
		},
		{
			name:           "Should execute the DeleteFairByID Function, however receive an internal_server_error with status code 500",
			input:          "1",
			expectedFairID: 1,
			warmUP: func(expectedFairID int64) {
				service = new(serviceMock)
				service.On("DeleteFairByID", expectedFairID).Return(fmt.Errorf("internal_server_error"))
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 500, result.Code)
				service.AssertNumberOfCalls(t, "DeleteFairByID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.expectedFairID)
			router := gin.Default()
			router.Use(middleware.ErrorHandle())
			handler := handlers.NewHandler(service)
			delete.NewFairByIDyHandler(handler, router)
			response := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/fairs/%v", tt.input), nil)
			router.ServeHTTP(response, req)
			tt.expected(response)
		})
	}
}

type serviceMock struct {
	mock.Mock
	service.FairService
}

func (sm *serviceMock) DeleteFairByID(id int64) error {
	args := sm.Called(id)
	return args.Error(0)
}
