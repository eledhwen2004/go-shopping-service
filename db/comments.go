package db

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID         string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Content    string `json:"content" gorm:"type:varchar(100)"`
	CustomerID string `json:"customer_id" gorm:"not null"`
	ProductID  string `json:"product_id" gorm:"not null"`
}
