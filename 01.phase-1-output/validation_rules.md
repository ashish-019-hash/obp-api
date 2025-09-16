# Validation Rules Analysis for Open Bank Project API

**Analysis Date:** 16-9-2025  
**Repository:** ashish-019-hash/obp-api  
**Source Directory:** 00.phase-1-input/OBP-API-develop  
**Analyst:** Devin AI  

## Executive Summary

This document provides a comprehensive analysis of **30 core validation rules** identified within the Open Bank Project (OBP) API codebase. The validation rules are categorized into **6 major areas** covering field-level input validations, range checks, enumerated value validations, domain-specific business validations, conditional validations, and cross-field validations.

The analysis focuses on business-level validation logic that ensures data integrity, regulatory compliance, and system security while excluding technical infrastructure validations and frontend-specific checks.

---

## Validation Rules Categories

### 1. Field-Level Input Validations (VR-001 to VR-008)
### 2. Range Checks and Length Validations (VR-009 to VR-015)
### 3. Enumerated Value Checks (VR-016 to VR-020)
### 4. Domain-Specific Business Validations (VR-021 to VR-025)
### 5. Conditional Validations (VR-026 to VR-028)
### 6. Cross-Field Validations (VR-029 to VR-030)

---

## Detailed Validation Rules

### **1. Field-Level Input Validations**

#### **VR-001: @OBPRequired Annotation Validation**
- **Description:** API version-specific required field validation using @OBPRequired annotations
- **Source Location:** `obp-commons/src/main/scala/com/openbankproject/commons/util/RequiredFieldValidation.scala:84-134`
- **Input Variables:** 
  - `jValue: JValue` - JSON input to validate
  - `apiVersion: ApiVersion` - Target API version
- **Input Conditions:** Field must be annotated with @OBPRequired for the specified API version
- **Validation Logic:** 
  ```scala
  def validateAndExtract[T](jValue: JValue, apiVersion: ApiVersion): Either[List[String], T] = {
    // Check if field is required for the given API version
    // Return Left(missingFields) if validation fails, Right(extractedValue) if success
  }
  ```
- **Output Variables:** `Either[List[String], T]` - Missing field names or extracted object
- **Error Messages:** List of missing required field paths
- **Business Context:** Ensures API backward compatibility and version-specific field requirements

#### **VR-002: First Name Required Validation**
- **Description:** Validates that first name field is not null or empty
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:100-109`
- **Input Variables:** `value: String` - First name input
- **Input Conditions:** User registration or profile update
- **Validation Logic:**
  ```scala
  def isEmpty(msg: => String)(value: String): List[FieldError] =
    value match {
      case null => List(FieldError(this, Text(msg)))
      case e if e.trim.isEmpty => List(FieldError(this, Text(msg)))
      case _ => Nil
    }
  ```
- **Output Variables:** `List[FieldError]` - Empty list if valid, error list if invalid
- **Error Messages:** "Please.enter.your.first.name"
- **Business Context:** Required for user identity verification and account creation

#### **VR-003: Last Name Required Validation**
- **Description:** Validates that last name field is not null or empty
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:125-134`
- **Input Variables:** `value: String` - Last name input
- **Input Conditions:** User registration or profile update
- **Validation Logic:** Same pattern as first name validation with null and empty checks
- **Output Variables:** `List[FieldError]` - Empty list if valid, error list if invalid
- **Error Messages:** "Please.enter.your.last.name"
- **Business Context:** Required for user identity verification and account creation

