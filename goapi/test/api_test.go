// Run command => go test test/api_test.go
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goapi/api/controllers"
	"goapi/api/middlewares"
	"goapi/api/models"
	"goapi/api/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// router from main.go
	r := gin.Default()

	// Database config
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_NAME,
	)

	// Connect to database
	db, _ := sqlx.Connect(DB_DRIVER, dsn)

	// APIkeys Entity
	apiRepo := models.NewAPIkeyRepositoryDB(db)
	apiService := service.NewAPIkeyService(apiRepo)
	apiHandler := controllers.NewAPIkeyHandler(apiService)

	apiRoute := r.Group("/api")
	apiRoute.Use(middlewares.CheckCache()).POST("/handshake", apiHandler.GetUserFromKey)

	return r
}

// Save mac address in redis by username (key)
// Get mac address from redis
func TestGetUserFromKey(t *testing.T) {
	r := setupRouter()

	body := []byte(`{
		"uname": "last_order",
		"api_key": "ZG9sbHk7pJIhnO3ppHQvMS2-P9VR5XS2"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/handshake", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.Bytes()), &resp)

	// Test case ; All true
	// Assert field of response body
	assert.Nil(t, err)
	if assert.Equal(t, "Ready", resp["message"]) {
		assert.Equal(t, "94:B9:7E:D5:AD:F4", resp["robot"])
	} else {
		assert.Equal(t, "Connected", resp["message"]) // For => Key not found in redis
		assert.Equal(t, "94:B9:7E:D5:AD:F4", resp["robot"])
	}
}
