# OBP-API Business Entity Catalog

## Document Purpose
This document provides a comprehensive catalog of all business entities found in the Open Bank Project (OBP) API codebase. Each entity represents a real-world business concept that the banking system needs to track and manage.

**Generated Date**: October 12, 2025  
**Codebase Location**: https://github.com/OpenBankProject/OBP-API  
**Main Entity Path**: `obp-api/src/main/scala/code/`  
**Total Entities Documented**: 100

---

## Executive Summary

The OBP-API is a comprehensive open banking platform managing a rich ecosystem of 100 interconnected business entities across multiple domains. The system follows a consistent architectural pattern using the Lift ORM framework with Provider/Mapped implementation classes, ensuring maintainable and scalable code organization.

### Key Architectural Patterns
- **Provider Pattern**: Each entity has a Provider trait defining the interface and a Mapped implementation
- **UUID-based Identity**: Entities use UUIDs for primary identification ensuring global uniqueness
- **Flexible Attributes**: Core entities can be extended with custom attributes for bank-specific needs
- **Relationship Management**: Foreign keys and junction tables manage complex entity relationships
- **Audit Trail**: Most entities include CreatedUpdated timestamps for tracking

---

## Business Domain Overview

The OBP-API system manages the following major business domains:

### 1. Core Banking Infrastructure
Banks, Branches, ATMs, Accounts, and fundamental banking structures

### 2. Customer Management
Customers, Addresses, Attributes, Customer-Account Links, Agents

### 3. Transaction Processing
Transactions, Transaction Requests, Transaction Types, Transaction Attributes

### 4. Card Services
Physical Cards, PIN Management, Card Attributes

### 5. Products & Services
Banking Products, Product Fees, Product Collections, Product Attributes

### 6. Payment Operations
Direct Debits, Standing Orders, Counterparty Limits, Counterparties

### 7. KYC & Compliance
KYC Checks, KYC Documents, KYC Status, Tax Residence, Regulated Entities

### 8. Consent & Authorization
Consents, Entitlements, Signing Baskets (PSD2), Scopes

### 9. User Management & Security
Users, API Consumers, Views (Access Control), Account Access, Login Tracking

### 10. Metadata & Enrichment
Transaction Images, Tags, Narratives, Comments, Geographic Tags (WhereTag)

### 11. System Management
Metrics, Rate Limiting, Webhooks, Meetings, CRM Events, Attribute Definitions

### 12. Financial Operations
Foreign Exchange Rates, Currencies, Account Applications

### 13. Technical Infrastructure
Dynamic Entities, Dynamic Endpoints, Connector Methods, Endpoint Mappings, JSON Schema Validation

### 14. Security & OAuth
Nonces, ETags, PEM Certificate Usage

---

## Complete Entity Catalog by Domain

### CORE BANKING INFRASTRUCTURE (9 entities)
1. **Bank** - Financial institutions in the system
2. **BankAccount** - Customer bank accounts with balances
3. **Branch** - Physical bank branch locations
4. **ATM** - Automated teller machine locations and capabilities
5. **AccountHolder** - Links between accounts and their holders
6. **BankAttribute** - Custom attributes for banks
7. **AccountAttribute** - Custom attributes for accounts
8. **AtmAttribute** - Custom attributes for ATMs
9. **AccountRouting** - Routing information for accounts (IBAN, account number, etc.)

### CUSTOMER MANAGEMENT (7 entities)
10. **Customer** - Bank customers with personal information
11. **CustomerAddress** - Customer physical addresses
12. **CustomerAttribute** - Custom attributes for customers
13. **CustomerAccountLink** - Junction table linking customers to accounts
14. **UserCustomerLink** - Links between users and customers
15. **CustomerDependant** - Customer dependent information
16. **CustomerMessage** - Messages sent to customers

### TRANSACTION PROCESSING (8 entities)
17. **Transaction** - Financial transactions on accounts
18. **TransactionRequest** - Requests to initiate transactions
19. **TransactionRequestTypeCharge** - Fees for transaction types
20. **TransactionRequestReasons** - Reasons for transaction requests
21. **TransactionAttribute** - Custom attributes for transactions
22. **TransactionRequestAttribute** - Custom attributes for transaction requests
23. **TransactionType** - Types/categories of transactions
24. **TransactionChallenge** - Security challenges for transactions

### CARD SERVICES (3 entities)
25. **PhysicalCard** - Physical payment cards (debit/credit)
26. **CardAction** - Actions/permissions for cards
27. **CardAttribute** - Custom attributes for cards

### PRODUCTS & SERVICES (5 entities)
28. **Product** - Banking products (accounts, loans, savings)
29. **ProductFee** - Fees associated with products
30. **ProductCollection** - Collections/bundles of products
31. **ProductCollectionItem** - Items within product collections
32. **ProductAttribute** - Custom attributes for products

### PAYMENT OPERATIONS (5 entities)
33. **DirectDebit** - Recurring payment authorizations
34. **StandingOrder** - Standing payment orders
35. **Counterparty** - Other parties in transactions
36. **CounterpartyMetadata** - Additional counterparty information
37. **CounterpartyLimit** - Transaction limits for counterparties

### KYC & COMPLIANCE (7 entities)
38. **KycCheck** - Know Your Customer verification checks
39. **KycDocument** - KYC supporting documents
40. **KycMedia** - Media files for KYC
41. **KycStatus** - KYC verification status
42. **TaxResidence** - Customer tax residence information
43. **RegulatedEntity** - Regulated financial entities
44. **RegulatedEntityAttribute** - Attributes for regulated entities

### CONSENT & AUTHORIZATION (7 entities)
45. **Consent** - Customer consent for data sharing
46. **ConsentAuthContext** - Authentication context for consents
47. **ExpectedChallengeAnswer** - Expected answers for transaction challenges
48. **Entitlement** - User permissions/roles
49. **EntitlementRequest** - Requests for entitlements
50. **SigningBasket** - PSD2 batch authorization baskets
51. **Scope** - OAuth/API scopes

### USER MANAGEMENT & SECURITY (14 entities)
52. **ResourceUser** - System users with authentication
53. **UserRefreshes** - User refresh tracking for cache management
54. **UserAttribute** - Custom attributes for users
55. **AuthUser** - Authentication credentials for users
56. **Consumer** - API consumer applications
57. **View** - Access control views on accounts
58. **ViewDefinition** - Definition of system/custom views
59. **AccountAccess** - User access to account views
60. **ViewPermission** - Granular view permissions
61. **BadLoginAttempt** - Failed login attempt tracking
62. **UserScope** - Junction table linking users to OAuth scopes
63. **UserInvitation** - User invitation workflow with secret links
64. **UserAgreement** - User agreement acceptance with hash verification
65. **UserInitAction** - User initialization action tracking

### METADATA & ENRICHMENT (6 entities)
66. **TransactionImage** - Images attached to transactions
67. **Tag** - Tags for categorizing transactions
68. **WhereTag** - Geographic location tags for transactions
69. **Narrative** - User-editable transaction descriptions
70. **Comment** - Comments on transactions
71. **SocialMedia** - Social media handles for entities

### SYSTEM MANAGEMENT (7 entities)
72. **Metric** - API usage metrics and analytics
73. **MetricArchive** - Archived metrics
74. **RateLimiting** - API rate limiting configuration
75. **Webhook** - Event webhooks for notifications
76. **Meeting** - Customer meetings
77. **CrmEvent** - CRM events and interactions
78. **AttributeDefinition** - Schema definitions for dynamic attributes

### FINANCIAL OPERATIONS (3 entities)
79. **FXRate** - Foreign exchange rates
80. **Currency** - Currency definitions
81. **AccountApplication** - Account opening applications

### TECHNICAL INFRASTRUCTURE (14 entities)
82. **DynamicEntity** - Runtime-configurable custom entities
83. **DynamicData** - Data storage for dynamic entities
84. **DynamicEndpoint** - Runtime-configurable API endpoints
85. **DynamicMessageDoc** - Message documentation for connectors
86. **DynamicResourceDoc** - Resource documentation for endpoints
87. **ConnectorMethod** - Custom connector method definitions
88. **MethodRouting** - API method routing configuration
89. **EndpointMapping** - Endpoint mapping configuration
90. **EndpointTag** - Tags for API endpoints
91. **WebUiProps** - Web UI properties and configuration
92. **JsonSchemaValidation** - JSON schema validation rules
93. **AccountIdMapping** - Internal account ID mappings
94. **TransactionIdMapping** - Internal transaction ID mappings
95. **CustomerIdMapping** - Internal customer ID mappings

### SECURITY & OAUTH (5 entities)
96. **OpenIDConnectToken** - OpenID Connect authentication tokens
97. **AuthenticationTypeValidation** - Validation rules for authentication types
98. **Nonce** - OAuth nonce for replay protection
99. **ETag** - Entity tags for HTTP caching
100. **PemUsage** - PEM certificate usage tracking

---

## Detailed Entity Documentation

### CORE BANKING INFRASTRUCTURE

---

### BANK ENTITY

**Entity Name**: Bank

**Business Description**: Represents a financial institution within the Open Bank Project system. A bank is the top-level organizational unit that owns branches, ATMs, accounts, and customers. Multiple banks can coexist in the system, enabling multi-bank scenarios.

**Key Fields**:
- `permalink` (String, max 255): Unique bank identifier used in URLs (serves as bankId)
- `fullBankName` (String, max 255): Official full name of the bank
- `shortBankName` (String, max 100): Abbreviated bank name
- `logoURL` (String, max 255): URL to bank's logo image
- `websiteURL` (String, max 255): Bank's website URL
- `swiftBIC` (String, max 255): SWIFT/BIC code for international transfers
- `national_identifier` (String, max 255): National bank identifier
- `mBankRoutingScheme` (String, max 255): Routing scheme (e.g., "IBAN", "AccountNumber")
- `mBankRoutingAddress` (String, max 255): Routing address value

**Relationships to Other Entities**:
- A Bank has many Branches (one-to-many)
- A Bank has many ATMs (one-to-many)
- A Bank has many BankAccounts (one-to-many)
- A Bank has many Customers (one-to-many)
- A Bank has many Products (one-to-many)
- A Bank can have BankAttributes (one-to-many)

**Business Rules**:
- `permalink` should be unique (indexed but not enforced as unique constraint in code)
- `permalink` serves as the external bank identifier (BankId)

**Notes**:
- The code comments indicate there should be a UniqueIndex on permalink but it's not implemented due to test dependencies
- Other models could foreign key to this but would need to expose IdPK

---

### BANK ATTRIBUTE ENTITY

**Entity Name**: BankAttribute

**Business Description**: Represents custom attributes that can be attached to banks to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing banks to add custom properties without modifying core code. Attributes can represent regulatory information, operational details, or any bank-specific configuration.

**Key Fields**:
- `BankId_` (UUIDString): The bank this attribute belongs to
- `BankAttributeId` (MappedUUID): Unique identifier for the attribute
- `Name` (String, max 50): Name/key of the attribute
- `Type` (String, max 50): Attribute type as defined by BankAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `Value` (String, max 255): The attribute value stored as string
- `IsActive` (Boolean): Whether this attribute is currently active (default: true)

**Relationships to Other Entities**:
- A BankAttribute belongs to one Bank (via BankId_)
- Multiple attributes can be attached to the same bank

**Business Rules**:
- BankId_ is indexed for efficient queries
- IsActive defaults to true
- Type determines how the value string should be interpreted
- Only active attributes are typically returned in API responses

**Notes**:
- Enables bank-specific customization without code changes
- Supports various data types through the Type field
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API

---
### ACCOUNT ATTRIBUTE ENTITY

**Entity Name**: AccountAttribute

**Business Description**: Represents custom attributes that can be attached to bank accounts to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing accounts to have custom properties defined dynamically. Attributes can represent account-specific configurations, regulatory requirements, product features, or any custom metadata needed for account management.

**Key Fields**:
- `mBankIdId` (UUIDString): The bank this attribute belongs to
- `mAccountId` (UUIDString): The account this attribute belongs to
- `mAccountAttributeId` (MappedUUID): Unique identifier for the attribute
- `mCode` (String, max 50): Product code associated with this attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by AccountAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string
- `mProductInstanceCode` (String, max 255): Optional product instance identifier

**Relationships to Other Entities**:
- An AccountAttribute belongs to one BankAccount (via mAccountId)
- An AccountAttribute belongs to one Bank (via mBankIdId)
- Multiple attributes can be attached to the same account
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and accountId is indexed for efficient queries
- AccountAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Attributes can be filtered based on View permissions
- Supports product code and product instance code for product-specific attributes

**Notes**:
- Enables account-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Supports querying accounts by attribute name/value pairs
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API

---



### BANK ACCOUNT ENTITY

**Entity Name**: BankAccount

**Business Description**: Represents a customer's bank account where money is stored and transactions occur. Accounts have a balance, currency, account type, and various routing identifiers for payments.

**Key Fields**:
- `bank` (UUIDString): The bank this account belongs to
- `theAccountId` (AccountIdString): Unique account identifier
- `accountCurrency` (String, max 10): Currency code (e.g., "USD", "EUR")
- `accountNumber` (MappedAccountNumber): Account number
- `accountBalance` (Long): Balance in smallest currency units (cents, pence, etc.)
- `accountName` (String, max 255): Account name/description
- `kind` (String, max 255): Account type/financial product name
- `accountLabel` (String, max 255): User-friendly account label
- `accountLastUpdate` (DateTime): Last transaction refresh date
- `mBranchId` (UUIDString): Associated branch
- `accountRuleScheme1/2` (String, max 10): Rule schemes
- `accountRuleValue1/2` (Long): Rule values in smallest currency units
- `holder` (String, max 100): Deprecated holder field

**Relationships to Other Entities**:
- A BankAccount belongs to one Bank (via bank foreign key)
- A BankAccount belongs to one Branch (via mBranchId)
- A BankAccount has many Transactions (one-to-many)
- A BankAccount has many AccountHolders (one-to-many via junction)
- A BankAccount has many CustomerAccountLinks (one-to-many)
- A BankAccount has AccountAttributes (one-to-many)
- A BankAccount has AccountRoutings (one-to-many)
- A BankAccount can have PhysicalCards (one-to-many)

**Business Rules**:
- Combination of bank and theAccountId must be unique
- Balance is stored as Long (smallest currency unit) and converted to BigDecimal for display
- Currency code is always uppercase

