package usecase

import (
	"job-posting/internal/dto"
)

type CompanyUsecase interface {
	GetCompanies(page, limit int, search string) (dto.CompaniesResponse, error)
	SaveCompany(company dto.RequestCompanies) error
}
