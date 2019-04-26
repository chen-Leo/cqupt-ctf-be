package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Name    string
	Score   uint
	Content string
	Solve   string
}

type QuestionType struct {
	ID   int
	Name string
}

var questionTypes []QuestionType

func init() {
	db.Order("id").Find(&questionTypes)
}

func (q *Question) FindAll(uid uint) (res map[string][]*Question) {
	res = make(map[string][]*Question)
	for i := 0; i < len(questionTypes); i++ {
		var questions []*Question
		db.Order("score").Where("type_id = ?", questionTypes[i].ID).Find(&questions)
		for j := 0; j < len(questions); j++ {
			ques := questions[j]
			ques.FindSolved(uid)
		}
		fmt.Println(questions)
		res[questionTypes[i].Name] = questions
	}
	return
}

func (q *Question) FindSolved(uid uint) {
	fmt.Println(q.ID,uid)
	notFound := db.Model(&Submit{}).Where("question_id = ?", q.ID).Where("solved = 1").Where("uid = ?",uid).First(&Submit{}).RecordNotFound()
	if !notFound {
		(*q).Solve = "pass"
	}
}
