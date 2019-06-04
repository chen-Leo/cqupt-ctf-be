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

	route.POST("/team/create",controller.CreateNewTeam)
	route.POST("/team/add",controller.AddNewTeam)
	route.DELETE("/team/exite",controller.ExitTeam)
	route.DELETE("/team/break",controller.ExitTeam)
	route.POST("/team/agreeadd",controller.AgreeAdd)

    route.DELETE("/team/kickpeople",controller.KickPeople)
	route.GET("/team/getmessage",controller.GetTeamMessage)
	route.POST("/team/changemessage",controller.TeamMessageChange)

	route.POST("/user/getmessage",controller.UserMessageGet)
	route.POST("/user/changepassword",controller.PasswordChange)
	route.POST("/user/changemessage",controller.UserMessageChange)


	g:=route.Group("")
	g.Use(middleware.Auth)

	g.GET("/questions", controller.Question)
	g.POST("/submit",controller.Submit)
	g.GET("/rank",controller.ScoreBoard)


	return route
}