**Notes**:
- The `holder` field is deprecated
- Balance conversion uses Helper.smallestCurrencyUnitToBigDecimal
- Account rules allow for flexible business logic (min/max balance, etc.)

---

### ATM ATTRIBUTE ENTITY

**Entity Name**: AtmAttribute

**Business Description**: Represents custom attributes that can be attached to ATMs to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing banks to customize ATM information without modifying core code.

**Key Fields**:
- `mAtmAttributeId` (UUIDString): Unique identifier for the attribute
- `mBankId` (UUIDString): The bank that owns the ATM
- `mAtmId` (UUIDString): The ATM this attribute belongs to
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 20): Attribute type - one of: STRING, INTEGER, DOUBLE, DATE_WITH_DAY
- `mValue` (String, max 10000): The attribute value stored as string
- `mIsActive` (Boolean): Whether this attribute is currently active

**Relationships to Other Entities**:
- An AtmAttribute belongs to one ATM (via mAtmId)
- An AtmAttribute belongs to one Bank (via mBankId)

**Business Rules**:
- Combination of BankId, AtmId, and Name should identify a unique attribute
- Type determines how the value string should be interpreted
- Only active attributes are typically returned in API responses
- Values stored as strings but typed according to mType field

**Notes**:
- Enables bank-specific customization of ATM data
- Supports multiple data types through the mType field
- Part of the broader attribute extensibility pattern in OBP-API

---

### BANK ACCOUNT ROUTING ENTITY

**Entity Name**: BankAccountRouting

**Business Description**: Stores routing information for bank accounts, such as IBAN, account numbers, sort codes, and other routing schemes. An account can have multiple routing addresses for different payment networks and schemes.

**Key Fields**:
- `BankId` (UUIDString): The bank identifier
- `AccountId` (AccountIdString): The account identifier
- `AccountRoutingScheme` (String, max 32): The routing scheme name (e.g., "IBAN", "AccountNumber", "SortCode")
- `AccountRoutingAddress` (String, max 128): The routing address value for the specified scheme

**Relationships to Other Entities**:
- A BankAccountRouting belongs to one BankAccount (via BankId and AccountId)
- An account can have multiple routing entries for different schemes

**Business Rules**:
- Combination of BankId, AccountId, and AccountRoutingScheme must be unique
- Combination of BankId, AccountRoutingScheme, and AccountRoutingAddress must also be unique
- Common routing schemes include: IBAN, AccountNumber, SortCode, RoutingNumber, BankCode

**Notes**:
- Supports international payment routing across different banking systems
- Critical for PSD2 and Open Banking compliance
- Allows accounts to be identified through multiple routing mechanisms
- Two unique indexes ensure data integrity from different access patterns

---

### ACCOUNT HOLDER ENTITY

**Entity Name**: AccountHolder

**Business Description**: Links a ResourceUser (system user) to a bank account, tracking the source of the relationship (e.g., created through an account application). This junction entity establishes ownership and access relationships between users and accounts.

**Key Fields**:
- `user` (MappedLongForeignKey): Foreign key to ResourceUser (the person who holds the account)
- `accountBankPermalink` (UUIDString): The bank identifier
- `accountPermalink` (AccountIdString): The account identifier
- `source` (String, max 50): How this holder relationship was created (e.g., "ACCOUNT_APPLICATION", "MANUAL")

**Relationships to Other Entities**:
- An AccountHolder links one ResourceUser to one BankAccount
- Multiple AccountHolders can exist for the same account (joint accounts)
- An AccountHolder is created when an AccountApplication is approved

**Business Rules**:
- Combination of user, bank, and account should be unique per holder
- Source field tracks the origination of the account holder relationship
- Links established through account application workflow have source = "ACCOUNT_APPLICATION"

**Notes**:
- Enables joint account scenarios where multiple users hold the same account
- The source field provides audit trail for how account holder relationships were created
- Part of the account ownership and access control system

---

### CUSTOMER MANAGEMENT

---

### 1. CUSTOMER ENTITY

**Entity Name**: Customer

**Business Description**: Represents a bank customer - an individual or entity that has a relationship with the bank. Customers can have multiple accounts across different banks and are the primary users of banking services.

**Key Fields**:
- `customerId` (String): Unique identifier for the customer (UUID)
- `bankId` (String): The bank this customer belongs to
- `number` (String): Customer number within the bank
- `legalName` (String, max 255): Full legal name of the customer
- `mobileNumber` (String, max 50): Customer's mobile phone number
- `email` (String, max 200): Customer's email address
- `faceImage` (CustomerFaceImageTrait): Customer's photo with URL and timestamp
- `dateOfBirth` (Date): Customer's date of birth
- `relationshipStatus` (String, max 16): Marital/relationship status
- `dependents` (Integer): Number of dependents
- `dobOfDependents` (List[Date]): Dates of birth of dependents (stored in separate table)
- `highestEducationAttained` (String, max 32): Education level
- `employmentStatus` (String, max 32): Current employment status
- `kycStatus` (Boolean): Know Your Customer verification status
- `lastOkDate` (Date): Last date KYC was confirmed OK
- `creditRating` (CreditRatingTrait): Credit rating information
  - `rating` (String): The credit rating value
  - `source` (String): Source of the credit rating
- `creditLimit` (AmountOfMoneyTrait): Credit limit assigned to customer
  - `currency` (String): Currency code
  - `amount` (String): Credit limit amount
- `title` (String, max 255): Customer's title (Mr., Mrs., Dr., etc.)
- `branchId` (String, max 255): Home branch for this customer
- `nameSuffix` (String, max 255): Name suffix (Jr., Sr., III, etc.)
- `isPendingAgent` (Boolean): Whether customer is pending agent approval
- `isConfirmedAgent` (Boolean): Whether customer is confirmed as an agent

**Relationships to Other Entities**:
- A Customer belongs to one Bank (via bankId)
- A Customer can have multiple Accounts (via UserCustomerLink)
- A Customer can have multiple Cards (via customerId)
- A Customer has one or more CustomerDependants (one-to-many)
- A Customer can be linked to a User (via UserCustomerLink)
- A Customer can have multiple CustomerAddresses
- A Customer can have multiple CustomerAttributes

**Business Rules**:
- Combination of bankId and customer number must be unique
- customerId (UUID) must be unique across all customers
- Email should be in valid email format (MappedEmail validation)
- Customer can also serve as an Agent (dual role)

**Notes**:
- The Customer entity also implements the Agent interface, allowing customers to act as bank agents
- Customer dependents are stored in a separate OneToMany table
- Credit information is optional and stored as embedded objects

---
### CUSTOMER ATTRIBUTE ENTITY

**Entity Name**: CustomerAttribute

**Business Description**: Represents custom attributes that can be attached to customers to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing customers to have custom properties defined dynamically. Attributes can represent customer-specific information, regulatory requirements, marketing preferences, or any custom metadata needed for customer management and compliance.

**Key Fields**:
- `mBankId` (UUIDString): The bank this attribute belongs to
- `mCustomerId` (UUIDString): The customer this attribute belongs to
- `mCustomerAttributeId` (MappedUUID): Unique identifier for the attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by CustomerAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string

**Relationships to Other Entities**:
- A CustomerAttribute belongs to one Customer (via mCustomerId)
- A CustomerAttribute belongs to one Bank (via mBankId)
- Multiple attributes can be attached to the same customer
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and customerId is indexed for efficient queries
- CustomerAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Attributes can be filtered based on View permissions
- Multiple attributes with the same name are allowed for the same customer

**Notes**:
- Enables customer-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Supports querying customers by attribute name/value pairs
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing customer preferences, tags, custom fields, and regulatory data

---



### CUSTOMER DEPENDANT ENTITY

**Entity Name**: CustomerDependant

**Business Description**: Stores information about a customer's dependants (children, family members, or others financially dependent on the customer). Each dependant has a date of birth recorded, which is used for various banking products and services.

**Key Fields**:
- `id` (Long): Primary key
- `customer` (MappedLongForeignKey): Foreign key to the Customer entity
- `dependant_dob` (MappedDate): Date of birth of the dependant

**Relationships to Other Entities**:
- A CustomerDependant belongs to one Customer (via customer foreign key)
- A Customer can have multiple CustomerDependants (one-to-many)

**Business Rules**:
- Each dependant is linked to exactly one customer
- Date of birth is stored for each dependant
- Multiple dependants can be associated with a single customer

**Notes**:
- Part of the customer profile information used for KYC and product eligibility
- Dependant information may be used for calculating credit limits, insurance products, or family banking packages
- The Customer entity's `dobOfDependents` field aggregates these records

---

### CUSTOMER MESSAGE ENTITY

**Entity Name**: CustomerMessage

**Business Description**: Represents messages sent to or created for customers by bank staff. These messages can be delivered through various transport methods and are used for customer communication, notifications, and record-keeping.

**Key Fields**:
- `mMessageId` (MappedUUID): Unique message identifier
- `user` (MappedLongForeignKey): Foreign key to ResourceUser (deprecated - use customer instead)
- `customer` (MappedLongForeignKey): Foreign key to Customer entity
- `bank` (UUIDString): Bank identifier
- `mFromPerson` (String, max 64): Name of the person sending the message
- `mFromDepartment` (String, max 64): Department sending the message
- `mMessage` (String, max 1024): The message content
- `mTransport` (String, max 64): Transport method (email, SMS, etc.)
- `createdAt` (DateTime): Message creation timestamp
- `updatedAt` (DateTime): Last update timestamp

**Relationships to Other Entities**:
- A CustomerMessage belongs to one Customer (via customer foreign key)
- A CustomerMessage belongs to one Bank (via bank identifier)
- A Customer can have multiple CustomerMessages (one-to-many)

**Business Rules**:
- Message ID must be unique (UUID)
- Messages are ordered by updatedAt in descending order (most recent first)
- Transport field is optional
- Original user field is deprecated in favor of customer field

**Notes**:
- Used for customer communication tracking and audit trail
- Messages can be retrieved by user or by customer
- Supports various transport methods for message delivery
- The system is transitioning from user-based to customer-based message linking

---

### SOCIAL MEDIA ENTITY

**Entity Name**: SocialMedia

**Business Description**: Represents social media handles and contact information for customers. This entity stores the customer's presence on various social media platforms, enabling banks to track and manage customer communication channels and digital presence.

**Key Fields**:
- `mCustomerId` (MappedUUID): The customer this social media handle belongs to
- `mBankId` (UUIDString): The bank identifier
- `mCustomerNumber` (String, max 50): Customer number within the bank
- `mType` (String, max 100): Type of social media platform (e.g., "Twitter", "Facebook", "LinkedIn", "Instagram")
- `mHandle` (String, max 100): The social media handle or username
- `mDateAdded` (DateTime): When this social media handle was added
- `mDateActivated` (DateTime): When this handle was activated/verified

**Relationships to Other Entities**:
- A SocialMedia entry belongs to one Customer (via mCustomerId and mCustomerNumber)
- A SocialMedia entry belongs to one Bank (via mBankId)
- A Customer can have multiple SocialMedia entries (one-to-many)

**Business Rules**:
- Combination of customer and social media type should identify unique social media accounts
- Date added is automatically set when the record is created
- Date activated tracks when the social media handle was verified or activated

**Notes**:
- Enables multi-channel customer communication tracking
- Supports various social media platforms through the flexible Type field
- Useful for customer engagement and marketing purposes
- Can be used for customer verification and identity confirmation

---

### COUNTERPARTY ENTITY

**Entity Name**: Counterparty

**Business Description**: Represents another party involved in a transaction - the person or organization that money is sent to or received from. Counterparties can be explicitly created by users or implicitly generated from transaction data.

**Key Fields**:
- `mCounterPartyId` (UUIDString): Unique counterparty identifier
- `mThisBankId` (String): Perspective bank ID
- `mThisAccountId` (String): Perspective account ID
- `mName` (String, max 255): Counterparty name
- `mOtherBankRoutingScheme` (String, max 50): Other bank's routing scheme
- `mOtherBankRoutingAddress` (String, max 255): Other bank's routing address
- `mOtherAccountRoutingScheme` (String, max 50): Other account routing scheme
- `mOtherAccountRoutingAddress` (String, max 255): Other account routing address (e.g., account number)
- `mOtherAccountSecondaryRoutingScheme` (String, max 50): Secondary routing scheme
- `mOtherAccountSecondaryRoutingAddress` (String, max 255): Secondary routing address (e.g., IBAN)
- `mIsBeneficiary` (Boolean): Whether this is a beneficiary

**Relationships to Other Entities**:
- A Counterparty is associated with a BankAccount (perspective account)
- A Counterparty has one CounterpartyMetadata (one-to-one)
- A Counterparty can have CounterpartyLimits (one-to-many)

**Business Rules**:
- Created explicitly via API or implicitly from transaction data
- Can be searched by IBAN or other routing information

**Notes**:
- Two types: explicit (stored in DB) and implicit (generated from transactions)
- Routing information supports multiple schemes for flexibility

---

### COUNTERPARTY METADATA ENTITY

**Entity Name**: CounterpartyMetadata  

**Business Description**: Additional metadata about a counterparty from the perspective of a specific account, including aliases, location information, corporate data, and other enrichment.

**Key Fields**:
- `counterpartyId` (String): Link to counterparty
- `thisBankId` (String): Perspective bank
- `thisAccountId` (String): Perspective account
- `counterpartyName` (String): Counterparty name
- `publicAlias` (String): Public alias for counterparty
- `privateAlias` (String): Private alias
- `moreInfo` (String): Additional information
- `url` (String): Related URL
- `imageUrl` (String): Image URL
- `openCorporatesUrl` (String): OpenCorporates lookup URL
- `physicalLocation` (CounterpartyWhereTag): Physical location
- `corporateLocation` (CounterpartyWhereTag): Corporate location

**Relationships to Other Entities**:
- Belongs to one Counterparty (via counterpartyId)
- Has physical and corporate locations (embedded CounterpartyWhereTag objects)

**Business Rules**:
- Public alias must be unique for the account
- Auto-generates public alias if not provided (format: "ALIAS_XXXXXX")

**Notes**:
- Metadata is from perspective of specific account
- Supports both personal and corporate information

---

### VIEW ENTITY

**Entity Name**: View (ViewDefinition)

**Business Description**: Defines what information a user can see and do with a bank account. Views control access to account data, implementing a sophisticated permission system. Views can be system-defined (standard views like "owner") or custom-created per account.

