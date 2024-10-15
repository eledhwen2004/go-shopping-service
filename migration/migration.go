package migration

import (
	"shopping-clone/database"
	"shopping-clone/postgre"
)

func MigrateDB() {
	postgre.DB.AutoMigrate(database.User{})
	postgre.DB.AutoMigrate(database.Customer{})
	postgre.DB.AutoMigrate(database.Supplier{})
	postgre.DB.AutoMigrate(database.Product{})
	postgre.DB.AutoMigrate(database.Order{})
	postgre.DB.AutoMigrate(database.Comment{})
	postgre.DB.AutoMigrate(database.CreditCard{})
}
