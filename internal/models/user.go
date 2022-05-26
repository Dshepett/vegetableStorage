package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	Role           int    `json:"role"`
}

func (u *User) EncryptPassword() error {
	err := validation.Validate(u.HashedPassword, validation.Length(8, 100), validation.Required)
	if err != nil {
		return err
	}
	res, err := bcrypt.GenerateFromPassword([]byte(u.HashedPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(res)
	return nil
}

func (u *User) CheckEmail() error {
	return validation.Validate(u.Email, validation.Required, is.Email)
}
