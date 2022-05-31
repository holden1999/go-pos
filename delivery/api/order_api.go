package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/usecase"
	"net/http"
)

type OrderApi struct {
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
	limit := c.DefaultQuery("limit", "10")
	skip := c.DefaultQuery("skip", "0")
	c.String(http.StatusOK, "show with %s %s %s %s", limit, skip)
}

func (api *OrderApi) DetailOrder(c *gin.Context) {
	id := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": id,
	})
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
