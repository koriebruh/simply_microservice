package entity

import "gorm.io/gorm"

// Order table schema
type Order struct {
	gorm.Model
	Items          []ProductOrder `gorm:"foreignKey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // Linking to ProductOrder
	Amount         int64          `gorm:"not null"`                                                         // Total amount of the order
	PaymentMethod  string         `gorm:"size:100"`                                                         // Restrict size to avoid large inputs
	ShippingAddr   string         `gorm:"size:255"`                                                         // Address for shipping
	ShippingStatus string         `gorm:"size:50"`                                                          // Status of the shipment (Pending, Delivered, etc.)
}

// ProductOrder table schema
type ProductOrder struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`                               // Primary key and auto-increment
	OrderID   uint  `gorm:"not null;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // Foreign key linking to Order table
	ProductID uint  `gorm:"not null"`                                               // Foreign key linking to Product table
	Quantity  int64 `gorm:"not null"`                                               // Quantity of the product ordered
}

// Product table schema
type Product struct {
	gorm.Model
	Name         string `gorm:"size:100;not null"` // Name of the product
	Stock        int64  `gorm:"not null"`          // Current stock level
	PricePerItem int64  `gorm:"not null"`          // Price per item
}

// cascade => - jika OnDelete table parents dihapus tabel terkait dihapus,
// 			  - jika di SetNull maka ketika table parrent di hapus maka children colm akan di ganti null
//			  - jika RESTRICT Melarang penghapusan baris di tabel parent jika ada baris terkait di tabel child
//			  - jika onUpdate maka nilai foreign key di tabel child juga ikut diperbarui otomatis.

// constraint => memastikan tidak boleh null dan harus fk/pk

// fk bebas boleh >= 1 percolum selagi tidak ti gabung dengan pk
