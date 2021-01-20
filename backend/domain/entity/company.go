package entity

// Company entity model
type Company struct {
	ID             int64  `json:"id,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	CommercialName string `json:"commercial_name,omitempty"`
	LegalName      string `json:"legal_name,omitempty"`
}
