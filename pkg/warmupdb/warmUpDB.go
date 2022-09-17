package warmupdb

import (
	"fmt"
	"os"
	"strings"

	"github.com/jeffersonto/feira-api/internal/entity"

	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
)

func WarmUp(db *sqlx.DB) error {
	err := CreateTablesInMemory(db)
	if err != nil {
		return err
	}

	err = InsertData(db)
	if err != nil {
		return err
	}

	return nil
}

func CreateTablesInMemory(db *sqlx.DB) error {
	table := "CREATE TABLE feiras_livres " +
		"  ( " +
		"     id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, " +
		"     longitude INTEGER NOT NULL, " +
		"     latitude INTEGER NOT NULL, " +
		"     setor_censitario BIGINT NOT NULL, " +
		"     area_ponderacao BIGINT NOT NULL, " +
		"     codigo_ibge VARCHAR(4) NOT NULL, " +
		"     distrito VARCHAR(32) NOT NULL, " +
		"     codigo_subprefeitura INTEGER NOT NULL, " +
		"     subprefeitura VARCHAR(32) NOT NULL, " +
		"     regiao5 VARCHAR(8) NOT NULL, " +
		"     regiao8 VARCHAR(8) NOT NULL, " +
		"     nome_feira VARCHAR(32) NOT NULL, " +
		"     registro VARCHAR(8) NOT NULL, " +
		"     logradouro VARCHAR(64) NOT NULL, " +
		"     numero VARCHAR(16), " +
		"     bairro VARCHAR(32), " +
		"     referencia VARCHAR(32) " +
		"  ) "

	stmt, err := db.Prepare(table)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func InsertData(db *sqlx.DB) error {
	fairsFile, err := os.OpenFile("resources/data/DEINFO_AB_FEIRASLIVRES_2014.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer fairsFile.Close()

	fairs := make([]entity.Fair, 0)

	err = gocsv.UnmarshalFile(fairsFile, &fairs)

	if err != nil {
		return err
	}

	valueStrings := make([]string, 0, len(fairs))
	valueArgs := make([]interface{}, 0, len(fairs)*3)
	for _, post := range fairs {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, post.ID)
		valueArgs = append(valueArgs, post.Longitude)
		valueArgs = append(valueArgs, post.Latitude)
		valueArgs = append(valueArgs, post.SetorCensitario)
		valueArgs = append(valueArgs, post.AreaPonderacao)
		valueArgs = append(valueArgs, post.CodigoIBGE)
		valueArgs = append(valueArgs, post.Distrito)
		valueArgs = append(valueArgs, post.CodigoSubPrefeitura)
		valueArgs = append(valueArgs, post.SubPrefeitura)
		valueArgs = append(valueArgs, post.Regiao5)
		valueArgs = append(valueArgs, post.Regiao8)
		valueArgs = append(valueArgs, post.NomeFeira)
		valueArgs = append(valueArgs, post.Registro)
		valueArgs = append(valueArgs, post.Logradouro)
		valueArgs = append(valueArgs, post.Numero)
		valueArgs = append(valueArgs, post.Bairro)
		valueArgs = append(valueArgs, post.Referencia)
	}

	stmt := fmt.Sprintf("INSERT INTO feiras_livres (id, longitude, latitude, setor_censitario, area_ponderacao, "+
		"codigo_ibge, distrito, codigo_subprefeitura, subprefeitura, regiao5, regiao8, nome_feira, registro, logradouro, "+
		"numero, bairro, referencia) VALUES %s", strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)

	if err != nil {
		return err
	}

	return nil
}
