package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Content    string `gorm:"type:varchar(100)"`
	CustomerID string `gorm:"not null"`
	ProductID  string `gorm:"not null"`
}
