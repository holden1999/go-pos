package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
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
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	result := api.cashierUseCase.ListCashier(limit, skip)
	api.Success(c, "Success", result)
}

func (api *CashierApi) CreateCashier(c *gin.Context) {
	var createCashier apprequest.Cashier
	err := c.ShouldBindJSON(&createCashier)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	data, err := api.cashierUseCase.CreateCashier(createCashier)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.Success(c, "Success", data)
}

func (api *CashierApi) DetailCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	result := api.cashierUseCase.DetailCashier(data)
	c.JSON(200, gin.H{
		"cashierId": result.ID,
		"name":      result.Name,
	})
}

func (api *CashierApi) UpdateCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	var updateCashier apprequest.Cashier
	err := c.ShouldBindJSON(&updateCashier)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = api.cashierUseCase.UpdateCashier(updateCashier, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func (api *CashierApi) DeleteCashier(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	err := api.cashierUseCase.DeleteCashier(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func NewCashierApi(publicRoute *gin.RouterGroup, cashierUseCase usecase.CashierUseCase) (*CashierApi, error) {
	cashierApi := CashierApi{
		publicRoute:    publicRoute,
		cashierUseCase: cashierUseCase,
	}
	cashierApi.InitRouter()
	return &cashierApi, nil
}
