package company_test

import (
	"job-posting/internal/dto"
	"job-posting/internal/repository/company"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompanyRepositoryImpl_GetCompanies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "Company A").
		AddRow("2", "Company B")

	mock.ExpectQuery("SELECT id, name FROM companies").WillReturnRows(mockRows)

	repo := company.NewCompanyRepository(db)

	companies, _ := repo.GetCompanies()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(companies))
	assert.Equal(t, "Company A", companies[0].Name)
	assert.Equal(t, "1", companies[0].ID)
	assert.Equal(t, "Company B", companies[1].Name)
	assert.Equal(t, "2", companies[1].ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCompanyRepositoryImpl_SaveCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := company.NewCompanyRepository(db)

	company := dto.RequestCompanies{Name: "Test Company"}

	mock.ExpectExec("INSERT INTO companies").
		WithArgs(sqlmock.AnyArg(), company.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SaveCompany(company)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
