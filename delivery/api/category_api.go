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

type CategoryApi struct {
	BaseApi
	publicRoute     *gin.RouterGroup
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryApi(publicRoute *gin.RouterGroup, categoryUseCase usecase.CategoryUseCase) {
	categoryApi := CategoryApi{
		publicRoute:     publicRoute,
		categoryUseCase: categoryUseCase,
	}
	categoryApi.InitRouter()
}

func (api *CategoryApi) InitRouter() {
	api.publicRoute.POST("", api.createCategory)
	api.publicRoute.PUT("/:categoryId", api.updateCategory)
	api.publicRoute.DELETE("", api.updateCategory)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("", api.listCategory)
	api.publicRoute.GET("/:categoryId", api.detailCategory)
}

func (api *CategoryApi) listCategory(c *gin.Context) {
	var meta model.Meta
	var data model.CategoryData
	meta.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	meta.Skip, _ = strconv.Atoi(c.DefaultQuery("skip", "0"))
	result := api.categoryUseCase.ListCategory(meta.Limit, meta.Skip)
	data.Category = result
	meta.Total = len(result)
	data.Meta = meta
	api.Success(c, "Success", data)
}

func (api *CategoryApi) detailCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, _ := strconv.Atoi(id)
	result := api.categoryUseCase.DetailCategory(data)
	api.Success(c, "Success", result)
}

func (api *CategoryApi) createCategory(c *gin.Context) {
	var createCategory apprequest.CategoryRequest
	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	data, err := api.categoryUseCase.CreateCategory(createCategory)
	if err != nil {
		c.AbortWithStatusJSON(401, err.Error())
	}
	c.JSON(200, data)
}

func (api *CategoryApi) updateCategory(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	var updateCategory apprequest.CategoryRequest
	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	err = api.categoryUseCase.UpdateCategory(updateCategory, data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}

func (api *CategoryApi) deleteCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, _ := strconv.Atoi(id)
	err := api.categoryUseCase.DeleteCategory(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	api.SuccessNotif(c, "Success")
}
