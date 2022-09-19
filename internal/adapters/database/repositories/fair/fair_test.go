package fair_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jeffersonto/feira-api/internal/adapters/database/repositories/fair"
	"github.com/jeffersonto/feira-api/internal/entity"
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
		" FROM feiras_livres " +
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
		" FROM feiras_livres " +
		" WHERE 1=1 "

	queryWithFilters := "SELECT id, longitude, latitude, setor_censitario, area_ponderacao," +
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5," +
		" regiao8, nome_feira, registro, logradouro, numero, bairro, referencia " +
		" FROM feiras_livres " +
		" WHERE 1=1 " +
		" AND UPPER(TRIM(distrito)) = UPPER(TRIM(?))" +
		" AND UPPER(TRIM(regiao5)) = UPPER(TRIM(?))" +
		" AND UPPER(TRIM(nome_feira)) = UPPER(TRIM(?))" +
		" AND UPPER(TRIM(bairro)) = UPPER(TRIM(?))"

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
				mock.ExpectQuery(regexp.QuoteMeta(queryWithFilters)).
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

func TestDeleteByID(t *testing.T) {

	sqlClient, mock := NewMock()
	query := "DELETE FROM feiras_livres" +
		" WHERE id = ? "

	tests := []struct {
		name     string
		input    int64
		warmUP   func(id int64)
		expected func(err error)
	}{
		{
			name:  "Should return error generic deleting rows",
			input: 1,
			warmUP: func(id int64) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnError(errors.New("error deleting rows"))
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error deleting rows", err.Error())
			},
		},
		{
			name:  "Should delete correctly",
			input: 1,
			warmUP: func(id int64) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: func(err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			err := sqlClient.DeleteByID(tt.input)
			tt.expected(err)
		})
	}
}

