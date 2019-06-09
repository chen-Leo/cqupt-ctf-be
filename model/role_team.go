//roleTeam table
//@author Doc.sao
package model

import (
	"github.com/jinzhu/gorm"
)

type RoleTeam struct {
	gorm.Model
	Uid    uint
	TeamId uint
	RoleId uint
}

type Result struct {
	RoleId uint
	TeamId uint
}

//创建一只队伍或者加入新的队伍 ( roleId->2 以队长身份创建队伍 || roleId->1 以队员身份加入队伍  )
func (roleTeam *RoleTeam) InsertNew() error {
	err := db.Create(roleTeam)
	if err != nil {
		return err.Error
	}
	return nil;
}

//查询角色是否加入过其他队伍
func (roleTeam *RoleTeam) IsAlone() bool {
	db.Where("uid = ? ", roleTeam.Uid).First(&roleTeam)
	if roleTeam.TeamId == 0 {
		return false
	}
	return true
}

//查询角色是某队队伍的队长
func (roleTeam *RoleTeam) IsLeader() bool {
	var count int
	db.Table("role_team").Where("uid = ?  AND role_id = 2", roleTeam.Uid).Count(&count)
	if count == 0 {
		return false
	}
	return true
}

//删除一条队伍角色记录
func (roleTeam *RoleTeam) DeleteByUid() error {
	err := db.Where("uid = ?", roleTeam.Uid).Delete(*roleTeam)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//删除所有team_id = * 的 队伍角色记录
func (roleTeam *RoleTeam) DeleteAllByTeamId() error {
	err := db.Where("team_id = ?", roleTeam.TeamId).Delete(RoleTeam{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//查询用户队员角色身份的改进版
func (roleTeam *RoleTeam) RoleAffirm() {
	db.Select("role_id,team_id").Where("uid = ? ", roleTeam.Uid).First(&roleTeam)
}

//获取队长名字及id
func (roleTeam *RoleTeam) GetLeaderId() (string, uint) {
	var leader User
	db.Table("user").Select("user.username,user.id").
		Joins("left join role_team on role_team.uid = user.id").
		Where("role_team.role_id = 2 AND role_team.team_id = ?", roleTeam.TeamId).
		Scan(&leader)
	return leader.Username, leader.ID
}
