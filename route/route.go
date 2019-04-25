package route

import (
	"cqupt-ctf-be/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	route = gin.Default()
	store := cookie.NewStore([]byte("SessionId"))
	route.Use(sessions.Sessions("session", store))

}

func SetupRoute() *gin.Engine {
	route.POST("/login", controller.Login)
	route.POST("/signup", controller.SignUp)

	route.GET("/questions", controller.Question)
	return route
}
