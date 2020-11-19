package anotherPostgres

import (
	"github.com/go-pg/pg/v10"
	"golang.org/x/net/context"
	"log"
)

type AnotherPG struct {
	*pg.DB
}

func Connect() *AnotherPG {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	//_, err := db.Exec("SELECT 1")
	ctx := context.Background()
	err := db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &AnotherPG{db}
}
