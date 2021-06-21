package jwtAuth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"thinknetica/14-api-2/pkg/users"
	"time"
)

type tokenClaims struct {
	UserID int
	Admin  bool
	jwt.StandardClaims
}

func GenerateToken(usr *users.UserInfo) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: usr.ID(),
		Admin:  usr.Admin(),
	})

	key, _ := jwtKey(nil)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}
	return []byte(tokenStr), nil
}

func ParseToken(tokenStr string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &tokenClaims{}, jwtKey)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token is not valid")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("Wrong token claims")
	}
	return claims, nil
}

func jwtKey(token *jwt.Token) (interface{}, error) {
	return []byte("VERY_SECRET_KEY"), nil
}
