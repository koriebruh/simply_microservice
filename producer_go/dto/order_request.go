package dto

type OrderRequest struct {
	CustomerId    int64     `json:"customer_id"`
	Items         []Product `json:"items"`
	ShippingAddr  string    `json:"shipping_addr"`
	PaymentMethod string    `json:"payment_method"`
	Amount        int64     `json:"amount"`
}