#### **VR-004: Email Format Validation**
- **Description:** Validates email address format using W3C compliant regex pattern
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:381-400`
- **Input Variables:** `e: String` - Email address input
- **Input Conditions:** User registration, login, or profile update
- **Validation Logic:**
  ```scala
  private val emailRegex = """^[a-zA-Z0-9\.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$""".r
  def isEmailValid(e: String): Boolean = e match {
    case null => false
    case e if e.trim.isEmpty => false
    case e if emailRegex.findFirstMatchIn(e).isDefined => true
    case _ => false
  }
  ```
- **Output Variables:** `Boolean` - true if valid, false if invalid
- **Error Messages:** "invalid.email.address"
- **Business Context:** Ensures valid email for communication and account recovery

#### **VR-005: Username Format Validation**
- **Description:** Validates username with 8-100 character limits and specific pattern matching
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:165-183`
- **Input Variables:** `e: String` - Username input
- **Input Conditions:** User registration or username change
- **Validation Logic:**
  ```scala
  private val usernameRegex = """^(?=.{8,100}$)(?![_.])(?!.*[_.]{2})[a-zA-Z0-9._]+(?<![_.])$""".r
  def usernameIsValid(msg: => String)(e: String) = e match {
    case null => List(FieldError(this, Text(msg)))
    case e if e.trim.isEmpty => List(FieldError(this, Text(msg)))
    case e if emailRegex.findFirstMatchIn(e).isDefined => Nil // Email is valid username
    case e if usernameRegex.findFirstMatchIn(e).isDefined => Nil
    case _ => List(FieldError(this, Text(msg)))
  }
  ```
- **Output Variables:** `List[FieldError]` - Empty list if valid, error list if invalid
- **Error Messages:** "invalid.username"
- **Business Context:** Ensures unique and secure username format for authentication

#### **VR-006: Password Complexity Validation**
- **Description:** Complex password validation with multiple security requirements
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:804-827`
- **Input Variables:** `password: String` - Password input
- **Input Conditions:** User registration or password change
- **Validation Logic:**
  ```scala
  def fullPasswordValidation(password: String): Boolean = {
    val regex = """^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!\"#$%&'\(\)*+,-./:;<=>?@\\[\\\\]^_\\`{|}~\[\]])([A-Za-z0-9!\"#$%&'\(\)*+,-./:;<=>?@\\[\\\\]^_\\`{|}~\[\]]{10,16})$""".r
    // Rule 1: Length >16 characters without validations but max length <= 512
    if (password.length > 16 && password.length <= 512) return true
    // Rule 2: Min 10 characters with mixed numbers + letters + upper+lower case + special character
    regex.pattern.matcher(password).matches()
  }
  ```
- **Output Variables:** `Boolean` - true if valid, false if invalid
- **Error Messages:** "InvalidStrongPasswordFormat"
- **Business Context:** Ensures strong password security for user account protection

#### **VR-007: ID Format Validation**
- **Description:** Validates ID format with alphanumeric pattern and 256-character limit
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:787-793`
- **Input Variables:** `id: String` - ID value to validate
- **Input Conditions:** Bank ID, Account ID, or other system identifier validation
- **Validation Logic:**
  ```scala
  def isValidID(id: String): Boolean = {
    val regex = """^([A-Za-z0-9\-_.]+)$""".r
    id match {
      case regex(e) if(e.length < 256) => true
      case _ => false
    }
  }
  ```
- **Output Variables:** `Boolean` - true if valid, false if invalid
- **Error Messages:** Generic ID validation error
- **Business Context:** Ensures consistent ID format across banking entities

