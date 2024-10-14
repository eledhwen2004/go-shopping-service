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

func CreateProduct(product *Product) {
	result := postgre.DB.Create(&product)
	if result.Error != nil {
		fmt.Printf("Error while creating user %v\n", result.Error)
		return
	}
}

func ReadProduct(productID string) Product {
	product := Product{}
	result := postgre.DB.Where("id = ?", productID).First(&product)
	if result.Error != nil {
		fmt.Printf("Error while finding customer %v\n", result.Error)
	}
	return product
}

func UpdateProduct(product *Product) {
	tx := postgre.DB.Begin()
	result := tx.Model(&Product{}).Where("id = ?", product.ID).Updates(&product)
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func DeleteProduct(productID string) {
	tx := postgre.DB.Begin()
	result := tx.Where("id = ?", productID).Delete(&Product{})
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func GetAllProductsBySupplierID(supplierID string) []Product {
	products := []Product{}
	result := postgre.DB.Where("supplier_id = ?", supplierID).Find(&products)
	if result.Error != nil {
		fmt.Printf("Error while searching products : %v", result.Error)
		return nil
	}
	return products
}
