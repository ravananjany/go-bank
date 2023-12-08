package server

import (
	"go-bank/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

var tk string

func setupDB() {
	ac := []models.Account{{
		Id:      "b2fd5a7d-d46c-4b57-a167-b802d4d37ffe",
		Name:    "Trupe",
		Balance: "4000",
	}, {
		Id:      "b2fd5a7d-d46c-4b57-a167-b802d4d37fde",
		Name:    "john",
		Balance: "5000.78",
	}}
	Repository.UploadAccounts(ac)
}

// func TestHealthRouter(t *testing.T) {
// 	router := setupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/health", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

func TestAccounts(t *testing.T) {
	r := gin.Default()
	r.GET("/accounts", GetAccounts)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts", nil)
	req.Header.Set("Auth", "Trupe")
	buildRoutes()
	setupDB()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