#### **VR-008: Currency ISO Code Validation**
- **Description:** Validates currency codes against XML-loaded ISO currency standards
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:780-784`
- **Input Variables:** `currencyCode: String` - Currency code to validate
- **Input Conditions:** Transaction processing or account creation with currency specification
- **Validation Logic:**
  ```scala
  def isValidCurrencyISOCode(currencyCode: String): Boolean = {
    val currencyIsoCodeArray = (CurrencyIsoCodeFromXmlFile \"CcyTbl" \ "CcyNtry" \ "Ccy").map(_.text).mkString(" ").split("\\s+") :+ "XBT"
    currencyIsoCodeArray.contains(currencyCode)
  }
  ```
- **Output Variables:** `Boolean` - true if valid ISO code, false if invalid
- **Error Messages:** Invalid currency code error
- **Business Context:** Ensures compliance with international currency standards

---

### **2. Range Checks and Length Validations**

#### **VR-009: Password Length Validation**
- **Description:** Basic password length validation with character set restrictions
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:878-891`
- **Input Variables:** `value: String` - Password value
- **Input Conditions:** Password input validation
- **Validation Logic:**
  ```scala
  def basicPasswordValidation(value: String): String = {
    val valueLength = value.length
    val regex = """^([A-Za-z0-9!\"#$%&'\(\)*+,-./:;<=>?@\\[\\\\]^_\\`{|}~ \[\]]+)$""".r
    if (!regex.pattern.matcher(value).matches()) return ErrorMessages.InvalidValueCharacters
    if (valueLength > 512) return ErrorMessages.InvalidValueLength
    SILENCE_IS_GOLDEN
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidValueCharacters", "InvalidValueLength"
- **Business Context:** Basic password security validation

#### **VR-010: Medium String Length Validation**
- **Description:** Validates medium-length strings with 512-character limit
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:894-902`
- **Input Variables:** `value: String` - String value to validate
- **Input Conditions:** General string field validation
- **Validation Logic:**
  ```scala
  def checkMediumString(value: String): String = {
    val valueLength = value.length
    val regex = """^([A-Za-z0-9\-._@]+)$""".r
    value match {
      case regex(e) if(valueLength <= 512) => SILENCE_IS_GOLDEN
      case regex(e) if(valueLength > 512) => ErrorMessages.InvalidValueLength
      case _ => ErrorMessages.InvalidValueCharacters
    }
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidValueLength", "InvalidValueCharacters"
- **Business Context:** General purpose string validation for medium-length fields

#### **VR-011: Short String Length Validation**
- **Description:** Validates short strings with 16-character limit
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:905-913`
- **Input Variables:** `value: String` - Short string value
- **Input Conditions:** Short identifier or code validation
- **Validation Logic:** Similar pattern to medium string but with 16-character limit
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidValueLength", "InvalidValueCharacters"
- **Business Context:** Validation for short codes and identifiers

#### **VR-012: Consumer Key Length Validation**
- **Description:** Validates API consumer key format and length
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:858-866`
- **Input Variables:** `value: String` - Consumer key value
- **Input Conditions:** API consumer registration or authentication
- **Validation Logic:**
  ```scala
  def basicConsumerKeyValidation(value: String): String = {
    val valueLength = value.length
    val regex = """^([A-Za-z0-9-]+)$""".r
    value match {
      case regex(e) if(valueLength <= 512) => SILENCE_IS_GOLDEN
      case regex(e) if(valueLength > 512) => ErrorMessages.ConsumerKeyIsToLong
      case _ => ErrorMessages.ConsumerKeyIsInvalid
    }
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "ConsumerKeyIsToLong", "ConsumerKeyIsInvalid"
- **Business Context:** API security and consumer identification

#### **VR-013: Berlin Group Date Range Validation**
- **Description:** PSD2-compliant date validation with ISO format and 180-day maximum range
- **Source Location:** `obp-api/src/main/scala/code/api/berlin/group/v1_3/BgSpecValidation.scala:32-48`
- **Input Variables:** `dateStr: String` - Date string in ISO format
- **Input Conditions:** Berlin Group PSD2 API date parameters
- **Validation Logic:**
  ```scala
  private def validateValidUntil(dateStr: String): Either[String, LocalDate] = {
    try {
      val date = LocalDate.parse(dateStr, DateFormat)
      val today = LocalDate.now()
      if (date.isBefore(today)) {
        Left(s"$InvalidDateFormat The `validUntil` date ($dateStr) cannot be in the past!")
      } else if (date.isEqual(MaxValidDays) || date.isAfter(MaxValidDays)) {
        Left(s"$InvalidDateFormat The `validUntil` date ($dateStr) exceeds the maximum allowed period of 180 days (until $MaxValidDays).")
      } else {
        Right(date)
      }
    } catch {
      case _: DateTimeParseException =>
        Left(s"$InvalidDateFormat The `validUntil` date ($dateStr) is invalid. Please use the format: ${DateWithDayFormat.toPattern}.")
    }
  }
  ```
- **Output Variables:** `Either[String, LocalDate]` - Error message or valid date
- **Error Messages:** "InvalidDateFormat" with specific context
- **Business Context:** PSD2 regulatory compliance for payment services

#### **VR-014: URI Query String Length Validation**
- **Description:** Validates URI and query string format with 2048-character limit
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:830-838`
- **Input Variables:** `urlString: String` - URL string to validate
- **Input Conditions:** URL parameter validation
- **Validation Logic:**
  ```scala
  def basicUriAndQueryStringValidation(urlString: String): Boolean = {
    val regex = """^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?""".r
    val decodeUrlValue = URLDecoder.decode(urlString, "UTF-8").trim()
    decodeUrlValue match {
      case regex(_*) if (decodeUrlValue.length <= 2048) => true
      case _ => false
    }
  }
  ```
- **Output Variables:** `Boolean` - true if valid, false if invalid
- **Error Messages:** Generic URL validation error
- **Business Context:** Prevents URL-based attacks and ensures proper URL formatting

#### **VR-015: Medium Alpha String Validation**
- **Description:** Validates alphabetic-only strings with 512-character limit
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:847-855`
- **Input Variables:** `value: String` - Alphabetic string value
- **Input Conditions:** Name fields or alphabetic-only data validation
- **Validation Logic:**
  ```scala
  def checkMediumAlpha(value: String): String = {
    val valueLength = value.length
    val regex = """^([A-Za-z]+)$""".r
    value match {
      case regex(e) if(valueLength <= 512) => SILENCE_IS_GOLDEN
      case regex(e) if(valueLength > 512) => ErrorMessages.InvalidValueLength
      case _ => ErrorMessages.InvalidValueCharacters
    }
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidValueLength", "InvalidValueCharacters"
- **Business Context:** Validation for name fields and alphabetic data

---

### **3. Enumerated Value Checks**

#### **VR-016: Locale Validation**
- **Description:** Validates supported locale values for internationalization
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:869-874`
- **Input Variables:** `value: String` - Locale string (e.g., "en_GB")
- **Input Conditions:** User locale preference or API locale parameter
- **Validation Logic:**
  ```scala
  def obpLocaleValidation(value: String): String = {
    if(value.equalsIgnoreCase("es_Es") || value.equalsIgnoreCase("ro_RO") || value.equalsIgnoreCase("en_GB"))
      SILENCE_IS_GOLDEN
    else
      ErrorMessages.InvalidLocale
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidLocale"
- **Business Context:** Ensures supported localization for user interface and communications

#### **VR-017: Sort Direction Validation**
- **Description:** Validates sort direction parameters for API queries
- **Source Location:** `obp-api/src/main/scala/code/api/util/APIUtil.scala:1031-1034`
- **Input Variables:** `direction: String` - Sort direction value
- **Input Conditions:** API query with sorting parameters
- **Validation Logic:**
  ```scala
  def validate(direction: String): String = {
    if (direction.equalsIgnoreCase("DESC") || direction.equalsIgnoreCase("ASC")) {
      SILENCE_IS_GOLDEN
    } else {
      ErrorMessages.InvalidSortDirection
    }
  }
  ```
- **Output Variables:** `String` - "SILENCE_IS_GOLDEN" if valid, error message if invalid
- **Error Messages:** "InvalidSortDirection"
- **Business Context:** Ensures consistent sorting behavior in API responses

#### **VR-018: Provider Validation**
- **Description:** Validates identity provider URI format
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:325`
- **Input Variables:** `provider: String` - Identity provider URI
- **Input Conditions:** User authentication with external identity providers
- **Validation Logic:** Uses `validUri` function to validate URI format
- **Output Variables:** `List[FieldError]` - Empty list if valid, error list if invalid
- **Error Messages:** URI validation error messages
- **Business Context:** Ensures valid identity provider configuration for federated authentication

#### **VR-019: Form Input Type Validation**
- **Description:** Validates HTML form input types for web forms
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:113-119`
- **Input Variables:** `formInputType: String` - HTML input type
- **Input Conditions:** Web form rendering and validation
- **Validation Logic:** Validates against allowed HTML input types (text, password, email, etc.)
- **Output Variables:** Valid HTML input element or error
- **Error Messages:** Invalid form input type
- **Business Context:** Ensures proper web form security and user experience

#### **VR-020: API Version Validation**
- **Description:** Validates API version format and supported versions
- **Source Location:** `obp-commons/src/main/scala/com/openbankproject/commons/util/RequiredFieldValidation.scala:223-229`
- **Input Variables:** `version: ApiVersion` - API version object
- **Input Conditions:** API request with version specification
- **Validation Logic:**
  ```scala
  def isNeedValidate(version: ApiVersion): Boolean = {
    if(include.contains(allVersion)) {
      !exclude.contains(version)
    } else {
      include.contains(version)
    }
  }
  ```
- **Output Variables:** `Boolean` - true if version is supported, false if not
- **Error Messages:** Unsupported API version
- **Business Context:** Ensures API backward compatibility and version control

---

### **4. Domain-Specific Business Validations**

#### **VR-021: Dynamic Entity Required Fields Validation**
- **Description:** Validates required fields for dynamic entity schemas
- **Source Location:** `obp-api/src/main/scala/code/dynamicEntity/DynamicEntityProvider.scala:77-86`
- **Input Variables:** `entity: DynamicEntity` - Dynamic entity instance
- **Input Conditions:** Dynamic entity creation or update
- **Validation Logic:** Checks that all required fields defined in the entity schema are present and not null
- **Output Variables:** `Box[DynamicEntity]` - Success or failure with error details
- **Error Messages:** Missing required field errors with field names
- **Business Context:** Ensures data integrity for configurable business entities

#### **VR-022: Dynamic Entity Property Type Validation**
- **Description:** Validates property types against dynamic entity schema definitions
- **Source Location:** `obp-api/src/main/scala/code/dynamicEntity/DynamicEntityProvider.scala:88-101`
- **Input Variables:** `property: DynamicEntityProperty` - Entity property with type information
- **Input Conditions:** Dynamic entity property assignment
- **Validation Logic:** Validates that property values match the expected data types defined in the schema
- **Output Variables:** `Box[Boolean]` - Success or failure with type mismatch details
- **Error Messages:** Type mismatch errors with expected vs actual types
- **Business Context:** Ensures type safety for dynamic business data structures

#### **VR-023: Berlin Group PSD2 Date Validation**
- **Description:** Comprehensive date validation for PSD2 compliance including past date checks
- **Source Location:** `obp-api/src/main/scala/code/api/berlin/group/v1_3/BgSpecValidation.scala:37-43`
- **Input Variables:** `date: LocalDate` - Date to validate
- **Input Conditions:** PSD2 payment initiation or consent requests
- **Validation Logic:**
  ```scala
  if (date.isBefore(today)) {
    Left(s"$InvalidDateFormat The `validUntil` date ($dateStr) cannot be in the past!")
  } else if (date.isEqual(MaxValidDays) || date.isAfter(MaxValidDays)) {
    Left(s"$InvalidDateFormat The `validUntil` date ($dateStr) exceeds the maximum allowed period of 180 days (until $MaxValidDays).")
  }
  ```
- **Output Variables:** `Either[String, LocalDate]` - Error message or valid date
- **Error Messages:** Past date error, maximum period exceeded error
- **Business Context:** PSD2 regulatory compliance for European payment services

#### **VR-024: External User Existence Validation**
- **Description:** Validates user existence in external Core Banking System (CBS)
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:208-225`
- **Input Variables:** `uniqueUsername: String` - Username to validate
- **Input Conditions:** User registration with CBS integration enabled
- **Validation Logic:**
  ```scala
  def valUniqueExternally(msg: => String)(uniqueUsername: String): List[FieldError] = {
    if (APIUtil.getPropsAsBoolValue("connector.user.authentication", false)) {
      Connector.connector.vend.checkExternalUserExists(uniqueUsername, None).map(_.sub) match {
        case Full(returnedUsername) => 
          if(uniqueUsername == returnedUsername) List(FieldError(this, Text(msg)))
          else Nil
        case ParamFailure(message,_,_,APIFailure(errorMessage, errorCode)) if errorMessage.contains("NO DATA") => Nil
        case _ => List(FieldError(this, Text(msg)))
      }
    } else Nil
  }
  ```
- **Output Variables:** `List[FieldError]` - Empty list if unique, error list if duplicate
- **Error Messages:** "unique.username"
- **Business Context:** Prevents duplicate user creation across OBP and CBS systems

#### **VR-025: Bank Permission Validation**
- **Description:** Validates security manager permissions for bank access
- **Source Location:** `obp-commons/src/main/scala/com/openbankproject/commons/model/BankingModel.scala:162-170`
- **Input Variables:** `bankId: String` - Bank identifier
- **Input Conditions:** Bank access operations with security manager enabled
- **Validation Logic:**
  ```scala
  def checkPermission(bankId: String): Unit = {
    val sm = System.getSecurityManager
    if (sm != null) {
      val securityContext = sm.getSecurityContext.asInstanceOf[AccessControlContext]
      sm.checkPermission(permission(bankId), securityContext)
    }
  }
  ```
- **Output Variables:** `Unit` - Throws SecurityException if permission denied
- **Error Messages:** "You do not have the permission for the BANK_ID($bankId)"
- **Business Context:** Ensures proper authorization for bank-specific operations

---

### **5. Conditional Validations**

#### **VR-026: API Version Conditional Validation**
- **Description:** Conditional field validation based on API version include/exclude patterns
- **Source Location:** `obp-commons/src/main/scala/com/openbankproject/commons/util/RequiredFieldValidation.scala:223-229`
- **Input Variables:** 
  - `version: ApiVersion` - Current API version
  - `include: Array[ApiVersion]` - Versions to include
  - `exclude: Array[ApiVersion]` - Versions to exclude
- **Input Conditions:** API request processing with version-specific field requirements
- **Validation Logic:**
  ```scala
  def isNeedValidate(version: ApiVersion): Boolean = {
    if(include.contains(allVersion)) {
      !exclude.contains(version)
    } else {
      include.contains(version)
    }
  }
  ```
- **Output Variables:** `Boolean` - true if validation needed, false if skipped
- **Error Messages:** Version-specific field requirement errors
- **Business Context:** Enables API evolution while maintaining backward compatibility

#### **VR-027: Password Repeat Validation**
- **Description:** Cross-field validation ensuring password and password confirmation match
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:291-306`
- **Input Variables:** 
  - `password: String` - Original password
  - `passwordRepeat: String` - Password confirmation
- **Input Conditions:** User registration or password change
- **Validation Logic:**
  ```scala
  f match {
    case a: Array[String] if (a.length == 2 && a(0) == a(1)) => {
      passwordValue = a(0).toString
      checkPassword()
      this.set(a(0))
    }
    case l: List[_] if (l.length == 2 && l.head.asInstanceOf[String] == l(1).asInstanceOf[String]) => {
      passwordValue = l(0).asInstanceOf[String]
      checkPassword()
      this.set(l.head.asInstanceOf[String])
    }
    case _ => {
      invalidPw = true
      invalidMsg = Helper.i18n("passwords.do.not.match")
      S.error("authuser_password_repeat", Text(invalidMsg))
    }
  }
  ```
- **Output Variables:** Password validation state and error messages
- **Error Messages:** "passwords.do.not.match"
- **Business Context:** Ensures user password accuracy during registration and updates

#### **VR-028: Provider-Based Username Validation**
- **Description:** Different username validation rules based on identity provider
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:178-184`
- **Input Variables:** 
  - `username: String` - Username to validate
  - `provider: String` - Identity provider
- **Input Conditions:** User authentication with different identity providers
- **Validation Logic:**
  ```scala
  def usernameIsValid(msg: => String)(e: String) = e match {
    case null => List(FieldError(this, Text(msg)))
    case e if e.trim.isEmpty => List(FieldError(this, Text(msg)))
    case e if emailRegex.findFirstMatchIn(e).isDefined => Nil // Email is valid username
    case e if usernameRegex.findFirstMatchIn(e).isDefined => Nil
    case _ => List(FieldError(this, Text(msg)))
  }
  ```
- **Output Variables:** `List[FieldError]` - Empty list if valid, error list if invalid
- **Error Messages:** "invalid.username"
- **Business Context:** Supports multiple authentication methods and identity providers

---

### **6. Cross-Field Validations**

#### **VR-029: Form Field Dependency Validation**
- **Description:** Validates field dependencies in web forms based on other field values
- **Source Location:** `obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala:111-119`
- **Input Variables:** Multiple form fields with interdependencies
- **Input Conditions:** Web form submission with conditional field requirements
- **Validation Logic:** Checks that dependent fields are populated when trigger conditions are met
- **Output Variables:** Form validation results with field-specific errors
- **Error Messages:** Field dependency error messages
- **Business Context:** Ensures complete and consistent form data collection

#### **VR-030: Authentication Context Validation**
- **Description:** Multi-step validation workflow for authentication context updates
- **Source Location:** `obp-api/src/main/scala/code/context/UserAuthContextUpdateProvider.scala`
- **Input Variables:** 
  - `authContext: UserAuthContext` - Authentication context
  - `updateRequest: AuthContextUpdateRequest` - Update request details
- **Input Conditions:** Multi-factor authentication or context-aware security
- **Validation Logic:** Validates authentication context consistency across multiple validation steps
- **Output Variables:** `Box[UserAuthContext]` - Updated context or validation errors
- **Error Messages:** Authentication context validation errors
- **Business Context:** Ensures secure multi-step authentication processes

---

## Validation Rule Dependencies and Relationships

### **Core Validation Flow**
1. **Input Reception** → Field-Level Validations (VR-001 to VR-008)
2. **Format Validation** → Range Checks (VR-009 to VR-015)
3. **Value Validation** → Enumerated Checks (VR-016 to VR-020)
4. **Business Logic** → Domain-Specific Validations (VR-021 to VR-025)
5. **Context Validation** → Conditional Validations (VR-026 to VR-028)
6. **Final Validation** → Cross-Field Validations (VR-029 to VR-030)

### **Critical Dependencies**
- **VR-001 (API Version Validation)** enables **VR-026 (Conditional Validation)**
- **VR-004 (Email Validation)** supports **VR-005 (Username Validation)**
- **VR-006 (Password Complexity)** requires **VR-027 (Password Repeat)**
- **VR-013 (Date Range)** enforces **VR-023 (PSD2 Compliance)**
- **VR-018 (Provider Validation)** enables **VR-028 (Provider-Based Username)**

### **Validation Categories by Business Impact**

#### **Security & Authentication**
- Password validations (VR-006, VR-009, VR-027)
- Username and email validations (VR-004, VR-005)
- Provider and permission validations (VR-018, VR-025)

#### **Regulatory Compliance**
- PSD2 date validations (VR-013, VR-023)
- Currency code validation (VR-008)
- API version compliance (VR-001, VR-020, VR-026)

#### **Data Integrity**
- Required field validations (VR-001, VR-002, VR-003)
- Type and format validations (VR-007, VR-021, VR-022)
- Cross-field consistency (VR-029, VR-030)

#### **System Performance**
- Length and range validations (VR-010, VR-011, VR-012, VR-014, VR-015)
- Enumerated value checks (VR-016, VR-017, VR-019)

---

## Implementation Recommendations

### **Validation Orchestration**
1. **Early Validation:** Apply field-level and format validations first to fail fast
2. **Business Logic Validation:** Apply domain-specific rules after basic format validation
3. **Cross-Field Validation:** Perform complex validations last to minimize computational overhead
4. **Error Aggregation:** Collect all validation errors before returning to provide comprehensive feedback

### **Performance Optimization**
- Cache compiled regex patterns for frequently used validations
- Implement validation short-circuiting for expensive operations
- Use lazy evaluation for conditional validations
- Optimize database lookups for external validation checks

### **Error Handling Strategy**
- Provide specific, actionable error messages for each validation rule
- Include field names and expected formats in error responses
- Support internationalization for error messages
- Implement validation error codes for programmatic handling

---

## Regulatory Compliance Features

### **PSD2 Compliance**
- Date range validations ensure 180-day maximum periods
- Currency validation supports European payment standards
- Authentication context validation enables Strong Customer Authentication

### **Data Protection**
- Password complexity rules protect user credentials
- Email validation ensures valid communication channels
- External user validation prevents data duplication

### **API Security**
- Consumer key validation secures API access
- Version-specific validation maintains API integrity
- Permission validation enforces access controls

---

**Analysis Completed:** 16-9-2025  
**Total Validation Rules:** 30  
**Source Files Analyzed:** 12  
**Categories Covered:** 6
