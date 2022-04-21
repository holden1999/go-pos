package api

import "github.com/gin-gonic/gin"

type ReportApi struct {
	publicRoute *gin.RouterGroup
}

func NewReportApi(publicRoute *gin.RouterGroup) {
	reportApi := ReportApi{
		publicRoute: publicRoute,
	}
	reportApi.InitRouter()
}

func (api *ReportApi) InitRouter() {
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
