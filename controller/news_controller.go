package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type NewsGet struct {
	Page int `json:"page" binding:"required"`
}

func NewsGetbyPage(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		response.ParamError(c)
		return
	}

	var newsAllJson []model.NewsReturn
	var totalLength int
	var firstNum, lastNum, lastPage int

	//当page为负数时，返回第一页的内容
	if page <= 0 {
		page = 1
	}
	//从redis取数据，没缓存从MySQL取
	if !(&model.Redis{}).Exists("newsAll") {
		newsAllJson, totalLength, err = (&model.News{}).FindAll()
		if err != nil {
			response.RedisError(c)
			return
		}
	} else {
		data, _ := (&model.Redis{}).Get("newsAll")
		_ = json.Unmarshal(data, &newsAllJson)
		totalLength = len(newsAllJson)
	}

	lastPage = int(math.Ceil(float64(totalLength) / 5.0))
	if lastPage <= page {
		firstNum = (lastPage - 1) * 5
		lastNum = totalLength
	} else
	{
		firstNum = (page - 1) * 5
		lastNum = firstNum + 5
	}

	if firstNum < 0 {
		firstNum = 0
	}
	res := make([]gin.H, lastNum-firstNum)
	//封装返回数据
	j := 0
	for i := firstNum; i < lastNum; i++ {
		res[j] = gin.H{
			"title":        newsAllJson[i].Title,
			"content":      newsAllJson[i].Content,
			"time":         newsAllJson[i].Time,
			"number":       newsAllJson[i].Number,
			"currentPage ": newsAllJson[i].CurrentPage,
			"totalPage":    newsAllJson[i].TotalPage,
		}
		j++
	}
	response.OkWithArray(c, res)
	return
}
