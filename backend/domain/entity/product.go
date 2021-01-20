package entity

// Product entity
type Product struct {
	ID           int64          `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Cost         float64        `json:"cost,omitempty"`
	Price        float64        `json:"price,omitempty"`
	Brand        Brand          `json:"brand,omitempty"`
	Gender       Gender         `json:"gender,omitempty"`
	BusinessID   int64          `json:"business_id,omitempty"`
	ProductStock []ProductStock `json:"product_stock,omitempty"`
}

// Brand entity
type Brand struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Gender entity
type Gender struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ProductStock entity
type ProductStock struct {
	ID       int64  `json:"id,omitempty"`
	Color    Color  `json:"color,omitempty"`
	Size     string `json:"size,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

// Color entity
type Color struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
