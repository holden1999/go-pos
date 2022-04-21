package model

import (
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	*gorm.Model
	Qty             int
	Type            string
	Result          int
	ExpiredAt       time.Time
	ExpiredAtFormat string
	StringFormat    string
}
