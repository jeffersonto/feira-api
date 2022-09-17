package fair_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jeffersonto/feira-api/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetByID(t *testing.T) {
	var (
		resultsExpected = entity.Fair{
			ID:        1234,
			Longitude: -46450426,
			Latitude:  -23602582,
			NomeFeira: "Feira Teste",
		}
	)
	sqlClient, mock := NewMock()
	query := "SELECT id, longitude, latitude, setor_censitario, area_ponderacao," +
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5, regiao8, nome_feira," +
		" registro, logradouro, numero, bairro, referencia " +
		" FROM fairs " +
		" WHERE id = ? "

	tests := []struct {
		name     string
		input    int64
		warmUP   func(id int64)
		expected func(result entity.Fair, err error)
	}{
		{
			name:  "Should return error generic executing query",
			input: 1,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnError(errors.New("error finding results"))
			},
			expected: func(result entity.Fair, err error) {
				assert.Equal(t, entity.Fair{}, result)
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error finding results", err.Error())
			},
		},
		{
			name:  "Should return the error sql: no rows in result set",
			input: 999,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnError(sql.ErrNoRows)
			},
			expected: func(result entity.Fair, err error) {
				assert.Equal(t, entity.Fair{}, result)
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "not found", err.Error())
			},
		},
		{
			name:  "Should return records correctly",
			input: 10,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "longitude", "latitude", "nome_feira"}).
							AddRow(1234, -46450426, -23602582, "Feira Teste"))
			},
			expected: func(result entity.Fair, err error) {
				assert.Equal(t, resultsExpected, result)
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			result, err := sqlClient.GetByID(tt.input)
			tt.expected(result, err)
		})
	}
}

func TestGetByQueryID(t *testing.T) {
	sqlClient, mock := NewMock()
	queryAllRows := "SELECT id, longitude, latitude, setor_censitario, area_ponderacao," +
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5," +
		" regiao8, nome_feira, registro, logradouro, numero, bairro, referencia " +
		" FROM fairs " +
		" WHERE 1=1 "

	tests := []struct {
		name     string
		input    entity.Filter
		warmUP   func(filters entity.Filter)
		expected func(result []entity.Fair, err error)
	}{
		{
			name:  "Should return error generic executing query",
			input: entity.Filter{},
			warmUP: func(filters entity.Filter) {
				mock.ExpectQuery(regexp.QuoteMeta(queryAllRows)).
					WillReturnError(errors.New("error finding results"))
			},
			expected: func(result []entity.Fair, err error) {
				assert.Len(t, result, 0)
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error finding results", err.Error())
			},
		},
		{
			name:  "Should return error converting row",
			input: entity.Filter{},
			warmUP: func(filters entity.Filter) {
				mock.ExpectQuery(regexp.QuoteMeta(queryAllRows)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "longitude", "latitude", "setor_censitario", "area_ponderacao",
							"codigo_ibge", "distrito", "codigo_subprefeitura", "subprefeitura", "regiao5", "regiao8",
							"nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
							AddRow(nil, -46450426, -23602582, 355030833000022, 3550308005274, "32", "IGUATEMI",
								30, "SAO MATEUS", "Leste", "Leste 2", "JD.BOA ESPERANCA", "5171-3", "RUA IGUPIARA",
								"S/N", "JD BOA ESPERANCA", ""))
			},
			expected: func(result []entity.Fair, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				assert.Len(t, result, 0)
				assert.Equal(t, "sql: Scan error on column index 0, name \"id\": "+
					"converting NULL to int64 is unsupported", err.Error())
			},
		},
		{
			name:  "Should return error converting row",
			input: entity.Filter{},
			warmUP: func(filters entity.Filter) {
				mock.ExpectQuery(regexp.QuoteMeta(queryAllRows)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "longitude", "latitude", "setor_censitario", "area_ponderacao",
							"codigo_ibge", "distrito", "codigo_subprefeitura", "subprefeitura", "regiao5", "regiao8",
							"nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
							AddRow(nil, -46450426, -23602582, 355030833000022, 3550308005274, "32", "IGUATEMI",
								30, "SAO MATEUS", "Leste", "Leste 2", "JD.BOA ESPERANCA", "5171-3", "RUA IGUPIARA",
								"S/N", "JD BOA ESPERANCA", "").
							RowError(0, errors.New("generic error after scan")))
			},
			expected: func(result []entity.Fair, err error) {
				assert.Nil(t, result)
				assert.Error(t, err)
				assert.Len(t, result, 0)
				assert.Equal(t, "generic error after scan", err.Error())
			},
		},
		{
			name:  "Should return rows correctly, without a filter",
			input: entity.Filter{},
			warmUP: func(filters entity.Filter) {
				mock.ExpectQuery(regexp.QuoteMeta(queryAllRows)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "longitude", "latitude", "setor_censitario", "area_ponderacao",
							"codigo_ibge", "distrito", "codigo_subprefeitura", "subprefeitura", "regiao5", "regiao8",
							"nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
							AddRow(880, -46450426, -23602582, 355030833000022, 3550308005274, "32", "IGUATEMI",
								30, "SAO MATEUS", "Leste", "Leste 2", "JD.BOA ESPERANCA", "5171-3", "RUA IGUPIARA",
								"S/N", "JD BOA ESPERANCA", ""))
			},
			expected: func(result []entity.Fair, err error) {
				assert.NotNil(t, result)
				assert.Nil(t, err)
				assert.Len(t, result, 1)
			},
		},
		{
			name: "Should return rows correctly, with filters",
			input: entity.Filter{
				Distrito:  "IGUATEMI",
				Regiao5:   "Leste",
				NomeFeira: "JD.BOA ESPERANCA",
				Bairro:    "JD BOA ESPERANCA",
			},
			warmUP: func(filters entity.Filter) {
				mock.ExpectQuery(regexp.QuoteMeta(queryAllRows)).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "longitude", "latitude", "setor_censitario", "area_ponderacao",
							"codigo_ibge", "distrito", "codigo_subprefeitura", "subprefeitura", "regiao5", "regiao8",
							"nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
							AddRow(880, -46450426, -23602582, 355030833000022, 3550308005274, "32", "IGUATEMI",
								30, "SAO MATEUS", "Leste", "Leste 2", "JD.BOA ESPERANCA", "5171-3", "RUA IGUPIARA",
								"S/N", "JD BOA ESPERANCA", ""))
			},
			expected: func(result []entity.Fair, err error) {
				assert.NotNil(t, result)
				assert.Nil(t, err)
				assert.Len(t, result, 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			result, err := sqlClient.GetByQueryID(tt.input)
			tt.expected(result, err)
		})
	}
}

func NewMock() (*fair.Repository, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	repo, _ := fair.NewRepository(dbx)
	return repo, mock
}
