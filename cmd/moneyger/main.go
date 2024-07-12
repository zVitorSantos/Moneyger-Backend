package main

import (
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/database"
    "github.com/zVitorSantos/Moneyger-Backend/api/v1/routes"
    
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/zVitorSantos/Moneyger-Backend/docs" 
)

// @title Moneyger API
// @version 1.0
// @description This is a erver for a financial management application.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1

func main() {
    // Initialize database
    db := database.Init()

    // Setup router
    r := routes.SetupRouter(db)

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}
