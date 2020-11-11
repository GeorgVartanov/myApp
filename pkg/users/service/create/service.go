package create

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

type RepositoryUserCreate interface {
	Create(u *User) error
}

type ServiceUserCreate interface {
	Create(r *routes.User) error
}

//service Create ...
type service struct {
	r RepositoryUserCreate
}

//NewCreateUserService Init
func NewCreateUserService(r RepositoryUserCreate) ServiceUserCreate {
	return &service{r: r}
}

func (s *service) Create(r *routes.User) error {
	var newUserCreate User
	newUserCreate.ConvertToServiceUser(r)
	if err := s.r.Create(&newUserCreate); err != nil {
		return err
	}
	return nil
}
