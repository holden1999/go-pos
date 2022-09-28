package commonresp

import (
	"encoding/json"
	"go-pos/model"
)

type AppHttpResponse interface {
	SendData(message ResponseMessage)
	SendListData(message ListResponseMessage)
	SendNotif(message ResponseMessageNoData)
	SendError(code int, errMessage ErrorMessage)
}

type ResponseMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponseMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    model.Meta  `json:"meta"`
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

func NewListResponseMessage(message string, data interface{}, meta model.Meta) ListResponseMessage {
	return ListResponseMessage{
		true, message, data, meta,
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