**Key Fields**:
- `viewId` (ViewId): Unique view identifier
- `bankId` (String): Bank this view belongs to
- `accountId` (String): Account this view is for
- `name` (String): View name
- `description` (String): View description
- `isPublic` (Boolean): Whether this is a public view
- `isPrivate` (Boolean): Whether this is a private view
- `isSystem` (Boolean): Whether this is a system view (vs custom)
- Various permission fields controlling what can be seen/done

**Relationships to Other Entities**:
- A View is associated with one BankAccount
- A View has many AccountAccess records (which users have access)
- A View has ViewPermissions defining granular permissions

**Business Rules**:
- System views are predefined (owner, public, accountant, etc.)
- Custom views are created per account
- Public views can be disabled system-wide via configuration
- Views control visibility of transaction details, counterparty info, etc.

**Notes**:
- Core security mechanism in OBP
- Implements fine-grained access control
- Supports Open Banking standard view types

---

### ACCOUNT ACCESS ENTITY

**Entity Name**: AccountAccess

**Business Description**: Junction table granting a user access to a specific view on a bank account. This is how the system tracks which users can access which accounts through which views.

**Key Fields**:
- `user_fk` (Long): Foreign key to ResourceUser
- `bank_id` (String): Bank identifier
- `account_id` (String): Account identifier  
- `view_id` (String): View identifier
- `consumer_id` (String): Consumer/application identifier

**Relationships to Other Entities**:
- Links ResourceUser to BankAccount via View
- Many-to-many relationship between Users and Accounts

**Business Rules**:
- Combination of user, bank, account, view should be unique
- Used to determine user permissions on accounts

---

### VIEW PERMISSION ENTITY

**Entity Name**: ViewPermission

**Business Description**: Defines granular permissions for what actions can be performed through a specific view. Each view has associated permissions that control fine-grained access to account operations like seeing balance, creating transactions, adding metadata, etc.

**Key Fields**:
- `permission` (MappedLongForeignKey): Foreign key to the ViewDefinition
- `permission_key` (String, max 255): The permission name/key (e.g., "canSeeBalance", "canCreateTransaction")
- `permission_value` (Boolean): Whether this permission is granted or denied

**Relationships to Other Entities**:
- A ViewPermission belongs to one ViewDefinition (via permission foreign key)
- A View has many ViewPermissions (one-to-many)

**Business Rules**:
- Each permission key within a view should be unique
- Permission keys follow naming convention "canDoSomething"
- System views have predefined permission sets
- Custom views can have customized permission configurations

**Notes**:
- Implements fine-grained access control at the view level
- Common permission keys include: canSeeBalance, canSeeTransactions, canCreateTransaction, canSeeCounterparty, canAddComment, canAddTag
- System-defined permissions vs custom permissions are distinguished by the ViewDefinition's isSystem flag
- This is the mechanism that makes Views flexible and powerful for access control

---

### RESOURCE USER ENTITY

**Entity Name**: ResourceUser (User)

**Business Description**: Represents a system user who can authenticate and access the API. Users can be created through various authentication providers and can be linked to customers.

**Key Fields**:
- `userId_` (String): Unique user identifier (UUID)
- `provider_` (String): Authentication provider (e.g., "localhost", "google")
- `providerId` (String): Provider-specific user ID
- `name_` (String): User's name
- `email` (String): User's email
- `Company` (String): User's company
- `CreatedByConsentId` (String): Consent that created this user
- `CreatedByUserInvitationId` (String): Invitation that created this user
- `IsDeleted` (Boolean): Soft delete flag
- `LastMarketingAgreementSignedDate` (Date): Marketing consent date

**Relationships to Other Entities**:
- A User can have many AccountAccess grants (many-to-many to Accounts via Views)
- A User can have many Entitlements (permissions/roles)
- A User can be linked to Customers via UserCustomerLink
- A User can create Consents, Consumers, etc.

**Business Rules**:
- userId_ must be unique
- Combination of provider and providerId should be unique
- Supports soft deletion via IsDeleted flag

**Notes**:
- Core authentication entity
- Supports multiple authentication providers
- Can be created by consents or invitations

---

### USER REFRESHES ENTITY

**Entity Name**: UserRefreshes

**Business Description**: Tracks when user information was last refreshed from external authentication providers. This entity implements a cache invalidation mechanism to ensure user data stays up-to-date without making excessive calls to external identity providers.

**Key Fields**:
- `mUserId` (UUIDString): User identifier (UUID)
- `createdAt` (DateTime): Initial creation timestamp
- `updatedAt` (DateTime): Last refresh timestamp

**Relationships to Other Entities**:
- One UserRefreshes record per ResourceUser
- Links to ResourceUser via userId

**Business Rules**:
- User ID must be unique (enforced by unique index)
- Refresh interval is configurable via "refresh_user.interval" property (default: 30 minutes)
- User is considered stale if: (lastUpdate + interval) < currentDate
- Automatically creates new record if user not found
- Updates timestamp when user is refreshed

**Notes**:
- Used to optimize external provider calls by caching user data
- Implements time-based cache invalidation strategy
- The needToRefreshUser method checks if refresh is needed based on configured interval
- createOrUpdateRefreshUser either updates existing record or creates new one
- Helps reduce load on external authentication providers while keeping data fresh

---

### AUTH USER ENTITY

**Entity Name**: AuthUser

**Business Description**: Represents authentication credentials for web-based user login to the OBP API management interface. AuthUser extends MegaProtoUser and handles username/password authentication, password reset workflows, email validation, and session management for the web UI. This is separate from ResourceUser which handles API authentication and business relationships.

**Key Fields**:
- `firstName` (String): User's first name
- `lastName` (String): User's last name  
- `email` (String): User's email address (unique)
- `password` (MappedPassword): Hashed password
- `validated` (Boolean): Whether email is validated
- `superUser` (Boolean): Administrator flag
- `user` (ForeignKey): Link to ResourceUser for business operations

**Relationships to Other Entities**:
- One AuthUser links to one ResourceUser (one-to-one relationship)
- AuthUser handles web authentication, ResourceUser handles API authentication and business data
- AuthUser provides login/logout functionality for the web interface
- ResourceUser links to all business entities (accounts, transactions, customers, etc.)

**Business Rules**:
- Email must be unique across all AuthUsers
- Password must meet minimum strength requirements
- Email must be validated before full access
- Supports password reset via email workflow
- Implements account lockout after failed login attempts (via BadLoginAttempt)
- Superuser flag grants administrative privileges

**Notes**:
- Part of the dual user model: AuthUser for web authentication, ResourceUser for API and business
- Extends Lift's MegaProtoUser providing standard authentication features
- Used primarily for OBP API web UI login, not for API client authentication
- API clients use OAuth, Direct Login, or OpenID Connect via ResourceUser
- The user field links AuthUser to ResourceUser for unified user management

---

### CONSUMER ENTITY

**Entity Name**: Consumer

**Business Description**: Represents an API consumer application that uses the Open Bank Project API. Consumers authenticate using OAuth or other methods and are subject to rate limiting.

**Key Fields**:
- `mConsumerId` (UUIDString): Unique consumer identifier
- `mKey` (String): Consumer key for OAuth
- `mSecret` (String): Consumer secret (hashed)
- `mIsActive` (Boolean): Whether consumer is active
- `mName` (String): Application name
- `mAppType` (String): Application type (Web, Mobile, etc.)
- `mDescription` (String): Application description
- `mDeveloperEmail` (String): Developer email
- `mRedirectURL` (String): OAuth redirect URL
- `mCreatedByUserId` (String): User who created this consumer
- `mClientCertificate` (String): PEM certificate for authentication
- `mCompany` (String): Company name
- `mLogoURL` (String): Logo URL

**Relationships to Other Entities**:
- A Consumer belongs to a ResourceUser (creator)
- A Consumer has RateLimiting configurations
- A Consumer makes API calls tracked in Metrics

**Business Rules**:
- Consumer key should be unique
- Active consumers can make API calls
- Rate limits can be configured per consumer

**Notes**:
- Represents third-party applications
- Subject to OAuth authentication
- Rate limiting applied per consumer

---

### USER SCOPE ENTITY

**Entity Name**: UserScope

**Business Description**: Represents authorization scopes granted to users for accessing specific APIs or resources. Scopes define what actions a user is permitted to perform, implementing fine-grained access control for API operations.

**Key Fields**:
- `mUserId` (String): The user this scope belongs to
- `mScopeId` (String): The scope identifier
- `mRoleNames` (String, max 2000): Comma-separated list of role names
- `mBankId` (String): Bank identifier (optional, for bank-specific scopes)

**Relationships to Other Entities**:
- A UserScope belongs to one ResourceUser (via mUserId)
- A UserScope may be bank-specific (via mBankId)

**Business Rules**:
- Combination of UserId, ScopeId, and BankId should be unique
- RoleNames stored as comma-separated string
- Scopes can be system-wide or bank-specific

**Notes**:
- Part of OAuth 2.0 and OpenID Connect implementation
- Enables fine-grained permission management
- Used for API authorization decisions

---

### USER ATTRIBUTE ENTITY

**Entity Name**: UserAttribute

**Business Description**: Represents custom attributes that can be attached to users to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing systems to add custom user properties without modifying core code. Attributes can represent user preferences, additional profile information, or system-specific user settings.

**Key Fields**:
- `UserId` (UUIDString): The user this attribute belongs to
- `UserAttributeId` (MappedUUID): Unique identifier for the attribute
- `Name` (String, max 50): Name/key of the attribute
- `Type` (String, max 50): Attribute type as defined by UserAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `Value` (String, max 255): The attribute value stored as string
- `IsActive` (Boolean): Whether this attribute is currently active (default: true)

**Relationships to Other Entities**:
- A UserAttribute belongs to one ResourceUser (via UserId)
- A ResourceUser can have multiple UserAttributes (one-to-many)

**Business Rules**:
- UserId is indexed for efficient queries
- IsActive defaults to true
- Type determines how the value string should be interpreted
- Only active attributes are typically returned in API responses
- Combination of UserId and Name should identify unique user attributes

**Notes**:
- Enables user-specific customization without code changes
- Supports various data types through the Type field
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing user preferences, settings, or custom profile data

---

### USER INVITATION ENTITY

**Entity Name**: UserInvitation

**Business Description**: Manages invitations sent to potential users to join the banking platform. Tracks the invitation lifecycle from creation through acceptance, including secret links for secure signup.

**Key Fields**:
- `mUserInvitationId` (UUIDString): Unique invitation identifier
- `mBankId` (UUIDString): Bank sending the invitation
- `mFirstName` (String, max 100): Invitee's first name
- `mLastName` (String, max 100): Invitee's last name
- `mEmail` (String, max 100): Invitee's email address
- `mCompany` (String, max 100): Invitee's company
- `mCountry` (String, max 100): Invitee's country
- `mStatus` (String, max 20): Invitation status (CREATED, SENT, ACCEPTED, etc.)
- `mPurpose` (String, max 20): Purpose of invitation (DEVELOPER, CUSTOMER, etc.)
- `mSecretLink` (UUIDString): Secret link for secure invitation acceptance

**Relationships to Other Entities**:
- A UserInvitation belongs to one Bank (via mBankId)
- Accepted invitations create ResourceUser accounts

**Business Rules**:
- Combination of BankId and Email should be unique
- Secret link must be unique and unpredictable (UUID)
- Status tracks invitation workflow lifecycle
- Email must be valid format

**Notes**:
- Secure invitation mechanism using secret links
- Tracks complete invitation workflow
- Prevents unauthorized user creation
- ResourceUser.CreatedByUserInvitationId links back to invitation

---

### USER AGREEMENT ENTITY

**Entity Name**: UserAgreement

**Business Description**: Records user agreements to terms of service, privacy policies, marketing consents, and other legal agreements. Uses hash verification to ensure agreement text hasn't been modified.

**Key Fields**:
- `mUserAgreementId` (UUIDString): Unique agreement identifier
- `mUserId` (UUIDString): User who agreed
- `mAgreementType` (String, max 100): Type of agreement (TERMS_OF_SERVICE, PRIVACY_POLICY, etc.)
- `mAgreementText` (String, max 100000): Full text of the agreement
- `mAgreementHash` (String, max 512): SHA-256 hash of agreement text for verification
- `mDate` (DateTime): When the agreement was accepted

**Relationships to Other Entities**:
- A UserAgreement belongs to one ResourceUser (via mUserId)
- Multiple agreements can exist per user for different types

**Business Rules**:
- Hash is computed from agreement text for integrity verification
- Date records when user accepted the agreement
- Users can have multiple agreements (different types or versions)

**Notes**:
- Critical for legal compliance and audit trail
- Hash ensures agreement text hasn't been tampered with
- Supports versioning through multiple agreements per type
- Provider includes getLatestUserAgreement method for current version

---

### USER INIT ACTION ENTITY

**Entity Name**: UserInitAction

**Business Description**: Tracks initialization actions performed by or for users, such as onboarding steps, setup tasks, or configuration actions. Monitors completion status of required user setup activities.

**Key Fields**:
- `UserId` (UUIDString): User performing the action
- `ActionName` (String, max 100): Name of the initialization action
- `ActionValue` (String, max 100): Value or parameter for the action
- `Success` (Boolean): Whether the action completed successfully

**Relationships to Other Entities**:
- A UserInitAction belongs to one ResourceUser (via UserId)
- Users can have multiple initialization actions

**Business Rules**:
- Combination of UserId, ActionName, and ActionValue is unique
- Success flag tracks completion status
- Used for tracking user onboarding progress

**Notes**:
- Enables tracking of multi-step user initialization workflows
- Can be used for onboarding checklists
- Success flag allows retry logic for failed actions
- Unique index prevents duplicate action tracking

---

### 2. PHYSICAL CARD ENTITY

**Entity Name**: PhysicalCard

**Business Description**: Represents a physical payment card (debit, credit, or other) issued to a customer. The card is linked to a bank account and contains information about its validity, status, and usage permissions.

