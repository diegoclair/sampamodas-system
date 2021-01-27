package viewmodel

//Sale viewmodel
type Sale struct {
	LeadID          int64         `json:"lead_id,omitempty"`
	Freight         float64       `json:"freight,omitempty"`
	PaymentMethodID int64         `json:"payment_method_id,omitempty"`
	SendMethodID    int64         `json:"send_method_id,omitempty"`
	SaleProduct     []SaleProduct `json:"sale_products,omitempty"`
}

//SaleProduct viewmodel
type SaleProduct struct {
	ProductStockID int64 `json:"product_stock_id,omitempty"`
	Quantity       int64 `json:"quantity,omitempty"`
}

// CreateSaleResponse viewmodel
type CreateSaleResponse struct {
	SaleID int64 `json:"sale_id,omitempty"`
}
