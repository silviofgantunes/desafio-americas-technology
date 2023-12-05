// ./crud-users/handlers/handler_test.go

package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"crud-users/models"

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
	err := db.AutoMigrate(&models.User{})
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

func TestGetUser(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := models.User{
		ID:          uuid.New().String(),
		Name:        "Test User",
		Email:       "testuser@example.com",
		PhoneNumber: "123456789",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.Create(&testUser)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testUser.ID}}
	c, w, err := createTestContext("GET", "/users/"+testUser.ID, nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	GetUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestCreateUser(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: "1"}}
	c, w, err := createTestContext("POST", "/users", models.CreateUserRequest{
		Name:        "Test User",
		Email:       "testuser@example.com",
		PhoneNumber: "123456789",
	}, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	CreateUser(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdUser models.User
	result := db.First(&createdUser, "email = ?", "testuser@example.com")
	assert.NoError(t, result.Error)
	assert.Equal(t, "Test User", createdUser.Name)
}

func TestUpdateUser(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := models.User{
		ID:          uuid.New().String(),
		Name:        "Test User",
		Email:       "testuser@example.com",
		PhoneNumber: "123456789",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.Create(&testUser)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testUser.ID}}
	c, w, err := createTestContext("PUT", "/users/"+testUser.ID, models.UpdateUserRequest{
		Name:        "Updated User",
		Email:       "updateduser@example.com",
		PhoneNumber: "987654321",
	}, pathParams)

	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	UpdateUser(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedUser models.User
	db.First(&updatedUser, "email = ?", "updateduser@example.com")
	assert.Equal(t, "Updated User", updatedUser.Name)
	// Add more assertions based on your expected behavior
}

func TestListUsers(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUsers := []models.User{
		{ID: uuid.New().String(), Name: "User1", Email: "user1@example.com", PhoneNumber: "111111111"},
		{ID: uuid.New().String(), Name: "User2", Email: "user2@example.com", PhoneNumber: "222222222"},
	}
	db.Create(&testUsers)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{
		gin.Param{Key: "id", Value: testUsers[0].ID},
		gin.Param{Key: "id", Value: testUsers[1].ID},
	}
	c, w, err := createTestContext("GET", "/users", nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	ListUsers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestDeleteUser(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testUser := models.User{
		ID:          uuid.New().String(),
		Name:        "Test User",
		Email:       "testuser@example.com",
		PhoneNumber: "123456789",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.Create(&testUser)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testUser.ID}}
	c, w, err := createTestContext("DELETE", "/users/"+testUser.ID, nil, pathParams)

	if err != nil {
		log.Fatal(err)
	}

	c.Set("db", db)

	DeleteUser(c)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var deletedUser models.User
	result := db.First(&deletedUser, "email = ?", "testuser@example.com")
	assert.ErrorIs(t, result.Error, gorm.ErrRecordNotFound)
	// Add more assertions based on your expected behavior
}
