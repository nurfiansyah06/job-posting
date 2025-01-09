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

func (c *companyRepositoryImpl) GetCompanies(page, limit int, search string) ([]entity.Company, int, error) {
	var (
		companies []entity.Company
		totalRows int
		offset    = (page - 1) * limit
	)

	countQuery := "SELECT COUNT(*) FROM companies WHERE name LIKE ?"
	searchTerm := "%" + search + "%"
	err := c.DB.QueryRow(countQuery, searchTerm).Scan(&totalRows)
	if err != nil {
		return nil, 0, err
	}

	query := "SELECT id, name FROM companies WHERE name LIKE ? LIMIT ? OFFSET ?"
	rows, err := c.DB.Query(query, searchTerm, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var company entity.Company
		err := rows.Scan(&company.ID, &company.Name)
		if err != nil {
			return nil, 0, err
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return companies, totalRows, nil
}

func (c *companyRepositoryImpl) SaveCompany(company dto.RequestCompanies) error {
	uuid := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO companies (id, name) VALUES (?, ?)", uuid, company.Name)
	if err != nil {
		return err
	}

	return nil
}
