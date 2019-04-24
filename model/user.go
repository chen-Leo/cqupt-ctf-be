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

func (u *User) FindByUsernameAndPassword() {
	err := db.Where(u).First(&u)
	if err.Error != nil {
		fmt.Println(err.Error.Error())
	}
}
