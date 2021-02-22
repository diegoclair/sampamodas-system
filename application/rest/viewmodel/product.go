package viewmodel

// Product viewmodel
type Product struct {
	ID           int64          `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Cost         float64        `json:"cost,omitempty"`
	Price        float64        `json:"price,omitempty"`
	BrandName    string         `json:"brand_name,omitempty" mapper:"Brand.Name"`
	GenderName   string         `json:"gender_name,omitempty" mapper:"Gender.Name"`
	BusinessID   int64          `json:"business_id,omitempty"`
	ProductStock []ProductStock `json:"product_stock,omitempty"`
}

// ProductStock viewmodel
type ProductStock struct {
	ID                int64  `json:"id,omitempty"`
	Color             string `json:"color,omitempty" mapper:"Color.Name"`
	Size              string `json:"size,omitempty"`
	AvailableQuantity int64  `json:"available_quantity,omitempty"`
	Quantity          int64  `json:"quantity,omitempty" mapper:"InputQuantity"`
}

// Brand viewmodel
type Brand struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Gender viewmodel
type Gender struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Color viewmodel
type Color struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
