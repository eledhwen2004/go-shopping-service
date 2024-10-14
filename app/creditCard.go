package app

import "shopping-clone/database"

func CreateCreditCard(creditCard *database.CreditCard) error {
	return database.CreateCreditCard(creditCard)
}

func ReadCreditCard(creditCardID string) (*database.CreditCard, error) {
	return database.ReadCreditCard(creditCardID)
}

func UpdateCreditCard(creditCard *database.CreditCard) error {
	return database.UpdateCreditCard(creditCard)
}

func DeteleCreditCard(creditCardID string) error {
	return database.DeleteCreditCard(creditCardID)
}
