package main

import (
	"job-posting/internal/controller/companiescont"
	"job-posting/internal/controller/jobscont"
	"job-posting/internal/db"
	"job-posting/internal/redis"
	"job-posting/internal/repository/company"
	"job-posting/internal/repository/job"
	"job-posting/internal/usecase/companies"
	"job-posting/internal/usecase/jobs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisClient := redis.NewRedisClient()
	validate := validator.New()

	companyRepository := company.NewCompanyRepository(db)
	companyUsecase := companies.NewCompaniesUsecaseImpl(companyRepository, redisClient)
	companyController := companiescont.NewCompaniesControllerImpl(companyUsecase, validate)

	jobRepository := job.NewJobRepositoryImpl(db)
	jobUsecase := jobs.NewJobsUsecaseImpl(jobRepository, redisClient)
	jobController := jobscont.NewJobsControllerImpl(jobUsecase, validate)

	router := gin.Default()

	router.GET("/api/v1/companies", companyController.GetCompanies)
	router.POST("/api/v1/companies", companyController.SaveCompany)

	router.GET("/api/v1/jobs", jobController.GetJobs)
	router.POST("/api/v1/jobs", jobController.SaveJob)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if err := router.Run(":8080"); err != nil {
		log.Print("Server failed to start", err.Error())
	}
}
