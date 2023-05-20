package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `validate:"-" json:"id"`
	Email             string `validate:"email" json:"email"`
	Password          string `validate:"required_without=EncryptedPassword,omitempty,min=6,max=30" json:"password,omitempty"`
	EncryptedPassword string `validate:"-" json:"-"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)

}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
