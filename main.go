package main

import (
	"shopping-clone/database"
	"shopping-clone/postgre"
)

func main() {
	postgre.ConnectDatabase()

	postgre.DB.AutoMigrate(database.User{})
	postgre.DB.AutoMigrate(database.Comment{})
	postgre.DB.AutoMigrate(database.Order{})
	postgre.DB.AutoMigrate(database.Customer{})
	postgre.DB.AutoMigrate(database.Supplier{})
	postgre.DB.AutoMigrate(database.Comment{})
	postgre.DB.AutoMigrate(database.CreditCard{})

	/*
			users := []database.User{
				{
					Username:    "admin_user",
					Password:    "admin_password",
					Email:       "admin@example.com",
					PhoneNumber: "+123456789",
					Role:        database.AdminRole,
					Warnings:    0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Username:    "supplier_user",
					Password:    "supplier_password",
					Email:       "supplier@example.com",
					PhoneNumber: "+987654321",
					Role:        database.SupplierRole,
					Warnings:    1,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Username:    "customer_user",
					Password:    "customer_password",
					Email:       "customer@example.com",
					PhoneNumber: "+555555555",
					Role:        database.CustomerRole,
					Warnings:    0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}

		 Mock customers

		customer := database.Customer{
			UserID:    users[2].ID, // Linking with customer user
			User:      users[2],
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 Customer St, Customer City",
		}
			database.CreateCustomer(&customer)

		supplier := database.Supplier{
			User:        users[1],
			CompanyName: "Supplier Co.",
			Address:     "456 Supplier Ave, Supplier City",
		}

		database.CreateSupplier(&supplier)

		 Mock products
		products := []database.Product{
			{
				Name:        "Gaming Laptop",
				Description: "High-end gaming laptop with NVIDIA graphics and Intel i9 processor.",
				Price:       1500.99,
				Amount:      20,
				Rating:      4.5,
			},
			{
				Name:        "Smartphone",
				Description: "Latest smartphone with 5G connectivity and high-resolution camera.",
				Price:       999.99,
				Amount:      50,
				Rating:      4.7,
			},
			{
				Name:        "Wireless Headphones",
				Description: "Noise-cancelling wireless headphones with 40 hours of battery life.",
				Price:       199.99,
				Amount:      100,
				Rating:      4.3,
			},
		}

			database.CreateProduct(&products[0])
			database.CreateProduct(&products[1])
			database.CreateProduct(&products[2])

		 Mock credit cards

			creditCard := database.CreditCard{
				CardholderName:  "John Doe",
				LastFourDigits:  "1234",
				ExpirationMonth: 12,
				ExpirationYear:  2025,
				CardType:        "Visa",
				BillingAddress:  "123 Customer St, Customer City",
				CustomerID:      customer.ID, // Linking to the first customer
				CreatedAt:       time.Now(),
			}
			database.CreateCreditCard(&creditCard)
	*/
	comment := database.Comment{
		Content:    "Bu nasil urun aq dolandırici misiniz nesiniz",
		CustomerID: "2ff82e13-914d-4000-8d4e-e420c023702a", // Linking to first customer (John Doe)
		ProductID:  "f343a9a2-5e3a-4040-b0d0-c5c62b7b2147", // Linking to the first product (Gaming Laptop)
	}

	database.CreateComment(&comment)

	// app := fiber.New()

	// app.Listen(":3000")
}
