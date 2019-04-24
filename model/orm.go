package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("mysql", "root:123456@/ctf?charset=utf8&parseTime=True&loc=Local")
	if (err != nil) {
		fmt.Println(err.Error())
	}
	d.DB().SetMaxIdleConns(10)
	d.DB().SetMaxOpenConns(100)
	d.SingularTable(true)
	db=d
}
