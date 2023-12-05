// ./auth-service/handlers/handler_test.go

package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"auth-service/models"

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
	err := db.AutoMigrate(&models.Admin{})
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

func TestGenerateToken(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	// Create a test admin user in the database
	testAdmin := models.Admin{
		Name:      "Test Admin",
		Email:     "testadmin@example.com",
		Password:  "test_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testAdmin)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testAdmin.ID}}
	c, w, err := createTestContext("GET", "/generate-token", map[string]interface{}{
		"email":    "testadmin@example.com",
		"password": "test_password",
	}, pathParams)

	if err != nil {
		log.Fatal(err)
	}

	c.Set("db", db)

	GenerateToken(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestCreateAdmin(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: "1"}}
	c, w, err := createTestContext("POST", "/admins", models.CreateAdminRequest{
		Name:     "Test Admin",
		Email:    "testadmin@example.com",
		Password: "test_password",
	}, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	CreateAdmin(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdAdmin models.Admin
	result := db.First(&createdAdmin, "email = ?", "testadmin@example.com")
	assert.NoError(t, result.Error)
	assert.Equal(t, "Test Admin", createdAdmin.Name)
}

func TestListAdmins(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testAdmins := []models.Admin{
		{ID: uuid.New().String(), Name: "Admin1", Email: "admin1@example.com", Password: "password1"},
		{ID: uuid.New().String(), Name: "Admin2", Email: "admin2@example.com", Password: "password2"},
	}
	db.Create(&testAdmins)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{
		gin.Param{Key: "id", Value: testAdmins[0].ID},
		gin.Param{Key: "id", Value: testAdmins[1].ID},
	}
	c, w, err := createTestContext("GET", "/admins", nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	ListAdmins(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestGetAdmin(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testAdmin := models.Admin{
		ID:        uuid.New().String(),
		Name:      "Test Admin",
		Email:     "testadmin@example.com",
		Password:  "test_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testAdmin)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testAdmin.ID}}
	c, w, err := createTestContext("GET", "/admins/"+testAdmin.ID, nil, pathParams)
	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	GetAdmin(c)

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestUpdateAdmin(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testAdmin := models.Admin{
		ID:        uuid.New().String(),
		Name:      "Test Admin",
		Email:     "testadmin@example.com",
		Password:  "test_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testAdmin)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testAdmin.ID}}
	c, w, err := createTestContext("PUT", "/admins/"+testAdmin.ID, models.UpdateAdminRequest{
		Name:     "Updated Admin",
		Email:    "updatedadmin@example.com",
		Password: "updated_password",
	}, pathParams)

	if err != nil {
		log.Fatal(err)
	}
	c.Set("db", db)

	UpdateAdmin(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedAdmin models.Admin
	db.First(&updatedAdmin, "email = ?", "updatedadmin@example.com")
	assert.Equal(t, "Updated Admin", updatedAdmin.Name)
	// Add more assertions based on your expected behavior
}

func TestDeleteAdmin(t *testing.T) {
	db := openSQLiteDB(t)
	AutoMigrateSQLiteDB(db)

	testAdmin := models.Admin{
		ID:        uuid.New().String(),
		Name:      "Test Admin",
		Email:     "testadmin@example.com",
		Password:  "test_password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&testAdmin)

	gin.SetMode(gin.TestMode)
	pathParams := gin.Params{gin.Param{Key: "id", Value: testAdmin.ID}}
	c, w, err := createTestContext("DELETE", "/admins/"+testAdmin.ID, nil, pathParams)

	if err != nil {
		log.Fatal(err)
	}

	c.Set("db", db)

	DeleteAdmin(c)

	assert.Equal(t, http.StatusNoContent, w.Code)

	var deletedAdmin models.Admin
	result := db.First(&deletedAdmin, "email = ?", "testadmin@example.com")
	assert.ErrorIs(t, result.Error, gorm.ErrRecordNotFound)
	// Add more assertions based on your expected behavior
}
