package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"job-posting/internal/constant"
	"job-posting/internal/dto"
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

func (j *JobsUsecaseImpl) GetJobs(page, limit int, search string) (dto.JobsResponse, error) {
	cachedJobs, err := j.redisClient.Client.Get(context.Background(), constant.KeyRedis).Result()
	if err == nil {
		var jobs dto.JobsResponse
		if err := json.Unmarshal([]byte(cachedJobs), &jobs); err == nil {
			return jobs, nil
		}
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	jobs, totalRows, err := j.jobRepo.GetJobs(page, limit, search)
	if err != nil {
		return dto.JobsResponse{}, err
	}

	if len(jobs) == 0 {
		return dto.JobsResponse{}, fmt.Errorf("No jobs found")
	}

	jobsJSON, err := json.Marshal(jobs)
	if err == nil {
		j.redisClient.Client.Set(context.Background(), constant.KeyRedis, jobsJSON, 5*time.Minute)
	}

	jobsResp := dto.JobsResponse{
		Status: constant.StatusSuccess,
		Data:   jobs,
		Pagination: dto.Pagination{
			TotalPages: page,
			TotalItems: totalRows,
		},
		Message: "Get All Job",
	}

	return jobsResp, nil
}

func (j *JobsUsecaseImpl) SaveJob(job dto.RequestJobs) error {
	err := j.jobRepo.SaveJob(job)
	if err != nil {
		return err
	}

	j.redisClient.Client.Del(context.Background(), constant.KeyRedis)
	return nil
}
