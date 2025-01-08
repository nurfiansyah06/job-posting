package repository

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type JobRepository interface {
	GetJobs() ([]entity.Job, error)
	SaveJob(job dto.RequestJobs) error
}
