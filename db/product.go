package db

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Price       float64    `gorm:"not null"`
	Amount      uint       `gorm:"not null"`
	Suppliers   []Supplier `gorm:"many2many:product_suppliers"`
	Rating      float64    `gorm:"default:0"`
	Comment     []Comment  `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   time.Time  `gorm:"autoUpdateTime"`
}
