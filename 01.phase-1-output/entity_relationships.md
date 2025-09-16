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

## Additional Business Entities (Comprehensive Analysis Update)

### 14. Entitlement
**Source:** `obp-api/src/main/scala/code/entitlement/Entilement.scala` (lines 34-40)

| Attribute | Type | Description |
|-----------|------|-------------|
| entitlementId | String | Unique entitlement identifier |
| bankId | String | Reference to Bank |
| userId | String | Reference to User |
| roleName | String | Role name for authorization |
| createdByProcess | String | Process that created entitlement |

### 15. Meeting
**Source:** `obp-api/src/main/scala/code/meetings/Meetings.scala` (lines 25-50)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| staffUser | User | Staff user participating |
| customerUser | User | Customer user participating |
| providerId | String | Meeting provider identifier |
| purposeId | String | Meeting purpose identifier |
| when | Date | Meeting date and time |
| sessionId | String | Meeting session identifier |
| customerToken | String | Customer authentication token |
| staffToken | String | Staff authentication token |
| creator | ContactDetails | Meeting creator details |
| invitees | List[Invitee] | List of meeting invitees |

### 16. ContactMedium
**Source:** `obp-api/src/main/scala/code/meetings/Meetings.scala` (lines 11-14)

| Attribute | Type | Description |
|-----------|------|-------------|
| type | String | Contact medium type |
| value | String | Contact medium value |

### 17. Consent
**Source:** `obp-api/src/main/scala/code/consent/ConsentProvider.scala` (lines 59-195)

| Attribute | Type | Description |
|-----------|------|-------------|
| consentId | String | Unique consent identifier |
| userId | String | Reference to User |
| secret | String | Consent secret |
| status | String | Consent status |
| challenge | String | Hashed challenge for verification |
| jsonWebToken | String | JWT token for consent |
| consumerId | String | Consumer identifier |
| consentRequestId | String | Consent request identifier |
| apiStandard | String | API standard (OBP, Berlin-Group, UK) |
| apiVersion | String | API version |
| recurringIndicator | Boolean | Recurring access indicator |
| validUntil | Date | Consent validity end date |
| frequencyPerDay | Int | Maximum daily access frequency |
| usesSoFarTodayCounter | Int | Current daily usage counter |
| usesSoFarTodayCounterUpdatedAt | Date | Counter update timestamp |
| combinedServiceIndicator | Boolean | Combined service indicator |
| lastActionDate | Date | Last action date |
| creationDateTime | Date | Consent creation date |
| statusUpdateDateTime | Date | Status update date |
| expirationDateTime | Date | Consent expiration date |
| transactionFromDateTime | Date | Transaction query start date |
| transactionToDateTime | Date | Transaction query end date |
| consentReferenceId | String | Consent reference identifier |
| note | String | Consent notes |

### 18. ProductFee
**Source:** `obp-api/src/main/scala/code/productfee/ProductFee.scala` (lines 31-53)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| productCode | String | Reference to Product |
| productFeeId | String | Unique fee identifier |
| name | String | Fee name |
| isActive | Boolean | Fee active status |
| moreInfo | String | Additional fee information |
| currency | String | Fee currency |
| amount | BigDecimal | Fee amount |
| frequency | String | Fee frequency |
| type | String | Fee type |

### 19. AccountWebhook
**Source:** `obp-api/src/main/scala/code/webhook/AccountWebhook.scala` (lines 35-53)

| Attribute | Type | Description |
|-----------|------|-------------|
| accountWebhookId | String | Unique webhook identifier |
| bankId | String | Reference to Bank |
| accountId | String | Reference to BankAccount |
| triggerName | String | Webhook trigger name |
| url | String | Webhook URL |
| httpMethod | String | HTTP method |
| httpProtocol | String | HTTP protocol |
| createdByUserId | String | Reference to User |
| isActive | Boolean | Webhook active status |

### 20. CustomerMessage
**Source:** `obp-api/src/main/scala/code/customer/CustomerMessage.scala` (lines 16-27)

| Attribute | Type | Description |
|-----------|------|-------------|
| user | User | Reference to User |
| bankId | String | Reference to Bank |
| message | String | Message content |
| fromDepartment | String | Sending department |
| fromPerson | String | Sending person |
| transport | String | Message transport method |

