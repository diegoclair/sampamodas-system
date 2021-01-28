package entity

// Product entity
type Product struct {
	ID           int64
	Name         string
	Cost         float64
	Price        float64
	Brand        Brand
	Gender       Gender
	BusinessID   int64
	ProductStock []ProductStock
}

// Brand entity
type Brand struct {
	ID   int64
	Name string
}

// Gender entity
type Gender struct {
	ID   int64
	Name string
}

// ProductStock entity
type ProductStock struct {
	ID                int64
	Color             Color
	Size              string
	AvailableQuantity int64
	InputQuantity     int64
}

// Color entity
type Color struct {
	ID   int64
	Name string
}
