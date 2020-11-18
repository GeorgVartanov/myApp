package pg

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/users/service/read"
)

func (u *UserPostgresStorage) SelectAll() ([]read.UserWithOutPassword, error) {
	var users []read.UserWithOutPassword

	err := u.Select(&users, `SELECT email, display_name FROM app_user`)

	//tx := u.MustBegin()
	//
	//fmt.Printf("SELECT ALL USERS %v \n", users)
	//
	//// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//err := tx.QueryRowx(`SELECT email, display_name FROM app_user `).StructScan(&users)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserPostgresStorage) SelectOne(email string) (*read.User, error) {
	var user read.User

	tx := u.MustBegin()

	fmt.Printf("SELECT ONE USER %v \n", user)

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	err := tx.QueryRowx(`SELECT email, password, display_name FROM app_user WHERE email=$1`, email).StructScan(&user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserPostgresStorage) SelectOneWithOutPassword(email string) (*read.UserWithOutPassword, error) {
	var user read.UserWithOutPassword

	tx := u.MustBegin()

	fmt.Printf("SELECT ONE USER %v \n", user)

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	err := tx.QueryRowx(`SELECT email, display_name FROM app_user WHERE email=$1`, email).StructScan(&user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
