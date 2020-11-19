package pg

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/users/service"
)

func (u *UserPostgresStorage) Create(user *service.User) (*service.User, error) {
	var dbUser service.User

	tx := u.MustBegin()

	fmt.Printf("INSERT VALUES %v \n", user)

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	err := tx.QueryRowx(`INSERT INTO app_user (email, password, display_name) VALUES ($1, $2, $3) RETURNING email, display_name`, user.Email, user.Password, user.DisplayName).StructScan(&dbUser)
	if err != nil {
		return &dbUser, err
	}

	err = tx.Commit()
	if err != nil {
		return &dbUser, nil
	}

	//fmt.Println(id)
	return &dbUser, nil
}


