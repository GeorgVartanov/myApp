package create

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
)

type RepositoryUserCreate interface {
	InsertUser(user *User) (*read.UserWithOutPassword, error)
}

type ServiceUserCreate interface {
	Create(r *routes.User) (*read.UserWithOutPassword, error)
}

//service Create ...
type service struct {
	repo RepositoryUserCreate
}

//NewCreateUserService Init
func NewCreateUserService(r RepositoryUserCreate) ServiceUserCreate {
	return &service{repo: r}
}

func (s *service) Create(r *routes.User) (*read.UserWithOutPassword, error) {
	var newUserCreate User
	newUserCreate.ConvertToServiceUser(r)
	fmt.Printf("Service VALUES %v \n", newUserCreate)
	userDB, err := s.repo.InsertUser(&newUserCreate)
	if err != nil {
		return userDB, err
	}
	return userDB, nil
}
