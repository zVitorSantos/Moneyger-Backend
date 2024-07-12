package main

import (
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/database"
    "github.com/zVitorSantos/Moneyger-Backend/api/v1/handlers"
)

func main() {
    r := gin.Default()

    // Initialize database
    db := database.Init()

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
        v1.GET("/accounts", accountHandler.GetAccounts)
        v1.GET("/budgets", budgetHandler.GetBudgets)
        v1.GET("/categories", categoryHandler.GetCategories)
        v1.GET("/debts-loans", debtLoanHandler.GetDebtLoans)
        v1.GET("/transactions", transactionHandler.GetTransactions)
    }

    r.Run(":8080")
}
