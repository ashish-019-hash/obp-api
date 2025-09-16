# Entity Relationship Model - OBP-API Business Entities

**Analysis Date:** 16-9-2025  
**Repository:** ashish-019-hash/obp-api  
**Input Source:** 00.phase-1-input folder

## Overview

This document presents a comprehensive Entity Relationship Model (ERM) for the Open Bank Project API, focusing on real-world business entities and their relationships. The analysis extracts domain models, case classes, and traits from the business logic and persistence layers while excluding UI-specific fields and technical implementation details.

## Core Banking Entities

### 1. Bank
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala` (lines 30-40)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | BankId | Unique identifier for the bank |
| shortName | String | Short display name of the bank |
| fullName | String | Full legal name of the bank |
| logoUrl | String | URL to bank's logo image |
| websiteUrl | String | Bank's official website URL |
| swiftBic | String | SWIFT Bank Identifier Code |
| nationalIdentifier | String | National banking identifier |
| bankRoutingScheme | String | Routing scheme used by bank |
| bankRoutingAddress | String | Routing address for the bank |

**Mapped Implementation:** `code/model/dataAccess/MappedBank.scala` (lines 6-29)

### 2. BankAccount
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala` (lines 42-60)

| Attribute | Type | Description |
|-----------|------|-------------|
| accountId | AccountId | Unique identifier for the account |
| bankId | BankId | Reference to the owning bank |
| accountType | String | Type of account (savings, checking, etc.) |
| balance | BigDecimal | Current account balance |
| currency | String | Currency code (USD, EUR, etc.) |
| name | String | Display name for the account |
| number | String | Account number |
| label | String | User-defined label for the account |
| lastUpdate | Date | Last transaction update timestamp |
| branchId | String | Reference to the branch |

**Mapped Implementation:** `code/model/dataAccess/MappedBankAccount.scala` (lines 11-75)

### 3. Transaction
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala` (lines 62-85)

| Attribute | Type | Description |
|-----------|------|-------------|
| id | TransactionId | Unique transaction identifier |
| thisAccount | BankIdAccountId | Source account reference |
| otherAccount | Counterparty | Destination account/counterparty |
| transactionType | String | Type of transaction |
| amount | BigDecimal | Transaction amount |
| currency | String | Transaction currency |
| description | Option[String] | Transaction description |
| startDate | Date | Transaction initiation date |
| finishDate | Date | Transaction completion date |
| balance | BigDecimal | Account balance after transaction |

### 4. Counterparty
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala` (lines 87-110)

| Attribute | Type | Description |
|-----------|------|-------------|
| counterpartyId | String | Unique counterparty identifier |
| counterpartyName | String | Display name of counterparty |
| thisBankId | String | Reference to this bank |
| thisAccountId | String | Reference to this account |
| otherBankRoutingScheme | String | Other bank's routing scheme |
| otherBankRoutingAddress | String | Other bank's routing address |
| otherAccountRoutingScheme | String | Other account's routing scheme |
| otherAccountRoutingAddress | String | Other account's routing address |
| isBeneficiary | Boolean | Whether this is a beneficiary |

## Customer Management Entities

### 5. User
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/UserModel.scala` (lines 30-50)

| Attribute | Type | Description |
|-----------|------|-------------|
| userPrimaryKey | UserPrimaryKey | Primary key for user |
| userId | String | Unique user identifier |
| provider | String | Identity provider |
| emailAddress | String | User's email address |
| name | String | User's display name |
| createdByConsentId | Option[String] | Consent that created this user |
| createdByUserInvitationId | Option[String] | Invitation that created this user |
| isDeleted | Option[Boolean] | Soft delete flag |
| lastMarketingAgreementSignedDate | Option[Date] | Marketing agreement date |

**Mapped Implementation:** `code/model/dataAccess/ResourceUser.scala` (lines 59-121)

### 6. Customer
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/CustomerDataModel.scala` (lines 30-60)

