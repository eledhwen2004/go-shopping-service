package db

import (
	"time"

	"gorm.io/gorm"
)

// CreditCard represents a user's credit card details in the system.
type CreditCard struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CardholderName  string         `json:"cardholder_name" gorm:"not null"`
	LastFourDigits  string         `json:"last_four_digits" gorm:"not null"`
	ExpirationMonth uint           `json:"expiration_month" gorm:"not null"`
	ExpirationYear  uint           `json:"expiration_year" gorm:"not null"`
	CardType        string         `json:"card_type" gorm:"not null"` // Visa, MasterCard, etc.
	BillingAddress  string         `json:"billing_address" gorm:"not null"`
	CustomerID      uint           `json:"customer_id" gorm:"index;not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func addCreditCard() {
}

func updateCreditCard() {
}

func deleteCreditCard() {
}
