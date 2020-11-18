package update

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
)

type RepositoryUserUpdate interface {
	UpdateUser(user *User) (*read.UserWithOutPassword, error)
}

type ServiceUserUpdate interface {
	Update(r *routes.User) (*read.UserWithOutPassword, error)
}

//service Create ...
type service struct {
	repo RepositoryUserUpdate
}

//NewCreateUserService Init
func NewCreateUserUpdate(r RepositoryUserUpdate) ServiceUserUpdate {
	return &service{repo: r}
}

func (s service) Update(r *routes.User) (*read.UserWithOutPassword, error) {
	var userUpdate User
	userUpdate.ConvertToServiceUser(r)
	fmt.Printf("Service VALUES %v \n", userUpdate)
	userDB, err := s.repo.UpdateUser(&userUpdate)
	if err != nil {
		return userDB, err
	}
	return userDB, nil
}
