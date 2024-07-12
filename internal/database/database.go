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
    // Load enviroment variables in the .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dsn := "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=" + os.Getenv("DB_SSLMODE")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // Migrate schema
    db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{}, &models.Category{}, &models.Budget{}, &models.DebtLoan{})

    return db
}
