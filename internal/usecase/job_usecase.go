package usecase

import (
	"job-posting/internal/dto"
	"job-posting/internal/entity"
)

type JobUsecase interface {
	GetJobs() ([]entity.Job, error)
	SaveJob(job dto.RequestJobs) error
}
