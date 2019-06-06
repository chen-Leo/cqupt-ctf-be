package model

import (
	"fmt"


	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("mysql", "root:0405duanQWER789@/ctf?charset=utf8mb4&parseTime=True&loc=Local")
	//d, err := gorm.Open("mysql", "root:zxc981201@tcp(127.0.0.1:3306)/ctf?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
	}
	d.DB().SetMaxIdleConns(10)
	d.DB().SetMaxOpenConns(100)
	d.SingularTable(true)
	db = d
}

func Close() {
	_ = db.Close()
}

















