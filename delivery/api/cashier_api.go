package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"go-pos/usecase"
	"strconv"
)

type CashierApi struct {
	BaseApi
	publicRoute    *gin.RouterGroup
	cashierUseCase usecase.CashierUseCase
}

func (api *CashierApi) InitRouter() {
	api.publicRoute.GET("", api.ListCashier)
	api.publicRoute.GET("/:cashierId", api.DetailCashier)
	api.publicRoute.POST("", api.CreateCashier)
	api.publicRoute.PUT("/:cashierId", api.UpdateCashier)
	api.publicRoute.DELETE("", api.DeleteCashier)
}

func (api *CashierApi) ListCashier(c *gin.Context) {
	var meta model.Meta
	var data model.CashierData
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	result := api.cashierUseCase.ListCashier(meta.Limit, meta.Skip)
	data.Cashiers = result
	meta.Total = len(result)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *CashierApi) CreateCashier(c *gin.Context) {
	var createCashier apprequest.Cashier
	c.BindJSON(&createCashier)
	data, err := api.cashierUseCase.CreateCashier(createCashier)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.Success(c, "Success", data)
}

func (api *CashierApi) DetailCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	result, err := api.cashierUseCase.DetailCashier(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.Success(c, "Success", result)
}

func (api *CashierApi) UpdateCashier(c *gin.Context) {
	id := c.Param("cashierId")
	if id == "" {
		c.AbortWithStatusJSON(404, "ID doesn't exist")
	}
	data, _ := strconv.Atoi(id)
	var updateCashier apprequest.Cashier
	c.BindJSON(&updateCashier)
	err := api.cashierUseCase.UpdateCashier(updateCashier, data)
	if err != nil {
		c.AbortWithStatusJSON(200, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func (api *CashierApi) DeleteCashier(c *gin.Context) {
	id := c.Param("cashierId")
	if id == "" {
		c.AbortWithStatusJSON(404, "ID doesn't exist")
	}
	data, _ := strconv.Atoi(id)
	err := api.cashierUseCase.DeleteCashier(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func NewCashierApi(publicRoute *gin.RouterGroup, cashierUseCase usecase.CashierUseCase) *CashierApi {
	cashierApi := CashierApi{
		publicRoute:    publicRoute,
		cashierUseCase: cashierUseCase,
	}
	cashierApi.InitRouter()
	return &cashierApi
}
