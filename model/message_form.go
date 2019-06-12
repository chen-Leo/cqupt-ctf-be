package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type MessageForm struct {
	gorm.Model
	Pid     uint
	Uid     uint
	Content string
}

type Message struct {
	Id        uint
	Pid       uint
	Username  string
	Content   string
	CreatedAt time.Time
}

type MessageFormReturns struct {
	Id                uint
	Content           string
	Username          string
	Time              string
	OthersMessageForm []MessageFormReturns
}

func (m *MessageForm) FindByPage() (message []*Message) {
	db.Table("message_form").Select("users.username,message_form.id,message_form.pid,message_form.content,message_form.created_at").
		Joins("left join users on users.id = message_form.uid").Order("message_form.created_at DESC").
		Find(&message)
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

func DFS(Parent *MessageFormReturns, messages []*Message) {
	//循环遍历找出待需结果集的子节点
	for _, son := range messages {
		if Parent.Id == son.Pid {
			sDfs := MessageFormReturns{
				son.Id,
				son.Content,
				son.Username,
				son.CreatedAt.Format("2006-01-02 15:04:05"),
				make([]MessageFormReturns, 0),
			}
			DFS(&sDfs, messages)
			Parent.OthersMessageForm = append(Parent.OthersMessageForm, sDfs)
		}
	}
}
