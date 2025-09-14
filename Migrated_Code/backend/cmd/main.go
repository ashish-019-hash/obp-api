package main

import (
	"log"

	"github.com/ashish-019-hash/obp-api-backend/internal/config"
	"github.com/ashish-019-hash/obp-api-backend/internal/routes"
)

func main() {
	cfg := config.Load()

	router := routes.SetupRoutes()

	port := ":" + cfg.Port
	log.Printf("Starting OBP-API Backend Server on port %s", cfg.Port)
	log.Println("Based on Open Bank Project API analysis from phase-01-output")
	log.Println("Server endpoints:")
	log.Println("  GET /health - Health check")
	log.Println("  GET /ping - Ping endpoint")
	log.Println("  GET /api/v1/health - API health check")

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
