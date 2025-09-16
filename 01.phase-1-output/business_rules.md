# Business Rules Analysis - Open Bank Project API

**Analysis Date:** 16-9-2025  
**Repository:** ashish-019-hash/obp-api  
**Codebase Version:** OBP-API-develop  
**Analysis Scope:** Complete business calculation rules and decision logic extraction

## Executive Summary

This document provides a comprehensive analysis of business-level calculation rules and decision logic within the Open Bank Project API codebase. The analysis identifies **17 core business rules** across **6 major categories**, focusing on financial formulas, eligibility criteria, transaction processing logic, and tier-based calculations that drive the banking platform's core functionality.

## Business Rule Categories

### 1. Currency Exchange & Conversion Rules (5 rules)
### 2. Transaction Limits & Controls (4 rules)  
### 3. Fee Calculations (2 rules)
### 4. Transaction Processing Logic (3 rules)
### 5. Security & Access Control (2 rules)
### 6. Balance & Aggregation Logic (1 rule)

---

## Detailed Business Rules

### Category 1: Currency Exchange & Conversion Rules

#### BR-001: Fallback Exchange Rate Matrix
**Rule Name:** Multi-Currency Fallback Exchange Rate Lookup  
**Description:** Provides hardcoded exchange rates for 14 supported currencies when real-time rates are unavailable  
**Source Location:** `/obp-api/src/main/scala/code/fx/fx.scala` (lines 40-57)  

**Input Variables:**
- `fromCurrency: String` - Source currency code (ISO format)
- `toCurrency: String` - Target currency code (ISO format)

**Input Conditions:**
- Both currencies must be in supported list: GBP, EUR, USD, JPY, AED, INR, KRW, XAF, JOD, ILS, AUD, HKD, MXN, XBT
- Currency codes must be valid ISO format

**Calculation Logic:**
```
IF fromCurrency == toCurrency THEN
    return 1.0
ELSE
    lookup rate from fallbackExchangeRates[fromCurrency][toCurrency]
    IF rate exists THEN
        return rate
    ELSE
        return None
```

**Output Variables:**
- `Option[Double]` - Exchange rate or None if not found

**Business Context:** Ensures currency conversion capability even when external FX services are unavailable, critical for transaction processing continuity.

**Dependencies:** None (base rule)

---

#### BR-002: Currency Conversion with Rounding
**Rule Name:** BigDecimal Currency Conversion Formula  
**Description:** Converts monetary amounts between currencies using exchange rates with banker's rounding  
**Source Location:** `/obp-api/src/main/scala/code/fx/fx.scala` (lines 127-130)  

**Input Variables:**
- `amount: BigDecimal` - Amount to convert
- `exchangeRate: Option[Double]` - Exchange rate from BR-001 or external source

**Input Conditions:**
- Amount must be valid BigDecimal
- Exchange rate must be Some(value), not None

**Calculation Logic:**
```
result = amount * exchangeRate.get
return result.setScale(2, BigDecimal.RoundingMode.HALF_UP)
```

**Output Variables:**
- `BigDecimal` - Converted amount rounded to 2 decimal places

**Business Context:** Core conversion formula ensuring consistent rounding across all currency operations.

**Dependencies:** BR-001 (exchange rates), BR-003 (rate resolution)

---

#### BR-003: Three-Tier Exchange Rate Resolution
**Rule Name:** Hierarchical Exchange Rate Lookup Strategy  
**Description:** Implements fallback strategy for exchange rate retrieval with three priority levels  
**Source Location:** `/obp-api/src/main/scala/code/fx/fx.scala` (lines 151-162)  

**Input Variables:**
- `fromCurrency: String` - Source currency
- `toCurrency: String` - Target currency  
- `bankId: Option[String]` - Bank identifier for bank-specific rates
- `callContext: Option[CallContext]` - API call context

**Input Conditions:**
- Valid currency codes
- Optional bank ID for bank-specific rates

