package viewmodel

// CreateBusiness viewmodel
type CreateBusiness struct {
	CompanyID int64  `json:"company_id,omitempty"`
	Name      string `json:"name,omitempty"`
}
