package usecase

import "go-pos/repository"

type ReportUseCase interface {
}

type reportUseCase struct {
	reportRepo repository.ReportRepo
}

func NewReportUseCase(reportRepo repository.ReportRepo) ReportUseCase {
	return &reportUseCase{reportRepo: reportRepo}
}
