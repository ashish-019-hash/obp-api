package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type DirectLoginClaims struct {
	UserID     string `json:"user_id"`
	ConsumerID string `json:"consumer_id"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
	jwt.RegisteredClaims
}

func GenerateDirectLoginToken(userID, consumerID, provider, providerID, secret string, expirationSeconds int) (string, error) {
	claims := DirectLoginClaims{
		UserID:     userID,
		ConsumerID: consumerID,
		Provider:   provider,
		ProviderID: providerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationSeconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateDirectLoginToken(tokenString, secret string) (*DirectLoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DirectLoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*DirectLoginClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
