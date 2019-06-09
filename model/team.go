//team table
//@author Doc.sao
package model

import (
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
func (team *Team) InsertNew() (uint, error) {
	err := db.Create(&team)
	if err.Error != nil {
		return 0, err.Error
	}
	return team.ID, nil
}

func (team *Team) Delete(teamId uint) error {
	err := db.Where("id = ? ", teamId).Delete(&team)
	if err.Error != nil {
		return err.Error
	}
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
