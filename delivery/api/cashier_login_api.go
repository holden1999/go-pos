package api

import "github.com/gin-gonic/gin"

type CashierLoginApi struct {
	BaseApi
	publicRoute *gin.RouterGroup
}
