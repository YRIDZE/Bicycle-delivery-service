package helper

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	jwt.StandardClaims
	ID  int32  `json:"id"`
	UID string `json:"uid"`
}
