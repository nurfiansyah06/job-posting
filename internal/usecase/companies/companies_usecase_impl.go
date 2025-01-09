package companies

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

type CompaniesUsecaseImpl struct {
	companyRepo repository.CompanyRepository
	redisClient *redis.RedisClient
}

func NewCompaniesUsecaseImpl(companyRepo repository.CompanyRepository, redisClient *redis.RedisClient) *CompaniesUsecaseImpl {
	return &CompaniesUsecaseImpl{
		companyRepo: companyRepo,
		redisClient: redisClient,
	}
}

func (c *CompaniesUsecaseImpl) GetCompanies(page, limit int, search string) (dto.CompaniesResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	companyRedis, err := c.redisClient.Client.Get(context.Background(), constant.KeyRedisCompany).Result()
	if err != nil {
		var companiesResp dto.CompaniesResponse
		if err := json.Unmarshal([]byte(companyRedis), &companiesResp); err == nil {
			return companiesResp, nil
		}
	}

	companies, totalRows, err := c.companyRepo.GetCompanies(page, limit, search)
	if err != nil {
		return dto.CompaniesResponse{}, fmt.Errorf("error when get companies", err.Error())
	}

	companiesResp := dto.CompaniesResponse{
		Status: constant.StatusSuccess,
		Data:   companies,
		Pagination: dto.Pagination{
			TotalPages: page,
			TotalItems: totalRows,
		},
		Message: "Get All Company",
	}

	companiesJSON, err := json.Marshal(companiesResp)
	if err == nil {
		c.redisClient.Client.Set(context.Background(), constant.KeyRedisCompany, companiesJSON, 5*time.Minute)
	}

	return companiesResp, nil
}

func (c *CompaniesUsecaseImpl) SaveCompany(company dto.RequestCompanies) error {
	err := c.companyRepo.SaveCompany(company)
	if err != nil {
		return err
	}

	c.redisClient.Client.Del(context.Background(), constant.KeyRedisCompany)
	return nil
}
