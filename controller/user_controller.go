package controller

import (
	"cqupt-ctf-be/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var login loginUser
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.String(http.StatusBadRequest, "param error")
		return
	}
	user := model.User{Username: login.Username, Password: login.Password}
	user.FindByUsernameAndPassword()
	if user.ID != 0 {
		c.String(http.StatusOK, "hello "+user.Username)
		return
	}
	c.String(http.StatusBadRequest,"password error")
}
