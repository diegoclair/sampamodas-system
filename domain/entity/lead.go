package entity

// Lead entity
type Lead struct {
	LeadID      int64         `json:"lead_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Email       string        `json:"email,omitempty"`
	PhoneNumber string        `json:"phone_number,omitempty"`
	Instagram   string        `json:"instagram,omitempty"`
	LeadAddress []LeadAddress `json:"lead_address,omitempty"`
}

// LeadAddress entity
type LeadAddress struct {
	LeadAddressID  int64  `json:"lead_address_id,omitempty"`
	LeadID         int64  `json:"lead_id,omitempty"`
	AddressType    string `json:"address_type,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Neighborhood   string `json:"neighborhood,omitempty"`
	Complement     string `json:"complement,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
}
