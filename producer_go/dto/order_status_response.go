package dto

type OrderStatusResponse struct {
	Items          []Product `json:"items"`
	Amount         int64     `json:"amount"`
	PaymentMethod  string    `json:"payment_method"`
	ShippingAddr   string    `json:"shipping_addr"`
	ShippingStatus string    `json:"shipping_status"`
}
