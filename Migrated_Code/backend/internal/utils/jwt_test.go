package utils

import (
	"testing"
	"time"
)

func TestGenerateAndValidateDirectLoginToken(t *testing.T) {
	secret := "test-secret-key-for-testing"
	userID := "user-123"
	consumerID := "consumer-456"
	provider := "local"
	providerID := "testuser"
	expirationSeconds := 3600

	token, err := GenerateDirectLoginToken(userID, consumerID, provider, providerID, secret, expirationSeconds)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Generated token is empty")
	}

	claims, err := ValidateDirectLoginToken(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %s, got %s", userID, claims.UserID)
	}

	if claims.ConsumerID != consumerID {
		t.Errorf("Expected ConsumerID %s, got %s", consumerID, claims.ConsumerID)
	}

	if claims.Provider != provider {
		t.Errorf("Expected Provider %s, got %s", provider, claims.Provider)
	}

	if claims.ProviderID != providerID {
		t.Errorf("Expected ProviderID %s, got %s", providerID, claims.ProviderID)
	}
}

func TestValidateDirectLoginTokenWithWrongSecret(t *testing.T) {
	secret := "test-secret-key"
	wrongSecret := "wrong-secret-key"

	token, err := GenerateDirectLoginToken("user-123", "consumer-456", "local", "testuser", secret, 3600)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	_, err = ValidateDirectLoginToken(token, wrongSecret)
	if err == nil {
		t.Fatal("Expected validation to fail with wrong secret, but it succeeded")
	}
}

func TestValidateDirectLoginTokenExpired(t *testing.T) {
	secret := "test-secret-key"

	token, err := GenerateDirectLoginToken("user-123", "consumer-456", "local", "testuser", secret, -1)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	time.Sleep(2 * time.Second)

	_, err = ValidateDirectLoginToken(token, secret)
	if err == nil {
		t.Fatal("Expected validation to fail for expired token, but it succeeded")
	}
}
