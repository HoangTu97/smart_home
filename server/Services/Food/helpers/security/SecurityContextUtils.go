package security

import (
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) *Token {
	return c.MustGet("UserInfo").(*Token)
}

func GetUserName(c *gin.Context) string {
	return GetUserInfo(c).GetUserName()
}

func GetUserID(c *gin.Context) string {
	return GetUserInfo(c).GetUserID()
}
