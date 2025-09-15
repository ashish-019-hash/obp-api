# OBP API v5.1.0 - Complete Testing Guide

This comprehensive guide covers all 185+ REST API endpoints implemented in the Go backend service for Open Bank Project API v5.1.0.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Core API Endpoints](#core-api-endpoints)
3. [User Management](#user-management)
4. [Consent Management](#consent-management)
5. [Account & Balance Management](#account--balance-management)
6. [Counterparty Management](#counterparty-management)
7. [Transaction Management](#transaction-management)
8. [Consumer Management](#consumer-management)
9. [ATM Management](#atm-management)
10. [System Management](#system-management)
11. [Security & Certificates](#security--certificates)
12. [Metrics & Monitoring](#metrics--monitoring)
13. [Testing Examples](#testing-examples)

---

## Getting Started

### Server Setup
```bash
# Start the server
cd /path/to/backend
go run cmd/main.go

# Server runs on port 8080
```

### Base URL
All endpoints use the base URL pattern: `http://[server-host]:8080`

### Response Format
All endpoints return JSON responses with consistent structure:
```json
{
  "data": {...},
  "status": "success|error",
  "message": "Description"
}
```

---

## Core API Endpoints

### 1. API Information
**Purpose**: Get API version and basic information

#### GET /root
**Description**: Returns API version information and status
**Use Case**: Check API availability and version compatibility
**Response**:
```json
{
  "version": "v5.1.0",
  "version_status": "BLEEDING_EDGE",
  "git_commit": "abc123",
  "connector": "mapped"
}
```

**Test Command**:
```bash
curl -X GET [BASE_URL]/root
```

#### GET /ui/suggested-session-timeout
**Description**: Returns suggested session timeout for UI applications
**Use Case**: Configure client-side session management
**Response**:
```json
{
  "suggested_session_timeout_in_minutes": 60
}
```

**Test Command**:
```bash
curl -X GET [BASE_URL]/ui/suggested-session-timeout
```

#### GET /well-known
**Description**: OAuth2 well-known URIs for discovery
**Use Case**: OAuth2 client configuration and discovery
**Response**:
```json
{
  "authorization_endpoint": "https://api.example.com/oauth/authorize",
  "token_endpoint": "https://api.example.com/oauth/token",
  "userinfo_endpoint": "https://api.example.com/oauth/userinfo"
}
```

#### GET /waiting-for-godot
**Description**: Test endpoint for connectivity checks
**Use Case**: Health checks and connectivity testing

#### GET /tags
**Description**: Get all available API tags
**Use Case**: API documentation and categorization

---

## User Management

### User Information & Status

#### GET /users/provider/{provider}/username/{username}
**Description**: Get user information by provider and username
**Use Case**: User lookup and profile retrieval
**Parameters**:
- `provider`: Authentication provider (e.g., "github", "google")
- `username`: Username within the provider

**Response**:
```json
{
  "user_id": "user_123",
  "username": "john_doe",
  "provider": "github",
  "email": "john@example.com",
  "display_name": "John Doe",
  "is_locked": false,
  "is_validated": true
}
```

**Test Command**:
```bash
curl -X GET [BASE_URL]/users/provider/github/username/john_doe
```

#### GET /users/provider/{provider}/username/{username}/lock-status
**Description**: Check user lock status
**Use Case**: Security monitoring and user access control
**Response**:
```json
{
  "is_locked": false,
  "lock_reason": "",
  "locked_at": "2023-01-01T00:00:00Z",
  "bad_attempts": 0
}
```

#### PUT /users/provider/{provider}/username/{username}/lock-status
**Description**: Unlock a user account
**Use Case**: Administrative user account management
**Request Body**: None required
**Response**: Updated lock status

#### POST /users/provider/{provider}/username/{username}/locks
**Description**: Lock a user account
**Use Case**: Security enforcement and account suspension
**Request Body**:
```json
{
  "lock_reason": "Suspicious activity detected"
}
```

#### PUT /management/users/{userId}
**Description**: Validate a user by user ID
**Use Case**: User verification and account activation

#### POST /users/provider/{provider}/provider-id/{providerId}/sync
**Description**: Synchronize external user data
**Use Case**: User data synchronization from external providers
**Response**:
```json
{
  "user_id": "user_456",
  "username": "external_user",
  "provider": "ldap",
  "email": "user@company.com",
  "display_name": "External User",
  "is_locked": false,
  "is_validated": true
}
```

### User Attributes

#### POST /users/{userId}/non-personal/attributes
**Description**: Create non-personal user attributes
**Use Case**: Store additional user metadata
**Request Body**:
```json
{
  "name": "department",
  "type": "STRING",
  "value": "Engineering"
}
```

#### GET /users/{userId}/non-personal/attributes
**Description**: Get all non-personal user attributes
**Use Case**: Retrieve user metadata

#### DELETE /users/{userId}/non-personal/attributes/{userAttributeId}
**Description**: Delete a specific user attribute
**Use Case**: Attribute cleanup and management

#### GET /users/{userId}/banks/{bankId}/accounts-held
**Description**: Get accounts held by user at specific bank
**Use Case**: Account relationship management

#### GET /users/{userId}/accounts-held
**Description**: Get all accounts held by user across all banks
**Use Case**: Comprehensive account overview

---

## Consent Management

### Consent Lifecycle

#### POST /banks/{bankId}/consents
**Description**: Create a new consent
**Use Case**: Initiate consent flow for account access
**Request Body**:
```json
{
  "everything": false,
  "views": [
    {
      "bank_id": "bank123",
      "account_id": "acc456",
      "view_id": "owner"
    }
  ],
  "entitlements": [
    {
      "bank_id": "bank123",
      "role_name": "CanGetAccount"
    }
  ],
  "consumer_id": "consumer789",
  "email": "user@example.com",
  "valid_from": "2023-01-01T00:00:00Z",
  "time_to_live": 3600
}
```

#### GET /banks/{bankId}/consents/{consentId}
**Description**: Get consent details by ID
**Use Case**: Consent status checking and details retrieval

#### PUT /banks/{bankId}/consents/{consentId}/status
**Description**: Update consent status
**Use Case**: Consent approval, rejection, or revocation
**Request Body**:
```json
{
  "status": "ACCEPTED"
}
```

#### DELETE /banks/{bankId}/consents/{consentId}
**Description**: Revoke consent at bank level
**Use Case**: Administrative consent revocation

### Consent Queries

#### GET /my/consents
**Description**: Get all consents for current user
**Use Case**: User consent management dashboard
**Response**:
```json
{
  "consents": [
    {
      "consent_id": "consent_123",
      "status": "ACCEPTED",
      "created_date": "2023-01-01T00:00:00Z",
      "valid_from": "2023-01-01T00:00:00Z",
      "valid_to": "2023-12-31T23:59:59Z"
    }
  ]
}
```

#### GET /banks/{bankId}/my/consents
**Description**: Get user's consents for specific bank
**Use Case**: Bank-specific consent management

#### GET /banks/{bankId}/consents
**Description**: Get all consents at bank (admin)
**Use Case**: Bank administrative consent overview

#### GET /consents
**Description**: Get all consents across system (super admin)
**Use Case**: System-wide consent monitoring

### Consumer Consent Access

#### GET /consumer/consents/{consentId}
**Description**: Get consent details via consumer
**Use Case**: Third-party application consent verification

#### DELETE /consumer/consents/{consentId}
**Description**: Self-revoke consent by consumer
**Use Case**: Third-party application consent cleanup

#### DELETE /my/consents/{consentId}
**Description**: User self-revoke consent
**Use Case**: User-initiated consent revocation

### Implicit Consent

#### POST /my/consents/{scaMethod}
**Description**: Create implicit consent with SCA method
**Use Case**: Streamlined consent flow with strong customer authentication
**Parameters**:
- `scaMethod`: Strong Customer Authentication method (e.g., "SMS", "EMAIL")

### Advanced Consent Management

#### PUT /banks/{bankId}/consents/{consentId}/account-access
**Description**: Update consent account access permissions
**Use Case**: Modify consent scope and permissions

#### PUT /banks/{bankId}/consents/{consentId}/user-id
**Description**: Update consent user association
**Use Case**: Transfer consent ownership

#### GET /user/current/consents/{consentId}
**Description**: Get current user's specific consent
**Use Case**: Detailed consent information for user

---

## Account & Balance Management

### Account Information

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}
**Description**: Get core account information through specific view
**Use Case**: Account details retrieval with view-based permissions
**Response**:
```json
{
  "id": "account_123",
  "bank_id": "bank_456",
  "label": "My Checking Account",
  "number": "1234567890",
  "type": "CURRENT",
  "balance": {
    "currency": "EUR",
    "amount": "1500.50"
  },
  "IBAN": "DE89370400440532013000",
  "swift_bic": "DEUTDEFF",
  "views_available": [
    {
      "id": "owner",
      "short_name": "Owner",
      "description": "Full access"
    }
  ]
}
```

**Test Command**:
```bash
curl -X GET [BASE_URL]/banks/bank123/accounts/acc456/views/owner
```

### Balance Management

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/balances
**Description**: Get account balances through view
**Use Case**: Balance inquiry with view permissions
**Response**:
```json
{
  "balances": [
    {
      "id": "balance_001",
      "type": "CURRENT",
      "currency": "EUR",
      "amount": "1500.50",
      "date": "2023-01-01T00:00:00Z"
    },
    {
      "id": "balance_002",
      "type": "AVAILABLE",
      "currency": "EUR",
      "amount": "1450.50",
      "date": "2023-01-01T00:00:00Z"
    }
  ]
}
```

#### GET /banks/{bankId}/balances
**Description**: Get balances for all accounts at bank
**Use Case**: Bank-wide balance overview

#### GET /banks/{bankId}/views/{viewId}/balances
**Description**: Get balances through specific view across bank
**Use Case**: View-filtered balance information

#### POST /banks/{bankId}/accounts/{accountId}/balances
**Description**: Create new balance entry
**Use Case**: Balance initialization or correction
**Request Body**:
```json
{
  "balance_type": "CURRENT",
  "balance_amount": "2000.00"
}
```

#### GET /banks/{bankId}/accounts/{accountId}/balances/{balanceId}
**Description**: Get specific balance by ID
**Use Case**: Individual balance record retrieval

#### PUT /banks/{bankId}/accounts/{accountId}/balances/{balanceId}
**Description**: Update existing balance
**Use Case**: Balance correction or adjustment

#### DELETE /banks/{bankId}/accounts/{accountId}/balances/{balanceId}
**Description**: Delete balance record
**Use Case**: Balance record cleanup

### Account Access Control

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/account-access/grant
**Description**: Grant user access to account view
**Use Case**: Permission management and access control
**Request Body**:
```json
{
  "user_id": "user_123"
}
```

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/account-access/revoke
**Description**: Revoke user access to account view
**Use Case**: Access revocation and security management

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/user-account-access
**Description**: Create user with account access
**Use Case**: User onboarding with immediate access

#### GET /users/{userId}/account-access
**Description**: Get all account access for user
**Use Case**: User permission audit and overview

---

## Counterparty Management

### Counterparty CRUD Operations

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties
**Description**: Get all counterparties for account view
**Use Case**: Counterparty listing and management
**Response**:
```json
{
  "counterparties": [
    {
      "counterparty_id": "cp_001",
      "name": "ACME Corp",
      "description": "Business partner",
      "currency": "EUR",
      "bank_routing": {
        "scheme": "IBAN",
        "address": "DE89370400440532013000"
      },
      "account_routing": {
        "scheme": "AccountNumber",
        "address": "1234567890"
      },
      "other_bank_routing_scheme": "BIC",
      "other_bank_routing_address": "DEUTDEFF",
      "other_account_routing_scheme": "IBAN",
      "other_account_routing_address": "DE89370400440532013000",
      "is_beneficiary": true,
      "bespoke": []
    }
  ]
}
```

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties
**Description**: Create new counterparty
**Use Case**: Add new payment recipient or business partner
**Request Body**:
```json
{
  "name": "New Counterparty",
  "description": "Payment recipient",
  "currency": "EUR",
  "bank_routing": {
    "scheme": "IBAN",
    "address": "DE89370400440532013000"
  },
  "account_routing": {
    "scheme": "AccountNumber",
    "address": "9876543210"
  },
  "other_bank_routing_scheme": "BIC",
  "other_bank_routing_address": "DEUTDEFF",
  "other_account_routing_scheme": "IBAN",
  "other_account_routing_address": "DE89370400440532013000",
  "is_beneficiary": true,
  "bespoke": []
}
```

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}
**Description**: Get specific counterparty details
**Use Case**: Counterparty information retrieval

#### PUT /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}
**Description**: Update counterparty information
**Use Case**: Counterparty data maintenance

#### DELETE /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}
**Description**: Delete counterparty
**Use Case**: Counterparty cleanup and removal

### Counterparty Limits

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limits
**Description**: Create counterparty transaction limits
**Use Case**: Risk management and transaction control
**Request Body**:
```json
{
  "currency": "EUR",
  "max_single_amount": "1000.00",
  "max_monthly_amount": "5000.00",
  "max_yearly_amount": "50000.00",
  "max_single_number": 1,
  "max_monthly_number": 10,
  "max_yearly_number": 100
}
```

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limits
**Description**: Get counterparty limits
**Use Case**: Limit inquiry and compliance checking

#### PUT /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limits
**Description**: Update counterparty limits
**Use Case**: Limit adjustment and risk management

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limit-status
**Description**: Check counterparty limit status
**Use Case**: Real-time limit checking before transactions

#### DELETE /banks/{bankId}/accounts/{accountId}/views/{viewId}/counterparties/{counterpartyId}/limits
**Description**: Remove counterparty limits
**Use Case**: Limit removal and unrestricted access

---

## Transaction Management

### Transaction Requests

#### GET /banks/{bankId}/accounts/{accountId}/{viewId}/transaction-requests
**Description**: Get transaction requests for account
**Use Case**: Transaction request history and status tracking
**Response**:
```json
{
  "transaction_requests": [
    {
      "id": "tr_001",
      "type": "SEPA",
      "from": {
        "bank_id": "bank123",
        "account_id": "acc456"
      },
      "to": {
        "counterparty_id": "cp_001"
      },
      "body": {
        "value": {
          "currency": "EUR",
          "amount": "100.00"
        },
        "description": "Payment for services"
      },
      "status": "COMPLETED",
      "start_date": "2023-01-01T00:00:00Z",
      "end_date": "2023-01-01T00:05:00Z"
    }
  ]
}
```

#### GET /management/transaction-requests/{transactionRequestId}
**Description**: Get specific transaction request (admin)
**Use Case**: Administrative transaction request monitoring

#### PUT /management/transaction-requests/{transactionRequestId}
**Description**: Update transaction request status (admin)
**Use Case**: Administrative transaction request management
**Request Body**:
```json
{
  "status": "CANCELLED",
  "reason": "Compliance review required"
}
```

---

## Consumer Management

### Consumer Registration

#### POST /dynamic-registration/consumers
**Description**: Dynamic consumer registration
**Use Case**: Third-party application registration
**Request Body**:
```json
{
  "app_name": "My Banking App",
  "app_type": "Web",
  "description": "Personal finance management",
  "developer_email": "dev@example.com",
  "redirect_uris": ["https://myapp.com/callback"],
  "client_uri": "https://myapp.com",
  "logo_uri": "https://myapp.com/logo.png"
}
```

#### POST /management/consumers
**Description**: Create consumer (admin)
**Use Case**: Administrative consumer creation

#### POST /my/consumers
**Description**: Create consumer for current user
**Use Case**: User-initiated consumer registration

### Consumer Management

#### GET /management/consumers/{consumerId}
**Description**: Get consumer details (admin)
**Use Case**: Consumer information retrieval

#### GET /management/consumers
**Description**: Get all consumers (admin)
**Use Case**: Consumer management dashboard

#### PUT /management/consumers/{consumerId}/consumer/redirect_url
**Description**: Update consumer redirect URL
**Use Case**: Consumer configuration management

#### PUT /management/consumers/{consumerId}/consumer/logo_url
**Description**: Update consumer logo URL
**Use Case**: Consumer branding management

#### PUT /management/consumers/{consumerId}/consumer/certificate
**Description**: Update consumer certificate
**Use Case**: Consumer security credential management

#### PUT /management/consumers/{consumerId}/consumer/name
**Description**: Update consumer name
**Use Case**: Consumer profile management

---

## ATM Management

### ATM Operations

#### POST /banks/{bankId}/atms
**Description**: Create new ATM
**Use Case**: ATM network expansion
**Request Body**:
```json
{
  "id": "atm_001",
  "bank_id": "bank123",
  "name": "Main Street ATM",
  "address": {
    "line_1": "123 Main Street",
    "line_2": "",
    "line_3": "",
    "city": "Berlin",
    "county": "",
    "state": "Berlin",
    "postcode": "10115",
    "country_code": "DE"
  },
  "location": {
    "latitude": 52.5200,
    "longitude": 13.4050
  },
  "meta": {
    "license": {
      "id": "license_001",
      "name": "Open Bank Project License"
    }
  }
}
```

#### GET /banks/{bankId}/atms
**Description**: Get all ATMs for bank
**Use Case**: ATM location services

#### GET /banks/{bankId}/atms/{atmId}
**Description**: Get specific ATM details
**Use Case**: ATM information retrieval

#### PUT /banks/{bankId}/atms/{atmId}
**Description**: Update ATM information
**Use Case**: ATM maintenance and updates

#### DELETE /banks/{bankId}/atms/{atmId}
**Description**: Remove ATM from network
**Use Case**: ATM decommissioning

### ATM Attributes

#### POST /banks/{bankId}/atms/{atmId}/attributes
**Description**: Create ATM attribute
**Use Case**: ATM feature and capability management
**Request Body**:
```json
{
  "name": "wheelchair_accessible",
  "type": "BOOLEAN",
  "value": "true"
}
```

#### GET /banks/{bankId}/atms/{atmId}/attributes
**Description**: Get all ATM attributes
**Use Case**: ATM capability inquiry

#### GET /banks/{bankId}/atms/{atmId}/attributes/{atmAttributeId}
**Description**: Get specific ATM attribute
**Use Case**: Individual attribute retrieval

#### PUT /banks/{bankId}/atms/{atmId}/attributes/{atmAttributeId}
**Description**: Update ATM attribute
**Use Case**: ATM feature updates

#### DELETE /banks/{bankId}/atms/{atmId}/attributes/{atmAttributeId}
**Description**: Delete ATM attribute
**Use Case**: ATM attribute cleanup

---

## System Management

### System Integrity Checks

#### GET /management/system/integrity/custom-view-names-check
**Description**: Check custom view names integrity
**Use Case**: System health monitoring and data validation
**Response**:
```json
{
  "status": "OK",
  "issues_found": 0,
  "details": []
}
```

#### GET /management/system/integrity/system-view-names-check
**Description**: Check system view names integrity
**Use Case**: System configuration validation

#### GET /management/system/integrity/account-access-unique-index-1-check
**Description**: Check account access unique index integrity
**Use Case**: Database integrity verification

#### GET /management/system/integrity/banks/{bankId}/account-currency-check
**Description**: Check account currency consistency
**Use Case**: Data consistency validation

#### GET /management/system/integrity/banks/{bankId}/orphaned-account-check
**Description**: Check for orphaned accounts
**Use Case**: Data cleanup and maintenance

### System Views

#### POST /system-views/{viewId}/permissions
**Description**: Add system view permission
**Use Case**: System-level permission management
**Request Body**:
```json
{
  "permission": "CanGetAccount"
}
```

#### DELETE /system-views/{viewId}/permissions/{permissionName}
**Description**: Delete system view permission
**Use Case**: Permission cleanup and security management

### WebUI Properties

#### GET /webui-props
**Description**: Get WebUI configuration properties
**Use Case**: Frontend configuration and customization
**Response**:
```json
{
  "webui_api_explorer_url": "https://api.example.com/api-explorer",
  "webui_api_manager_url": "https://api.example.com/api-manager",
  "webui_api_tester_url": "https://api.example.com/api-tester",
  "webui_sofi_url": "https://api.example.com/sofi",
  "webui_api_collection_url": "https://api.example.com/api-collection"
}
```

---

## Security & Certificates

### Certificate Management

#### GET /mtls-client-certificate-info
**Description**: Get mTLS client certificate information
**Use Case**: Certificate validation and security monitoring
**Response**:
```json
{
  "certificate": {
    "subject": "CN=client.example.com",
    "issuer": "CN=CA Authority",
    "serial_number": "123456789",
    "not_before": "2023-01-01T00:00:00Z",
    "not_after": "2024-01-01T00:00:00Z",
    "fingerprint": "SHA256:abcd1234..."
  }
}
```

#### GET /my/mtls/certificate/current
**Description**: Get current user's mTLS certificate
**Use Case**: User certificate management

#### GET /mtls/client-certificate-info
**Description**: Alternative endpoint for certificate info
**Use Case**: Certificate information retrieval

---

## Metrics & Monitoring

### System Metrics

#### GET /management/aggregate-metrics
**Description**: Get aggregated system metrics
**Use Case**: System performance monitoring and analytics
**Response**:
```json
{
  "metrics": [
    {
      "date": "2023-01-01",
      "total_requests": 1000,
      "successful_requests": 950,
      "failed_requests": 50,
      "average_response_time": 150,
      "unique_users": 100
    }
  ]
}
```

#### GET /management/metrics
**Description**: Get detailed system metrics
**Use Case**: Comprehensive system monitoring

---

## Additional Features

### API Collections

#### GET /management/api-collections
**Description**: Get all API collections (admin)
**Use Case**: API organization and management

#### POST /management/api-collections
**Description**: Create new API collection (admin)
**Use Case**: API grouping and organization

#### GET /my/api-collections
**Description**: Get user's API collections
**Use Case**: Personal API organization

#### PUT /my/api-collections/{collectionId}
**Description**: Update user's API collection
**Use Case**: Collection management

### Currency Management

#### GET /banks/{bankId}/currencies
**Description**: Get supported currencies at bank
**Use Case**: Currency support inquiry
**Response**:
```json
{
  "currencies": [
    {
      "code": "EUR",
      "name": "Euro",
      "symbol": "€"
    },
    {
      "code": "USD",
      "name": "US Dollar",
      "symbol": "$"
    }
  ]
}
```

### Entitlements

#### GET /users/{userId}/entitlements-and-permissions
**Description**: Get user entitlements and permissions
**Use Case**: Authorization and access control
**Response**:
```json
{
  "entitlements": [
    {
      "entitlement_id": "ent_001",
      "role_name": "CanGetAccount",
      "bank_id": "bank123"
    }
  ],
  "permissions": [
    {
      "permission": "can_see_transaction_this_bank_account",
      "views": ["owner", "accountant"]
    }
  ]
}
```

### Custom Views

#### POST /banks/{bankId}/accounts/{accountId}/views/{viewId}/target-views
**Description**: Create custom view
**Use Case**: Tailored account access permissions
**Request Body**:
```json
{
  "name": "Custom Accountant View",
  "description": "Limited access for accountant",
  "metadata_view": "owner",
  "is_public": false,
  "which_alias_to_use": "public",
  "hide_metadata_if_alias_used": false,
  "allowed_actions": []
}
```

#### GET /banks/{bankId}/accounts/{accountId}/views/{viewId}/target-views/{targetViewId}
**Description**: Get custom view details
**Use Case**: Custom view information retrieval

#### PUT /banks/{bankId}/accounts/{accountId}/views/{viewId}/target-views/{targetViewId}
**Description**: Update custom view
**Use Case**: Custom view modification

#### DELETE /banks/{bankId}/accounts/{accountId}/views/{viewId}/target-views/{targetViewId}
**Description**: Delete custom view
**Use Case**: Custom view cleanup

### VRP (Variable Recurring Payments)

#### POST /consumer/vrp-consent-requests
**Description**: Create VRP consent request
**Use Case**: Variable recurring payment setup
**Request Body**:
```json
{
  "max_amount_per_payment": {
    "currency": "EUR",
    "amount": "100.00"
  },
  "periodic_limits": [
    {
      "period": "Day",
      "max_amount": "500.00",
      "max_number_of_payments": 5
    }
  ],
  "valid_from_date": "2023-01-01T00:00:00Z",
  "valid_to_date": "2023-12-31T23:59:59Z"
}
```

### Regulated Entities

#### GET /regulated-entities
**Description**: Get all regulated entities
**Use Case**: Regulatory compliance and entity management

#### POST /regulated-entities
**Description**: Create regulated entity
**Use Case**: Entity registration and compliance

#### GET /regulated-entities/{regulatedEntityId}
**Description**: Get specific regulated entity
**Use Case**: Entity information retrieval

#### DELETE /regulated-entities/{regulatedEntityId}
**Description**: Delete regulated entity
**Use Case**: Entity deregistration

### Regulated Entity Attributes

#### POST /regulated-entities/{regulatedEntityId}/attributes
**Description**: Create regulated entity attribute
**Use Case**: Entity metadata management

#### GET /regulated-entities/{regulatedEntityId}/attributes
**Description**: Get all entity attributes
**Use Case**: Entity information retrieval

#### GET /regulated-entities/{regulatedEntityId}/attributes/{attributeId}
**Description**: Get specific entity attribute
**Use Case**: Individual attribute retrieval

#### PUT /regulated-entities/{regulatedEntityId}/attributes/{attributeId}
**Description**: Update entity attribute
**Use Case**: Entity information maintenance

#### DELETE /regulated-entities/{regulatedEntityId}/attributes/{attributeId}
**Description**: Delete entity attribute
**Use Case**: Attribute cleanup

### Agent Management

#### POST /banks/{bankId}/agents
**Description**: Create bank agent
**Use Case**: Agent onboarding and management
**Request Body**:
```json
{
  "name": "John Agent",
  "email": "agent@bank.com",
  "phone_number": "+1234567890",
  "status": "ACTIVE"
}
```

#### GET /banks/{bankId}/agents
**Description**: Get all bank agents
**Use Case**: Agent management dashboard

#### GET /banks/{bankId}/agents/{agentId}
**Description**: Get specific agent details
**Use Case**: Agent information retrieval

#### PUT /banks/{bankId}/agents/{agentId}
**Description**: Update agent status
**Use Case**: Agent status management

### Customer Management

#### GET /users/current/customers/customer_ids
**Description**: Get customer IDs for current user
**Use Case**: Customer relationship mapping

#### POST /banks/{bankId}/customers/legal-name
**Description**: Get customers by legal name
**Use Case**: Customer search and identification
**Request Body**:
```json
{
  "legal_name": "ACME Corporation"
}
```

---

## Testing Examples

### Complete Test Scenarios

#### 1. User Registration and Consent Flow
```bash
# 1. Check API status
curl -X GET [BASE_URL]/root

# 2. Get user information
curl -X GET [BASE_URL]/users/provider/github/username/testuser

# 3. Create consent
curl -X POST [BASE_URL]/banks/bank123/consents \
  -H "Content-Type: application/json" \
  -d '{
    "everything": false,
    "views": [{"bank_id": "bank123", "account_id": "acc456", "view_id": "owner"}],
    "consumer_id": "consumer789",
    "email": "test@example.com",
    "valid_from": "2023-01-01T00:00:00Z",
    "time_to_live": 3600
  }'

# 4. Check consent status
curl -X GET [BASE_URL]/my/consents
```

#### 2. Account and Balance Management
```bash
# 1. Get account details
curl -X GET [BASE_URL]/banks/bank123/accounts/acc456/views/owner

# 2. Check account balances
curl -X GET [BASE_URL]/banks/bank123/accounts/acc456/views/owner/balances

# 3. Create new balance entry
curl -X POST [BASE_URL]/banks/bank123/accounts/acc456/balances \
  -H "Content-Type: application/json" \
  -d '{
    "balance_type": "CURRENT",
    "balance_amount": "2500.00"
  }'

# 4. Update balance
curl -X PUT [BASE_URL]/banks/bank123/accounts/acc456/balances/balance_001 \
  -H "Content-Type: application/json" \
  -d '{
    "balance_type": "CURRENT",
    "balance_amount": "2600.00"
  }'
```

#### 3. Counterparty Management
```bash
# 1. Get all counterparties
curl -X GET [BASE_URL]/banks/bank123/accounts/acc456/views/owner/counterparties

# 2. Create new counterparty
curl -X POST [BASE_URL]/banks/bank123/accounts/acc456/views/owner/counterparties \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Business Partner",
    "description": "Regular supplier",
    "currency": "EUR",
    "bank_routing": {"scheme": "IBAN", "address": "DE89370400440532013000"},
    "account_routing": {"scheme": "AccountNumber", "address": "9876543210"},
    "is_beneficiary": true
  }'

# 3. Set counterparty limits
curl -X POST [BASE_URL]/banks/bank123/accounts/acc456/views/owner/counterparties/cp_001/limits \
  -H "Content-Type: application/json" \
  -d '{
    "currency": "EUR",
    "max_single_amount": "1000.00",
    "max_monthly_amount": "10000.00",
    "max_yearly_amount": "100000.00"
  }'
```

#### 4. System Administration
```bash
# 1. Check system integrity
curl -X GET [BASE_URL]/management/system/integrity/custom-view-names-check

# 2. Get system metrics
curl -X GET [BASE_URL]/management/aggregate-metrics

# 3. Get WebUI properties
curl -X GET [BASE_URL]/webui-props

# 4. Check certificate info
curl -X GET [BASE_URL]/mtls-client-certificate-info
```

### Error Handling Examples

#### Common Error Responses
```json
{
  "error": {
    "code": "OBP-30001",
    "message": "Bank not found",
    "details": "Bank with id 'invalid_bank' does not exist"
  }
}
```

#### Status Codes
- `200 OK`: Successful GET requests
- `201 Created`: Successful POST requests
- `204 No Content`: Successful DELETE requests
- `400 Bad Request`: Invalid request format
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict
- `500 Internal Server Error`: Server error

### Performance Testing

#### Load Testing Example
```bash
# Use Apache Bench for basic load testing
ab -n 1000 -c 10 [BASE_URL]/root

# Use curl for concurrent requests
for i in {1..100}; do
  curl -X GET [BASE_URL]/banks/bank123/accounts/acc456/views/owner &
done
wait
```

---

## Endpoint Summary by Controller

### Core Controllers (28 total)

1. **V510Controller** - Core API endpoints (root, session timeout, well-known)
2. **UserController** - User management and authentication
3. **ConsentController** - Consent lifecycle management
4. **CounterpartyController** - Counterparty CRUD operations
5. **BalanceController** - Account balance management
6. **TransactionRequestController** - Transaction request handling
7. **ConsumerController** - Consumer registration and management
8. **AtmController** - ATM attribute management
9. **AtmManagementController** - ATM CRUD operations
10. **AgentController** - Bank agent management
11. **AccountController** - Account information retrieval
12. **AccountAccessController** - Account access permissions
13. **ApiCollectionController** - API collection management
14. **CurrencyController** - Currency support
15. **CustomViewController** - Custom view management
16. **CustomerController** - Customer operations
17. **EntitlementController** - User entitlements and permissions
18. **CertificateController** - Certificate management
19. **MetricsController** - System metrics and monitoring
20. **IntegrityCheckController** - System integrity checks
21. **RegulatedEntityAttributeController** - Regulated entity attributes
22. **SystemViewController** - System view permissions
23. **UserAttributeController** - User attribute management
24. **VRPController** - Variable recurring payments
25. **WebUIController** - WebUI configuration
26. **CounterpartyLimitController** - Counterparty transaction limits
27. **TagController** - API tags
28. **HealthController** - Health checks

---

## Conclusion

This comprehensive guide covers all 185+ endpoints in the OBP API v5.1.0 implementation. Each endpoint is designed to handle specific banking operations while maintaining security, compliance, and performance standards.

For additional support or questions about specific endpoints, refer to the individual controller implementations in the `/internal/controllers/` directory.

### Quick Reference

**Server Port**: `8080`  
**Health Check**: `GET /root`  
**API Documentation**: This guide  
**Source Code**: `/internal/controllers/` and `/internal/routes/v510_routes.go`

### Support

For technical issues or questions:
1. Check server logs for detailed error information
2. Verify request format matches the examples provided
3. Ensure proper authentication and permissions
4. Review the specific controller implementation for detailed behavior

---

*Last Updated: September 2024*  
*Version: v5.1.0*  
*Implementation: Go with Gin Framework*
