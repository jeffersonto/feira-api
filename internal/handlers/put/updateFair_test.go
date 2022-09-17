package put_test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/cmd/server/middleware"
	"github.com/jeffersonto/feira-api/internal/dto"
	"github.com/jeffersonto/feira-api/internal/handlers"
	"github.com/jeffersonto/feira-api/internal/handlers/put"
	"github.com/jeffersonto/feira-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateFair(t *testing.T) {

	var service *serviceMock

	tests := []struct {
		name          string
		pathParameter string
		body          string
		warmUP        func()
		expected      func(result *httptest.ResponseRecorder)
	}{
		{
			name:          "Should successfully get and return status code 204",
			pathParameter: "1",
			body: `{
						"longitude": -46550164,
						"latitude": -23558733,
						"setor_censitario": 355030885000091,
						"area_ponderacao": 3550308005040,
						"codigo_ibge": "87",
						"distrito": "VILA FORMOSA",
						"codigo_subprefeitura": 26,
						"subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
						"regiao5": "Leste",
						"regiao8": "Leste 1",
						"nome_feira": "VILA FORMOSA",
						"registro": "4041-0",
						"logradouro": "RUA MARAGOJIPE",
						"numero": "S/N",
						"bairro": "VL FORMOSA",
						"referencia": "TV RUA PRETORIA"
				}`,
			warmUP: func() {
				service = new(serviceMock)
				service.On("UpdateFairByID", mock.Anything, mock.Anything).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 204, result.Code)
				service.AssertNumberOfCalls(t, "UpdateFairByID", 1)
			},
		},
		{
			name:          "Should miss the longitude field and return status code 400",
			pathParameter: "1",
			body: `{
						"latitude": -23558733,
						"setor_censitario": 355030885000091,
						"area_ponderacao": 3550308005040,
						"codigo_ibge": "87",
						"distrito": "VILA FORMOSA",
						"codigo_subprefeitura": 26,
						"subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
						"regiao5": "Leste",
						"regiao8": "Leste 1",
						"nome_feira": "VILA FORMOSA",
						"registro": "4041-0",
						"logradouro": "RUA MARAGOJIPE",
						"numero": "S/N",
						"bairro": "VL FORMOSA",
						"referencia": "TV RUA PRETORIA"
				}`,
			warmUP: func() {
				service = new(serviceMock)
				service.On("UpdateFairByID", mock.Anything, mock.Anything).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 400, result.Code)
				service.AssertNumberOfCalls(t, "UpdateFairByID", 0)
			},
		},
		{
			name:          "Should not be able to convert the path parameter and return status code 400",
			pathParameter: "A",
			body: `{
						"longitude": -46550164,
						"latitude": -23558733,
						"setor_censitario": 355030885000091,
						"area_ponderacao": 3550308005040,
						"codigo_ibge": "87",
						"distrito": "VILA FORMOSA",
						"codigo_subprefeitura": 26,
						"subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
						"regiao5": "Leste",
						"regiao8": "Leste 1",
						"nome_feira": "VILA FORMOSA",
						"registro": "4041-0",
						"logradouro": "RUA MARAGOJIPE",
						"numero": "S/N",
						"bairro": "VL FORMOSA",
						"referencia": "TV RUA PRETORIA"
				}`,
			warmUP: func() {
				service = new(serviceMock)
				service.On("UpdateFairByID", mock.Anything, mock.Anything).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 400, result.Code)
				service.AssertNumberOfCalls(t, "UpdateFairByID", 0)
			},
		},
		{
			name:          "Should execute the SaveFair Function, however receive an internal_server_error with status code 500",
			pathParameter: "1",
			body: `{
						"longitude": -46550164,
						"latitude": -23558733,
						"latitude": -23558733,
						"setor_censitario": 355030885000091,
						"area_ponderacao": 3550308005040,
						"codigo_ibge": "87",
						"distrito": "VILA FORMOSA",
						"codigo_subprefeitura": 26,
						"subprefeitura": "ARICANDUVA-FORMOSA-CARRAO",
						"regiao5": "Leste",
						"regiao8": "Leste 1",
						"nome_feira": "VILA FORMOSA",
						"registro": "4041-0",
						"logradouro": "RUA MARAGOJIPE",
						"numero": "S/N",
						"bairro": "VL FORMOSA",
						"referencia": "TV RUA PRETORIA"
				}`,
			warmUP: func() {
				service = new(serviceMock)
				service.On("UpdateFairByID", mock.Anything, mock.Anything).Return(fmt.Errorf("internal_server_error"))
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 500, result.Code)
				service.AssertNumberOfCalls(t, "UpdateFairByID", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP()
			router := gin.Default()
			router.Use(middleware.ErrorHandle())
			handler := handlers.NewHandler(service)
			put.NewUpdateHandler(handler, router)
			response := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/feiras/%v", tt.pathParameter), bytes.NewBufferString(tt.body))
			router.ServeHTTP(response, req)
			tt.expected(response)
		})
	}
}

type serviceMock struct {
	mock.Mock
	service.FairService
}

func (sm *serviceMock) UpdateFairByID(fairID int64, fairToBeUpdated dto.Fair) error {
	args := sm.Called(fairID, fairToBeUpdated)
	return args.Error(0)
}
