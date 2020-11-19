package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"net/http"
)

func (read userControllers) ReadAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userDB, err := read.serv.ReadAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func (read userControllers) ReadAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user routes.UserIn

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := user.ValidateEmail(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDB, err := read.serv.ReadByEmail(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}
