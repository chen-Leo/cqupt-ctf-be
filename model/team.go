//team table
//@author Doc.sao
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
)

//create by Doc.sao

type Team struct {
	gorm.Model
	Name         string
	Score        uint
	Introduction string `gorm:"default:nothing to say"` //队伍简介
	Application  int                                    //是否接受申请 1->接受，-1->不接受
}

type TeamAllMessage struct {
	Name         string
	Score        uint
	Introduction string                                   //队伍简介
	Application  int                                    //是否接受申请 1->接受，-1->不接受
	Members     []string
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

//获取team全部信息及所有队员
func (team *Team) GetTeamMessageAndMember(teamId uint) TeamAllMessage {
	var users []Users
	var members []string

	db.Where("id = ?", teamId).First(&team)
	db.Table("users").Select("users.username,users.id").
		Joins("left join role_team on role_team.uid = users.id").
		Where("role_team.deleted_at  is null AND role_team.team_id = ?", team.ID).
		Find(&users)

	for i := 0; i < len(users); i++ {
		members = append(members, users[i].Username)
	}
	teamAllMessage := TeamAllMessage{
		Name:         team.Name,
		Score:        team.Score,
		Introduction: team.Introduction,
		Application:  team.Application,
		Members:      members,
	}
	return teamAllMessage
}

func (team *Team)FindByPage(page int) (TeamAllMessages []TeamAllMessage, lastPage int) {
	var teams []Team
	var count,firstNum,lastNum int

	db.Table("team").Where("deleted_at is null").Count(&count)

	lastPage = int(math.Ceil(float64(count) / 10.0))
	fmt.Println(lastPage)
    //page小于0 返回第一页
	if page <= 0 {
		page = 1
	}
	if lastPage <= page {
		firstNum = (lastPage - 1) * 10
		lastNum =  count
	} else {
		firstNum = (page - 1) * 10
		lastNum = firstNum + 10
	}

	fmt.Println(firstNum)
	fmt.Println(lastNum)


	db.Order("score").Order("created_at DESC ").Offset(firstNum).Limit(lastNum).Find(&teams)
	totalLength := len(teams)
	for i := 0; i < totalLength; i++ {
		TeamAllMessages = append(TeamAllMessages, TeamAllMessage{
			teams[i].Name,
			teams[i].Score,
			teams[i].Introduction,
			teams[i].Application,
			make([]string, 0), //这里不需要队员名字列表，所以直接置的空
		})
	}
	return
}

