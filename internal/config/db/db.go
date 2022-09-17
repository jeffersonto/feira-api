package config

import (
	"github.com/jeffersonto/feira-api/pkg/warmupdb"
	"github.com/jmoiron/sqlx"
)

func DB() (*sqlx.DB, error) {
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

	return db, nil
}
