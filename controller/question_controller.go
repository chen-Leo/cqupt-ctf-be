package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SubmitFlag struct {
	QuestionId uint   `json:"questionId" binding:"required"`
	Flag       string `json:"flag" binding:"required"`
}

func Question(c *gin.Context) {
	var questions map[string][]*model.Question
	q := model.Question{}
	s := sessions.Default(c)
	uidInterface := s.Get("uid")
	//TODO:测试
	var uid uint
	if uidInterface == nil {
		uid = 1
	} else {
		uid = uint(uidInterface.(int))
	}
	questions = q.FindAll(uid)
	res := make([]gin.H, len(questions))
	i := 0
	for key, value := range questions {
		res[i] = gin.H{
			"name":      key,
			"questions": value}
		i++
	}
	response.OkWithArray(c, res)
}

func Submit(c *gin.Context) {
	var f SubmitFlag
	err := c.ShouldBindJSON(&f)
	if err != nil {
		fmt.Println(err.Error())
		response.ParamError(c)
		return
	}
	session := sessions.Default(c)
	uidInterface := session.Get("uid")
	//TODO:测试
	var uid uint
	if uidInterface == nil {
		uid = 1
	} else {
		uid = uint(uidInterface.(int))
	}
	s := &model.Submit{
		Uid:        uid,
		QuestionId: f.QuestionId,
	}
	fmt.Println(s.Uid)
	accept := s.Submit(f.Flag)

	switch accept {
	case 1:
		response.Ok(c)
	case 2:
		response.IsSolved(c)
	case 0:
		response.FlagErr(c)
	default:
		response.ParamError(c)
	}
}
