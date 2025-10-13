package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashish-019-hash/obp-api-backend/internal/routes"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	router := routes.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}

func TestPingEndpoint(t *testing.T) {
	router := routes.SetupRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
