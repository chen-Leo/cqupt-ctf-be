package model

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Name    string
	Score   uint
	Content string
}

type QuestionType struct {
	ID   int
	Name string
}

var questionTypes []QuestionType

func init() {
	db.Find(&questionTypes)
}

func (q *Question) FindAll() (res map[string][]Question) {
	res = make(map[string][]Question)
	for i := 0; i < len(questionTypes); i++ {
		var questions []Question
		db.Where("type_id = ?", questionTypes[i].ID).Find(&questions)
		res[questionTypes[i].Name] = questions
	}
	return
}
