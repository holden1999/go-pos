package repository

import "gorm.io/gorm"

type ReportRepo interface {
}

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo(db *gorm.DB) ReportRepo {
	return &reportRepo{db: db}
}
