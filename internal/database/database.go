package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct DSN for PostgreSQL connection
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")

	// Connect to PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate schemas
	err = db.AutoMigrate(
		&models.Transaction{},
		&models.DebtLoan{},
		&models.Category{},
		&models.Account{},
		&models.Budget{},
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	return db
}
