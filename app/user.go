package app

import "shopping-clone/database"

func CreateUser(user *database.User) error {
	return database.CreateUser(user)
}

func ReadUser(id string) (*database.User, error) {
	return database.ReadUser(id)
}

func UpdateUser(user *database.User) error {
	return database.UpdateUser(user)
}

func DeleteUser(id string) error {
	return database.DeleteUser(id)
}

func CreateCustomer(customer *database.Customer) error {
	return database.CreateCustomer(customer)
}

func ReadCustomer(customerID string) (*database.Customer, error) {
	return database.ReadCustomer(customerID)
}

func UpdateCustomer(customer *database.Customer) error {
	return database.UpdateCustomer(customer)
}

func DeleteCustomer(customerID string) error {
	return database.DeleteCustomer(customerID)
}

func CreateSupplier(supplier *database.Supplier) error {
	return database.CreateSupplier(supplier)
}

func ReadSupplier(supplierID string) (*database.Supplier, error) {
	return database.ReadSupplier(supplierID)
}

func UpdateSupplier(supplier *database.Supplier) error {
	return database.UpdateSupplier(supplier)
}

func DeleteSupplier(supplierID string) error {
	return database.DeleteSupplier(supplierID)
}
