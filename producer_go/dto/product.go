package dto

type Product struct {
	Id           int64 `json:"id"`
	Quantity     int64 `json:"quantity"`
	PricePerItem int64 `json:"price_per_item"`
}
