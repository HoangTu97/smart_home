package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Token struct {
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Roles    []string `json:"roles"`
	*jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userID string, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Token{
		UserID: userID,
		Name: EncodeMD5(username),
		Password: EncodeMD5(password),
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

func (t *Token) GetUserID() (string) {
	return t.UserID
}

func (t *Token) GetUserName() (string) {
	return t.Name
}
