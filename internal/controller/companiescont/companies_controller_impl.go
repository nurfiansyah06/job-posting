package companiescont

import (
	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CompaniesControllerImpl struct {
	companiesUsecase usecase.CompanyUsecase
	validate         *validator.Validate
}

func NewCompaniesControllerImpl(companiesUsecase usecase.CompanyUsecase, validator *validator.Validate) *CompaniesControllerImpl {
	return &CompaniesControllerImpl{
		companiesUsecase: companiesUsecase,
		validate:         validator,
	}
}

func (company *CompaniesControllerImpl) GetCompanies(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 2
	}

	search := c.Request.URL.Query().Get("search")

	companies, err := company.companiesUsecase.GetCompanies(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseCompanies{
			Status:  constant.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (company *CompaniesControllerImpl) SaveCompany(c *gin.Context) {
	var companyModel dto.RequestCompanies
	err := c.BindJSON(&companyModel)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate request
	err = company.validate.Struct(companyModel)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = company.companiesUsecase.SaveCompany(companyModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseCompanies{
		Status:  constant.StatusSuccess,
		Message: "Company saved",
	})
}
