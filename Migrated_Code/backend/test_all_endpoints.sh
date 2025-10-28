#!/bin/bash


echo "=== OBP-API Authentication Testing ==="
echo "Testing all API endpoints for proper authentication protection"
echo "Server should be running on localhost:8080"
echo ""

BASE_URL="http://localhost:8080"
TOKEN=""

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

test_endpoint() {
    local method=$1
    local endpoint=$2
    local expected_status=$3
    local auth_header=$4
    local description=$5
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    echo -n "Testing: $method $endpoint - $description ... "
    
    if [ -n "$auth_header" ]; then
        response=$(curl -s -w "%{http_code}" -X "$method" "$BASE_URL$endpoint" -H "$auth_header" -o /dev/null)
    else
        response=$(curl -s -w "%{http_code}" -X "$method" "$BASE_URL$endpoint" -o /dev/null)
    fi
    
    if [ "$response" = "$expected_status" ]; then
        echo -e "${GREEN}PASS${NC} (Status: $response)"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    else
        echo -e "${RED}FAIL${NC} (Expected: $expected_status, Got: $response)"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
}

get_directlogin_token() {
    echo "=== Getting DirectLogin Token ==="
    response=$(curl -s -X POST "$BASE_URL/auth/direct-login" \
        -H "Content-Type: application/json" \
        -d '{
            "username": "testuser",
            "password": "password123",
            "consumer_key": "test_consumer_key_123"
        }')
    
    TOKEN=$(echo "$response" | grep -o '"token":"[^"]*' | cut -d'"' -f4)
    
    if [ -n "$TOKEN" ]; then
        echo -e "${GREEN}✓ DirectLogin token obtained${NC}"
        echo "Token: ${TOKEN:0:20}..."
    else
        echo -e "${RED}✗ Failed to get DirectLogin token${NC}"
        echo "Response: $response"
        exit 1
    fi
    echo ""
}

echo "=== 1. Testing Public Endpoints (Should work without auth) ==="
test_endpoint "GET" "/health" "200" "" "Health check"
test_endpoint "GET" "/ping" "200" "" "Ping endpoint"
test_endpoint "GET" "/obp/v5.1.0/root" "200" "" "API root info"
test_endpoint "GET" "/obp/v5.1.0/well-known" "200" "" "OAuth2 well-known URIs"
test_endpoint "GET" "/obp/v5.1.0/ui/suggested-session-timeout" "200" "" "Session timeout"
test_endpoint "GET" "/obp/v5.1.0/waiting-for-godot" "200" "" "Waiting for Godot"
test_endpoint "GET" "/api/v1/health" "200" "" "API v1 health"
echo ""

echo "=== 2. Testing Authentication Endpoints ==="
test_endpoint "POST" "/auth/direct-login" "200" "" "DirectLogin authentication (with valid creds)"
test_endpoint "POST" "/auth/consumers" "201" "" "Consumer registration"
test_endpoint "POST" "/auth/users" "201" "" "User registration"
test_endpoint "POST" "/oauth/initiate" "200" "" "OAuth initiate"
test_endpoint "GET" "/oauth/authorize" "400" "" "OAuth authorize (missing token)"
test_endpoint "POST" "/oauth/token" "400" "" "OAuth token (missing params)"
echo ""

get_directlogin_token

echo "=== 3. Testing Protected Endpoints Without Auth (Should return 401) ==="
test_endpoint "GET" "/obp/v5.1.0/banks" "401" "" "Get banks (no auth)"
test_endpoint "POST" "/obp/v5.1.0/banks" "401" "" "Create bank (no auth)"
test_endpoint "GET" "/obp/v5.1.0/users" "401" "" "Get users (no auth)"
test_endpoint "POST" "/obp/v5.1.0/users" "401" "" "Create user (no auth)"
test_endpoint "GET" "/obp/v5.1.0/tags" "401" "" "Get API tags (no auth)"
test_endpoint "GET" "/my/user" "401" "" "Get current user (no auth)"
echo ""

