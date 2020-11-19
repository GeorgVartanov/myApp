package routes

import (
	"errors"
	"strings"
)

var (
	emailIsEmpty       = errors.New("email field is empty, should be filled")
	passwordIsEmpty    = errors.New("password field is empty, should be filled")
	emailContainsChar  = errors.New("email doesn't contains @, please fill email field properly")
	passwordLength     = errors.New("password is less than 6 characters")
	passwordMatch      = errors.New("passwords doesn't match")
	displayNameIsEmpty = errors.New("display name field is empty, should be filled")
)

//UserIn from router
type UserIn struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordCheck string `json:"passwordCheck"`
	DisplayName   string `json:"displayName"`
}

func (u *UserIn) ValidateFieldsUpdate() error {
	if u.Email == "" {
		return emailIsEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return emailContainsChar
	}
	if u.DisplayName == "" {
		return displayNameIsEmpty
	}

	return nil
}

//ValidateFields All UserIn fields
func (u *UserIn) ValidateFields() error {
	if err := u.ValidateEmail(); err != nil {
		return err
	}
	if err := u.validatePassword(); err != nil {
		return err
	}
	u.validateDisplayName()

	return nil
}

func (u *UserIn) ValidateEmail() error {
	if u.Email == "" {
		return emailIsEmpty
	}
	if !strings.Contains(u.Email, "@") {
		return emailContainsChar
	}

	return nil
}

func (u *UserIn) validatePassword() error {
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

func (u *UserIn) validateDisplayName() {
	if u.DisplayName == "" {
		u.DisplayName = u.Email
	}
}
