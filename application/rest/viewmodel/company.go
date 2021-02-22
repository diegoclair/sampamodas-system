package viewmodel

// Company viewmodel
type Company struct {
	CompanyID      int64  `json:"company_id,omitempty" mapper:"ID"`
	DocumentNumber string `json:"document_number,omitempty"`
	CommercialName string `json:"commercial_name,omitempty"`
	LegalName      string `json:"legal_name,omitempty"`
}
