package controller

import "github.com/gin-gonic/gin"

type JobsController interface {
	GetJobs(c *gin.Context)
	SaveJob(c *gin.Context)
}
