package db

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Amount      uint           `json:"amount" gorm:"not null"`
	Suppliers   []Supplier     `gorm:"many2many:product_suppliers"`
	Rating      float64        `json:"rating" gorm:"default:0"`
	Comments    []Comment      `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func addProduct() {
}

func deleteProduct() {
}

func updateProduct() {}

func getProduct() {}
