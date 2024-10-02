package models

import (
	"time"

	"gorm.io/gorm"
)

type userRole int

const (
	admin userRole = iota
	supplier
	customer
)

type User struct {
	gorm.Model
	ID          string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Username    string         `gorm:"not null"`
	Password    string         `gorm:"not null"`
	Email       string         `gorm:"not null"`
	PhoneNumber string         `gorm:"not null"`
	Role        userRole       `gorm:"not null"`
	Warnings    uint           `gorm:"default:0"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Customer struct {
	gorm.Model
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	UserID    string    `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Address   string    `gorm:"not null"`
	Comments  []Comment `gorm:"foreignKey:CommentID"`
}

type Supplier struct {
	gorm.Model
	ID          string    `gorm:type:uuid;default:gen_random_uuid();primaryKey;index`
	UserID      string    `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;not null"`
	CompanyName string    `gorm:"not null"`
	Address     string    `gorm:"not null"`
	Products    []Product `gorm:"many2many:product_suppliers"`
}
