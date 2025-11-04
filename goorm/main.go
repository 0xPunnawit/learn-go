package main

import (
	"goorm/database"
	"goorm/repository"

	"log"
)

func main() {

	database.ConnectDB()

	log.Println("🚀 Running Migrations")

	// err := database.DB.AutoMigrate(&models.Gender{}, &models.Test{}, &models.Customer{})
	// if err != nil {
	// 	log.Fatalf("Failed td migrate database: %v", err)
	// }

	// if err := repository.CreateGenders("Female"); err != nil {
	// 	log.Fatalf("❌ Failed to create gender: %v", err)
	// }

	// repository.GetGenders()

	// repository.GetGender(1)

	// repository.GetGenderByName("Female")

	// repository.DeleteGender(4)

	// repository.CreateCustomer("Ying", 2)

	repository.GetCustomers()

	log.Println("✅ Database migrated successfully!")

}
