package model

type List struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}
