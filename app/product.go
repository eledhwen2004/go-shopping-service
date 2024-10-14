package app

import (
	"shopping-clone/database"
)

func CreateProduct(product *database.Product) error {
	return database.CreateProduct(product)
}

func ReadProduct(productID string) (*database.Product, error) {
	return database.ReadProduct(productID)
}

func UpdateProduct(product *database.Product) error {
	return database.UpdateProduct(product)
}

func DeleteProduct(productID string) error {
	return database.DeleteCustomer(productID)
}

func GetAllProductsBySuppliedID(supplierID string) (*[]database.Product, error) {
	return database.GetAllProductsBySupplierID(supplierID)
}
