package api

import (
	"go-pos/authenticator"
	"go-pos/controller/apprequest"
	"go-pos/controller/middleware"
	"go-pos/model"
	"go-pos/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderApi struct {
	BaseApi
	publicRoute  *gin.RouterGroup
	orderUseCase usecase.OrderUseCase
}

func NewOrderApi(publicRoute *gin.RouterGroup, orderUseCase usecase.OrderUseCase) {
	orderApi := OrderApi{
		publicRoute:  publicRoute,
		orderUseCase: orderUseCase,
	}
	orderApi.InitRouter()
}

func (api *OrderApi) InitRouter() {
	api.publicRoute.GET("", api.ListOrder)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("/:orderId", api.DetailOrder)
	api.publicRoute.POST("", api.AddOrder)
	api.publicRoute.POST("/subtotal", api.SubTotalOrder)
	api.publicRoute.GET("/:orderId/download", api.DownloadOrder)
	api.publicRoute.GET("/:orderId/check-download", api.CheckDownloadOrder)
}

func (api *OrderApi) ListOrder(c *gin.Context) {
	var meta model.Meta
	var data model.ListOrder
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	result, err := api.orderUseCase.ListOrder(meta.Limit, meta.Skip)
	if err != nil {
		api.Error(c, 404, "Success")
	}
	data.List = result
	meta.Total = len(result)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *OrderApi) DetailOrder(c *gin.Context) {
	var order model.DetailOrder
	id := c.Param("orderId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
	}
	result, _ := api.orderUseCase.DetailOrder(data)
	order.Order = result
	api.Success(c, "Success", result)
}

func (api *OrderApi) SubTotalOrder(c *gin.Context) {
	productId := c.PostForm("productId")
	qty := c.PostForm("qty")
	c.JSON(200, gin.H{
		"productId": productId,
		"qty":       qty,
	})
}

func (api *OrderApi) AddOrder(c *gin.Context) {
	var order apprequest.Order
	err := c.BindJSON(&order)
	if err != nil {
		api.Error(c, http.StatusBadRequest, "empty body")
		return
	}
	result, err := api.orderUseCase.CreateOrder(order)
	if err != nil {
		return
	}
	api.Success(c, "Order placed", result)
}

func (api *OrderApi) DownloadOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": orderId,
	})
}

func (api *OrderApi) CheckDownloadOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": orderId,
	})
}
