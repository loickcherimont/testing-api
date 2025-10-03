package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/loickcherimont/testing-api/model"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Load env variables
func init() {
	gotenv.Load()
	var err error
	var (
		DBUSER     string = os.Getenv("DBUSER")
		DBPASSWORD string = os.Getenv("DBPASSWORD")
		DBNAME     string = os.Getenv("DBNAME")
	)
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Paris", DBUSER, DBPASSWORD, DBNAME)
	fmt.Println(dsn)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create table if not exists
	db.AutoMigrate(&model.Test{})

	// Mock data in DB
	tests := []model.Test{
		{
			Status: "Passed",
			Title:  "Database connection check",
			Details: []string{
				"Verify that the application can successfully connect to the database using the configured credentials.",
				"Ensure that connection pooling works correctly and no errors occur during multiple concurrent connections.",
				"Test the database timeout behavior when the database is slow or unreachable.",
			},
		},
		{
			Status: "Failed",
			Title:  "Login with invalid credentials",
			Details: []string{
				"Attempt to login using incorrect username and/or password combinations.",
				"Check that the system displays appropriate error messages without revealing sensitive information.",
				"Ensure that account lockout policies are enforced after multiple failed attempts.",
			},
		},
		{
			Status: "Open",
			Title:  "Search performance test",
			Details: []string{
				"Measure the response time of the search functionality under normal load conditions.",
				"Validate that search results are accurate and relevant according to the query parameters.",
				"Test search performance under high load and ensure system stability and acceptable response times.",
			},
		},
		{
			Status: "Blocked",
			Title:  "Stripe payment validation",
			Details: []string{
				"Verify that the payment process integrates correctly with Stripe's API in test mode.",
				"Check that successful transactions update order status correctly in the application.",
				"Handle error scenarios such as declined payments or network failures and ensure proper user feedback.",
			},
		},
		{
			Status: "Open",
			Title:  "CSV data export",
			Details: []string{
				"Ensure that users can export data to CSV format from different sections of the application.",
				"Verify that the CSV files contain correct headers and all expected data fields.",
				"Test large data exports for performance and file integrity, ensuring no data is truncated or corrupted.",
			},
		},
	}

	// batch creation
	db.Create(&tests)

	fmt.Println("Tables created with success")
}

func FindAll() []model.Test {
	var tests []model.Test
	result := db.Find(&tests)
	if err := result.Error; err != nil {
		log.Printf("FindAll error: %s", err.Error())
		return nil
	}
	fmt.Println(tests)
	return tests
}

func FindById(id uint) model.Test {
	var test model.Test
	result := db.First(&test, id)
	if err := result.Error; err != nil {
		log.Printf("FindById error: %s", err.Error())
		return model.Test{}
	}
	return test
}
