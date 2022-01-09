package utils

import (
	"log"
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/repositories"
)

func SeedAdminUser(repository repositories.UserRepository) {
	// Create Admin 1
	adminOne := models.User{Name: "Admin 1", UserName: "admin1", Password: "admin1"}

	// Create Admin 2
	adminTwo := models.User{Name: "Admin 2", UserName: "admin2", Password: "admin2"}

	// List Admin
	listAdmin := []models.User{
		adminOne, adminTwo,
	}

	// Loop to Create Admin
	for _, admin := range listAdmin {
		// Check if admin is already on database
		isAdminExists, err := repository.CheckIfUserExistsByUsername(admin.UserName)
		if isAdminExists {
			return
		}

		// If there is Query error
		if err != nil {
			log.Fatal("Failed to create Admin: query")
		}

		// Hash Password
		hashedPassword, err := GeneratePassword(admin.Password)

		// Assign Hashed Password
		admin.Password = hashedPassword

		// Check if there is error when hashing password
		if err != nil {
			log.Fatal("Failed to create Admin: hash")
		}

		// Save to Database
		errorCreate := repository.CreateUser(admin)

		// Check if there is error when create admin
		if errorCreate != nil {
			log.Fatal("Failed to create Admin: database")
		}
	}
}
