package service

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

func (s service) UpdateByEmail(user *routes.UserIn) error {
	var updUser User
	updUser.ConvertToServiceUser(user)
	err := s.repo.UpdateByEmail(&updUser)
	if err != nil {
		return err
	}
	return nil
}
