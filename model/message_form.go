package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type MessageForm struct {
	gorm.Model
	Pid      uint
	Username string
	Content  string
}

type MessageFormReturns struct {
	Id                uint
	Content           string
	Username          string
	Time              string
	OthersMessageForm []MessageFormReturns
}

func (m *MessageForm) FindAll() (messageForms []*MessageForm) {
	db.Order("created_at DESC").Find(&messageForms)
	return
}

//插入事务
func (m *MessageForm) InsertNew() error {
	tx := db.Begin()
	if m.Pid != 0 {
		var m1 MessageForm
		//查询数据库中是否有夫留言或评论的存在
		err := tx.Where("id = ?", m.Pid).First(&m1).Error
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	err := db.Create(&m)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

func DFS(Parent *MessageFormReturns, messageForms []*MessageForm) {
	//循环遍历找出待需结果集的子节点
	for _, son := range messageForms {
		if Parent.Id == son.Pid {
			sDfs := MessageFormReturns{
				son.ID,
				son.Content,
				son.Username,
				son.CreatedAt.Format("2006-01-02 15:04:05"),
				make([]MessageFormReturns, 0),
			}
			DFS(&sDfs, messageForms)
			Parent.OthersMessageForm = append(Parent.OthersMessageForm, sDfs)
		}
	}
}
