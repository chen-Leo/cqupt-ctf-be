//team controller
//@author doc.sao
package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//队伍成员对应表
type RoleTeam struct {
	TeamId uint `json:"teamId" binding:"required"`
}

//team表
type CreateTeam struct {
	Name         string `json:"name" binding:"required"`
	Introduction string `json:"introduction" `
}

//队员申请是否同意表
type TeamApplication struct {
	NewUserName string `json:"newusername" binding:"required"`
	AgreeOrNot  int    `json:"agreeornot" binding:"required"`
}

type KickUid struct {
	PoorUid uint `json:"pooruid" binding:"required"`
}

type NewTeamMessage struct {
	Name         string `json:"name" `
	Introduction string `json:"introduction" `
	Application  int    `json:"application" `
}

//根据teamId返回的team详细信息
type TeamMessageAll struct {
	Name             string   //  -> ->数据库取
	Score            uint     //  -> ->数据库取
	LeaderName       string   //队长
	Introduction     string   //队伍简介 ->数据库取
	Application      int      //是否接受申请 1->接受，-1- // >不接受 ->数据库取
	LsLeader         int      //是否是队长，1->是，-1->不是
	ApplicationUsers []string //申请人列表
}

//获当前用户的队伍信息并判断当前用户是否是leader
func GetTeamMessage(c *gin.Context) {
	var isLeader int              //是否是队长 （1->是，-1->不是)
	var applicationUsers []string //申请该队的用户名切片

	//获取当前用户的id
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	//根据用户的id查询用户所在的team,和其个人身份
	roleTeam := model.RoleTeam{Uid: uid}
	roleTeam.RoleAffirm()
	roleId, teamId := roleTeam.RoleId, roleTeam.TeamId

	//team == 0 无队伍返回空
	if teamId == 0 {
		response.OkWithData(c, gin.H{"team": ""})
		return
	}

	//确认是否是队长
	switch {
	case roleId == 2:
		isLeader = 1
	case roleId == 1:
		isLeader = -1
	}

	//查找用户team的队长姓名
	leaderName, _ := roleTeam.GetLeaderId()
	fmt.Println(leaderName)
	//获取该用户所在队伍的信息
	team := model.Team{}
	team.FindByTeamId(teamId)

	//获取申请该队的用户切片
	teamApplication := model.TeamApplication{TeamId: teamId}
	userApplication := teamApplication.FindNameByTeamId()
	for i := 0; i < len(userApplication); i++ {
		applicationUsers = append(applicationUsers, userApplication[i].Username)
	}

	//封装数据
	teamMessageAll := TeamMessageAll{
		Name:             team.Name,
		Score:            team.Score,
		LeaderName:       leaderName,        //队长
		Introduction:     team.Introduction, //队伍简介 ->数据库取
		Application:      team.Application,  //是否接受申请 1->接受，-1->不接受 ->数据库取
		LsLeader:         isLeader,          //是否是队长，1->是，-1->不是
		ApplicationUsers: applicationUsers,  //申请人名字列表
	}
	response.OkWithData(c, gin.H{"team": teamMessageAll})
}

