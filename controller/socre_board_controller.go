package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"sort"

	"github.com/gin-gonic/gin"
)

type RankEntity struct {
	Name      string
	Motto     string
	Score     uint
	Solved    uint
	Submitted uint
}

type Ranks []RankEntity

//Len()
func (s Ranks) Len() int {
	return len(s)
}

//Less():降序
func (s Ranks) Less(i, j int) bool {
	return s[j].Score < s[i].Score
}

//Swap()
func (s Ranks) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ScoreBoard(c *gin.Context) {
	users := (&model.User{}).FindAll()
	rank := Ranks(make([]RankEntity, len(users)))
	for i := 0; i < len(users); i++ {
		u := users[i]
		solved, submitted, score := u.FindRank()
		r := RankEntity{
			Name:      u.Username,
			Motto:     u.Motto,
			Score:     score,
			Solved:    solved,
			Submitted: submitted}
		rank[i] = r
	}
	sort.Sort(rank)
	response.OkWithData(c, gin.H{"rank": rank})
}
