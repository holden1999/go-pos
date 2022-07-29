package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/authenticator"
	"go-pos/delivery/apprequest"
	"go-pos/delivery/middleware"
	"go-pos/model"
	"go-pos/usecase"
	"strconv"
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
	api.publicRoute.POST("", api.createPayment)
	api.publicRoute.PUT("/:payments", api.updatePayment)
	api.publicRoute.DELETE("", api.deletePayment)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("", api.listPayment)
	api.publicRoute.GET("/:payments", api.detailPayment)
}

func (api *PaymentApi) listPayment(c *gin.Context) {
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

func (api *PaymentApi) detailPayment(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	result := api.paymentUseCase.DetailPayment(data)
	api.Success(c, "Success", result)
}

func (api *PaymentApi) createPayment(c *gin.Context) {
	var createPayment apprequest.PaymentRequest
	err := c.ShouldBindJSON(&createPayment)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	data, err := api.paymentUseCase.CreatePayment(createPayment)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.Success(c, "Success", data)
}

func (api *PaymentApi) updatePayment(c *gin.Context) {
	id := c.Param("paymentId")
	data, _ := strconv.Atoi(id)
	var updatePayment apprequest.PaymentRequest
	err := c.ShouldBindJSON(&updatePayment)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = api.paymentUseCase.UpdatePayment(updatePayment, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func (api *PaymentApi) deletePayment(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	err := api.paymentUseCase.DeletePayment(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}
