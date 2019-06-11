//team table
//@author Doc.sao
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//create by Doc.sao

type Team struct {
	gorm.Model
	Name         string
	Score        uint
	Introduction string `gorm:"default:nothing to say"` //队伍简介
	Application  int                                    //是否接受申请 1->接受，-1->不接受
}





//创建了一只新队伍
func (team *Team) InsertNew(uid uint) (int, error) {
	//开启事务
	tx := db.Begin()
	//创建队伍
	err := tx.Create(&team)
	if err.Error != nil {
		tx.Rollback()
		return -1, err.Error
	}
	//根据Uid判断是否加入或创建过其他队伍
	roleTeam := RoleTeam{Uid: uid, RoleId: 2}
	if roleTeam.IsAlone() {
		tx.Rollback()
		err := fmt.Errorf("%s", "you joined a team before")
		return -2, err
	}
	//加入角色对应表
	roleTeam.TeamId = team.ID
	err = tx.Create(&roleTeam)
	if err.Error != nil {
		tx.Rollback()
		return -3, err.Error
	}
	tx.Commit()
	return 0, nil
}

//删除队伍
func (team *Team) Delete(teamId uint) error {
	//删队伍
	//删队伍角色信息表
	//开启事务
	tx := db.Begin()
	err :=   tx.Table("role_team").Where("team_id = ?",teamId).Delete(RoleTeam{})
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	err = tx.Where("id = ? ", teamId).Delete(&team)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	tx.Commit()
	return nil
}

//查询所有队伍
func (team *Team) FindAll(teams []*Team) {
	db.Find(&teams)
	return
}

//通过teamId查询某一队伍信息
func (team *Team) FindByTeamId(teamId uint) {
	db.Where("id = ?", teamId).First(&team)
}

//通过teamname查询某一队伍信息
func (team *Team) FindByTeamName() {
	db.Where("name = ?", team.Name).First(&team)
}


//修改队伍是否允许申请的状态
func (team *Team) ApplicationChange(teamId uint) error {
	err := db.Model(&team).Where("id = ?", teamId).UpdateColumn("application", gorm.Expr("application * ?", -1))
	return err.Error
}

//队伍信息修改
func (team *Team) TeamMessageChange() error {
	err := db.Model(&team).Updates(team)
	return err.Error
}




