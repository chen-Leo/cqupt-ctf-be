package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Compete struct {
	gorm.Model
	Name         string
	Introduction string //简介或者公告
	Type        string
	EndTime    time.Time
}

func (c *Compete) FindAll() (competes []*Compete) {
	db.Order("end_time ").Order("created_at DESC").Find(&competes)
	return
}