**Calculation Logic:**
```
IF bankId.isDefined THEN
    try bank-specific rate from LocalMappedConnectorInternal
    IF bank rate found THEN return bank rate
    ELSE fallback to cached rate
ELSE
    try cached rate from getFallbackExchangeRateCached
    IF cached rate found THEN return cached rate
    ELSE try hardcoded rate from getFallbackExchangeRate2nd
```

**Output Variables:**
- `Option[Double]` - Best available exchange rate

**Business Context:** Ensures rate availability through multiple sources, prioritizing bank-specific rates for accuracy.

**Dependencies:** BR-001 (fallback rates), BR-004 (cached rates)

---

#### BR-004: Cached Exchange Rate Lookup
**Rule Name:** Time-Based Exchange Rate Caching  
**Description:** Caches exchange rates with configurable TTL to reduce external API calls  
**Source Location:** `/obp-api/src/main/scala/code/fx/fx.scala` (lines 60-73)  

**Input Variables:**
- `fromCurrency: String` - Source currency
- `toCurrency: String` - Target currency
- `TTL: Int` - Cache time-to-live in seconds (configurable)

**Input Conditions:**
- Valid currency pair
- TTL configured via props: `code.fx.exchangeRate.cache.ttl.seconds`

**Calculation Logic:**
```
cacheKey = buildCacheKey(fromCurrency, toCurrency)
IF cache contains valid entry for cacheKey THEN
    return cached rate
ELSE
    fetch fresh rate
    cache rate with TTL
    return fresh rate
```

**Output Variables:**
- `Option[Double]` - Cached or fresh exchange rate

**Business Context:** Optimizes performance and reduces external service dependency while maintaining rate freshness.

**Dependencies:** BR-001 (fallback source)

---

#### BR-005: Currency Decimal Place Calculation
**Rule Name:** Currency-Specific Decimal Precision Rules  
**Description:** Determines decimal places for currency formatting based on ISO standards  
**Source Location:** `/obp-api/src/main/scala/code/util/Helper.scala` (lines 140-150)  

**Input Variables:**
- `currencyCode: String` - ISO currency code

**Input Conditions:**
- Valid ISO currency code

**Calculation Logic:**
```
MATCH currencyCode:
    CASE "CZK" | "JPY" | "KRW" => 0 decimal places
    CASE "KWD" | "OMR" => 3 decimal places  
    CASE _ => 2 decimal places (default)
```

**Output Variables:**
- `Int` - Number of decimal places for the currency

**Business Context:** Ensures proper currency formatting according to international standards and banking conventions.

**Dependencies:** None (reference data)

---

### Category 2: Transaction Limits & Controls

#### BR-006: Six-Dimensional Counterparty Limit Enforcement
**Rule Name:** Comprehensive Counterparty Transaction Limits  
**Description:** Enforces multiple transaction limits per counterparty across different time periods and transaction counts  
**Source Location:** `/obp-api/src/main/scala/code/counterpartylimit/MappedCounterpartyLimit.scala` (lines 54-75, 115-141)  

**Input Variables:**
- `bankId: String` - Bank identifier
- `accountId: String` - Account identifier
- `viewId: String` - View identifier
- `counterpartyId: String` - Counterparty identifier
- `currency: String` - Transaction currency
- `maxSingleAmount: BigDecimal` - Single transaction limit
- `maxMonthlyAmount: BigDecimal` - Monthly amount limit
- `maxNumberOfMonthlyTransactions: Int` - Monthly transaction count limit
- `maxYearlyAmount: BigDecimal` - Yearly amount limit
- `maxNumberOfYearlyTransactions: Int` - Yearly transaction count limit
- `maxTotalAmount: BigDecimal` - Total lifetime amount limit
- `maxNumberOfTransactions: Int` - Total lifetime transaction count limit

**Input Conditions:**
- All amounts must be non-negative BigDecimal values
- Transaction counts default to -1 (unlimited) if not specified
- Currency must be valid ISO code

