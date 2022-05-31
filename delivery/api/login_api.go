package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
	"go-pos/usecase"
)

type LoginApi interface {
	Login(c *gin.Context) string
}

type loginApi struct {
	loginUseCase usecase.LoginUseCase
	jwtUseCase   usecase.JWTUseCase
}

func (l *loginApi) Login(c *gin.Context) string {
	var credential apprequest.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := l.loginUseCase.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return l.jwtUseCase.GenerateToken(credential.Email, true)
	}
	return ""
}

func NewLoginApi(loginUseCase usecase.LoginUseCase, jwtUseCase usecase.JWTUseCase) LoginApi {
	return &loginApi{
		loginUseCase: loginUseCase,
		jwtUseCase:   jwtUseCase,
	}
}
