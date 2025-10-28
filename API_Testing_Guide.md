# Open Bank Project API - Comprehensive Testing Guide

## Overview

This guide provides complete testing information for all implemented REST API endpoints in the Open Bank Project (OBP) API backend service. The server runs on **port 8080** with base URL: `http://YOUR_SERVER:8080`

## Server Information

- **Base URL**: `http://YOUR_SERVER:8080` (replace YOUR_SERVER with your actual server address)
- **Total Endpoints**: 700+ across all regulatory frameworks
- **Authentication**: Bearer token required for most endpoints
- **Content-Type**: `application/json`
- **Response Format**: JSON

## Health Check Endpoints

### System Health
```bash
# Health check
GET http://YOUR_SERVER:8080/health

# API root information
GET http://YOUR_SERVER:8080/
```

**Expected Response:**
```json
{
  "status": "OK",
  "message": "OBP API Backend is running"
}
```

## OBP Core API v5.1.0 Endpoints

**Base Path**: `/obp/v5.1.0`
**Total Endpoints**: ~200+

### API Information
```bash
# Get API information
GET http://YOUR_SERVER:8080/obp/v5.1.0/root

# Get configuration
GET http://YOUR_SERVER:8080/obp/v5.1.0/config

# Get adapter information
GET http://YOUR_SERVER:8080/obp/v5.1.0/adapter

# Get rate limiting info
GET http://YOUR_SERVER:8080/obp/v5.1.0/rate-limiting

# Waiting for Godot endpoint
GET http://YOUR_SERVER:8080/obp/v5.1.0/waiting-for-godot

# Get suggested session timeout
GET http://YOUR_SERVER:8080/obp/v5.1.0/ui/suggested-session-timeout
```

### Bank Management
```bash
# Get all banks
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks

# Get specific bank
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}

# Create new bank
POST http://YOUR_SERVER:8080/obp/v5.1.0/banks
Content-Type: application/json
{
  "id": "bank123",
  "short_name": "Test Bank",
  "full_name": "Test Bank Limited",
  "logo": "https://example.com/logo.png",
  "website": "https://testbank.com"
}

# Update bank
PUT http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}
Content-Type: application/json
{
  "short_name": "Updated Bank Name",
  "full_name": "Updated Bank Limited"
}

# Delete bank
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}
```

### Account Management
```bash
# Get all accounts for a bank
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts

# Get specific account
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}/{viewId}/account

# Create new account
POST http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts
Content-Type: application/json
{
  "user_id": "user123",
  "label": "Savings Account",
  "type": "SAVINGS",
  "balance": {
    "currency": "USD",
    "amount": "1000.00"
  }
}

# Update account
PUT http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}
Content-Type: application/json
{
  "label": "Updated Account Label"
}

# Delete account
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}
```

### Transaction Management
```bash
# Get transactions for account
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}/{viewId}/transactions

# Get specific transaction
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}/{viewId}/transactions/{transactionId}/transaction

# Create transaction
POST http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/accounts/{accountId}/{viewId}/transactions
Content-Type: application/json
{
  "type": "TRANSFER",
  "amount": {
    "currency": "USD",
    "amount": "100.00"
  },
  "description": "Test transaction"
}
```

### Customer Management
```bash
# Get all customers for a bank
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/customers

# Get specific customer
GET http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/customers/{customerId}

# Create new customer
POST http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/customers
Content-Type: application/json
{
  "customer_number": "CUST001",
  "legal_name": "John Doe",
  "mobile_phone_number": "+1234567890",
  "email": "john.doe@example.com"
}

# Update customer
PUT http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/customers/{customerId}
Content-Type: application/json
{
  "legal_name": "John Smith",
  "email": "john.smith@example.com"
}

# Delete customer
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/banks/{bankId}/customers/{customerId}
```

