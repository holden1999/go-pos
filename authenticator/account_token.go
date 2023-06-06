package authenticator

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type MiddlewareToken interface {
	GenerateToken() (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

type TokenConfig struct {
	secretKey string
	issue     string
}

func (t *TokenConfig) GenerateToken() (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Subject:  "1",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(t.secretKey))
	return ss, err
}

func (t *TokenConfig) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secretKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, err
	}
	return nil, err
}

func NewTokenConfig() TokenConfig {
	return TokenConfig{
		secretKey: getSecretKey(),
		issue:     "Holden",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
