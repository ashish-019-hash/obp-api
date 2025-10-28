#!/bin/bash


BASE_URL="http://localhost:8080/obp/v5.1.0"
RESULTS_FILE="comprehensive_test_results.txt"

echo "=== OBP API v5.1.0 Comprehensive Testing ===" > $RESULTS_FILE
echo "Started: $(date)" >> $RESULTS_FILE
echo "" >> $RESULTS_FILE

test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local expected_status=$4
    local description=$5
    
    echo "Testing: $method $endpoint - $description"
    echo "Testing: $method $endpoint - $description" >> $RESULTS_FILE
    
    if [ -n "$data" ]; then
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X $method "$BASE_URL$endpoint" \
                  -H "Content-Type: application/json" \
                  -d "$data")
    else
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X $method "$BASE_URL$endpoint")
    fi
    
    http_code=$(echo $response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    body=$(echo $response | sed -e 's/HTTPSTATUS:.*//g')
    
    if [ "$http_code" = "$expected_status" ]; then
        echo "  ✓ PASS - Status: $http_code" >> $RESULTS_FILE
        echo "  Response: $body" >> $RESULTS_FILE
    else
        echo "  ✗ FAIL - Expected: $expected_status, Got: $http_code" >> $RESULTS_FILE
        echo "  Response: $body" >> $RESULTS_FILE
    fi
    echo "" >> $RESULTS_FILE
}

echo "=== 1. Core API Information Endpoints ==="
test_endpoint "GET" "/root" "" "200" "API root information"
test_endpoint "GET" "/ui/suggested-session-timeout" "" "200" "Session timeout"
test_endpoint "GET" "/well-known" "" "200" "OAuth2 well-known URIs"
test_endpoint "GET" "/tags" "" "200" "API tags"

echo "=== 2. Database-Integrated Endpoints (Bank Controller) ==="
test_endpoint "POST" "/banks" '{"id":"test_comprehensive_bank","short_name":"Test Comprehensive Bank","full_name":"Test Comprehensive Bank Ltd","logo":"","website":""}' "201" "Create bank"
test_endpoint "GET" "/banks" "" "200" "Get all banks"
test_endpoint "GET" "/banks/test_comprehensive_bank" "" "200" "Get specific bank"

echo "=== 3. Database-Integrated Endpoints (User Controller) ==="
test_endpoint "POST" "/users" '{"username":"testcomprehensive","email":"comprehensive@example.com","password":"password123","first_name":"Test","last_name":"Comprehensive"}' "201" "Create user"
test_endpoint "GET" "/users" "" "200" "Get all users"
test_endpoint "GET" "/users/provider/github/username/testuser" "" "200" "Get user by provider"
test_endpoint "GET" "/users/provider/github/username/testuser/lock-status" "" "200" "Get user lock status"
test_endpoint "POST" "/users/provider/github/provider-id/test123/sync" "" "201" "Sync external user"

echo "=== 4. Consent Management Endpoints ==="
test_endpoint "GET" "/my/consents" "" "200" "Get user consents"
test_endpoint "GET" "/consents" "" "200" "Get all consents"
test_endpoint "POST" "/banks/test_comprehensive_bank/consents" '{"everything":true,"views":[],"entitlements":[]}' "201" "Create consent"
test_endpoint "PUT" "/consents/consent123/status" '{"status":"ACCEPTED"}' "200" "Update consent status"

echo "=== 5. Balance Management Endpoints ==="
test_endpoint "GET" "/banks/test_comprehensive_bank/accounts/test_account/views/owner/balances" "" "200" "Get account balances"
test_endpoint "GET" "/banks/test_comprehensive_bank/balances" "" "200" "Get bank balances"
test_endpoint "POST" "/banks/test_comprehensive_bank/accounts/test_account/balances" '{"balance_type":"CURRENT","balance_amount":"1000.00"}' "201" "Create balance"

echo "=== 6. Counterparty Management Endpoints ==="
test_endpoint "GET" "/banks/test_comprehensive_bank/accounts/test_account/views/owner/counterparties" "" "200" "Get counterparties"
test_endpoint "POST" "/banks/test_comprehensive_bank/accounts/test_account/views/owner/counterparties" '{"name":"Test Counterparty","bank_id":"test_comprehensive_bank","account_id":"test_account"}' "201" "Create counterparty"
test_endpoint "POST" "/banks/test_comprehensive_bank/accounts/test_account/views/owner/counterparties/cp123/limits" '{"currency_code":"USD","max_single_amount":"1000","max_monthly_amount":"5000","max_yearly_amount":"10000","max_number_of_monthly_transactions":50,"max_number_of_yearly_transactions":500}' "201" "Create counterparty limit"

echo "=== 7. ATM Management Endpoints ==="
test_endpoint "GET" "/banks/test_comprehensive_bank/atms" "" "200" "Get ATMs"
test_endpoint "POST" "/banks/test_comprehensive_bank/atms" '{"bank_id":"test_comprehensive_bank","name":"Test ATM","address":{"city":"Test City","country_code":"US"},"location":{"latitude":40.7128,"longitude":-74.0060},"meta":{"license":{"id":"license_001","name":"Standard License"}}}' "201" "Create ATM"

echo "=== 8. Consumer Management Endpoints ==="
test_endpoint "GET" "/management/consumers" "" "200" "Get consumers"
test_endpoint "POST" "/management/consumers" '{"app_name":"Test App","app_type":"Confidential","description":"Test application","developer_email":"dev@example.com"}' "201" "Create consumer"

echo "=== 9. Account Access Endpoints ==="
test_endpoint "POST" "/banks/test_comprehensive_bank/accounts/test_account/views/owner/account-access/grant" '{"user_id":"user123","view_id":"owner"}' "201" "Grant account access"
test_endpoint "GET" "/users/user123/account-access" "" "200" "Get user account access"

echo "=== 10. Transaction Request Endpoints ==="
test_endpoint "GET" "/banks/test_comprehensive_bank/accounts/test_account/owner/transaction-requests" "" "200" "Get transaction requests"
test_endpoint "GET" "/management/transaction-requests/tr123" "" "200" "Get transaction request by ID"

echo "=== 11. User Attribute Endpoints ==="
test_endpoint "POST" "/users/user123/non-personal/attributes" '{"name":"test_attribute","type":"STRING","value":"test_value"}' "201" "Create user attribute"

echo "=== 12. Regulated Entity Endpoints ==="
test_endpoint "POST" "/regulated-entities" '{"name":"Test Entity","website":"https://test.com","email":"contact@test.com","entity_type":"BANK"}' "201" "Create regulated entity"

echo "=== 13. WebUI Endpoints ==="
test_endpoint "GET" "/webui-props" "" "200" "Get WebUI properties"

echo "=== 14. System View Endpoints ==="
test_endpoint "POST" "/system-views/system/permissions" '{"permission_name":"can_see_transaction_amount"}' "201" "Add system view permission"

echo "Completed: $(date)" >> $RESULTS_FILE
echo "=== Testing Complete - Check $RESULTS_FILE for detailed results ==="
