package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/controller/apprequest"
	"go-pos/model"
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
	if id == ":cashierId" {
		api.Error(c, 404, "ID empty")
		return
	}
	data, _ := strconv.Atoi(id)
	result := api.loginUseCase.GetPasscode(data)
	if result.Passcode == "" {
		api.Error(c, 404, "Cashier Not Found")
		return
	}
	api.Success(c, "Success", result)
}

func (api *LoginApi) VerifyLogin(c *gin.Context) {
	var (
		credential apprequest.LoginCredentials
		result     model.LoginResp
	)
	id := c.Param("cashierId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	cashierId := uint(intId)
	c.BindJSON(&credential)
	isUserAuthenticated := api.loginUseCase.LoginUser(cashierId, credential.Passcode)
	if !isUserAuthenticated {
		api.Error(c, 401, "Passcode Not Match")
		return
	}
	token, _ := api.jwtUseCase.GenerateToken()
	result.Token = token
	api.Success(c, "Success", result)
}

func NewLoginApi(publicRoute *gin.RouterGroup, loginUseCase usecase.LoginUseCase, jwtUseCase usecase.JWTUseCase) {
	loginApi := LoginApi{
		publicRoute:  publicRoute,
		loginUseCase: loginUseCase,
		jwtUseCase:   jwtUseCase,
	}
	loginApi.InitRouter()
}
