package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/usecase"
	"net/http"
)

type PaymentApi struct {
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
	api.publicRoute.GET("", api.listPayment)
	api.publicRoute.GET("/:payments", api.detailPayment)
	api.publicRoute.POST("", api.createPayment)
	api.publicRoute.PUT("/:payments", api.updatePayment)
	api.publicRoute.DELETE("", api.updatePayment)
}

func (api *PaymentApi) listPayment(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	skip := c.DefaultQuery("skip", "0")
	c.String(http.StatusOK, "show with %s %s", limit, skip)
}

func (api *PaymentApi) detailPayment(c *gin.Context) {
	id := c.Param("cashierId")
	c.JSON(200, gin.H{
		"categoryId": id,
	})
}

func (api *PaymentApi) createPayment(c *gin.Context) {
	name := c.PostForm("name")
	tipe := c.PostForm("type")
	logo := c.PostForm("logo")
	c.JSON(200, gin.H{
		"name": name,
		"type": tipe,
		"logo": logo,
	})
}

func (api *PaymentApi) updatePayment(c *gin.Context) {
	id := c.Param("cashierId")
	name := c.PostForm("name")
	tipe := c.PostForm("type")
	logo := c.PostForm("logo")
	c.JSON(200, gin.H{
		"id":   id,
		"name": name,
		"type": tipe,
		"logo": logo,
	})
}

func (api *PaymentApi) deletePayment(c *gin.Context) {
	name := c.Param("cashierId")
	c.String(http.StatusOK, "delete %s", name)
}
