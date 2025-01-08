package repository

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type CompanyRepository interface {
	GetCompanies() ([]entity.Company, error)
	SaveCompany(company dto.RequestCompanies) error
}
