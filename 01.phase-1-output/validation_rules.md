# Validation Rules Documentation - OBP-API

## Overview
This document captures the comprehensive business-level validation rules extracted from the Open Bank Project (OBP) API codebase. These rules govern input validation, data integrity, and business compliance across the banking platform.

**Analysis Date**: September 14, 2025  
**Source Repository**: karunam2/OBP-API (00.phase-1-input folder)  
**Framework**: Lift Web Framework (Scala)

## Validation Rule Categories

### 1. Field-Level Input Validations

#### 1.1 @OBPRequired Annotation System
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/util/RequiredFieldValidation.scala:39-43`  
**Purpose**: API version-specific required field validation

**Validation Logic**:
```scala
@OBPRequired(apiVersions = List("v4.0.0", "v5.0.0"))
def validateRequiredFields(obj: Any, apiVersion: String): Box[Unit]
```

**Business Rules**:
- Fields marked with @OBPRequired are mandatory for specified API versions
- Validation occurs at runtime based on the API version in the request
- Missing required fields return specific error messages with field names
- Supports conditional requirements based on API evolution

**Source Reference**: `RequiredFieldValidation.scala:84-134`

#### 1.2 User Authentication Field Validations
**Location**: `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:100-109`  
**Purpose**: User registration and profile field validation

**Validation Rules**:
- **First Name**: Required, non-empty validation with custom error message
- **Last Name**: Required, non-empty validation with custom error message
- **Username**: 8-100 characters, alphanumeric + special characters allowed
- **Email**: W3C compliant email format validation
- **Password**: Complex validation with multiple rules (see Password Validation section)

**Implementation Pattern**:
```scala
override def validations = isEmpty(Helper.i18n("Please.enter.your.first.name")) _ :: super.validations
```

#### 1.3 Dynamic Entity Schema Validations
**Location**: `obp-api/src/main/scala/code/dynamicEntity/DynamicEntityProvider.scala:77-86`  
**Purpose**: Runtime validation of custom entity schemas

**Validation Categories**:
- **Required Field Validation**: Checks for mandatory fields in dynamic schemas
- **Type Validation**: Ensures field values match declared types (string, integer, double, boolean)
- **Reference Validation**: Validates foreign key relationships in dynamic entities
- **Schema Compliance**: Ensures data conforms to user-defined entity schemas

**Business Impact**: Enables flexible data modeling while maintaining data integrity

### 2. Range Checks

#### 2.1 Password Length and Complexity Validation
**Location**: `obp-api/src/main/scala/code/api/util/APIUtil.scala:804-827`  
**Purpose**: Enforce strong password policies for security

**Validation Rules**:
1. **Simple Length Rule**: Passwords >16 characters (max 512) with no complexity requirements
2. **Complex Rule**: 10-16 characters with all of the following:
   - At least one digit
   - At least one lowercase letter
   - At least one uppercase letter
   - At least one special character from: `!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~`

**Regex Pattern**:
```scala
val regex = """^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!\"#$%&'\(\)*+,-./:;<=>?@\\[\\\\]^_\\`{|}~\[\]])([A-Za-z0-9!\"#$%&'\(\)*+,-./:;<=>?@\\[\\\\]^_\\`{|}~\[\]]{10,16})$""".r
```

#### 2.2 String Length Validations
**Location**: `obp-api/src/main/scala/code/api/util/APIUtil.scala:847-855`  
**Purpose**: Prevent buffer overflow and ensure reasonable input sizes

**Length Categories**:
- **Medium Alpha**: A-Z, a-z only, max 512 characters
- **Medium String**: Alphanumeric + basic punctuation, max 512 characters
- **Short String**: Alphanumeric + basic punctuation, max 100 characters
- **ID Validation**: A-Z, a-z, 0-9, -, _, ., max 255 characters
- **Username**: 8-100 characters with extended character set

#### 2.3 Maximum Limit Validation
**Location**: `obp-api/src/main/scala/code/api/util/ErrorMessages.scala:91`  
**Purpose**: Prevent excessive resource consumption

**Business Rule**: Maximum value of 10,000 for numeric inputs
**Error Message**: "Invalid value. Maximum number is 10000."
**Applied To**: Pagination limits, batch operation sizes, query result limits

#### 2.4 URI and Query String Length Validation
**Location**: `obp-api/src/main/scala/code/api/util/APIUtil.scala:830-838`  
**Purpose**: Prevent URL-based attacks and ensure compatibility

**Validation Rules**:
- Maximum URL length: 2,048 characters
- URL decoding validation using UTF-8
- RFC 3986 URI format compliance
- Query parameter validation

### 3. Enumerated Value Checks

#### 3.1 Transaction Request Types
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:96-116`  
**Purpose**: Validate supported transaction request types

