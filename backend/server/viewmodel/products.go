package viewmodel

// CreateProduct viewmodel
type CreateProduct struct {
	ID           int64                `json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Cost         float64              `json:"cost,omitempty"`
	Price        float64              `json:"price,omitempty"`
	BrandName    string               `json:"brand_name,omitempty" mapper:"-"`
	GenderName   string               `json:"gender_name,omitempty" mapper:"-"`
	BusinessID   int64                `json:"business_id,omitempty"`
	ProductStock []CreateProductStock `json:"product_stock,omitempty"`
}

// CreateProductStock viewmodel
type CreateProductStock struct {
	ID       int64  `json:"id,omitempty"`
	Color    string `json:"color,omitempty"`
	Size     string `json:"size,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

// Product viewmodel
type Product struct {
	ID           int64          `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Cost         float64        `json:"cost,omitempty"`
	Price        float64        `json:"price,omitempty"`
	Brand        Brand          `json:"brand,omitempty"`
	Gender       Gender         `json:"gender,omitempty"`
	ProductStock []ProductStock `json:"product_stock,omitempty"`
	BusinessID   int64          `json:"business_id,omitempty"`
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

// ProductStock viewmodel
type ProductStock struct {
	ID       int64  `json:"id,omitempty"`
	Color    Color  `json:"color,omitempty"`
	Size     string `json:"size,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

// Color viewmodel
type Color struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