### 21. CustomerAttribute
**Source:** `obp-api/src/main/scala/code/customerattribute/CustomerAttribute.scala` (lines 33-60)

| Attribute | Type | Description |
|-----------|------|-------------|
| customerId | String | Reference to Customer |
| customerAttributeId | String | Unique attribute identifier |
| name | String | Attribute name |
| attributeType | CustomerAttributeType | Attribute type enumeration |
| value | String | Attribute value |

### 22. BankAttribute
**Source:** `obp-api/src/main/scala/code/bankattribute/BankAttribute.scala` (lines 31-47)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| bankAttributeId | String | Unique attribute identifier |
| name | String | Attribute name |
| attributeType | BankAttributeType | Attribute type enumeration |
| value | String | Attribute value |
| isActive | Boolean | Attribute active status |

### 23. ProductAttribute
**Source:** `obp-api/src/main/scala/code/productattribute/ProductAttribute.scala` (lines 32-49)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| productCode | String | Reference to Product |
| productAttributeId | String | Unique attribute identifier |
| name | String | Attribute name |
| attributeType | ProductAttributeType | Attribute type enumeration |
| value | String | Attribute value |
| isActive | Boolean | Attribute active status |

### 24. CardAttribute
**Source:** `obp-api/src/main/scala/code/cardattribute/CardAttribute.scala` (lines 30-49)

| Attribute | Type | Description |
|-----------|------|-------------|
| cardId | String | Reference to PhysicalCard |
| cardAttributeId | String | Unique attribute identifier |
| name | String | Attribute name |
| attributeType | CardAttributeType | Attribute type enumeration |
| value | String | Attribute value |

### 25. AtmAttribute
**Source:** `obp-api/src/main/scala/code/atmattribute/AtmAttribute.scala` (lines 30-49)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| atmId | String | Reference to ATM |
| atmAttributeId | String | Unique attribute identifier |
| name | String | Attribute name |
| attributeType | AtmAttributeType | Attribute type enumeration |
| value | String | Attribute value |
| isActive | Boolean | Attribute active status |

## Extended Entity Relationships

### Additional Relationships (Comprehensive Analysis)

| From Entity | To Entity | Cardinality | Description |
|-------------|-----------|-------------|-------------|
| Bank | Entitlement | 1:N | One bank can grant multiple entitlements |
| Bank | Meeting | 1:N | One bank can host multiple customer meetings |
| Bank | BankAttribute | 1:N | One bank can have multiple attributes |
| BankAccount | AccountWebhook | 1:N | One account can have multiple webhooks |
| User | Entitlement | 1:N | One user can have multiple entitlements/roles |
| User | Meeting | N:M | Users (staff/customers) can participate in multiple meetings |
| User | Consent | 1:N | One user can have multiple consents |
| User | CustomerMessage | 1:N | One user can receive multiple messages |
| User | AccountWebhook | 1:N | One user can create multiple webhooks |
| Customer | Meeting | N:M | Customers can participate in multiple meetings |
| Customer | CustomerMessage | 1:N | One customer can receive multiple messages |
| Customer | CustomerAttribute | 1:N | One customer can have multiple attributes |
| Product | ProductFee | 1:N | One product can have multiple fees |
| Product | ProductAttribute | 1:N | One product can have multiple attributes |
| PhysicalCard | CardAttribute | 1:N | One card can have multiple attributes |
| ATM | AtmAttribute | 1:N | One ATM can have multiple attributes |

### Extended Business Rules

8. **Authorization Control**: Entitlements define user permissions within specific banks
9. **Consent Management**: Consents control API access with regulatory compliance features
10. **Meeting Coordination**: Meetings require both staff and customer participants
11. **Webhook Integration**: Account webhooks enable real-time event notifications
12. **Extensible Attributes**: All major entities support custom attributes for flexibility
13. **Fee Management**: Products can have multiple fees with different frequencies and types
14. **Message Communication**: Customer messages enable bank-customer communication

## Updated Data Sources Summary

