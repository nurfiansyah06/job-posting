package job

import (
	"database/sql"
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

func (j *jobRepositoryImpl) GetJobs() ([]entity.Job, error) {
	var jobs []entity.Job

	rows, err := j.DB.Query("SELECT id, title, description, company_id FROM jobs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var job entity.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.CompanyId)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (j *jobRepositoryImpl) SaveJob(job dto.RequestJobs) error {
	uuid := uuid.New().String()

	_, err := j.DB.Exec("INSERT INTO jobs (id, title, description, company_id) VALUES (?, ?, ?, ?)", uuid, job.Title, job.Description, job.CompanyID)
	if err != nil {
		return err
	}

	return nil
}
