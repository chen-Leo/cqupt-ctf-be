package model

import "github.com/jinzhu/gorm"

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
	OthersMessageForm []MessageFormReturns
}

func (m *MessageForm) FindAll() (messageForms []*MessageForm) {
	db.Order("created_at DESC").Find(&messageForms)
	return
}

func (m *MessageForm) InsertNew() error {
	err := db.Create(&m)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func DFS(Parent *MessageFormReturns, messageForms []*MessageForm) {
	//循环遍历找出待需结果集的子节点
	for _, son := range messageForms {
		if Parent.Id == son.Pid {
			sDfs := MessageFormReturns{
				son.ID,
				son.Username,
				son.Content,
				make([]MessageFormReturns, 0),
			}
			DFS(&sDfs, messageForms)
			Parent.OthersMessageForm = append(Parent.OthersMessageForm, sDfs)
		}
	}
}
