package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderApi struct {
	publicRoute *gin.RouterGroup
}

func NewOrderApi(publicRoute *gin.RouterGroup) {
	orderApi := OrderApi{
		publicRoute: publicRoute,
	}
	orderApi.InitRouter()
}

func (api *OrderApi) InitRouter() {
	api.publicRoute.GET("", api.listOrder)
	api.publicRoute.GET("/:orderId", api.detailOrder)
	api.publicRoute.POST("", api.addOrder)
	api.publicRoute.POST("/subtotal", api.subTotalOrder)
	api.publicRoute.GET("/:orderId/download", api.downloadOrder)
	api.publicRoute.GET("/:orderId/check-download", api.checkDownloadOrder)
}

func (api *OrderApi) listOrder(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	skip := c.DefaultQuery("skip", "0")
	c.String(http.StatusOK, "show with %s %s %s %s", limit, skip)
}

func (api *OrderApi) detailOrder(c *gin.Context) {
	id := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": id,
	})
}

func (api *OrderApi) subTotalOrder(c *gin.Context) {
	productId := c.PostForm("productId")
	qty := c.PostForm("qty")
	c.JSON(200, gin.H{
		"productId": productId,
		"qty":       qty,
	})
}

func (api *OrderApi) addOrder(c *gin.Context) {
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

func (api *OrderApi) downloadOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": orderId,
	})
}

func (api *OrderApi) checkDownloadOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	c.JSON(200, gin.H{
		"orderId": orderId,
	})
}
