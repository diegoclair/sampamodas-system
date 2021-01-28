package entity

//Sale entity
type Sale struct {
	ID              int64
	LeadID          int64
	TotalPrice      float64
	Freight         float64
	PaymentMethodID int64
	SendMethodID    int64
	SaleProduct     []SaleProduct
}

//SaleProduct entity
type SaleProduct struct {
	ID             int64
	SaleID         int64
	ProductStockID int64
	Quantity       int64
	Price          float64
}
