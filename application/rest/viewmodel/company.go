package viewmodel

// Company viewmodel
type Company struct {
	UUID           string `json:"uuid,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	CommercialName string `json:"commercial_name,omitempty"`
	LegalName      string `json:"legal_name,omitempty"`
}
