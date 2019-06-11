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
	Username string `json:"username" binding:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"oldpassword" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}

type ChangeUserMessage struct {
	Name        string `json:"name"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
	Email       string `json:"email"`
	Motto       string `json:"motto"`

}

func Login(c *gin.Context) {
	var login loginUser
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.Users{Username: login.Username, Password: login.Password}
	secret.ToSha256(&user.Password)
	err = user.FindByUsernameAndPassword()
	if err == nil {

		roleTeam := model.RoleTeam{Uid: user.ID}
		roleTeam.RoleAffirm()
		team := model.Team{}
		team.FindByTeamId(roleTeam.TeamId)

		token, jwtErr := jwt_utils.GenerateToken(user.ID)
		if jwtErr != nil {
			response.ParamError(c)
			return
		}
		response.OkWithData(c, gin.H{
			"teamname": team.Name,
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
	user := model.Users{
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
			"teamname": "",
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
	//获取用户id
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	user := model.Users{}
	user.GetUserMessageByUid(uid)

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
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
	//获取用户原信息
	user := model.Users{}
	user.GetUserMessageByUid(uid)
	//获取用户队伍信息
	roleTeam := model.RoleTeam{Uid: user.ID}
	roleTeam.RoleAffirm()
	team := model.Team{}
	team.FindByTeamId(roleTeam.TeamId)

	if changeUserMessage.OldPassword != "" && changeUserMessage.NewPassword != "" {

		secret.ToSha256(&changeUserMessage.OldPassword)
		secret.ToSha256(&changeUserMessage.NewPassword)

		if user.Password != changeUserMessage.OldPassword {
			response.PasswordError(c)
			return
		}

		user.Password = changeUserMessage.NewPassword
	}

	if changeUserMessage.Name != "" {
		user.Username = changeUserMessage.Name
	}
	if changeUserMessage.Email != "" {
		user.Email = changeUserMessage.Email
	}
	if changeUserMessage.Motto != "" {
		user.Motto = changeUserMessage.Motto
	}

	err = user.UserMessageChange()
	if err != nil {
		response.UsernameOrEmailExist(c)
		return
	}
	response.OkWithData(c, gin.H{
		"teamname": team.Name,
		"username": user.Username,
		"email":    user.Email,
		"motto":    user.Motto,
	})

}

//通过用户名获得用户信息
func UserMessageGet(c *gin.Context) {
	var getUserMessage GetUserMessage
	err := c.ShouldBindJSON(&getUserMessage)
	if err != nil {
		response.ParamError(c)
		return
	}
	user := model.Users{Username: getUserMessage.Username}
	user.GetUserMessageByUsername()
	if user.ID != 0 {
		roleTeam := model.RoleTeam{Uid: user.ID}
		roleTeam.RoleAffirm()
		team := model.Team{}
		team.FindByTeamId(roleTeam.TeamId)

		response.OkWithData(c, gin.H{
			"teamname": team.Name,
			"username": user.Username,
			"email":    user.Email,
			"motto":    user.Motto,
		})
		return
	}
	response.OkWithData(c, gin.H{})
}


