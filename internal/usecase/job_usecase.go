package usecase

import (
	"job-posting/internal/dto"
)

type JobUsecase interface {
	GetJobs(page, limit int, search string) (dto.JobsResponse, error)
	SaveJob(job dto.RequestJobs) error
}
