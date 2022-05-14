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
