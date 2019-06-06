package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type NewsGet struct {
	Page int `json:"page" binding:"required"`
}

//从数据库中找出所有的公告并存入redis返回前端第一页（5个)
func GetNews(c *gin.Context) {

	newsAllJson, totalLength, err := (&model.News{}).FindAll()
	if err != nil {
		response.RedisError(c)
		return
	}
	res := make([]gin.H, totalLength)
	if totalLength >= 5 {
		for i := 0; i < 5; i++ {
			res[i] = gin.H{
				"content":      newsAllJson[i].Content,
				"title":        newsAllJson[i].Title,
				"number":       newsAllJson[i].Number,
				"currentPage ": newsAllJson[i].CurrentPage,
				"TotalPage":    newsAllJson[i].TotalPage,
			}
		}
		response.OkWithArray(c, res)
		return
	}
	for i := 0; i < totalLength; i++ {
		res[i] = gin.H{
			"content":      newsAllJson[i].Content,
			"title":        newsAllJson[i].Title,
			"number":       newsAllJson[i].Number,
			"currentPage ": newsAllJson[i].CurrentPage,
			"TotalPage":    newsAllJson[i].TotalPage,
		}
	}
	response.OkWithArray(c, res)
	return

}

func GetNewsbyPage(c *gin.Context) {
	var newGet NewsGet
	err := c.ShouldBindJSON(&newGet)
	if err != nil || newGet.Page <= 0 {
		response.ParamError(c)
		return
	}

	var newsAllJson []model.NewsReturn
	var totalLength int

	//从redis取数据，没缓存从MySQL取
	if !(&model.Redis{}).Exists("newsAll") {
		newsAllJson, totalLength, err = (&model.News{}).FindAll()
		if err != nil {
			response.RedisError(c)
			return
		}
	} else {
		date, _ := (&model.Redis{}).Get("newsAll")
		err = json.Unmarshal(date, newsAllJson)
		if err != nil {
			response.RedisError(c)
			return
		}
		totalLength = len(newsAllJson)
	}

	var firstNum, lastNum int

	if totalLength/5 <= newGet.Page {
		firstNum = (totalLength/5 - 1) * 5
		lastNum = totalLength
	} else {
		firstNum = (newGet.Page/5 - 1) * 5
		lastNum = lastNum + 5
	}

	res := make([]gin.H, 5)

	for i := firstNum; i < lastNum; i++ {
		res[i] = gin.H{
			"content":      newsAllJson[i].Content,
			"title":        newsAllJson[i].Title,
			"number":       newsAllJson[i].Number,
			"currentPage ": newsAllJson[i].CurrentPage,
			"TotalPage":    newsAllJson[i].TotalPage,
		}
	}
	response.OkWithArray(c, res)
	return
}
