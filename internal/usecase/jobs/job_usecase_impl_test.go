package jobs

import (
	"errors"
	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	"testing"

	redisLibrary "job-posting/internal/redis"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) GetJobs(page, limit int, search string) ([]entity.Job, int, error) {
	args := m.Called(page, limit, search)
	return args.Get(0).([]entity.Job), args.Get(1).(int), args.Error(2)
}

func (m *MockJobRepository) SaveJob(job dto.RequestJobs) error {
	args := m.Called(job)
	return args.Error(0)
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Del(keys string) *redis.IntCmd {
	args := m.Called(keys)
	return args.Get(0).(*redis.IntCmd)
}

func TestJobsUsecaseImpl_GetJobs(t *testing.T) {
	mockRepo := new(MockJobRepository)
	MockRedisClient := new(MockRedisClient)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6479",
	})

	redisWrapper := &redisLibrary.RedisClient{
		Client: rdb,
	}

	usecase := NewJobsUsecaseImpl(mockRepo, redisWrapper)

	t.Run("Success - Get Jobs", func(t *testing.T) {
		mockJobs := []entity.Job{
			{
				ID:          "1",
				CompanyId:   "2",
				Title:       "Software Engineer",
				Description: "Jr. Software Engineer",
			},
			{
				ID:          "1",
				CompanyId:   "2",
				Title:       "Software Engineer",
				Description: "Jr. Software Engineer",
			},
		}

		mockTotalRows := 2

		mockRepo.On("GetJobs", 1, 5, "").Return(mockJobs, mockTotalRows, nil).Once()

		expectedResp := dto.JobsResponse{
			Status: constant.StatusSuccess,
			Data:   mockJobs,
			Pagination: dto.Pagination{
				TotalPages: 1,
				TotalItems: mockTotalRows,
			},
			Message: "Get All Job",
		}

		result, err := usecase.GetJobs(1, 5, "")
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, result)
		mockRepo.AssertExpectations(t)
		MockRedisClient.AssertExpectations(t)
	})
}

func TestJobsUsecaseImpl_SaveJob(t *testing.T) {
	mockRepo := new(MockJobRepository)
	mockRedisClient := new(MockRedisClient)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	redisWrapper := &redisLibrary.RedisClient{
		Client: rdb,
	}

	usecase := NewJobsUsecaseImpl(mockRepo, redisWrapper)

	t.Run("Success - Save Job", func(t *testing.T) {
		jobRequest := dto.RequestJobs{
			CompanyID:   "1",
			Title:       "Test Job A",
			Description: "Test Job A",
		}

		mockRepo.On("SaveJob", jobRequest).Return(nil).Once()
		err := usecase.SaveJob(jobRequest)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockRedisClient.AssertExpectations(t)
	})

	t.Run("Error - Save Job Failed", func(t *testing.T) {
		jobRequest := dto.RequestJobs{
			CompanyID:   "1",
			Title:       "Test Job A",
			Description: "Test Job A",
		}

		expectedErr := errors.New("database error")

		mockRepo.On("SaveJob", jobRequest).Return(expectedErr).Once()

		err := usecase.SaveJob(jobRequest)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
