package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type SecureRandomService struct{}

func NewSecureRandomService() *SecureRandomService {
	return &SecureRandomService{}
}

func (srs *SecureRandomService) Alphanumeric(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return srs.generateRandomString(charset, length)
}

func (srs *SecureRandomService) Numeric(maxValue int) (int, error) {
	if maxValue <= 0 {
		maxValue = 999999 // Default 6-digit challenge
	}
	
	n, err := rand.Int(rand.Reader, big.NewInt(int64(maxValue)))
	if err != nil {
		return 0, err
	}
	
	return int(n.Int64()), nil
}

func (srs *SecureRandomService) NumericString(length int) (string, error) {
	const charset = "0123456789"
	return srs.generateRandomString(charset, length)
}

func (srs *SecureRandomService) Token(length int) (string, error) {
	if length <= 0 {
		length = 24 // Default token length
	}
	return srs.Alphanumeric(length)
}

func (srs *SecureRandomService) ChallengeCode(digits int) (string, error) {
	if digits <= 0 {
		digits = 6 // Default challenge code length
	}
	return srs.NumericString(digits)
}

func (srs *SecureRandomService) SessionID() (string, error) {
	return srs.Alphanumeric(32)
}

func (srs *SecureRandomService) ConsumerKey() (string, error) {
	return srs.Alphanumeric(32)
}

func (srs *SecureRandomService) ConsumerSecret() (string, error) {
	return srs.Alphanumeric(64)
}

func (srs *SecureRandomService) Nonce() (string, error) {
	return srs.Alphanumeric(16)
}

func (srs *SecureRandomService) Salt() (string, error) {
	return srs.Alphanumeric(32)
}

func (srs *SecureRandomService) generateRandomString(charset string, length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive")
	}
	
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))
	
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		result[i] = charset[randomIndex.Int64()]
	}
	
	return string(result), nil
}

func (srs *SecureRandomService) Bytes(length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be positive")
	}
	
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	
	return bytes, nil
}

func (srs *SecureRandomService) HexString(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive")
	}
	
	byteLength := (length + 1) / 2
	bytes, err := srs.Bytes(byteLength)
	if err != nil {
		return "", err
	}
	
	hexString := fmt.Sprintf("%x", bytes)
	
	if len(hexString) > length {
		hexString = hexString[:length]
	}
	
	return strings.ToUpper(hexString), nil
}

func (srs *SecureRandomService) UUID() (string, error) {
	bytes, err := srs.Bytes(16)
	if err != nil {
		return "", err
	}
	
	bytes[6] = (bytes[6] & 0x0f) | 0x40 // Version 4
	bytes[8] = (bytes[8] & 0x3f) | 0x80 // Variant 10
	
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:16]), nil
}

func (srs *SecureRandomService) ValidateRandomness(value string, minLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("value too short: minimum length %d, got %d", minLength, len(value))
	}
	
	if srs.hasObviousPattern(value) {
		return fmt.Errorf("value appears to have obvious patterns")
	}
	
	return nil
}

func (srs *SecureRandomService) hasObviousPattern(value string) bool {
	if len(value) < 3 {
		return false
	}
	
	firstChar := value[0]
	allSame := true
	for _, char := range value {
		if char != rune(firstChar) {
			allSame = false
			break
		}
	}
	
	return allSame
}
