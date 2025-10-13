package services

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWKSService struct {
	configService *ConfigService
	httpClient    *http.Client
	jwksCache     map[string]*JWKSResponse
	cacheExpiry   map[string]time.Time
}

type JWKSResponse struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kty string `json:"kty"` // Key Type
	Use string `json:"use"` // Public Key Use
	Kid string `json:"kid"` // Key ID
	N   string `json:"n"`   // RSA Modulus
	E   string `json:"e"`   // RSA Exponent
	Alg string `json:"alg"` // Algorithm
}

func NewJWKSService(configService *ConfigService) *JWKSService {
	return &JWKSService{
		configService: configService,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		jwksCache:   make(map[string]*JWKSResponse),
		cacheExpiry: make(map[string]time.Time),
	}
}

func (j *JWKSService) ValidateJWTWithJWKS(tokenString, jwksURL string) (*jwt.Token, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("token missing kid header")
	}

	jwks, err := j.getJWKS(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get JWKS: %w", err)
	}

	var jwk *JWK
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			jwk = &key
			break
		}
	}

	if jwk == nil {
		return nil, fmt.Errorf("key with kid %s not found in JWKS", kid)
	}

	publicKey, err := j.jwkToRSAPublicKey(jwk)
	if err != nil {
		return nil, fmt.Errorf("failed to convert JWK to RSA key: %w", err)
	}

	validatedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	return validatedToken, nil
}

func (j *JWKSService) getJWKS(jwksURL string) (*JWKSResponse, error) {
	if cached, exists := j.jwksCache[jwksURL]; exists {
		if time.Now().Before(j.cacheExpiry[jwksURL]) {
			return cached, nil
		}
		delete(j.jwksCache, jwksURL)
		delete(j.cacheExpiry, jwksURL)
	}

	resp, err := j.httpClient.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("JWKS endpoint returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKS response: %w", err)
	}

	var jwks JWKSResponse
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, fmt.Errorf("failed to parse JWKS response: %w", err)
	}

	j.jwksCache[jwksURL] = &jwks
	j.cacheExpiry[jwksURL] = time.Now().Add(time.Hour)

	return &jwks, nil
}

func (j *JWKSService) jwkToRSAPublicKey(jwk *JWK) (*rsa.PublicKey, error) {
	if jwk.Kty != "RSA" {
		return nil, fmt.Errorf("unsupported key type: %s", jwk.Kty)
	}

	
	return nil, errors.New("JWK to RSA conversion not fully implemented - requires base64url decoding and big.Int conversion")
}

func (j *JWKSService) ValidateOIDCToken(tokenString string) (*jwt.Token, map[string]interface{}, error) {
	oidcEnabled := j.configService.GetConfigBool("oauth2.enabled", false)
	if !oidcEnabled {
		return nil, nil, errors.New("OIDC authentication is disabled")
	}

	jwksURL := j.configService.GetConfig("oauth2.jwk_set.url", "")
	if jwksURL == "" {
		return nil, nil, errors.New("JWKS URL not configured")
	}

	token, err := j.ValidateJWTWithJWKS(tokenString, jwksURL)
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("invalid token claims")
	}

	if err := j.validateOIDCClaims(claims); err != nil {
		return nil, nil, fmt.Errorf("OIDC claim validation failed: %w", err)
	}

	return token, claims, nil
}

func (j *JWKSService) validateOIDCClaims(claims jwt.MapClaims) error {
	requiredClaims := []string{"iss", "sub", "aud", "exp", "iat"}
	for _, claim := range requiredClaims {
		if _, exists := claims[claim]; !exists {
			return fmt.Errorf("missing required claim: %s", claim)
		}
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return errors.New("token expired")
		}
	}

	if iat, ok := claims["iat"].(float64); ok {
		if time.Now().Unix() < int64(iat) {
			return errors.New("token used before issued")
		}
	}

	return nil
}
