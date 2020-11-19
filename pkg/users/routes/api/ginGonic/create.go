package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"net/http"
)

func (u userControllers) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newApiUser routes.UserIn
		if err := c.ShouldBindJSON(&newApiUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newApiUser.ValidateFields(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDB, err := u.serv.Create(&newApiUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}
