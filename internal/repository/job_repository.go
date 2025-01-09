package repository

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type JobRepository interface {
	GetJobs(page, limit int, search string) ([]entity.Job, int, error)
	SaveJob(job dto.RequestJobs) error
}
