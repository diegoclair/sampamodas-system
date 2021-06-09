package viewmodel

// Business viewmodel
type Business struct {
	UUID        int64  `json:"uuid,omitempty"`
	CompanyUUID int64  `json:"company_uuid,omitempty"`
	Name        string `json:"name,omitempty"`
}
