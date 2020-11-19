package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"net/http"
)

func (del userControllers) DeleteAUser() gin.HandlerFunc {
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

		err := del.serv.DeleteByEmail(user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "Deleted"})
		return
	}
}
