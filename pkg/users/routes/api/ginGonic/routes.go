package ginGonic

import (
	"github.com/gin-gonic/gin"
)



func UserHandler(router *gin.RouterGroup, u UserControllers) {

	router.POST("/create/", u.CreateUser())
	router.GET("/all/", u.ReadAllUser())
	router.POST("/one/", u.ReadAUser())
	router.DELETE("/delete/", u.DeleteAUser())
	router.PUT("/update/", u.UpdateUser())

}
