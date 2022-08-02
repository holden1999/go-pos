package commonresp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	c *gin.Context
}

func (j *JsonResponse) SendListData(message ListResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendNotif(message ResponseMessageNoData) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendData(message ResponseMessage) {
	j.c.JSON(http.StatusOK, message)
}

func (j *JsonResponse) SendError(code int, errMessage ErrorMessage) {
	j.c.JSON(code, errMessage)
}

func NewJsonResponse(c *gin.Context) AppHttpResponse {
	return &JsonResponse{c: c}
}
