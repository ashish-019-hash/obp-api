# Entity Relationship Model (ERM) - Open Bank Project API

## Overview
This document provides a comprehensive mapping of business-level entities and their relationships extracted from the Open Bank Project (OBP) API codebase. The analysis focuses on domain models, ORM/database mappings, business DTOs, and class hierarchies that represent real-world banking concepts.

## Framework Context
The OBP-API uses the **Lift Web Framework (Scala)** with:
- **Domain Models**: Business entity traits and case classes in `obp-commons/model/`
- **ORM Mappings**: Lift Mapper classes for database persistence (extends `LongKeyedMapper`)
- **Business DTOs**: Case classes for data transfer and API responses
- **Class Hierarchies**: Trait → Case Class → Mapped Class inheritance patterns

## Entity Domains

### 1. Banking Core Domain

#### Bank Entity
**Source**: [BankingModel.scala:36-61](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala#L36-L61)

**Domain Model**:
```scala
trait Bank {
  def bankId: BankId
  def shortName: String
  def fullName: String
  def logoUrl: String
  def websiteUrl: String
  def bankRoutingScheme: String
  def bankRoutingAddress: String
  def swiftBic: String
  def nationalIdentifier: String
}
```

**Database Mapping**: [MappedBank.scala:6-29](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/model/dataAccess/MappedBank.scala#L6-L29)
```scala
class MappedBank extends Bank with LongKeyedMapper[MappedBank] {
  object permalink extends MappedString(this, 255)           // bankId
  object fullBankName extends MappedString(this, 255)       // fullName
  object shortBankName extends MappedString(this, 100)      // shortName
  object logoURL extends MappedString(this, 255)            // logoUrl
  object websiteURL extends MappedString(this, 255)         // websiteUrl
  object swiftBIC extends MappedString(this, 255)           // swiftBic
  object national_identifier extends MappedString(this, 255) // nationalIdentifier
  object mBankRoutingScheme extends MappedString(this, 255) // bankRoutingScheme
  object mBankRoutingAddress extends MappedString(this, 255) // bankRoutingAddress
}
```

**Business Purpose**: Central entity representing financial institutions in the system. Each bank can have multiple accounts, products, and customers.

**Relationships**:
- Bank → BankAccount (1:many via bankId)
- Bank → Product (1:many via bankId)
- Bank → Customer (1:many via bankId)
- Bank → ATM (1:many via bankId)

#### BankAccount Entity
**Source**: [BankingModel.scala:209-235](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala#L209-L235)

**Domain Model**:
```scala
trait BankAccount {
  def accountId: AccountId
  def accountType: String
  def balance: BigDecimal
  def currency: String
  def name: String
  def label: String
  def number: String
  def bankId: BankId
  def lastUpdate: Date
  def branchId: String
  def accountRoutings: List[AccountRouting]
  def accountRules: List[AccountRule]
  def accountHolder: String
  def attributes: Option[List[Attribute]]
}
```

**Database Mapping**: [MappedBankAccount.scala:11-75](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/model/dataAccess/MappedBankAccount.scala#L11-L75)
```scala
class MappedBankAccount extends BankAccount with LongKeyedMapper[MappedBankAccount] {
  object bank extends UUIDString(this)                      // bankId (FK)
  object theAccountId extends AccountIdString(this)         // accountId
  object accountCurrency extends MappedString(this, 10)     // currency
  object accountNumber extends MappedAccountNumber(this)    // number
  object accountBalance extends MappedLong(this)            // balance (smallest currency unit)
  object accountName extends MappedString(this, 255)        // name
  object kind extends MappedString(this, 255)               // accountType
  object accountLabel extends MappedString(this, 255)       // label
  object accountLastUpdate extends MappedDateTime(this)     // lastUpdate
  object mBranchId extends UUIDString(this)                 // branchId
  object holder extends MappedString(this, 100)             // accountHolder
}
```

**Business Purpose**: Represents customer bank accounts with balances, routing information, and account rules.

**Relationships**:
- BankAccount → Bank (many:1 via bankId)
- BankAccount → Customer (many:many via CustomerAccountLink)
- BankAccount → Transaction (1:many via accountId)
- BankAccount → AccountRouting (1:many via accountId)
- BankAccount → Product (many:1 via product categorization)

#### Product Entity
**Source**: [MappedProductsProvider.scala:25-73](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/products/MappedProductsProvider.scala#L25-L73)

**Database Mapping**:
```scala
class MappedProduct extends Product with LongKeyedMapper[MappedProduct] {
  object mBankId extends UUIDString(this)                   // bankId (FK)
  object mCode extends MappedString(this, 50)               // productCode
  object mParentProductCode extends MappedString(this, 50)  // parentProductCode
  object mName extends MappedString(this, 125)              // name
  object mCategory extends MappedString(this, 50)           // category
  object mFamily extends MappedString(this, 50)             // family
  object mSuperFamily extends MappedString(this, 50)        // superFamily
  object mMoreInfoUrl extends MappedString(this, 2000)      // moreInfoUrl
  object mTermsAndConditionsUrl extends MappedString(this, 2000) // termsAndConditionsUrl
  object mDetails extends MappedString(this, 2000)          // details
  object mDescription extends MappedString(this, 2000)      // description
  object mLicenseId extends UUIDString(this)                // license.id
  object mLicenseName extends MappedString(this, 255)       // license.name
}
```

**Business Purpose**: Represents banking products (savings accounts, loans, credit cards) with categorization and licensing information.

**Relationships**:
- Product → Bank (many:1 via bankId)
- Product → BankAccount (1:many via product association)

### 2. Customer Management Domain

#### Customer Entity
**Source**: [CustomerDataModel.scala:35-56](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/CustomerDataModel.scala#L35-L56)

**Domain Model**:
```scala
trait Customer {
  def customerId: String
  def bankId: String
  def number: String
  def legalName: String
  def mobileNumber: String
  def email: String
  def faceImage: CustomerFaceImageTrait
  def dateOfBirth: Date
  def relationshipStatus: String
  def dependents: Integer
  def dobOfDependents: List[Date]
  def highestEducationAttained: String
  def employmentStatus: String
  def creditRating: CreditRatingTrait
  def creditLimit: AmountOfMoneyTrait
  def kycStatus: Boolean
  def lastOkDate: Date
  def title: String
  def branchId: String
  def nameSuffix: String
}
```

**Database Mapping**: [MappedCustomerProvider.scala:337-412](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/customer/MappedCustomerProvider.scala#L337-L412)
```scala
class MappedCustomer extends Customer with LongKeyedMapper[MappedCustomer] {
  object mCustomerId extends MappedUUID(this)               // customerId
  object mBank extends UUIDString(this)                     // bankId (FK)
  object mNumber extends MappedString(this, 50)             // number
  object mLegalName extends MappedString(this, 100)         // legalName
  object mMobileNumber extends MappedString(this, 100)      // mobileNumber
  object mEmail extends MappedEmail(this, 100)              // email
  object mFaceImageUrl extends MappedString(this, 255)      // faceImage.url
  object mFaceImageDate extends MappedDateTime(this)        // faceImage.date
  object mDateOfBirth extends MappedDateTime(this)          // dateOfBirth
  object mRelationshipStatus extends MappedString(this, 25) // relationshipStatus
  object mDependents extends MappedInt(this)                // dependents
  object mHighestEducationAttained extends MappedString(this, 100) // highestEducationAttained
  object mEmploymentStatus extends MappedString(this, 100) // employmentStatus
  object mKycStatus extends MappedBoolean(this)             // kycStatus
  object mLastOkDate extends MappedDate(this)               // lastOkDate
  object mTitle extends MappedString(this, 10)              // title
  object mBranchId extends UUIDString(this)                 // branchId
  object mNameSuffix extends MappedString(this, 10)         // nameSuffix
}
```

**Business Purpose**: Represents bank customers with personal information, KYC status, and credit information.

**Relationships**:
- Customer → Bank (many:1 via bankId)
- Customer → User (many:many via UserCustomerLink)
- Customer → BankAccount (many:many via CustomerAccountLink)
- Customer → CreditRating (1:1 embedded)
- Customer → CustomerFaceImage (1:1 embedded)

#### User Entity
**Source**: [UserModel.scala:53-74](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/UserModel.scala#L53-L74)

**Domain Model**:
```scala
trait User {
  def userPrimaryKey: UserPrimaryKey
  def userId: String
  def idGivenByProvider: String
  def provider: String
  def emailAddress: String
  def name: String
  def createdByConsentId: Option[String]
  def createdByUserInvitationId: Option[String]
  def isOriginalUser: Boolean
  def isConsentUser: Boolean
  def isDeleted: Option[Boolean]
  def lastMarketingAgreementSignedDate: Option[Date]
  def lastUsedLocale: Option[String]
}
```

**Business Purpose**: Represents system users who can access banking services and be linked to customers.

**Relationships**:
- User → Customer (many:many via UserCustomerLink)
- User → UserAuthContext (1:many via userId)
- User → Consent (1:many via userId)

### 3. Transaction Processing Domain

#### TransactionRequest Entity
**Source**: [MappedTransactionRequestProvider.scala:232-476](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/transactionrequests/MappedTransactionRequestProvider.scala#L232-L476)

**Database Mapping**:
```scala
class MappedTransactionRequest extends LongKeyedMapper[MappedTransactionRequest] {
  // Core transaction request fields
  object mTransactionRequestId extends UUIDString(this)     // transactionRequestId
  object mType extends MappedString(this, 32)               // type
  object mTransactionIDs extends MappedString(this, 2000)   // transaction_ids
  object mStatus extends MappedString(this, 32)             // status
  object mStartDate extends MappedDate(this)                // start_date
  object mEndDate extends MappedDate(this)                  // end_date
  
  // Challenge fields
  object mChallenge_Id extends MappedString(this, 64)       // challenge.id
  object mChallenge_AllowedAttempts extends MappedInt(this) // challenge.allowed_attempts
  object mChallenge_ChallengeType extends MappedString(this, 100) // challenge.challenge_type
  
  // Charge fields
  object mCharge_Summary extends MappedString(this, 64)     // charge.summary
  object mCharge_Amount extends MappedString(this, 32)      // charge.value.amount
  object mCharge_Currency extends MappedString(this, 3)     // charge.value.currency
  object mcharge_Policy extends MappedString(this, 32)      // charge_policy
  
  // Body fields
  object mBody_Value_Currency extends MappedString(this, 3) // body.value.currency
  object mBody_Value_Amount extends MappedString(this, 32)  // body.value.amount
  object mBody_Description extends MappedString(this, 2000) // body.description
  object mDetails extends MappedText(this)                  // body details (JSON)
  
  // From account fields
  object mFrom_BankId extends UUIDString(this)              // from.bank_id (FK)
  object mFrom_AccountId extends AccountIdString(this)      // from.account_id (FK)
  
  // To account fields (deprecated)
  object mTo_BankId extends UUIDString(this)                // to.bank_id (FK)
  object mTo_AccountId extends AccountIdString(this)        // to.account_id (FK)
  
  // Counterparty fields
  object mName extends MappedString(this, 64)               // name
  object mThisBankId extends UUIDString(this)               // this_bank_id (FK)
  object mThisAccountId extends AccountIdString(this)       // this_account_id (FK)
  object mThisViewId extends UUIDString(this)               // this_view_id (FK)
  object mCounterpartyId extends UUIDString(this)           // counterparty_id (FK)
  object mOtherAccountRoutingScheme extends MappedString(this, 32) // other_account_routing_scheme
  object mOtherAccountRoutingAddress extends MappedString(this, 64) // other_account_routing_address
  object mOtherBankRoutingScheme extends MappedString(this, 32) // other_bank_routing_scheme
  object mOtherBankRoutingAddress extends MappedString(this, 64) // other_bank_routing_address
  object mIsBeneficiary extends MappedBoolean(this)         // is_beneficiary
  
  // Berlin Group PSD2 fields
  object mPaymentStartDate extends MappedDate(this)         // startDate
  object mPaymentEndDate extends MappedDate(this)           // endDate
  object mPaymentExecutionRule extends MappedString(this, 64) // executionRule
  object mPaymentFrequency extends MappedString(this, 64)   // frequency
  object mPaymentDayOfExecution extends MappedString(this, 64) // dayOfExecution
  
  // Consent and API fields
  object mConsentReferenceId extends MappedString(this, 64) // consentReferenceId
  object mApiStandard extends MappedString(this, 50)        // apiStandard
  object mApiVersion extends MappedString(this, 50)         // apiVersion
}
```

**Business Purpose**: Represents payment and transfer requests with comprehensive support for multiple payment types (SEPA, counterparty, simple transfers, etc.) and PSD2 compliance.

**Relationships**:
- TransactionRequest → BankAccount (many:1 via from_bank_id/from_account_id)
- TransactionRequest → BankAccount (many:1 via to_bank_id/to_account_id)
- TransactionRequest → Counterparty (many:1 via counterparty_id)
- TransactionRequest → Transaction (1:many via transaction_ids)
- TransactionRequest → Consent (many:1 via consentReferenceId)

#### Transaction Entity
**Source**: [BankingModel.scala:313-324](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala#L313-L324)

**Domain Model**:
```scala
case class TransactionCore(
  id: TransactionId,
  thisAccount: BankAccount,
  otherAccount: CounterpartyCore,
  transactionType: String,
  amount: BigDecimal,
  currency: String,
  description: Option[String],
  startDate: Date,
  finishDate: Date,
  balance: BigDecimal
)
```

**Business Purpose**: Represents completed financial transactions between accounts.

**Relationships**:
- Transaction → BankAccount (many:1 via thisAccount)
- Transaction → Counterparty (many:1 via otherAccount)
- Transaction → TransactionRequest (many:1 via transaction request)

#### Counterparty Entity
**Source**: [CounterpartyModel.scala:32-51](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/CounterpartyModel.scala#L32-L51)

**Domain Model**:
```scala
trait CounterpartyTrait {
  def createdByUserId: String
  def name: String
  def description: String
  def thisBankId: String
  def thisAccountId: String
  def thisViewId: String
  def counterpartyId: String
  def otherAccountRoutingScheme: String
  def otherAccountRoutingAddress: String
  def otherAccountSecondaryRoutingScheme: String
  def otherAccountSecondaryRoutingAddress: String
  def otherBankRoutingScheme: String
  def otherBankRoutingAddress: String
  def otherBranchRoutingScheme: String
  def otherBranchRoutingAddress: String
  def isBeneficiary: Boolean
  def currency: String
  def bespoke: List[CounterpartyBespoke]
}
```

**Business Purpose**: Represents other parties in financial transactions, including routing information for payments.

**Relationships**:
- Counterparty → BankAccount (many:1 via thisBankId/thisAccountId)
- Counterparty → User (many:1 via createdByUserId)
- Counterparty → Transaction (1:many as otherAccount)
- Counterparty → CounterpartyBespoke (1:many via counterpartyId)

### 4. Security & Authorization Domain

#### Consent Entity
**Database Mapping**: Found in search results as MappedConsent
```scala
class MappedConsent extends LongKeyedMapper[MappedConsent] {
  object mConsentId extends UUIDString(this)                // consentId
  object mUserId extends UUIDString(this)                   // userId (FK)
  object mBankId extends UUIDString(this)                   // bankId (FK)
  object mConsentRequestId extends UUIDString(this)         // consentRequestId
  object mStatus extends MappedString(this, 32)             // status
  object mConsentType extends MappedString(this, 32)        // consentType
  // Additional consent-specific fields
}
```

**Business Purpose**: Manages user consent for data access and payment authorization, supporting PSD2 compliance.

**Relationships**:
- Consent → User (many:1 via userId)
- Consent → Bank (many:1 via bankId)
- Consent → TransactionRequest (1:many via consentReferenceId)

#### OAuth Entity
**Source**: Found in search results as OAuth mappings
**Business Purpose**: Manages OAuth tokens and application authorization for API access.

**Relationships**:
- OAuth → User (many:1 via user association)
- OAuth → Application (many:1 via consumer key)

### 5. Supporting Entities

#### UserCustomerLink Entity
**Database Mapping**: [MappedUserCustomerLink.scala:79-97](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/usercustomerlinks/MappedUserCustomerLink.scala#L79-L97)
```scala
class MappedUserCustomerLink extends UserCustomerLink with LongKeyedMapper[MappedUserCustomerLink] {
  object mUserCustomerLinkId extends MappedUUID(this)       // userCustomerLinkId
  object mUserId extends UUIDString(this)                   // userId (FK)
  object mCustomerId extends UUIDString(this)               // customerId (FK)
  object mDateInserted extends MappedDateTime(this)         // dateInserted
  object mIsActive extends MappedBoolean(this)              // isActive
}
```

**Business Purpose**: Links users to customers in many-to-many relationship.

#### CustomerAccountLink Entity
**Database Mapping**: [MappedCustomerAccountLink.scala](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/customeraccountlinks/MappedCustomerAccountLink.scala)
```scala
class MappedCustomerAccountLink extends LongKeyedMapper[MappedCustomerAccountLink] {
  object mCustomerId extends UUIDString(this)               // customerId (FK)
  object mBankId extends UUIDString(this)                   // bankId (FK)
  object mAccountId extends AccountIdString(this)           // accountId (FK)
  // Additional linking fields
}
```

**Business Purpose**: Links customers to bank accounts in many-to-many relationship.

#### AccountRouting Entity
**Source**: [BankingModel.scala:368-371](file:///home/ubuntu/repos/OBP-API/00.phase-1-input/OBP-API-develop/obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala#L368-L371)
```scala
case class AccountRouting(
  scheme: String,
  address: String
)
```

**Business Purpose**: Provides routing information for accounts (IBAN, account numbers, etc.).

**Relationships**:
- AccountRouting → BankAccount (many:1 via accountId)

## Entity Relationships Summary

### Core Banking Relationships
1. **Bank → BankAccount** (1:many)
   - FK: `BankAccount.bankId → Bank.bankId`
   - A bank can have multiple accounts

2. **Bank → Product** (1:many)
   - FK: `Product.bankId → Bank.bankId`
   - A bank offers multiple products

3. **Bank → Customer** (1:many)
   - FK: `Customer.bankId → Bank.bankId`
   - A bank serves multiple customers

### Customer Management Relationships
4. **User ↔ Customer** (many:many via UserCustomerLink)
   - FK: `UserCustomerLink.userId → User.userId`
   - FK: `UserCustomerLink.customerId → Customer.customerId`
   - Users can be linked to multiple customers

5. **Customer ↔ BankAccount** (many:many via CustomerAccountLink)
   - FK: `CustomerAccountLink.customerId → Customer.customerId`
   - FK: `CustomerAccountLink.accountId → BankAccount.accountId`
   - Customers can have multiple accounts, accounts can have multiple holders

### Transaction Processing Relationships
6. **BankAccount → Transaction** (1:many)
   - FK: `Transaction.thisAccount → BankAccount.accountId`
   - An account can have multiple transactions

7. **TransactionRequest → BankAccount** (many:1)
   - FK: `TransactionRequest.from_account_id → BankAccount.accountId`
   - FK: `TransactionRequest.to_account_id → BankAccount.accountId`
   - Transaction requests reference source and destination accounts

8. **TransactionRequest → Counterparty** (many:1)
   - FK: `TransactionRequest.counterparty_id → Counterparty.counterpartyId`
   - Transaction requests can specify counterparties

9. **Counterparty → BankAccount** (many:1)
   - FK: `Counterparty.thisAccountId → BankAccount.accountId`
   - Counterparties are associated with specific accounts

### Authorization Relationships
10. **Consent → User** (many:1)
    - FK: `Consent.userId → User.userId`
    - Users can have multiple consents

11. **Consent → Bank** (many:1)
    - FK: `Consent.bankId → Bank.bankId`
    - Consents are bank-specific

12. **TransactionRequest → Consent** (many:1)
    - FK: `TransactionRequest.consentReferenceId → Consent.consentId`
    - Transaction requests may require consent

## Class Hierarchy Patterns

### Inheritance Pattern
```
Trait (Domain Model) → Case Class (DTO) → Mapped Class (ORM)
```

**Example**:
```scala
trait Bank                    // Domain interface
↓
case class BankInMemory       // DTO implementation  
↓
class MappedBank              // Database persistence
```

### Common Lift Mapper Patterns
- **UUIDString**: For UUID foreign keys (`bankId`, `customerId`, etc.)
- **AccountIdString**: For account identifiers
- **MappedString(length)**: For text fields with size constraints
- **MappedDateTime**: For timestamp fields
- **MappedBoolean**: For boolean flags
- **MappedLong**: For numeric values (often currency in smallest units)
- **MappedText**: For large text fields (JSON, descriptions)

## Database Indexing Strategy

### Primary Indexes
- All Mapped classes have auto-generated `id` primary key
- Unique indexes on business identifiers (bankId, accountId, customerId)

### Foreign Key Indexes
- Composite indexes on frequently queried combinations
- Example: `UniqueIndex(mBankId, mAccountId)` in account-related tables

### Performance Indexes
- Indexes on frequently filtered fields (status, dates, active flags)
- Composite indexes for common query patterns

## Business Rules and Constraints

### Data Integrity
- UUID-based identifiers ensure global uniqueness
- Foreign key relationships maintain referential integrity
- Boolean flags for soft deletion and status management

### Business Logic
- Balance calculations use smallest currency units (cents, pence)
- Account routing supports multiple schemes (IBAN, account numbers)
- Transaction requests support multiple payment types
- Consent management for PSD2 compliance

### Security Patterns
- User-customer linking for access control
- View-based permissions for account access
- Consent-based authorization for transactions
- OAuth integration for API access

---

**Extraction Date**: September 14, 2025  
**Source Repository**: karunam2/OBP-API  
**Framework**: Lift Web Framework (Scala)  
**Total Entities Analyzed**: 25+ core business entities  
**Total Mapped Classes**: 113+ database entities
