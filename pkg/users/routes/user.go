package routes

import (
	"errors"
	"strings"
)

var (
	emailIsEmpty      = errors.New("email field is empty, should be filled")
	passwordIsEmpty   = errors.New("password field is empty, should be filled")
	emailContainsChar = errors.New("email doesn't contains @, please fill email field properly")
	passwordLength    = errors.New("password is less than 6 characters")
	passwordMatch     = errors.New("passwords doesn't match")
)

//User from router
type User struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordCheck string `json:"passwordCheck"`
	DisplayName   string `json:"displayName"`
}

//ValidateFields All User fields
func (u *User) ValidateFields() error {
	if err := u.validateEmail(); err != nil {
		return err
	}
	if err := u.validatePassword(); err != nil {
		return err
	}
	u.validateDisplayName()

	return nil
}

func (u *User) validateEmail() error {
	if u.Email == "" {
		return emailIsEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return emailContainsChar
	}

	return nil
}

func (u *User) validatePassword() error {
	if u.Password == "" {
		return passwordIsEmpty
	}
	if len(u.Password) < 6 {
		return passwordLength
	}
	if u.Password != u.PasswordCheck {
		return passwordMatch
	}
	return nil
}

func (u *User) validateDisplayName() {
	if u.DisplayName == "" {
		u.DisplayName = u.Email
	}
}


