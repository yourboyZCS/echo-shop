package common

import (
	"echo_shop/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_create")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string,error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "jkdev.cn",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
//从tokenString中截取出claims，然后返回  （！或对应消息体（payload））
func ParseToken(tokenString string) (*jwt.Token,*Claims,error){
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey,nil
	})
	return token,claims,err
}