**Key Fields**:
- `cardId` (String, max 255): Unique identifier for the card (auto-generated UUID)
- `bankId` (String, max 50): The bank that issued this card
- `bankCardNumber` (String, max 50): The card number (PAN)
- `nameOnCard` (String, max 128): Name printed on the card
- `issueNumber` (String, max 10): Card issue number
- `serialNumber` (String, max 50): Card serial number
- `validFrom` (Date): Date the card becomes valid
- `expires` (Date): Expiration date of the card
- `enabled` (Boolean): Whether the card is currently enabled for use
- `cancelled` (Boolean): Whether the card has been cancelled
- `onHotList` (Boolean): Whether the card is on the hot list (blocked/stolen)
- `technology` (String, max 255): Card technology (chip, contactless, magnetic stripe, etc.)
- `networks` (String, max 255): Payment networks (Visa, MasterCard, etc.) - stored as comma-separated
- `allows` (String, max 255): Allowed actions (Credit, Debit, Cash) - stored as comma-separated
- `account` (MappedLongForeignKey): Foreign key to the bank account this card is linked to
- `replacement` (CardReplacementInfo): Information about card replacement
  - `requestedDate` (Date): When replacement was requested
  - `reasonRequested` (CardReplacementReason): Reason for replacement
- `pinResets` (List[PinResetInfo]): History of PIN reset requests
- `collected` (CardCollectionInfo): When/if card was collected
  - `date` (Date): Collection date
- `posted` (CardPostedInfo): When/if card was posted to customer
  - `date` (Date): Posted date
- `cardType` (String, max 255): Type of card (Debit, Credit, Prepaid, etc.)
- `customerId` (String, max 255): The customer who owns this card
- `brand` (String, max 255): Card brand (Visa, MasterCard, Amex, etc.)
- `cvv` (String, max 255): Card CVV (stored as SHA-256 hash)

**Relationships to Other Entities**:
- A Card belongs to one Bank (via bankId)
- A Card belongs to one Customer (via customerId)
- A Card is linked to one BankAccount (via account foreign key)
- A Card has multiple PinReset records (one-to-many relationship)
- A Card can have multiple CardAction records (one-to-many relationship)

**Business Rules**:
- Combination of bankId, bankCardNumber, and issueNumber must be unique
- Card must be associated with a valid bank account (foreign key constraint)
- CVV is stored as a SHA-256 hash for security
- Networks and allows are stored as comma-separated strings

**Notes**:
- The file is named "PhisicalCard" but should be "PhysicalCard" (typo in original code)
- PIN resets are tracked in a separate PinReset table
- Card actions are tracked in a separate CardAction table

---

### 3. PIN RESET ENTITY

**Entity Name**: PinReset

**Business Description**: Tracks PIN reset requests and actions for physical cards, maintaining a history of when and why PINs were reset.

**Key Fields**:
- `id` (Long): Primary key
- `card` (MappedLongForeignKey): Foreign key to the physical card
- `mReplacementDate` (Date): When the PIN reset was requested
- `mReplacementReason` (String, max 255): Reason for the PIN reset

**Relationships to Other Entities**:
- A PinReset belongs to one PhysicalCard (via card foreign key)

**Business Rules**:
- Each reset is linked to a specific card
- Reason must be a valid PinResetReason enum value

---

### CARD ACTION ENTITY

**Entity Name**: CardAction

**Business Description**: Records specific actions or permissions associated with a physical card.

**Key Fields**:
- `id` (Long): Primary key
- `post` (MappedLongForeignKey): Foreign key to the physical card
- `cardAction` (String, max 140): The action type or permission

**Relationships to Other Entities**:
- A CardAction belongs to one PhysicalCard (via post foreign key)

---
### CARD ATTRIBUTE ENTITY

**Entity Name**: CardAttribute

**Business Description**: Represents custom attributes that can be attached to physical cards to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing cards to have custom properties defined dynamically. Attributes can represent card-specific features, usage restrictions, loyalty programs, or any custom metadata needed for card management and personalization.

**Key Fields**:
- `mBankId` (UUIDString): The bank this attribute belongs to
- `mCardId` (UUIDString): The card this attribute belongs to
- `mCardAttributeId` (MappedUUID): Unique identifier for the attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by CardAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string

**Relationships to Other Entities**:
- A CardAttribute belongs to one PhysicalCard (via mCardId)
- A CardAttribute belongs to one Bank (via mBankId)
- Multiple attributes can be attached to the same card
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and cardId is indexed for efficient queries
- CardAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Attributes can be filtered based on View permissions
- Multiple attributes with the same name are allowed for the same card

**Notes**:
- Enables card-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing card preferences, spending limits, rewards programs, and travel notifications

---



### TRANSACTION PROCESSING

---

### TRANSACTION ENTITY

**Entity Name**: Transaction

**Business Description**: Represents a financial transaction on a bank account - money moving in or out. Contains core transaction details like amount, date, description, and links to counterparty information.

**Key Fields**:
- `bank` (UUIDString): Bank identifier
- `account` (AccountIdString): Account identifier
- `transactionId` (String): Unique transaction identifier
- `amount` (Long): Transaction amount in smallest currency units
- `transactionType` (String): Transaction type/category
- `currency` (String): Currency code
- `description` (String): Transaction description
- `startDate` (DateTime): Transaction start date
- `finishDate` (DateTime): Transaction finish date
- `balance` (Long): Account balance after transaction
- `counterpartyAccountHolder` (String): Counterparty account holder name
- `counterpartyAccount` (String): Counterparty account number
- `counterpartyNationalId` (String): Counterparty national ID

**Relationships to Other Entities**:
- A Transaction belongs to one BankAccount
- A Transaction can have TransactionAttributes
- A Transaction can have TransactionImages
- A Transaction can have Tags
- A Transaction can have Comments
- A Transaction can have a Narrative
- A Transaction can have a WhereTag (location)
- A Transaction links to Counterparty information

**Business Rules**:
- Transaction amount and balance stored in smallest currency units
- startDate and finishDate track transaction lifecycle

**Notes**:
- Rich metadata support through related entities
- Core financial record in the system

---

### TRANSACTION REQUEST ENTITY

**Entity Name**: TransactionRequest

**Business Description**: Represents a request to initiate a transaction. Transaction requests go through a workflow (initiated, pending, completed, failed) and may require challenges for security.

**Key Fields**:
- `mTransactionRequestId` (UUIDString): Unique request identifier
- `mType` (String): Request type (TRANSFER, PAYMENT, etc.)
- `mFrom_BankId` (String): Source bank
- `mFrom_AccountId` (String): Source account
- `mTo_BankId` (String): Destination bank
- `mTo_AccountId` (String): Destination account
- `mAmount_Value` (BigDecimal): Transaction amount
- `mAmount_Currency` (String): Currency
- `mStatus` (String): Request status (INITIATED, COMPLETED, etc.)
- `mCharge_Policy` (String): Who pays fees
- `mCharge_Summary` (String): Fee summary

**Relationships to Other Entities**:
- Links source and destination BankAccounts
- Can have TransactionChallenges for security
- Can have TransactionRequestAttributes
- Can have TransactionRequestReasons

**Business Rules**:
- Status must be valid (INITIATED, PENDING, COMPLETED, FAILED, etc.)
- May require challenge completion before execution

**Notes**:
- Supports PSD2/Open Banking payment initiation
- Workflow-based transaction processing

---
### TRANSACTION ATTRIBUTE ENTITY

**Entity Name**: TransactionAttribute

**Business Description**: Represents custom attributes that can be attached to transactions to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing transactions to have custom properties defined dynamically. Attributes can represent transaction classifications, regulatory tags, internal references, or any custom metadata needed for transaction management and compliance.

**Key Fields**:
- `mBankId` (UUIDString): The bank this attribute belongs to
- `mTransactionId` (UUIDString): The transaction this attribute belongs to
- `mTransactionAttributeId` (MappedUUID): Unique identifier for the attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by TransactionAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string

**Relationships to Other Entities**:
- A TransactionAttribute belongs to one Transaction (via mTransactionId)
- A TransactionAttribute belongs to one Bank (via mBankId)
- Multiple attributes can be attached to the same transaction
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and transactionId is indexed for efficient queries
- TransactionAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Attributes can be filtered based on View permissions
- Multiple attributes with the same name are allowed for the same transaction

**Notes**:
- Enables transaction-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Supports querying transactions by attribute name/value pairs
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing transaction categories, internal codes, regulatory flags, and custom metadata

---

### TRANSACTION REQUEST ATTRIBUTE ENTITY

**Entity Name**: TransactionRequestAttribute

**Business Description**: Represents custom attributes that can be attached to transaction requests to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing transaction requests to have custom properties defined dynamically. Attributes can represent request-specific information, workflow metadata, approval chains, or any custom metadata needed for payment initiation and transaction request management.

**Key Fields**:
- `mBankId` (UUIDString): The bank this attribute belongs to
- `mTransactionRequestId` (UUIDString): The transaction request this attribute belongs to
- `mTransactionRequestAttributeId` (MappedUUID): Unique identifier for the attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by TransactionRequestAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string

**Relationships to Other Entities**:
- A TransactionRequestAttribute belongs to one TransactionRequest (via mTransactionRequestId)
- A TransactionRequestAttribute belongs to one Bank (via mBankId)
- Multiple attributes can be attached to the same transaction request
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and transactionRequestId is indexed for efficient queries
- TransactionRequestAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Attributes can be filtered based on View permissions
- Multiple attributes with the same name are allowed for the same transaction request

**Notes**:
- Enables transaction request-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing approval metadata, workflow state, PSD2/Open Banking specific fields, and custom metadata
- Supports payment initiation use cases with custom business logic

---



### PAYMENT OPERATIONS

---

### DIRECT DEBIT ENTITY

**Entity Name**: DirectDebit

**Business Description**: Represents an authorization for recurring debits from an account - automatic payments that pull money regularly.

**Key Fields**:
- `directDebitId` (UUIDString): Unique identifier
- `bankId` (String): Bank identifier
- `accountId` (String): Account identifier
- `customerId` (String): Customer identifier
- `userId` (String): User identifier
- `counterpartyId` (String): Counterparty identifier
- `dateSigned` (Date): When authorization was signed
- `dateCancelled` (Date): When cancelled (if applicable)
- `dateStarts` (Date): When direct debit starts
- `dateExpires` (Date): When direct debit expires
- `active` (Boolean): Whether currently active

**Relationships to Other Entities**:
- Links BankAccount, Customer, User, and Counterparty

**Business Rules**:
- Requires signed authorization
- Can have start and end dates
- Can be cancelled

---

### STANDING ORDER ENTITY

**Entity Name**: StandingOrder

**Business Description**: Represents a standing order for regular payments from an account - automatic payments that push money on a schedule.

**Key Fields**:
- `standingOrderId` (UUIDString): Unique identifier
- `bankId`, `accountId`, `customerId`, `userId`: Entity identifiers
- `counterpartyId` (String): Payment recipient
- `amountValue` (Long): Payment amount (smallest units)
- `amountCurrency` (String): Currency
- `whenFrequency` (String): Payment frequency (monthly, weekly, etc.)
- `whenDetail` (String): Frequency details
- `dateSigned`, `dateStarts`, `dateExpires`: Lifecycle dates
- `active` (Boolean): Whether active

**Relationships to Other Entities**:
- Links BankAccount, Customer, User, Counterparty

**Business Rules**:
- Must have frequency specification
- Requires signed authorization

---

### KYC & COMPLIANCE

---

### KYC CHECK ENTITY

**Entity Name**: KycCheck

**Business Description**: Know Your Customer verification check performed on a customer. Records the verification details, status, and results.

**Key Fields**:
- `mBankId` (String): Bank performing check
- `mCustomerId` (String): Customer being checked
- `mId` (String): Unique check identifier
- `mCustomerNumber` (String): Customer number
- `mDate` (Date): Check date
- `mHow` (String): Verification method
- `mStaffUserId` (String): Staff who performed check
- `mStaffName` (String): Staff name
- `mSatisfied` (Boolean): Whether check passed
- `mComments` (String): Additional comments

**Relationships to Other Entities**:
- Belongs to Customer
- Performed by User (staff)

**Business Rules**:
- Records compliance verification
- Satisfied status indicates pass/fail

---

### KYC DOCUMENT ENTITY

**Entity Name**: KycDocument

