package pg

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/users/service"
)

//need method: GetByEmail(email string) (User, error) have method: GetByEmail(email string) (*service.User, error)

func (u *UserPostgresStorage) GetAll() ([]service.User, error) {
	var users []service.User
	err := u.Select(&users, `SELECT email, display_name FROM app_user`)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserPostgresStorage) GetByEmail(email string) (*service.User, error) {
	var user service.User
	tx := u.MustBegin()
	fmt.Printf("SELECT ONE USER %v \n", user)
	result := tx.QueryRowx(`SELECT email, password, display_name FROM app_user WHERE email=$1`, email) //.StructScan(&user)
	err := result.StructScan(&user)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