| Entity Category | Primary Source Files | Line References |
|----------------|---------------------|-----------------|
| Core Banking | BankingModel.scala | 30-110 |
| Customer Management | UserModel.scala, CustomerDataModel.scala | 30-85 |
| Products | Products.scala, PhysicalCardModel.scala | 10-92 |
| Infrastructure | Branches.scala, Atms.scala | 16-66 |
| Authorization | Entilement.scala, ConsentProvider.scala | 34-195 |
| Customer Interaction | Meetings.scala, CustomerMessage.scala | 11-50 |
| Financial Products | ProductFee.scala | 31-53 |
| Integration | AccountWebhook.scala | 35-53 |
| Extensible Attributes | CustomerAttribute.scala, BankAttribute.scala, ProductAttribute.scala, CardAttribute.scala, AtmAttribute.scala | 30-60 |
| Persistence Layer | MappedBank.scala, MappedBankAccount.scala, ResourceUser.scala | 6-121 |

## Comprehensive Analysis Notes

- **Business Focus**: This comprehensive analysis prioritizes business-relevant entities over technical implementation details
- **Extended Coverage**: Added 12 additional business entities covering authorization, regulatory compliance, customer interaction, integration, and extensible data domains
- **Relationship Inference**: Relationships are derived from foreign key references, collection fields, and provider method signatures
- **Attribute Selection**: Only business-meaningful attributes are included; technical fields like timestamps are excluded unless they impact business rules
- **Source Traceability**: All entities include precise source file references for verification and maintenance
- **Regulatory Compliance**: Consent entities support Berlin Group and UK Open Banking standards
- **Extensible Design**: Attribute entities enable flexible data extension across all major business entities

### 26. EntitlementRequest
**Source:** `obp-api/src/main/scala/code/entitlementrequest/EntilementRequest.scala` (lines 29-39)

| Attribute | Type | Description |
|-----------|------|-------------|
| entitlementRequestId | String | Unique entitlement request identifier |
| bankId | String | Reference to Bank |
| user | User | Reference to User |
| roleName | String | Requested role name |
| created | Date | Request creation date |

### 27. KycStatus
**Source:** `obp-api/src/main/scala/code/kycstatus/KycStatus.scala` (lines 18-24)

| Attribute | Type | Description |
|-----------|------|-------------|
| customerId | String | Reference to Customer |
| customerNumber | String | Customer number |
| ok | Boolean | KYC verification status |
| date | Date | KYC status date |

### 28. UserScope
**Source:** `obp-api/src/main/scala/code/scope/UserScope.scala` (lines 13-16)

| Attribute | Type | Description |
|-----------|------|-------------|
| scopeId | String | Unique scope identifier |
| userId | String | Reference to User |

### 29. CounterpartyLimit
**Source:** `obp-api/src/main/scala/code/counterpartylimit/CounterpartyLimit.scala` (lines 13-41)

| Attribute | Type | Description |
|-----------|------|-------------|
| counterpartyLimitId | String | Unique counterparty limit identifier |
| bankId | String | Reference to Bank |
| accountId | String | Reference to BankAccount |
| viewId | String | View identifier |
| counterpartyId | String | Reference to Counterparty |
| currency | String | Limit currency |
| maxSingleAmount | BigDecimal | Maximum single transaction amount |
| maxMonthlyAmount | BigDecimal | Maximum monthly amount |
| maxNumberOfMonthlyTransactions | Int | Maximum monthly transaction count |
| maxYearlyAmount | BigDecimal | Maximum yearly amount |
| maxNumberOfYearlyTransactions | Int | Maximum yearly transaction count |
| maxTotalAmount | BigDecimal | Maximum total amount |
| maxNumberOfTransactions | Int | Maximum total transaction count |

### 30. DynamicResourceDoc
**Source:** `obp-api/src/main/scala/code/dynamicResourceDoc/DynamicResourceDoc.scala` (lines 10-27)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| dynamicResourceDocId | String | Unique resource doc identifier |
| partialFunctionName | String | Function name |
| requestVerb | String | HTTP request verb |
| requestUrl | String | Request URL |
| summary | String | Resource summary |
| description | String | Resource description |
| exampleRequestBody | String | Example request body |
| successResponseBody | String | Success response body |
| errorResponseBodies | String | Error response bodies |
| tags | String | Resource tags |
| roles | String | Required roles |
| methodBody | String | Method implementation body |

