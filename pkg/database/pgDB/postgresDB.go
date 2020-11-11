package pgDB

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres driver
	"log"
)

// PostgresDB Postgres class struct
type PostgresDB struct {
	*sqlx.DB
}

// NewPostgresDB get Postgres class struct
func NewPostgresDB() *PostgresDB {
	dbPath := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
	pg := PostgresDB{}
	if err := pg.Connect(dbPath); err != nil {
		log.Fatalln(err)
	}

	return &pg
}

// Connect connection to Postgresql
func (s *PostgresDB) Connect(dbPath string) error {
	db, err := sqlx.Open("postgres", dbPath)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.DB = db
	return nil
}
