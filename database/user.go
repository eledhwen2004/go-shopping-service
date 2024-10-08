package database

import (
	"fmt"
	"time"

	"shopping-clone/postgre"

	"gorm.io/gorm"
)

type user_role int

const (
	AdminRole user_role = iota
	SupplierRole
	CustomerRole
)

type User struct {
	gorm.Model
	ID          string         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	Username    string         `json:"username" gorm:"not null"`
	Password    string         `json:"password" gorm:"not null"`
	Email       string         `json:"email" gorm:"not null"`
	PhoneNumber string         `json:"phone_number" gorm:"not null"`
	Role        user_role      `json:"role" gorm:"not null"`
	Warnings    uint           `json:"warning" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Customer struct {
	gorm.Model
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	UserID    string    `json:"user_id" gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;not null;constraint:OnDelete:CASCADE"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Comments  []Comment `gorm:"constraint:OnDelete:CASCADE"`
	Orders    []Order   `gorm:"constraint:OnDelete:CASCADE"`
}

type Supplier struct {
	gorm.Model
	ID          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey;index"`
	UserID      string    `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;not null;constraint:OnDelete:CASCADE"`
	CompanyName string    `json:"company_name" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	Products    []Product `gorm:"many2many:product_suppliers"`
}

// USER CRUD

func CreateUser(user *User) {
	result := postgre.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("Error while creating user %v\n", result.Error)
	}
}

func ReadUser(id string) User {
	user := User{}
	result := postgre.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		fmt.Printf("Error while finding user %v\n", result.Error)
	}
	return user
}

func UpdateUser(user *User) {
	result := postgre.DB.Model(&User{}).Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
	}
}

func DeleteUser(id string) {
	result := postgre.DB.Where("id = ?", id).Delete(&User{})
	if result.Error != nil {
		fmt.Printf("Error while deleting user %v\n", result.Error)
	}
}

// CUSTOMER CRUD

func CreateCustomer(customer *Customer) {
	tx := postgre.DB.Begin()
	result := tx.Create(&customer)
	if result.Error != nil {
		fmt.Printf("Error while creating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func ReadCustomer(customerID string) Customer {
	customer := Customer{}
	result := postgre.DB.Where("id = ?", customerID).First(&customer)
	if result.Error != nil {
		fmt.Printf("Error while finding customer %v\n", result.Error)
	}
	return customer
}

func UpdateCustomer(customer *Customer) {
	tx := postgre.DB.Begin()
	result := tx.Model(&Customer{}).Where("id = ?", customer.ID).Updates(&customer)
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func DeleteCustomer(customerID string) {
	tx := postgre.DB.Begin()
	result := tx.Where("id = ?", customerID).Delete(&Customer{})
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

// SUPPLIER CRUD

func CreateSupplier(supplier *Supplier) {
	tx := postgre.DB.Begin()
	result := tx.Create(&supplier)
	if result.Error != nil {
		fmt.Printf("Error while creating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func ReadSupplier(supplierID string) Supplier {
	supplier := Supplier{}
	result := postgre.DB.Where("id = ?", supplierID).First(&supplier)
	if result.Error != nil {
		fmt.Printf("Error while finding customer %v\n", result.Error)
	}
	return supplier
}

func UpdateSupplier(supplier *Supplier) {
	tx := postgre.DB.Begin()
	result := tx.Model(&Customer{}).Where("id = ?", supplier.ID).Updates(&supplier)
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func DeleteSupplier(supplierID string) {
	tx := postgre.DB.Begin()
	result := tx.Where("id = ?", supplierID).Delete(&Supplier{})
	if result.Error != nil {
		fmt.Printf("Error while updating user %v\n", result.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}
