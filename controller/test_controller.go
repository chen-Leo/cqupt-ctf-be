package controller

import (
	"cqupt-ctf-be/model"
	"fmt"
	"github.com/gin-gonic/gin"
)


func Post(c *gin.Context) {
	user := model.User{}
	user.GetUserMessageByUid(8)
    fmt.Println(user)
	user.Username = "test7"
	_ = user.UserMessageChange()
}