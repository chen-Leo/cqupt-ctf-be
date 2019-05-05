package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

//SubmitFlag 提交的flag后端的接收格式
type SubmitFlag struct {
	QuestionID uint   `json:"questionId" binding:"required"`
	Flag       string `json:"flag" binding:"required"`
}

//Question 获得全部问题 并根据uid给出是否已解决
func Question(c *gin.Context) {
	var questions map[string][]*model.Question
	q := model.Question{}
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
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
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	s := &model.Submit{
		Uid:        uid,
		QuestionId: f.QuestionID,
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
