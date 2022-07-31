package apprequest

import (
	"time"
)

type Discount struct {
	Qty       int       `json:"qty"`
	Type      string    `json:"type"`
	Result    int       `json:"result"`
	ExpiredAt time.Time `json:"expiredAt"`
}
