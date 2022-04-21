package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/usecase"
	"net/http"
)

type CashierApi struct {
	publicRoute    *gin.RouterGroup
	cashierUseCase usecase.CashierUseCase
}

func NewCashierApi(publicRoute *gin.RouterGroup, cashierUseCase usecase.CashierUseCase) {
	cashierApi := CashierApi{
		publicRoute:    publicRoute,
		cashierUseCase: cashierUseCase,
	}
	cashierApi.InitRouter()
}

func (api CashierApi) InitRouter() {
	api.publicRoute.GET("", api.listCashier)
	api.publicRoute.GET("/:cashierId", api.detailCashier)
	api.publicRoute.POST("", api.createCashier)
	api.publicRoute.PUT("/:cashierId", api.updateCashier)
	api.publicRoute.DELETE("", api.deleteCashier)
}

func (api *CashierApi) listCashier(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	skip := c.DefaultQuery("skip", "0")
	c.String(http.StatusOK, "show with %s %s", limit, skip)
}

func (api *CashierApi) detailCashier(c *gin.Context) {
	id := c.Param("cashierId")
	c.JSON(200, gin.H{
		"cashierId": id,
	})
}

func (api *CashierApi) createCashier(c *gin.Context) {
	name := c.PostForm("name")
	passcode := c.PostForm("passcode")
	c.JSON(200, gin.H{
		"name":     name,
		"passcode": passcode,
	})
}

func (api *CashierApi) updateCashier(c *gin.Context) {
	id := c.Param("cashierId")
	name := c.PostForm("name")
	passcode := c.PostForm("passcode")
	c.JSON(200, gin.H{
		"id":       id,
		"name":     name,
		"passcode": passcode,
	})
}

func (api *CashierApi) deleteCashier(c *gin.Context) {
	name := c.Param("cashierId")
	c.String(http.StatusOK, "delete %s", name)
}
