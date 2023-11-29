// ./auth-service/handlers/handler.go

package handlers

import (
	"auth-service/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GenerateToken generates a Bearer Token using admin credentials
// @Summary Generate a Bearer Token
// @Description Generate a Bearer Token using admin credentials and returns the token
// @Accept json
// @Produce json
// @Router /generate-token [post]
func GenerateToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the admin with the provided credentials exists in the database
	var admin models.Admin
	if err := db.Where("email = ? AND password = ?", credentials.Email, credentials.Password).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": credentials.Email,
		"exp":   time.Now().Add(time.Minute * 3).Unix(), // Token valid for 3 minutes
	})

	// Sign the token with a secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")

	// Convert the secretKey to []byte
	secretKeyBytes := []byte(secretKey)

	tokenString, err := token.SignedString(secretKeyBytes)

	log.Printf(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// CreateAdmin creates a new admin user
// @Summary Create a new admin user
// @Description Create a new admin user
// @ID create-order
// @Accept json
// @Produce json
// @Param admin body models.Admin true "Admin details"
// @Success 201 {object} models.Admin
// @Router /admins [post]
func CreateAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req models.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin := models.Admin{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

// ListAdmins lists all admin users
// @Summary List all admin users
// @Description List all admin users
// @Produce json
// @Success 200 {array} models.Admin
// @Router /admins [get]
func ListAdmins(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var admins []models.Admin
	if err := db.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admins)
}

// GetAdmin displays an admin user by ID
// @Summary Get an admin user by ID
// @Description Get an admin user by ID
// @Produce json
// @Param id path string true "Admin ID"
// @Success 200 {object} models.Admin
// @Router /admins/{id} [get]
func GetAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	adminID := c.Param("id")

	var admin models.Admin
	if err := db.First(&admin, "id = ?", adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// UpdateAdmin updates an admin user by ID
// @Summary Update an admin user by ID
// @Description Update an admin user by ID
// @Accept json
// @Produce json
// @Param id path string true "Admin ID"
// @Param admin body models.UpdateAdminRequest true "Updated admin details"
// @Success 200 {object} models.Admin
// @Router /admins/{id} [put]
func UpdateAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	adminID := c.Param("id")

	// Check if the admin with the given ID exists
	var existingAdmin models.Admin
	if err := db.First(&existingAdmin, "id = ?", adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	var updateReq models.UpdateAdminRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the admin details
	existingAdmin.Name = updateReq.Name
	existingAdmin.Email = updateReq.Email
	existingAdmin.Password = updateReq.Password
	existingAdmin.UpdatedAt = time.Now()

	if err := db.Save(&existingAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAdmin)
}

// DeleteAdmin deletes an admin user by ID
// @Summary Delete an admin user by ID
// @Description Delete an admin user by ID
// @Produce json
// @Param id path string true "Admin ID"
// @Success 204
// @Router /admins/{id} [delete]
func DeleteAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	adminID := c.Param("id")

	// Check if the admin with the given ID exists
	var existingAdmin models.Admin
	if err := db.First(&existingAdmin, "id = ?", adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	// Delete the admin
	if err := db.Delete(&existingAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
