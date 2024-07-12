package handlers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type AuthHandler struct {
    db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
    return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}
