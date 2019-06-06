package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func Post(c *gin.Context) {
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

    fmt.Println(uid)

}