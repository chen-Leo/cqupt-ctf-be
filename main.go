package main

import "cqupt-ctf-be/route"

func main(){
	route:=route.SetupRoute()
	route.Run(":1234")
	}