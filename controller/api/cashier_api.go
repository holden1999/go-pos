package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/controller/apprequest"
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
	api.publicRoute.DELETE("/:cashierId", api.DeleteCashier)
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
		api.Error(c, 400, "Incomplete data")
		return
	}
	api.Success(c, "Success", data)
}

func (api *CashierApi) DetailCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	result, err := api.cashierUseCase.DetailCashier(data)
	if err != nil {
		api.Error(c, 400, "Error detail cashier")
		return
	}
	api.Success(c, "Success", result)
}

func (api *CashierApi) UpdateCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 404, "ID doesn't exist")
		return
	}
	var updateCashier apprequest.Cashier
	err = c.BindJSON(&updateCashier)
	if err != nil {
		api.Error(c, 400, "Error update cashier")
		return
	}
	err = api.cashierUseCase.UpdateCashier(updateCashier, data)
	if err != nil {
		api.Error(c, 404, err.Error())
		return
	}
	api.SuccessNotif(c, "Success")
}

func (api *CashierApi) DeleteCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 404, "ID doesn't exist")
		return
	}
	err = api.cashierUseCase.DeleteCashier(data)
	if err != nil {
		api.Error(c, 404, err.Error())
		return
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