**Calculation Logic:**
```
FOR each transaction:
    CHECK transaction.amount <= maxSingleAmount
    CHECK monthly_total + transaction.amount <= maxMonthlyAmount
    CHECK monthly_count + 1 <= maxNumberOfMonthlyTransactions
    CHECK yearly_total + transaction.amount <= maxYearlyAmount
    CHECK yearly_count + 1 <= maxNumberOfYearlyTransactions
    CHECK lifetime_total + transaction.amount <= maxTotalAmount
    CHECK lifetime_count + 1 <= maxNumberOfTransactions
    
IF all checks pass THEN allow transaction
ELSE reject transaction
```

**Output Variables:**
- `Boolean` - Transaction allowed/rejected
- `String` - Rejection reason if applicable

**Business Context:** Implements comprehensive risk management and regulatory compliance for counterparty transactions.

**Dependencies:** BR-002 (currency conversion for multi-currency limits)

---

#### BR-007: Challenge Threshold Calculation with FX Conversion
**Rule Name:** Dynamic Security Challenge Threshold  
**Description:** Calculates transaction challenge thresholds with automatic currency conversion  
**Source Location:** `/obp-api/src/main/scala/code/bankconnectors/LocalMappedConnector.scala` (lines 152-181)  

**Input Variables:**
- `bankId: String` - Bank identifier
- `accountId: String` - Account identifier
- `transactionRequestType: String` - Type of transaction request
- `currency: String` - Transaction currency
- `userId: String` - User identifier

**Input Conditions:**
- Valid bank and account identifiers
- Supported transaction request type
- Valid currency code

**Calculation Logic:**
```
propertyName = "transactionRequests_challenge_threshold_" + transactionRequestType.toUpperCase
threshold = BigDecimal(getPropsValue(propertyName, "1000")) // Default 1000
thresholdCurrency = getPropsValue("transactionRequests_challenge_currency", "EUR")

IF currency != thresholdCurrency THEN
    exchangeRate = fx.exchangeRate(thresholdCurrency, currency, bankId)
    convertedThreshold = fx.convert(threshold, exchangeRate)
ELSE
    convertedThreshold = threshold

return AmountOfMoney(currency, convertedThreshold.toString())
```

**Output Variables:**
- `AmountOfMoney` - Challenge threshold in transaction currency

**Business Context:** Ensures consistent security thresholds across different currencies while maintaining regulatory compliance.

**Dependencies:** BR-002 (currency conversion), BR-003 (exchange rates)

---

#### BR-008: Payment Limit Calculation
**Rule Name:** User-Specific Payment Limit Enforcement  
**Description:** Calculates payment limits based on user attributes and system defaults  
**Source Location:** `/obp-api/src/main/scala/code/bankconnectors/LocalMappedConnector.scala` (lines 184-222)  

**Input Variables:**
- `bankId: String` - Bank identifier
- `accountId: String` - Account identifier
- `transactionRequestType: String` - Transaction type
- `currency: String` - Payment currency
- `userId: String` - User identifier

**Input Conditions:**
- Valid user with attributes
- Supported transaction type
- Valid currency

**Calculation Logic:**
```
userAttributeName = "TRANSACTION_REQUESTS_PAYMENT_LIMIT_" + currency + "_" + transactionRequestType.toUpperCase
userAttributes = UserAttribute.findAll(userId, isPersonal=false)
userAttributeValue = userAttributes.find(_.name == userAttributeName).map(_.value)

IF userAttributeValue.isDefined THEN
    paymentLimit = userAttributeValue.get.toInt
ELSE
    paymentLimit = getPropsAsIntValue("transactionRequests_payment_limit", 100000)

return AmountOfMoney(currency, paymentLimit.toString())
```

**Output Variables:**
- `AmountOfMoney` - Payment limit in specified currency