func TestSave(t *testing.T) {
	var (
		fairInput = entity.Fair{
			ID:                  880,
			Longitude:           -46450426,
			Latitude:            -23602582,
			SetorCensitario:     355030833000022,
			AreaPonderacao:      3550308005274,
			CodigoIBGE:          "32",
			Distrito:            "IGUATEMI",
			CodigoSubPrefeitura: 30,
			SubPrefeitura:       "SAO MATEUS",
			Regiao5:             "Leste",
			Regiao8:             "Leste 2",
			NomeFeira:           "JD.BOA ESPERANCA",
			Registro:            "5171-3",
			Logradouro:          "RUA IGUPIARA",
			Numero:              "S/N",
			Bairro:              "JD BOA ESPERANCA",
			Referencia:          "",
		}
	)

	sqlClient, mock := NewMock()
	query := "INSERT INTO feiras_livres (longitude, latitude, setor_censitario, area_ponderacao," +
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5, regiao8, nome_feira," +
		"registro, logradouro, numero, bairro, referencia) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tests := []struct {
		name     string
		input    entity.Fair
		warmUP   func(fair entity.Fair)
		expected func(result int64, err error)
	}{
		{
			name:  "Should return a generic error when trying to insert",
			input: fairInput,
			warmUP: func(fair entity.Fair) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(fair.Longitude, fair.Latitude, fair.SetorCensitario, fair.AreaPonderacao,
						fair.CodigoIBGE, fair.Distrito, fair.CodigoSubPrefeitura, fair.SubPrefeitura,
						fair.Regiao5, fair.Regiao8, fair.NomeFeira, fair.Registro, fair.Logradouro,
						fair.Numero, fair.Bairro, fair.Referencia).
					WillReturnError(errors.New("error saving rows"))
			},
			expected: func(result int64, err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error saving rows", err.Error())
			},
		},
		{
			name:  "Should save correctly",
			input: fairInput,
			warmUP: func(fair entity.Fair) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(fair.Longitude, fair.Latitude, fair.SetorCensitario, fair.AreaPonderacao,
						fair.CodigoIBGE, fair.Distrito, fair.CodigoSubPrefeitura, fair.SubPrefeitura,
						fair.Regiao5, fair.Regiao8, fair.NomeFeira, fair.Registro, fair.Logradouro,
						fair.Numero, fair.Bairro, fair.Referencia).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: func(result int64, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			results, err := sqlClient.Save(tt.input)
			tt.expected(results, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	var (
		fairInput = entity.Fair{
			ID:                  880,
			Longitude:           -46450426,
			Latitude:            -23602582,
			SetorCensitario:     355030833000022,
			AreaPonderacao:      3550308005274,
			CodigoIBGE:          "32",
			Distrito:            "IGUATEMI",
			CodigoSubPrefeitura: 30,
			SubPrefeitura:       "SAO MATEUS",
			Regiao5:             "Leste",
			Regiao8:             "Leste 2",
			NomeFeira:           "JD.BOA ESPERANCA",
			Registro:            "5171-3",
			Logradouro:          "RUA IGUPIARA",
			Numero:              "S/N",
			Bairro:              "JD BOA ESPERANCA",
			Referencia:          "",
		}
	)

	type input struct {
		fairID int64
		fair   entity.Fair
	}

	sqlClient, mock := NewMock()
	query := "UPDATE feiras_livres SET" +
		" longitude = ?," +
		" latitude = ?," +
		" setor_censitario = ?," +
		" area_ponderacao = ?," +
		" codigo_ibge = ?, " +
		" distrito = ?," +
		" codigo_subprefeitura = ?, " +
		" subprefeitura = ?," +
		" regiao5 = ?," +
		" regiao8 = ?," +
		" nome_feira = ?," +
		" registro = ?," +
		" logradouro = ?," +
		" numero = ?," +
		" bairro = ?," +
		" referencia = ?" +
		" WHERE id = ? "

	tests := []struct {
		name     string
		input    input
		warmUP   func(fair input)
		expected func(err error)
	}{
		{
			name: "Should return a generic error when trying to update",
			input: input{
				fairID: 10,
				fair:   fairInput,
			},
			warmUP: func(data input) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(data.fair.Longitude, data.fair.Latitude, data.fair.SetorCensitario, data.fair.AreaPonderacao,
						data.fair.CodigoIBGE, data.fair.Distrito, data.fair.CodigoSubPrefeitura, data.fair.SubPrefeitura,
						data.fair.Regiao5, data.fair.Regiao8, data.fair.NomeFeira, data.fair.Registro, data.fair.Logradouro,
						data.fair.Numero, data.fair.Bairro, data.fair.Referencia, data.fairID).
					WillReturnError(errors.New("error updating rows"))
			},
			expected: func(err error) {
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error updating rows", err.Error())
			},
		},
		{
			name: "Should update correctly",
			input: input{
				fairID: 10,
				fair:   fairInput,
			},
			warmUP: func(data input) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(data.fair.Longitude, data.fair.Latitude, data.fair.SetorCensitario, data.fair.AreaPonderacao,
						data.fair.CodigoIBGE, data.fair.Distrito, data.fair.CodigoSubPrefeitura, data.fair.SubPrefeitura,
						data.fair.Regiao5, data.fair.Regiao8, data.fair.NomeFeira, data.fair.Registro, data.fair.Logradouro,
						data.fair.Numero, data.fair.Bairro, data.fair.Referencia, data.fairID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: func(err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			err := sqlClient.Update(tt.input.fairID, tt.input.fair)
			tt.expected(err)
		})
	}
}

func TestAlreadyAnID(t *testing.T) {
	sqlClient, mock := NewMock()
	query := "SELECT id " +
		" FROM feiras_livres " +
		" WHERE id = ?"

	tests := []struct {
		name     string
		input    int64
		warmUP   func(id int64)
		expected func(result bool, err error)
	}{
		{
			name:  "Should return error generic executing query",
			input: 1,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnError(errors.New("error finding results"))
			},
			expected: func(result bool, err error) {
				assert.Equal(t, false, result)
				assert.NotNil(t, err)
				assert.Error(t, err)
				assert.Equal(t, "error finding results", err.Error())
			},
		},
		{
			name:  "Should return false because it found no records",
			input: 999,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnError(sql.ErrNoRows)
			},
			expected: func(result bool, err error) {
				assert.Equal(t, false, result)
				assert.Nil(t, err)
			},
		},
		{
			name:  "Should return record correctly and return true",
			input: 10,
			warmUP: func(id int64) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(id).
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).
							AddRow(1234))
			},
			expected: func(result bool, err error) {
				assert.Equal(t, true, result)
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.warmUP(tt.input)
			result, err := sqlClient.AlreadyAnID(tt.input)
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
