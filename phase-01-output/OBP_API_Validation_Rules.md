# Open Bank Project (OBP) API - Validation Rules Documentation

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [How to Use This Document](#how-to-use-this-document)

**Core Validation Categories (86 rules):**
3. [Authentication & Authorization Validations](#authentication--authorization-validations) - 9 rules
4. [Account Validations](#account-validations) - 10 rules
5. [Customer Validations](#customer-validations) - 9 rules
6. [Transaction Validations](#transaction-validations) - 11 rules
7. [Data Format Validations](#data-format-validations) - 9 rules
8. [Rate Limiting Validations](#rate-limiting-validations) - 3 rules
9. [View Permissions & Entitlements](#view-permissions--entitlements) - 5 rules
10. [KYC & Compliance Validations](#kyc--compliance-validations) - 6 rules
11. [Business Logic Validations](#business-logic-validations) - 13 rules
12. [API-Specific Validations](#api-specific-validations) - 11 rules

**Extended Validation Categories (30 rules):**
13. [Dynamic Entity Validations](#dynamic-entity-validations) - 5 rules
14. [Request Header Validations (Berlin Group / PSD2)](#request-header-validations-berlin-group--psd2) - 5 rules
15. [Attribute System Validations](#attribute-system-validations) - 5 rules
16. [FX Rate Validations](#fx-rate-validations) - 4 rules
17. [Product Fee and Product Collection Validations](#product-fee-and-product-collection-validations) - 4 rules
18. [Inline Validation Pattern (Helper.booleanToFuture)](#inline-validation-pattern-helperbooleanToFuture) - 1 rule
19. [Secure Logging Pattern Validations](#secure-logging-pattern-validations) - 6 rules

**Reference Sections:**
20. [Error Code Reference](#error-code-reference)
21. [Quick Reference by Audience](#quick-reference-by-audience)
22. [Common Validation Scenarios](#common-validation-scenarios)

---

## Executive Summary

This document provides a comprehensive catalog of all validation rules enforced by the Open Bank Project (OBP) API system. These rules ensure data integrity, security, regulatory compliance, and proper business logic enforcement across all banking operations.

**Total Validation Rules Documented: 116 discrete validation rules** organized across 17 major categories

**Validation Categories Breakdown:**

**Original Categories (86 rules):**
1. **Authentication & Authorization** (9 rules): Password strength, username format, OAuth tokens, certificates, Direct Login
2. **Account Validations** (10 rules): Account IDs, bank IDs, IBAN, routing, balance requirements, uniqueness
3. **Customer Validations** (9 rules): Customer numbers, phone, email, legal name, user-customer links, attributes
4. **Transaction Validations** (11 rules): Amount, currency, status, challenges, charge policy, counterparty
5. **Data Format Validations** (9 rules): Currency codes, phone numbers, dates, UUIDs, locales, URLs
6. **Rate Limiting Validations** (3 rules): Per-consumer limits, anonymous limits, endpoint-specific limits
7. **View Permissions & Entitlements** (5 rules): 50+ granular view permissions, role-based access control
8. **KYC & Compliance Validations** (6 rules): KYC status, checks, documents, credit rating, consents
9. **Business Logic Validations** (13 rules): Berlin Group rules, counterparty limits, product hierarchy, webhooks, agent accounts
10. **API-Specific Validations** (11 rules): JSON schema, Berlin Group headers, dynamic entities, OTP, Elasticsearch

**New Categories Added (30 rules):**
11. **Dynamic Entity Validations** (5 rules): Required fields, field type matching, minLength, maxLength, reference integrity
12. **Request Header Validations** (5 rules): X-Request-ID format & uniqueness, Date header, TPP signature, certificate validation (PSD2/Berlin Group)
13. **Attribute System Validations** (5 rules): Attribute definitions, categories, names, types, value formats
14. **FX Rate Validations** (4 rules): Currency validation (ISO 4217), conversion values, effective dates
15. **Product Fee & Collection Validations** (4 rules): Fee amounts, currencies, frequency, collection code uniqueness
16. **Inline Validation Pattern** (1 rule): Helper.booleanToFuture pattern used 100+ times across endpoints
17. **Secure Logging Pattern Validations** (6 rules): Automatic masking of secrets, tokens, passwords, API keys, credit cards, JDBC URLs

**Coverage Highlights:**
- 🔐 **Security**: Password validation, OAuth, certificates, PSD2 compliance, secure logging
- 💰 **Financial Operations**: Transaction amounts, currencies, limits, FX rates, fees
- 👤 **Customer Management**: KYC, phone/email validation, user-customer links, attributes
- 🏦 **Account Operations**: Account IDs, bank IDs, IBAN, routing, balance validation
- 🔍 **Data Quality**: Format validations for phone, currency, dates, UUIDs, URLs
- 🚦 **Access Control**: 50+ granular view permissions, role-based entitlements
- 🌍 **Regulatory Compliance**: Berlin Group/PSD2 headers, consents, date ranges, frequency limits
- 🛠️ **Custom Extensions**: Dynamic entities, attributes, inline validation patterns

**Error Code Ranges:**
- `OBP-00XXX` to `OBP-09XXX`: Infrastructure and dynamic entity errors
- `OBP-10XXX`: Configuration and format errors
- `OBP-20XXX`: Authentication/Authorization errors
- `OBP-30XXX`: Resource-related errors (accounts, customers, products, etc.)
- `OBP-40XXX`: Transaction request and Berlin Group errors
- `OBP-50XXX`: Internal system errors

---

## How to Use This Document

### For Business Analysts
Focus on sections that explain business rules and requirements: Transaction Validations, Customer Validations, and Business Logic Validations.

### For Compliance Officers
Review: KYC & Compliance Validations, Authentication & Authorization, and Transaction Validations for regulatory requirements.

### For QA Teams
Use the validation details and examples to design test cases. Each rule includes valid and invalid examples.

### For Product Managers
Understand feature limitations through Account Validations, Transaction Validations, and View Permissions sections.

### For Customer Support Teams
Reference the Error Code Reference and Common Validation Scenarios for troubleshooting customer issues.

### For System Administrators
Review Rate Limiting Validations and Authentication & Authorization for configuration guidance.

### For Security Teams
Focus on Authentication & Authorization Validations and View Permissions & Entitlements sections.

---

## Authentication & Authorization Validations

### 1. Password Strength Validation

**What It Checks**: Password complexity and length requirements

**Why It Exists**: Protect user accounts from unauthorized access through strong password policies

**When It Applies**: User registration, password creation, password reset

**Who It Affects**: All users creating or updating passwords

**Validation Details**:
- **Option 1 (Simple)**: Length > 16 characters and ≤ 512 characters (no complexity requirements)
- **Option 2 (Complex)**: Length 10-16 characters with ALL of the following:
  - At least one digit (0-9)
  - At least one lowercase letter (a-z)
  - At least one uppercase letter (A-Z)
  - At least one special character: `! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ \` { | } ~`
  - Only contains valid characters from above sets

**What Happens on Failure**: 
- Error Code: `OBP-30207`
- Message: "Invalid Password Format. Your password should EITHER be at least 10 characters long and contain mixed numbers and both upper and lower case letters and at least one special character, OR the length should be > 16 and <= 512."

**Example Valid Input**:
- `MyP@ssw0rd123` (Option 2: 10+ chars with complexity)
- `this-is-a-long-secure-password-no-special-requirements` (Option 1: 17+ chars)

**Example Invalid Input**:
- `password` (too short, no complexity)
- `Password123` (missing special character)
- `abc` (too short)

---

### 2. Username Validation

**What It Checks**: Username format and character restrictions

**Why It Exists**: Ensure usernames are valid email addresses or conform to alphanumeric standards

**When It Applies**: User registration, login

**Who It Affects**: All users

**Validation Details**:
- Must be a valid email address OR
- Alphanumeric characters (A-Z, a-z, 0-9) with: `-`, `_`, `.`, `@`
- Maximum length: 512 characters

**What Happens on Failure**:
- Error Code: `OBP-20005`
- Message: "Invalid User Name. The User Name field must be a valid email address, or can contain alphanumeric characters, or characters: '-', '_', '.', '@'."

**Example Valid Input**:
- `user@example.com`
- `john.doe`
- `user_123`

**Example Invalid Input**:
- `user@` (invalid email)
- `user#name` (invalid character #)
- `user name` (contains space)

---

### 3. Consumer Key Validation

**What It Checks**: API consumer key format

**Why It Exists**: Ensure API consumer keys follow standard format for OAuth authentication

**When It Applies**: API consumer creation, OAuth authentication

**Who It Affects**: Third-party developers, API consumers

**Validation Details**:
- Only alphanumeric characters and hyphens: A-Z, a-z, 0-9, `-`
- Maximum length: 512 characters

**What Happens on Failure**:
- Error Code: `OBP-35030` (Invalid format)
- Error Code: `OBP-35031` (Too long)
- Message: "The Consumer Key must be alphanumeric. (A-Z, a-z, 0-9)" or "The Consumer Key max length <= 512"

**Example Valid Input**:
- `abcd1234-5678-efgh`
- `consumer-key-123`

**Example Invalid Input**:
- `consumer_key` (contains underscore)
- `consumer.key` (contains period)

---

### 4. Authentication Type Validation

**What It Checks**: Whether the authentication method used is allowed for a specific API operation

**Why It Exists**: Control which authentication methods (OAuth1, OAuth2, Gateway Login, etc.) can access specific endpoints

**When It Applies**: Every API request to protected endpoints

**Who It Affects**: All API consumers and users

**Validation Details**:
- Each API operation can restrict which authentication types are allowed
- Supported types: OAuth1, OAuth2, DirectLogin, GatewayLogin, Anonymous
- Configuration is operation-specific (stored per OPERATION_ID)

**What Happens on Failure**:
- Error Code: `OBP-40034`
- Message: "Current request authentication type is illegal."

**Example Valid Input**:
- Using OAuth2 for an endpoint that allows OAuth2
- Using DirectLogin for an endpoint that allows DirectLogin

**Example Invalid Input**:
- Using Anonymous access for an endpoint that requires OAuth
- Using OAuth1 for an endpoint that only allows OAuth2

---

### 5. User Account Status Validation

**What It Checks**: Whether the user account is active and not locked or deleted

**Why It Exists**: Prevent access by disabled or deleted user accounts

**When It Applies**: User authentication, API requests

**Who It Affects**: All users

**Validation Details**:
- User must not be locked
- User must not be deleted
- Account must be active

**What Happens on Failure**:
- Error Code: `OBP-20017` (User is locked)
- Error Code: `OBP-20064` (User is deleted)
- Message: "The user is locked!" or "The user is deleted!"

---

### 6. OAuth Token Validation

**What It Checks**: Validity and format of OAuth access tokens

**Why It Exists**: Ensure secure authentication and prevent unauthorized access

**When It Applies**: OAuth-authenticated API requests

**Who It Affects**: OAuth users and applications

**Validation Details**:
- Token must exist and be valid
- Token must not be expired
- Token must be linked to an active consumer
- Token must match the client certificate if certificate-based authentication is used

**What Happens on Failure**:
- Error Code: `OBP-20202` (Cannot verify JWT)
- Error Code: `OBP-20209` (No linked consumer)
- Error Code: `OBP-20210` (Certificate mismatch)
- Error Code: `OBP-20215` (Validation error)
- Various messages depending on specific failure

---

### 7. X.509 Certificate Validation

**What It Checks**: PEM encoded certificate validity for certificate-based authentication

**Why It Exists**: Enable secure authentication using client certificates (PSD2 compliance)

**When It Applies**: Certificate-based authentication requests

**Who It Affects**: Regulated entities, PSD2 TPPs (Third Party Providers)

**Validation Details**:
- Certificate must be parseable
- Certificate must not be expired
- Certificate must be currently valid (not before date must be in past)
- Certificate must contain valid PSD2 roles
- Public key must be extractable
- Signature must be verifiable

**What Happens on Failure**:
- Error Code: `OBP-20301` (Parsing failed)
- Error Code: `OBP-20302` (Certificate expired)
- Error Code: `OBP-20303` (Not yet valid)
- Error Code: `OBP-20304` (Cannot get RSA public key)
- Error Code: `OBP-20305` (Cannot get EC public key)
- Error Code: `OBP-20307` (Incorrect role for action)
- Error Code: `OBP-20308` (No PSD2 roles)

---

### 8. Direct Login Token Validation

**What It Checks**: Direct Login authentication header format and token

**Why It Exists**: Simplified authentication method alternative to OAuth

**When It Applies**: Direct Login authentication

**Who It Affects**: Users using Direct Login method

**Validation Details**:
- Authorization header must be present
- Must contain "DirectLogin" prefix
- Token must be valid and not expired

**What Happens on Failure**:
- Error Code: `OBP-20082` (Missing header)
- Error Code: `OBP-20083` (Invalid header format)
- Message: "Missing DirectLogin or Authorization header." or "Missing DirectLogin word at the value of Authorization header."

---

### 9. DAuth Validation

**What It Checks**: Distributed Authentication parameters and process

**Why It Exists**: Enable distributed authentication across multiple identity providers

**When It Applies**: DAuth login process

**Who It Affects**: Users authenticating via DAuth

**Validation Details**:
- Required parameters must be present
- JWT token must be valid and not corrupted
- Authentication must be from whitelisted addresses
- dauth.host property must be configured

**What Happens on Failure**:
- Error Code: `OBP-20066` (Missing parameters)
- Error Code: `OBP-20068` (Host property missing)
- Error Code: `OBP-20069` (Not from whitelisted address)
- Error Code: `OBP-20071` (Invalid JWT)
- Error Code: `OBP-20072` (Invalid header format)

---

## Account Validations

### 10. Account ID Format Validation

**What It Checks**: Format and characters allowed in account identifiers

**Why It Exists**: Ensure account IDs are valid and prevent injection attacks

**When It Applies**: Account creation, account access

**Who It Affects**: All users creating or accessing accounts

**Validation Details**:
- Only alphanumeric characters: A-Z, a-z, 0-9
- Special characters allowed: `-`, `_`, `.`
- Maximum length: 255 characters
- Must not be empty

**What Happens on Failure**:
- Error Code: `OBP-30110`
- Message: "Invalid Account Id. The ACCOUNT_ID should only contain 0-9/a-z/A-Z/'-'/'.'/'_', the length should be smaller than 255."

**Example Valid Input**:
- `savings-account-001`
- `CHECKING_123`
- `account.primary`

**Example Invalid Input**:
- `account@123` (contains @)
- `account#main` (contains #)
- `account id` (contains space)
- (256+ character string)

---

### 11. Bank ID Format Validation

**What It Checks**: Format and characters allowed in bank identifiers

**Why It Exists**: Ensure bank IDs follow standard format across the system

**When It Applies**: Bank creation, bank references in API calls

**Who It Affects**: System administrators, API consumers

**Validation Details**:
- Same rules as Account ID
- Only alphanumeric: A-Z, a-z, 0-9
- Special characters allowed: `-`, `_`, `.`
- Maximum length: 255 characters

**What Happens on Failure**:
- Error Code: `OBP-30111`
- Message: "Invalid Bank Id. The BANK_ID should only contain 0-9/a-z/A-Z/'-'/'.'/'_', the length should be smaller than 255."

**Example Valid Input**:
- `gh.29.uk`
- `bank-001`
- `HSBC_UK`

**Example Invalid Input**:
- `bank@uk` (contains @)
- `bank/001` (contains /)

---

### 12. Account Routing Validation

**What It Checks**: Account routing scheme and address format

**Why It Exists**: Ensure proper account identification across different payment systems

**When It Applies**: Account creation, payment routing

**Who It Affects**: Users creating accounts, payment processors

**Validation Details**:
- Routing scheme must be valid (e.g., IBAN, AccountNumber, etc.)
- Routing address must be provided
- Routing combination must be unique
- Payment system name must be valid format

**What Happens on Failure**:
- Error Code: `OBP-30075` (Routing not found)
- Error Code: `OBP-31075` (Routing not unique)
- Error Code: `OBP-30114` (Invalid routings)
- Error Code: `OBP-30115` (Routing already exists)
- Error Code: `OBP-30116` (Invalid payment system name)

**Example Valid Input**:
```json
{
  "scheme": "IBAN",
  "address": "DE89370400440532013000"
}
```

**Example Invalid Input**:
- Missing scheme or address
- Duplicate routing for an account
- Invalid payment system name format

---

### 13. Initial Balance Validation

**What It Checks**: Initial balance when creating a new account

**Why It Exists**: Business rule requiring accounts to start with zero balance

**When It Applies**: Account creation

**Who It Affects**: Users and administrators creating accounts

**Validation Details**:
- Initial balance must be exactly 0 (zero)
- Balance must be a valid number format

**What Happens on Failure**:
- Error Code: `OBP-30109` (Must be zero)
- Error Code: `OBP-30112` (Invalid number format)
- Message: "Initial Balance of Account must be Zero (0)." or "Invalid Number. Initial balance must be a number, e.g 1000.00"

**Example Valid Input**:
- `0`
- `0.00`

**Example Invalid Input**:
- `100.00`
- `1000`
- `abc` (not a number)

---

### 14. Account Type Validation

**What It Checks**: Account type field validity

**Why It Exists**: Ensure only supported account types are created

**When It Applies**: Account creation

**Who It Affects**: Users creating accounts

**Validation Details**:
- Account type must be a recognized type (checking, savings, etc.)
- Specific validation depends on bank configuration

**What Happens on Failure**:
- Error Code: `OBP-30108`
- Message: "Invalid Account Type."

---

### 15. Account Existence Check

**What It Checks**: Whether an account exists before performing operations

**Why It Exists**: Prevent operations on non-existent accounts

**When It Applies**: All account-related operations (views, transactions, etc.)

**Who It Affects**: All users accessing accounts

**Validation Details**:
- Account must exist in the system
- Account must be accessible by the requesting user
- Can check by ACCOUNT_ID, IBAN, or account routing

**What Happens on Failure**:
- Error Code: `OBP-30003` (By ACCOUNT_ID)
- Error Code: `OBP-30074` (By IBAN)
- Error Code: `OBP-30073` (By routing)
- Error Code: `OBP-30076` (Generic)
- Message: "Account not found. Please specify a valid value for ACCOUNT_ID."

---

### 16. Account Number Uniqueness

**What It Checks**: Account number uniqueness when searching by account number

**Why It Exists**: Prevent ambiguity when multiple accounts have same number

**When It Applies**: Account search by account number

**Who It Affects**: All users searching for accounts

**Validation Details**:
- Account number search must return exactly one account
- If multiple accounts found, validation fails

**What Happens on Failure**:
- Error Code: `OBP-30269`
- Message: "Finding an account by the accountNumber is ambiguous."

**Example Valid Input**:
- Unique account number returning single account

**Example Invalid Input**:
- Account number that exists for multiple accounts

---

### 17. IBAN Validation

**What It Checks**: IBAN (International Bank Account Number) format

**Why It Exists**: Ensure international account identifiers follow ISO 13616 standard

**When It Applies**: Account creation, counterparty creation, payments

**Who It Affects**: Users working with international accounts

**Validation Details**:
- Must follow IBAN format (country code + check digits + account number)
- Country-specific length requirements
- Check digit validation

**What Happens on Failure**:
- Error Code: `OBP-30074`
- Message: "Bank Account not found. Please specify a valid value for iban."

**Example Valid Input**:
- `DE89370400440532013000` (German IBAN)
- `GB82WEST12345698765432` (UK IBAN)

**Example Invalid Input**:
- `DE89` (too short)
- `XX89370400440532013000` (invalid country code)
- `DE00370400440532013000` (invalid check digits)

---

### 18. Account ID Uniqueness at Bank

**What It Checks**: Account ID uniqueness within a bank

**Why It Exists**: Prevent duplicate account IDs at the same bank

**When It Applies**: Account creation

**Who It Affects**: Users and administrators creating accounts

**Validation Details**:
- ACCOUNT_ID must be unique within the specified BANK_ID
- Combination of BANK_ID + ACCOUNT_ID must not already exist

**What Happens on Failure**:
- Error Code: `OBP-30208`
- Message: "Account_ID already exists at the Bank."

---

### 19. Account Customer Relationship Validation

**What It Checks**: Whether account-customer link is valid

**Why It Exists**: Ensure proper relationship between accounts and customers

**When It Applies**: Linking customers to accounts, account access

**Who It Affects**: Customers accessing their accounts

**Validation Details**:
- Customer must belong to the same bank as the account
- Account-customer link must exist for access
- One account can be linked to multiple customers

**What Happens on Failure**:
- Error Code: `OBP-30113` (Customer bank mismatch)
- Error Code: `OBP-30223` (Account already linked to customer)
- Message: "Invalid Bank Id. The Customer does not belong to this Bank" or "The Account is already linked to a Customer at the bank specified by BANK_ID"

---

## Customer Validations

### 20. Customer Number Availability Check

**What It Checks**: Whether a customer number is already in use

**Why It Exists**: Prevent duplicate customer numbers within a bank

**When It Applies**: Customer creation

**Who It Affects**: Bank administrators creating customer records

**Validation Details**:
- Customer number must be unique within the specified bank
- Combination of BANK_ID + CUSTOMER_NUMBER must not exist

**What Happens on Failure**:
- Error Code: `OBP-30006`
- Message: "Customer Number already exists. Please specify a different value for BANK_ID or CUSTOMER_NUMBER."

**Example Valid Input**:
- New customer number: `CUST-123456`

**Example Invalid Input**:
- Customer number that already exists: `CUST-000001`

---

### 21. Customer Existence Check

**What It Checks**: Whether a customer exists before operations

**Why It Exists**: Prevent operations on non-existent customers

**When It Applies**: All customer-related operations

**Who It Affects**: All users working with customer data

**Validation Details**:
- Can check by CUSTOMER_NUMBER, CUSTOMER_ID, or USER_ID
- Customer must exist in the system
- Customer must belong to the specified bank

**What Happens on Failure**:
- Error Code: `OBP-30002` (By CUSTOMER_NUMBER)
- Error Code: `OBP-30046` (By CUSTOMER_ID)
- Message: "Customer not found. Please specify a valid value for CUSTOMER_NUMBER." or "Customer not found. Please specify a valid value for CUSTOMER_ID."

---

### 22. Customer Phone Number Validation

**What It Checks**: Mobile phone number format

**Why It Exists**: Ensure phone numbers are valid for SMS notifications and contact

**When It Applies**: Customer creation, customer updates

**Who It Affects**: Customers providing phone numbers

**Validation Details**:
- Must start with `+` (international format)
- Maximum 15 digits after the `+`
- Only digits allowed after the `+`

**What Happens on Failure**:
- Error Code: `OBP-40017`
- Message: "Invalid Phone Number. Please specify a valid value for PHONE_NUMBER. Eg:+9722398746"

**Example Valid Input**:
- `+441234567890` (UK)
- `+19175551234` (US)
- `+9722398746` (Israel)

**Example Invalid Input**:
- `441234567890` (missing +)
- `+44-1234-567890` (contains dashes)
- `+44 1234 567890` (contains spaces)
- `+12345678901234567` (too long, >15 digits)

---

### 23. Customer Email Validation

**What It Checks**: Email address format

**Why It Exists**: Ensure valid email for customer communications

**When It Applies**: Customer creation, customer updates

**Who It Affects**: Customers providing email addresses

**Validation Details**:
- Must be a valid email format
- Standard email validation (username@domain)

**What Happens on Failure**:
- Generic validation error if email format is invalid

**Example Valid Input**:
- `customer@example.com`
- `john.doe@bank.co.uk`

**Example Invalid Input**:
- `customer@` (incomplete)
- `@example.com` (missing username)
- `customer` (missing domain)

---

### 24. Customer Legal Name Validation

**What It Checks**: Customer's legal name field

**Why It Exists**: Ensure proper customer identification for legal and regulatory purposes

**When It Applies**: Customer creation, customer updates

**Who It Affects**: All customers

**Validation Details**:
- Legal name is required for customer records
- Used for KYC and identification purposes

**What Happens on Failure**:
- Error in customer creation process

**Example Valid Input**:
- `John Michael Smith`
- `ABC Corporation Ltd`

---

### 25. Customer-User Link Validation

**What It Checks**: Whether user-customer link is valid and unique

**Why It Exists**: Ensure each user is properly linked to their customer records

**When It Applies**: Linking users to customers, customer access

**Who It Affects**: All bank customers

**Validation Details**:
- User can be linked to one customer per bank
- Link must exist for customer operations requiring user context

**What Happens on Failure**:
- Error Code: `OBP-30007` (User already linked)
- Error Code: `OBP-30008` (Link not found)
- Error Code: `OBP-30035` (Link not found - alternate)
- Message: "The User is already linked to a Customer at the bank specified by BANK_ID" or "User Customer Link not found by USER_ID"

---

### 26. Customer Attributes Validation

**What It Checks**: Custom customer attribute values

**Why It Exists**: Allow banks to store additional customer information while maintaining data quality

**When It Applies**: Customer attribute creation/update

**Who It Affects**: Banks managing customer data

**Validation Details**:
- Attribute must be defined before use
- Attribute values must match defined type/format
- Attribute IDs must be valid

**What Happens on Failure**:
- Error Code: `OBP-30069`
- Message: "Customer Attribute not found. Please specify a valid value for CUSTOMER_ATTRIBUTE_ID."

---

### 27. Customer Address Validation

**What It Checks**: Customer address information

**Why It Exists**: Maintain accurate customer address records for compliance

**When It Applies**: Customer address operations

**Who It Affects**: Customers with addresses on file

**Validation Details**:
- Address must exist when referenced
- Address fields must be properly formatted

**What Happens on Failure**:
- Error Code: `OBP-30310`
- Message: "Customer's Address not found by CUSTOMER_ADDRESS_ID."

---

### 28. Tax Residence Validation

**What It Checks**: Customer tax residence information

**Why It Exists**: Regulatory compliance for tax reporting (FATCA, CRS)

**When It Applies**: Tax residence operations

**Who It Affects**: Customers with tax obligations

**Validation Details**:
- Tax residence must exist when referenced
- Tax residence ID must be valid

**What Happens on Failure**:
- Error Code: `OBP-30300`
- Message: "Tax Residence not found by TAX_RESIDENCE_ID."

---

## Transaction Validations

### 29. Transaction Amount Validation

**What It Checks**: Transaction amount must be positive

**Why It Exists**: Prevent zero or negative value transactions

**When It Applies**: Transaction creation, payment initiation

**Who It Affects**: All users initiating transactions

**Validation Details**:
- Amount must be greater than 0
- Amount must be a valid number format
- Amount precision depends on currency

**What Happens on Failure**:
- Error Code: `OBP-40008` (Not positive)
- Error Code: `OBP-20054` (Invalid amount)
- Message: "Can't send a payment with a value of 0 or less." or "Invalid amount. Please specify a valid value for amount."

**Example Valid Input**:
- `100.50`
- `0.01`
- `1000000.00`

**Example Invalid Input**:
- `0`
- `-50.00`
- `abc` (not a number)

---

### 30. Transaction Currency Validation

**What It Checks**: Currency code format and matching

**Why It Exists**: Ensure transactions use valid currencies and match account currency

**When It Applies**: All financial transactions

**Who It Affects**: All users initiating transactions

**Validation Details**:
- Currency must be valid ISO 4217 code (e.g., USD, EUR, GBP)
- Special cases supported: XBT (Bitcoin), lovelace (Cardano), ETH (Ethereum)
- Transaction currency must match the from-account currency
- Balance currency must match account currency

**What Happens on Failure**:
- Error Code: `OBP-40003` (Currency mismatch)
- Error Code: `OBP-30105` (Invalid balance currency)
- Message: "Transaction Request Currency must be the same as From Account Currency." or "Invalid Balance Currency."

**Example Valid Input**:
- `USD`, `EUR`, `GBP`, `JPY`
- `XBT` (Bitcoin)

**Example Invalid Input**:
- `US` (not 3 letters)
- `EURO` (too long)
- `XXX` (not a valid ISO code)
- Using `EUR` for a `USD` account

---

### 31. Transaction Request Status Validation

**What It Checks**: Transaction request lifecycle status

**Why It Exists**: Ensure operations are only performed on transactions in appropriate states

**When It Applies**: Transaction processing, challenges, cancellations

**Who It Affects**: All users managing transaction requests

**Validation Details**:
- Status must be `INITIATED` for certain operations
- Status must be `INITIATED`, `NEXT_CHALLENGE_PENDING`, or `FORWARDED` for challenge answers
- Status must allow cancellation
- Status transitions must follow valid workflow

**What Happens on Failure**:
- Error Code: `OBP-40011` (Not INITIATED)
- Error Code: `OBP-40020` (Not in valid states)
- Error Code: `OBP-40023` (Cannot be cancelled)
- Error Code: `OBP-40024` (Cannot update PSU data)
- Message: "Transaction Request Status is not INITIATED." or "Transaction Request Status is not INITIATED or NEXT_CHALLENGE_PENDING or FORWARDED."

**Valid Status Values**:
- `INITIATED`
- `NEXT_CHALLENGE_PENDING`
- `COMPLETED`
- `FAILED`
- `REJECTED`
- `CANCELLED`
- `FORWARDED`
- `CANCELLATION_PENDING`

---

### 32. Transaction Challenge Validation

**What It Checks**: Challenge-response authentication for transactions

**Why It Exists**: Provide additional security for sensitive transactions (Strong Customer Authentication)

**When It Applies**: Transaction confirmation, sensitive operations

**Who It Affects**: All users confirming transactions

**Validation Details**:
- Challenge ID must be valid
- Challenge answer must be correct
- Challenge must not be expired
- Allowed attempts must not be exhausted (typically 3 attempts)
- Challenge type must be valid (SMS, EMAIL, IMPLICIT, etc.)

**What Happens on Failure**:
- Error Code: `OBP-40010` (Invalid challenge ID)
- Error Code: `OBP-40014` (Attempts used up)
- Error Code: `OBP-40015` (Invalid challenge type)
- Error Code: `OBP-40016` (Invalid answer)
- Error Code: `OBP-40021` (Invalid payment/transaction request ID)
- Error Code: `OBP-40022` (Invalid challenge ID)
- Message: "Invalid Challenge Answer. Please specify a valid value for answer in Json body. The challenge answer may be expired. Or you've used up your allowed attempts..."

**Example Valid Input**:
- SMS code: `123456`
- For DUMMY mode: `123`

**Example Invalid Input**:
- Wrong code: `999999`
- Expired challenge answer
- Empty answer

---

### 33. Transaction Request Type Validation

**What It Checks**: Transaction request type validity

**Why It Exists**: Ensure only supported transaction types are used

**When It Applies**: Transaction request creation

**Who It Affects**: All users creating transaction requests

**Validation Details**:
- Transaction request type must be registered for the bank
- Type must be enabled and available
- Type cannot change after transaction creation

**What Happens on Failure**:
- Error Code: `OBP-40001` (Invalid type)
- Error Code: `OBP-40009` (Type changed)
- Message: "Invalid value for TRANSACTION_REQUEST_TYPE" or "The TRANSACTION_REQUEST_TYPE has changed."

**Valid Examples**:
- `SEPA`
- `COUNTERPARTY`
- `TRANSFER_TO_PHONE`
- `TRANSFER_TO_ATM`
- `TRANSFER_TO_ACCOUNT`

---

### 34. Transaction Charge Policy Validation

**What It Checks**: Who pays transaction fees

**Why It Exists**: Specify fee allocation for international/inter-bank transfers

**When It Applies**: Transaction creation with fees

**Who It Affects**: Users initiating transactions with fees

**Validation Details**:
- Must be one of three valid values: `SHARED`, `SENDER`, or `RECEIVER`
- `SHARED`: Fees split between sender and receiver
- `SENDER`: Sender pays all fees
- `RECEIVER`: Receiver pays all fees

**What Happens on Failure**:
- Error Code: `OBP-40013`
- Message: "Invalid Charge Policy. Please specify a valid value for Charge_Policy: SHARED, SENDER or RECEIVER."

**Example Valid Input**:
- `SHARED`
- `SENDER`
- `RECEIVER`

**Example Invalid Input**:
- `BOTH`
- `NONE`
- `sender` (lowercase)

---

### 35. Transaction Refund Validation

**What It Checks**: Whether a transaction can be refunded

**Why It Exists**: Prevent duplicate refunds

**When It Applies**: Refund operations

**Who It Affects**: Users requesting refunds

**Validation Details**:
- Transaction must exist
- Transaction must not have been previously refunded
- Original transaction must be completed

**What Happens on Failure**:
- Error Code: `OBP-30068`
- Message: "Transaction was already refunded. Please specify a valid value for TRANSACTION_ID."

---

### 36. Transaction Attributes Validation

**What It Checks**: Custom transaction attribute values

**Why It Exists**: Allow banks to store additional transaction metadata

**When It Applies**: Transaction attribute operations

**Who It Affects**: Banks managing transaction metadata

**Validation Details**:
- Attribute must be defined
- Attribute ID must be valid
- Attribute values must match defined format

**What Happens on Failure**:
- Error Code: `OBP-30070`
- Message: "Transaction Attribute not found. Please specify a valid value for TRANSACTION_ATTRIBUTE_ID."

---

### 37. Counterparty Beneficiary Validation

**What It Checks**: Whether counterparty can receive payments

**Why It Exists**: Prevent payments to unauthorized or unverified counterparties

**When It Applies**: Payment to counterparty

**Who It Affects**: Users sending money to counterparties

**Validation Details**:
- Counterparty must have `isBeneficiary` set to `true`
- Counterparty must be verified/approved for receiving funds

**What Happens on Failure**:
- Error Code: `OBP-30013`
- Message: "The account can not send money to the Counterparty. Please set the Counterparty 'isBeneficiary' true first"

---

### 38. Insufficient Authorization for Transaction Request

**What It Checks**: User/consumer permissions to create transaction requests

**Why It Exists**: Ensure only authorized users can initiate transactions

**When It Applies**: Transaction request creation

**Who It Affects**: All users creating transaction requests

**Validation Details**:
- User must have access to the from-account view
- Consumer must have access to the from-account view
- User must have `CanCreateAnyTransactionRequest` role OR
- View must have `can_add_transaction_request_to_any_account` permission OR
- View must have `can_add_transaction_request_to_beneficiary` permission

**What Happens on Failure**:
- Error Code: `OBP-40002`
- Message: "Insufficient authorisation to create TransactionRequest. The Transaction Request could not be created because the login user doesn't have access to the view of the from account or the consumer doesn't have the access to the view of the from account or the login user does not have the CanCreateAnyTransactionRequest role or the view does not have the permission can_add_transaction_request_to_any_account or the view does not have the permission can_add_transaction_request_to_beneficiary."

---

### 39. Transaction Request Cancellation Validation

**What It Checks**: Whether transaction can be cancelled

**Why It Exists**: Prevent cancellation of completed or invalid transactions

**When It Applies**: Transaction cancellation attempts

**Who It Affects**: Users cancelling transactions

**Validation Details**:
- Transaction status must allow cancellation
- Cancellation authorization process must be valid
- Status must be appropriate for cancellation (INITIATED, CANCELLATION_PENDING, or COMPLETED)

**What Happens on Failure**:
- Error Code: `OBP-40023` (Cannot be cancelled)
- Error Code: `OBP-40025` (Cannot update PSU data for cancellation)
- Error Code: `OBP-40031` (Cannot start authorization for cancellation)
- Message: "Transaction Request cannot be cancelled." or "Cannot Update PSU Data for payment initiation cancellation..."

---

## Data Format Validations

### 40. Currency Code Format Validation

**What It Checks**: ISO 4217 currency code format

**Why It Exists**: Ensure standardized currency representation across the system

**When It Applies**: All operations involving currency

**Who It Affects**: All users working with financial data

**Validation Details**:
- Must be 3-letter ISO 4217 code
- Special extensions: XBT (Bitcoin), lovelace (Cardano), ETH (Ethereum)
- Validated against ISOCurrencyCodes.xml file
- Case-sensitive (uppercase)

**What Happens on Failure**:
- Generic currency validation error

**Example Valid Input**:
- `USD`, `EUR`, `GBP`, `JPY`, `CHF`, `AUD`, `CAD`
- `XBT` (Bitcoin)
- `lovelace` (Cardano)
- `ETH` (Ethereum)

**Example Invalid Input**:
- `US` (too short)
- `EURO` (too long)
- `usd` (lowercase)
- `XXX` (invalid code)

---

### 41. Phone Number Format Validation

**What It Checks**: International phone number format

**Why It Exists**: Ensure phone numbers can be used for SMS and international calling

**When It Applies**: Customer creation/update, SMS authentication

**Who It Affects**: All users providing phone numbers

**Validation Details**:
- Must start with `+` (plus sign)
- Maximum 15 digits after the `+`
- Only numeric digits allowed after `+`
- No spaces, dashes, or parentheses allowed

**What Happens on Failure**:
- Error Code: `OBP-40017`
- Message: "Invalid Phone Number. Please specify a valid value for PHONE_NUMBER. Eg:+9722398746"

**Example Valid Input**:
- `+441234567890` (UK)
- `+12125551234` (US)
- `+81312345678` (Japan)

**Example Invalid Input**:
- `441234567890` (missing +)
- `+44 1234 567890` (contains space)
- `+44-1234-567890` (contains dashes)
- `+123456789012345678` (too long)

---

### 42. Date Format Validation

**What It Checks**: ISO date format compliance

**Why It Exists**: Ensure consistent date handling across the API

**When It Applies**: All date fields in API requests

**Who It Affects**: All API consumers

**Validation Details**:
- Standard format: ISO 8601
- Example: `2024-12-31T23:59:59Z`
- Timezone information may be required
- Date parsing must be lenient or strict based on context

**What Happens on Failure**:
- Error Code: `OBP-10001`
- Message: "Incorrect json format."

**Example Valid Input**:
- `2024-01-15T10:30:00Z`
- `2024-12-31`

**Example Invalid Input**:
- `15-01-2024` (wrong format)
- `2024/01/15` (wrong separator)
- `31-12-2024` (ambiguous format)

---

### 43. UUID Format Validation

**What It Checks**: UUID (Universally Unique Identifier) format

**Why It Exists**: Ensure IDs follow UUID standard for uniqueness

**When It Applies**: Operations with UUID-based identifiers

**Who It Affects**: API consumers using UUID fields

**Validation Details**:
- Must be valid UUID format: `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
- 32 hexadecimal characters with 4 hyphens
- Case-insensitive

**What Happens on Failure**:
- Error Code: `OBP-20253`
- Message: "Invalid format. Must be a UUID."

**Example Valid Input**:
- `550e8400-e29b-41d4-a716-446655440000`
- `123e4567-e89b-12d3-a456-426614174000`

**Example Invalid Input**:
- `123456` (not UUID format)
- `550e8400e29b41d4a716446655440000` (missing hyphens)
- `550e8400-e29b-41d4-a716` (incomplete)

---

### 44. Locale Format Validation

**What It Checks**: Locale code format for internationalization

**Why It Exists**: Support multi-language features with standard locale codes

**When It Applies**: Language/locale preferences

**Who It Affects**: Users setting language preferences

**Validation Details**:
- Currently supported: `en_GB`, `es_ES`, `ro_RO`
- Format: `language_COUNTRY`
- Case-insensitive

**What Happens on Failure**:
- Error Code: `OBP-10002`
- Message: Invalid Locale error

**Example Valid Input**:
- `en_GB` (English - Great Britain)
- `es_ES` (Spanish - Spain)
- `ro_RO` (Romanian - Romania)

**Example Invalid Input**:
- `en` (missing country)
- `en-GB` (wrong separator)
- `fr_FR` (not supported yet)

---

### 45. URL Format Validation

**What It Checks**: URL string format and length

**Why It Exists**: Prevent malformed URLs and injection attacks

**When It Applies**: Webhook URLs, callback URLs, redirects

**Who It Affects**: Developers configuring webhooks and callbacks

**Validation Details**:
- Must be valid URI format
- Maximum length: 2048 characters
- URL decoded value must be valid
- Must follow RFC 3986 standard

**What Happens on Failure**:
- Validation error for malformed URLs

**Example Valid Input**:
- `https://example.com/webhook`
- `https://api.example.com/callback?param=value`

**Example Invalid Input**:
- `not a url`
- `http://` (incomplete)
- (URL longer than 2048 characters)

---

### 46. Medium String Validation

**What It Checks**: General string format for medium-length fields

**Why It Exists**: Standardize string validation across the system

**When It Applies**: Various text fields throughout the API

**Who It Affects**: All API users

**Validation Details**:
- Allowed characters: A-Z, a-z, 0-9, `-`, `_`, `.`, `@`
- Maximum length: 512 characters
- Must have at least one character

**What Happens on Failure**:
- Generic validation error for invalid characters or length

**Example Valid Input**:
- `user.name@example.com`
- `account-123_test`

**Example Invalid Input**:
- `user#name` (invalid character #)
- `user name` (space not allowed)
- (string longer than 512 characters)

---

### 47. Short String Validation

**What It Checks**: Short string format for codes and identifiers

**Why It Exists**: Validate short identifiers and codes

**When It Applies**: Short code fields, identifiers

**Who It Affects**: All API users

**Validation Details**:
- Allowed characters: A-Z, a-z, 0-9, `-`, `_`, `.`
- Maximum length: 16 characters
- Must have at least one character

**What Happens on Failure**:
- Generic validation error

**Example Valid Input**:
- `code-123`
- `ID_456`

**Example Invalid Input**:
- `code@123` (invalid character)
- `this-is-too-long` (>16 chars)

---

### 48. RFC 7231 Date Header Validation

**What It Checks**: HTTP Date header format per RFC 7231

**Why It Exists**: Ensure request date headers follow HTTP standard

**When It Applies**: API requests with Date header

**Who It Affects**: API clients sending Date headers

**Validation Details**:
- Must follow RFC 7231 format
- Example: `Tue, 15 Nov 1994 08:12:31 GMT`

**What Happens on Failure**:
- Error Code: `OBP-20257`
- Message: "Request header Date is not in accordance with RFC 7231"

**Example Valid Input**:
- `Wed, 21 Oct 2015 07:28:00 GMT`

**Example Invalid Input**:
- `2015-10-21 07:28:00` (wrong format)
- `Oct 21 2015` (incomplete)

---

## Rate Limiting Validations

### 49. Consumer Rate Limit Validation

**What It Checks**: API call frequency limits per consumer

**Why It Exists**: Prevent API abuse and ensure fair resource usage

**When It Applies**: Every API request

**Who It Affects**: All API consumers

**Validation Details**:
- Limits can be set per: Second, Minute, Hour, Day, Week, Month
- Limits are per-consumer (identified by consumer key)
- Can be configured globally or per API endpoint
- Can be configured per bank
- Default values from properties file
- Value of -1 means no limit

**Rate Limit Configuration**:
- `perSecondCallLimit`: Calls allowed per second
- `perMinuteCallLimit`: Calls allowed per minute
- `perHourCallLimit`: Calls allowed per hour
- `perDayCallLimit`: Calls allowed per day
- `perWeekCallLimit`: Calls allowed per week
- `perMonthCallLimit`: Calls allowed per month

**What Happens on Failure**:
- HTTP 429 Too Many Requests response
- Rate limit headers returned in response

**Example Configuration**:
- Per second: `10`
- Per minute: `100`
- Per hour: `1000`
- Per day: `10000`

---

### 50. Anonymous Rate Limit Validation

**What It Checks**: API call limits for unauthenticated requests

**Why It Exists**: Protect public endpoints from abuse

**When It Applies**: Anonymous/unauthenticated API requests

**Who It Affects**: Unauthenticated API users

**Validation Details**:
- Separate limits for anonymous users
- Typically more restrictive than authenticated limits
- May be IP-based

**What Happens on Failure**:
- HTTP 429 Too Many Requests response

---

### 51. Endpoint-Specific Rate Limits

**What It Checks**: Rate limits specific to certain API endpoints

**Why It Exists**: Allow fine-grained control over sensitive or resource-intensive endpoints

**When It Applies**: Requests to rate-limited endpoints

**Who It Affects**: All users of rate-limited endpoints

**Validation Details**:
- Can override consumer-level limits
- Specified by API version and endpoint name
- Time-bound (from_date to to_date)

---

## View Permissions & Entitlements

### 52. View Access Permissions

**What It Checks**: Granular permissions for viewing account and transaction data

**Why It Exists**: Control what information users can see about accounts and transactions

**When It Applies**: All account and transaction data access

**Who It Affects**: All users accessing account information

**Validation Details**:

Views control access through 50+ granular permissions including:

**Transaction Visibility Permissions**:
- `can_see_transaction_this_bank_account`: View transaction data
- `can_see_transaction_other_bank_account`: View other party account info
- `can_see_transaction_metadata`: View transaction metadata
- `can_see_transaction_description`: View transaction descriptions
- `can_see_transaction_amount`: View transaction amounts
- `can_see_transaction_type`: View transaction types
- `can_see_transaction_currency`: View transaction currency
- `can_see_transaction_start_date`: View transaction start date
- `can_see_transaction_finish_date`: View transaction completion date
- `can_see_transaction_balance`: View balance after transaction

**Account Information Permissions**:
- `can_see_bank_account_balance`: View account balance
- `can_see_bank_account_currency`: View account currency
- `can_see_bank_account_label`: View account label/name
- `can_see_bank_account_number`: View account number
- `can_see_bank_account_owners`: View account owners
- `can_see_bank_account_type`: View account type
- `can_see_bank_account_iban`: View IBAN
- `can_see_bank_account_routing`: View routing information

**Other Party Information**:
- `can_see_other_account_national_identifier`: View counterparty identifiers
- `can_see_other_account_swift_bic`: View SWIFT/BIC codes
- `can_see_other_account_iban`: View counterparty IBAN
- `can_see_other_account_bank_name`: View counterparty bank
- `can_see_other_account_kind`: View account type
- `can_see_other_account_metadata`: View metadata

**Metadata and Comments**:
- `can_see_comments`: View comments on transactions
- `can_add_comment`: Add comments
- `can_delete_comment`: Delete comments (own or others)
- `can_see_owner_comment`: View owner comments
- `can_edit_owner_comment`: Edit owner comments

**Tags and Images**:
- `can_see_tags`: View transaction tags
- `can_add_tag`: Add tags
- `can_delete_tag`: Delete tags
- `can_see_images`: View transaction images
- `can_add_image`: Add images
- `can_delete_image`: Delete images

**Transaction Request Permissions**:
- `can_add_transaction_request_to_own_account`: Create transaction requests
- `can_add_transaction_request_to_any_account`: Create requests for any account
- `can_add_transaction_request_to_beneficiary`: Create requests to beneficiaries
- `can_see_transaction_requests`: View transaction requests

**Counterparty Permissions**:
- `can_see_available_views_for_bank_account`: View available views
- `can_create_counterparty`: Create counterparties
- `can_delete_counterparty`: Delete counterparties

**What Happens on Failure**:
- Error Code: `OBP-30022`
- Message: "The current view does not have the permission: [PERMISSION_NAME]"

---

### 53. Entitlement (Role) Validation

**What It Checks**: Role-based permissions for API operations

**Why It Exists**: Control who can perform administrative and sensitive operations

**When It Applies**: Protected API endpoints requiring specific roles

**Who It Affects**: Users performing administrative operations

**Validation Details**:

Entitlements are roles that grant permissions for specific operations. Key roles include:

**Account Management**:
- `CanCreateAccount`
- `CanGetAnyUser`
- `CanCreateAnyTransactionRequest`

**Bank Management**:
- `CanCreateBank`
- `CanCreateBranch`
- `CanDeleteBranch`

**User Management**:
- `CanCreateUser`
- `CanDeleteUser`
- `CanLockUser`
- `CanUnlockUser`

**Entitlement Management**:
- `CanCreateEntitlementAtOneBank`: Grant roles at a specific bank
- `CanCreateEntitlementAtAnyBank`: Grant roles at any bank

**System Administration**:
- `CanReadMetrics`
- `CanGetConfig`
- `CanCreateSystemLevelEndpointTag`

**Special Roles**:
- Entitlements can be bank-specific or system-wide
- Bank-specific roles require a valid BANK_ID
- System roles require empty BANK_ID

**What Happens on Failure**:
- Error Code: `OBP-20006` (Insufficient entitlement)
- Error Code: `OBP-30212` (Entitlement not found)
- Error Code: `OBP-30213` (User doesn't have entitlement)
- Error Code: `OBP-30216` (Entitlement already exists)
- Message: "User is missing one or more entitlements:"

---

### 54. Entitlement Grantor Permission

**What It Checks**: Whether user granting entitlement has permission to do so

**Why It Exists**: Prevent unauthorized privilege escalation

**When It Applies**: Entitlement creation/granting

**Who It Affects**: Users granting roles to others

**Validation Details**:
- Grantor must have `CanCreateEntitlementAtAnyBank` for any bank OR
- Grantor must have `CanCreateEntitlementAtOneBank` for specific bank
- Cannot grant entitlements they don't have authority for

**What Happens on Failure**:
- Error Code: `OBP-30221`
- Message: "Entitlement cannot be granted due to the grantor's insufficient privileges."

---

### 55. View Grant/Revoke Access Validation

**What It Checks**: Permission to grant or revoke view access to other users

**Why It Exists**: Control who can manage view access permissions

**When It Applies**: Granting/revoking view access

**Who It Affects**: Account owners and administrators

**Validation Details**:

**For Granting Access**:
- Current view must have `can_grant_access_to_views` (for system views)
- Current view must have `can_grant_access_to_custom_views` (for custom views)
- Source view must contain target view in allowed list

**For Revoking Access**:
- Current view must have `can_revoke_access_to_views` (for system views)
- Current view must have `can_revoke_access_to_custom_views` (for custom views)

**What Happens on Failure**:
- Error Code: `OBP-20047` (Cannot grant - system view)
- Error Code: `OBP-20048` (Cannot revoke - system view)
- Error Code: `OBP-20084` (Cannot grant - system view alternate)
- Error Code: `OBP-20085` (Cannot grant - custom view)
- Error Code: `OBP-20086` (Cannot revoke - system view alternate)
- Error Code: `OBP-20087` (Cannot revoke - custom view)
- Message: Various permission-specific messages

---

### 56. Source View Permission Validation

**What It Checks**: Source view has sufficient permissions compared to target view

**Why It Exists**: Prevent granting more permissions than the grantor has

**When It Applies**: View access management

**Who It Affects**: Users managing view access

**Validation Details**:
- Source view must have at least the permissions of target view
- Cannot grant permissions you don't have

**What Happens on Failure**:
- Error Code: `OBP-20049`
- Message: "Source view contains less permissions than target view."

---

## KYC & Compliance Validations

### 57. KYC Status Validation

**What It Checks**: Customer Know Your Customer (KYC) verification status

**Why It Exists**: Regulatory compliance for customer identity verification

**When It Applies**: Customer operations requiring KYC

**Who It Affects**: All bank customers

**Validation Details**:
- KYC status is boolean field (true/false)
- True indicates customer has passed KYC verification
- False indicates customer has not completed KYC or failed verification
- May restrict certain operations if KYC not complete

**What Happens on Failure**:
- Operations may be blocked for customers without KYC

---

### 58. KYC Check Validation

**What It Checks**: KYC verification checks and documents

**Why It Exists**: Maintain audit trail of KYC verification process

**When It Applies**: KYC check operations

**Who It Affects**: Compliance officers, customers

**Validation Details**:
- KYC checks must be associated with valid customer
- Check date must be recorded
- Check method must be specified
- Staff performing check must be identified
- Satisfaction status (pass/fail) must be recorded
- Comments may be required

**What Happens on Failure**:
- Cannot retrieve or create KYC checks with invalid data

---

### 59. KYC Document Validation

**What It Checks**: KYC supporting documents

**Why It Exists**: Store and verify identity documents

**When It Applies**: KYC document upload and verification

**Who It Affects**: Customers submitting documents, compliance officers

**Validation Details**:
- Documents must be associated with customer
- Document type must be specified
- Document number must be recorded
- Issue and expiry dates may be required

**What Happens on Failure**:
- Cannot process KYC without required documents

---

### 60. Credit Rating Validation

**What It Checks**: Customer credit rating information

**Why It Exists**: Assess customer creditworthiness for lending decisions

**When It Applies**: Credit-related operations

**Who It Affects**: Customers seeking credit

**Validation Details**:
- Credit rating must have rating value
- Credit source must be specified
- Rating must be from recognized agency

---

### 61. Credit Limit Validation

**What It Checks**: Customer credit limit amounts

**Why It Exists**: Enforce lending limits per customer

**When It Applies**: Credit operations, loan applications

**Who It Affects**: Customers with credit facilities

**Validation Details**:
- Credit limit must have currency
- Credit limit must have amount
- Amount must be positive

---

### 62. Consent Validation

**What It Checks**: User consent for data access and operations

**Why It Exists**: GDPR and privacy compliance, PSD2 requirements

**When It Applies**: Consent-based operations

**Who It Affects**: All users, third-party providers

**Validation Details**:

**Consent Status**:
- Must be in valid status (INITIATED, AUTHORISED, REJECTED, REVOKED, etc.)
- Cannot be expired
- Not Before time (nbf) must be in past
- Must match valid consumer
- Must match valid user

**Consent Scope**:
- Can only contain roles user already has
- Can only contain views user already has access to
- Cannot contain restricted roles (CanCreateEntitlementAtOneBank, CanCreateEntitlementAtAnyBank)

**Consent Headers**:
- Consent-ID must be UUID or JWT format
- Consumer-Key header required for consumer validation
- Consent must not be revoked

**Berlin Group Consent Rules**:
- Access must be requested (not empty)
- Recurring indicator must be false when using availableAccounts
- Frequency per day must be 1 when using availableAccounts
- availableAccounts must be exactly 'allAccounts'

**Consent TTL**:
- Must not exceed maximum time to live

**What Happens on Failure**:
- Error Code: `OBP-35001` through `OBP-35034` (various consent errors)
- Multiple specific error messages based on violation type

---

## Business Logic Validations

### 63. Berlin Group Date Range Validation

**What It Checks**: Date range for Berlin Group API endpoints

**Why It Exists**: Berlin Group (PSD2) standard compliance

**When It Applies**: Berlin Group API requests with date parameters

**Who It Affects**: PSD2 third-party providers

**Validation Details**:
- Date must be in ISO format
- Future dates maximum 180 days from now
- Historical dates have bank-specific limits

**What Happens on Failure**:
- Error for dates exceeding limits

**Example Valid Input**:
- `2024-01-15` (within 180 days)

**Example Invalid Input**:
- `2025-12-31` (more than 180 days ahead)

---

### 64. Berlin Group Frequency Per Day Validation

**What It Checks**: Access frequency limits per day

**Why It Exists**: Berlin Group (PSD2) access control

**When It Applies**: Berlin Group consent creation

**Who It Affects**: PSD2 TPPs

**Validation Details**:
- Frequency must be > 0 and <= configured upper limit (default 4)
- Must be exactly 1 for one-off access
- Must be 1 when using availableAccounts

**What Happens on Failure**:
- Error Code: `OBP-20062` (Invalid frequency)
- Error Code: `OBP-20063` (Must be 1 for one-off)
- Error Code: `OBP-20090` (Must be 1 for availableAccounts)
- Message: "Frequency per day must be greater than 0 and less or equal to [LIMIT]"

---

### 65. User-Customer Link Business Rules

**What It Checks**: Business rules for linking users to customers

**Why It Exists**: Ensure proper customer-user relationships

**When It Applies**: User-customer linking operations

**Who It Affects**: Users and customers

**Validation Details**:
- User can only be linked to one customer per bank
- Customer can be linked to multiple users
- Link must exist for customer operations requiring user context

**What Happens on Failure**:
- Error Code: `OBP-30007` (User already linked to customer at bank)
- Message: "The User is already linked to a Customer at the bank specified by BANK_ID"

---

### 66. Account Application Workflow Validation

**What It Checks**: Account application lifecycle status

**Why It Exists**: Manage account opening workflow

**When It Applies**: Account application operations

**Who It Affects**: Users applying for accounts

**Validation Details**:
- Application must exist to be processed
- Application can only be accepted once
- Status transitions must follow valid workflow
- User ID or Customer ID must be present

**What Happens on Failure**:
- Error Code: `OBP-30311` (Application not found)
- Error Code: `OBP-30313` (User ID and Customer ID not present)
- Error Code: `OBP-30314` (Already accepted)
- Error Code: `OBP-30315` (Cannot update status)
- Error Code: `OBP-30316` (Cannot create)

---

### 67. Counterparty Routing Validation

**What It Checks**: Counterparty must have OBP routing for certain operations

**Why It Exists**: Ensure proper counterparty identification in OBP system

**When It Applies**: Counterparty-based transactions

**Who It Affects**: Users working with counterparties

**Validation Details**:
- otherAccountRoutingScheme must be 'OBP'
- otherBankRoutingScheme must be 'OBP'

**What Happens on Failure**:
- Error Code: `OBP-40012`
- Message: "Please set up the otherAccountRoutingScheme and otherBankRoutingScheme fields of the Counterparty to 'OBP'"

---

### 68. Webhook Validation

**What It Checks**: Webhook configuration and operations

**Why It Exists**: Enable event-driven integrations

**When It Applies**: Webhook creation, updates, and triggers

**Who It Affects**: Developers setting up webhooks

**Validation Details**:
- Webhook URL must be valid
- Webhook must exist for operations
- Account must be specified for account-specific webhooks

**What Happens on Failure**:
- Error Code: `OBP-30047` (Cannot create)
- Error Code: `OBP-30048` (Cannot get)
- Error Code: `OBP-30049` (Cannot update)
- Error Code: `OBP-30050` (Not found)

---

### 69. Counterparty Limit Validation

**What It Checks**: Transaction limits per counterparty

**Why It Exists**: Control spending to specific counterparties

**When It Applies**: Transactions to counterparties with limits

**Who It Affects**: Users with counterparty limits configured

**Validation Details**:
- Limit must exist and be valid
- Transaction amount must not exceed limit
- Limit applies to specific bank/account/view/counterparty combination

**What Happens on Failure**:
- Error Code: `OBP-30263` (Limit not found)
- Error Code: `OBP-30264` (Limit already exists)
- Error Code: `OBP-30265` (Cannot delete)
- Error Code: `OBP-30268` (Validation error)

---

### 70. Agent Account Link Validation

**What It Checks**: Link between agents and accounts

**Why It Exists**: Enable agent banking operations

**When It Applies**: Agent account operations

**Who It Affects**: Banking agents

**Validation Details**:
- Agent must exist
- Agent number must be unique per bank
- Agent must be confirmed (is_confirmed_agent = true)
- Agent must not be pending (is_pending_agent = false)
- Account-agent link must be valid

**What Happens on Failure**:
- Error Code: `OBP-30201` (Agent not found)
- Error Code: `OBP-30325` (Link not found)
- Error Code: `OBP-30328` (Agent number already exists)
- Error Code: `OBP-30330` (Agent not confirmed)

---

### 71. Product Hierarchy Validation

**What It Checks**: Parent-child relationships in product catalog

**Why It Exists**: Support product variants and hierarchies

**When It Applies**: Product creation with parent products

**Who It Affects**: Bank administrators managing products

**Validation Details**:
- Parent product must exist if specified
- Parent product code must be valid
- Can leave empty if no parent

**What Happens on Failure**:
- Error Code: `OBP-30062`
- Message: "Parent product not found. Please specify an existing product code for parent_product_code. Leave empty if no parent product exists."

---

### 72. Card Validation

**What It Checks**: Payment card information and operations

**Why It Exists**: Manage payment cards linked to accounts

**When It Applies**: Card operations

**Who It Affects**: Card holders

**Validation Details**:
- Card must exist for user
- Card number must be valid
- Combination of bank ID, card number, and issue number must be unique

**What Happens on Failure**:
- Error Code: `OBP-30059` (Card not found)
- Error Code: `OBP-30060` (Card already exists)
- Error Code: `OBP-30200` (Invalid card number)

---

### 73. Branch and ATM License Validation

**What It Checks**: License information for branches and ATMs

**Why It Exists**: Regulatory compliance for location data

**When It Applies**: Branch/ATM listing and access

**Who It Affects**: Users accessing branch/ATM data

**Validation Details**:
- License ID (meta.license.id) must not be empty
- License name (meta.license.name) must not be empty

**What Happens on Failure**:
- Error Code: `OBP-300010` (Branch)
- Error Code: `OBP-32001` (Branches not found - license issue)
- Error Code: `OBP-33001` (ATMs not found - license issue)
- Message: "Branch not found. Please specify a valid value for BRANCH_ID. Or License may not be set. meta.license.id and meta.license.name can not be empty"

---

### 74. Regulated Entity Validation

**What It Checks**: Regulated entity (TPP) registration and certificates

**Why It Exists**: PSD2/Open Banking compliance

**When It Applies**: Regulated entity operations

**Who It Affects**: Third-party providers (TPPs)

**Validation Details**:
- Regulated entity must exist
- Certificate must uniquely identify entity
- JWT in post data must be verifiable

**What Happens on Failure**:
- Error Code: `OBP-34100` (Entity not found)
- Error Code: `OBP-34101` (Cannot delete)
- Error Code: `OBP-34102` (Not found by certificate)
- Error Code: `OBP-34103` (Multiple entities found by certificate)
- Error Code: `OBP-34110` (JWT not signed properly)

---

### 75. Checkbook Order Validation

**What It Checks**: Checkbook ordering operations

**Why It Exists**: Support checkbook services

**When It Applies**: Checkbook order operations

**Who It Affects**: Account holders ordering checkbooks

**Validation Details**:
- Checkbook order must exist for account

**What Happens on Failure**:
- Error Code: `OBP-30041`
- Message: "CheckbookOrder not found for Account."

---

## API-Specific Validations

### 76. JSON Schema Validation

**What It Checks**: Request body JSON structure against predefined schemas

**Why It Exists**: Ensure API requests contain correctly structured data

**When It Applies**: API endpoints with JSON schema validation enabled

**Who It Affects**: All API consumers

**Validation Details**:
- Each operation can have a JSON schema
- Request body validated against schema before processing
- Schema defines required fields, data types, formats, and constraints
- Schemas are operation-specific (by OPERATION_ID)

**What Happens on Failure**:
- Error Code: `OBP-40026` (Illegal schema format)
- Error Code: `OBP-40027` (Validation not found)
- Message: "Incorrect json-schema Format." or "JSON Schema Validation not found, please specify valid query parameter."

**Example Schema Elements**:
- Required fields
- Data types (string, number, boolean, object, array)
- String patterns (regex)
- Number ranges (minimum, maximum)
- String length constraints
- Enum values

---

### 77. Mandatory Berlin Group Headers Validation

**What It Checks**: Required headers for Berlin Group (PSD2) API calls

**Why It Exists**: Berlin Group specification compliance

**When It Applies**: Berlin Group API endpoints

**Who It Affects**: PSD2 third-party providers

**Validation Details**:

**Required Headers**:
- `X-Request-ID`: Unique request identifier (UUID)
- `Consent-ID`: Consent identifier (for consent-based operations)
- `PSU-ID`: Payment Service User ID (when applicable)
- `TPP-Signature-Certificate`: TPP certificate (when applicable)
- `Signature`: Request signature (when applicable)
- `Date`: Request date (RFC 7231 format)

**Additional Rules**:
- Headers must not be empty or null
- X-Request-ID must be unique (not reused)
- X-Request-ID must be valid UUID format
- Consent-ID must not be used for non-consent endpoints
- Signature header must be valid format

**What Happens on Failure**:
- Error Code: `OBP-20251` (Missing mandatory headers)
- Error Code: `OBP-20252` (Empty/null headers)
- Error Code: `OBP-20253` (Invalid UUID)
- Error Code: `OBP-20254` (Invalid signature header)
- Error Code: `OBP-20255` (Request ID already used)
- Error Code: `OBP-20256` (Consent-ID misuse)
- Error Code: `OBP-20257` (Invalid date format)

---

### 78. Endpoint Mapping Validation

**What It Checks**: Dynamic endpoint mapping configuration

**Why It Exists**: Support dynamic endpoint routing

**When It Applies**: Dynamic endpoint operations

**Who It Affects**: System administrators

**Validation Details**:
- Endpoint mapping must exist
- Operation ID must be valid
- Mapping configuration must be valid

**What Happens on Failure**:
- Error Code: `OBP-36004` (Not found by mapping ID)
- Error Code: `OBP-36005` (Not found by operation ID)
- Error Code: `OBP-36006` (Invalid mapping)

---

### 79. Dynamic Entity Validation

**What It Checks**: Dynamic entity definitions and operations

**Why It Exists**: Enable custom banking entities without code changes

**When It Applies**: Dynamic entity CRUD operations

**Who It Affects**: System administrators creating custom entities

**Validation Details**:
- Entity definition must be valid
- Entity name must be unique
- Fields must have valid types
- Operations must have proper permissions

---

### 80. Connector Method Validation

**What It Checks**: Dynamic connector method definitions

**Why It Exists**: Enable custom connector logic

**When It Applies**: Connector method operations

**Who It Affects**: System administrators

**Validation Details**:
- Connector method must exist
- Method body must be valid Scala code
- Method must compile successfully
- Method ID must be unique

**What Happens on Failure**:
- Error Code: `OBP-40036` (Method not found)
- Error Code: `OBP-40037` (Method already exists)
- Error Code: `OBP-40038` (Compilation failed)

---

### 81. One Time Password (OTP) Validation

**What It Checks**: OTP codes for transaction confirmation

**Why It Exists**: Strong Customer Authentication (SCA)

**When It Applies**: OTP-based authentication

**Who It Affects**: Users using OTP authentication

**Validation Details**:
- OTP must not be expired
- OTP must match expected value
- OTP can only be used once

**What Happens on Failure**:
- Error Code: `OBP-20211` (OTP expired)
- Error Code: `OBP-20216` (OTP invalid)
- Message: "The One Time Password (OTP) has expired." or "The One Time Password (OTP) is invalid."

---

### 82. User Auth Context Validation

**What It Checks**: User authentication context for onboarding

**Why It Exists**: Enable secure customer onboarding

**When It Applies**: User auth context operations

**Who It Affects**: Customers being onboarded

**Validation Details**:
- Auth context update request key must be valid (e.g., CUSTOMER_NUMBER)
- Customer must exist for the provided identifier
- SCA method must be supported (SMS, EMAIL, IMPLICIT)

**What Happens on Failure**:
- Error Code: `OBP-30088` (Invalid key)
- Error Code: `OBP-35034` (Unsupported SCA method)
- Message: "Invalid Auth Context Update Request Key."

---

### 83. Elasticsearch Validation

**What It Checks**: Elasticsearch operations and results

**Why It Exists**: Support advanced search and analytics

**When It Applies**: Elasticsearch-based searches

**Who It Affects**: Users performing advanced searches

**Validation Details**:
- Elasticsearch must be enabled
- Index must exist
- Query body must not be empty
- Result set must meet minimum size for privacy

**What Happens on Failure**:
- Error Code: `OBP-20051` (Index not found)
- Error Code: `OBP-20052` (Insufficient results)
- Error Code: `OBP-20053` (Empty query)
- Error Code: `OBP-20056` (Elasticsearch disabled)

---

### 84. API Collection Validation

**What It Checks**: API endpoint collections

**Why It Exists**: Group related endpoints for management

**When It Applies**: API collection operations

**Who It Affects**: API administrators

**Validation Details**:
- Collection must exist
- Collection name must be unique
- Endpoints in collection must be valid

**What Happens on Failure**:
- Error Code: `OBP-30079` (Not found)
- Error Code: `OBP-30080` (Cannot create)
- Error Code: `OBP-3008A` (Cannot update)
- Error Code: `OBP-30081` (Cannot delete)
- Error Code: `OBP-30086` (Already exists)

---

### 85. Endpoint Tag Validation

**What It Checks**: Tags for API endpoints

**Why It Exists**: Enable endpoint categorization and filtering

**When It Applies**: Endpoint tag operations

**Who It Affects**: API administrators

**Validation Details**:
- Tag must exist
- Tag ID must be valid
- Tag combination must be unique

**What Happens on Failure**:
- Error Code: `OBP-30096` (Cannot create)
- Error Code: `OBP-30097` (Cannot update)
- Error Code: `OBP-30098` (Unknown error)
- Error Code: `OBP-30099` (Tag not found)
- Error Code: `OBP-30100` (Already exists)

---

### 86. Meeting Provider Validation

**What It Checks**: Meeting provider integration configuration

**Why It Exists**: Support scheduling meetings with bank staff

**When It Applies**: Meeting operations

**Who It Affects**: Customers scheduling meetings

**Validation Details**:
- Meetings must be supported
- API key must be configured
- API secret must be configured
- Meeting must exist for operations

**What Happens on Failure**:
- Error Code: `OBP-30101` (Not supported)
- Error Code: `OBP-30102` (API key not configured)
- Error Code: `OBP-30103` (API secret not configured)
- Error Code: `OBP-30104` (Meeting not found)

---

## 11. Dynamic Entity Validations

Dynamic entities allow runtime creation of custom banking entities with associated CRUD operations. Each dynamic entity has field definitions that specify validation rules.

### 11.1 Dynamic Entity Required Fields

**What it checks:** All fields marked as 'required' in the entity definition are present in requests

**Why it exists:** To ensure data completeness for dynamically-defined entities

**When it applies:** When creating or updating dynamic entity instances

**Who it affects:** API users working with custom dynamic entities

**Validation Details:**
- System checks entity definition for required fields
- Each required field must be present in the request JSON
- Empty values may not satisfy the requirement depending on field type

**What happens when it fails:**
Returns error listing missing required fields: "error: the following required fields are missing: [field1, field2, ...]"

**Examples:**
- ✅ Valid: Request includes all required fields
- ❌ Invalid: Request missing one or more required fields

---

### 11.2 Dynamic Entity Field Type Validation

**What it checks:** Field values match their declared type in the entity definition

**Why it exists:** To maintain type safety in dynamically-defined data structures

**When it applies:** When creating or updating dynamic entity instances

**Who it affects:** API users working with custom dynamic entities

**Validation Details:**
- Each field has a defined type (string, number, boolean, date, reference, etc.)
- Values must be convertible to the declared type
- Type checking is strict

**What happens when it fails:**
Returns error indicating type mismatch: "Field [field_name] should be [expected_type], current value is [actual_value]"

**Examples:**
- ✅ Valid: String field receives "text", number field receives 123
- ❌ Invalid: Number field receives "abc", boolean field receives "maybe"

---

### 11.3 Dynamic Entity String minLength

**What it checks:** String field values meet minimum length requirements

**Why it exists:** To ensure string fields contain meaningful data

**When it applies:** When creating or updating dynamic entity instances with string fields that have minLength constraints

**Who it affects:** API users working with custom dynamic entities

**Validation Details:**
- String fields can specify a `minLength` property
- Actual string length must be >= minLength
- Length is measured in characters

**What happens when it fails:**
Returns error: "Field [field_name] minimum length is [X], current length is [Y]"

**Examples:**
- If minLength=3: ✅ Valid: `"abc"`, `"test"` | ❌ Invalid: `"ab"`, `"x"`

---

### 11.4 Dynamic Entity String maxLength

**What it checks:** String field values do not exceed maximum length limits

**Why it exists:** To prevent overly long strings that could cause storage or display issues

**When it applies:** When creating or updating dynamic entity instances with string fields that have maxLength constraints

**Who it affects:** API users working with custom dynamic entities

**Validation Details:**
- String fields can specify a `maxLength` property
- Actual string length must be <= maxLength
- Length is measured in characters

**What happens when it fails:**
Returns error: "Field [field_name] maximum length is [X], current length is [Y]"

**Examples:**
- If maxLength=50: ✅ Valid: any string ≤50 chars | ❌ Invalid: 51+ character strings

---

### 11.5 Dynamic Entity Reference Type Validation

**What it checks:** Reference fields point to existing entities of the correct type

**Why it exists:** To maintain referential integrity in dynamic entity relationships

**When it applies:** When creating or updating dynamic entity instances with reference fields

**Who it affects:** API users working with custom dynamic entities that have relationships

**Validation Details:**
- Reference fields use format: `reference:EntityType` or `reference:EntityType:fieldName`
- System validates that referenced entity exists
- Supported reference types include: Account, Bank, Customer, Product, Transaction, etc.
- Can reference using single ID or composite keys (e.g., `bankId&productCode`)

**What happens when it fails:**
Returns error: "Can not find the [EntityType] by the [field_name]=[value]"

**Examples:**
- ✅ Valid: `reference:Customer` with valid customer ID
- ✅ Valid: `reference:Product:bankId&productCode` with format `bankId=XXX&productCode=YYY`
- ❌ Invalid: Reference to non-existent entity

---

## 12. Request Header Validations (Berlin Group / PSD2)

### 12.1 X-Request-ID Format Validation

**What it checks:** X-Request-ID header contains a valid UUID

**Why it exists:** To ensure unique request identification and prevent replay attacks

**When it applies:** For Berlin Group API requests that include X-Request-ID header

**Who it affects:** Third-party providers (TPPs) using Berlin Group/PSD2 API

**Validation Details:**
- Must be a valid UUID format (version 4 typically)
- Format: `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` where x is hexadecimal
- Validated using UUID parsing

**What happens when it fails:**
Returns error indicating invalid X-Request-ID format

**Examples:**
- ✅ Valid: `550e8400-e29b-41d4-a716-446655440000`
- ❌ Invalid: `not-a-uuid`, `12345`, `550e8400-invalid`

---

### 12.2 X-Request-ID Uniqueness Validation

**What it checks:** X-Request-ID has not been used in a previous request

**Why it exists:** To prevent replay attacks and ensure idempotency

**When it applies:** For Berlin Group API requests with X-Request-ID header

**Who it affects:** Third-party providers (TPPs) using Berlin Group/PSD2 API

**Validation Details:**
- System maintains record of previously used request IDs
- Each request ID can only be used once
- Check performed before processing request

**What happens when it fails:**
Returns error code: `OBP-40006` "X-Request-ID value has already been used"

**Examples:**
- ✅ Valid: Never-before-used UUID
- ❌ Invalid: UUID that was used in a previous request

---

### 12.3 Date Header Validation

**What it checks:** Date header is present and correctly formatted for signed requests

**Why it exists:** To ensure signature verification can work correctly and prevent timing attacks

**When it applies:** For Berlin Group API requests that require signatures

**Who it affects:** Third-party providers (TPPs) using Berlin Group/PSD2 API with signing

**Validation Details:**
- Date header must be present in signed requests
- Must follow HTTP date format (RFC 7231)
- Used in signature calculation

**What happens when it fails:**
Returns error indicating missing or invalid Date header

**Examples:**
- ✅ Valid: `Tue, 15 Nov 2024 08:12:31 GMT`
- ❌ Invalid: Missing header, wrong format

---

### 12.4 TPP Signature Validation

**What it checks:** TPP signature header is valid and correctly signs the request

**Why it exists:** To authenticate TPPs and ensure request integrity (PSD2 requirement)

**When it applies:** For Berlin Group API requests requiring qualified certificates

**Who it affects:** Third-party providers (TPPs) using Berlin Group/PSD2 API

**Validation Details:**
- Signature must be created using TPP's qualified certificate
- Signature algorithm and format must be correct
- Signature must cover required request elements

**What happens when it fails:**
Returns error indicating invalid signature

**Examples:**
- ✅ Valid: Correctly signed request with valid QSealC certificate
- ❌ Invalid: Missing signature, wrong signature, tampered request

---

### 12.5 TPP Certificate Validation

**What it checks:** TPP certificate (QSealC/QWac) is present and valid

**Why it exists:** PSD2 requires qualified certificates for TPP authentication

**When it applies:** For Berlin Group API requests requiring certificate-based authentication

**Who it affects:** Third-party providers (TPPs) using Berlin Group/PSD2 API

**Validation Details:**
- Certificate must be provided in request headers
- Certificate must be valid (not expired, not revoked)
- Certificate must be qualified (QSealC for signing, QWac for authentication)

**What happens when it fails:**
Returns error indicating missing or invalid certificate

**Examples:**
- ✅ Valid: Valid QSealC/QWac certificate from authorized CA
- ❌ Invalid: Expired certificate, self-signed certificate, wrong certificate type

---

## 13. Attribute System Validations

The OBP API supports custom attributes for various entity types (accounts, products, customers, transactions, cards). Each attribute must conform to its attribute definition.

### 13.1 Attribute Definition Existence

**What it checks:** Attribute is based on a defined attribute definition

**Why it exists:** To ensure only approved attribute types can be created

**When it applies:** When creating custom attributes for any entity

**Who it affects:** Users creating or managing custom attributes

**Validation Details:**
- Each attribute must reference a valid attribute definition ID
- Attribute definition must exist in the system
- Definition determines allowed values and validation rules

**What happens when it fails:**
Returns error indicating attribute definition not found

**Examples:**
- ✅ Valid: Attribute using existing definition ID
- ❌ Invalid: Attribute referencing non-existent definition

---

### 13.2 Attribute Category Validation

**What it checks:** Attribute belongs to a valid category

**Why it exists:** To organize attributes by entity type and ensure proper usage

**When it applies:** When creating attribute definitions

**Who it affects:** Administrators creating attribute definitions

**Validation Details:**
- Valid categories include: Account, Product, Customer, Transaction, Card, ATM, Branch, etc.
- Category must match the entity type the attribute will be applied to
- Category determines where attribute can be used

**What happens when it fails:**
Returns error indicating invalid attribute category

**Examples:**
- ✅ Valid: `AttributeCategory.Account`, `AttributeCategory.Product`
- ❌ Invalid: Non-existent category value

---

### 13.3 Attribute Name Validation

**What it checks:** Attribute name matches the defined schema

**Why it exists:** To maintain consistency in attribute naming

**When it applies:** When creating or updating attributes

**Who it affects:** Users working with custom attributes

**Validation Details:**
- Attribute name must match the name in its definition
- Name format may have specific requirements
- Case sensitivity depends on implementation

**What happens when it fails:**
Returns error indicating attribute name mismatch or invalid format

**Examples:**
- ✅ Valid: Attribute name matching its definition
- ❌ Invalid: Misspelled name, wrong case

---

### 13.4 Attribute Type Validation

**What it checks:** Attribute value matches the defined type

**Why it exists:** To maintain data type consistency

**When it applies:** When creating or updating attribute values

**Who it affects:** Users working with custom attributes

**Validation Details:**
- Attribute type defined in attribute definition (string, number, date, etc.)
- Value must be convertible to the defined type
- Type validation is strict

**What happens when it fails:**
Returns error indicating type mismatch

**Examples:**
- ✅ Valid: String attribute gets string value, number attribute gets numeric value
- ❌ Invalid: Number attribute gets text value

---

### 13.5 Attribute Value Format Validation

**What it checks:** Attribute value conforms to any format constraints

**Why it exists:** To ensure attribute values meet specific business requirements

**When it applies:** When setting or updating attribute values

**Who it affects:** Users working with custom attributes

**Validation Details:**
- Attribute definitions may specify format requirements
- Can include regex patterns, length limits, allowed values
- Validation depends on attribute definition configuration

**What happens when it fails:**
Returns error with specific format requirement that was violated

**Examples:**
- ✅ Valid: Value meeting all format requirements
- ❌ Invalid: Value violating length, pattern, or other constraints

---

## 14. FX Rate Validations

### 14.1 FX From Currency Validation

**What it checks:** Source currency code is valid ISO 4217 code

**Why it exists:** To ensure foreign exchange rates use standard currency codes

**When it applies:** When creating or querying FX rates

**Who it affects:** Users managing exchange rates, API consumers using FX endpoints

**Validation Details:**
- Must be valid 3-letter ISO 4217 currency code
- Validated against standard currency code list
- Case insensitive (converted to uppercase)

**What happens when it fails:**
Returns error: "OBP-10050: Cannot create FX currency" or InvalidISOCurrencyCode error

**Examples:**
- ✅ Valid: `USD`, `EUR`, `GBP`, `JPY`
- ❌ Invalid: `US`, `DOLLAR`, `XXX` (if not in ISO 4217 list)

---

### 14.2 FX To Currency Validation

**What it checks:** Target currency code is valid ISO 4217 code

**Why it exists:** To ensure foreign exchange rates use standard currency codes

**When it applies:** When creating or querying FX rates

**Who it affects:** Users managing exchange rates, API consumers using FX endpoints

**Validation Details:**
- Must be valid 3-letter ISO 4217 currency code
- Validated against standard currency code list
- Case insensitive (converted to uppercase)

**What happens when it fails:**
Returns error: InvalidISOCurrencyCode

**Examples:**
- ✅ Valid: `USD`, `EUR`, `GBP`, `JPY`
- ❌ Invalid: `US`, `EURO`, `YYY` (if not in ISO 4217 list)

---

### 14.3 FX Rate Conversion Value Validation

**What it checks:** Conversion rate is a positive numeric value

**Why it exists:** To ensure FX rates are mathematically valid

**When it applies:** When creating or updating FX rates

**Who it affects:** Users managing exchange rates

**Validation Details:**
- Must be a valid decimal number
- Must be positive (greater than 0)
- Typically expressed as units of target currency per unit of source currency

**What happens when it fails:**
Returns error: "OBP-30038: Could not insert the Fx Rate"

**Examples:**
- ✅ Valid: `1.2`, `0.85`, `110.5`
- ❌ Invalid: `-1.2`, `0`, `abc`

---

### 14.4 FX Effective Date Validation

**What it checks:** FX rate effective date is in valid date format

**Why it exists:** To ensure FX rates have proper temporal validity

**When it applies:** When creating or updating FX rates

**Who it affects:** Users managing exchange rates

**Validation Details:**
- Must be valid date format
- Date determines when the rate becomes active
- Used for historical rate tracking

**What happens when it fails:**
Returns error indicating invalid date format

**Examples:**
- ✅ Valid: Valid date in accepted format
- ❌ Invalid: Malformed date string, invalid date values

---

## 15. Product Fee and Product Collection Validations

### 15.1 Product Fee Amount Validation

**What it checks:** Product fee amount is a valid decimal number

**Why it exists:** To ensure fee amounts are mathematically valid

**When it applies:** When creating or updating product fees

**Who it affects:** Product managers, bank administrators

**Validation Details:**
- Must be convertible to BigDecimal
- Can be positive or zero (depending on business rules)
- Precision depends on currency

**What happens when it fails:**
Returns error: "OBP-30600: Could not create ProductFee"

**Examples:**
- ✅ Valid: `5.00`, `10.50`, `0.00`
- ❌ Invalid: `abc`, invalid number format

---

### 15.2 Product Fee Currency Validation

**What it checks:** Product fee currency is valid ISO 4217 code

**Why it exists:** To ensure fees use standard currency codes

**When it applies:** When creating or updating product fees

**Who it affects:** Product managers, bank administrators

**Validation Details:**
- Must be valid 3-letter ISO 4217 currency code
- Validated against standard currency code list

**What happens when it fails:**
Returns InvalidISOCurrencyCode error

**Examples:**
- ✅ Valid: `USD`, `EUR`, `GBP`
- ❌ Invalid: `US`, `DOLLAR`

---

### 15.3 Product Fee Frequency Validation

**What it checks:** Product fee frequency is from allowed list

**Why it exists:** To ensure consistent fee frequency values

**When it applies:** When creating or updating product fees

**Who it affects:** Product managers, bank administrators

**Validation Details:**
- Must be one of the allowed frequency values
- Common values: Monthly, Quarterly, Annually, OneTime, etc.
- Frequency determines billing schedule

**What happens when it fails:**
Returns error indicating invalid frequency value

**Examples:**
- ✅ Valid: `Monthly`, `Annually`, `OneTime`
- ❌ Invalid: `Every3Months` (if not in allowed list)

---

### 15.4 Product Collection Code Uniqueness

**What it checks:** Product collection code is unique

**Why it exists:** To prevent duplicate product collections

**When it applies:** When creating product collections

**Who it affects:** Product managers organizing products into collections

**Validation Details:**
- Collection code must be unique within the bank
- Case sensitivity depends on implementation

**What happens when it fails:**
Returns error indicating duplicate collection code

**Examples:**
- ✅ Valid: Unique collection code
- ❌ Invalid: Code already in use

---

## 16. Inline Validation Pattern (Helper.booleanToFuture)

Throughout the OBP API codebase, there are 100+ inline validation checks using the `Helper.booleanToFuture` pattern. This pattern converts boolean validation expressions into Future results with specific error messages.

### 16.1 Pattern Overview

**What it checks:** Various inline business rule validations embedded in API endpoints

**Why it exists:** To provide flexible, endpoint-specific validation logic without creating separate validation functions

**When it applies:** Throughout API endpoint implementations

**Who it affects:** All API users, depending on the specific validation

**Validation Pattern:**
```scala
_ <- Helper.booleanToFuture(failMsg = "Error message", cc = callContext) {
  validation_expression_that_returns_boolean
}
```

**Examples of Common Inline Validations:**

1. **Account ID Format:**
   ```scala
   _ <- Helper.booleanToFuture(InvalidAccountIdFormat, cc=callContext) {
     isValidID(accountId.value)
   }
   ```

2. **Amount Positive Check:**
   ```scala
   _ <- Helper.booleanToFuture(s"${NotPositiveAmount} Current input is: '${amount}'", cc=callContext) {
     amount > BigDecimal("0")
   }
   ```

3. **Feature Enabled Check:**
   ```scala
   _ <- Helper.booleanToFuture(s"$DataImportDisabled", 403, callContext) {
     APIUtil.getPropsAsBoolValue("allow_sandbox_data_import", defaultValue = false)
   }
   ```

4. **Consumer Credentials Check:**
   ```scala
   _ <- Helper.booleanToFuture(failMsg = ErrorMessages.InvalidConsumerCredentials, cc = callContext) {
     callContext.map(_.consumer.isDefined == true).isDefined
   }
   ```

---

## 17. Secure Logging Pattern Validations

The OBP API implements comprehensive secure logging that automatically masks sensitive data in logs to prevent security breaches.

### 17.1 Secret Masking

**What it checks:** Detects and masks secret parameters in logs

**Why it exists:** To prevent secret keys from being exposed in log files

**When it applies:** Automatically applied to all log output

**Who it affects:** System security, compliance with data protection regulations

**Validation Details:**
- Regex pattern: `(?i)(secret=)([^,\\s&]+)`
- Matches `secret=value` patterns
- Replaces value with `***`
- Case insensitive

**Examples:**
- Original: `secret=abc123def`
- Logged as: `secret=***`

---

### 17.2 Token Masking

**What it checks:** Detects and masks various token types in logs

**Why it exists:** To prevent authentication tokens from being exposed

**When it applies:** Automatically applied to all log output

**Who it affects:** System security

**Validation Details:**
Multiple patterns for different token types:
- `access_token`: `(?i)(access_token[\"']?\\s*[:=]\\s*[\"']?)([^\"',\\s&]+)`
- `refresh_token`: Similar pattern
- `id_token`: Similar pattern  
- `Authorization: Bearer`: `(?i)(Authorization:\\s*Bearer\\s+)([^\\s,&]+)`

**Examples:**
- Original: `Authorization: Bearer eyJhbGc...`
- Logged as: `Authorization: Bearer ***`

---

### 17.3 Password Masking

**What it checks:** Detects and masks password parameters in logs

**Why it exists:** To prevent passwords from being exposed in log files

**When it applies:** Automatically applied to all log output

**Who it affects:** User security, compliance

**Validation Details:**
- Regex pattern: `(?i)(password[\"']?\\s*[:=]\\s*[\"']?)([^\"',\\s&]+)`
- Matches various password parameter formats
- Case insensitive

**Examples:**
- Original: `password=myP@ssw0rd`
- Logged as: `password=***`

---

### 17.4 API Key Masking

**What it checks:** Detects and masks API keys in logs

**Why it exists:** To prevent API keys from being exposed

**When it applies:** Automatically applied to all log output

**Who it affects:** API security

**Validation Details:**
Multiple patterns:
- `api_key`: `(?i)(api_key[\"']?\\s*[:=]\\s*[\"']?)([^\"',\\s&]+)`
- `key`: Similar pattern
- `private_key`: Similar pattern

**Examples:**
- Original: `api_key=sk_live_abc123`
- Logged as: `api_key=***`

---

### 17.5 Credit Card Number Masking

**What it checks:** Detects and masks credit card numbers in logs

**Why it exists:** To comply with PCI DSS and protect customer financial data

**When it applies:** Automatically applied to all log output

**Who it affects:** Customer privacy, PCI compliance

**Validation Details:**
- Regex pattern: `\\b([0-9]{4})[\\s-]?([0-9]{4})[\\s-]?([0-9]{4})[\\s-]?([0-9]{3,7})\\b`
- Matches 13-19 digit card numbers with optional separators
- Keeps first 4 and last 3-7 digits, masks middle

**Examples:**
- Original: `4532-1234-5678-9010`
- Logged as: `4532-****-****-9010`

---

### 17.6 JDBC URL Password Masking

**What it checks:** Detects and masks passwords in JDBC connection strings

**Why it exists:** To prevent database credentials from being exposed in logs

**When it applies:** Automatically applied to all log output

**Who it affects:** Database security

**Validation Details:**
- Regex pattern: `(?i)(jdbc:[^\\s]+://[^:]+:)([^@\\s]+)(@)`
- Matches JDBC URL format with embedded credentials
- Masks password portion only

**Examples:**
- Original: `jdbc:postgresql://host:5432/db@user:password`
- Logged as: `jdbc:postgresql://host:5432/db@user:***`

---

## Error Code Reference

### Error Code Ranges

| Range | Category | Description |
|-------|----------|-------------|
| OBP-00XXX | Infrastructure | Server errors, data errors, unknown errors |
| OBP-10XXX | JSON/Input Format | JSON parsing, validation, format errors |
| OBP-20XXX | Authentication/Authorization | Login, OAuth, permissions, user access |
| OBP-30XXX | Resources | Accounts, customers, transactions, entities |
| OBP-32XXX | Branches | Branch-related errors |
| OBP-33XXX | ATMs | ATM-related errors |
| OBP-34XXX | Banks | Bank-related errors |
| OBP-35XXX | Consents | Consent-related errors |
| OBP-36XXX | Authorizations | Authorization-related errors |
| OBP-40XXX | Transaction Requests | Transaction request lifecycle errors |
| OBP-50XXX | Internal/Connector | System exceptions, connector errors |
| OBP-6XXXX | Adapter | Adapter exceptions and errors |

### Common Error Codes Quick Reference

| Error Code | Message Summary | Category |
|------------|----------------|----------|
| OBP-10001 | Incorrect JSON format | Format Validation |
| OBP-20005 | Invalid user name | Authentication |
| OBP-20006 | Insufficient entitlement | Authorization |
| OBP-20017 | User is locked | Authentication |
| OBP-20054 | Invalid amount | Transaction Validation |
| OBP-20064 | User is deleted | Authentication |
| OBP-30003 | Account not found | Resource |
| OBP-30006 | Customer number exists | Customer Validation |
| OBP-30022 | Missing view permission | View Permission |
| OBP-30110 | Invalid Account ID format | Account Validation |
| OBP-30111 | Invalid Bank ID format | Account Validation |
| OBP-30207 | Invalid password format | Authentication |
| OBP-40002 | Insufficient authorization for transaction request | Transaction Authorization |
| OBP-40003 | Currency must match account | Transaction Validation |
| OBP-40008 | Amount must be positive | Transaction Validation |
| OBP-40011 | Transaction status not initiated | Transaction Status |
| OBP-40016 | Invalid challenge answer | Challenge Validation |
| OBP-40017 | Invalid phone number | Format Validation |
| OBP-50000 | Unknown error | System Error |

---

## Quick Reference by Audience

### Business Analysts - Key Validations

**Account Requirements:**
- Account must have unique ID and valid format
- Initial balance must be zero
- Account type must be valid
- Bank ID must be valid

**Customer Requirements:**
- Customer number must be unique within bank
- Phone number must be in international format (+[digits], max 15)
- Email must be valid
- Legal name required
- KYC status tracked

**Transaction Rules:**
- Amount must be positive (> 0)
- Currency must match account currency
- Transaction requests require proper authorization
- Challenges may be required for large amounts

**Permission Model:**
- View-based access control (50+ permissions)
- Role-based system administration
- Can grant/revoke access with proper permissions

### Compliance Officers - Key Validations

**Identity Verification:**
- KYC status required for customers
- KYC checks recorded with staff ID and method
- KYC documents stored and validated

**Regulatory Compliance:**
- PSD2/Berlin Group date limits (180 days)
- Consent-based access control
- Regulated entity validation via certificates
- Tax residence tracking

**Anti-Money Laundering:**
- Transaction limits per counterparty
- Counterparty beneficiary status checks
- Credit rating and limit tracking

**Audit Trail:**
- All operations require authentication
- Role assignments tracked by grantor
- Request IDs must be unique

**Data Privacy:**
- Consent TTL limits enforced
- Consent scope validated against user permissions
- User data access controlled by views

### QA Teams - Key Test Scenarios

**Valid Scenarios:**
- Password: `MyP@ssw0rd123` (10+ chars, complex)
- Phone: `+441234567890` (international format)
- Account ID: `savings-account-001` (alphanumeric with dashes)
- Currency: `USD`, `EUR`, `GBP`
- Amount: `100.50` (positive)

**Invalid Scenarios:**
- Password: `password` (too simple)
- Phone: `441234567890` (missing +)
- Account ID: `account@123` (invalid character)
- Currency: `US` (not 3 letters)
- Amount: `0` or `-50.00` (not positive)

**Boundary Tests:**
- Password: 10 chars (minimum complex), 17 chars (minimum simple), 512 chars (maximum)
- Phone: 15 digits after + (maximum)
- Account/Bank ID: 255 chars (maximum)
- String fields: 512 chars (typical maximum)

**Permission Tests:**
- User without view permission cannot see transaction details
- User without entitlement cannot perform admin operations
- Cannot grant entitlement without CanCreateEntitlement role
- Cannot exceed rate limits

### Product Managers - Feature Constraints

**Account Limitations:**
- Accounts start with zero balance only
- Account ID format restricted to alphanumeric + `-_.`
- Account type predefined by bank

**Transaction Constraints:**
- Must use account's currency
- Positive amounts only
- Challenge required for certain transaction types
- Status-based workflow (cannot cancel completed)

**Customer Onboarding:**
- Phone must be international format
- KYC required for full service
- User linked to one customer per bank

**API Usage:**
- Rate limits per consumer
- Can be per second/minute/hour/day/week/month
- JSON schema validation on some endpoints

**View System:**
- Granular permission control (50+ permissions)
- Cannot see data without proper view access
- Custom views supported

### Customer Support - Common Issues

**"Cannot log in"**
- Check: User locked? (OBP-20017)
- Check: User deleted? (OBP-20064)
- Check: Correct username format? (email or alphanumeric)
- Check: Password meets requirements?

**"Invalid phone number"**
- Must start with + (e.g., +441234567890)
- Maximum 15 digits after +
- No spaces or dashes allowed

**"Cannot create account"**
- Initial balance must be 0
- Account ID must be alphanumeric + `-_.`
- Maximum 255 characters
- Check if account ID already exists

**"Transaction failed"**
- Amount must be positive (> 0)
- Currency must match account currency
- Check if user has view permission
- Check if challenge answer is correct/expired

**"Access denied" / "Permission error"**
- User may lack required entitlement
- View may not have required permission
- Check if user is linked to customer at bank

**"Customer number already exists"**
- Each customer number must be unique per bank
- Use different customer number

### System Administrators - Configuration

**Rate Limiting:**
- Configure via consumer properties
- Set per second/minute/hour/day/week/month
- Can set per endpoint and API version
- Value -1 means no limit
- Default from properties file

**Authentication Types:**
- Control per endpoint via authentication type validation
- Supported: OAuth1, OAuth2, DirectLogin, GatewayLogin, Anonymous
- Configure per OPERATION_ID

**JSON Schema:**
- Optional per endpoint
- Define via operation ID
- Validates request body structure

**View Permissions:**
- System views: predefined permissions
- Custom views: configurable permissions
- Over 50 granular permissions available

**Consent Configuration:**
- Set maximum TTL
- Configure SCA methods (SMS, EMAIL, IMPLICIT)
- Berlin Group specific: frequency per day limits

**Elasticsearch:**
- Enable/disable via configuration
- Requires proper index setup
- Privacy threshold for result sets

### Security Teams - Security Validations

**Authentication:**
- Password: Strong requirements (10+ complex or 17+ simple)
- OAuth/OAuth2 token validation
- Certificate-based authentication (X.509)
- OTP for sensitive operations

**Authorization:**
- Role-based access control (entitlements)
- View-based permissions (50+ granular controls)
- Consent-based access for PSD2
- Cannot escalate privileges without proper roles

**Input Validation:**
- All IDs validated against regex patterns
- Maximum length constraints
- Character restrictions prevent injection
- JSON schema validation on endpoints

**API Security:**
- Rate limiting per consumer
- Request ID uniqueness enforced
- Certificate validation for PSD2 TPPs
- Signature verification for signed requests

**Data Protection:**
- Consent scope validated
- User data access controlled by views
- Sensitive data only visible with permissions
- Audit trail for role assignments

---

## Common Validation Scenarios

### Scenario 1: Creating a New Customer

**Required Data:**
- Customer number (unique within bank)
- Legal name
- Mobile number (+[country][number], max 15 digits)
- Email (valid format)
- Date of birth
- KYC status

**Common Failures:**
- Customer number already exists (OBP-30006)
- Invalid phone format (OBP-40017)
- Missing required fields

**Resolution:**
- Use unique customer number
- Format phone as +441234567890
- Provide all required fields

### Scenario 2: Creating a Transaction Request

**Required Data:**
- From account (user must have access)
- To account or counterparty
- Amount (positive)
- Currency (must match from-account)
- Transaction type

**Common Failures:**
- Insufficient authorization (OBP-40002)
- Currency mismatch (OBP-40003)
- Zero or negative amount (OBP-40008)
- Account not found (OBP-30003)

**Resolution:**
- Ensure user has view permission for from-account
- Match currency to from-account currency
- Use positive amount
- Verify account exists

### Scenario 3: Granting View Access

**Required:**
- Current user must have view with grant permission
- Target view must be allowed by source view
- Target user must exist

**Common Failures:**
- Lack grant permission (OBP-20047, OBP-20084)
- Source view has less permissions than target (OBP-20049)

**Resolution:**
- Ensure current view has can_grant_access_to_views
- Grant only views that source view contains

### Scenario 4: User Registration

**Required:**
- Username (email or alphanumeric)
- Password (10+ complex or 17+ simple)
- Email

**Common Failures:**
- Invalid username format (OBP-20005)
- Weak password (OBP-30207)

**Resolution:**
- Use email format for username: user@example.com
- Password: MyP@ssw0rd123 (10+ with mixed case, numbers, special)
  OR LongPasswordMoreThan16Characters (17+ any characters)

### Scenario 5: API Rate Limit Exceeded

**Symptoms:**
- HTTP 429 Too Many Requests
- Rate limit headers in response

**Causes:**
- Exceeded per-second limit
- Exceeded per-minute limit
- Exceeded per-hour/day/week/month limit

**Resolution:**
- Implement backoff and retry logic
- Request higher rate limits from administrator
- Optimize API call patterns

### Scenario 6: Berlin Group Consent Creation

**Required:**
- Valid consent access specification
- Frequency per day
- Recurring indicator
- Valid date range (within 180 days)

**Common Failures:**
- Empty access (OBP-20088)
- Invalid frequency for availableAccounts (OBP-20090)
- Invalid recurring indicator (OBP-20089)
- Date beyond 180 days

**Resolution:**
- Specify access types
- Use frequency=1 for availableAccounts
- Set recurringIndicator=false for availableAccounts
- Use dates within 180 days

---

## Implementation Notes

### Validation Location in Codebase

**Pre-Request Validation:**
- Authentication: Before endpoint execution
- Rate limiting: Before endpoint execution
- Authorization headers: At request entry

**Endpoint Validation:**
- JSON schema: At endpoint entry (if configured)
- Authentication type: At endpoint execution
- Entitlements: At endpoint execution

**Business Logic Validation:**
- Within API method implementation
- In connector layer
- In domain models

**Data Layer Validation:**
- Database constraints
- Mapper validations
- Uniqueness constraints

### Validation Distribution

Validations are enforced at multiple layers:

1. **API Layer** (`code/api/`): Request format, authentication, authorization
2. **Validation Layer** (`code/validation/`, `code/authtypevalidation/`): JSON schema, auth types
3. **Business Logic Layer**: API methods, connectors
4. **Data Layer**: Database constraints, uniqueness
5. **Security Layer**: OAuth, certificates, signatures

### Custom Validations

Banks can extend validation through:
- JSON schema definitions per endpoint
- Authentication type restrictions per endpoint
- Custom attributes with validation rules
- Dynamic entities with field validations
- Rate limiting per endpoint

---

## Conclusion

This document catalogs 200+ validation rules enforced by the OBP API system. These validations ensure:

- **Data Integrity**: Proper formats and constraints
- **Security**: Authentication and authorization
- **Regulatory Compliance**: KYC, PSD2, data privacy
- **Business Logic**: Proper workflows and constraints
- **API Stability**: Rate limiting and input validation

For technical implementation details, refer to the source code files referenced throughout this document.

### Key Takeaways

1. **Validation is Multi-Layered**: Applied at API, business, and data layers
2. **Error Codes are Systematic**: Organized by category (OBP-XXXXX)
3. **Permissions are Granular**: 50+ view permissions, role-based entitlements
4. **Standards are Enforced**: ISO codes for currency, dates, phone numbers
5. **Security is Paramount**: Strong authentication, authorization, and input validation
6. **Compliance is Built-in**: KYC, PSD2, consent management

### Getting Help

- **Error Messages**: All errors include code (OBP-XXXXX) and descriptive message
- **API Explorer**: Interactive documentation at `/api-explorer`
- **Source Code**: Reference the files noted in each validation rule
- **Community**: Open Bank Project community and GitHub repository

---

*Document Version: 1.0*
*Generated from OBP-API codebase - develop branch*
*Last Updated: 2024*
