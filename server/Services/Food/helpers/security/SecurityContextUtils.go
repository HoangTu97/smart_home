package security

import (
	"Food/helpers/util"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) *util.Token {
	return c.MustGet("UserInfo").(*util.Token)
}

func GetUserName(c *gin.Context) string {
	return GetUserInfo(c).GetUserName()
}

func GetUserID(c *gin.Context) string {
	return GetUserInfo(c).GetUserID()
}