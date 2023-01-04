package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vithor-vbs/API-Requests-Example/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestGetUsers(t *testing.T) {
	r := SetUpRouter()
	r.GET("/user", GetUsers)
	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}
