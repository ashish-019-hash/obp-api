package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	alphanumericChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numericChars      = "0123456789"
)

func GenerateSecureRandomString(length int, charset string) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result), nil
}

func GenerateConsumerKey() (string, error) {
	return GenerateSecureRandomString(32, alphanumericChars)
}

func GenerateConsumerSecret() (string, error) {
	return GenerateSecureRandomString(64, alphanumericChars)
}
