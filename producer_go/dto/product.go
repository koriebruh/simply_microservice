package dto

type Product struct {
	ProductId int64 `json:"product_id" validate:"required,gt=0"`
	Quantity  int64 `json:"quantity" validate:"required,gt=0"`
}
