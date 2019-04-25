package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SubmitFlag struct {
	QuestionId uint   `json:"questionId" binding:"required"`
	Flag       string `json:"flag" binding:"required"`
}

func Question(c *gin.Context) {
	var questions map[string][]model.Question
	q := model.Question{}
	questions = q.FindAll()
	res := gin.H{}
	for key, value := range questions {
		res[key] = value
	}
	response.OkWithData(c, res)
}

func Submit(c *gin.Context) {
	var f SubmitFlag
	err := c.ShouldBindJSON(&f)
	if err != nil {
		response.ParamError(c)
		return
	}
	session := sessions.Default(c)
	uid := session.Get("uid").(uint)
	s := &model.Submit{
		Uid:        uid,
		QuestionId: f.QuestionId,
	}
	accept := s.Submit(f.Flag)
	if accept {
		response.Ok(c)
	} else {
		response.FlagErr(c)
	}
}
