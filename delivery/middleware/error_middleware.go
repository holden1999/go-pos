package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-pos/delivery/commonresp"
)

func ErrorMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		errResp := commonresp.ErrorMessage{}
		json.Unmarshal([]byte(e), &errResp)
		commonresp.NewJsonResponse(c).SendError(errResp)
	}
}
