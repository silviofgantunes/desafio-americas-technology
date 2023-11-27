package handlers

import (
	"net/http"
	"order-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

// @Summary Get an order by ID
// @Description Get an order by ID
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
func GetOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orderID := c.Param("id")

	var order models.Order
	if err := db.First(&order, "id = ?", orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary Create a new order
// @Description Create a new order
// @Accept json
// @Produce json
// @Param order body models.CreateOrderRequest true "Order details"
// @Success 201 {object} models.Order
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req models.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
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

// @Summary Delete a limit order by ID
// @Description Delete a limit order by ID
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} gin.H
// @Router /orders/limit/{id} [delete]
func DeleteLimitOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	orderID := c.Param("id")

	var order models.Order
	if err := db.First(&order, "id = ? AND type = ?", orderID, "limit").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Limit order not found"})
		return
	}

	if err := db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Limit order deleted successfully"})
}
