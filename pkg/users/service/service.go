package service

import (
	"github.ru/GeorgVartanov/myApp/pkg/users/routes"
)

type RepositoryUser interface {
	Create(user *User) (*User, error)
	GetAll() ([]User, error)
	GetByEmail(email string) (*User, error)
	UpdateByEmail(user *User) error
	DeleteByEmail(email string) error
}

type ServiceUser interface {
	Create(user *routes.UserIn) (*routes.UserOut, error)
	ReadAll() ([]routes.UserOut, error)
	ReadByEmail(email string) (*routes.UserOut, error)
	UpdateByEmail(user *routes.UserIn) error
	DeleteByEmail(email string) error
}

func NewService(repository RepositoryUser) ServiceUser {
	return &service{repo: repository}
}

type service struct {
	repo RepositoryUser
}




