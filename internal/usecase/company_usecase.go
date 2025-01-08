package usecase

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type CompanyUsecase interface {
	GetCompanies() ([]entity.Company, error)
	SaveCompany(company dto.RequestCompanies) error
}
