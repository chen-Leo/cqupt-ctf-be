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
