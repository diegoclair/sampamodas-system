package viewmodel

// CreateCompany viewmodel
type CreateCompany struct {
	DocumentNumber string `json:"document_number,omitempty"`
	CommercialName string `json:"commercial_name,omitempty"`
	LegalName      string `json:"legal_name,omitempty"`
}
