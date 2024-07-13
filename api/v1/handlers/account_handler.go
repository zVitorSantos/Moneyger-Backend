package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/gorm"
)

type AccountHandler struct {
	DB *gorm.DB
}

func NewAccountHandler(db *gorm.DB) *AccountHandler {
	return &AccountHandler{DB: db}
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.Account true "Account"
// @Success 201 {object} models.Account
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

// GetAccounts godoc
// @Summary Get all accounts
// @Description Get all accounts
// @Tags accounts
// @Produce json
// @Success 200 {array} models.Account
// @Failure 500 {object} string "Internal Server Error"
// @Router /accounts [get]
func (h *AccountHandler) GetAccounts(c *gin.Context) {
	var accounts []models.Account
	if err := h.DB.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// GetAccount godoc
// @Summary Get an account by ID
// @Description Get an account by ID
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account
// @Failure 404 {object} string "Account not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := h.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
// @Summary Update an account by ID
// @Description Update an account by ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body models.Account true "Account"
// @Success 200 {object} models.Account
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Account not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /accounts/{id} [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.Account
	if err := h.DB.First(&account, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
// @Summary Delete an account by ID
// @Description Delete an account by ID
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} string "Account deleted"
// @Failure 500 {object} string "Internal Server Error"
// @Router /accounts/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Account{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}
