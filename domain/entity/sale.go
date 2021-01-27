package entity

//Sale entity
type Sale struct {
	ID              int64         `json:"id,omitempty"`
	LeadID          int64         `json:"lead_id,omitempty"`
	TotalPrice      float64       `json:"total_price,omitempty"`
	Freight         float64       `json:"freight,omitempty"`
	PaymentMethodID int64         `json:"payment_method_id,omitempty"`
	SendMethodID    int64         `json:"send_method_id,omitempty"`
	SaleProduct     []SaleProduct `json:"sale_products,omitempty"`
}

//SaleProduct entity
type SaleProduct struct {
	ID             int64   `json:"id,omitempty"`
	SaleID         int64   `json:"sale_id,omitempty"`
	ProductStockID int64   `json:"product_stock_id,omitempty"`
	Quantity       int64   `json:"quantity,omitempty"`
	Price          float64 `json:"price,omitempty"`
}
