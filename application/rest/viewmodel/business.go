package viewmodel

// Business viewmodel
type Business struct {
	UUID        string `json:"uuid,omitempty"`
	CompanyUUID string `json:"company_uuid,omitempty"`
	Name        string `json:"name,omitempty"`
}
