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

	g := route.Group("")
	g.Use(middleware.Auth)

	g.GET("/questions", controller.Question)
	g.POST("/submit", controller.Submit)
	g.GET("/rank", controller.ScoreBoard)

	g.POST("/team/create", controller.CreateNewTeam)
	g.POST("/team/add", controller.AddNewTeam)
	g.GET("/team/getmessage", controller.GetTeamMessage)

	g.DELETE("/team/exite", controller.ExitTeam) //解散或退出队伍

	g.POST("/team/agreeadd", controller.AgreeAdd)

	g.DELETE("/team/kickpeople", controller.KickPeople)

	g.POST("/team/changemessage", controller.TeamMessageChange)

	g.POST("/user/getmessage", controller.UserMessageGet)
	g.POST("/user/changemessage", controller.UserMessageChange)
	g.POST("/user/changepassword", controller.PasswordChange)

	g.GET("/news/get", controller.NewsGetbyPage)

	g.GET("/messageform/get", controller.MessageFormAll)
	g.POST("/messageform/add", controller.MessageFormAdd)

	g.GET("/compete/get", controller.CompeteAll)
	return route
}
