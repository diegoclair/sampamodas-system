package viewmodel

// Address model
type Address struct {
	ID             int64  `json:"id,omitempty"`
	Country        string `json:"country,omitempty"`
	Street         string `json:"street,omitempty"`
	Number         string `json:"number,omitempty"`
	Complement     string `json:"complement,omitempty"`
	ZipCode        string `json:"zip_code,omitempty"`
	City           string `json:"city,omitempty"`
	FederativeUnit string `json:"federative_unit,omitempty"`
}

// Sale model
type Sale struct {
	PaymentMethodID int64   `json:"payment_method_id,omitempty"`
	AddressID       int64   `json:"address_id,omitempty"`
	Freight         float64 `json:"freight,omitempty"`
}

// SaleSummary model
type SaleSummary struct {
	CompanyName string  `json:"company_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Freight     float64 `json:"freight,omitempty"`
}