### User Management
```bash
# Get all users
GET http://YOUR_SERVER:8080/obp/v5.1.0/users

# Get specific user
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}

# Create new user
POST http://YOUR_SERVER:8080/obp/v5.1.0/users
Content-Type: application/json
{
  "username": "testuser",
  "email": "test@example.com",
  "first_name": "Test",
  "last_name": "User"
}

# Update user
PUT http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}
Content-Type: application/json
{
  "first_name": "Updated",
  "last_name": "Name"
}

# Delete user
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}
```

### User Attributes Management
```bash
# Create user attribute (non-personal)
POST http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/non-personal/attributes
Content-Type: application/json
{
  "name": "preference",
  "type": "STRING",
  "value": "email_notifications"
}

# Get user attributes (non-personal)
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/non-personal/attributes

# Delete user attribute (non-personal)
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/non-personal/attributes/{userAttributeId}

# Get user attributes by user
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/attributes

# Create user attribute for user
POST http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/attributes
Content-Type: application/json
{
  "name": "department",
  "type": "STRING",
  "value": "finance"
}

# Update user attribute
PUT http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/attributes/{userAttributeId}
Content-Type: application/json
{
  "value": "updated_value"
}

# Delete user attribute for user
DELETE http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/attributes/{userAttributeId}
```

### User Synchronization
```bash
# Sync user with external provider
POST http://YOUR_SERVER:8080/obp/v5.1.0/users/sync/{provider}/{providerId}
Content-Type: application/json
{
  "username": "external_user",
  "email": "external@example.com"
}

# Sync user (new endpoint)
POST http://YOUR_SERVER:8080/obp/v5.1.0/users/sync-new/{provider}/{providerId}
Content-Type: application/json
{
  "user_data": {
    "username": "new_external_user",
    "email": "new_external@example.com"
  }
}
```

### User Account Access
```bash
# Get user accounts at specific bank
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/accounts-at-bank/{bankId}

# Get all user accounts
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/accounts

# Get user entitlements and permissions
GET http://YOUR_SERVER:8080/obp/v5.1.0/users/{userId}/entitlements-and-permissions
```

## OBP Core API v4.0.0 Endpoints

**Base Path**: `/obp/v4.0.0`
**Total Endpoints**: ~150+

### API Information
```bash
# Get API information
GET http://YOUR_SERVER:8080/obp/v4.0.0/root

# Get database information
GET http://YOUR_SERVER:8080/obp/v4.0.0/database

# Get logout link
GET http://YOUR_SERVER:8080/obp/v4.0.0/logout-link
```

### User Management
```bash
# Create user with roles
POST http://YOUR_SERVER:8080/obp/v4.0.0/users
Content-Type: application/json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "first_name": "New",
  "last_name": "User",
  "roles": ["customer"]
}

# Get entitlements
GET http://YOUR_SERVER:8080/obp/v4.0.0/entitlements

# Get entitlements for bank
GET http://YOUR_SERVER:8080/obp/v4.0.0/banks/{bankId}/entitlements

# Lock user
POST http://YOUR_SERVER:8080/obp/v4.0.0/users/{userId}/lock
```

### Dynamic Entities
```bash
# Get system dynamic entities
GET http://YOUR_SERVER:8080/obp/v4.0.0/management/system-dynamic-entities

# Create system dynamic entity
POST http://YOUR_SERVER:8080/obp/v4.0.0/management/system-dynamic-entities
Content-Type: application/json
{
  "entity_name": "CustomEntity",
  "fields": [
    {"name": "field1", "type": "string"},
    {"name": "field2", "type": "number"}
  ]
}

# Update system dynamic entity
PUT http://YOUR_SERVER:8080/obp/v4.0.0/management/system-dynamic-entities/{entityId}
Content-Type: application/json
{
  "entity_name": "UpdatedEntity"
}

# Delete system dynamic entity
DELETE http://YOUR_SERVER:8080/obp/v4.0.0/management/system-dynamic-entities/{entityId}
```

