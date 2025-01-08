package controller

import (
	"github.com/gin-gonic/gin"
)

type CompaniesController interface {
	GetCompanies(c *gin.Context)
	SaveCompany(c *gin.Context)
}
