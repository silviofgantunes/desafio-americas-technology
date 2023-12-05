// ./order-service/handlers/handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ListOrders lists all orders
// @Summary List all orders
// @Description List all orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func ListOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var orders []models.Order
	if err := db.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// ListOrdersByUser lists user's orders by its ID
// @Summary List orders by user
// @Description List orders by user
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} models.Order
// @Router /orders/user/{user_id} [get]
func ListOrdersByUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.Param("user_id")

	var orders []models.Order
	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder gets an order by ID
// @Summary Get an order by ID
// @Description Get order details by providing the order ID
// @ID get-order-by-id
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /orders/{id} [get]
func GetOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orderID := c.Param("id")

	var order models.Order
	if err := db.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found."})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order
// @Summary Create a new order
// @Description Create a new order with the provided details
// @ID create-order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order details"
// @Success 201 {object} models.Order
// @Failure 400 {object} utils.ErrorResponse
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req models.CreateOrderRequest

	// Retrieve the Bearer token from the Authorization header
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing."})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar a existência do usuário
	userName, err := checkUserExistence(req.UserID, bearerToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		UserName:  userName, // Atribuir o nome do usuário verificado
		Pair:      req.Pair,
		Amount:    req.Amount,
		Direction: req.Direction,
		Type:      req.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func checkUserExistence(userID string, bearerToken string) (string, error) {
	// URL do endpoint check-user no serviço crud-users
	hostIP := "host.docker.internal"

	url := fmt.Sprintf("http://%s:8080/api/v1/users/%s", hostIP, userID)

	// Fazer uma requisição HTTP GET para o endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Adicionar o token ao cabeçalho da requisição
	req.Header.Set("Authorization", bearerToken)

	// Executar a requisição
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Ler o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Verificar se o status da resposta é OK (200)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to check user's existence: %s", body)
	}

	// Extrair o nome do usuário do corpo da resposta
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	userName, ok := response["name"].(string)
	if !ok {
		return "", fmt.Errorf("UserName not found in response.")
	}

	return userName, nil
}

// DeleteLimitOrder deletes an order by ID
// @Summary Delete an order by ID
// @Description Delete an order by providing the order ID
// @ID delete-order
// @Produce json
// @Param id path int true "Order ID"
// @Success 204 "No Content"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /orders/{id} [delete]
func DeleteLimitOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orderID := c.Param("id")

	var order models.Order

	// Check if the order exists with "type" as "market"
	if err := db.First(&order, "id = ? AND type = ?", orderID, "market").Error; err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Market orders not supposed to be deleted."})
		return
	}

	// Check if the order exists with "type" as "limit"
	if err := db.First(&order, "id = ? AND type = ?", orderID, "limit").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Limit order not found."})
		return
	}

	// Delete the limit order
	if err := db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Limit order deleted successfully."})
}