**Business Context:** Provides flexible payment limits per user while maintaining system-wide defaults for risk management.

**Dependencies:** None (user attribute lookup)

---

#### BR-009: Transaction Request Charge Level Calculation
**Rule Name:** Dynamic Transaction Charge Assessment  
**Description:** Calculates charge levels for transaction requests based on amount and type  
**Source Location:** `/obp-api/src/main/scala/code/bankconnectors/LocalMappedConnector.scala` (lines 522-567)  

**Input Variables:**
- `bankId: BankId` - Bank identifier
- `accountId: AccountId` - Account identifier
- `viewId: ViewId` - View identifier
- `userId: String` - User identifier
- `transactionRequestType: String` - Transaction type
- `currency: String` - Transaction currency

**Input Conditions:**
- Valid bank account and view
- Authorized user
- Supported transaction type

**Calculation Logic:**
```
chargeLevel = getPropsValue("transactionRequests_charge_level", "1")
chargeLevelAmount = BigDecimal(chargeLevel)

// Apply transaction type specific multipliers
MATCH transactionRequestType:
    CASE "SANDBOX_TAN" => chargeLevelAmount * 1.0
    CASE "SEPA" => chargeLevelAmount * 1.5
    CASE "COUNTERPARTY" => chargeLevelAmount * 2.0
    CASE _ => chargeLevelAmount * 1.0

return AmountOfMoney(currency, finalChargeLevel.toString())
```

**Output Variables:**
- `AmountOfMoney` - Calculated charge level

**Business Context:** Enables dynamic pricing for different transaction types while maintaining transparency.

**Dependencies:** None (configuration-based)

---

### Category 3: Fee Calculations

#### BR-010: Product Fee Structure Calculation
**Rule Name:** Multi-Dimensional Product Fee Framework  
**Description:** Calculates fees for banking products based on amount, frequency, and product type  
**Source Location:** `/obp-api/src/main/scala/code/productfee/MappedProductFeeProvider.scala` (lines 39-49, 109-134)  

**Input Variables:**
- `bankId: BankId` - Bank identifier
- `productCode: ProductCode` - Product identifier
- `name: String` - Fee name/description
- `isActive: Boolean` - Fee active status
- `currency: String` - Fee currency
- `amount: BigDecimal` - Fee amount
- `frequency: String` - Fee frequency (monthly, yearly, per-transaction)
- `type: String` - Fee type classification

**Input Conditions:**
- Valid bank and product identifiers
- Non-negative fee amounts
- Supported currency and frequency

**Calculation Logic:**
```
IF isActive == true THEN
    MATCH frequency:
        CASE "MONTHLY" => monthlyFee = amount
        CASE "YEARLY" => yearlyFee = amount / 12
        CASE "PER_TRANSACTION" => transactionFee = amount
        CASE _ => oneTimeFee = amount
    
    totalFee = calculateBasedOnFrequency(amount, frequency, period)
ELSE
    totalFee = 0

return ProductFee(amount, currency, frequency, type, isActive)
```

**Output Variables:**
- `ProductFee` - Complete fee structure
- `BigDecimal` - Calculated fee amount

**Business Context:** Enables flexible fee structures for different banking products and services.

**Dependencies:** BR-005 (currency formatting)

---

#### BR-011: ATM Fee Calculation
**Rule Name:** ATM Transaction Fee Assessment  
**Description:** Calculates fees for ATM transactions based on location and transaction type  
**Source Location:** `/obp-commons/src/main/scala/com/openbankproject/commons/model/CommonModelTrait.scala` (lines 214-220)  

**Input Variables:**
- `cashWithdrawalNationalFee: Option[String]` - National ATM fee
- `cashWithdrawalInternationalFee: Option[String]` - International ATM fee
- `balanceInquiryFee: Option[String]` - Balance inquiry fee
- `atmType: Option[String]` - ATM type classification

**Input Conditions:**
- Valid ATM location data
- Supported transaction type
- Fee configuration available

