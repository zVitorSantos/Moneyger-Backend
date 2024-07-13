package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/models"
    "gorm.io/gorm"
)

type DebtLoanHandler struct {
    DB *gorm.DB
}

func NewDebtLoanHandler(db *gorm.DB) *DebtLoanHandler {
    return &DebtLoanHandler{DB: db}
}

func (h *DebtLoanHandler) CreateDebtLoan(c *gin.Context) {
    var debtLoan models.DebtLoan
    if err := c.ShouldBindJSON(&debtLoan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&debtLoan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, debtLoan)
}

func (h *DebtLoanHandler) GetDebtLoans(c *gin.Context) {
    var debtLoans []models.DebtLoan
    if err := h.DB.Find(&debtLoans).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, debtLoans)
}

func (h *DebtLoanHandler) GetDebtLoan(c *gin.Context) {
    id := c.Param("id")
    var debtLoan models.DebtLoan
    if err := h.DB.First(&debtLoan, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "DebtLoan not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, debtLoan)
}

func (h *DebtLoanHandler) UpdateDebtLoan(c *gin.Context) {
    id := c.Param("id")
    var debtLoan models.DebtLoan
    if err := h.DB.First(&debtLoan, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "DebtLoan not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := c.ShouldBindJSON(&debtLoan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Save(&debtLoan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, debtLoan)
}

func (h *DebtLoanHandler) DeleteDebtLoan(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.DebtLoan{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "DebtLoan deleted"})
}