echo "=== 4. Testing Protected Endpoints With Valid Auth (Should work) ==="
AUTH_HEADER="Authorization: DirectLogin token=$TOKEN"
test_endpoint "GET" "/obp/v5.1.0/banks" "200" "$AUTH_HEADER" "Get banks (with auth)"
test_endpoint "GET" "/obp/v5.1.0/tags" "200" "$AUTH_HEADER" "Get API tags (with auth)"
test_endpoint "GET" "/my/user" "200" "$AUTH_HEADER" "Get current user (with auth)"
echo ""

echo "=== 5. Testing Management Endpoints (Require special entitlements) ==="
test_endpoint "GET" "/obp/v5.1.0/management/api-collections" "403" "$AUTH_HEADER" "Get API collections (insufficient entitlement)"
test_endpoint "GET" "/obp/v5.1.0/management/metrics" "403" "$AUTH_HEADER" "Get metrics (insufficient entitlement)"
test_endpoint "GET" "/obp/v5.1.0/management/consumers" "403" "$AUTH_HEADER" "Get consumers (insufficient entitlement)"
echo ""

echo "=== 6. Testing My Endpoints (User-specific, require auth) ==="
test_endpoint "GET" "/obp/v5.1.0/my/api-collections" "200" "$AUTH_HEADER" "Get my API collections"
test_endpoint "GET" "/obp/v5.1.0/my/consents" "200" "$AUTH_HEADER" "Get my consents"
echo ""

echo "=== 7. Testing Consumer Endpoints (OAuth-specific auth) ==="
test_endpoint "GET" "/obp/v5.1.0/consumer/consents/test-consent-id" "401" "" "Get consent via consumer (no OAuth)"
echo ""

echo "=== 8. Testing Bank-specific Endpoints ==="
test_endpoint "GET" "/obp/v5.1.0/banks/test-bank/accounts/test-account/views/owner" "401" "" "Get account view (no auth)"
test_endpoint "GET" "/obp/v5.1.0/banks/test-bank/accounts/test-account/views/owner" "404" "$AUTH_HEADER" "Get account view (with auth, account not found)"
test_endpoint "GET" "/obp/v5.1.0/banks/test-bank/currencies" "404" "$AUTH_HEADER" "Get bank currencies (with auth, bank not found)"
echo ""

echo "=== 9. Testing Rate Limiting ==="
echo "Testing rate limiting (100 requests per minute)..."
rate_limit_hit=false
for i in {1..105}; do
    response=$(curl -s -w "%{http_code}" "$BASE_URL/obp/v5.1.0/root" -o /dev/null)
    if [ "$response" = "429" ]; then
        echo -e "${GREEN}✓ Rate limiting triggered at request $i${NC}"
        rate_limit_hit=true
        break
    fi
done

if [ "$rate_limit_hit" = false ]; then
    echo -e "${YELLOW}⚠ Rate limiting not triggered within 105 requests${NC}"
fi
echo ""

echo "=== 10. Testing User Lockout (Failed Login Attempts) ==="
echo "Testing user lockout after 5 failed attempts..."
lockout_triggered=false
for i in {1..6}; do
    response=$(curl -s -X POST "$BASE_URL/auth/direct-login" \
        -H "Content-Type: application/json" \
        -d '{
            "username": "testuser",
            "password": "wrongpassword",
            "consumer_key": "test_consumer_key_123"
        }')
    
    if echo "$response" | grep -q "locked\|lockout"; then
        echo -e "${GREEN}✓ User lockout triggered at attempt $i${NC}"
        lockout_triggered=true
        break
    fi
done

if [ "$lockout_triggered" = false ]; then
    echo -e "${YELLOW}⚠ User lockout not triggered within 6 failed attempts${NC}"
fi
echo ""

echo "=== Test Summary ==="
echo "Total Tests: $TOTAL_TESTS"
echo -e "Passed: ${GREEN}$PASSED_TESTS${NC}"
echo -e "Failed: ${RED}$FAILED_TESTS${NC}"
echo ""

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}🎉 All tests passed! Authentication system is working correctly.${NC}"
    exit 0
else
    echo -e "${RED}❌ Some tests failed. Please review the authentication implementation.${NC}"
    exit 1
fi
