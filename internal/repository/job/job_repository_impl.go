package job

import (
	"database/sql"
	"fmt"
	"job-posting/internal/dto"
	"job-posting/internal/entity"

	"github.com/google/uuid"
)

type jobRepositoryImpl struct {
	DB *sql.DB
}

func NewJobRepositoryImpl(db *sql.DB) *jobRepositoryImpl {
	return &jobRepositoryImpl{
		DB: db,
	}
}

func (j *jobRepositoryImpl) GetJobs(page, limit int, search string) ([]entity.Job, int, error) {
	var (
		jobs      []entity.Job
		totalRows int
		offset    = (page - 1) * limit
	)

	countQuery := "SELECT COUNT(*) FROM jobs WHERE title LIKE ? OR description LIKE ?"
	err := j.DB.QueryRow(countQuery, "%"+search+"%", "%"+search+"%").Scan(&totalRows)
	if err != nil {
		return nil, 0, fmt.Errorf("error counting rows: %w", err)
	}

	query := "SELECT id,title, description, company_id FROM jobs WHERE title LIKE ? OR description LIKE ? ORDER BY created_at DESC LIMIT ? OFFSET ? "
	rows, err := j.DB.Query(query, "%"+search+"%", "%"+search+"%", limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching jobs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var job entity.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.CompanyId)
		if err != nil {
			return nil, 0, err
		}
		jobs = append(jobs, job)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return jobs, totalRows, nil
}

func (j *jobRepositoryImpl) SaveJob(job dto.RequestJobs) error {
	uuid := uuid.New().String()

	_, err := j.DB.Exec("INSERT INTO jobs (id, title, description, company_id) VALUES (?, ?, ?, ?)", uuid, job.Title, job.Description, job.CompanyID)
	if err != nil {
		return err
	}

	return nil
}
