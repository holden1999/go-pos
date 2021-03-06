package commonresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	c *gin.Context
}

func (j *JsonResponse) SendData(message ResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func NewJsonResponse(c *gin.Context) AppHttpResponse {
	return &JsonResponse{c: c}
}
