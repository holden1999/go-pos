package authenticator

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"go-pos/model"
	"log"
	"time"
)

type Token interface {
	CreateAccessToken(cred *model.Credential) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type token struct {
	Config TokenConfig
}

func (t token) CreateAccessToken(cred *model.Credential) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.Config.AccessTokenLifeTime)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.Config.ApplicationName,
		},
		Id:       cred.CashierId,
		Passcode: cred.Passcode,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		t.Config.JwtSigningMethod,
		claims,
	)
	return token.SignedString([]byte(t.Config.JwtSignatureKey))
}

func (t token) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != t.Config.JwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return []byte(t.Config.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != t.Config.ApplicationName {
		log.Println("Token Invalid")
		return nil, err
	}
	return claims, nil
}

func NewTokenService(config TokenConfig) Token {
	return &token{Config: config}
}
