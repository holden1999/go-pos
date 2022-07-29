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
	api.publicRoute.POST("/:cashierId/login", api.VerifyLogin)
	api.publicRoute.GET("/:cashierId/passcode", api.Passcode)
}

func (api *LoginApi) Passcode(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	result := api.loginUseCase.GetPasscode(data)
	api.Success(c, "Success", result)
}

func (api *LoginApi) VerifyLogin(c *gin.Context) {
	id := c.Param("cashierId")
	intId, _ := strconv.Atoi(id)
	cashierId := uint(intId)
	var credential apprequest.LoginCredentials
	err := c.BindJSON(&credential)
	if err != nil {
		api.Error(c, "Passcode does not match")
		return
	}
	isUserAuthenticated := api.loginUseCase.LoginUser(cashierId, credential.Passcode)
	if !isUserAuthenticated {
		api.Error(c, "Passcode does not match")
		return
	}
	token, _ := api.jwtUseCase.GenerateToken()
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
