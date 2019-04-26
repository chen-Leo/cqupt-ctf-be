package route

import (
	"cqupt-ctf-be/controller"
	"cqupt-ctf-be/middleware"
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
	route.Use(middleware.CORS)

	route.POST("/login", controller.Login)
	route.POST("/signup", controller.SignUp)

	g:=route.Group("")
	g.Use(middleware.Auth)

	g.GET("/questions", controller.Question)
	g.POST("/submit",controller.Submit)
	g.GET("/rank",controller.ScoreBoard)

	return route
}
