package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Create Setup Database Function
func SetupDatabase() *sql.DB {
	// Load ENV
	err := godotenv.Load()

	// Check if there is error when load ENV
	if err != nil {
		log.Fatal("Failed to Load ENV")
	}

	// Assign ENV value to Local Variable
	dbUsername := os.Getenv("MYSQL_USERNAME")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")

	// Database URL ( Data Source Name )
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to Database
	db, err := sql.Open("mysql", dsn)

	// Check if there is error when connecting to database
	if err != nil {
		log.Fatal("Failed connect to database: ", err.Error())
	}

	// Double Check Database Connection ( As Recomended by Official Documentation )
	if errorPing := db.Ping(); err != nil {
		log.Fatal("Failed connect to database: ", errorPing.Error())
	}

	// Return DB Instance
	return db
}
