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

	messages := (&model.MessageForm{}).FindByPage()
	length := len(messages)
	var messageFormReturns []model.MessageFormReturns
	for i := 0; i < length; i++ {
		if messages[i].Pid == 0 {
			firstpDfs := model.MessageFormReturns{
				Id:       messages[i].Id,
				Content:  messages[i].Content,
				Username: messages[i].Username,
				Time:     messages[i].CreatedAt.Format("2006-01-02 15:04:05"),

				OthersMessageForm: make([]model.MessageFormReturns, 0),
			}
			model.DFS(&firstpDfs, messages)
			messageFormReturns = append(messageFormReturns, firstpDfs)
			length = len(messages)
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

	if messageLeave.Pid == 0 {
		messageForm := model.MessageForm{
			Uid:     uid,
			Content: messageLeave.Content,
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
		Pid:     messageLeave.Pid,
		Uid:     uid,
		Content: messageLeave.Content,
	}

	err = messageForm.InsertNew()
	if err != nil {
		response.MessageError(c)
		return
	}
	response.Ok(c)
	return
}
