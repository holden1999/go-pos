package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/authenticator"
	"go-pos/controller/apprequest"
	"go-pos/controller/middleware"
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
		api.Error(c, 400, "Error Create Product")
	}
	api.Success(c, "Success", data)
}

func (api *ProductApi) UpdateProduct(c *gin.Context) {
	id := c.Param("productId")
	if id == "productId" {
		api.Error(c, 404, "ID doesn't match")
		return
	}
	data, _ := strconv.Atoi(id)
	var updateProduct apprequest.ProductRequest
	c.BindJSON(&updateProduct)
	err := api.productUseCase.UpdateProduct(updateProduct, data)
	if err != nil {
		api.Error(c, 400, "Product data doesn't match")
		return
	}
	api.SuccessNotif(c, "Success")
}

func (api *ProductApi) DeleteProduct(c *gin.Context) {
	id := c.Param("productId")
	if id == "productId" {
		api.Error(c, 404, "ID doesn't exist")
		return
	}
	data, _ := strconv.Atoi(id)
	err := api.productUseCase.DeleteProduct(data)
	if err != nil {
		api.Error(c, 404, "Data not found")
		return
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
