package services

import (
	"log"
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"github.com/ashish-019-hash/obp-api-backend/internal/repositories"
	"gorm.io/gorm"
)

func SeedAuthenticationData(db *gorm.DB, authRepo repositories.AuthRepository) error {
	log.Println("Seeding authentication data...")

	testConsumer := models.NewConsumer(
		"test_consumer_key_123",
		"test_consumer_secret_456",
		"Test Banking App",
		"developer@testbank.com",
	)
	testConsumer.Description = "Test consumer for development and testing"
	testConsumer.RedirectURL = "http://localhost:3000/callback"
	testConsumer.AppType = "Web"

	if err := authRepo.CreateConsumer(testConsumer); err != nil {
		log.Printf("Consumer already exists or error creating: %v", err)
	} else {
		log.Printf("Created test consumer with key: %s", testConsumer.ConsumerKey)
	}

	testUser := &models.User{
		UserID:       "test_user_001",
		Email:        "testuser@example.com",
		FirstName:    "Test",
		LastName:     "User",
		Provider:     "local",
		ProviderID:   "testuser",
		IsActive:     true,
		ConsentGiven: true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := db.Create(testUser).Error; err != nil {
		log.Printf("Test user already exists or error creating: %v", err)
	} else {
		log.Printf("Created test user: %s", testUser.Email)
	}

	testCredential, err := models.NewUserCredential(testUser.UserID, "testuser", "password123")
	if err != nil {
		log.Printf("Error creating test credentials: %v", err)
		return err
	}

	if err := authRepo.CreateUserCredential(testCredential); err != nil {
		log.Printf("Test credentials already exist or error creating: %v", err)
	} else {
		log.Printf("Created test credentials for user: %s", testCredential.Username)
	}

	testEntitlement := models.NewEntitlement(testUser.UserID, "CanGetBanks", nil)
	if err := authRepo.CreateEntitlement(testEntitlement); err != nil {
		log.Printf("Test entitlement already exists or error creating: %v", err)
	} else {
		log.Printf("Created test entitlement: %s", testEntitlement.RoleName)
	}

	adminEntitlement := models.NewEntitlement(testUser.UserID, "CanGetApiCollections", nil)
	if err := authRepo.CreateEntitlement(adminEntitlement); err != nil {
		log.Printf("Admin entitlement already exists or error creating: %v", err)
	} else {
		log.Printf("Created admin entitlement: %s", adminEntitlement.RoleName)
	}

	log.Println("Authentication data seeding completed!")
	log.Println("Test credentials:")
	log.Println("  Username: testuser")
	log.Println("  Password: password123")
	log.Println("  Consumer Key: test_consumer_key_123")
	log.Println("  Consumer Secret: test_consumer_secret_456")

	return nil
}
