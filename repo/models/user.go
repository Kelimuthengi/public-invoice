package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrPasswordTooLong = errors.New(bcrypt.ErrPasswordTooLong.Error())
)

type User struct {
	gorm.Model
	Username    string    `gorm:"size:255;not null " json:"username" `
	Email       string    `gorm:"size:255;not null; unique" json:"email"`
	Address     string    `gorm:"size:255;not null" json:"address"`
	Phonenumber string    `gorm:"size:255;not null; unique" json:"phonenumber"`
	Products    []Product `gorm:"foreignKey:UserID"`
	Password    string    `gorm:"size:255"`
	LPassword   string
}

// func to hash password

func (u *User) HashUserPassword(userPassword string) ([]byte, error) {

	if len(userPassword) > 72 {
		return []byte(""), ErrPasswordTooLong
	}

	hashedByte, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

	return hashedByte, err
}

func (u *User) CompareUserPasswords(hashedPassword,password []byte) error {

	err := bcrypt.CompareHashAndPassword(hashedPassword,password)
	return err
}
