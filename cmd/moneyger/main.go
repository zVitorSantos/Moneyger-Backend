package main

import (
	"github.com/zVitorSantos/Moneyger-Backend/api/v1/routes"
	"github.com/zVitorSantos/Moneyger-Backend/internal/database"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zVitorSantos/Moneyger-Backend/docs"
)

func main() {
	// Initialize database
	db := database.Init()

	// Setup router
	r := routes.SetupRouter(db)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
