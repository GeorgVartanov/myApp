package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes/api/ginGonic"
	"log"
	"os"
)

// StartGinServer ...
func StartGinServer() {
	r := gin.Default()
	MyApi := r.Group("/api/")
	UserRouter := MyApi.Group("/user/")
	ginGonic.UserRoutes(UserRouter)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
