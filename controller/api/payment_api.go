package api

import (
	"go-pos/authenticator"
	"go-pos/controller/apprequest"
	"go-pos/controller/middleware"
	"go-pos/model"
	"go-pos/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentApi struct {
	BaseApi
	publicRoute    *gin.RouterGroup
	paymentUseCase usecase.PaymentUseCase
}

func NewPaymentApi(publicRoute *gin.RouterGroup, paymentUseCase usecase.PaymentUseCase) {
	paymentApi := PaymentApi{
		publicRoute:    publicRoute,
		paymentUseCase: paymentUseCase,
	}
	paymentApi.InitRouter()
}

func (api *PaymentApi) InitRouter() {
	api.publicRoute.POST("", api.CreatePayment)
	api.publicRoute.PUT("/:payments", api.UpdatePayment)
	api.publicRoute.DELETE("/:payments", api.DeletePayment)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("", api.ListPayment)
	api.publicRoute.GET("/:payments", api.DetailPayment)
}

func (api *PaymentApi) ListPayment(c *gin.Context) {
	var meta model.Meta
	var data model.PaymentData
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	subtotal, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	result := api.paymentUseCase.ListPayment(meta.Limit, meta.Skip, subtotal)
	data.Payment = result
	meta.Total = len(result)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *PaymentApi) DetailPayment(c *gin.Context) {
	id := c.Param("paymentId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	result, err := api.paymentUseCase.DetailPayment(data)
	api.Success(c, "Success", result)
	if err != nil {
		api.Error(c, 500, "Server Error")
		return
	}
}

func (api *PaymentApi) CreatePayment(c *gin.Context) {
	var createPayment apprequest.PaymentRequest
	err := c.BindJSON(&createPayment)
	if err != nil {
		api.Error(c, 400, err.Error())
		return
	}
	data, err := api.paymentUseCase.CreatePayment(createPayment)
	if err != nil {
		api.Error(c, 400, err.Error())
		return
	}
	api.Success(c, "Success", data)
}

func (api *PaymentApi) UpdatePayment(c *gin.Context) {
	id := c.Param("paymentId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	var updatePayment apprequest.PaymentRequest
	err = c.ShouldBindJSON(&updatePayment)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}
	err = api.paymentUseCase.UpdatePayment(updatePayment, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}
	api.SuccessNotif(c, "Success")
}

func (api *PaymentApi) DeletePayment(c *gin.Context) {
	id := c.Param("paymentId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	err = api.paymentUseCase.DeletePayment(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}
	api.SuccessNotif(c, "Success")
}
