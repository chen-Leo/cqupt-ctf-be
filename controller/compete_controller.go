package controller

import (
	"cqupt-ctf-be/model"
	response "cqupt-ctf-be/utils/response_utils"
	"github.com/gin-gonic/gin"
)

type CompeteReturn struct {
	Name         string
	Introduction string
	CreateTime   string
}

type CompeteReturns [] CompeteReturn

func CompeteAll(c *gin.Context) {
	competes := (&model.Compete{}).FindAll()
	competeReturns := CompeteReturns(make([]CompeteReturn, len(competes)))
	for i := 0; i < len(competes); i++ {
		u := competes[i]
		r := CompeteReturn{
			u.Name,
			u.Introduction,
			u.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		competeReturns[i] = r
	}
	response.OkWithData(c, gin.H{"competes": competeReturns})
}
