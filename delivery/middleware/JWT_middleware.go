package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-pos/authenticator"
	"go-pos/usecase"
	"net/http"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware struct {
	accToken authenticator.Token
}

func (a *AuthTokenMiddleware) AuthorizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := usecase.NewJWTUseCase().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func NewTokenValidator(accToken authenticator.Token) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{accToken: accToken}
}
