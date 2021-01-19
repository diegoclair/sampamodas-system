package entity

// Address entity
type Address struct {
	ID             int64  `json:"id,omitempty"`
	LeadID         int64  `json:"sale_id,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}

// Sale entity
type Sale struct {
	ID              int64   `json:"id,omitempty"`
	CompanyID       int64   `json:"company_id,omitempty"`
	LeadID          int64   `json:"lead_id,omitempty"`
	ProductID       int64   `json:"product_id,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Freight         float64 `json:"freight,omitempty"`
	PaymentMethodID int64   `json:"payment_method_id,omitempty"`
	AddressID       int64   `json:"address_id,omitempty"`
}

// SaleSummary entity
type SaleSummary struct {
	CompanyName string  `json:"company_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
}
