package route

import (
	"github.com/gin-gonic/gin"
	"cqupt-ctf-be/controller"
)

var route *gin.Engine

func init(){
	gin.SetMode(gin.ReleaseMode)
	route =gin.Default()

}

func SetupRoute() *gin.Engine{
	route.POST("/login", controller.Login)
	return route
}
