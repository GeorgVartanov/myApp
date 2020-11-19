package service

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

//User from service
type User struct {
	tableName struct{} `pg:"app_user,alias:app_user"`

	Email       string `db:"email"`
	Password    string `db:"password"`
	DisplayName string `db:"display_name"`
}

func (u *User) ConvertToServiceUser(rUser *routes.UserIn) {
	u.Email = rUser.Email
	u.Password = rUser.Password
	u.DisplayName = rUser.DisplayName
}

func (u *User) ConvertToAPIUserOut() *routes.UserOut {
	var userOut routes.UserOut
	userOut.Email = u.Email
	userOut.DisplayName = u.DisplayName

	return &userOut
}
