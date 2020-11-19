package ginGonic

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/myApp/pkg/users/service"
)

type UserService interface {
	service.ServiceUser
}

//type UserService interface {
//	create.ServiceUserCreate
//	read.ServiceUserRead
//	update.ServiceUserUpdate
//	delete.ServiceUserDelete
//}

type userControllers struct {
	serv UserService
}

func NewUserControllers(userService UserService) UserControllers {
	return &userControllers{serv: userService}
}

type UserControllers interface {
	CreateUser() gin.HandlerFunc
	ReadAllUser() gin.HandlerFunc
	ReadAUser() gin.HandlerFunc
	DeleteAUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
}

//func NewUserControllers(create create.ServiceUserCreate, read read.ServiceUserRead, update update.ServiceUserUpdate, delete delete.ServiceUserDelete) UserControllers {
//	return userControllers{create: create, read: read, update: update, delete: delete}
//}
