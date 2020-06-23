package middlewares

import (
	"Food/dto/response"
	"Food/pkg/e"
	"Food/pkg/util"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(c *gin.Context) {
	var code int

	code = e.SUCCESS
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		response.CreateErrorResponse(c, e.GetMsg(e.INVALID_PARAMS))
		c.Abort()
		return
	}
	token = strings.Fields(token)[1]

	_, err := util.ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
	}

	if code != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(code))
		c.Abort()
		return
	}

	c.Next()
}