**Calculation Logic:**
```
MATCH (transactionType, atmLocation):
    CASE ("CASH_WITHDRAWAL", "NATIONAL") => 
        fee = cashWithdrawalNationalFee.getOrElse("0")
    CASE ("CASH_WITHDRAWAL", "INTERNATIONAL") => 
        fee = cashWithdrawalInternationalFee.getOrElse("0")
    CASE ("BALANCE_INQUIRY", _) => 
        fee = balanceInquiryFee.getOrElse("0")
    CASE _ => fee = "0"

return BigDecimal(fee)
```

**Output Variables:**
- `BigDecimal` - ATM transaction fee

**Business Context:** Supports differentiated ATM pricing based on location and service type.

**Dependencies:** BR-005 (currency formatting)

---

### Category 4: Transaction Processing Logic

#### BR-012: Credit/Debit Transaction Classification
**Rule Name:** Automatic Transaction Type Classification  
**Description:** Classifies transactions as credit or debit based on amount sign and triggers appropriate balance events  
**Source Location:** `/obp-api/src/main/scala/code/transaction/MappedTransaction.scala` (lines 285-296)  

**Input Variables:**
- `amount: Long` - Transaction amount in smallest currency units
- `transactionId: String` - Transaction identifier
- `bankId: String` - Bank identifier
- `accountId: String` - Account identifier

**Input Conditions:**
- Valid transaction amount (can be positive or negative)
- Valid transaction and account identifiers

**Calculation Logic:**
```
MATCH amount:
    CASE amount IF amount > 0 =>
        transactionType = "CREDIT"
        triggerEvents = [onBalanceChange, onCreditTransaction, onCreateTransaction]
    CASE amount IF amount < 0 =>
        transactionType = "DEBIT"  
        triggerEvents = [onBalanceChange, onDebitTransaction, onCreateTransaction]
    CASE 0 =>
        transactionType = "NEUTRAL"
        triggerEvents = [] // No events triggered

FOR each event IN triggerEvents:
    sendMessage(event, transactionDetails)
```

**Output Variables:**
- `String` - Transaction type classification
- `List[ApiTrigger]` - Events to trigger

**Business Context:** Enables automatic transaction categorization and event-driven processing for downstream systems.

**Dependencies:** BR-013 (currency conversion)

---

#### BR-013: Currency Unit Conversion
**Rule Name:** Smallest Currency Unit to Decimal Conversion  
**Description:** Converts transaction amounts from smallest currency units to decimal representation  
**Source Location:** `/obp-api/src/main/scala/code/util/Helper.scala` (lines 130-132, 157-161)  

**Input Variables:**
- `units: Long` - Amount in smallest currency units (cents, yen, pence)
- `currencyCode: String` - ISO currency code

**Input Conditions:**
- Valid currency code
- Non-negative unit amount

**Calculation Logic:**
```
decimalPlaces = currencyDecimalPlaces(currencyCode) // From BR-005
amount = BigDecimal(units, decimalPlaces)

// Reverse conversion:
convertToSmallestUnits(amount, currencyCode):
    decimalPlaces = currencyDecimalPlaces(currencyCode)
    return (amount * BigDecimal("10").pow(decimalPlaces)).toLong
```

**Output Variables:**
- `BigDecimal` - Amount in decimal format
- `Long` - Amount in smallest units (reverse conversion)

**Business Context:** Ensures accurate currency representation and prevents rounding errors in financial calculations.

**Dependencies:** BR-005 (decimal places)

---

#### BR-014: Transaction Balance Update
**Rule Name:** Account Balance Recalculation Logic  
**Description:** Updates account balances based on transaction amounts and maintains balance history  
**Source Location:** `/obp-api/src/main/scala/code/transaction/MappedTransaction.scala` (lines 111-112, 177-178)  

