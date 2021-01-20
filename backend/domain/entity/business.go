package entity

// Business viewmodel
type Business struct {
	ID        int64  `json:"id,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
	Name      string `json:"name,omitempty"`
}
