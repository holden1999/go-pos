package middleware

import (
	"go-pos/controller/commonresp"
	"go-pos/model"

	"github.com/gin-gonic/gin"
)

type CustomResp struct {
}

func (b *CustomResp) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
}

func (b *CustomResp) Success(c *gin.Context, message string, data interface{}) {
	commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage(message, data))
}

func (b *CustomResp) SuccessList(c *gin.Context, message string, data interface{}, meta model.Meta) {
	commonresp.NewJsonResponse(c).SendListData(commonresp.NewListResponseMessage(message, data, meta))
}

func (b *CustomResp) SuccessNotif(c *gin.Context, message string) {
	commonresp.NewJsonResponse(c).SendNotif(commonresp.NewResponseMessageNoData(message))
}

func (b *CustomResp) Error(c *gin.Context, code int, message string) {
	commonresp.NewJsonResponse(c).SendError(code, commonresp.NewErrorMessage(message))
}