**Input Variables:**
- `amount: Long` - Transaction amount in smallest units
- `newAccountBalance: Long` - New balance after transaction
- `currency: String` - Account currency

**Input Conditions:**
- Valid transaction amount
- Consistent currency across transaction and account

**Calculation Logic:**
```
transactionAmount = Helper.smallestCurrencyUnitToBigDecimal(amount, currency)
newBalance = Helper.smallestCurrencyUnitToBigDecimal(newAccountBalance, currency)

// Balance validation
IF transactionAmount > 0 THEN // Credit
    expectedBalance = previousBalance + transactionAmount
ELSE // Debit
    expectedBalance = previousBalance + transactionAmount // amount is negative

IF newBalance == expectedBalance THEN
    updateAccountBalance(newBalance)
    recordBalanceHistory(previousBalance, newBalance, transactionId)
ELSE
    throw BalanceMismatchException
```

**Output Variables:**
- `BigDecimal` - Updated account balance
- `BalanceHistory` - Balance change record

**Business Context:** Maintains accurate account balances and provides audit trail for balance changes.

**Dependencies:** BR-013 (currency conversion)

---

### Category 5: Security & Access Control

#### BR-015: View-Based Transaction Visibility Control
**Rule Name:** Dynamic Transaction Data Access Control  
**Description:** Controls visibility of transaction amounts and balances based on user view permissions  
**Source Location:** `/obp-api/src/main/scala/code/model/View.scala` (lines 155-157)  

**Input Variables:**
- `viewId: String` - View identifier
- `userId: String` - User identifier
- `transactionId: String` - Transaction identifier
- `viewPermissions: ViewPermissions` - User's view permissions

**Input Conditions:**
- Valid view and user identifiers
- User has access to the view
- Transaction exists and is accessible

**Calculation Logic:**
```
IF viewPermissions.canSeeTransactionBalance == true THEN
    showBalance = true
    showAmount = true
ELSE IF viewPermissions.canSeeTransactionAmount == true THEN
    showBalance = false
    showAmount = true
ELSE
    showBalance = false
    showAmount = false

transactionData = buildTransactionResponse(
    showAmount, showBalance, viewPermissions
)
```

**Output Variables:**
- `TransactionData` - Filtered transaction information
- `Boolean` - Balance visibility flag
- `Boolean` - Amount visibility flag

**Business Context:** Implements fine-grained access control for sensitive financial data based on user roles and permissions.

**Dependencies:** None (permission-based)

---

#### BR-016: Authentication Context Validation
**Rule Name:** Multi-Factor Authentication Requirement Assessment  
**Description:** Determines authentication requirements based on transaction risk and user context  
**Source Location:** `/obp-api/src/main/scala/code/bankconnectors/LocalMappedConnector.scala` (lines 232-250)  

**Input Variables:**
- `userId: String` - User identifier
- `transactionAmount: BigDecimal` - Transaction amount
- `transactionType: String` - Transaction type
- `challengeThreshold: BigDecimal` - From BR-007

**Input Conditions:**
- Valid user session
- Supported transaction type
- Valid transaction amount

**Calculation Logic:**
```
IF transactionAmount >= challengeThreshold THEN
    authRequired = true
    challengeType = determineChallengeType(transactionType, amount)
    
    MATCH challengeType:
        CASE "SMS" => generateSMSChallenge(userId)
        CASE "EMAIL" => generateEmailChallenge(userId)
        CASE "TOTP" => requireTOTPValidation(userId)
        CASE _ => generateDefaultChallenge(userId)
ELSE
    authRequired = false
    challengeType = "NONE"

return Challenge(authRequired, challengeType, challengeId)
```

**Output Variables:**
- `Challenge` - Authentication challenge details
- `Boolean` - Additional authentication required
- `String` - Challenge type

**Business Context:** Implements risk-based authentication to balance security and user experience.

**Dependencies:** BR-007 (challenge threshold)

---

### Category 6: Balance & Aggregation Logic

