package api

import (
	"go-pos/authenticator"
	"go-pos/controller/apprequest"
	"go-pos/controller/middleware"
	"go-pos/model"
	"go-pos/usecase"
	"log"

	"github.com/gin-gonic/gin"

	"strconv"
)

type CategoryApi struct {
	BaseApi
	publicRoute     *gin.RouterGroup
	categoryUseCase usecase.CategoryUseCase
}

func (api *CategoryApi) InitRouter() {
	api.publicRoute.POST("", api.CreateCategory)
	api.publicRoute.PUT("/:categoryId", api.UpdateCategory)
	api.publicRoute.DELETE("/:categoryId", api.DeleteCategory)

	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("", api.ListCategory)
	api.publicRoute.GET("/:categoryId", api.DetailCategory)
}

func (api *CategoryApi) ListCategory(c *gin.Context) {
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

func (api *CategoryApi) DetailCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	result := api.categoryUseCase.DetailCategory(data)
	api.Success(c, "Success", result)
}

func (api *CategoryApi) CreateCategory(c *gin.Context) {
	var createCategory apprequest.CategoryRequest
	c.BindJSON(&createCategory)
	log.Println(createCategory.Name)
	data, err := api.categoryUseCase.CreateCategory(createCategory)
	if err != nil {
		api.Error(c, 400, "Error create category")
		return
	}
	api.Success(c, "Success", data)
}

func (api *CategoryApi) UpdateCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	var updateCategory apprequest.CategoryRequest
	c.BindJSON(&updateCategory)
	err = api.categoryUseCase.UpdateCategory(updateCategory, data)
	if err != nil {
		api.Error(c, 404, "Error update category")
		return
	}
	api.SuccessNotif(c, "Success")
}

func (api *CategoryApi) DeleteCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, err := strconv.Atoi(id)
	if err != nil {
		api.Error(c, 400, "ID doesn't exist")
		return
	}
	err = api.categoryUseCase.DeleteCategory(data)
	if err != nil {
		api.Error(c, 404, "Error delete category")
		return
	}
	api.SuccessNotif(c, "Success")
}

func NewCategoryApi(publicRoute *gin.RouterGroup, categoryUseCase usecase.CategoryUseCase) {
	categoryApi := CategoryApi{
		publicRoute:     publicRoute,
		categoryUseCase: categoryUseCase,
	}
	categoryApi.InitRouter()
}
