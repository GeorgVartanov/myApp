package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/create"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/delete"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/update"
)

type UserService interface {
	create.ServiceUserCreate
	read.ServiceUserRead
	update.ServiceUserUpdate
	delete.ServiceUserDelete
}

type userConttrolers struct {
	create create.ServiceUserCreate
	read   read.ServiceUserRead
	update update.ServiceUserUpdate
	delete delete.ServiceUserDelete
}

type UserControllers interface {
	CreateUser() gin.HandlerFunc
	ReadAllUser() gin.HandlerFunc
	ReadAUser() gin.HandlerFunc
	DeleteAUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
}

func NewUserConttrolers(create create.ServiceUserCreate, read read.ServiceUserRead, update update.ServiceUserUpdate, delete delete.ServiceUserDelete) UserControllers {
	return userConttrolers{create: create, read: read, update: update, delete: delete}
}

func UserHandler(router *gin.RouterGroup, u UserControllers) {

	router.POST("/create/", u.CreateUser())
	router.GET("/all/", u.ReadAllUser())
	router.POST("/one/", u.ReadAUser())
	router.DELETE("/delete/", u.DeleteAUser())
	router.PUT("/update/", u.UpdateUser())

}
