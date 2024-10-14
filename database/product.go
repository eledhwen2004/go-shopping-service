package database

import (
	"fmt"
	"time"

	"shopping-clone/postgre"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Amount      uint           `json:"amount" gorm:"not null"`
	Suppliers   []Supplier     `gorm:"many2many:product_suppliers"`
	Rating      float64        `json:"rating" gorm:"default:0"`
	Comments    []Comment      `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func CreateProduct(product *Product) error {
	result := postgre.DB.Create(&product)
	if result.Error != nil {
		fmt.Printf("Error while creating user %v\n", result.Error)
	}
	return result.Error
}

func ReadProduct(productID string) (*Product, error) {
	product := Product{}
	result := postgre.DB.Where("id = ?", productID).First(&product)
	if result.Error != nil {
		fmt.Printf("Error while finding customer %v\n", result.Error)
	}
	return &product, result.Error
}

func UpdateProduct(product *Product) error {
	tx := postgre.DB.Begin()
	result := tx.Model(&Product{}).Where("id = ?", product.ID).Updates(&product)
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func DeleteProduct(productID string) error {
	tx := postgre.DB.Begin()
	result := tx.Where("id = ?", productID).Delete(&Product{})
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return result.Error
}

func GetAllProductsBySupplierID(supplierID string) (*[]Product, error) {
	products := []Product{}
	result := postgre.DB.Where("supplier_id = ?", supplierID).Find(&products)
	if result.Error != nil {
		fmt.Printf("Error while searching products : %v", result.Error)
		return nil, result.Error
	}
	return &products, result.Error
}