### Settlement Accounts
```bash
# Get settlement accounts
GET http://YOUR_SERVER:8080/obp/v4.0.0/settlement-accounts

# Create settlement account
POST http://YOUR_SERVER:8080/obp/v4.0.0/settlement-accounts
Content-Type: application/json
{
  "account_id": "settlement_001",
  "bank_id": "bank123",
  "currency": "USD"
}

# Get specific settlement account
GET http://YOUR_SERVER:8080/obp/v4.0.0/settlement-accounts/{accountId}

# Update settlement account
PUT http://YOUR_SERVER:8080/obp/v4.0.0/settlement-accounts/{accountId}
Content-Type: application/json
{
  "status": "ACTIVE"
}

# Delete settlement account
DELETE http://YOUR_SERVER:8080/obp/v4.0.0/settlement-accounts/{accountId}
```

### Transaction Requests
```bash
# Get transaction request types
GET http://YOUR_SERVER:8080/obp/v4.0.0/banks/{bankId}/transaction-request-types

# Create REFUND transaction request
POST http://YOUR_SERVER:8080/obp/v4.0.0/banks/{bankId}/accounts/{accountId}/{viewId}/transaction-request-types/REFUND/transaction-requests
Content-Type: application/json
{
  "value": {
    "currency": "USD",
    "amount": "50.00"
  },
  "description": "Refund for transaction",
  "original_transaction_id": "trans123"
}

# Get REFUND transaction request
GET http://YOUR_SERVER:8080/obp/v4.0.0/banks/{bankId}/accounts/{accountId}/{viewId}/transaction-request-types/REFUND/transaction-requests/{transactionRequestId}

# Answer REFUND transaction request challenge
POST http://YOUR_SERVER:8080/obp/v4.0.0/banks/{bankId}/accounts/{accountId}/{viewId}/transaction-request-types/REFUND/transaction-requests/{transactionRequestId}/challenge
Content-Type: application/json
{
  "id": "challenge123",
  "answer": "12345"
}
```

## OBP Core API v3.1.0 Endpoints

**Base Path**: `/obp/v3.1.0`
**Total Endpoints**: ~200+

### API Information
```bash
# Get API information
GET http://YOUR_SERVER:8080/obp/v3.1.0/root

# Get configuration
GET http://YOUR_SERVER:8080/obp/v3.1.0/config

# Get adapter information
GET http://YOUR_SERVER:8080/obp/v3.1.0/adapter

# Get rate limiting information
GET http://YOUR_SERVER:8080/obp/v3.1.0/rate-limiting
```

### Product Management
```bash
# Create product
POST http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products
Content-Type: application/json
{
  "code": "SAVINGS_001",
  "name": "Premium Savings Account",
  "category": "SAVINGS",
  "family": "DEPOSIT",
  "super_family": "BANKING"
}

# Get all products
GET http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products

# Get specific product
GET http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}

# Update product
PUT http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}
Content-Type: application/json
{
  "name": "Updated Premium Savings Account"
}

# Delete product
DELETE http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}
```

### Product Attributes
```bash
# Create product attribute
POST http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}/attributes
Content-Type: application/json
{
  "name": "interest_rate",
  "type": "DOUBLE",
  "value": "2.5"
}

# Get product attributes
GET http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}/attributes

# Update product attribute
PUT http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}/attributes/{attributeId}
Content-Type: application/json
{
  "value": "3.0"
}

# Delete product attribute
DELETE http://YOUR_SERVER:8080/obp/v3.1.0/banks/{bankId}/products/{productCode}/attributes/{attributeId}
```

