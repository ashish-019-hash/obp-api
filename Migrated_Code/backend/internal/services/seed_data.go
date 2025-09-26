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

	testScope := models.NewScope(testConsumer.ConsumerID, "CanGetBanks", nil)
	if err := authRepo.CreateScope(testScope); err != nil {
		log.Printf("Test scope already exists or error creating: %v", err)
	} else {
		log.Printf("Created test scope: %s", testScope.RoleName)
	}

	adminScope := models.NewScope(testConsumer.ConsumerID, "CanGetApiCollections", nil)
	if err := authRepo.CreateScope(adminScope); err != nil {
		log.Printf("Admin scope already exists or error creating: %v", err)
	} else {
		log.Printf("Created admin scope: %s", adminScope.RoleName)
	}

	testViewPermission := models.NewViewPermission("owner", "can_see_transaction_amount", nil, nil)
	if err := authRepo.CreateViewPermission(testViewPermission); err != nil {
		log.Printf("Test view permission already exists or error creating: %v", err)
	} else {
		log.Printf("Created test view permission: %s", testViewPermission.PermissionName)
	}

	testAuthContext := models.NewUserAuthContext(testUser.UserID, testConsumer.ConsumerID, "auth_method", "DirectLogin")
	if err := authRepo.CreateUserAuthContext(testAuthContext); err != nil {
		log.Printf("Test auth context already exists or error creating: %v", err)
	} else {
		log.Printf("Created test auth context: %s=%s", testAuthContext.Key, testAuthContext.Value)
	}

	testAuthTypeValidation := models.NewAuthenticationTypeValidation("GetBanks", []string{"JWT", "DirectLogin", "OAuth2"})
	if err := authRepo.CreateAuthTypeValidation(testAuthTypeValidation); err != nil {
		log.Printf("Test auth type validation already exists or error creating: %v", err)
	} else {
		log.Printf("Created test auth type validation for operation: %s", testAuthTypeValidation.OperationID)
	}

	testConsumerRateLimit := models.NewConsumerRateLimit(testConsumer.ConsumerID, 200, 2000, 20000)
	if err := db.Create(testConsumerRateLimit).Error; err != nil {
		log.Printf("Test consumer rate limit already exists or error creating: %v", err)
	} else {
		log.Printf("Created test consumer rate limit: %d req/min, %d req/hour", testConsumerRateLimit.RequestsPerMinute, testConsumerRateLimit.RequestsPerHour)
	}

	directLoginTokenConfig := models.NewTokenConfiguration("DirectLogin", 2419200, false, 0)
	if err := db.Create(directLoginTokenConfig).Error; err != nil {
		log.Printf("DirectLogin token config already exists or error creating: %v", err)
	}

	oauthTokenConfig := models.NewTokenConfiguration("OAuth", 3600, true, 86400)
	if err := db.Create(oauthTokenConfig).Error; err != nil {
		log.Printf("OAuth token config already exists or error creating: %v", err)
	}

	bcryptCostSetting := models.NewSecuritySettings("bcrypt.cost", "12", "int", "Bcrypt hashing cost for password security")
	if err := db.Create(bcryptCostSetting).Error; err != nil {
		log.Printf("Bcrypt cost setting already exists or error creating: %v", err)
	}

	log.Println("Authentication data seeding completed!")
	log.Println("Test credentials:")
	log.Println("  Username: testuser")
	log.Println("  Password: password123")
	log.Println("  Consumer Key: test_consumer_key_123")
	log.Println("  Consumer Secret: test_consumer_secret_456")
	log.Println("Configuration features enabled:")
	log.Println("  - Database-backed authentication configuration")
	log.Println("  - Consumer-specific rate limiting")
	log.Println("  - Configurable token expiration times")
	log.Println("  - Configurable security settings (bcrypt cost, lock duration)")
	log.Println("  - Environment variable overrides with database fallbacks")
	log.Println("Authentication features enabled:")
	log.Println("  - Entitlement checking (role-based access control)")
	log.Println("  - Scope-based consumer permissions")
	log.Println("  - View-based permission system")
	log.Println("  - User authentication context tracking")
	log.Println("  - User lock system")
	log.Println("  - Authentication type validation")

	return nil
}
