package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Items          []Product `gorm:"foreignKey:OrderID"`
	Amount         int64
	PaymentMethod  string
	ShippingAddr   string
	ShippingStatus string
}

type Product struct {
	Id           int64 `gorm:"primaryKey;autoIncrement"`
	Name         string
	Quantity     int64
	PricePerItem int64

	OrderID uint
}