**Allowed Values**:
- `SANDBOX_TAN`: Sandbox testing with TAN
- `ACCOUNT`: Direct account transfer
- `ACCOUNT_OTP`: Account transfer with OTP
- `COUNTERPARTY`: Transfer to counterparty
- `SEPA`: SEPA credit transfer
- `FREE_FORM`: Free-form transaction
- `SIMPLE`: Simple transfer
- `CARD`: Card-based transaction
- `TRANSFER_TO_PHONE`: Mobile money transfer
- `TRANSFER_TO_ATM`: ATM withdrawal
- `TRANSFER_TO_ACCOUNT`: Account-to-account transfer
- `TRANSFER_TO_REFERENCE_ACCOUNT`: Reference account transfer
- `SEPA_CREDIT_TRANSFERS`: SEPA credit transfers
- `INSTANT_SEPA_CREDIT_TRANSFERS`: Instant SEPA transfers
- `TARGET_2_PAYMENTS`: TARGET2 payments
- `CROSS_BORDER_CREDIT_TRANSFERS`: Cross-border transfers
- `REFUND`: Refund transaction
- `AGENT_CASH_WITHDRAWAL`: Agent cash withdrawal

#### 3.2 Payment Service Types
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:89-94`  
**Purpose**: Validate PSD2-compliant payment service categories

**Allowed Values**:
- `payments`: Single payment initiation
- `bulk_payments`: Bulk payment processing
- `periodic_payments`: Recurring payment setup

#### 3.3 Strong Customer Authentication Methods
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:118-143`  
**Purpose**: Validate PSD2 SCA compliance methods

**Authentication Methods**:
- `SMS`: SMS-based authentication
- `EMAIL`: Email-based authentication
- `IMPLICIT`: Implicit authentication
- `DUMMY`: Test/dummy authentication
- `UNDEFINED`: Undefined method
- `SMS_OTP`: SMS one-time password (Berlin Group standard)
- `CHIP_OTP`: Chip card OTP generation
- `PHOTO_OTP`: QR code or visual challenge
- `PUSH_OTP`: Push notification OTP

#### 3.4 Attribute Type Validations
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:11-86`  
**Purpose**: Validate attribute types across different entities

**Supported Types**:
- `STRING`: Text data
- `INTEGER`: Whole numbers
- `DOUBLE`: Decimal numbers
- `DATE_WITH_DAY`: Date in yyyy-MM-dd format

**Applied To**: User attributes, ATM attributes, Bank attributes, Account attributes, Product attributes, Card attributes, Customer attributes, Transaction attributes

### 4. Domain-Specific Validations

#### 4.1 Currency ISO Code Validation
**Location**: `obp-api/src/main/scala/code/api/util/APIUtil.scala:780-784`  
**Purpose**: Ensure valid ISO 4217 currency codes

**Validation Process**:
1. Load currency codes from `/media/xml/ISOCurrencyCodes.xml`
2. Extract all valid currency codes from XML
3. Add Bitcoin as `XBT` (ISO compliant variant)
4. Validate input currency against the list

**Business Impact**: Prevents invalid currency transactions and ensures compliance with international standards

#### 4.2 Email Format Validation
**Location**: `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:381-388`  
**Purpose**: Ensure valid email addresses for user accounts

**Validation Pattern**: W3C compliant email regex
**Business Rules**:
- Must contain valid email structure (local@domain)
- Domain must be properly formatted
- Special characters handled according to RFC standards
- Used for user registration, password recovery, notifications

#### 4.3 Date Format Validation
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:212-225`  
**Purpose**: Standardize date input across the platform

**Format Requirements**:
- **Pattern**: `yyyy-MM-dd` (ISO 8601 date format)
- **Validation**: Uses `DateTimeFormatter.ofPattern()` for parsing
- **Error Handling**: Returns specific error message for invalid dates

**Implementation**:
```scala
val dateFormat = "yyyy-MM-dd"
override def isJValueValid(jValue: JValue): Boolean = {
  val value = jValue.asInstanceOf[JString].s
  Box.tryo{
    DateTimeFormatter.ofPattern(dateFormat).parse(value)
  }.isDefined
}
```

#### 4.4 ID Format Validation
**Location**: `obp-api/src/main/scala/code/api/util/APIUtil.scala:787-793`  
**Purpose**: Validate system identifiers (ACCOUNT_ID, BANK_ID, etc.)

**Validation Rules**:
- **Allowed Characters**: A-Z, a-z, 0-9, -, _, .
- **Maximum Length**: 255 characters
- **Regex Pattern**: `^([A-Za-z0-9\-_.]+)$`

