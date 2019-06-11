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

	g.POST("/team/create", controller.CreateNewTeam)   //ok
	g.POST("/team/application/add", controller.AddNewTeam)

	g.GET("/team/message", controller.GetTeamMessage)
	g.POST("/team/message",controller.TeamMessageGetByName)

	g.DELETE("/team/exit", controller.ExitTeam) //解散或退出队伍 ok
	g.POST("/team/application/agree", controller.AgreeAdd)
	g.DELETE("/team/kick", controller.KickPeople)

	g.PUT("/team/message/change", controller.TeamMessageChange)
	g.PUT("/team/application/change",controller.ApplicationChange)

	g.POST("/user/message/get", controller.UserMessageGet)
	g.PUT("/user/message/change", controller.UserMessageChange)

	g.PATCH("/user/password", controller.PasswordChange)

	g.GET("/news/get", controller.NewsGetbyPage)

	g.GET("/messageform/get", controller.MessageFormAll)
	g.POST("/messageform/add", controller.MessageFormAdd)

	g.GET("/compete/get", controller.CompeteAll)

	return route
}
