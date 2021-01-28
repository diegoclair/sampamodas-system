package viewmodel

// Business viewmodel
type Business struct {
	BusinessID int64  `json:"business_id,omitempty"`
	CompanyID  int64  `json:"company_id,omitempty"`
	Name       string `json:"name,omitempty"`
}
