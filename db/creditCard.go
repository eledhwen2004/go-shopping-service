package db

import (
	"time"
)

// CreditCard represents a user's credit card details in the system.
type CreditCard struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CardholderName  string    `json:"cardholder_name"`
	LastFourDigits  string    `json:"last_four_digits"`
	ExpirationMonth int       `json:"expiration_month"`
	ExpirationYear  int       `json:"expiration_year"`
	CardType        string    `json:"card_type"` // Visa, MasterCard, etc.
	BillingAddress  string    `json:"billing_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UserID          uint      `json:"user_id" gorm:"index"` // Foreign key referencing the user who owns the card
}

// TableName sets a custom table name for the CreditCard model in the database.
func (CreditCard) TableName() string {
	return "credit_cards"
}