### 31. CustomerAccountLink
**Source:** `obp-api/src/main/scala/code/customeraccountlinks/CustomerAccountLink.scala` (lines 17-28)

| Attribute | Type | Description |
|-----------|------|-------------|
| customerAccountLinkId | String | Unique link identifier |
| customerId | String | Reference to Customer |
| bankId | String | Reference to Bank |
| accountId | String | Reference to BankAccount |
| relationshipType | String | Type of customer-account relationship |

### 32. KycDocument
**Source:** `obp-api/src/main/scala/code/kycdocuments/MappedKycDocumentsProvider.scala` (lines 50-77)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| customerId | String | Reference to Customer |
| idKycDocument | String | Unique KYC document identifier |
| customerNumber | String | Customer number |
| type | String | Document type |
| number | String | Document number |
| issueDate | Date | Document issue date |
| issuePlace | String | Document issue place |
| expiryDate | Date | Document expiry date |

### 33. DynamicMessageDoc
**Source:** `obp-api/src/main/scala/code/dynamicMessageDoc/DynamicMessageDoc.scala` (lines 8-26)

| Attribute | Type | Description |
|-----------|------|-------------|
| bankId | String | Reference to Bank |
| dynamicMessageDocId | String | Unique message doc identifier |
| process | String | Message process name |
| messageFormat | String | Message format |
| description | String | Message description |
| outboundTopic | String | Outbound message topic |
| inboundTopic | String | Inbound message topic |
| exampleOutboundMessage | String | Example outbound message |
| exampleInboundMessage | String | Example inbound message |
| outboundAvroSchema | String | Outbound Avro schema |
| inboundAvroSchema | String | Inbound Avro schema |
| adapterImplementation | String | Adapter implementation |
| methodBody | String | Method body |
| programmingLang | String | Programming language |

