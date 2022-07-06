package commonresp

import "encoding/json"

type AppHttpResponse interface {
	SendData(message ResponseMessage)
	SendNotif(message ResponseMessageNoData)
	SendError(errMessage ErrorMessage)
}

type ResponseMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseMessageNoData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (e ErrorMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func NewResponseMessage(message string, data interface{}) ResponseMessage {
	return ResponseMessage{
		true, message, data,
	}
}

func NewResponseMessageNoData(message string) ResponseMessageNoData {
	return ResponseMessageNoData{
		Success: true,
		Message: message,
	}
}

func NewErrorMessage(message string) ErrorMessage {
	return ErrorMessage{
		Success: false,
		Message: message,
	}
}
