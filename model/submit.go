package model

import (
	"github.com/jinzhu/gorm"
)

type Submit struct {
	gorm.Model
	Uid        uint
	QuestionId uint
	Soleved    bool
}

func (s *Submit) Submit(flag string) (accept bool) {
	tx := db.Begin()

	tx.Create(&s)

	q := &Question{}
	q.ID = s.QuestionId

	flagFail := tx.Model(&q).Where("id = ?", q.ID).Where("flag = ?", flag).First(&q).RecordNotFound()

	if !flagFail {
		s.Soleved = true
		accept = true
		tx.Model(&s).Where("id = ?", s.ID).Update(&s)
	}
	tx.Commit()
	return
}