### Webhook Management
```bash
# Get webhooks
GET http://YOUR_SERVER:8080/obp/v3.1.0/webhooks

# Create webhook
POST http://YOUR_SERVER:8080/obp/v3.1.0/webhooks
Content-Type: application/json
{
  "url": "https://example.com/webhook",
  "http_method": "POST",
  "http_protocol": "HTTP/1.1"
}

# Get specific webhook
GET http://YOUR_SERVER:8080/obp/v3.1.0/webhooks/{webhookId}

# Update webhook
PUT http://YOUR_SERVER:8080/obp/v3.1.0/webhooks/{webhookId}
Content-Type: application/json
{
  "url": "https://updated.example.com/webhook"
}

# Delete webhook
DELETE http://YOUR_SERVER:8080/obp/v3.1.0/webhooks/{webhookId}
```

## UK Open Banking v3.1.0 Endpoints

**Base Path**: `/open-banking/v3.1.0`
**Total Endpoints**: ~60+

### Account Information Service Provider (AISP)
```bash
# Get accounts
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts

# Get specific account
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}

# Get account balances
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/balances

# Get account transactions
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/transactions

# Get account statements
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/statements

# Get account direct debits
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/direct-debits

# Get account standing orders
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/standing-orders

# Get account products
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/product

# Get account offers
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/offers

# Get account party
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/party

# Get account beneficiaries
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/aisp/accounts/{AccountId}/beneficiaries
```

### Payment Initiation Service Provider (PISP)
```bash
# Create domestic payment consent
POST http://YOUR_SERVER:8080/open-banking/v3.1.0/pisp/domestic-payment-consents
Content-Type: application/json
{
  "Data": {
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "EndToEndIdentification": "FRESCO.21302.GFX.20",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "GBP"
      },
      "CreditorAccount": {
        "SchemeName": "UK.OBIE.SortCodeAccountNumber",
        "Identification": "08080021325698",
        "Name": "ACME Inc"
      }
    }
  }
}

# Get domestic payment consent
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/pisp/domestic-payment-consents/{ConsentId}

# Create domestic payment
POST http://YOUR_SERVER:8080/open-banking/v3.1.0/pisp/domestic-payments
Content-Type: application/json
{
  "Data": {
    "ConsentId": "58923",
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "EndToEndIdentification": "FRESCO.21302.GFX.20",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "GBP"
      }
    }
  }
}

# Get domestic payment
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/pisp/domestic-payments/{DomesticPaymentId}
```

### Confirmation of Funds Service Provider (CBPII)
```bash
# Create funds confirmation consent
POST http://YOUR_SERVER:8080/open-banking/v3.1.0/cbpii/funds-confirmation-consents
Content-Type: application/json
{
  "Data": {
    "DebtorAccount": {
      "SchemeName": "UK.OBIE.SortCodeAccountNumber",
      "Identification": "11280001234567",
      "Name": "Andrea Smith"
    }
  }
}

# Get funds confirmation consent
GET http://YOUR_SERVER:8080/open-banking/v3.1.0/cbpii/funds-confirmation-consents/{ConsentId}

# Create funds confirmation
POST http://YOUR_SERVER:8080/open-banking/v3.1.0/cbpii/funds-confirmations
Content-Type: application/json
{
  "Data": {
    "ConsentId": "88379",
    "Reference": "Purchase01",
    "InstructedAmount": {
      "Amount": "20.00",
      "Currency": "GBP"
    }
  }
}
```

## Bahrain Open Banking Framework v1.0.0 Endpoints

**Base Path**: `/bahrain-obf/v1.0.0`
**Total Endpoints**: ~80+

### Account Information
```bash
# Get accounts
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts

# Get specific account
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}

# Get account balances
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/balances

# Get account supplementary info
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/supplementary-account-info

# Get account transactions
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/transactions

# Get account statements
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/statements

# Get account standing orders
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/standing-orders

# Get account direct debits
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/direct-debits

# Get account beneficiaries
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/beneficiaries

# Get account party
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/accounts/{AccountId}/party
```