| Attribute | Type | Description |
|-----------|------|-------------|
| customerId | String | Unique customer identifier |
| bankId | String | Reference to customer's bank |
| number | String | Customer number |
| legalName | String | Legal name of customer |
| mobileNumber | String | Customer's mobile phone |
| email | String | Customer's email address |
| faceImage | CustomerFaceImage | Customer's face image data |
| dateOfBirth | Date | Customer's date of birth |
| relationshipStatus | String | Relationship with bank |
| dependents | Integer | Number of dependents |
| dobOfDependents | List[Date] | Dependents' dates of birth |
| highestEducationAttained | String | Education level |
| employmentStatus | String | Current employment status |
| kycStatus | Boolean | Know Your Customer status |
| lastOkDate | Date | Last KYC verification date |

### 7. Agent
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/CustomerDataModel.scala` (lines 70-85)

| Attribute | Type | Description |
|-----------|------|-------------|
| agentId | String | Unique agent identifier |
| bankId | String | Reference to agent's bank |
| legalName | String | Legal name of agent |
| mobileNumber | String | Agent's mobile phone |
| email | String | Agent's email address |
| isConfirmedAgent | Boolean | Agent confirmation status |

## Product Entities

### 8. Product
**Source:** `code/products/Products.scala` (lines 10-50) and `obp-commons/src/main/scala/com/openbankproject/commons/model/Product.scala`

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | BankId | Reference to offering bank |
| productCode | ProductCode | Unique product code |
| name | String | Product name |
| category | String | Product category |
| family | String | Product family |
| superFamily | String | Product super family |
| moreInfoUrl | String | URL for more information |
| termsAndConditionsUrl | String | Terms and conditions URL |
| description | String | Product description |
| meta | Meta | Metadata information |

### 9. PhysicalCard
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/PhysicalCardModel.scala` (lines 42-92)

| Attribute | Type | Description |
|-----------|------|-------------|
| cardId | String | Unique card identifier |
| bankId | String | Reference to issuing bank |
| bankCardNumber | String | Card number |
| cardType | String | Type of card (credit, debit) |
| nameOnCard | String | Cardholder name |
| issueNumber | String | Card issue number |
| serialNumber | String | Card serial number |
| validFrom | Date | Card validity start date |
| expires | Date | Card expiration date |
| enabled | Boolean | Card enabled status |
| cancelled | Boolean | Card cancellation status |
| onHotList | Boolean | Card security status |
| technology | String | Card technology (chip, magnetic) |
| networks | List[String] | Payment networks |
| allows | List[CardAction] | Allowed card actions |
| customerId | String | Reference to card owner |
| cvv | Option[String] | Card verification value |
| brand | Option[String] | Card brand |

## Infrastructure Entities

### 10. Branch
**Source:** `code/branches/Branches.scala` (lines 17-36)

| Attribute | Type | Description |
|-----------|------|-------------|
| branchId | BranchId | Unique branch identifier |
| bankId | BankId | Reference to owning bank |
| name | String | Branch name |
| address | Address | Branch physical address |
| location | Location | Geographic coordinates |
| lobby | Option[Lobby] | Lobby operating hours |
| driveUp | Option[DriveUp] | Drive-up operating hours |
| isAccessible | Option[Boolean] | Accessibility features |
| accessibleFeatures | Option[String] | Accessibility details |
| branchType | Option[String] | Type of branch |
| moreInfo | Option[String] | Additional information |
| phoneNumber | Option[String] | Branch phone number |
| branchRouting | Option[Routing] | Branch routing information |

### 11. ATM
**Source:** `code/atms/Atms.scala` (lines 16-66)

