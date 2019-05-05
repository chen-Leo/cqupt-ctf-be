package controller

import (
	"cqupt-ctf-be/model"
	"cqupt-ctf-be/utils/jwt_utils"
	response "cqupt-ctf-be/utils/response_utils"
	secret "cqupt-ctf-be/utils/secret_utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
		token, jwtErr := jwt_utils.GenerateToken(user.ID)
		if jwtErr != nil {
			response.ParamError(c)
			return
		}
		response.OkWithData(c, gin.H{
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
			"jwt":      token})
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
		token, jwtErr := jwt_utils.GenerateToken(user.ID)
		if jwtErr != nil {
			response.ParamError(c)
			return
		}
		response.OkWithData(c, gin.H{
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
			"jwt":      token})
	} else {
		if strings.Contains(err.Error(), "1062") {
			response.UsernameExist(c)
		}
	}
	response.ParamError(c)
}