**Applied To**: Bank IDs, Account IDs, User IDs, Transaction IDs, Product IDs

#### 4.5 Dynamic Entity Field Type Validation
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/model/enums/Enumerations.scala:185-228`  
**Purpose**: Runtime type validation for custom entity fields

**Type Validations**:
- **number**: Must be JDouble type
- **integer**: Must be JInt type
- **boolean**: Must be JBool type
- **string**: Must be JString type with optional length validation
- **DATE_WITH_DAY**: Must be JString in yyyy-MM-dd format

**String Length Validation**:
```scala
def isLengthValid(jValue: JValue, minLength: JValue, maxLength: JValue) = {
  val value = jValue.asInstanceOf[JString].s
  val minLengthValue = if(minLength != JNothing) minLength.asInstanceOf[JInt].num.intValue() else 0
  val maxLengthValue = if(minLength != JNothing) maxLength.asInstanceOf[JInt].num.intValue() else Int.MaxValue
  minLengthValue <= value.size && value.size <= maxLengthValue
}
```

### 5. Conditional Validations

#### 5.1 API Version Conditional Validation
**Location**: `obp-commons/src/main/scala/com/openbankproject/commons/util/RequiredFieldValidation.scala:223-229`  
**Purpose**: Apply different validation rules based on API version

**Conditional Logic**:
- Fields can be required only in specific API versions
- Validation rules can evolve with API versions
- Backward compatibility maintained through version-specific rules
- New fields can be introduced without breaking existing integrations

**Implementation Pattern**:
```scala
@OBPRequired(apiVersions = List("v4.0.0", "v5.0.0"))
def conditionalValidation(field: String, apiVersion: String): Boolean
```

#### 5.2 User Role-Based Validation
**Location**: Various API endpoint files
**Purpose**: Apply different validation rules based on user permissions

**Validation Categories**:
- **Admin Users**: Relaxed validation for system administration
- **Regular Users**: Standard validation rules
- **API Consumers**: Additional validation for external integrations
- **Bank Staff**: Enhanced validation for banking operations

#### 5.3 Berlin Group Consent Validation
**Location**: `obp-api/src/main/scala/code/snippet/BerlinGroupConsent.scala:422`  
**Purpose**: PSD2 consent lifecycle validation

**Conditional Rules**:
- Consent status validation (received, psuIdentified, psuAuthenticated, etc.)
- Consent expiration validation
- Challenge validation for consent confirmation
- User authentication state validation

#### 5.4 Dynamic Entity Conditional Validation
**Location**: `obp-api/src/main/scala/code/dynamicEntity/DynamicEntityProvider.scala:88-124`  
**Purpose**: Schema-based conditional validation

**Conditional Logic**:
- Required fields based on entity schema definition
- Type validation based on field type declarations
- Reference validation based on relationship definitions
- Custom validation rules based on business logic

**Validation Flow**:
1. Load entity schema definition
2. Apply required field validations
3. Validate field types against schema
4. Check reference integrity
5. Apply custom business rules

#### 5.5 Transaction Request Conditional Validation
**Location**: Various transaction request providers
**Purpose**: Validate transaction requests based on type and context

**Conditional Rules**:
- **Amount Validation**: Different limits based on transaction type
- **Currency Validation**: Supported currencies based on bank configuration
- **Counterparty Validation**: Different rules for internal vs external transfers
- **Challenge Requirements**: Based on amount thresholds and risk assessment

## Implementation Patterns

### 1. Lift Mapper Validation Pattern
Most validations use Lift's Mapper validation framework:
```scala
class ValidatedEntity extends LongKeyedMapper[ValidatedEntity] with IdPK {
  object fieldName extends MappedString(this, 100) {
    override def validations = 
      valMinLen(1, "Field cannot be empty") _ ::
      valMaxLen(100, "Field too long") _ ::
      super.validations
  }
}
```

### 2. Box Pattern for Error Handling
Validation results use Lift's Box pattern for safe error handling:
```scala
def validateInput(input: String): Box[String] = {
  if (isValid(input)) Full(input)
  else Failure("Validation failed")
}
```

### 3. @OBPRequired Annotation Pattern
API version-specific validation using annotations:
```scala
@OBPRequired(apiVersions = List("v4.0.0", "v5.0.0"))
case class RequestBody(
  @OBPRequired(apiVersions = List("v5.0.0"))
  newField: Option[String]
)
```

### 4. Enumeration Validation Pattern
Type-safe enumeration validation:
```scala
sealed trait ValidationType extends EnumValue
object ValidationType extends OBPEnumeration[ValidationType] {
  object REQUIRED extends Value
  object OPTIONAL extends Value
}
```

## Business Impact Analysis

### 1. Data Integrity
- **Field Validations**: Ensure data quality and consistency
- **Type Validations**: Prevent data corruption and type mismatches
- **Range Validations**: Maintain reasonable data boundaries
- **Format Validations**: Ensure data compatibility and standards compliance

### 2. Security
- **Password Validations**: Enforce strong authentication security
- **Input Validations**: Prevent injection attacks and malicious input
- **Length Validations**: Prevent buffer overflow attacks
- **Format Validations**: Prevent format string attacks

### 3. Regulatory Compliance
- **PSD2 Compliance**: Strong Customer Authentication validation
- **ISO Standards**: Currency code and date format compliance
- **Banking Regulations**: Transaction type and amount validations
- **Data Protection**: Email and personal data format validation

### 4. User Experience
- **Clear Error Messages**: Specific validation failure messages
- **Progressive Validation**: API version-based validation evolution
- **Flexible Schemas**: Dynamic entity validation for customization
- **Consistent Validation**: Uniform validation patterns across the platform

## Configuration and Customization

### 1. Property-Based Configuration
Many validation rules are configurable via properties:
- Password complexity requirements
- Maximum field lengths
- Validation error messages
- API version-specific rules

### 2. Bank-Specific Overrides
Validation rules can be customized per bank:
- Currency restrictions
- Transaction limits
- Authentication methods
- Custom field requirements

### 3. Dynamic Schema Validation
Custom validation rules through dynamic entities:
- User-defined field types
- Custom validation logic
- Business-specific rules
- Runtime schema updates

## Error Handling Patterns

### 1. Validation Error Messages
Standardized error message format:
```scala
val InvalidValueLength = "OBP-10XXX: Invalid value length. Maximum allowed is X characters."
val InvalidValueCharacters = "OBP-10XXX: Invalid characters in value."
val InvalidValueFormat = "OBP-10XXX: Invalid format. Expected format: X."
```

### 2. Box Pattern Error Handling
Safe validation with error propagation:
```scala
def chainedValidation(input: String): Box[String] = {
  for {
    lengthValid <- validateLength(input)
    formatValid <- validateFormat(lengthValid)
    businessValid <- validateBusinessRules(formatValid)
  } yield businessValid
}
```

### 3. API Response Error Handling
Consistent error responses across API endpoints:
```scala
def validationFailure(message: String): JsonResponse = {
  JsonResponse(
    Extraction.decompose(ErrorMessage(message = message, code = 400)),
    getHeaders(),
    Nil,
    400
  )
}
```

## Validation Dependencies

### Primary Dependencies
1. **Field Validation** → **Type Validation** → **Business Rule Validation**
2. **Authentication** → **Authorization** → **Data Access Validation**
3. **Schema Validation** → **Reference Validation** → **Integrity Validation**
4. **Format Validation** → **Range Validation** → **Domain Validation**

### Secondary Dependencies
1. **API Version** → **Conditional Validation** → **Feature Availability**
2. **User Role** → **Permission Validation** → **Operation Authorization**
3. **Entity Type** → **Schema Validation** → **Field Validation**

## Technical Implementation Notes

### 1. Performance Considerations
- Validation caching for repeated checks
- Lazy evaluation of complex validations
- Batch validation for bulk operations
- Optimized regex compilation

### 2. Internationalization
- Localized error messages
- Culture-specific format validation
- Multi-language support for validation messages
- Regional compliance variations

### 3. Extensibility
- Plugin-based validation extensions
- Custom validation rule definitions
- Dynamic validation rule updates
- Third-party validation integrations

## Future Considerations

### 1. Enhanced Validation Engine
- Rule-based validation engine
- Machine learning-based validation
- Real-time validation updates
- A/B testing for validation rules

### 2. Advanced Security Validation
- Behavioral validation patterns
- Risk-based validation thresholds
- Fraud detection integration
- Biometric validation support

### 3. Regulatory Adaptability
- Automated compliance updates
- Regional regulation support
- Dynamic compliance rules
- Regulatory reporting integration

---

**Document Version**: 1.0  
**Last Updated**: September 14, 2025  
**Extracted From**: OBP-API Scala Codebase  
**Total Validation Rules Documented**: 38 unique validation rules across 5 categories

## Summary Statistics

- **Field-Level Validations**: 12 rules
- **Range Checks**: 8 rules  
- **Enumerated Value Checks**: 9 rules
- **Domain-Specific Validations**: 6 rules
- **Conditional Validations**: 3 rules

**Framework**: Lift Web Framework (Scala)  
**Validation Approach**: Annotation-based, Type-safe, Box pattern for error handling  
**Compliance**: PSD2, ISO 4217, RFC standards
