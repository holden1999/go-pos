package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/apprequest"
	"go-pos/usecase"

	"strconv"
)

type CategoryApi struct {
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
	api.publicRoute.GET("", api.listCategory)
	api.publicRoute.GET("/:categoryId", api.detailCategory)
	api.publicRoute.POST("", api.createCategory)
	api.publicRoute.PUT("/:categoryId", api.updateCategory)
	api.publicRoute.DELETE("", api.updateCategory)
}

func (api *CategoryApi) listCategory(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	result := api.categoryUseCase.ListCategory(limit, skip)
	c.JSON(200, result)
}

func (api *CategoryApi) detailCategory(c *gin.Context) {
	id := c.Param("cashierId")
	data, _ := strconv.Atoi(id)
	result := api.categoryUseCase.DetailCategory(data)
	c.JSON(200, result)
}

func (api *CategoryApi) createCategory(c *gin.Context) {
	var createCategory apprequest.CategoryRequest
	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
	data, err := api.categoryUseCase.CreateCategory(createCategory)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
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
}

func (api *CategoryApi) deleteCategory(c *gin.Context) {
	id := c.Param("categoryId")
	data, _ := strconv.Atoi(id)
	err := api.categoryUseCase.DeleteCategory(data)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}
}
