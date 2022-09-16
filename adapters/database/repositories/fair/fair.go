package fair

import (
	"database/sql"

	"github.com/jeffersonto/feira-api/entity"
	"github.com/jeffersonto/feira-api/util/exceptions"
	"github.com/jeffersonto/feira-api/util/warmupdb"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import driver for sqlite connection
)

type FairRepository interface {
	GetByID(fairID int64) (entity.Fair, error)
	DeleteByID(fairID int64) error
	Save(fair entity.Fair) error
	Update(id int64, fair entity.Fair) error
}

type Repository struct {
	DB *sqlx.DB
}

func NewRepository() (*Repository, error) {
	db, err := sqlx.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(-1)

	if err != nil {
		return nil, err
	}

	err = warmupdb.WarmUp(db)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: db,
	}, nil
}

func (repo *Repository) GetByID(fairID int64) (entity.Fair, error) {
	var (
		fair = entity.Fair{}
	)
	err := repo.DB.Get(&fair, "SELECT * "+
		"FROM fairs "+
		"WHERE id = ? ", fairID)
	switch {
	case err == sql.ErrNoRows:
		return fair, exceptions.NewNoContent(err)
	case err != nil:
		return fair, err
	default:
		return fair, nil
	}
}

func (repo *Repository) DeleteByID(fairID int64) error {
	_, err := repo.DB.Exec(
		"DELETE FROM fairs"+
			" WHERE id = ? ", fairID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Save(fair entity.Fair) error {
	_, err := repo.DB.Exec(
		"INSERT INTO fairs (longitude, latitude, setor_censitario, area_ponderacao,"+
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
		"UPDATE fairs SET"+
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
