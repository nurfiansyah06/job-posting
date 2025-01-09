package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	"job-posting/internal/redis"
	"job-posting/internal/repository"
	"time"
)

type JobsUsecaseImpl struct {
	jobRepo     repository.JobRepository
	redisClient *redis.RedisClient
}

func NewJobsUsecaseImpl(jobRepo repository.JobRepository, redisClient *redis.RedisClient) *JobsUsecaseImpl {
	return &JobsUsecaseImpl{
		jobRepo:     jobRepo,
		redisClient: redisClient,
	}
}

func (j *JobsUsecaseImpl) GetJobs() ([]entity.Job, error) {
	cachedJobs, err := j.redisClient.Client.Get(context.Background(), constant.KeyRedis).Result()
	if err == nil {
		var jobs []entity.Job
		if err := json.Unmarshal([]byte(cachedJobs), &jobs); err == nil {
			return jobs, nil
		}
	}

	jobs, err := j.jobRepo.GetJobs()
	if err != nil {
		return nil, err
	}

	if len(jobs) == 0 {
		return nil, fmt.Errorf("No jobs found")
	}

	jobsJSON, err := json.Marshal(jobs)
	if err == nil {
		j.redisClient.Client.Set(context.Background(), constant.KeyRedis, jobsJSON, 5*time.Minute) // Cache for 5 minutes
	}

	return jobs, nil
}

func (j *JobsUsecaseImpl) SaveJob(job dto.RequestJobs) error {
	err := j.jobRepo.SaveJob(job)
	if err != nil {
		return err
	}

	j.redisClient.Client.Del(context.Background(), constant.KeyRedis)
	return nil
}