//申请加入新队伍(添加新的加入队伍申请表)
func AddNewTeam(c *gin.Context) {
	var add RoleTeam
	err := c.ShouldBindJSON(&add)
	if err != nil {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
	//构建新的队伍申请表
	application := model.TeamApplication{Uid: uid, TeamId: add.TeamId}
	roleTeam := model.RoleTeam{Uid: application.Uid}

	//根据Uid判断是否加入或创建过其他队伍
	if roleTeam.IsAlone() {
		response.TeamRoleErr(c)
		return
	}
	//判断是否申请过该队伍防止重复
	if application.AppliedBefore() {
		response.ApplicationAlreadyError(c)
		return
	}
	//判断所加入的队伍是否开放申请
	if !application.IsAllowJoin() {
		response.ApplicationError(c)
		return
	}

	//插入新的成员信息表
	err = application.InsertNew()
	if err == nil {
		response.Ok(c)
		return
	}
	response.ParamError(c)
}

//创建一只队伍
func CreateNewTeam(c *gin.Context) {
	var createTeam CreateTeam
	err := c.ShouldBindJSON(&createTeam)
	if err != nil {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
	//创建新的队伍表(默认允许其他申请加入)
	newTeam := model.Team{Name: createTeam.Name, Score: 0, Application: 1}
	//构建新的team_role表（以队长身份) 这里暂时无teamId
	roleTeam := model.RoleTeam{Uid: uid, RoleId: 2}

	//根据Uid判断是否加入或创建过其他队伍
	if roleTeam.IsAlone() {
		response.TeamRoleErr(c)
		return
	}
	//插入一张新的队伍表,无错误返回teamId
	teamId, err := newTeam.InsertNew()
	if err != nil {
		response.TeamNameExist(c)
		return
	}

	//加入teamId
	roleTeam.TeamId = teamId
	//插入新的成员信息表
	err = roleTeam.InsertNew()
	if err == nil {
		response.Ok(c)
		return
	}

	_ = newTeam.Delete(newTeam.ID)
	response.ParamError(c)
}

//退出或解散该队伍
func ExitTeam(c *gin.Context) {

	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	roleTeam := model.RoleTeam{Uid: uid}
	roleTeam.RoleAffirm()
	if roleTeam.TeamId == 0 {
		response.NotJoinTeamError(c)
		return
	}
	//如果是队长，解散该队伍
	if roleTeam.IsLeader() {
		teamId := roleTeam.TeamId
		//删除该队伍的所有队伍成员信息表
		err := roleTeam.DeleteAllByTeamId()
		if err != nil {
			response.ParamError(c)
			return
		}
		team := model.Team{}
		//删掉该队伍信息
		err = team.Delete(teamId)
		if err != nil {
			response.ParamError(c)
			return
		}
	}
	//不是队长，退队，删掉该成员自己的队伍成员信息表
	err := roleTeam.DeleteByUid()
	if err == nil {
		response.Ok(c)
		return
	}
	response.ParamError(c)
}

//同意申请
func AgreeAdd(c *gin.Context) {
	var teamApplication TeamApplication
	err := c.ShouldBindJSON(&teamApplication)
	if err != nil {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	roleTeam := model.RoleTeam{Uid: uid}
	roleTeam.RoleAffirm()
	roleId, teamId := roleTeam.RoleId, roleTeam.TeamId

	user := model.User{Username: teamApplication.NewUserName}
	user.GetUserMessageByUsername()

	application := model.TeamApplication{Uid: uid}
	application.GetApplicationByUid()

	//判断是否加入的本队（恶意构造表单)
	if teamId != application.TeamId {
		response.NotYourTeamApplicationError(c)
		return
	}
	//判断是否是队长
	if roleId == 2 {
		//不同意申请，有内鬼取消交易
		if teamApplication.AgreeOrNot == 1 {
			//同意申请，开始交易
			roleTeam := model.RoleTeam{Uid: application.Uid, TeamId: teamId, RoleId: 1} //1->队员
			err = roleTeam.InsertNew()
			if err != nil {
				response.ParamError(c)
				return
			}
			err := application.Delete()
			if err != nil {
				response.ParamError(c)
				return
			}
			response.Ok(c)
			return

		}
		//不同意申请，有内鬼取消交易
		err := application.Delete()
		if err == nil {
			response.Ok(c)
			return
		}
		response.ParamError(c)
		return
	}
	//不是队长，权限不足
	response.PermissionError(c)
}

//踢出某人出队伍
func KickPeople(c *gin.Context) {
	var kickUid KickUid
	err := c.ShouldBindJSON(&kickUid)
	if err != nil {
		response.ParamError(c)
		return
	}

	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	nowRoleTeam := model.RoleTeam{Uid: uid}
	//判断是否是队长
	if nowRoleTeam.IsLeader() {
		kickRoleTeam := model.RoleTeam{Uid: kickUid.PoorUid}
		err := kickRoleTeam.DeleteByUid()
		if err != nil {
			response.ParamError(c)
			return
		}
		response.Ok(c)
		return
	}
	//不是队长，权限不足
	response.PermissionError(c)
}

//队伍是否同意申请状态修改
func ApplicationChange(c *gin.Context) {
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	nowRoleTeam := model.RoleTeam{Uid: uid}
	nowRoleTeam.RoleAffirm()
	nowRoleId, teamId := nowRoleTeam.RoleId, nowRoleTeam.TeamId

	//判断是否是队长
	if nowRoleId == 2 {
		applicationChangeTeamTable := model.Team{}
		err := applicationChangeTeamTable.ApplicationChange(teamId)
		if err != nil {
			response.ParamError(c)
			return
		}
		response.Ok(c)
		return
	}
	//不是队长，权限不足
	response.PermissionError(c)
}

//修改队伍信息
func TeamMessageChange(c *gin.Context) {
	var newTeamMessage NewTeamMessage
	err := c.ShouldBindJSON(&newTeamMessage)
	if err != nil ||
		newTeamMessage.Application > 1 || newTeamMessage.Application < -1 {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	nowRoleTeam := model.RoleTeam{Uid: uid}
	nowRoleTeam.RoleAffirm()
	nowRoleId, teamId := nowRoleTeam.RoleId, nowRoleTeam.TeamId

	//判断是否是队长
	if nowRoleId == 2 {
		//获取team原信息
		team := model.Team{}
		team.FindByTeamId(teamId)
		//更新数据
		team.Name = newTeamMessage.Name
		team.Application = newTeamMessage.Application
		team.Introduction = newTeamMessage.Introduction

		err := team.TeamMessageChange()
		if err != nil {
			response.ParamError(c)
			return
		}
		response.Ok(c)
		return
	}
	//不是队长，权限不足
	response.PermissionError(c)
}
