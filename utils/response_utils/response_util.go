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

func UsernameExist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10012,
		"message": "username exist",
		"time":    time.Now(),
	})
}

func FlagErr(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status":  10021,
		"message": "flag error",
		"time":    time.Now(),
	})
}

func IsSolved(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status":  10022,
		"message": "is solved",
		"time":    time.Now(),
	})
}


func AuthErr(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status":  10031,
		"message": "not login",
		"time":    time.Now(),
	})
}
