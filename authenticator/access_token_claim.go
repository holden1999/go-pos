package authenticator

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	jwt.StandardClaims
	Id       int
	Passcode string `json:"passcode"`
}
