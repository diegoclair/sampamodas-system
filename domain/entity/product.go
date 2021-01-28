package entity

// Product entity
type Product struct {
	ProductID    int64          `json:"product_id,omitempty"`
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
	BrandID int64  `json:"brand_id,omitempty"`
	Name    string `json:"name,omitempty"`
}

// Gender entity
type Gender struct {
	GenderID int64  `json:"gender_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

// ProductStock entity
type ProductStock struct {
	ProductStockID int64  `json:"product_stock_id,omitempty"`
	Color          Color  `json:"color,omitempty"`
	Size           string `json:"size,omitempty"`
	Quantity       int64  `json:"quantity,omitempty"`
}

// Color entity
type Color struct {
	ColorID int64  `json:"color_id,omitempty"`
	Name    string `json:"name,omitempty"`
}
