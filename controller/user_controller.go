package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	secret "cqupt-ctf-be/utils/secret_utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type loginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpUser struct {
	loginUser
	Email string `json:"email" binding:"required"`
}

func Login(c *gin.Context) {
	var login loginUser
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.User{Username: login.Username, Password: login.Password}
	secret.ToSha256(&user.Password)
	err = user.FindByUsernameAndPassword()
	if err == nil {
		session := sessions.Default(c)
		session.Set("uid", user.ID)
		err=session.Save()
		if err!=nil {
			fmt.Println(err.Error())
		}
		response.OkWithData(c, gin.H{
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  10011,
		"message": "password error",
	})
}

func SignUp(c *gin.Context) {
	var signUpUser signUpUser
	err := c.ShouldBindJSON(&signUpUser)
	if err != nil {
		fmt.Println(err.Error())
		response.ParamError(c)
		return
	}
	user := model.User{
		Username: signUpUser.Username,
		Password: signUpUser.Password,
		Email:    signUpUser.Email,
	}
	secret.ToSha256(&user.Password)
	err = user.InsertNew()
	if err == nil {
		s := sessions.Default(c)
		s.Set("uid", user.ID)
		err=s.Save()
		if err!=nil {
			fmt.Println(err.Error())
		}
		response.OkWithData(c, gin.H{
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
		})
		return
	} else {
		if strings.Contains(err.Error(), "1062") {
			response.UsernameExist(c)
			return
		}
	}
	response.ParamError(c)
}