| Attribute | Type | Description |
|-----------|------|-------------|
| atmId | AtmId | Unique ATM identifier |
| bankId | BankId | Reference to owning bank |
| name | String | ATM name/location |
| address | Address | ATM physical address |
| location | Location | Geographic coordinates |
| isAccessible | Option[Boolean] | Accessibility features |
| locatedAt | Option[String] | Location description |
| moreInfo | Option[String] | Additional information |
| hasDepositCapability | Option[Boolean] | Deposit functionality |
| supportedLanguages | Option[List[String]] | Supported languages |
| services | Option[List[String]] | Available services |
| accessibilityFeatures | Option[List[String]] | Accessibility features |
| supportedCurrencies | Option[List[String]] | Supported currencies |
| minimumWithdrawal | Option[String] | Minimum withdrawal amount |
| phone | Option[String] | ATM contact phone |

## Supporting Entities

### 12. Address
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala`

| Attribute | Type | Description |
|-----------|------|-------------|
| line1 | String | Address line 1 |
| line2 | String | Address line 2 |
| line3 | String | Address line 3 |
| city | String | City name |
| county | String | County/state |
| state | String | State/province |
| postCode | String | Postal code |
| countryCode | String | Country code |

### 13. Location
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala`

| Attribute | Type | Description |
|-----------|------|-------------|
| latitude | Double | Geographic latitude |
| longitude | Double | Geographic longitude |
| date | Option[Date] | Location timestamp |
| user | Option[BasicUser] | User who provided location |

## Entity Relationships

### Primary Relationships

| From Entity | To Entity | Cardinality | Description |
|-------------|-----------|-------------|-------------|
| Bank | BankAccount | 1:N | One bank has many accounts |
| Bank | Branch | 1:N | One bank has many branches |
| Bank | ATM | 1:N | One bank has many ATMs |
| Bank | Product | 1:N | One bank offers many products |
| Bank | Customer | 1:N | One bank serves many customers |
| BankAccount | Transaction | 1:N | One account has many transactions |
| BankAccount | PhysicalCard | 1:N | One account can have multiple cards |
| Customer | BankAccount | N:M | Customers can hold multiple accounts |
| Customer | PhysicalCard | 1:N | One customer can have multiple cards |
| User | Customer | N:M | Users can be linked to multiple customers |
| BankAccount | Counterparty | 1:N | One account can have multiple counterparties |
| Transaction | Counterparty | N:1 | Many transactions can involve same counterparty |
| Branch | Address | 1:1 | Each branch has one address |
| ATM | Address | 1:1 | Each ATM has one address |
| Branch | Location | 1:1 | Each branch has geographic coordinates |
| ATM | Location | 1:1 | Each ATM has geographic coordinates |

### Key Business Rules

1. **Account Ownership**: Accounts can have multiple owners (joint accounts) through Customer-BankAccount many-to-many relationship
2. **User-Customer Linking**: Users can be associated with multiple customers across different banks
3. **Card Issuance**: Physical cards are issued to specific customers for specific accounts
4. **Transaction Flow**: Transactions always involve a source account and a counterparty (which may represent another account)
5. **Geographic Distribution**: Banks operate through branches and ATMs with specific geographic locations
6. **Product Offering**: Banks offer various financial products that can be associated with accounts

## Data Sources Summary

| Entity Category | Primary Source Files | Line References |
|----------------|---------------------|-----------------|
| Core Banking | BankingModel.scala | 30-110 |
| Customer Management | UserModel.scala, CustomerDataModel.scala | 30-85 |
| Products | Products.scala, PhysicalCardModel.scala | 10-92 |
| Infrastructure | Branches.scala, Atms.scala | 16-66 |
| Persistence Layer | MappedBank.scala, MappedBankAccount.scala, ResourceUser.scala | 6-121 |

## Analysis Notes

- **Business Focus**: This analysis prioritizes business-relevant entities over technical implementation details
- **Relationship Inference**: Relationships are derived from foreign key references, collection fields, and provider method signatures
- **Attribute Selection**: Only business-meaningful attributes are included; technical fields like timestamps are excluded unless they impact business rules
- **Source Traceability**: All entities include precise source file references for verification and maintenance

---

**Generated by:** Devin AI  
**Analysis Completion:** 16-9-2025  
**Total Business Entities Identified:** 13 core entities + 2 supporting entities  
**Total Relationships Mapped:** 15 primary relationships
