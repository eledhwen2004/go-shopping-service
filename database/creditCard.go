package database

import (
	"fmt"
	"time"

	"shopping-clone/postgre"

	"gorm.io/gorm"
)

// CreditCard represents a user's credit card details in the system.
type CreditCard struct {
	ID              string         `json:"id" gorm:"primaryKey"`
	CardholderName  string         `json:"cardholder_name" gorm:"not null"`
	LastFourDigits  string         `json:"last_four_digits" gorm:"not null"`
	ExpirationMonth uint           `json:"expiration_month" gorm:"not null"`
	ExpirationYear  uint           `json:"expiration_year" gorm:"not null"`
	CardType        string         `json:"card_type" gorm:"not null"` // Visa, MasterCard, etc.
	BillingAddress  string         `json:"billing_address" gorm:"not null"`
	CustomerID      string         `json:"customer_id" gorm:"index;not null"`
	Customer        Customer       `gorm:"foreignKey:CustomerID;not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func CreateCreditCard(creditCard *CreditCard) error {
	result := postgre.DB.Create(&creditCard)
	if result.Error != nil {
		fmt.Printf("Error while creating credit card %v\n", result.Error)
	}
	return result.Error
}

func ReadCreditCard(creditCardID string) (*CreditCard, error) {
	creditCard := CreditCard{}
	result := postgre.DB.Where("id = ?", creditCardID).First(&creditCard)
	if result.Error != nil {
		fmt.Printf("Error while finding credit card %v\n", result.Error)
	}
	return &creditCard, result.Error
}

func UpdateCreditCard(creditCard *CreditCard) error {
	result := postgre.DB.Model(&CreditCard{}).Where("id = ?", creditCard.ID).Updates(&creditCard)
	if result.Error != nil {
		fmt.Printf("Error while updating credit card %v\n", result.Error)
	}
	return result.Error
}

func DeleteCreditCard(creditCardID string) error {
	result := postgre.DB.Where("id = ?", creditCardID).Delete(&CreditCard{})
	if result.Error != nil {
		fmt.Printf("Error while deleting credit card %v\n", result.Error)
	}
	return result.Error
}
