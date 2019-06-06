package middleware

import (
	"cqupt-ctf-be/utils/jwt_utils"
	response "cqupt-ctf-be/utils/response_utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	jwtStr := c.GetHeader("Authorization")
	jwtStr = strings.Replace(jwtStr, "Bearer ", "", 7)
	u, err := jwt_utils.ParseToken(jwtStr)
	if err == nil {

		if u.Uid != 0 {
			c.Set("uid", u.Uid)
			c.Next()
			return
		}
	}
	c.Abort()
	response.AuthErr(c)
}