### 34. StandingOrder
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/CommonModelTrait.scala` (lines 699-715)

| Attribute | Type | Description |
|-----------|------|-------------|
| standingOrderId | String | Unique standing order identifier |
| bankId | String | Reference to Bank |
| accountId | String | Reference to BankAccount |
| customerId | String | Reference to Customer |
| userId | String | Reference to User |
| counterpartyId | String | Reference to Counterparty |
| amountValue | BigDecimal | Payment amount |
| amountCurrency | String | Payment currency |
| whenFrequency | String | Payment frequency |
| whenDetail | String | Payment timing details |
| dateSigned | Date | Order signing date |
| dateCancelled | Date | Order cancellation date |
| dateStarts | Date | Order start date |
| dateExpires | Date | Order expiry date |
| active | Boolean | Order active status |

### 35. EndpointTag
**Source:** `obp-commons/src/main/scala/com/openbankproject/commons/model/CommonModelTrait.scala` (lines 692-697)

| Attribute | Type | Description |
|-----------|------|-------------|
| endpointTagId | String | Unique endpoint tag identifier |
| operationId | String | API operation identifier |
| tagName | String | Tag name |
| bankId | String | Reference to Bank |

## Extended Entity Relationships (Comprehensive Final Update)

### Additional Relationships (Complete Analysis)

| From Entity | To Entity | Cardinality | Description |
|-------------|-----------|-------------|-------------|
| Bank | EntitlementRequest | 1:N | One bank can have multiple entitlement requests |
| Bank | CounterpartyLimit | 1:N | One bank can have multiple counterparty limits |
| Bank | DynamicResourceDoc | 1:N | One bank can have multiple dynamic resource docs |
| Bank | CustomerAccountLink | 1:N | One bank can have multiple customer-account links |
| Bank | KycDocument | 1:N | One bank can have multiple KYC documents |
| Bank | DynamicMessageDoc | 1:N | One bank can have multiple dynamic message docs |
| Bank | StandingOrder | 1:N | One bank can have multiple standing orders |
| Bank | EndpointTag | 1:N | One bank can have multiple endpoint tags |
| User | EntitlementRequest | 1:N | One user can have multiple entitlement requests |
| User | UserScope | 1:N | One user can have multiple scopes |
| User | StandingOrder | 1:N | One user can have multiple standing orders |
| Customer | KycStatus | 1:N | One customer can have multiple KYC status records |
| Customer | CustomerAccountLink | 1:N | One customer can be linked to multiple accounts |
| Customer | KycDocument | 1:N | One customer can have multiple KYC documents |
| Customer | StandingOrder | 1:N | One customer can have multiple standing orders |
| BankAccount | CounterpartyLimit | 1:N | One account can have multiple counterparty limits |
| BankAccount | CustomerAccountLink | 1:N | One account can be linked to multiple customers |
| BankAccount | StandingOrder | 1:N | One account can have multiple standing orders |
| Counterparty | CounterpartyLimit | 1:N | One counterparty can have multiple limits |
| Counterparty | StandingOrder | 1:N | One counterparty can receive multiple standing orders |

### Extended Business Rules (Final Comprehensive Update)

15. **Entitlement Requests**: Users can request specific roles/entitlements for banks
16. **KYC Management**: Customers have verification status and supporting documents
17. **User Scopes**: Users have specific permission scopes for API access
18. **Counterparty Limits**: Transaction limits can be set per counterparty relationship
19. **Dynamic Documentation**: Banks can have custom API and message documentation
20. **Customer-Account Links**: Flexible relationship types between customers and accounts
21. **Document Management**: KYC documents support customer verification processes
22. **Message Integration**: Dynamic message documentation for system integrations
23. **Standing Orders**: Recurring payment instructions with scheduling and lifecycle management
24. **API Management**: Endpoint tagging system for API organization and categorization

## Final Comprehensive Data Sources Summary

| Entity Category | Primary Source Files | Line References |
|----------------|---------------------|-----------------|
| Core Banking | BankingModel.scala | 30-110 |
| Customer Management | UserModel.scala, CustomerDataModel.scala | 30-85 |
| Products | Products.scala, PhysicalCardModel.scala | 10-92 |
| Infrastructure | Branches.scala, Atms.scala | 16-66 |
| Authorization | Entilement.scala, ConsentProvider.scala, EntilementRequest.scala | 34-195 |
| Customer Interaction | Meetings.scala, CustomerMessage.scala | 11-50 |
| Financial Products | ProductFee.scala | 31-53 |
| Integration | AccountWebhook.scala, DynamicResourceDoc.scala, DynamicMessageDoc.scala | 35-53 |
| Extensible Attributes | CustomerAttribute.scala, BankAttribute.scala, ProductAttribute.scala, CardAttribute.scala, AtmAttribute.scala | 30-60 |
| Compliance & Verification | KycStatus.scala, KycDocuments.scala | 18-77 |
| Access Control | UserScope.scala, CounterpartyLimit.scala | 13-41 |
| Customer Relations | CustomerAccountLink.scala | 17-28 |
| Payment Automation | CommonModelTrait.scala (StandingOrder) | 699-715 |
| API Management | CommonModelTrait.scala (EndpointTag) | 692-697 |
| Persistence Layer | MappedBank.scala, MappedBankAccount.scala, ResourceUser.scala | 6-121 |

## Final Comprehensive Analysis Notes

- **Business Focus**: This final comprehensive analysis prioritizes business-relevant entities over technical implementation details
- **Complete Coverage**: Added 10 additional business entities covering authorization requests, KYC management, user scopes, counterparty limits, dynamic documentation, customer-account relationships, document management, message integration, standing orders, and API management
- **Relationship Inference**: Relationships are derived from foreign key references, collection fields, and provider method signatures
- **Attribute Selection**: Only business-meaningful attributes are included; technical fields like timestamps are excluded unless they impact business rules
- **Source Traceability**: All entities include precise source file references for verification and maintenance
- **Regulatory Compliance**: Comprehensive KYC and consent management with document support
- **Extensible Design**: Multiple attribute and documentation entities enable flexible data extension and customization
- **Access Control**: Complete permission and scope management system
- **Integration Support**: Dynamic documentation for both API resources and message formats
- **Payment Automation**: Standing order support for recurring payment management
- **API Organization**: Endpoint tagging for API categorization and management

**Generated by:** Devin AI  
**Analysis Completion:** 16-9-2025  
**Total Business Entities Identified:** 35 core entities + 2 supporting entities  
**Total Relationships Mapped:** 52 primary relationships
