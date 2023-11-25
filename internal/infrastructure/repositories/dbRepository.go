package repositories

import (
	"database/sql"
	"time"
	"volleyapp/internal/core/ports"
	"volleyapp/internal/infrastructure/config"

	_ "github.com/lib/pq"
)

type database struct {
	*sql.DB
}

func NewDB(config config.DatabaseConfig) (ports.Database, error) {
	db, err := newDatabase(config)
	if err != nil {
		return nil, err
	}
	return &database{
		db,
	}, nil
}

func newDatabase(config config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.Url)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(
		time.Minute * time.Duration(config.ConnMaxLifetimeInMinutes),
	)
	db.SetMaxOpenConns(config.MaxOppenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d database) Close() error {
	return d.DB.Close()
}

func (d database) GetDB() *sql.DB {
	return d.DB
}
