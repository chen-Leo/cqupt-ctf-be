//加入队伍申请表
//@author doc.sao
package model

import (
	"github.com/jinzhu/gorm"
)

type TeamApplication struct {
	gorm.Model
	Uid uint
	TeamId uint
}

type UidResult struct {
	uid uint
	teamId uint
}


func (application *TeamApplication) InsertNew() error{
	err := db.Create(&application )
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//删除用户申请
func (application *TeamApplication) Delete() error{
	err := db.Where("uid = ? AND team_id = ?",application.Uid,application.TeamId).Delete(&application)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

//判断是否重复申请
func (application *TeamApplication) AppliedBefore() bool{

	db.Where("uid = ? AND team_id = ?",application.Uid,application.TeamId).First(&application)
	if application.ID == 0 {
		return false
	}
     return true
}

//查找所有申请该队的用户名及id
func (application *TeamApplication) FindNameByTeamId() (results []User){
	db.Table("user").
		Joins("left join team_application on team_application.uid = user.id").
		Where("team_application.team_id = ?",application.TeamId).Scan(&results)
	return
}

//判断所申请的队伍是否开放申请(查的是team表)
func (application *TeamApplication) IsAllowJoin() bool{
	var result Team
	db.Table("team").
		Where("id = ?",application.TeamId).Scan(&result)
	if result.Application == 1 && result.DeletedAt == nil {
		return true
	}
	return false
}


func (application *TeamApplication) GetApplicationByUid() {
	db.Where("uid = ?",application.Uid).First(&application)
}





