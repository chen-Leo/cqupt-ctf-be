package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Motto    string
}

func (u *User) FindByUsernameAndPassword() error {
	err := db.Where(u).First(&u)
	if err.Error != nil {
		fmt.Println(err.Error.Error())
		return err.Error
	}
	return nil
}

func (u *User) InsertNew() error {
	err := db.Create(&u)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
