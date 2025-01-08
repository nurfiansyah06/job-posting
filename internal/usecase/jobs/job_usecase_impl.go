package jobs

import (
	"fmt"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	"job-posting/internal/repository"
)

type JobsUsecaseImpl struct {
	jobRepo repository.JobRepository
}

func NewJobsUsecaseImpl(jobRepo repository.JobRepository) *JobsUsecaseImpl {
	return &JobsUsecaseImpl{
		jobRepo: jobRepo,
	}
}

func (j *JobsUsecaseImpl) GetJobs() ([]entity.Job, error) {
	jobs, err := j.jobRepo.GetJobs()
	if err != nil {
		return nil, err
	}

	if len(jobs) == 0 {
		return nil, fmt.Errorf("No jobs found")
	}

	return jobs, nil
}

func (j *JobsUsecaseImpl) SaveJob(job dto.RequestJobs) error {
	err := j.jobRepo.SaveJob(job)
	if err != nil {
		return err
	}

	return nil
}
