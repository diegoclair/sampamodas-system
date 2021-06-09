package entity

// Lead entity
type Lead struct {
	ID          int64
	UUID        string
	Name        string
	Email       string
	PhoneNumber string
	Instagram   string
	LeadAddress []LeadAddress
}

// LeadAddress entity
type LeadAddress struct {
	ID             int64
	LeadID         int64
	LeadUUID       string
	AddressType    string
	Street         string
	Number         string
	Neighborhood   string
	Complement     string
	City           string
	FederativeUnit string
	ZipCode        string
}
