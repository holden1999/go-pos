package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/model"
	"go-pos/usecase"
	"strconv"
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
	api.publicRoute.GET("/:orderId", api.DetailOrder)
	api.publicRoute.POST("", api.AddOrder)
	api.publicRoute.POST("/subtotal", api.SubTotalOrder)
	api.publicRoute.GET("/:orderId/download", api.DownloadOrder)
	api.publicRoute.GET("/:orderId/check-download", api.CheckDownloadOrder)
}

func (api *OrderApi) ListOrder(c *gin.Context) {
	var meta model.List
	var data model.OrderData
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	data.Order = api.orderUseCase.ListOrder(meta.Limit, meta.Skip)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *OrderApi) DetailOrder(c *gin.Context) {
	id := c.Param("orderId")
	data, _ := strconv.Atoi(id)
	result := api.orderUseCase.DetailOrder(data)
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
	categoryId := c.PostForm("categoryId")
	name := c.PostForm("name")
	image := c.PostForm("image")
	price := c.PostForm("price")
	stock := c.PostForm("stock")
	discount := c.PostForm("discount")
	c.JSON(200, gin.H{
		"categoryId": categoryId,
		"name":       name,
		"image":      image,
		"price":      price,
		"stock":      stock,
		"discount":   discount,
	})
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
