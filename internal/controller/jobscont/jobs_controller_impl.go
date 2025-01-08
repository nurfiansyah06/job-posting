package jobscont

import (
	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JobsControllerImpl struct {
	jobsUsecase usecase.JobUsecase
	validate    *validator.Validate
}

func NewJobsControllerImpl(jobsUsecase usecase.JobUsecase, validator *validator.Validate) *JobsControllerImpl {
	return &JobsControllerImpl{
		jobsUsecase: jobsUsecase,
		validate:    validator,
	}
}

func (j *JobsControllerImpl) GetJobs(c *gin.Context) {
	jobs, err := j.jobsUsecase.GetJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseJob{
			Status:  constant.StatusError,
			Message: err.Error(),
		})
		return
	}

	responseJobs := dto.ResponseJob{
		Status:  constant.StatusSuccess,
		Job:     jobs,
		Message: "Success get jobs",
	}

	c.JSON(http.StatusOK, responseJobs)
}

func (j *JobsControllerImpl) SaveJob(c *gin.Context) {
	var jobModel dto.RequestJobs
	err := c.BindJSON(&jobModel)
	if err != nil {
		c.JSON(http.StatusConflict, dto.ResponseJob{
			Status:  constant.StatusError,
			Message: err.Error(),
		})
		return
	}

	// Validate request
	err = j.validate.Struct(jobModel)
	if err != nil {
		c.JSON(http.StatusConflict, dto.ResponseJob{
			Status:  constant.StatusError,
			Message: err.Error(),
		})
		return
	}

	err = j.jobsUsecase.SaveJob(jobModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseJob{
		Status:  constant.StatusSuccess,
		Message: "Success save job",
	})
}
