package response_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  10001,
		"message": "param error",
		"time":    time.Now(),
	})
}

func OkWithData(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10000,
		"message": "success",
		"time":    time.Now(),
		"data":    data,
	})
}

func OkWithArray(c *gin.Context, data []gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10000,
		"message": "success",
		"time":    time.Now(),
		"data":    data,
	})
}

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10000,
		"message": "success",
		"time":    time.Now(),
	})
}

func UsernameOrEmailExist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10012,
		"message": "username exist or or email is used",
		"time":    time.Now(),
	})
}

//
func PasswordError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10013,
		"message": "change password error, old password error",
		"time":    time.Now(),
	})
}

func FlagErr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10021,
		"message": "flag error",
		"time":    time.Now(),
	})
}

func IsSolved(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10022,
		"message": "is solved",
		"time":    time.Now(),
	})
}

func AuthErr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10031,
		"message": "not login",
		"time":    time.Now(),
	})
}

//加入或创建新队伍前不允许必须以前是单身
//create by sao
func TeamRoleErr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10041,
		"message": "error,you already join a team, you can not join or create other team",
		"time":    time.Now(),
	})
}

//队伍名已存在
//create by sao
func TeamNameExist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10042,
		"message": "team name exist",
		"time":    time.Now(),
	})
}

//你不是队长，权限不足
//create by sao
func PermissionError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10043,
		"message": "you are not the team leader ",
		"time":    time.Now(),
	})
}

//队伍未开放加入申请
//create by sao
func ApplicationError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10044,
		"message": "the team is not open the application or team is not exit",
		"time":    time.Now(),
	})
}
func NotJoinTeamError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10045,
		"message": "you do not have a team",
		"time":    time.Now(),
	})
}

//create by sao
func ApplicationAlreadyError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10046,
		"message": "Application Already before",
		"time":    time.Now(),
	})
}

//create by sao
func TeamApplicationNotExist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10047,
		"message": "the team application do not exist ",
		"time":    time.Now(),
	})
}

func RedisError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  10051,
		"message": "redis error",
		"time":    time.Now(),
	})
}

func MessageError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  10061,
		"message": "leave message error,the message doesn't exist.",
		"time":    time.Now(),
	})
}
