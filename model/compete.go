package model

import "github.com/jinzhu/gorm"

type Compete struct {
	gorm.Model
	Name         string
	Introduction string //简介或者公告
	Type        string
}

func (c *Compete) FindAll() (competes []*Compete) {
	db.Order("created_at DESC").Find(&competes)
	return
}
