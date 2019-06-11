package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Motto    string
}

func (u *Users) FindByUsernameAndPassword() error {
	err := db.Where(u).First(&u)
	if err.Error != nil {
		fmt.Println(err.Error.Error())
		return err.Error
	}
	return nil
}

func (u *Users) InsertNew() error {
	err := db.Create(&u)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (u *Users) FindAll() (users []*Users) {
	db.Find(&users)
	return
}

func (u *Users) FindRank() (solved uint, submitted uint, score uint) {
	var submits []Submit
	db.Where("uid = ?", u.ID).Find(&submits)
	submitted = uint(len(submits))
	for i := 0; i < len(submits); i++ {
		s := submits[i]
		if s.Solved {
			solved++
			q := Question{}
			q.ID = s.QuestionId
			db.Where(q).Find(&q)
			score += q.Score
		}
	}
	return
}

//根据uid返回用户user信息
func (u *Users) GetUserMessageByUid(uid uint) *Users {
	db.Where("id = ?", uid).First(&u)
	return u
}

//更具username返回user信息
func (u *Users) GetUserMessageByUsername() {
	db.Where("username = ?", u.Username).First(&u)
}

func (u *Users) UserMessageChange() error {
	err := db.Model(&u).Updates(u)
	return err.Error
}
