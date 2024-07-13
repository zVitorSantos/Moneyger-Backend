package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/models"
    "gorm.io/gorm"
)

type TransactionHandler struct {
    DB *gorm.DB
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
    return &TransactionHandler{DB: db}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
    var transactions []models.Transaction
    if err := h.DB.Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
    id := c.Param("id")
    var transaction models.Transaction
    if err := h.DB.First(&transaction, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
    id := c.Param("id")
    var transaction models.Transaction
    if err := h.DB.First(&transaction, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Save(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Transaction{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
