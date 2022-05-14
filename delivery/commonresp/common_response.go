package commonresp

type AppHttpResponse interface {
	SendData(message ResponseMessage)
}

type ResponseMessage struct {
	Success string      `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseMessage(message string, data interface{}) ResponseMessage {
	return ResponseMessage{
		"true", message, data,
	}
}
