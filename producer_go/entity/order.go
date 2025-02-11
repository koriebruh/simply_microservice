package entity

import (
	"gorm.io/gorm"
)

// Order table schema
type Order struct {
	gorm.Model
	Items          []ProductOrder `gorm:"foreignKey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // Linking to ProductOrder
	Amount         int64          `gorm:"not null"`
	PaymentMethod  PaymentMethod  `gorm:"size:100"`
	PaymentStatus  PaymentStatus  `gorm:"size:100"`
	ShippingAddr   string         `gorm:"size:255"`
	ShippingStatus ShippingStatus `gorm:"size:50"`
}

// ProductOrder table schema
type ProductOrder struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	OrderID   uint  `gorm:"not null;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // Foreign key linking to Order table
	ProductID uint  `gorm:"not null"`                                               // Foreign key linking to Product table
	Quantity  int64 `gorm:"not null"`
}

// Product table schema
type Product struct {
	gorm.Model
	Name         string `gorm:"size:100;not null"`
	Stock        int64  `gorm:"not null"`
	PricePerItem int64  `gorm:"not null"`
}

// cascade => - jika OnDelete table parents dihapus tabel terkait dihapus,
// 			  - jika di SetNull maka ketika table parrent di hapus maka children colm akan di ganti null
//			  - jika RESTRICT Melarang penghapusan baris di tabel parent jika ada baris terkait di tabel child
//			  - jika onUpdate maka nilai foreign key di tabel child juga ikut diperbarui otomatis.

// constraint => memastikan tidak boleh null dan harus fk/pk

// fk bebas boleh >= 1 percolum selagi tidak ti gabung dengan pk

// PaymentStatus Enum for Payment Status
type PaymentStatus string

const (
	Pending   PaymentStatus = "pending"
	Completed PaymentStatus = "completed"
	Failed    PaymentStatus = "failed"
)

// ShippingStatus Enum for Shipping Status
type ShippingStatus string

const (
	PendingShipment ShippingStatus = "pending"
	Shipped         ShippingStatus = "shipped"
	Delivered       ShippingStatus = "delivered"
	Accepted        ShippingStatus = "accepted by user"
)

// PaymentMethod p Enum for Shipping Status
type PaymentMethod string

const (
	BankTransfer PaymentMethod = "bank_transfer"
	COD          PaymentMethod = "COD"
	PayLater     PaymentMethod = "paylater"
)
