package pg

import (
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/update"
)

func (u *UserPostgresStorage) UpdateUser(user *update.User) (*read.UserWithOutPassword, error) {
	var dbUser read.UserWithOutPassword

	_, err := u.Exec(`UPDATE app_user SET email=$1, display_name=$2 WHERE email=$3`, user.Email, user.DisplayName, user.Email)

	//tx := u.MustBegin()
	//
	//fmt.Printf("UPDATE VALUES %v \n", user)
	//
	//// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//err := tx.QueryRowx(`UPDATE app_user SET email=$1, display_name=$2 WHERE email=$1  RETURNING email, display_name `, user.Email, user.DisplayName).StructScan(&dbUser)
	//if err != nil {
	//	return &dbUser, err
	//}

	//err = tx.Commit()
	if err != nil {
		return &dbUser, nil
	}



	return &dbUser, nil
}
