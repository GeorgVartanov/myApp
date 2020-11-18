package main

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/database/postgres"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes/api/ginGonic"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/create"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/delete"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/update"
	"github.ru/GeorgVartanov/myApp/pkg/users/storage/pg"
	"log"
)

//docker run -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres --name postgres-server -p 5432:5432 -v pgdata:/var/lib/postgresql/data --restart=always postgres:11
//docker run -p 80:80 -e "PGADMIN_DEFAULT_EMAIL=user@domain.com" -e "PGADMIN_DEFAULT_PASSWORD=SuperSecret" -d dpage/pgadmin

func main() {
	pgDB := postgres.NewPostgresDB()

	defer func() {
		if err := pgDB.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	pgStorage := pg.NewUserPostgresStorage(pgDB)
	createUserService := create.NewCreateUserService(pgStorage)
	readUserService := read.NewCreateUserService(pgStorage)
	deleteUserService := delete.NewCreateUserService(pgStorage)
	updateUserService := update.NewCreateUserUpdate(pgStorage)

	userServ := ginGonic.NewUserConttrolers(createUserService, readUserService, updateUserService, deleteUserService)
	router := gin.Default()
	API := router.Group("/api/")
	UserRouter := API.Group("/user/")

	ginGonic.UserHandler(UserRouter, userServ)

	err := pgStorage.CreateUserTableORDrop(true)
	if err != nil {
		log.Fatal(err)
	}
	if err := router.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
