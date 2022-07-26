package api

import (
	"github.com/gin-gonic/gin"
	"go-pos/delivery/commonresp"
)

type BaseApi struct {
}

func (b *BaseApi) Success(c *gin.Context, message string, data interface{}) {
	commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage(message, data))
}

func (b *BaseApi) SuccessList(c *gin.Context, message string, data interface{}, meta interface{}) {
	commonresp.NewJsonResponse(c).SendListData(commonresp.NewListResponseMessage(message, data, meta))
}

func (b *BaseApi) SuccessNotif(c *gin.Context, message string) {
	commonresp.NewJsonResponse(c).SendNotif(commonresp.NewResponseMessageNoData(message))
}

func (b *BaseApi) Error(c *gin.Context, message string) {
	commonresp.NewJsonResponse(c).SendError(commonresp.NewErrorMessage(message))
}
