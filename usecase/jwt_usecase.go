package usecase

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type JWTUseCase interface {
	GenerateToken() (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

type JwtUseCase struct {
	secretKey string
	issue     string
}

func (j *JwtUseCase) GenerateToken() (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Subject:  "1",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.secretKey))
	return ss, err
}

func (j *JwtUseCase) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, err
	}
	return nil, err
}

func NewJWTUseCase() JWTUseCase {
	return &JwtUseCase{
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
