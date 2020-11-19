package main

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/database/anotherPostgres"
	"github.ru/GeorgVartanov/myApp/pkg/users/service"
)

//docker run -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres --name postgres-server -p 5432:5432 -v pgdata:/var/lib/postgresql/data --restart=always postgres:11
//docker run -p 80:80 -e "PGADMIN_DEFAULT_EMAIL=user@domain.com" -e "PGADMIN_DEFAULT_PASSWORD=SuperSecret" -d dpage/pgadmin

//func main() {
//	pgDB := postgres.NewPostgresDB()
//
//	defer func() {
//		if err := pgDB.Close(); err != nil {
//			log.Fatal(err.Error())
//		}
//	}()
//	pgStorage := pg.NewUserPostgresStorage(pgDB)
//
//	UserService := service.NewService(pgStorage)
//
//	userServ := ginGonic.NewUserControllers(UserService)
//	router := gin.Default()
//	API := router.Group("/api/")
//	UserRouter := API.Group("/user/")
//
//	ginGonic.UserHandler(UserRouter, userServ)
//
//	err := pgStorage.CreateUserTableORDrop(true)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := router.Run(); err != nil {
//		log.Fatal(err.Error())
//	}
//}


func main() {
	db := anotherPostgres.Connect()
	//ctx := context.Background()

	var app_user service.User
	err := db.Model(&app_user).Where("id = ?", 11).Select()

	//_, err := db.QueryOneContext(ctx, pg.Scan(&user), "SELECT * FROM app_user WHERE id=1")
	if err != nil {
		panic(err)
	}
	fmt.Println(app_user)

	defer db.Close()
}
