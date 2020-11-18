package update

import "github.ru/GeorgVartanov/myApp/pkg/users/routes"

//User from service Create
type User struct {
	Email       string `db:"email"`
	Password    string `db:"password"`
	DisplayName string `db:"display_name"`
}

func (u *User) ConvertToServiceUser(rUser *routes.User) {
	u.Email = rUser.Email
	u.Password = rUser.Password
	u.DisplayName = rUser.DisplayName
}
