package app

import "shopping-clone/database"

func CreateOrder(order *database.Order) error {
	return database.CreateOrder(order)
}

func ReadOrder(orderID string) (*database.Order, error) {
	return database.ReadOrder(orderID)
}

func UpdateOrder(order *database.Order) error {
	return database.UpdateOrder(order)
}

func DeleteOrder(orderID string) error {
	return database.DeleteOrder(orderID)
}
