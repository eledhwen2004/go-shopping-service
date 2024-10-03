package db

import (
	"time"

	"gorm.io/gorm"
)

type order_status uint

const (
	Pending order_status = iota
	Shipper
	Delivered
	Canceled
)

type Order struct {
	gorm.Model
	ID          string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	ProductID   string         `json:"product_id" gorm:"not null"`
	Product     Product        `gorm:"foreignKey:ProductID;not null"`
	SupplierID  string         `json:"supplier_id" gorm:"not null"`
	Supplier    Supplier       `gorm:"foreignKey:SupplierID;not null"`
	CustomerID  string         `json:"custome_id"`
	Customer    Customer       `gorm:"foreignKey:CustomerID;not null"`
	TotalPrice  float64        `json:"total_price"`
	TotalAmount uint           `json:"total_amount" gorm:"not null"`
	Status      order_status   `json:"status" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func createOrder() {}

func deleteOrder() {}

func updateOrderStatus() {}
