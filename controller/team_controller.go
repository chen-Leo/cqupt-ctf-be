//team controller
//@author doc.sao
package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//获取队伍名
type TeamName struct {
	TeamName string `json:"teamname" binding:"required"`
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

type KickName struct {
	PoorName string `json:"poorname" binding:"required"`
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
	Members          []string //队员名字
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
	switch roleId {
	case 2:
		isLeader = 1
	case 1:
		isLeader = -1
	}
	//查找用户team的队长姓名
	leaderName, _ := roleTeam.GetLeaderId()
	//获取该用户所在队伍的信息
	team := model.Team{}
	//获取该队的其他队员
	teamAllMessage := team.GetTeamMessageAndMember(teamId)
	//获取申请该队的用户切片
	teamApplication := model.TeamApplication{TeamId: teamId}
	userApplication := teamApplication.FindNameByTeamId()

	for i := 0; i < len(userApplication); i++ {
		applicationUsers = append(applicationUsers, userApplication[i].Username)
	}
	//封装数据
	teamMessageAll := TeamMessageAll{
		Name:             teamAllMessage.Name,
		Score:            teamAllMessage.Score,
		LeaderName:       leaderName,                  //队长
		Introduction:     teamAllMessage.Introduction, //队伍简介 ->数据库取
		Application:      teamAllMessage.Application,  //是否接受申请 1->接受，-1->不接受 ->数据库取
		LsLeader:         isLeader,                    //是否是队长，1->是，-1->不是
		ApplicationUsers: applicationUsers,            //申请人名字列表
		Members:          teamAllMessage.Members,
	}
	response.OkWithData(c, gin.H{"team": teamMessageAll})
}

//申请加入新队伍(添加新的加入队伍申请表)
func AddNewTeam(c *gin.Context) {
	var add TeamName
	err := c.ShouldBindJSON(&add)
	if err != nil {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)

	//获取队伍信息
	team := model.Team{Name: add.TeamName}
	team.FindByTeamName()
	//构建新的队伍申请表
	application := model.TeamApplication{Uid: uid, TeamId: team.ID}
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

	//判断所加入的队伍是否开放申请或者是否存在
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

	//创建队伍，errInt为错误参数
	errInt, _ := newTeam.InsertNew(uid)
	switch errInt {
	case -1:
		response.TeamNameExist(c)
	case -2:
		response.TeamRoleErr(c)
	case -3:
		response.ParamError(c)
	case 0:
		response.OkWithData(c, gin.H{
			"name":         newTeam.Name,
			"score":        newTeam.Score,
			"introduction": newTeam.Introduction,
			"application":  newTeam.Application})
	}
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
	if roleTeam.IsLeader(){
		teamId := roleTeam.TeamId
		//删掉该队伍信息
		err := (&model.Team{}).Delete(teamId)
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

	//获取申请者的信息
	user := model.Users{Username: teamApplication.NewUserName}
	user.GetUserMessageByUsername()
	//获取申请表信息
	application := model.TeamApplication{Uid: user.ID}
	application.GetApplicationByUid()

	errInt, _ := application.AgreeJoin(uid, teamApplication.AgreeOrNot)
	switch errInt {
	case -1:
		response.PermissionError(c)
	case -2:
		response.TeamApplicationNotExist(c)
	case -3:
		response.ParamError(c)
	case -4:
		response.ParamError(c)
	case -5:
		response.TeamRoleErr(c)
	case 0:
		response.Ok(c)
	}
	return
}

//踢出某人出队伍
func KickPeople(c *gin.Context) {
	var kickName KickName
	err := c.ShouldBindJSON(&kickName)
	if err != nil {
		response.ParamError(c)
		return
	}
	//获取用户uid
	uidInterface, _ := c.Get("uid")
	uid := uidInterface.(uint)
	//获取队员信息
	poorUser := model.Users{Username: kickName.PoorName}
	poorUser.GetUserMessageByUsername()

	nowRoleTeam := model.RoleTeam{Uid: uid}
	nowRoleTeam.RoleAffirm()
	if poorUser.ID == 0 {
		response.UserNameNotExist(c)
		return
	}
	if poorUser.ID == uid {
		response.KickYourself(c)
		return
	}
	//判断是否是队长
	fmt.Println(nowRoleTeam)
	if nowRoleTeam.RoleId == 2  {
		kickRoleTeam := model.RoleTeam{Uid: poorUser.ID}
		kickRoleTeam.RoleAffirm()
		if nowRoleTeam.TeamId == kickRoleTeam.TeamId {
			err := kickRoleTeam.DeleteByUid()
			if err != nil {
				response.ParamError(c)
				return
			}
			response.Ok(c)
			return
		}
		response.NotYourMember(c)
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
		if newTeamMessage.Name != "" {
			team.Name = newTeamMessage.Name
		}
		if newTeamMessage.Application != 0 {
			team.Application = newTeamMessage.Application
		}
		if newTeamMessage.Introduction != "" {
			team.Introduction = newTeamMessage.Introduction
		}

		err := team.TeamMessageChange()
		if err != nil {
			response.TeamNameExist(c)
			return
		}
		teamAllMessage := team.GetTeamMessageAndMember(uid)
		response.OkWithData(c,gin.H{"team":teamAllMessage})
		return
	}
	//不是队长，权限不足
	response.PermissionError(c)
}



func TeamMessageGetByName(c *gin.Context) {
	var TeamName TeamName
	err := c.ShouldBindJSON(&TeamName)
	if err != nil {
		response.ParamError(c)
		return
	}
	team := model.Team{Name: TeamName.TeamName}
	team.FindByTeamName()
	teamAllMessage := (&model.Team{}).GetTeamMessageAndMember(team.ID)
	response.OkWithData(c, gin.H{"team": teamAllMessage})

}