package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/zVitorSantos/Moneyger-Backend/internal/models"
    "gorm.io/gorm"
)

type CategoryHandler struct {
    DB *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
    return &CategoryHandler{DB: db}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Create(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
    var categories []models.Category
    if err := h.DB.Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
    id := c.Param("id")
    var category models.Category
    if err := h.DB.First(&category, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
    id := c.Param("id")
    var category models.Category
    if err := h.DB.First(&category, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.DB.Save(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Category{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
