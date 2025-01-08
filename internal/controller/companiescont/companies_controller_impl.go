package companiescont

import (
	"job-posting/internal/constant"
	"job-posting/internal/dto"
	"job-posting/internal/usecase"
	"net/http"

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
	companies, err := company.companiesUsecase.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseCompanies{
			Status:  constant.StatusError,
			Message: err.Error(),
		})
		return
	}

	responseCompanies := dto.ResponseCompanies{
		Status:    constant.StatusSuccess,
		Companies: companies,
		Message:   "Success get companies",
	}

	c.JSON(http.StatusOK, responseCompanies)
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
