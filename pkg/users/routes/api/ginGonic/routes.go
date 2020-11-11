package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"net/http"
)

func UserRoutes(r *gin.RouterGroup) {

	r.POST("/create/", CreateUser())

}

func CreateUser() gin.HandlerFunc {
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


		c.JSON(http.StatusOK, gin.H{"data": newApiUser})
		return
	}
}
