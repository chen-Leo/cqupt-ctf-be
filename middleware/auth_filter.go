package middleware

import (
	response "cqupt-ctf-be/utils/response_utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	s := sessions.Default(c)
	if s.Get("uid") != nil {
		c.Next()
		return
	}
	c.Abort()
	response.AuthErr(c)
}
