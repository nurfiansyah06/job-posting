package companies

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/entity"
	redisLibrary "job-posting/internal/redis"
)

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) GetCompanies(page, limit int, search string) ([]entity.Company, int, error) {
	args := m.Called(page, limit, search)
	return args.Get(0).([]entity.Company), args.Get(1).(int), args.Error(2)
}

func (m *MockCompanyRepository) SaveCompany(company dto.RequestCompanies) error {
	args := m.Called(company)
	return args.Error(0)
}

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Del(keys string) *redis.IntCmd {
	args := m.Called(keys)
	return args.Get(0).(*redis.IntCmd)
}

func TestCompaniesUsecaseImpl_GetCompanies(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	MockRedisClient := new(MockRedisClient)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6479",
	})

	redisWrapper := &redisLibrary.RedisClient{
		Client: rdb,
	}

	usecase := NewCompaniesUsecaseImpl(mockRepo, redisWrapper)

	t.Run("Success - Get Companies", func(t *testing.T) {
		mockCompanies := []entity.Company{
			{
				ID:   "1",
				Name: "Test Company 1",
			},
			{
				ID:   "2",
				Name: "Test Company 2",
			},
		}
		mockTotalRows := 2

		mockRepo.On("GetCompanies", 1, 5, "").Return(mockCompanies, mockTotalRows, nil).Once()

		expectedResp := dto.CompaniesResponse{
			Status: constant.StatusSuccess,
			Data:   mockCompanies,
			Pagination: dto.Pagination{
				TotalPages: 1,
				TotalItems: mockTotalRows,
			},
			Message: "Get All Company",
		}

		result, err := usecase.GetCompanies(1, 5, "")
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, result)

		mockRepo.AssertExpectations(t)
		MockRedisClient.AssertExpectations(t)
	})

	t.Run("Success - Get Companies With Search", func(t *testing.T) {
		mockCompanies := []entity.Company{
			{
				ID:   "1",
				Name: "Test Company 1",
			},
			{
				ID:   "2",
				Name: "Test Company 2",
			},
		}
		mockTotalRows := 2

		mockRepo.On("GetCompanies", 1, 5, "company").Return(mockCompanies, mockTotalRows, nil).Once()

		expectedResp := dto.CompaniesResponse{
			Status: constant.StatusSuccess,
			Data:   mockCompanies,
			Pagination: dto.Pagination{
				TotalPages: 1,
				TotalItems: mockTotalRows,
			},
			Message: "Get All Company",
		}

		result, err := usecase.GetCompanies(1, 5, "company")
		assert.NoError(t, err)
		assert.Equal(t, expectedResp, result)

		mockRepo.AssertExpectations(t)
		MockRedisClient.AssertExpectations(t)
	})
}

func TestCompaniesUsecaseImpl_SaveCompany(t *testing.T) {
	mockRepo := new(MockCompanyRepository)
	mockRedisClient := new(MockRedisClient)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	redisWrapper := &redisLibrary.RedisClient{
		Client: rdb,
	}

	usecase := NewCompaniesUsecaseImpl(mockRepo, redisWrapper)

	t.Run("Success - Save Company", func(t *testing.T) {
		companyRequest := dto.RequestCompanies{
			Name: "Test Company",
		}

		mockRepo.On("SaveCompany", companyRequest).Return(nil).Once()
		err := usecase.SaveCompany(companyRequest)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
		mockRedisClient.AssertExpectations(t)
	})

	t.Run("Error - Save Company Failed", func(t *testing.T) {
		companyRequest := dto.RequestCompanies{
			Name: "Test Company",
		}
		expectedErr := errors.New("database error")

		mockRepo.On("SaveCompany", companyRequest).Return(expectedErr).Once()

		err := usecase.SaveCompany(companyRequest)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
