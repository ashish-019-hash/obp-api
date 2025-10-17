package utils

import (
	"testing"
)

func TestGenerateConsumerKey(t *testing.T) {
	key, err := GenerateConsumerKey()
	if err != nil {
		t.Fatalf("Failed to generate consumer key: %v", err)
	}

	if len(key) != 32 {
		t.Errorf("Expected consumer key length 32, got %d", len(key))
	}

	key2, err := GenerateConsumerKey()
	if err != nil {
		t.Fatalf("Failed to generate second consumer key: %v", err)
	}

	if key == key2 {
		t.Fatal("Expected two consecutive keys to be different")
	}
}

func TestGenerateConsumerSecret(t *testing.T) {
	secret, err := GenerateConsumerSecret()
	if err != nil {
		t.Fatalf("Failed to generate consumer secret: %v", err)
	}

	if len(secret) != 64 {
		t.Errorf("Expected consumer secret length 64, got %d", len(secret))
	}

	secret2, err := GenerateConsumerSecret()
	if err != nil {
		t.Fatalf("Failed to generate second consumer secret: %v", err)
	}

	if secret == secret2 {
		t.Fatal("Expected two consecutive secrets to be different")
	}
}

func TestGenerateSecureRandomString(t *testing.T) {
	length := 16
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"

	str, err := GenerateSecureRandomString(length, charset)
	if err != nil {
		t.Fatalf("Failed to generate secure random string: %v", err)
	}

	if len(str) != length {
		t.Errorf("Expected length %d, got %d", length, len(str))
	}

	for _, char := range str {
		found := false
		for _, c := range charset {
			if char == c {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Character %c is not in charset", char)
		}
	}
}