### Domestic Payments
```bash
# Create domestic payment consent
POST http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-payment-consents
Content-Type: application/json
{
  "Data": {
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "EndToEndIdentification": "FRESCO.21302.GFX.20",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "BHD"
      },
      "CreditorAccount": {
        "SchemeName": "BH.CBB.IBAN",
        "Identification": "BH67BMAG00001299123456",
        "Name": "ACME Inc"
      }
    }
  }
}

# Get domestic payment consent
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-payment-consents/{ConsentId}

# Create domestic payment
POST http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-payments
Content-Type: application/json
{
  "Data": {
    "ConsentId": "58923",
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "EndToEndIdentification": "FRESCO.21302.GFX.20",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "BHD"
      }
    }
  }
}

# Get domestic payment
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-payments/{DomesticPaymentId}

# Get domestic payment details
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-payments/{DomesticPaymentId}/payment-details
```

### Future Dated Payments
```bash
# Create domestic future dated payment consent
POST http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payment-consents
Content-Type: application/json
{
  "Data": {
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "EndToEndIdentification": "FRESCO.21302.GFX.20",
      "RequestedExecutionDateTime": "2024-05-02T00:00:00+00:00",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "BHD"
      }
    }
  }
}

# Get domestic future dated payment consent
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payment-consents/{ConsentId}

# Create domestic future dated payment
POST http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payments
Content-Type: application/json
{
  "Data": {
    "ConsentId": "58923",
    "Initiation": {
      "InstructionIdentification": "ACME412",
      "RequestedExecutionDateTime": "2024-05-02T00:00:00+00:00",
      "InstructedAmount": {
        "Amount": "165.88",
        "Currency": "BHD"
      }
    }
  }
}

# Get domestic future dated payment
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payments/{DomesticFutureDatedPaymentId}

# Update domestic future dated payment
PATCH http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payments/{DomesticFutureDatedPaymentId}
Content-Type: application/json
{
  "Data": {
    "Status": "Cancelled"
  }
}

# Get domestic future dated payment details
GET http://YOUR_SERVER:8080/bahrain-obf/v1.0.0/domestic-future-dated-payments/{DomesticFutureDatedPaymentId}/payment-details
```

## Berlin Group PSD2 v1.3 Endpoints

**Base Path**: `/berlin-group/v1.3`
**Total Endpoints**: ~20+

### Account Information
```bash
# Get accounts
GET http://YOUR_SERVER:8080/berlin-group/v1.3/accounts

# Get specific account
GET http://YOUR_SERVER:8080/berlin-group/v1.3/accounts/{account-id}

# Get account balances
GET http://YOUR_SERVER:8080/berlin-group/v1.3/accounts/{account-id}/balances

# Get account transactions
GET http://YOUR_SERVER:8080/berlin-group/v1.3/accounts/{account-id}/transactions
```

### Payment Initiation
```bash
# Initiate payment
POST http://YOUR_SERVER:8080/berlin-group/v1.3/payments/{payment-product}
Content-Type: application/json
{
  "instructedAmount": {
    "currency": "EUR",
    "amount": "123.50"
  },
  "creditorAccount": {
    "iban": "DE23100120020123456789"
  },
  "creditorName": "Merchant123",
  "remittanceInformationUnstructured": "Ref Number Merchant"
}

# Get payment information
GET http://YOUR_SERVER:8080/berlin-group/v1.3/payments/{payment-product}/{paymentId}

# Get payment status
GET http://YOUR_SERVER:8080/berlin-group/v1.3/payments/{payment-product}/{paymentId}/status
```

## Australian Consumer Data Right (CDR) v1.0.0 Endpoints

**Base Path**: `/cds-au/v1.0.0`
**Total Endpoints**: ~20+

### Banking Data
```bash
# Get accounts
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/accounts

# Get specific account
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/accounts/{accountId}

# Get account balance
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/accounts/{accountId}/balance

# Get account transactions
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/accounts/{accountId}/transactions

# Get products
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/products

# Get specific product
GET http://YOUR_SERVER:8080/cds-au/v1.0.0/banking/products/{productId}
```

