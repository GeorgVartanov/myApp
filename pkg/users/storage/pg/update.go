package pg

import (
	"github.ru/GeorgVartanov/myApp/pkg/users/service"
)

func (u *UserPostgresStorage) UpdateByEmail(user *service.User) error {

	_, err := u.Exec(`UPDATE app_user SET email=$1, display_name=$2 WHERE email=$3`, user.Email, user.DisplayName, user.Email)

	if err != nil {
		return err
	}

	return  nil
}
