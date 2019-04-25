package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	name    string
	score   uint16
	content string
}

type QuestionType struct {
	ID   int
	name string
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
		fmt.Println(questions[0].name)
		res[questionTypes[i].name] = questions
	}
	return
}
