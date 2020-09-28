package security

import (
	"Food/helpers/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Token struct {
	UserID string              `json:"userId"`
	Name   string              `json:"name"`
	Roles  []UserRole `json:"roles"`
	*jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userID string, username string, roles []UserRole) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Token{
		UserID: userID,
		Name:   util.EncodeMD5(username),
		Roles:  roles,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Token); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (t *Token) GetUserID() string {
	return t.UserID
}

func (t *Token) GetUserName() string {
	return t.Name
}
