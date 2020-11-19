package service

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

func (s service) Create(user *routes.UserIn) (*routes.UserOut, error) {

	var newUser User
	newUser.ConvertToServiceUser(user)
	createdUser, err := s.repo.Create(&newUser)
	if err != nil {
		return nil, err
	}
	//var userOut routes.UserOut
	//userOut.ConvertToAPIUserOut(createdUser)
	userOUT := createdUser.ConvertToAPIUserOut()
	return userOUT, nil
}
