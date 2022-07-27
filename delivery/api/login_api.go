package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
	"go-pos/usecase"
	"strconv"
)

type LoginApi struct {
	BaseApi
	publicRoute  *gin.RouterGroup
	loginUseCase usecase.LoginUseCase
	jwtUseCase   usecase.JWTUseCase
}

func (api *LoginApi) InitRouter() {
	api.publicRoute.POST("/:cashierId/login", api.Login)
}

func (api *LoginApi) Login(c *gin.Context) {
	id := c.Param("cashierId")
	intId, _ := strconv.Atoi(id)
	cashierId := uint(intId)
	var credential apprequest.LoginCredentials
	err := c.BindJSON(&credential)
	if err != nil {
		api.Error(c, "Wrong input")
		return
	}
	isUserAuthenticated := api.loginUseCase.LoginUser(cashierId, credential.Passcode)
	if !isUserAuthenticated {
		//c.AbortWithStatusJSON(http.StatusUnauthorized, "Passcode does not match")
		api.Error(c, "Passcode does not match")
		return
	}
	token := api.jwtUseCase.GenerateToken(id, true)
	api.Success(c, "Success", token)
}

func NewLoginApi(publicRoute *gin.RouterGroup, loginUseCase usecase.LoginUseCase, jwtUseCase usecase.JWTUseCase) {
	loginApi := LoginApi{
		publicRoute:  publicRoute,
		loginUseCase: loginUseCase,
		jwtUseCase:   jwtUseCase,
	}
	loginApi.InitRouter()
}
