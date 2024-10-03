package db

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID         string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Content    string         `json:"content" gorm:"type:varchar(100)"`
	CustomerID string         `json:"customer_id" gorm:"not null"`
	ProductID  string         `json:"product_id" gorm:"not null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func sendComment() {}

func updateComment() {}

func deleteComment() {
}
