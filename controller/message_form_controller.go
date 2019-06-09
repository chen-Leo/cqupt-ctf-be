package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"github.com/gin-gonic/gin"
)

type MessageLeave struct {
	Pid     uint   `json:"pid"`
	Content string `json:"content" binding:"required"`
}

func MessageFormAll(c *gin.Context) {
	messageForms := (&model.MessageForm{}).FindAll()
	length := len(messageForms)
	var messageFormReturns []model.MessageFormReturns
	for i := 0; i < length; i++ {
		if messageForms[i].Pid == 0 {
			firstpDfs := model.MessageFormReturns{
				messageForms[i].ID,
				messageForms[i].Username,
				messageForms[i].Content,
				make([]model.MessageFormReturns, 0),
			}
			model.DFS(&firstpDfs, messageForms)
			messageFormReturns = append(messageFormReturns, firstpDfs)
			length = len(messageForms)
		}
	}
	response.OkWithData(c, gin.H{"message": messageFormReturns})
}

func MessageFormAdd(c *gin.Context) {
	var messageLeave MessageLeave
	err := c.ShouldBindJSON(&messageLeave)
	if err != nil {
		response.ParamError(c)
		return
	}

	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
	user := (&model.User{}).GetUserMessageByUid(uid)

	if messageLeave.Pid == 0 {
		messageForm := model.MessageForm{
			Pid:      0,
			Username: user.Username,
			Content:  messageLeave.Content,
		}
		err := messageForm.InsertNew()
		if err != nil {
			response.ParamError(c)
			return
		}
		response.Ok(c)
		return
	}
	messageForm := model.MessageForm{
		Pid:      messageLeave.Pid,
		Username: user.Username,
		Content:  messageLeave.Content,
	}
	err = messageForm.InsertNew()
	if err != nil {
		response.MessageError(c)
		return
	}
	response.Ok(c)
	return

}
