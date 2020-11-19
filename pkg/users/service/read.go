package service

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

func (s service) ReadAll() ([]routes.UserOut, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var userOut []routes.UserOut
	for _, v := range users {
		userOUT := v.ConvertToAPIUserOut()
		userOut = append(userOut, *userOUT)
	}
	return userOut, nil

}

func (s service) ReadByEmail(email string) (*routes.UserOut, error) {
	//var userOut routes.UserOut
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	//userOut.ConvertToAPIUserOut(user)
	userOUT := user.ConvertToAPIUserOut()
	return userOUT, nil

}
