package ginGonic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/create"
	"net/http"
)

func CreateUser(create create.ServiceUserCreate) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newApiUser routes.User
		if err := c.ShouldBindJSON(&newApiUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newApiUser.ValidateFields(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("ROUTER VALUES %v \n", newApiUser)
		userDB, err := create.Create(&newApiUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func (create userConttrolers) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newApiUser routes.User
		if err := c.ShouldBindJSON(&newApiUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newApiUser.ValidateFields(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("ROUTER VALUES %v \n", newApiUser)
		userDB, err := create.create.Create(&newApiUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}
