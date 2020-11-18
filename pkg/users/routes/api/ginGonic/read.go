package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
	"net/http"
)

func ReadAllUser(read read.ServiceUserRead) gin.HandlerFunc {
	return func(c *gin.Context) {

		userDB, err := read.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func ReadAUser(read read.ServiceUserRead) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user routes.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := user.ValidateEmail(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDB, err := read.GetOne(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func (read userConttrolers) ReadAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userDB, err := read.read.GetAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func (read userConttrolers) ReadAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user routes.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := user.ValidateEmail(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDB, err := read.read.GetOne(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}
