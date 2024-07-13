package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/models"
    "gorm.io/gorm"
)

type BudgetHandler struct {
    DB *gorm.DB
}

func NewBudgetHandler(db *gorm.DB) *BudgetHandler {
    return &BudgetHandler{DB: db}
}

func (h *BudgetHandler) CreateBudget(c *gin.Context) {
    var budget models.Budget
    if err := c.ShouldBindJSON(&budget); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&budget).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, budget)
}

func (h *BudgetHandler) GetBudgets(c *gin.Context) {
    var budgets []models.Budget
    if err := h.DB.Find(&budgets).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, budgets)
}

func (h *BudgetHandler) GetBudget(c *gin.Context) {
    id := c.Param("id")
    var budget models.Budget
    if err := h.DB.First(&budget, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, budget)
}

func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
    id := c.Param("id")
    var budget models.Budget
    if err := h.DB.First(&budget, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := c.ShouldBindJSON(&budget); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Save(&budget).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, budget)
}

func (h *BudgetHandler) DeleteBudget(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Budget{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Budget deleted"})
}
