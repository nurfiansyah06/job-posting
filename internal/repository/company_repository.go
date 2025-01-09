package repository

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type CompanyRepository interface {
	GetCompanies(page, limit int, search string) ([]entity.Company, int, error)
	SaveCompany(company dto.RequestCompanies) error
}
