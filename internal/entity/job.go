package entity

type Job struct {
	ID          string `json:"id"`
	CompanyId   string `json:"company_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"-"`
}
