package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword123!@#"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hash == "" {
		t.Fatal("Generated hash is empty")
	}

	if hash == password {
		t.Fatal("Hash should not be equal to the original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword123!@#"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatal("Expected password to match hash, but it didn't")
	}

	if CheckPasswordHash("wrongPassword", hash) {
		t.Fatal("Expected wrong password to not match hash, but it did")
	}
}

func TestCheckPasswordHashEmptyPassword(t *testing.T) {
	password := ""
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash empty password: %v", err)
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatal("Expected empty password to match its hash")
	}
}
