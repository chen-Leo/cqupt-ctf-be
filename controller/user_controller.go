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

type GetUserMessage struct {
	Username string`json:"username" binding:"required"`
}

type ChangePassword struct {
	OldPassword  string  `json:"oldpassword" binding:"required"`
	NewPassword  string  `json:"newpassword" binding:"required"`
}

type ChangeUserMessage struct {
	Name  string         `json:"name"`
	OldPassword  string  `json:"oldpassword"`
	NewPassword  string  `json:"newpassword"`
	Email      string    `json:"email"`
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

		roleTeam := model.RoleTeam{Uid:user.ID}
		roleTeam.RoleAffirm()

		token, jwtErr := jwt_utils.GenerateToken(user.ID)
		if jwtErr != nil {
			response.ParamError(c)
			return
		}
		response.OkWithData(c, gin.H{
			"teamid"  : roleTeam.TeamId,
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
			"teamid"  :0,
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
			"jwt":      token})
		return
	} else {
		if strings.Contains(err.Error(), "1062") {
			response.UsernameOrEmailExist(c)
			return
		}
	}
	response.ParamError(c)
}

//修改密码
//create by sao
func PasswordChange(c *gin.Context) {
	 var changePassword ChangePassword
	 err := c.ShouldBindJSON(&changePassword)
	 if err != nil {
		response.ParamError(c)
		return
	 }
	jwtStr := c.GetHeader("Authorization")
	jwtStr = strings.Replace(jwtStr, "Bearer ", "", 7)
	u, err := jwt_utils.ParseToken(jwtStr)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.User{}
	user.GetUserMessageByUid(u.Uid)
	secret.ToSha256(&changePassword.OldPassword)
	secret.ToSha256(&changePassword.NewPassword)
	if user.Password == changePassword.OldPassword {
		user.Password = changePassword.NewPassword
		err := user.UserMessageChange()
			if err != nil {
				response.UsernameOrEmailExist(c)
				return
			}
		response.Ok(c)
		return
		}
	response.PasswordError(c)

	}

//用户信息修改
//create by sao
func UserMessageChange(c *gin.Context) {
	var changeUserMessage ChangeUserMessage
	err := c.ShouldBindJSON(&changeUserMessage)
	if err != nil {
		response.ParamError(c)
		return
	}
	jwtStr := c.GetHeader("Authorization")
	jwtStr = strings.Replace(jwtStr, "Bearer ", "", 7)
	u, err := jwt_utils.ParseToken(jwtStr)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.User{}
	user.GetUserMessageByUid(u.Uid)
	secret.ToSha256(&changeUserMessage.OldPassword)
	secret.ToSha256(&changeUserMessage.NewPassword)
	if user.Password != changeUserMessage.OldPassword {
		response.PasswordError(c)
		return
	}
	user.Password = changeUserMessage.NewPassword
	if changeUserMessage.Name != "" {
	user.Username = changeUserMessage.Name
    }
	if changeUserMessage.Email != ""{
		user.Email = changeUserMessage.Email
	}
	err = user.UserMessageChange()
	if err != nil {
		response.UsernameOrEmailExist(c)
		return
	}
    response.Ok(c)
}


func UserMessageGet(c *gin.Context) {
	var getUserMessage GetUserMessage
	err := c.ShouldBindJSON(&getUserMessage)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.User{Username:getUserMessage.Username}
	user.GetUserMessageByUsername()
	if user.ID != 0 {
		roleTeam := model.RoleTeam{Uid:user.ID}
		roleTeam.RoleAffirm()
		response.OkWithData(c, gin.H{
			"teamid"  : roleTeam.TeamId,
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
		})
		return
	}
	response.OkWithData(c,gin.H{})
}






