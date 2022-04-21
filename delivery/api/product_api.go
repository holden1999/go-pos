package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
	"go-pos/usecase"
	"strconv"
)

type ProductApi struct {
	publicRoute    *gin.RouterGroup
	productUseCase usecase.ProductUseCase
}

func NewProductApi(publicRoute *gin.RouterGroup, productUseCase usecase.ProductUseCase) {
	productApi := ProductApi{
		publicRoute:    publicRoute,
		productUseCase: productUseCase,
	}
	productApi.InitRouter()
}

func (api *ProductApi) InitRouter() {
	api.publicRoute.GET("", api.listProduct)
	api.publicRoute.GET("/:productId", api.detailProduct)
	api.publicRoute.POST("", api.createProduct)
	api.publicRoute.PUT("/:productId", api.updateProduct)
	api.publicRoute.DELETE("/:productId", api.deleteProduct)
}

func (api *ProductApi) listProduct(c *gin.Context) {

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	result := api.productUseCase.ListProduct(limit, skip, categoryId)
	c.JSON(200, result)
}

func (api *ProductApi) detailProduct(c *gin.Context) {
	id := c.Param("productId")
	data, _ := strconv.Atoi(id)
	result := api.productUseCase.DetailProduct(data)
	c.JSON(200, result)
}

func (api *ProductApi) createProduct(c *gin.Context) {
	var createProduct apprequest.ProductRequest
	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	data, err := api.productUseCase.CreateProduct(createProduct)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	c.JSON(200, data)
}

func (api *ProductApi) updateProduct(c *gin.Context) {
	id := c.Param("productId")
	data, _ := strconv.Atoi(id)
	var updateProduct apprequest.ProductRequest
	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = api.productUseCase.UpdateProduct(updateProduct, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
}

func (api *ProductApi) deleteProduct(c *gin.Context) {
	id := c.Param("productId")
	data, _ := strconv.Atoi(id)
	err := api.productUseCase.DeleteProduct(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
}
