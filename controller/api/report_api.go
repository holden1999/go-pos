package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/authenticator"
	"go-pos/controller/middleware"
	"go-pos/usecase"
)

type ReportApi struct {
	publicRoute   *gin.RouterGroup
	reportUseCase usecase.ReportUseCase
}

func NewReportApi(publicRoute *gin.RouterGroup, reportUseCase usecase.ReportUseCase) {
	reportApi := ReportApi{
		publicRoute:   publicRoute,
		reportUseCase: reportUseCase,
	}
	reportApi.InitRouter()
}

func (api *ReportApi) InitRouter() {
	tokenService := authenticator.NewTokenConfig()
	api.publicRoute.Use(middleware.NewTokenValidator(&tokenService).RequireToken())
	api.publicRoute.GET("revenues", api.revenues)
	api.publicRoute.GET("solds", api.solds)
}

func (api *ReportApi) revenues(c *gin.Context) {
	c.JSON(200, gin.H{
		"revenues": "revenues",
	})
}

func (api *ReportApi) solds(c *gin.Context) {
	c.JSON(200, gin.H{
		"solds": "solds",
	})
}
