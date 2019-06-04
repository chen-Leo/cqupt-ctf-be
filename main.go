package main

import (
	"cqupt-ctf-be/model"
	"cqupt-ctf-be/route"
	"fmt"
)

func main() {
	route := route.SetupRoute()
	defer model.Close()
	err := route.Run(":8888")
	if err != nil {
		fmt.Println(err.Error())
	}
}
