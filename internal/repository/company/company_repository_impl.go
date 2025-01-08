package company

import (
	"database/sql"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	"job-posting/internal/repository"

	"github.com/google/uuid"
)

type companyRepositoryImpl struct {
	DB *sql.DB
}

func NewCompanyRepository(DB *sql.DB) repository.CompanyRepository {
	return &companyRepositoryImpl{
		DB: DB,
	}
}

func (c *companyRepositoryImpl) GetCompanies() ([]entity.Company, error) {
	var companies []entity.Company

	rows, err := c.DB.Query("SELECT id, name FROM companies")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var company entity.Company
		err := rows.Scan(&company.ID, &company.Name)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (c *companyRepositoryImpl) SaveCompany(company dto.RequestCompanies) error {
	uuid := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO companies (id, name) VALUES (?, ?)", uuid, company.Name)
	if err != nil {
		return err
	}

	return nil
}
