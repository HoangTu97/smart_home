package middlewares

import (
	"Food/dto/response"
	"Food/helpers/util"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Security is Security middleware
func Security(c *gin.Context) {

	if regexp.MustCompile(`/api/public/.*`).Match([]byte(c.Request.URL.Path)) {
		c.Next()
		return
	}

	if regexp.MustCompile(`/api/private/.*`).Match([]byte(c.Request.URL.Path)) {
		iUserInfo, exists := c.Get("UserInfo")
		if !exists {
			response.CreateErrorResponse(c, "UNAUTHORIZED")
			c.Abort()
			return
		}

		userInfo := iUserInfo.(*util.Token)
		if err := userInfo.Valid(); err != nil {
			response.CreateErrorResponse(c, "INVALID_TOKEN")
			c.Abort()
			return
		}

		c.Next()
		return
	}

	c.Next()
}
