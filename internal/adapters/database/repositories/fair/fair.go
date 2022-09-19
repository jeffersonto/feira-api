package fair

import (
	"database/sql"

	"github.com/jeffersonto/feira-api/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import driver for sqlite connection
)

type FairRepository interface {
	GetByID(fairID int64) (entity.Fair, error)
	DeleteByID(fairID int64) error
	Save(fair entity.Fair) error
	Update(id int64, fair entity.Fair) error
	GetByQueryID(filters entity.Filter) ([]entity.Fair, error)
	AlreadyAnID(userID int64) (bool, error)
}

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{
		DB: db,
	}, nil
}

func (repo *Repository) GetByID(fairID int64) (entity.Fair, error) {
	var (
		fair = entity.Fair{}
	)
	err := repo.DB.Get(&fair, "SELECT id, longitude, latitude, setor_censitario, area_ponderacao,"+
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5, regiao8, nome_feira,"+
		" registro, logradouro, numero, bairro, referencia "+
		" FROM feiras_livres "+
		" WHERE id = ? ", fairID)

	if err != nil {
		return fair, err
	}

	return fair, nil
}

func (repo *Repository) GetByQueryID(filters entity.Filter) ([]entity.Fair, error) {
	parametersForQuery := make([]interface{}, 0)

	query := "SELECT id, longitude, latitude, setor_censitario, area_ponderacao," +
		" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5," +
		" regiao8, nome_feira, registro, logradouro, numero, bairro, referencia " +
		" FROM feiras_livres " +
		" WHERE 1=1 "

	if filters.Distrito != "" {
		query += " AND UPPER(TRIM(distrito)) = UPPER(TRIM(?))"
		parametersForQuery = append(parametersForQuery, filters.Distrito)
	}

	if filters.Regiao5 != "" {
		query += " AND UPPER(TRIM(regiao5)) = UPPER(TRIM(?))"
		parametersForQuery = append(parametersForQuery, filters.Regiao5)
	}

	if filters.NomeFeira != "" {
		query += " AND UPPER(TRIM(nome_feira)) = UPPER(TRIM(?))"
		parametersForQuery = append(parametersForQuery, filters.NomeFeira)
	}

	if filters.Bairro != "" {
		query += " AND UPPER(TRIM(bairro)) = UPPER(TRIM(?))"
		parametersForQuery = append(parametersForQuery, filters.Bairro)
	}

	fairs := make([]entity.Fair, 0)

	var rows *sql.Rows

	rows, err := repo.DB.Query(query, parametersForQuery...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data entity.Fair
		err = rows.Scan(&data.ID, &data.Longitude, &data.Latitude, &data.SetorCensitario, &data.AreaPonderacao,
			&data.CodigoIBGE, &data.Distrito, &data.CodigoSubPrefeitura, &data.SubPrefeitura, &data.Regiao5,
			&data.Regiao8, &data.NomeFeira, &data.Registro, &data.Logradouro, &data.Numero, &data.Bairro, &data.Referencia)
		if err != nil {
			return nil, err
		}
		fairs = append(fairs, data)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return fairs, nil
}

func (repo *Repository) DeleteByID(fairID int64) error {
	_, err := repo.DB.Exec(
		"DELETE FROM feiras_livres"+
			" WHERE id = ? ", fairID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Save(fair entity.Fair) error {
	_, err := repo.DB.Exec(
		"INSERT INTO feiras_livres (longitude, latitude, setor_censitario, area_ponderacao,"+
			" codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5, regiao8, nome_feira,"+
			"registro, logradouro, numero, bairro, referencia) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		fair.Longitude, fair.Latitude, fair.SetorCensitario,
		fair.AreaPonderacao, fair.CodigoIBGE, fair.Distrito,
		fair.CodigoSubPrefeitura, fair.SubPrefeitura,
		fair.Regiao5, fair.Regiao8, fair.NomeFeira,
		fair.Registro, fair.Logradouro, fair.Numero,
		fair.Bairro, fair.Referencia)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Update(id int64, fair entity.Fair) error {
	_, err := repo.DB.Exec(
		"UPDATE feiras_livres SET"+
			" longitude = ?,"+
			" latitude = ?,"+
			" setor_censitario = ?,"+
			" area_ponderacao = ?,"+
			" codigo_ibge = ?, "+
			" distrito = ?,"+
			" codigo_subprefeitura = ?, "+
			" subprefeitura = ?,"+
			" regiao5 = ?,"+
			" regiao8 = ?,"+
			" nome_feira = ?,"+
			" registro = ?,"+
			" logradouro = ?,"+
			" numero = ?,"+
			" bairro = ?,"+
			" referencia = ?"+
			" WHERE id = ? ",
		fair.Longitude, fair.Latitude, fair.SetorCensitario,
		fair.AreaPonderacao, fair.CodigoIBGE, fair.Distrito,
		fair.CodigoSubPrefeitura, fair.SubPrefeitura,
		fair.Regiao5, fair.Regiao8, fair.NomeFeira,
		fair.Registro, fair.Logradouro, fair.Numero,
		fair.Bairro, fair.Referencia, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) AlreadyAnID(userID int64) (bool, error) {
	var idFetched int64
	err := repo.DB.QueryRow("SELECT id "+
		" FROM feiras_livres "+
		" WHERE id = ?", userID).Scan(&idFetched)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}
