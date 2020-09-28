package middlewares

import (
	"Food/helpers/e"
	"Food/helpers/security"
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
		c.Next()
		return
	}
	token = strings.Fields(token)[1]

	parsedToken, err := security.ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
	}

	if code != e.SUCCESS {
		c.Next()
		return
	}

	c.Set("UserInfo", parsedToken)

	c.Next()
}
