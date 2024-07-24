package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/Moneyger-Backend/api/v1/handlers"
	"github.com/zVitorSantos/Moneyger-Backend/internal/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize handlers with database
	accountHandler := handlers.NewAccountHandler(db)
	authHandler := handlers.NewAuthHandler(db)
	budgetHandler := handlers.NewBudgetHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	debtLoanHandler := handlers.NewDebtLoanHandler(db)
	transactionHandler := handlers.NewTransactionHandler(db)

	// Define routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)

		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/accounts", accountHandler.GetAccounts)
			protected.GET("/accounts/:id", accountHandler.GetAccount)
			protected.POST("/accounts", accountHandler.CreateAccount)
			protected.PUT("/accounts/:id", accountHandler.UpdateAccount)
			protected.DELETE("/accounts/:id", accountHandler.DeleteAccount)
			protected.GET("/budgets", budgetHandler.GetBudgets)
			protected.GET("/categories", categoryHandler.GetCategories)
			protected.GET("/debts-loans", debtLoanHandler.GetDebtLoans)
			protected.GET("/transactions", transactionHandler.GetTransactions)
		}
	}

	return r
}
