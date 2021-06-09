package viewmodel

import "github.com/diegoclair/sampamodas-system/backend/util/format"

type Lead struct {
	UUID        string        `json:"uuid,omitempty"`
	Name        string        `json:"name,omitempty"`
	Email       string        `json:"email,omitempty"`
	PhoneNumber string        `json:"phone_number,omitempty"`
	Instagram   string        `json:"instagram,omitempty"`
	LeadAddress []LeadAddress `json:"lead_address,omitempty"`
}

func (l *Lead) Validate() error {
	format.FirstLetterUpperCase(&l.Name)
	return nil
}

type LeadAddress struct {
	LeadAddressID  int64  `json:"lead_address_id,omitempty"`
	LeadUUID       string `json:"lead_uuid,omitempty"`
	AddressType    string `json:"address_type,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Neighborhood   string `json:"neighborhood,omitempty"`
	Complement     string `json:"complement,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
}

func (l *LeadAddress) Validate() error {
	format.FirstLetterUpperCase(&l.Street)
	format.FirstLetterUpperCase(&l.Neighborhood)
	format.FirstLetterUpperCase(&l.City)
	format.ToUpperCase(&l.FederativeUnit)
	return nil
}

// CreateLeadResponse viewmodel
type CreateLeadResponse struct {
	LeadID int64 `json:"lead_id,omitempty"`
}
