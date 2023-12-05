// ./order-service/handlers/handler_test.go

package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	user_models "crud-users/models"
	"order-service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openSQLiteDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to set up database: %v", err)
	}
	return db
}

func AutoMigrateSQLiteDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Order{})
	if err != nil {
		panic("failed to run migrations")
	}
	err = db.AutoMigrate(&user_models.User{})
	if err != nil {
		panic("failed to run migrations")
	}
}

func createTestContext(method, url string, payload interface{}, pathParams gin.Params) (*gin.Context, *httptest.ResponseRecorder, error) {
	// Convert payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	// Create HTTP request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}

	// Create Gin context
	w := httptest.NewRecorder()

	if w.Code == 200 && method == "DELETE" {
		w.Code = 204
	} // There's a technical limitation when using httptest's NewRecorder function for "DELETE" requests. That's why
	// this if is here.

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set path parameters
	c.Params = pathParams

	return c, w, nil
}

func TestListOrders(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := user_models.User{
		ID:          uuid.New().String(),
		Name:        "User1",
		Email:       "user1@example.com",
		PhoneNumber: "111111111",
	}
	db.Create(&testUser)

	testOrders := []models.Order{
		{ID: uuid.New().String(), UserID: testUser.ID, UserName: testUser.Name, Pair: "BTC/USD", Amount: 1.5, Direction: "buy", Type: "limit", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New().String(), UserID: testUser.ID, UserName: testUser.Name, Pair: "ETH/USD", Amount: 2.0, Direction: "sell", Type: "market", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	db.Create(&testOrders)

	gin.SetMode(gin.TestMode)
	c, w, err := createTestContext("GET", "/orders", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	ListOrders(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestListOrdersByUser(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUsers := []user_models.User{
		{ID: uuid.New().String(), Name: "User1", Email: "user1@example.com", PhoneNumber: "111111111"},
		{ID: uuid.New().String(), Name: "User2", Email: "user2@example.com", PhoneNumber: "222222222"},
	}
	db.Create(&testUsers)

	testOrders := []models.Order{
		{ID: uuid.New().String(), UserID: testUsers[0].ID, UserName: "User1", Pair: "BTC/USD", Amount: 1.5, Direction: "buy", Type: "limit", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: uuid.New().String(), UserID: testUsers[1].ID, UserName: "User2", Pair: "ETH/USD", Amount: 2.0, Direction: "sell", Type: "market", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	db.Create(&testOrders)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{
		gin.Param{Key: "id", Value: testUsers[0].ID},
		gin.Param{Key: "id", Value: testUsers[1].ID},
	}
	c, w, err := createTestContext("GET", "/orders/user/"+testUsers[0].ID, nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	ListOrdersByUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestGetOrder(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := user_models.User{
		ID:          uuid.New().String(),
		Name:        "User1",
		Email:       "user1@example.com",
		PhoneNumber: "111111111",
	}
	db.Create(&testUser)

	testOrder := models.Order{
		ID:        uuid.New().String(),
		UserID:    testUser.ID,
		UserName:  testUser.Name,
		Pair:      "BTC/USD",
		Amount:    1.5,
		Direction: "buy",
		Type:      "limit",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testOrder)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testOrder.ID}}
	c, w, err := createTestContext("GET", "/orders/"+testOrder.ID, nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	GetOrder(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestDeleteOrder(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := user_models.User{
		ID:          uuid.New().String(),
		Name:        "User1",
		Email:       "user1@example.com",
		PhoneNumber: "111111111",
	}
	db.Create(&testUser)

	testOrder := models.Order{
		ID:        uuid.New().String(),
		UserID:    testUser.ID,
		Pair:      "BTC/USD",
		Amount:    1.5,
		Direction: "buy",
		Type:      "limit",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testOrder)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testOrder.ID}}
	c, w, err := createTestContext("DELETE", "/orders/"+testOrder.ID, nil, pathParams)

	if err != nil {
		log.Fatal(err)
	}

	c.Set("db", db)

	DeleteLimitOrder(c)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var deletedOrder models.Order
	result := db.First(&deletedOrder, "id = ?", testOrder.ID)
	assert.ErrorIs(t, result.Error, gorm.ErrRecordNotFound)
	// Add more assertions based on your expected behavior

	testOrderMarket := models.Order{
		ID:        uuid.New().String(),
		UserID:    testUser.ID,
		Pair:      "ETH/USD",
		Amount:    2.0,
		Direction: "sell",
		Type:      "market", // Order with "type" as "market"
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testOrderMarket)

	// Try to delete an order with "type" as "market"
	pathParamsMarket := gin.Params{gin.Param{Key: "id", Value: testOrderMarket.ID}}
	cMarket, wMarket, errMarket := createTestContext("DELETE", "/orders/"+testOrderMarket.ID, nil, pathParamsMarket)

	if errMarket != nil {
		log.Fatal(errMarket)
	}

	cMarket.Set("db", db)

	DeleteLimitOrder(cMarket)

	// Assert that deletion of orders with "type" as "market" is not allowed
	assert.Equal(t, http.StatusForbidden, wMarket.Code)

	var stillExists models.Order
	result = db.First(&stillExists, "id = ?", testOrderMarket.ID)
	assert.NoError(t, result.Error) // The order should still exist
}
