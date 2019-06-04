package model

import "github.com/jinzhu/gorm"

type News struct {
	gorm.Model
	title string
	content string
}

func (n *News) FindAll() (news []*News) {
	db.Find(&news).Order("created _at DESC")
	return
}


