package middleware

import (
	"github.com/gin-gonic/gin"
	"go-pos/authenticator"
	"strings"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware struct {
	accToken authenticator.MiddlewareToken
}

func (a *AuthTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.BindHeader(&h)
		if err != nil {
			c.AbortWithStatusJSON(401, "No token")
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer", "", -1)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, "No token")
			return
		}
		token, _ := a.accToken.ValidateToken(tokenString)
		c.JSON(200, token)
	}
}

func NewTokenValidator(accToken authenticator.MiddlewareToken) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{accToken: accToken}
}
