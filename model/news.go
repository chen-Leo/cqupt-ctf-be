package model

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title   string
	Content string
}

type NewsReturn struct {
	Title       string
	Content     string
	Number      int
	Time        string
	CurrentPage int
	TotalPage   int
}

//查找数据库中所有的公告，存入redis，并返回封装好的结构体数组,公告个数，存redis可能的错误
func (n *News) FindAll() (newsAllReturn []NewsReturn, totalLength int, err error) {
	var news []*News
	var currentPage int

	db.Order("weight_time DESC").Order("created_at DESC ").Find(&news)
	totalLength = len(news)

	for i := 0; i < totalLength; i++ {
		if i%5 == 0 {
			currentPage = i/5 + 1;
		}
		newsAllReturn = append(newsAllReturn, NewsReturn{
			news[i].Title,
			news[i].Content,
			i + 1,
			news[i].CreatedAt.Format("2006-01-02 15:04:05"),
			currentPage,
			totalLength / 5,
		})
	}

	err = (&Redis{}).Set("newsAll", newsAllReturn, 900)
	return
}
