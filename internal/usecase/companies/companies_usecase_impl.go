package companies

import (
	"fmt"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	"job-posting/internal/repository"
)

type CompaniesUsecaseImpl struct {
	companyRepo repository.CompanyRepository
}

func NewCompaniesUsecaseImpl(companyRepo repository.CompanyRepository) *CompaniesUsecaseImpl {
	return &CompaniesUsecaseImpl{
		companyRepo: companyRepo,
	}
}

func (c *CompaniesUsecaseImpl) GetCompanies() ([]entity.Company, error) {
	companies, err := c.companyRepo.GetCompanies()
	if err != nil {
		return nil, err
	}

	if len(companies) == 0 {
		return nil, fmt.Errorf("No companies found")
	}

	return companies, nil
}

func (c *CompaniesUsecaseImpl) SaveCompany(company dto.RequestCompanies) error {
	err := c.companyRepo.SaveCompany(company)
	if err != nil {
		return err
	}

	return nil
}
