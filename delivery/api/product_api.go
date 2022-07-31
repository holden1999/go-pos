package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/authenticator"
	"go-pos/delivery/apprequest"
	"go-pos/delivery/middleware"
	"go-pos/model"
	"go-pos/usecase"
	"strconv"
)

type ProductApi struct {
	BaseApi
	publicRoute    *gin.RouterGroup
	productUseCase usecase.ProductUseCase
}

func (api *ProductApi) InitRouter() {
	api.publicRoute.POST("", api.CreateProduct)
	api.publicRoute.PUT("/:productId", api.UpdateProduct)
	api.publicRoute.DELETE("/:productId", api.DeleteProduct)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("", api.ListProduct)
	api.publicRoute.GET("/:productId", api.DetailProduct)
}

func (api *ProductApi) ListProduct(c *gin.Context) {
	var meta model.Meta
	var data model.ProductData
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	categoryId, _ := strconv.Atoi(c.Query("categoryId"))
	result := api.productUseCase.ListProduct(meta.Limit, meta.Skip, categoryId)
	data.Products = result
	meta.Total = len(result)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *ProductApi) DetailProduct(c *gin.Context) {
	id := c.Param("productId")
	data, _ := strconv.Atoi(id)
	result := api.productUseCase.DetailProduct(data)
	api.Success(c, "Success", result)
}

func (api *ProductApi) CreateProduct(c *gin.Context) {
	var createProduct apprequest.ProductRequest
	c.BindJSON(&createProduct)
	data, err := api.productUseCase.CreateProduct(createProduct)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.Success(c, "Success", data)
}

func (api *ProductApi) UpdateProduct(c *gin.Context) {
	id := c.Param("productId")
	if id == "" {
		c.AbortWithStatusJSON(404, "ID doesn't exist")
	}
	data, _ := strconv.Atoi(id)
	var updateProduct apprequest.ProductRequest
	err := c.BindJSON(&updateProduct)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = api.productUseCase.UpdateProduct(updateProduct, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func (api *ProductApi) DeleteProduct(c *gin.Context) {
	id := c.Param("productId")
	if id == "" {
		c.AbortWithStatusJSON(404, "ID doesn't exist")
	}
	data, _ := strconv.Atoi(id)
	err := api.productUseCase.DeleteProduct(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func NewProductApi(publicRoute *gin.RouterGroup, productUseCase usecase.ProductUseCase) {
	productApi := ProductApi{
		publicRoute:    publicRoute,
		productUseCase: productUseCase,
	}
	productApi.InitRouter()
}
