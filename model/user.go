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

func (u *User) FindAll() (users []*User) {
	db.Find(&users)
	return
}

func (u *User) FindRank() (solved uint, submitted uint, score uint) {
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
func (u *User) GetUserMessageByUid(uid uint) {
   db.Where("id = ?",uid).First(&u)
}

//更具username返回user信息
func (u *User) GetUserMessageByUsername() {
	db.Where("username = ?",u.Username).First(&u)
}

func (u *User) UserMessageChange() error {
	err  := db.Model(&u).Updates(u)
    return err.Error
}






