package ginGonic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/update"
	"net/http"
)

func UpdateUser(upd update.ServiceUserUpdate) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser routes.User
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := updateUser.ValidateFieldsUpdate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("UPDATE VALUES %v \n", updateUser)
		userDB, err := upd.Update(&updateUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}



func (upd userConttrolers) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser routes.User
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := updateUser.ValidateFieldsUpdate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("UPDATE VALUES %v \n", updateUser)
		userDB, err := upd.update.Update(&updateUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}