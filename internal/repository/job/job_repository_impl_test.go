package job_test

import (
	"job-posting/internal/dto"
	"job-posting/internal/repository/job"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestJobRepositoryImpl_GetJobs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "title", "description", "company_id"}).
		AddRow("1", "Software Engineer", "Software Engineer", "1").
		AddRow("2", "Software Engineer", "Software Engineer", "1")

	mock.ExpectQuery("SELECT id, title, description, company_id FROM jobs").WillReturnRows(mockRows)

	repo := job.NewJobRepositoryImpl(db)

	jobs, _ := repo.GetJobs()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(jobs))
	assert.Equal(t, "Software Engineer", jobs[0].Title)
	assert.Equal(t, "1", jobs[0].ID)
	assert.Equal(t, "Software Engineer", jobs[1].Title)
	assert.Equal(t, "2", jobs[1].ID)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestJobRepositoryImpl_SaveJob(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := job.NewJobRepositoryImpl(db)

	mock.ExpectExec("INSERT INTO jobs").
		WithArgs(sqlmock.AnyArg(), "Software Engineer", "Software Engineer", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SaveJob(dto.RequestJobs{Title: "Software Engineer", Description: "Software Engineer", CompanyID: "1"})

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
