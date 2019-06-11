//加入队伍申请表
//@author doc.sao
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type TeamApplication struct {
	gorm.Model
	Uid    uint
	TeamId uint
}

type UidResult struct {
	uid    uint
	teamId uint
}

func (application *TeamApplication) InsertNew() error {
	err := db.Create(&application)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//判断是否重复申请
func (application *TeamApplication) AppliedBefore() bool {
	db.Where("uid = ? AND team_id = ?", application.Uid, application.TeamId).First(&application)
	if application.ID == 0 {
		return false
	}
	return true
}

//查找所有申请该队的用户名及id
func (application *TeamApplication) FindNameByTeamId() (results []Users) {
	db.Table("users").
		Joins("left join team_application on team_application.uid = users.id").
		Where("team_application.team_id = ?", application.TeamId).Scan(&results)
	return
}

func (application *TeamApplication) GetApplicationByUid() {
	db.Where("uid = ?", application.Uid).First(&application)
}

//同意用户申请
func (application *TeamApplication) AgreeJoin(uid uint, ifAgree int) (int, error) {

	//获取当前登陆用户的队伍id，及其角色id（是否是队长）
	nowRoleTeam := RoleTeam{Uid: uid}
	nowRoleTeam.RoleAffirm()
	roleId, teamId := nowRoleTeam.RoleId, nowRoleTeam.TeamId

	if roleId != 2 {
		err := fmt.Errorf("%s", "you are not the leader")
		return -1, err
	}
	if teamId != application.TeamId {
		err := fmt.Errorf("%s", "the team application do not exist")
		return -2, err
	}

	//开始事务
	tx := db.Begin()

	if ifAgree == 1 {

		checkRoleTeam := RoleTeam{Uid: application.Uid}
		if checkRoleTeam.IsAlone() {
			err := tx.Delete(&application).Error
			if err != nil {
				tx.Rollback()
				return -4, err
			}
			err = fmt.Errorf("%s", "you have joined a team ")
			return -5, err
		}

		newRoleTeam := RoleTeam{Uid: application.Uid, TeamId: teamId, RoleId: 1} //1->队员
		err := newRoleTeam.InsertNew()
		if err != nil {
			tx.Rollback()
			return -3, err
		}

	}

	err := tx.Delete(&application).Error
	if err != nil {
		tx.Rollback()
		return -4, err
	}

	tx.Commit()
	return 0, nil

}

//判断所申请的队伍是否开放申请或是否存在(查的是team表)
func (application *TeamApplication) IsAllowJoin() bool {
	var result Team
	db.Table("team").
		Where("id = ?", application.TeamId).Scan(&result)
	if result.Application == 1 && result.DeletedAt == nil {
		return true
	}
	return false
}

//删除用户申请
func (application *TeamApplication) Delete() error {
	err := db.Where("uid = ? AND team_id = ?", application.Uid, application.TeamId).Delete(&application)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