## Polish API v2.1.1.1 Endpoints

**Base Path**: `/polish-api/v2.1.1.1`
**Total Endpoints**: ~10+

### Account Information
```bash
# Get accounts
GET http://YOUR_SERVER:8080/polish-api/v2.1.1.1/accounts

# Get account details
GET http://YOUR_SERVER:8080/polish-api/v2.1.1.1/accounts/{accountId}

# Get account balances
GET http://YOUR_SERVER:8080/polish-api/v2.1.1.1/accounts/{accountId}/balances

# Get account transactions
GET http://YOUR_SERVER:8080/polish-api/v2.1.1.1/accounts/{accountId}/transactions
```

## STET API v1.4 Endpoints

**Base Path**: `/stet/v1.4`
**Total Endpoints**: ~10+

### Account Information
```bash
# Get accounts
GET http://YOUR_SERVER:8080/stet/v1.4/accounts

# Get account details
GET http://YOUR_SERVER:8080/stet/v1.4/accounts/{accountId}

# Get account balances
GET http://YOUR_SERVER:8080/stet/v1.4/accounts/{accountId}/balances

# Get account transactions
GET http://YOUR_SERVER:8080/stet/v1.4/accounts/{accountId}/transactions
```

## MxOF API v1.0.0 Endpoints

**Base Path**: `/mxof/v1.0.0`
**Total Endpoints**: ~10+

### Account Information
```bash
# Get accounts
GET http://YOUR_SERVER:8080/mxof/v1.0.0/accounts

# Get account details
GET http://YOUR_SERVER:8080/mxof/v1.0.0/accounts/{accountId}

# Get account balances
GET http://YOUR_SERVER:8080/mxof/v1.0.0/accounts/{accountId}/balances

# Get account transactions
GET http://YOUR_SERVER:8080/mxof/v1.0.0/accounts/{accountId}/transactions
```

## Authentication

Most endpoints require authentication. Include the Authorization header:

```bash
Authorization: Bearer {your_access_token}
```

## Common HTTP Status Codes

- **200 OK**: Request successful
- **201 Created**: Resource created successfully
- **400 Bad Request**: Invalid request data
- **401 Unauthorized**: Authentication required
- **403 Forbidden**: Insufficient permissions
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server error

## Testing Tools

### Using curl
```bash
# Example GET request
curl -X GET "http://YOUR_SERVER:8080/obp/v5.1.0/banks" \
  -H "Authorization: Bearer your_token" \
  -H "Content-Type: application/json"

# Example POST request
curl -X POST "http://YOUR_SERVER:8080/obp/v5.1.0/banks" \
  -H "Authorization: Bearer your_token" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "bank123",
    "short_name": "Test Bank",
    "full_name": "Test Bank Limited"
  }'
```

### Using Postman
1. Import the API collection
2. Set base URL to `http://YOUR_SERVER:8080`
3. Configure authentication in the Authorization tab
4. Set Content-Type to `application/json` for POST/PUT requests

## Error Handling

All endpoints return consistent error responses:

```json
{
  "error": "Error description",
  "code": "ERROR_CODE",
  "message": "Detailed error message"
}
```

## Rate Limiting

The API implements rate limiting. Check the following headers in responses:
- `X-RateLimit-Limit`: Request limit per time window
- `X-RateLimit-Remaining`: Remaining requests in current window
- `X-RateLimit-Reset`: Time when the rate limit resets

## Support

For technical support or questions about the API:
- Check the comprehensive verification report for endpoint coverage
- Review the swagger specifications in the 01.phase-1-output folder
- Test endpoints using the examples provided in this guide

---

**Last Updated**: September 2025  
**API Version**: OBP Core v5.1.0  
**Total Endpoints**: 700+  
**Server**: http://YOUR_SERVER:8080
