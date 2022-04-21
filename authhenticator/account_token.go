package authhenticator

import "github.com/golang-jwt/jwt"

type Token interface {
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type TokenConfig struct {
	ApplicationName string
}
