package database

import (
	"fmt"
	"time"

	"shopping-clone/postgre"

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

func CreateOrder(order *Order) error {
	result := postgre.DB.Create(&order)
	if result.Error != nil {
		fmt.Printf("Error while creating order %v\n", result.Error)
	}
	return result.Error
}

func ReadOrder(orderID string) (*Order, error) {
	order := Order{}
	result := postgre.DB.Where("id = ?", orderID).First(&order)
	if result.Error != nil {
		fmt.Printf("Error while finding order %v\n", result.Error)
	}
	return &order, result.Error
}

func UpdateOrder(order *Order) error {
	result := postgre.DB.Model(&Order{}).Where("id = ?", order.ID).Updates(&order)
	if result.Error != nil {
		fmt.Printf("Error while updating order %v\n", result.Error)
	}
	return result.Error
}

func DeleteOrder(orderID string) error {
	result := postgre.DB.Where("id = ?", orderID).Delete(&Order{})
	if result.Error != nil {
		fmt.Printf("Error while deleting order %v\n", result.Error)
	}
	return result.Error
}
