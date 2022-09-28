package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-pos/controller/commonresp"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	if err := c.ShouldBindJSON(body); err != nil {
		return err
	}
	return nil
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

func (b *BaseApi) Error(c *gin.Context, code int, message string) {
	commonresp.NewJsonResponse(c).SendError(code, commonresp.NewErrorMessage(message))
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
