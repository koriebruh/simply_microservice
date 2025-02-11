package dto

import "github.com/koriebruh/simply_microservice/entity"

type OrderRequest struct {
	CustomerId    int64                `json:"customer_id" validate:"required,gt=0"`
	Items         []Product            `json:"items" validate:"required,min=1,dive"` // Minimal 1 item dalam array
	ShippingAddr  string               `json:"shipping_addr" validate:"required,min=10"`
	PaymentMethod entity.PaymentMethod `json:"payment_method" validate:"required,oneof=bank_transfer COD paylater"` // Validasi dengan `oneof`
	Amount        int64                `json:"amount" validate:"required,gt=0"`                                     // Amount harus lebih besar dari 0
}

//SEMENTATA AMMOUNT FORMALITAS NANTI KALO DAH JADI SEMUA KITA  PERBAIKI AMMOUNT INI KITA BUATKAN TABLE USER YG PUNYA SALDO AJA
