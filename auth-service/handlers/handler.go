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

// GenerateToken gera um Bearer Token
// @Summary Generate a Bearer Token
// @Description Generate a Bearer Token using admin credentials
// @Accept json
// @Produce json
// @Param credentials body GenerateTokenRequest true "Admin credentials"
// @Success 200 {object} GenerateTokenResponse
// @Router /generate-token [post]
// GenerateToken gera um Bearer Token
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

	// Verificar se o admin com as credenciais fornecidas existe no banco de dados
	var admin models.Admin
	if err := db.Where("email = ? AND password = ?", credentials.Email, credentials.Password).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Criar o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": credentials.Email,
		"exp":   time.Now().Add(time.Minute * 3).Unix(), // Token válido por 3 minutos
	})

	// Assinar o token com uma chave secreta
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

// CreateAdmin cria um novo usuário admin
// @Summary Create a new admin user
// @Description Create a new admin user
// @Accept json
// @Produce json
// @Param admin body CreateAdminRequest true "Admin details"
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

// ListAdmins lista todos os usuários admin
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

// GetAdmin exibe um usuário admin por ID
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

// UpdateAdmin atualiza um usuário admin por ID
// @Summary Update an admin user by ID
// @Description Update an admin user by ID
// @Accept json
// @Produce json
// @Param id path string true "Admin ID"
// @Param admin body UpdateAdminRequest true "Admin details"
// @Success 200 {object} models.Admin
// @Router /admins/{id} [put]
func UpdateAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	adminID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var req models.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	if err := db.First(&admin, "id = ?", adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	admin.Name = req.Name
	admin.Email = req.Email
	admin.Password = req.Password
	admin.UpdatedAt = time.Now()

	if err := db.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// DeleteAdmin exclui um usuário admin por ID
// @Summary Delete an admin user by ID
// @Description Delete an admin user by ID
// @Produce json
// @Param id path string true "Admin ID"
// @Success 200 {object}
// @Router /admins/{id} [delete]
func DeleteAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	adminID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var admin models.Admin
	if err := db.First(&admin, "id = ?", adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	if err := db.Delete(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
