package model

import (
	"github.com/jinzhu/gorm"
)

type Submit struct {
	gorm.Model
	Uid        uint
	QuestionId uint
	Solved    bool
}

func (s *Submit) Submit(flag string) (accept uint8) {
	tx := db.Begin()

	q := &Question{}
	q.ID = s.QuestionId

	isNotSolved := tx.Where("question_id = ?", q.ID).Where("solved = 1").Where("uid = ?", s.Uid).First(&Submit{}).RecordNotFound()

	if isNotSolved {
		tx.Create(&s)

		flagFail := tx.Model(&q).Where("id = ?", q.ID).Where("flag = ?", flag).First(&q).RecordNotFound()

		if !flagFail {
			s.Solved = true
			accept = 1
			tx.Model(&s).Where("id = ?", s.ID).Update(&s)
		}
	} else {
		accept = 2
	}
	tx.Commit()
	return
}
