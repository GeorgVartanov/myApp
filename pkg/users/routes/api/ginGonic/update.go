package ginGonic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"net/http"
)

func (upd userControllers) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser routes.UserIn
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := updateUser.ValidateFieldsUpdate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("UPDATE VALUES %v \n", updateUser)
		err := upd.serv.UpdateByEmail(&updateUser)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": "Updated"})
		return
	}
}
