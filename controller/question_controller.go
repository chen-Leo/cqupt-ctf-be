package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"github.com/gin-gonic/gin"
)

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