**Business Description**: Documents submitted for KYC verification (passport, driver's license, utility bills, etc.)

**Key Fields**:
- `mBankId`, `mCustomerId`: Entity identifiers
- `mId` (String): Document identifier
- `mCustomerNumber` (String): Customer number
- `mType` (String): Document type
- `mNumber` (String): Document number
- `mIssueDate` (Date): Issue date
- `mIssuePlace` (String): Issue location
- `mExpiryDate` (Date): Expiry date

**Relationships to Other Entities**:
- Belongs to Customer

**Business Rules**:
- Different document types for different jurisdictions
- Tracks expiry for compliance

---

### KYC STATUS ENTITY

**Entity Name**: KycStatus

**Business Description**: Overall KYC verification status for a customer.

**Key Fields**:
- `mCustomerId` (String): Customer identifier
- `mOk` (Boolean): Overall KYC status
- `mDate` (Date): Status date

**Relationships to Other Entities**:
- One per Customer

**Business Rules**:
- Simplified status tracking
- Boolean for pass/fail

---

### TAX RESIDENCE ENTITY

**Entity Name**: TaxResidence

**Business Description**: Records a customer's tax residence information for compliance and reporting purposes.

**Key Fields**:
- `mCustomerId` (String): Customer identifier
- `mDomain` (String): Tax domain (e.g., country code)
- `mTaxNumber` (String): Tax identification number

**Relationships to Other Entities**:
- Belongs to Customer

**Business Rules**:
- Customers can have multiple tax residences
- Required for international banking compliance

---

### CONSENT ENTITY

**Entity Name**: Consent

**Business Description**: Records customer consent for data sharing with third parties. Critical for Open Banking / PSD2 compliance.

**Key Fields**:
- `mConsentId` (UUIDString): Unique consent identifier
- `mUserId` (String): User granting consent
- `mConsumerId` (String): Consumer application receiving access
- `mStatus` (String): Consent status (ACCEPTED, REVOKED, etc.)
- `mConsentType` (String): Type of consent
- `mJsonWebToken` (String): JWT for authentication
- Various time-based fields for validity period

**Relationships to Other Entities**:
- Links User and Consumer
- Can create Users via consent

**Business Rules**:
- Time-bounded consent validity
- Can be revoked
- Status tracking through lifecycle

---

### EXPECTED CHALLENGE ANSWER ENTITY

**Entity Name**: ExpectedChallengeAnswer

**Business Description**: Stores expected answers for transaction challenges in the challenge-response authentication system. When a transaction requires additional verification, the system generates a challenge and stores the expected answer(s) that the user must provide to authorize the transaction. This is a critical security mechanism for high-value or sensitive transactions.

**Key Fields**:
- `mChallengeId` (String, max 100): Unique identifier for the challenge this answer belongs to
- `mSalt` (String, max 100): Cryptographic salt used for answer hashing
- `mExpectedAnswer` (String, max 100): The expected answer (typically hashed for security)
- `mExpectedUserId` (String, max 100): The user ID expected to provide this answer
- `mSuccessful` (Boolean): Whether the challenge was successfully answered
- `createdAt` (DateTime): When the expected answer was created
- `updatedAt` (DateTime): Last update timestamp

**Relationships to Other Entities**:
- An ExpectedChallengeAnswer belongs to one TransactionChallenge (via mChallengeId)
- An ExpectedChallengeAnswer is associated with one ResourceUser (via mExpectedUserId)

**Business Rules**:
- Challenge ID uniquely identifies the expected answer
- Salt is used to securely hash and verify the answer
- Successful flag tracks whether the challenge was completed
- Expected answer is typically stored hashed for security
- Timestamps track answer creation and updates for audit purposes

**Notes**:
- Part of the transaction authorization security system
- Implements challenge-response authentication for sensitive operations
- Answer verification uses salt and hashing for security
- Used in conjunction with TransactionChallenge and TransactionRequest entities
- Critical for PSD2 Strong Customer Authentication (SCA) requirements

---

### REGULATED ENTITY ENTITY

**Entity Name**: RegulatedEntity

**Business Description**: Represents a regulated financial entity with certificate-based identity, used for TPP (Third Party Provider) registration in Open Banking.

**Key Fields**:
- `entityId` (UUIDString): Unique identifier
- `certificateAuthorityCaOwnerId` (String): Certificate authority ID
- `entityCertificatePublicKey` (Text): Public key certificate (PEM)
- `entityName` (String): Entity name
- `entityCode` (String): Entity code
- `entityType` (String): Type (Bank, Payment Institution, etc.)
- `entityAddress`, `entityTownCity`, `entityPostCode`, `entityCountry`: Address fields
- `entityWebSite` (String): Website URL
- `services` (Text): Services provided

**Relationships to Other Entities**:
- Can have RegulatedEntityAttributes

**Business Rules**:
- Certificate-based authentication
- Required for PSD2 TPP registration

---

### SYSTEM MANAGEMENT

---

### METRIC ENTITY

**Entity Name**: Metric

**Business Description**: Records every API call made to the system for analytics, monitoring, and billing purposes. Comprehensive audit trail of API usage.

**Key Fields**:
- `userId` (String): User making the call
- `url` (String): API endpoint URL
- `date` (DateTime): Call timestamp
- `duration` (Long): Call duration in milliseconds
- `userName`, `appName`, `developerEmail`: User/app identification
- `consumerId` (String): Consumer application ID
- `implementedByPartialFunction` (String): Code function handling request
- `implementedInVersion` (String): API version
- `verb` (String): HTTP verb (GET, POST, etc.)
- `httpCode` (Int): HTTP response code
- `correlationId` (String): Request correlation ID
- `responseBody` (String): Response content
- `sourceIp`, `targetIp`: Network information

**Relationships to Other Entities**:
- Links to User, Consumer
- References API endpoints

**Business Rules**:
- Every API call generates a metric
- Archived periodically for performance
- Used for analytics and billing

**Notes**:
- Critical for monitoring and analytics
- Supports aggregate queries for dashboards
- Can identify top APIs and consumers

---

### METRIC ARCHIVE ENTITY

**Entity Name**: MetricArchive

**Business Description**: Stores archived API metrics that have been moved from the active Metric table for performance optimization. This allows the system to maintain historical API usage data while keeping the active metrics table performant.

**Key Fields**:
- `userId` (String): User making the call
- `url` (String): API endpoint URL
- `date` (DateTime): Call timestamp
- `duration` (Long): Call duration in milliseconds
- `userName`, `appName`, `developerEmail`: User/app identification
- `consumerId` (String): Consumer application ID
- `implementedByPartialFunction` (String): Code function handling request
- `implementedInVersion` (String): API version
- `verb` (String): HTTP verb (GET, POST, etc.)
- `httpCode` (Int): HTTP response code
- `correlationId` (String): Request correlation ID
- `responseBody` (String): Response content
- `sourceIp`, `targetIp`: Network information

**Relationships to Other Entities**:
- Same structure and relationships as Metric entity
- Links to User, Consumer
- References API endpoints

**Business Rules**:
- Contains older metrics that have been archived from the active Metric table
- Periodically populated from Metric table based on retention policies
- Used for historical analysis and long-term trend reporting

**Notes**:
- Enables performance optimization by keeping active Metric table smaller
- Critical for long-term analytics and compliance reporting
- Supports historical API usage analysis without impacting real-time performance
- Archiving typically done based on date/age of metrics

---

### RATE LIMITING ENTITY

**Entity Name**: RateLimiting

**Business Description**: Defines API call rate limits for consumers at various time granularities. Prevents API abuse and ensures fair usage.

**Key Fields**:
- `rateLimitingId` (UUIDString): Unique identifier
- `consumerId` (String): Consumer being limited
- `bankId` (String, optional): Bank-specific limit
- `apiVersion` (String, optional): Version-specific limit
- `apiName` (String, optional): Endpoint-specific limit
- `perSecondCallLimit` (Long): Calls per second
- `perMinuteCallLimit` (Long): Calls per minute
- `perHourCallLimit` (Long): Calls per hour
- `perDayCallLimit` (Long): Calls per day
- `perWeekCallLimit` (Long): Calls per week
- `perMonthCallLimit` (Long): Calls per month
- `fromDate`, `toDate` (DateTime): Validity period

**Relationships to Other Entities**:
- Associated with Consumer
- Can be bank/version/endpoint specific

**Business Rules**:
- Value of -1 means no limit
- More specific limits override general ones
- Hierarchy: endpoint > version > consumer > global
- Time-bounded validity

**Notes**:
- Essential for API management
- Prevents DOS attacks
- Supports tiered service levels

---

### ATTRIBUTE DEFINITION ENTITY

**Entity Name**: AttributeDefinition

**Business Description**: Defines the schema for dynamic attributes that can be attached to various entity types. Allows banks to extend core entities with custom fields without code changes.

**Key Fields**:
- `attributeDefinitionId` (UUIDString): Unique identifier
- `bankId` (String): Bank defining this attribute
- `name` (String): Attribute name
- `category` (String): Entity category (Account, Product, Customer, Transaction, Card, etc.)
- `typeOfValue` (String): Data type (STRING, INTEGER, DATE, etc.)
- `description` (String): Attribute description
- `alias` (String): Alias name
- `canBeSeenOnViews` (List[String]): Which views can see this attribute
- `isActive` (Boolean): Whether attribute is active

**Relationships to Other Entities**:
- Defines schema for:
  - AccountAttributes
  - ProductAttributes
  - CustomerAttributes
  - TransactionAttributes
  - CardAttributes
  - And more...

**Business Rules**:
- Combination of bankId, name, and category must be unique
- Controls visibility via views
- Type enforcement for values

**Notes**:
- Key extensibility mechanism
- Allows bank-specific customization
- Supports compliance requirements

---

### WEBHOOK ENTITY

**Entity Name**: AccountWebhook

**Business Description**: Configures webhooks that trigger HTTP callbacks when specific events occur on bank accounts (e.g., new transactions, balance changes).

**Key Fields**:
- `accountWebhookId` (String): Unique identifier
- `bankId`, `accountId`: Target account
- `userId`: User who created webhook
- `triggerName` (String): Event trigger type
- `url` (String): Callback URL
- `httpMethod` (String): HTTP method for callback
- `httpProtocol` (String): HTTP protocol version
- `isActive` (Boolean): Whether webhook is active

**Relationships to Other Entities**:
- Associated with BankAccount
- Created by User

**Business Rules**:
- Triggers on account events
- HTTP callback to configured URL
- Can be activated/deactivated

---

### CRM EVENT ENTITY

**Entity Name**: CrmEvent

**Business Description**: Customer relationship management events tracking customer interactions, meetings, and communications.

**Key Fields**:
- `mBankId`, `mUserId`, `mCrmEventId`: Entity identifiers
- `mCategory` (String): Event category
- `mDetail` (String, max 1024): Event details
- `mChannel` (String): Communication channel
- `mScheduledDate` (DateTime): Scheduled date
- `mActualDate` (DateTime): Actual date
- `mResult` (String): Event result
- `mCustomerName`, `mCustomerNumber`: Customer identification

**Relationships to Other Entities**:
- Links Bank, User (customer), User (staff)

**Business Rules**:
- Tracks scheduled vs actual dates
- Records outcome/result

---

### MEETING ENTITY

**Entity Name**: Meeting

**Business Description**: Represents scheduled meetings between bank staff and customers, including meeting details, invitees, keys for secure access, and session management. Supports both customer-initiated and staff-initiated meetings.

**Key Fields**:
- `mMeetingId` (UUIDString): Unique meeting identifier
- `mBankId` (UUIDString): Bank hosting the meeting
- `mProviderId` (String, max 128): Authentication provider
- `mPurposeId` (String, max 20): Purpose of the meeting
- `mWhen` (DateTime): Scheduled meeting time
- `mSessionId` (UUIDString): Session identifier for the meeting
- `mCustomerUserId` (UUIDString): Customer user attending the meeting
- `mStaffUserId` (UUIDString): Staff user conducting the meeting
- `mCustomerToken` (String, max 32): Token for customer to join meeting
- `mStaffToken` (String, max 32): Token for staff to join meeting
- `mMeetingOwnerUserId` (UUIDString): User who created/owns the meeting
- `mKeys` (String, max 8192): Encrypted keys or meeting access credentials
- `mPresent` (MeetingPresent): Present participants
  - `staffUserId` (String): Staff member present
  - `customerUserId` (String): Customer present
- `mInvitees` (List[Invitee]): List of invited participants
  - `contactNumber` (String): Contact number of invitee
  - `status` (String): Invitation status

**Relationships to Other Entities**:
- A Meeting belongs to one Bank (via mBankId)
- A Meeting has one customer User (via mCustomerUserId)
- A Meeting has one staff User (via mStaffUserId)
- A Meeting is owned by one User (via mMeetingOwnerUserId)

**Business Rules**:
- MeetingId must be unique
- SessionId links related meeting sessions
- Separate tokens for customer and staff access provide security
- Invitees stored as JSON list in separate field
- Keys field can store encrypted meeting credentials

**Notes**:
- Supports video/phone/in-person meetings between customers and bank staff
- Token-based access control for meeting participants
- Can track multiple invitees and their invitation status
- Provider field supports multiple authentication systems
- Present field tracks actual attendance vs invited participants

---

### TECHNICAL INFRASTRUCTURE

---

### DYNAMIC MESSAGE DOC ENTITY

**Entity Name**: DynamicMessageDoc

**Business Description**: Documents message formats and schemas for connector communications. This entity enables dynamic configuration of connector messages, supporting multiple message formats (JSON, Avro) and documenting the structure of inbound/outbound messages between the OBP-API and core banking connectors.

**Key Fields**:
- `BankId` (String, max 255): Optional bank identifier for bank-specific configurations
- `DynamicMessageDocId` (UUID): Unique identifier
- `Process` (String, max 255): Process name this message doc applies to (unique)
- `MessageFormat` (String, max 255): Format type (e.g., "json", "avro")
- `Description` (String, max 255): Human-readable description
- `OutboundTopic` (String, max 255): Kafka topic for outbound messages
- `InboundTopic` (String, max 255): Kafka topic for inbound messages
- `ExampleOutboundMessage` (Text): JSON example of outbound message
- `ExampleInboundMessage` (Text): JSON example of inbound message
- `OutboundAvroSchema` (Text): Avro schema for outbound messages
- `InboundAvroSchema` (Text): Avro schema for inbound messages
- `AdapterImplementation` (String, max 255): Adapter implementation identifier
- `MethodBody` (Text): Optional method body code
- `Lang` (String, max 50): Programming language

**Relationships to Other Entities**:
- Links to ConnectorMethod for connector implementations
- Referenced by MethodRouting for message processing

**Business Rules**:
- Process name must be unique across all message docs
- Supports both bank-specific and global configurations
- TTL cache of 40 seconds for performance

**Notes**:
- Critical for dynamic connector configuration without code changes
- Supports multiple message formats for flexibility
- Examples help developers understand message structure

---

### DYNAMIC RESOURCE DOC ENTITY

**Entity Name**: DynamicResourceDoc

**Business Description**: Enables runtime creation and documentation of API endpoints without code deployment. This powerful feature allows administrators to define new REST endpoints dynamically, complete with request/response documentation, validation, and authorization rules.

**Key Fields**:
- `BankId` (String, max 255): Optional bank identifier for bank-specific endpoints
- `DynamicResourceDocId` (UUID): Unique identifier
- `PartialFunctionName` (String, max 255): Function name for the endpoint
- `RequestVerb` (String, max 255): HTTP verb (GET, POST, PUT, DELETE)
- `RequestUrl` (String, max 255): URL pattern for the endpoint
- `Summary` (String, max 255): Short endpoint description
- `Description` (String, max 255): Detailed endpoint description
- `ExampleRequestBody` (String, max 255): JSON example request
- `SuccessResponseBody` (String, max 255): JSON example success response
- `ErrorResponseBodies` (String, max 255): JSON example error responses
- `Tags` (String, max 255): Endpoint tags for categorization
- `Roles` (String, max 255): Required roles for access
- `MethodBody` (Text): Scala code implementing the endpoint logic

**Relationships to Other Entities**:
- Can enforce Entitlement requirements via Roles
- Integrates with API documentation/explorer
- May interact with DynamicEntity for data access

**Business Rules**:
- Combination of RequestUrl and RequestVerb must be unique
- TTL cache of 40 seconds
- Supports bank-specific and global endpoints

**Notes**:
- Extremely powerful for rapid API development
- Allows banks to customize APIs without OBP-API code changes
- Security validated through Roles field

---

### DYNAMIC ENTITY ENTITY

**Entity Name**: DynamicEntity

**Business Description**: Defines custom entity schemas that can be created at runtime without code changes. This powerful feature allows banks to define their own custom business entities with their own fields and data structures, stored as JSON metadata. Each dynamic entity gets automatic CRUD endpoints generated for managing instances of that entity type.

**Key Fields**:
- `DynamicEntityId` (UUID): Unique identifier for this entity definition
- `EntityName` (String, max 255): Name of the custom entity type
- `MetadataJson` (Text): JSON schema defining the entity's structure, fields, and validation rules
- `UserId` (String, max 255): User who created this entity definition
- `BankId` (String, max 255): Optional bank identifier (null for system-level entities)
- `HasPersonalEntity` (Boolean): Whether instances of this entity contain personal user data

**Relationships to Other Entities**:
- A DynamicEntity can be system-wide (BankId = null) or bank-specific
- A DynamicEntity is created by a ResourceUser (via UserId)
- DynamicData instances are created based on DynamicEntity definitions
- DynamicEndpoint entities may reference DynamicEntity for data access

**Business Rules**:
- DynamicEntityId must be unique
- EntityName must be unique within a bank (or globally for system entities)
- MetadataJson defines the schema for data validation
- Bank-specific entities (BankId not null) are isolated per bank
- System-level entities (BankId = null) are available across all banks
- HasPersonalEntity determines access control for data instances

**Notes**:
- Core extensibility feature enabling runtime entity creation
- Eliminates need for database migrations when adding new entity types
- MetadataJson follows JSON Schema standard for validation
- Automatic CRUD endpoint generation for each entity
- Supports both shared (system) and isolated (bank-specific) entities
- Personal entities enforce user-level access control on data

---

### DYNAMIC DATA ENTITY

**Entity Name**: DynamicData

**Business Description**: Stores actual data instances for custom entity types defined by DynamicEntity. This is the runtime data storage for dynamically created entities, with each record containing JSON-formatted data conforming to its entity's schema definition.

**Key Fields**:
- `DynamicDataId` (UUID): Unique identifier for this data record
- `DynamicEntityName` (String, max 255): Name of the entity type this data belongs to
- `DataJson` (Text): JSON data for this record, conforming to entity schema
- `BankId` (String, max 255): Optional bank identifier (null for system-level data)
- `UserId` (String, max 255): Optional user identifier (for personal entity instances)
- `IsPersonalEntity` (Boolean): Whether this is personal user data requiring access control

**Relationships to Other Entities**:
- Each DynamicData record is an instance of a DynamicEntity type
- Personal data (IsPersonalEntity = true) belongs to a specific ResourceUser
- Bank-specific data is isolated per Bank
- Referenced by DynamicEndpoint for API operations

**Business Rules**:
- DynamicDataId must be unique
- DynamicEntityName must reference an existing DynamicEntity
- DataJson must conform to the schema defined in the DynamicEntity's MetadataJson
- Personal entities require UserId and enforce user-level access control
- Bank-specific data is isolated (BankId must match DynamicEntity's BankId)
- System-level data (BankId = null) is globally accessible
- Query access depends on IsPersonalEntity flag:
  - Personal: User can only access their own data
  - Non-personal: Data shared within bank or system scope

**Notes**:
- Flexible JSON storage enables arbitrary data structures
- Access control automatically enforced based on IsPersonalEntity flag
- Supports both multi-tenant (bank-specific) and global (system) data
- Schema validation performed against DynamicEntity definition
- Critical for custom banking products and services
- Enables banks to store custom data without OBP-API code changes

---

### DYNAMIC ENDPOINT ENTITY

**Entity Name**: DynamicEndpoint

**Business Description**: Enables runtime creation of custom API endpoints without code deployment. Banks and administrators can define new REST endpoints with custom Swagger documentation, including complete endpoint specifications with request/response examples and role-based access control.

**Key Fields**:
- `DynamicEndpointId` (UUID): Unique identifier for this endpoint definition
- `UserId` (String, max 255): User who created this endpoint
- `BankId` (String, max 255): Optional bank identifier (null for system-level endpoints)
- `SwaggerString` (Text): Complete Swagger/OpenAPI specification for the endpoint including:
  - HTTP method (GET, POST, PUT, DELETE)
  - URL path pattern
  - Request/response schemas
  - Authentication requirements
  - Role-based access control
  - Example requests and responses
  - Documentation and descriptions

**Relationships to Other Entities**:
- A DynamicEndpoint is created by a ResourceUser (via UserId)
- Can be system-wide (BankId = null) or bank-specific
- May interact with DynamicEntity/DynamicData for custom data operations
- Enforces Entitlement requirements through role specifications
- Integrated with API Explorer for documentation

**Business Rules**:
- DynamicEndpointId must be unique
- SwaggerString must be valid Swagger/OpenAPI specification
- Bank-specific endpoints only accessible within that bank
- System-level endpoints available across all banks
- Endpoints cached for 40 seconds (TTL) for performance
- Role-based access control enforced at runtime

**Notes**:
- Extremely powerful for rapid API customization
- Allows banks to extend OBP-API without code changes
- SwaggerString provides complete endpoint definition
- Security validated through role requirements in Swagger spec
- Integrated with API documentation and explorer
- Supports custom business logic implementation
- Enables bank-specific API extensions while maintaining core API stability

---

### CONNECTOR METHOD ENTITY

**Entity Name**: ConnectorMethod

**Business Description**: Stores custom connector method implementations that can be dynamically loaded at runtime. This allows banks to implement custom connector logic without modifying the OBP-API codebase, supporting various programming languages.

**Key Fields**:
- `ConnectorMethodId` (UUID): Unique identifier
- `MethodName` (String, max 255): Unique method name
- `MethodBody` (Text): Source code of the method implementation
- `Lang` (String, max 50): Programming language (default: "Scala")

**Relationships to Other Entities**:
- Referenced by MethodRouting for connector routing
- Used by DynamicMessageDoc for connector communications
- Integrates with connector framework

**Business Rules**:
- MethodName must be unique
- ConnectorMethodId must be unique
- Supports multiple programming languages
- TTL cache of 40 seconds

**Notes**:
- Enables dynamic connector customization
- Critical for extending connector capabilities
- Language support allows flexibility in implementation

---

### JSON SCHEMA VALIDATION ENTITY

**Entity Name**: JsonSchemaValidation

**Business Description**: Stores JSON schemas for validating API endpoint requests and responses. This entity enables runtime validation of JSON payloads against predefined schemas, ensuring data quality and API contract compliance.

**Key Fields**:
- `OperationId` (String, max 200): Unique operation identifier (API endpoint identifier)
- `JsonSchema` (Text): JSON Schema definition

**Relationships to Other Entities**:
- Referenced by API endpoints for request/response validation
- May be used by DynamicResourceDoc for dynamic endpoint validation

**Business Rules**:
- OperationId must be unique
- Schema must be valid JSON Schema format
- Applied at runtime to validate payloads

**Notes**:
- Improves API reliability through schema validation
- Prevents invalid data from entering the system
- Supports OpenAPI/JSON Schema standards

---

### ACCOUNT ID MAPPING ENTITY

**Entity Name**: AccountIdMapping

**Business Description**: Internal mapping table that abstracts account IDs between the OBP-API layer and core banking systems. This enables the API to use consistent UUIDs while mapping to various core banking account identifier formats.

**Key Fields**:
- `mAccountId` (UUID): OBP-API account identifier
- `mAccountPlainTextReference` (String, max 255): Core banking account reference
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Maps to BankAccount entities
- Used by connector layer for ID translation

**Business Rules**:
- mAccountId must be unique
- Combination of mAccountId and mAccountPlainTextReference must be unique
- Immutable once created (no updates to mapping)

**Notes**:
- Critical for connector abstraction layer
- Enables support for multiple core banking systems
- Protects internal account IDs from external exposure

---

### TRANSACTION ID MAPPING ENTITY

**Entity Name**: TransactionIdMapping

**Business Description**: Internal mapping table that abstracts transaction IDs between the OBP-API layer and core banking systems, similar to AccountIdMapping but for transactions.

**Key Fields**:
- `TransactionId` (UUID): OBP-API transaction identifier
- `TransactionPlainTextReference` (String, max 255): Core banking transaction reference
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Maps to Transaction entities
- Used by connector layer for ID translation

**Business Rules**:
- TransactionId must be unique
- Combination of TransactionId and TransactionPlainTextReference must be unique
- Immutable once created

**Notes**:
- Enables consistent transaction referencing across connectors
- Protects internal transaction IDs
- Critical for transaction history lookups

---

### CUSTOMER ID MAPPING ENTITY

**Entity Name**: CustomerIdMapping

**Business Description**: Internal mapping table that abstracts customer IDs between the OBP-API layer and core banking systems, similar to AccountIdMapping but for customers.

**Key Fields**:
- `mCustomerId` (UUID): OBP-API customer identifier
- `mCustomerPlainTextReference` (String, max 255): Core banking customer reference
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Maps to Customer entities
- Used by connector layer for ID translation

**Business Rules**:
- mCustomerId must be unique
- Combination of mCustomerId and mCustomerPlainTextReference must be unique
- Immutable once created

**Notes**:
- Protects customer PII by abstracting core banking IDs
- Enables multi-connector support
- Critical for customer data portability

---

### MIGRATION SCRIPT LOG ENTITY

**Entity Name**: MigrationScriptLog

**Business Description**: Tracks execution history of database migration scripts, ensuring migrations run only once and providing audit trail of database schema changes.

**Key Fields**:
- `MigrationScriptLogId` (UUID): Unique identifier
- `Name` (String, max 100): Migration script name
- `CommitId` (String, max 100): Git commit ID when migration ran
- `IsSuccessful` (Boolean): Whether migration succeeded
- `StartDate` (Long): Migration start timestamp (milliseconds)
- `EndDate` (Long): Migration end timestamp (milliseconds)
- `Remark` (String, max 1024): Notes or error messages
- `createdAt`, `updatedAt` (DateTime): Record timestamps

**Relationships to Other Entities**:
- Standalone entity for migration tracking

**Business Rules**:
- Combination of Name and IsSuccessful must be unique
- Prevents duplicate migration execution
- Failed migrations can be retried

**Notes**:
- Critical for database version control
- Helps troubleshoot migration issues
- Provides audit trail for schema changes

---
### PRODUCT ATTRIBUTE ENTITY

**Entity Name**: ProductAttribute

**Business Description**: Represents custom attributes that can be attached to banking products to extend their base properties with additional metadata. This follows the flexible attribute pattern used throughout OBP-API, allowing products to have custom properties defined dynamically. Attributes can represent product-specific features, regulatory requirements, pricing details, or any custom metadata needed for product management and marketing.

**Key Fields**:
- `mBankId` (UUIDString): The bank this attribute belongs to
- `mProductCode` (String, max 50): The product code this attribute belongs to
- `mProductAttributeId` (MappedUUID): Unique identifier for the attribute
- `mName` (String, max 50): Name/key of the attribute
- `mType` (String, max 50): Attribute type as defined by ProductAttributeType enum (e.g., STRING, INTEGER, DOUBLE, DATE_WITH_DAY)
- `mValue` (String, max 255): The attribute value stored as string
- `mIsActive` (Boolean): Whether this attribute is currently active (default: true)

**Relationships to Other Entities**:
- A ProductAttribute belongs to one Product (via mProductCode)
- A ProductAttribute belongs to one Bank (via mBankId)
- Multiple attributes can be attached to the same product
- Attributes can be filtered by View permissions via AttributeDefinition

**Business Rules**:
- Combination of bankId and productCode is indexed for efficient queries
- ProductAttributeId is indexed for lookups
- Type determines how the value string should be interpreted
- Only active attributes are typically returned in API responses
- mIsActive defaults to true

**Notes**:
- Enables product-specific customization without code changes
- Integrates with AttributeDefinition for view-based access control
- Supports querying products by attribute name/value pairs
- Part of the broader attribute extensibility pattern in OBP-API
- Can be created, updated, or deleted through the API
- Useful for storing product features, pricing tiers, eligibility criteria, and marketing tags

---



### PRODUCT COLLECTION ITEM ENTITY

**Entity Name**: ProductCollectionItem

**Business Description**: Represents the membership relationship between products and product collections. This entity enables banks to group multiple products into collections (like "Premium Services" or "Student Banking Package") for marketing and organizational purposes.

**Key Fields**:
- `mCollectionCode` (String, max 50): Collection identifier
- `mMemberProductCode` (String, max 50): Product code that belongs to this collection
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links ProductCollection to Product entities
- Many-to-many relationship enabler between collections and products

**Business Rules**:
- Combination of mCollectionCode and mMemberProductCode must be unique
- Products can belong to multiple collections
- Collections can contain multiple products

**Notes**:
- Junction table for product organization
- Supports product bundling and marketing strategies
- Can be fully replaced when updating collection membership

---

### CONSENT AUTH CONTEXT ENTITY

**Entity Name**: ConsentAuthContext

**Business Description**: Stores authentication context information associated with consent grants. This entity captures additional authentication details (like IP address, device ID, authentication method) that were present when a consent was authorized, providing audit trail and security context.

**Key Fields**:
- `ConsentAuthContextId` (UUID): Unique identifier
- `ConsentId` (UUID): Reference to the consent
- `Key` (String, max 255): Context key (e.g., "IP_ADDRESS", "USER_AGENT")
- `Value` (String, max 255): Context value
- `createdAt` (DateTime): When context was captured

**Relationships to Other Entities**:
- Links to Consent entity
- Multiple context entries per consent (key-value pairs)

**Business Rules**:
- Combination of ConsentId, Key, and createdAt must be unique
- Immutable once created (audit trail)
- Captures point-in-time authentication state

**Notes**:
- Critical for PSD2 and regulatory compliance
- Provides forensic evidence for consent authorization
- Supports fraud detection and security analysis

---

### ENDPOINT MAPPING ENTITY

**Entity Name**: EndpointMapping

**Business Description**: Enables runtime transformation of API endpoint requests and responses without code changes. This powerful feature allows banks to customize API behavior by defining JSONata transformation rules that map between the OBP-API format and their specific requirements.

**Key Fields**:
- `EndpointMappingId` (UUID): Unique identifier
- `OperationId` (String, max 255): API operation identifier (unique per bank or globally)
- `RequestMapping` (Text): JSONata expression for request transformation
- `ResponseMapping` (Text): JSONata expression for response transformation
- `BankId` (String, max 255): Optional bank identifier for bank-specific mappings

**Relationships to Other Entities**:
- References API operations/endpoints
- Can be bank-specific or global

**Business Rules**:
- EndpointMappingId must be unique
- OperationId must be unique (within scope)
- Supports both bank-specific and global mappings
- JSONata expressions must be valid

**Notes**:
- Extremely powerful for API customization
- Enables data format translation without code deployment
- Supports multi-tenant API variations

---

### TRANSACTION REQUEST TYPE CHARGE ENTITY

**Entity Name**: TransactionRequestTypeCharge

**Business Description**: Defines the fees/charges associated with specific transaction request types. This entity enables banks to configure different fee structures for various payment types (SEPA, wire transfer, instant payment, etc.).

**Key Fields**:
- `mTransactionRequestTypeId` (UUID): Transaction request type identifier
- `mBankId` (UUID): Bank that sets this charge
- `mChargeCurrency` (String, max 3): Currency of the charge
- `mChargeAmount` (String, max 32): Charge amount
- `mChargeSummary` (String, max 255): Description of the charge
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to TransactionRequestType (implicitly)
- Links to Bank entity
- Referenced by TransactionRequest for fee calculation

**Business Rules**:
- Banks can set different charges for different transaction types
- Charge amounts stored as strings to preserve precision
- Multiple charges can exist for same type at different banks

**Notes**:
- Critical for transparent fee disclosure (PSD2 requirement)
- Supports dynamic fee calculation
- May include mock data for testing when no real charges configured

---

### ENTITLEMENT REQUEST ENTITY

**Entity Name**: EntitlementRequest

**Business Description**: Tracks user requests for specific entitlements (permissions/roles). This workflow entity enables a request-approval process where users can request access to specific API operations or roles, and administrators can review and approve/deny these requests.

**Key Fields**:
- `mEntitlementRequestId` (UUID): Unique identifier
- `mBankId` (UUID): Bank context for the request
- `mUserId` (UUID): User requesting the entitlement
- `mRoleName` (String, max 64): Role/entitlement being requested
- `createdAt` (DateTime): When request was made

**Relationships to Other Entities**:
- Links to ResourceUser entity
- Links to Bank entity
- Results in Entitlement when approved

**Business Rules**:
- EntitlementRequestId must be unique
- Combination of bankId, userId, and roleName identifies a unique request
- Request can be deleted once processed
- No direct status field - existence means "pending"

**Notes**:
- Supports access request workflow
- Administrators review and approve/deny via API
- Complements direct Entitlement assignment
- Deletion implies request was processed (approved or denied)

---

### METHOD ROUTING ENTITY

**Entity Name**: MethodRouting

**Business Description**: Configures dynamic routing of connector method calls to specific connector implementations. This powerful entity enables runtime selection of which connector handles specific operations, supporting multi-connector deployments where different banks or operations use different backend systems.

**Key Fields**:
- `MethodRoutingId` (UUID): Unique identifier
- `MethodName` (String, max 255): Name of the connector method to route
- `BankIdPattern` (String, max 255): Regex pattern to match bank IDs (default: ".*" for match-any)
- `IsBankIdExactMatch` (Boolean): Whether bank ID must match exactly
- `ConnectorName` (String, max 255): Target connector name to route to
- `Parameters` (Text): JSON array of key-value parameters for the routing

**Relationships to Other Entities**:
- References ConnectorMethod implementations
- Links to specific connector configurations
- May reference DynamicMessageDoc for message handling

**Business Rules**:
- MethodRoutingId must be unique
- BankIdPattern supports regex for flexible matching
- Default pattern ".*" matches any bank
- Parameters stored as JSON for flexible configuration
- If no bankIdPattern supplied, isExactMatch must be false

**Notes**:
- Critical for multi-connector deployments
- Enables per-bank or per-operation connector selection
- Supports A/B testing and gradual migration to new connectors
- Regex patterns enable sophisticated routing rules

---

### WEB UI PROPS ENTITY

**Entity Name**: WebUiProps

**Business Description**: Stores configurable web UI properties that customize the API explorer, landing page, and other web interfaces. This entity enables runtime configuration of branding, text, and behavior without code deployment, supporting multi-brand deployments and internationalization.

**Key Fields**:
- `WebUiPropsId` (UUID): Unique identifier
- `Name` (String, max 255): Property name (must start with "webui_", unique)
- `Value` (Text): Property value

**Relationships to Other Entities**:
- Standalone configuration entity
- Properties can be brand-specific or language-specific via naming conventions

**Business Rules**:
- WebUiPropsId must be unique
- Name must be unique
- Property names starting with "webui_" can be stored in database
- Supports brand-specific properties: `{name}_FOR_BRAND_{brand}`
- Supports language-specific properties: `{name}_{language}`
- Priority: branded+translated > branded > translated > default
- Cached with configurable TTL (default from props file)

**Notes**:
- Enables white-label deployments
- Supports internationalization (i18n)
- Falls back to props file if property not in database
- Critical for UI customization without code changes

---

### BANK ACCOUNT DATA ENTITY

**Entity Name**: BankAccountData

**Business Description**: Internal utility entity for storing additional bank account metadata, particularly account labels. This appears to be a supplementary data store for account information that may not be available in the core banking system or needs to be cached locally.

**Key Fields**:
- `bankId` (String, max 255): Bank identifier
- `accountId` (String, max 255): Account identifier
- `accountLabel` (String, max 255): Human-readable account label
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- References BankAccount entities (via bankId and accountId)
- Supplementary to main BankAccount entity

**Business Rules**:
- Combination of bankId and accountId must be unique
- Appears to be a cache or supplementary data store

**Notes**:
- Simple supplementary entity for account metadata
- May be used when core banking doesn't provide labels
- Could be used for caching or local overrides
- Distinct from main MappedBankAccount entity

---

### TRANSACTION REQUEST REASONS ENTITY

**Entity Name**: TransactionRequestReasons

**Business Description**: Stores detailed reason codes and supporting documentation for payment transactions. This entity is critical for regulatory compliance, particularly for cross-border payments and large transactions that require detailed justification under anti-money laundering (AML) and know-your-transaction regulations.

**Key Fields**:
- `TransactionRequestReasonId` (UUID): Unique identifier
- `TransactionRequestId` (UUID): Link to the transaction request
- `Code` (String, max 8): Standardized reason code
- `DocumentNumber` (String, max 100): Supporting document reference
- `Currency` (String, max 3): Currency code
- `Amount` (String, max 32): Amount related to this reason
- `Description` (String, max 2048): Detailed explanation
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to TransactionRequest entity
- Multiple reasons can be associated with one transaction request

**Business Rules**:
- Auto-generates UUID on creation
- Supports detailed regulatory reporting
- Amounts stored as strings to preserve precision

**Notes**:
- Critical for PSD2 and AML compliance
- Enables structured reason documentation
- Supports regulatory reporting requirements
- Can include invoice numbers, contract references, etc.

---

### BAD LOGIN ATTEMPT ENTITY

**Entity Name**: BadLoginAttempt

**Business Description**: Tracks failed login attempts for security monitoring and account protection. This entity helps prevent brute-force attacks by recording failed authentication attempts and enabling account lockout policies.

**Key Fields**:
- `mUsername` (String, max 100, required): Username that attempted login
- `Provider` (String, max 100): Authentication provider
- `mBadAttemptsSinceLastSuccessOrReset` (Integer): Count of consecutive failed attempts
- `mLastFailureDate` (DateTime): Timestamp of most recent failure

**Relationships to Other Entities**:
- Links to ResourceUser authentication
- Tracks attempts per provider-username combination

**Business Rules**:
- Combination of Provider and Username must be unique
- Counter resets on successful login
- Username field is mandatory
- Tracks consecutive failures only

**Notes**:
- Critical security entity
- Enables account lockout policies
- Helps detect brute-force attacks
- Can trigger alerts for suspicious activity
- Counter increments with each failure

---

### COUNTERPARTY LIMIT ENTITY

**Entity Name**: CounterpartyLimit

**Business Description**: Defines transaction limits for specific counterparties from a particular account view. This entity enables fine-grained control over payment amounts and frequencies to individual counterparties, supporting fraud prevention, spend control, and regulatory compliance.

**Key Fields**:
- `CounterpartyLimitId` (UUID): Unique identifier
- `BankId` (String, max 255, required): Bank identifier
- `AccountId` (String, max 255, required): Account identifier
- `ViewId` (String, max 255, required): View identifier
- `CounterpartyId` (String, max 255, required): Counterparty identifier
- `Currency` (String, max 255): Currency code
- `MaxSingleAmount` (Decimal): Maximum per-transaction amount (default: 0)
- `MaxMonthlyAmount` (Decimal): Maximum monthly total (default: 0)
- `MaxNumberOfMonthlyTransactions` (Integer): Monthly transaction count limit (default: -1)
- `MaxYearlyAmount` (Decimal): Maximum yearly total (default: 0)
- `MaxNumberOfYearlyTransactions` (Integer): Yearly transaction count limit (default: -1)
- `MaxTotalAmount` (Decimal): Maximum lifetime total (default: 0)
- `MaxNumberOfTransactions` (Integer): Lifetime transaction count limit (default: -1)
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to BankAccount entity
- Links to Counterparty entity
- Links to View for access control
- Part of the view-based permission system

**Business Rules**:
- Combination of BankId, AccountId, ViewId, and CounterpartyId must be unique
- CounterpartyLimitId must be unique
- Supports create-or-update pattern
- Value of -1 for count limits means "no limit"
- Value of 0 for amount limits means "blocked"

**Notes**:
- Powerful fraud prevention tool
- Enables corporate spend control policies
- Supports regulatory velocity checks
- Can prevent unauthorized large payments
- Limits are view-specific (different limits for different access levels)

---

### OPENID CONNECT TOKEN ENTITY

**Entity Name**: OpenIDConnectToken

**Business Description**: Stores OpenID Connect (OIDC) authentication tokens including access tokens, ID tokens, and refresh tokens. This entity manages OAuth 2.0 / OIDC token lifecycle for users who authenticate through external identity providers that support the OpenID Connect protocol.

**Key Fields**:
- `AccessToken` (MappedText): OAuth 2.0 access token for API authentication
- `IDToken` (MappedText): OpenID Connect ID token containing user identity claims
- `RefreshToken` (MappedText): Refresh token for obtaining new access tokens
- `Scope` (String, max 250): OAuth scopes granted to this token
- `TokenType` (String, max 250): Token type (typically "Bearer")
- `ExpiresIn` (Long): Token expiration time in seconds
- `AuthUserPrimaryKey` (Long): Foreign key to the authenticated user
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to ResourceUser via AuthUserPrimaryKey
- One user can have multiple tokens over time (new token replaces old)
- Most recent token per user is retrieved for authentication

**Business Rules**:
- Tokens are stored securely and associated with authenticated users
- Most recent token by creation date is used for each user
- Supports full OpenID Connect token response structure
- Access tokens enable API authentication
- Refresh tokens enable token renewal without re-authentication
- ID tokens contain user identity information

**Notes**:
- Enables Single Sign-On (SSO) through OpenID Connect providers
- Supports OAuth 2.0 authorization code flow
- Tokens are typically obtained from external identity providers (e.g., Google, Azure AD, Keycloak)
- System retrieves most recent token when user authenticates
- Critical for integrating with enterprise identity management systems
- Supports modern authentication standards (OIDC, OAuth 2.0)

---

### ETAG ENTITY

**Entity Name**: ETag

**Business Description**: Stores HTTP ETag values for API response caching and conditional requests. This technical entity enables efficient caching by tracking resource versions through entity tags, reducing bandwidth and improving API performance through conditional GET requests.

**Key Fields**:
- `ETagResource` (String, max 1000, unique): Resource identifier/path
- `ETagValue` (String, max 256): Current ETag value for the resource
- `LastUpdatedMSSinceEpoch` (Long): Timestamp in milliseconds since epoch

**Relationships to Other Entities**:
- Can reference any API resource
- Standalone caching mechanism

**Business Rules**:
- ETagResource must be unique
- Used for HTTP ETag headers
- Updates when resource changes
- Enables HTTP 304 Not Modified responses

**Notes**:
- Technical infrastructure entity
- Critical for API performance optimization
- Supports HTTP caching standards
- Reduces unnecessary data transfer
- Enables efficient polling patterns

---

### PEM USAGE ENTITY

**Entity Name**: PemUsage

**Business Description**: Tracks the usage of PEM (Privacy Enhanced Mail) certificates for API authentication. This entity maintains a record of which consumer applications have used specific certificates and which users last authenticated with them, supporting certificate-based authentication workflows.

**Key Fields**:
- `PemHash` (String, max 50, unique): Hash of the PEM certificate
- `ConsumerId` (String, max 50): Consumer application using this certificate
- `LastUserId` (String, max 50): Last user who authenticated with this certificate
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to Consumer entity via ConsumerId
- Links to ResourceUser via LastUserId
- Tracks certificate usage patterns

**Business Rules**:
- PemHash must be unique
- Certificate can be used by multiple consumers over time
- LastUserId tracks most recent authentication

**Notes**:
- Technical security entity
- Supports PEM certificate authentication method
- Enables certificate rotation tracking
- Helps identify certificate misuse

---

### AUTHENTICATION TYPE VALIDATION ENTITY

**Entity Name**: AuthenticationTypeValidation

**Business Description**: Defines which authentication methods are allowed for specific API operations. This entity implements operation-level authentication policy, enabling fine-grained control over which authentication types (Direct Login, OAuth, Gateway Login, etc.) can be used to access particular API endpoints.

**Key Fields**:
- `OperationId` (String, max 200, unique): Unique identifier for the API operation/endpoint
- `AllowedAuthTypes` (String, max 300): Comma-separated list of allowed authentication types

**Relationships to Other Entities**:
- References API operations/endpoints by OperationId
- Enforces authentication policy at the API gateway level
- Works in conjunction with Consumer and ResourceUser authentication

**Business Rules**:
- OperationId must be unique (one validation rule per operation)
- AllowedAuthTypes contains comma-separated authentication method names
- Common auth types: "DirectLogin", "OAuth", "GatewayLogin", "Anonymous"
- If no validation exists for an operation, default authentication rules apply
- Empty or null allowed types may block all access to that operation

**Notes**:
- Security configuration entity
- Enables operation-specific authentication requirements
- Useful for enforcing stricter authentication on sensitive operations
- Can require OAuth for data modification while allowing DirectLogin for read operations
- Supports compliance requirements for different authentication strengths
- Part of the API security policy framework

---

### NONCE ENTITY

**Entity Name**: Nonce

**Business Description**: Stores OAuth 1.0 nonce values used for preventing replay attacks. A nonce (number used once) is a unique value included in each OAuth request that the server tracks to ensure the same request cannot be replayed by an attacker who intercepts the request. This is a critical security mechanism in OAuth 1.0 authentication.

**Key Fields**:
- `id` (Long, indexed): Primary key identifier
- `consumerkey` (String, max 250): OAuth consumer key making the request
- `tokenKey` (String, max 250): OAuth token key (defaults to empty string for unsigned requests)
- `timestamp` (DateTime): Timestamp when the request was made
- `value` (String, max 250): The unique nonce value for this request

**Relationships to Other Entities**:
- Links to Consumer via consumerkey (stored as string, not foreign key)
- Links to Token via tokenKey (stored as string, not foreign key)
- Part of OAuth 1.0 authentication flow
- Works in conjunction with Consumer and Token for request validation

**Business Rules**:
- Each combination of consumerKey, tokenKey, timestamp, and value should be unique
- Nonces are counted to detect replay attempts (count > 0 means replay)
- Expired nonces are periodically deleted to prevent database bloat
- Token key can be empty string for unsigned requests
- Timestamp is stored as DateTime but converted to milliseconds for validation

**Notes**:
- OAuth 1.0 security mechanism (OAuth 2.0 uses different replay prevention)
- Prevents replay attacks where intercepted requests are re-submitted
- System checks if nonce was already used before accepting request
- Old nonces are deleted after expiration to maintain database performance
- Critical for OAuth 1.0 security but not used in OAuth 2.0 or OpenID Connect flows
- The nonce value combined with timestamp must be unique per consumer/token combination
- Supports both signed requests (with tokenKey) and unsigned requests (empty tokenKey)

---

### USER AUTH CONTEXT UPDATE ENTITY

**Entity Name**: UserAuthContextUpdate

**Business Description**: Manages updates to user authentication context with challenge-response verification. This entity enables secure updates to authentication context by requiring users to solve a challenge before the update is applied, preventing unauthorized modification of authentication settings.

**Key Fields**:
- `mUserAuthContextUpdateId` (UUID): Unique identifier
- `mUserId` (UUID): User requesting the update
- `mConsumerId` (String, max 255): Consumer application making the request
- `mKey` (String, max 50): Context key to update
- `mValue` (String, max 50): New value for the context
- `mChallenge` (String, max 10): Random challenge code (auto-generated)
- `mStatus` (String, max 20): Status of the update request
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- Links to ResourceUser entity
- Links to Consumer entity
- Updates UserAuthContext upon successful challenge completion

**Business Rules**:
- Challenge auto-generated as random number (up to 99999999)
- Status tracks lifecycle: pending, verified, failed, etc.
- User must provide correct challenge to complete update
- One-time use per update request

**Notes**:
- Critical for secure authentication context management
- Prevents unauthorized context tampering
- Challenge sent via secure channel (SMS, email, etc.)
- Supports SCA (Strong Customer Authentication) requirements

---

### ENDPOINT TAG ENTITY

**Entity Name**: EndpointTag

**Business Description**: Assigns tags to API endpoints for organization, categorization, and filtering. This entity enables flexible grouping of endpoints by functionality, business domain, or other criteria, supporting API discovery and documentation organization.

**Key Fields**:
- `EndpointTagId` (UUID): Unique identifier
- `OperationId` (String, max 255): API operation identifier
- `TagName` (String, max 255): Tag name/label
- `BankId` (String, max 255, optional): Bank-specific tag
- `createdAt`, `updatedAt` (DateTime): Timestamps

**Relationships to Other Entities**:
- References API endpoints via OperationId
- Can be bank-specific or global
- Multiple tags can be applied to one endpoint

**Business Rules**:
- EndpointTagId must be unique
- Tags can be bank-specific or apply globally
- Supports hierarchical tag organization
- Create-or-update pattern supported

**Notes**:
- Technical organization entity
- Enhances API explorer usability
- Supports endpoint filtering and search
- Enables business domain grouping
- Can support API versioning strategies

---

### FINANCIAL OPERATIONS

---

### FX RATE ENTITY

**Entity Name**: FXRate

**Business Description**: Foreign exchange rates for currency conversion between different currencies.

**Key Fields**:
- `mBankId` (UUIDString): Bank offering this rate
- `mFromCurrencyCode` (String, max 3): Source currency
- `mToCurrencyCode` (String, max 3): Target currency
- `mConversionValue` (Double): Conversion rate
- `mInverseConversionValue` (Double): Inverse rate
- `mEffectiveDate` (DateTime): When rate is effective

**Relationships to Other Entities**:
- Links to Currency entities (foreign keys)
- Associated with Bank

**Business Rules**:
- Maintains forward and inverse rates
- Time-based effective dates

---

### CURRENCY ENTITY

**Entity Name**: Currency

**Business Description**: Supported currencies in the system.

**Key Fields**:
- Currency code and related metadata

**Relationships to Other Entities**:
- Referenced by FXRate
- Used in accounts, transactions, amounts throughout system

---

### ACCOUNT APPLICATION ENTITY

**Entity Name**: AccountApplication

**Business Description**: Application for opening a new bank account. Tracks the application workflow from submission to approval/rejection.

**Key Fields**:
- Application details including applicant information
- Product being applied for
- Status (REQUESTED, ACCEPTED, REJECTED)
- Workflow timestamps

**Relationships to Other Entities**:
- Links to Product
- Links to User (applicant)
- Creates Account when approved

**Business Rules**:
- Workflow-based processing
- Status transitions tracked

---

### ENTITLEMENT ENTITY

**Entity Name**: Entitlement

**Business Description**: Permissions/roles assigned to users defining what they can do in the system. Controls access to API endpoints and operations.

**Key Fields**:
- `entitlementId` (String): Unique identifier
- `userId` (String): User receiving entitlement
- `roleName` (String): Role name (e.g., CanCreateAccount, CanGetCustomers)
- `bankId` (String, optional): Bank-specific entitlement

**Relationships to Other Entities**:
- Belongs to ResourceUser
- Can be bank-specific or system-wide

**Business Rules**:
- Role-based access control
- Can be scoped to specific banks
- Combination of userId + roleName + bankId should be unique

**Notes**:
- Core authorization mechanism
- Hundreds of granular roles available
- Supports delegation and admin management

---

## Entity Relationship Map

This section describes the major relationships between entities in the OBP-API system.

### Core Banking Hierarchy
```
Bank
├── Has Many: Branches
├── Has Many: ATMs
├── Has Many: BankAccounts
│   ├── Has Many: Transactions
│   ├── Has Many: AccountHolders
│   ├── Has Many: CustomerAccountLinks → Customers
│   ├── Has Many: AccountAttributes
│   ├── Has Many: Views → AccountAccess → Users
│   └── Has Many: PhysicalCards
├── Has Many: Customers
│   ├── Has Many: CustomerAddresses
│   ├── Has Many: CustomerAttributes
│   ├── Has One: UserCustomerLink → ResourceUser
│   └── Has Many: KycChecks, KycDocuments
└── Has Many: Products
    ├── Has Many: ProductFees
    ├── Has Many: ProductAttributes
    └── Has Many: ProductCollections

```

### Transaction Ecosystem
```
Transaction
├── Belongs To: BankAccount
├── Has Many: TransactionAttributes
├── Has Many: TransactionImages
├── Has Many: Tags
├── Has Many: Comments
├── Has One: Narrative
├── Has One: WhereTag
└── References: Counterparty

TransactionRequest
├── Links: Source Account
├── Links: Destination Account
├── Has Many: TransactionChallenges
├── Has Many: TransactionRequestAttributes
└── Creates: Transaction (when completed)
```

### User & Security
```
ResourceUser
├── Has Many: AccountAccess (via Views)
├── Has Many: Entitlements
├── Has One: UserCustomerLink → Customer
├── Has Many: Consents (granted)
├── Creates: Consumers (applications)
└── Performs: CrmEvents, Meetings

Consumer (Application)
├── Belongs To: ResourceUser (creator)
├── Has: RateLimiting configuration
├── Generates: Metrics (API calls)
└── Can Have: Consents (for data access)

View
├── Belongs To: BankAccount
├── Has Many: AccountAccess → Users
└── Has: ViewPermissions (granular permissions)
```

### Payment Operations
```
BankAccount
├── Can Have: DirectDebits (recurring in)
├── Can Have: StandingOrders (recurring out)
└── Can Have: CounterpartyLimits

Counterparty
├── Has One: CounterpartyMetadata
├── Can Have: CounterpartyLimits
└── Referenced By: Transactions
```

### KYC & Compliance
```
Customer
├── Has Many: KycChecks
├── Has Many: KycDocuments
├── Has One: KycStatus
├── Has Many: TaxResidences
└── Linked Via: Consents (for third-party access)

RegulatedEntity
├── Has: Certificate (for authentication)
└── Has Many: RegulatedEntityAttributes
```

### Extensibility Pattern
```
AttributeDefinition (Schema)
├── Defines: AccountAttributes
├── Defines: ProductAttributes
├── Defines: CustomerAttributes
├── Defines: TransactionAttributes
├── Defines: CardAttributes
├── Defines: BankAttributes
├── Defines: AtmAttributes
└── Controls: Visibility via Views
```

---

## Business Rules Summary

### Identity & Uniqueness
- Most entities use UUIDs for primary identification
- Unique constraints typically combine organizational scope (bankId) with entity identifier
- Examples:
  - Bank: permalink (bankId) is unique
  - BankAccount: (bankId, accountId) is unique
  - Customer: (bankId, customerNumber) is unique

### Soft Deletion
- Many entities support soft deletion (IsDeleted flag) rather than hard deletion
- Allows audit trail and historical analysis
- Example: ResourceUser has IsDeleted field

### Temporal Validity
- Time-based validity common in:
  - RateLimiting: fromDate/toDate
  - Consent: validity periods
  - FXRate: effectiveDate
  - Cards: validFrom/expires

### Monetary Values
- Amounts stored in smallest currency units (cents, pence, etc.) as Long integers
- Converted to BigDecimal for display
- Prevents floating-point precision errors
- Examples: Transaction.amount, BankAccount.accountBalance

### Audit Trail
- Most entities inherit CreatedUpdated trait providing:
  - createdAt timestamp
  - updatedAt timestamp
- Metrics table provides comprehensive API usage audit trail

### Access Control
- View-based permission system
- Fine-grained control over what users can see/do
- Public vs Private views
- System vs Custom views

### Extensibility
- Dynamic attributes allow customization without schema changes
- AttributeDefinition defines schema
- Actual values stored in typed attribute tables
- Supports bank-specific and regulatory requirements

---

## Questions & Uncertainties

### Clarifications Needed

1. **Counterparty Dual Nature**: The codebase comments indicate two types of counterparties (explicit and implicit). What are the exact business rules for when each type is used?

2. **Card Entity Typo**: The PhysicalCard implementation is in files named "PhisicalCard" (typo). Should this be corrected, or is it maintained for backwards compatibility?

3. **Account Holder vs Customer**: There appear to be overlapping concepts. What is the precise business distinction between AccountHolder and Customer entities?

4. **View Permission Granularity**: The exact permission fields in Views/ViewPermissions need validation with business stakeholders to ensure they match current Open Banking requirements.

5. **Metric Archival**: What is the archival policy for metrics? How long are they kept in the active Metric table before moving to MetricArchive?

6. **Attribute Type Enforcement**: How strictly are attribute type definitions enforced when creating attribute values? Is there runtime validation?

7. **Consumer vs User**: What are the exact scenarios where a Consumer is needed vs just a User? When does an application need both?

8. **Currency Support**: Is there a complete list of supported currencies in the Currency table, or are they added dynamically?

9. **Transaction Request Lifecycle**: What are all possible status transitions for TransactionRequest entities? Is there a state machine diagram?

10. **Consent Revocation**: When a consent is revoked, what cascading effects occur (e.g., are created users disabled, are views revoked)?

### Data Migration Considerations

1. **Historical Data**: Some entities (like Metric) likely contain large amounts of historical data. Migration strategies need consideration.

2. **UUID Migration**: Comments in code mention migration from integer PKs to UUIDs. Are there still legacy integer references that need handling?

3. **Unique Constraints**: Several entities have comments about unique constraints not being enforced. Should these be added during migration?

### Integration Points

1. **External Systems**: How do DynamicEntities integrate with core banking systems?

2. **Webhook Reliability**: What retry/failure handling exists for webhooks?

3. **FX Rate Sources**: Where do FX rates come from? Are they manually entered or automatically synced from external sources?

4. **Certificate Management**: For RegulatedEntity certificates, is there automated validation/renewal?

---

## Implementation Notes

### Technology Stack
- **Language**: Scala
- **Web Framework**: Lift
- **ORM**: Lift Mapper
- **Database**: PostgreSQL, H2, MySQL (via JDBC)
- **Caching**: Redis (for some operations)
- **Search**: Elasticsearch (optional)

### Code Organization Pattern
All entities follow a consistent pattern:
1. **Provider Trait**: Defines the interface (e.g., `CustomerProvider`)
2. **Mapped Implementation**: Implements the interface (e.g., `MappedCustomerProvider`)
3. **Entity Class**: The actual ORM entity (e.g., `MappedCustomer`)
4. **Companion Object**: Provides database metadata (e.g., `object MappedCustomer`)

### Common Mapped Types
- `UUIDString`: String field storing UUIDs
- `MappedUUID`: UUID field with automatic UUID generation
- `AccountIdString`: Account ID field
- `MappedAccountNumber`: Account number field
- `MappedLongForeignKey`: Foreign key to another entity
- `CreatedUpdated`: Trait adding created/updated timestamps

### Database Indexing
Most entities define custom indexes via `dbIndexes` override:
- Unique indexes for business keys
- Foreign key indexes for relationships
- Query optimization indexes

---

## Conclusion

The OBP-API represents a sophisticated, well-architected banking platform with comprehensive entity coverage across all major banking domains. The system demonstrates strong design principles including:

- **Separation of Concerns**: Clear domain boundaries
- **Extensibility**: Dynamic attributes for customization
- **Security**: Multi-layered access control via Views and Entitlements
- **Compliance**: Rich KYC, consent, and regulatory entities
- **Audit**: Comprehensive tracking via metrics and timestamps
- **Flexibility**: Support for multiple banks, products, and regulatory frameworks

The entity model supports modern Open Banking standards (PSD2, UK Open Banking) while maintaining backwards compatibility and allowing for bank-specific customization. The consistent architectural patterns make the codebase maintainable despite its complexity.

---

*Document Generated: October 12, 2025*  
*Total Entities Documented: 100*  
*Primary Code Path: obp-api/src/main/scala/code/*
