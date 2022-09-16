package post_test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffersonto/feira-api/dto"
	"github.com/jeffersonto/feira-api/handlers"
	"github.com/jeffersonto/feira-api/handlers/post"
	"github.com/jeffersonto/feira-api/server/middleware"
	"github.com/jeffersonto/feira-api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewFair(t *testing.T) {

	var service *serviceMock

	tests := []struct {
		name     string
		body     string
		warmUP   func()
		expected func(result *httptest.ResponseRecorder)
	}{
		{
			name: "Should successfully get and return status code 204",
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
				service.On("SaveFair", mock.Anything).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 201, result.Code)
				service.AssertNumberOfCalls(t, "SaveFair", 1)
			},
		},
		{
			name: "Should miss the longitude field and return status code 400",
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
				service.On("SaveFair", mock.Anything).Return(nil)
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 400, result.Code)
				service.AssertNumberOfCalls(t, "SaveFair", 0)
			},
		},
		{
			name: "Should execute the SaveFair Function, however receive an internal_server_error with status code 500",
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
				service.On("SaveFair", mock.Anything).Return(fmt.Errorf("internal_server_error"))
			},
			expected: func(result *httptest.ResponseRecorder) {
				assert.Equal(t, 500, result.Code)
				service.AssertNumberOfCalls(t, "SaveFair", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP()
			router := gin.Default()
			router.Use(middleware.ErrorHandle())
			handler := handlers.NewHandler(service)
			post.NewFairHandler(handler, router)
			response := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/fairs", bytes.NewBufferString(tt.body))
			router.ServeHTTP(response, req)
			tt.expected(response)
		})
	}
}

type serviceMock struct {
	mock.Mock
	service.FairService
}

func (sm *serviceMock) SaveFair(newFair dto.Fair) error {
	args := sm.Called(newFair)
	return args.Error(0)
}
