package models

import "gorm.io/gorm"

type order_status uint

const (
	Pending order_status = iota
	Shipper
	Delivered
	Canceled
)

type Order struct {
	gorm.Model
	ID          string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	ProductID   string       `gorm:"not null"`
	Product     Product      `gorm:"foreignKey:ProductID;not null"`
	SupplierID  string       `gorm:"not null"`
	Supplier    Supplier     `gorm:"foreignKey:SupplierID;not null"`
	CustomerID  string       `gorm:"not null"`
	Customer    Customer     `gorm:"foreignKey:CustomerID;not null"`
	TotalPrice  float64      `gorm:"not null"`
	TotalAmount uint         `gorm:"not null"`
	Status      order_status `gorm:"not null"`
}
