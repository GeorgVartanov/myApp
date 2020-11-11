package main

import (
	"github.ru/GeorgVartanov/myApp/pkg/database/pgDB"
	"github.ru/GeorgVartanov/myApp/pkg/users/storage/pgStorage"
	"log"
)

//docker run -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres --name postgres-server -p 5432:5432 -v pgdata:/var/lib/postgresql/data --restart=always postgres:11
//docker run -p 80:80 -e "PGADMIN_DEFAULT_EMAIL=user@domain.com" -e "PGADMIN_DEFAULT_PASSWORD=SuperSecret" -d dpage/pgadmin

func main() {
	pgDB := pgDB.NewPostgresDB()

	defer func() {
		if err := pgDB.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	pgStorageDB := pgStorage.NewUserPostgresStorage(pgDB)
	err:= pgStorageDB.CreateUserTableORDrop(true)
	if err != nil {
		log.Fatal(err)
	}
}