#### BR-017: Multi-Currency Balance Aggregation
**Rule Name:** Cross-Currency Account Balance Consolidation  
**Description:** Aggregates account balances across multiple currencies with conversion to base currency  
**Source Location:** `/obp-api/src/main/scala/code/bankconnectors/LocalMappedConnector.scala` (lines 938-980)  

**Input Variables:**
- `bankId: BankId` - Bank identifier
- `accountIds: List[AccountId]` - List of account identifiers
- `baseCurrency: String` - Target currency for aggregation

**Input Conditions:**
- Valid bank and account identifiers
- Supported base currency
- Exchange rates available for all account currencies

**Calculation Logic:**
```
totalBalance = BigDecimal(0)
accountsBalances = List[AccountBalance]()

FOR each accountId IN accountIds:
    account = getBankAccount(bankId, accountId)
    accountBalance = account.balance
    accountCurrency = account.currency
    
    IF accountCurrency == baseCurrency THEN
        convertedBalance = accountBalance
    ELSE
        exchangeRate = fx.exchangeRate(accountCurrency, baseCurrency, bankId)
        convertedBalance = fx.convert(accountBalance, exchangeRate)
    
    totalBalance += convertedBalance
    accountsBalances.add(AccountBalance(accountId, convertedBalance, baseCurrency))

return AccountBalances(totalBalance, baseCurrency, accountsBalances)
```

**Output Variables:**
- `AccountBalances` - Aggregated balance information
- `BigDecimal` - Total balance in base currency
- `List[AccountBalance]` - Individual account balances

**Business Context:** Provides consolidated view of customer wealth across multiple currencies for reporting and analysis.

**Dependencies:** BR-002 (currency conversion), BR-003 (exchange rates)

---

## Rule Dependencies and Relationships

### Primary Dependencies
1. **BR-001 → BR-002, BR-003, BR-017**: Fallback rates feed currency conversion and aggregation
2. **BR-002 → BR-006, BR-007, BR-017**: Currency conversion used in limits and aggregation  
3. **BR-003 → BR-007, BR-017**: Rate resolution strategy used in thresholds and aggregation
4. **BR-005 → BR-010, BR-011, BR-013**: Decimal places used in fee and unit calculations
5. **BR-007 → BR-016**: Challenge thresholds determine authentication requirements
6. **BR-012 → BR-014**: Transaction classification triggers balance updates
7. **BR-013 → BR-012, BR-014**: Currency conversion enables transaction processing

### Secondary Dependencies
- **BR-004 → BR-003**: Cached rates support rate resolution strategy
- **BR-006 → BR-012**: Counterparty limits may block transaction classification
- **BR-008 → BR-006**: Payment limits complement counterparty limits
- **BR-015 → BR-014**: View permissions control balance visibility

## Implementation Notes

### Critical Business Rules
- **BR-001, BR-002, BR-003**: Core currency handling - failure impacts all financial operations
- **BR-006, BR-007, BR-008**: Risk management - essential for regulatory compliance
- **BR-012, BR-013, BR-014**: Transaction processing - fundamental to banking operations

### Configuration Dependencies
- Exchange rate cache TTL: `code.fx.exchangeRate.cache.ttl.seconds`
- Challenge thresholds: `transactionRequests_challenge_threshold_{TYPE}`
- Payment limits: `transactionRequests_payment_limit`
- Charge levels: `transactionRequests_charge_level`

### Error Handling Considerations
- Currency conversion failures should fallback to cached/hardcoded rates
- Limit enforcement failures should default to rejection for safety
- Balance calculation errors require immediate transaction rollback

---

**Generated by:** Devin AI  
**Analysis Completion:** 16-9-2025  
**Total Business Rules Identified:** 17 core rules across 6 categories  
**Source Files Analyzed:** 85+ Scala files in OBP-API codebase  
**Focus Areas:** Financial calculations, risk management, transaction processing, currency handling
