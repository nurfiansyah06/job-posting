package dto

import "job-posting/internal/entity"

type ResponseCompanies struct {
	Status    string           `json:"status"`
	Companies []entity.Company `json:"companies"`
	Message   string           `json:"message"`
}

type ResponseJob struct {
	Status  string       `json:"status"`
	Job     []entity.Job `json:"jobs"`
	Message string       `json:"message"`
}

type RequestCompanies struct {
	Name string `json:"name" validate:"required"`
}

type RequestJobs struct {
	CompanyID   string `json:"company_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
