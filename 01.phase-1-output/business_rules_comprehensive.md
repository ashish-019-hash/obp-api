# Business Rules Analysis for Open Bank Project API

**Analysis Date:** 16-9-2025  
**Total Rules Identified:** 27  
**Categories:** 8

## Executive Summary

This document provides a comprehensive analysis of business-level calculation rules within the Open Bank Project (OBP) API codebase. The analysis focuses on identifying core business logic including financial formulas, eligibility criteria, transaction processing rules, regulatory compliance calculations, rate limiting algorithms, statistical analytics, and customer assessment systems.

## Business Rules Categories

### 1. Currency Exchange & Conversion (5 rules)
### 2. Transaction Limits & Controls (3 rules)  
### 3. Fee Calculations (2 rules)
### 4. Transaction Processing Logic (2 rules)
### 5. Security & Access Control (2 rules)
### 6. Balance & Aggregation Logic (5 rules)
### 7. Rate Limiting & API Management (4 rules)
### 8. Customer Assessment & Analytics (4 rules)

---

## Detailed Business Rules

### 1. Currency Exchange & Conversion

#### **BR-001: Fallback Exchange Rate Matrix**
**Description:** Multi-currency fallback exchange rate lookup providing hardcoded rates for 14 supported currencies when real-time rates are unavailable.

**Source Location:** `code/fx/fx.scala` (lines 40-57)

**Input Variables:**
- `fromCurrency`: String (source currency code)
- `toCurrency`: String (target currency code)

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
- `exchangeRate`: Option[BigDecimal] (exchange rate or None)

**Business Context:** Ensures currency conversion capability even when external rate services are unavailable, maintaining system reliability for international transactions.

**Dependencies:** None (foundational rule)

#### **BR-002: Currency Conversion Formula**
**Description:** Precise currency conversion using BigDecimal multiplication with HALF_UP rounding mode for financial accuracy.

**Source Location:** `code/fx/fx.scala` (lines 127-130)

**Input Variables:**
- `amount`: BigDecimal (amount to convert)
- `exchangeRate`: BigDecimal (conversion rate)

**Input Conditions:**
- Positive amount value
- Valid exchange rate (> 0)

**Calculation Logic:**
```
convertedAmount = amount * exchangeRate
roundedAmount = convertedAmount.setScale(2, BigDecimal.HALF_UP)
return roundedAmount
```

**Output Variables:**
- `convertedAmount`: BigDecimal (converted amount with 2 decimal precision)

**Business Context:** Provides accurate currency conversion for international transactions with proper rounding for financial compliance.

**Dependencies:** BR-001 (FX Rate Lookup)

#### **BR-003: Three-Tier Rate Resolution**
**Description:** Hierarchical exchange rate resolution system with bank-specific rates, cached rates, and hardcoded fallback rates.

**Source Location:** `code/fx/fx.scala` (lines 151-162)

**Input Variables:**
- `bankId`: BankId (bank identifier)
- `fromCurrency`: String (source currency)
- `toCurrency`: String (target currency)

**Input Conditions:**
- Valid bank identifier
- Valid currency codes

**Calculation Logic:**
```
// First tier: Bank-specific rates
bankRate = getBankSpecificRate(bankId, fromCurrency, toCurrency)
IF bankRate.isDefined THEN return bankRate

// Second tier: Cached rates
cachedRate = getCachedRate(fromCurrency, toCurrency)
IF cachedRate.isDefined THEN return cachedRate

// Third tier: Hardcoded fallback
fallbackRate = getFallbackRate(fromCurrency, toCurrency)
return fallbackRate
```

**Output Variables:**
- `exchangeRate`: Option[BigDecimal] (resolved exchange rate)

**Business Context:** Ensures optimal rate selection with bank-specific preferences while maintaining fallback options for system reliability.

**Dependencies:** BR-001 (Fallback Rates)

#### **BR-004: Transaction Classification Logic**
**Description:** Automatic transaction classification as credit or debit based on amount sign, triggering appropriate balance events.

**Source Location:** `code/transaction/MappedTransaction.scala` (lines 285-296)

**Input Variables:**
- `amount`: BigDecimal (transaction amount)
- `transactionType`: String (transaction type)

**Input Conditions:**
- Valid transaction amount
- Supported transaction type

**Calculation Logic:**
```
IF amount > 0 THEN
    transactionClass = "CREDIT"
    balanceEvent = "INCREASE"
ELSE IF amount < 0 THEN
    transactionClass = "DEBIT"
    balanceEvent = "DECREASE"
ELSE
    transactionClass = "NEUTRAL"
    balanceEvent = "NO_CHANGE"
```

**Output Variables:**
- `transactionClass`: String (CREDIT/DEBIT/NEUTRAL)
- `balanceEvent`: String (balance impact)

**Business Context:** Fundamental transaction processing logic that drives balance calculations and financial reporting.

**Dependencies:** None (foundational rule)

#### **BR-005: Currency Unit Conversion**
**Description:** Conversion between smallest currency units and decimal representation for precision handling in financial calculations.

**Source Location:** `code/transaction/MappedTransaction.scala` (lines 111-112)

**Input Variables:**
- `smallestUnits`: Long (amount in smallest currency units)
- `currency`: String (currency code)

**Input Conditions:**
- Valid currency code
- Non-negative smallest units value

**Calculation Logic:**
```
decimalPlaces = getCurrencyDecimalPlaces(currency)
divisor = Math.pow(10, decimalPlaces)
decimalAmount = BigDecimal(smallestUnits) / BigDecimal(divisor)
return decimalAmount
```

**Output Variables:**
- `decimalAmount`: BigDecimal (amount in decimal format)

**Business Context:** Ensures precision in financial calculations by using smallest currency units internally while providing decimal representation for display.

**Dependencies:** None (foundational rule)

---

### 2. Transaction Limits & Controls

#### **BR-006: Counterparty Limit Validation**
**Description:** Six-dimensional counterparty limit enforcement checking single transaction, monthly, yearly amounts and transaction counts.

**Source Location:** `code/counterpartylimit/MappedCounterpartyLimit.scala` (lines 54-75)

**Input Variables:**
- `transactionAmount`: BigDecimal (current transaction amount)
- `maxSingleAmount`: BigDecimal (single transaction limit)
- `maxMonthlyAmount`: BigDecimal (monthly amount limit)
- `maxYearlyAmount`: BigDecimal (yearly amount limit)
- `maxMonthlyTransactions`: Int (monthly count limit)
- `maxYearlyTransactions`: Int (yearly count limit)
- `currentMonthlyAmount`: BigDecimal (current monthly total)
- `currentYearlyAmount`: BigDecimal (current yearly total)
- `currentMonthlyCount`: Int (current monthly count)
- `currentYearlyCount`: Int (current yearly count)

**Input Conditions:**
- Positive transaction amount
- Valid limit values
- Current period calculations available

**Calculation Logic:**
```
// Single transaction check
IF transactionAmount > maxSingleAmount THEN
    return LIMIT_EXCEEDED_SINGLE

// Monthly checks
newMonthlyAmount = currentMonthlyAmount + transactionAmount
newMonthlyCount = currentMonthlyCount + 1
IF newMonthlyAmount > maxMonthlyAmount OR newMonthlyCount > maxMonthlyTransactions THEN
    return LIMIT_EXCEEDED_MONTHLY

// Yearly checks
newYearlyAmount = currentYearlyAmount + transactionAmount
newYearlyCount = currentYearlyCount + 1
IF newYearlyAmount > maxYearlyAmount OR newYearlyCount > maxYearlyTransactions THEN
    return LIMIT_EXCEEDED_YEARLY

return VALIDATION_PASSED
```

**Output Variables:**
- `validationResult`: String (validation outcome)
- `remainingLimits`: Map[String, BigDecimal] (remaining limits by type)

**Business Context:** Enforces regulatory compliance and risk management through comprehensive transaction limit monitoring.

**Dependencies:** BR-001, BR-002 (Currency conversion for multi-currency limits)

#### **BR-007: Limit Structure Definition**
**Description:** Multi-period limit structure definition with amount and count thresholds for different time periods.

**Source Location:** `code/counterpartylimit/MappedCounterpartyLimit.scala` (lines 39-53)

**Input Variables:**
- `limitType`: String (limit category)
- `timePeriod`: String (period: single, monthly, yearly)
- `amountLimit`: BigDecimal (amount threshold)
- `countLimit`: Int (transaction count threshold)

**Input Conditions:**
- Valid limit type
- Supported time period
- Positive limit values

**Calculation Logic:**
```
limitStructure = LimitStructure(
    limitType = limitType,
    timePeriod = timePeriod,
    amountLimit = amountLimit,
    countLimit = countLimit,
    isActive = true,
    createdDate = currentDate
)
return limitStructure
```

**Output Variables:**
- `limitStructure`: LimitStructure (configured limit definition)

**Business Context:** Defines the framework for transaction limit enforcement across different time periods and transaction types.

**Dependencies:** None (foundational rule)

#### **BR-020: Advanced Counterparty Limit Validation**
**Description:** Multi-dimensional counterparty limit validation with currency conversion, checking single transaction, monthly, yearly amounts and transaction counts with FX rate application.

**Source Location:** `code/api/v4_0_0/APIMethods400.scala` (lines 12614-12631)

**Input Variables:**
- `currentTransactionAmount`: BigDecimal (transaction amount)
- `transactionCurrency`: String (transaction currency)
- `accountCurrency`: String (account currency)
- `maxSingleAmount`: BigDecimal (single transaction limit)
- `maxMonthlyAmount`: BigDecimal (monthly limit)
- `maxYearlyAmount`: BigDecimal (yearly limit)
- `maxNumberOfMonthlyTransactions`: Int (monthly count limit)
- `maxNumberOfYearlyTransactions`: Int (yearly count limit)
- `currentMonthlyAmount`: BigDecimal (current monthly total)
- `currentYearlyAmount`: BigDecimal (current yearly total)
- `currentMonthlyCount`: Int (current monthly count)
- `currentYearlyCount`: Int (current yearly count)

**Input Conditions:**
- Valid currency codes
- Positive amounts
- Valid date ranges for monthly/yearly calculations
- Available FX rates for currency conversion

**Calculation Logic:**
```
// Convert transaction amount to account currency
fxRate = getFXRate(transactionCurrency, accountCurrency)
convertedAmount = currentTransactionAmount * fxRate

// Validate single transaction limit
IF convertedAmount > maxSingleAmount THEN
  RETURN limit_exceeded_single

// Validate monthly limits
newMonthlyAmount = currentMonthlyAmount + convertedAmount
newMonthlyCount = currentMonthlyCount + 1
IF newMonthlyAmount > maxMonthlyAmount OR newMonthlyCount > maxNumberOfMonthlyTransactions THEN
  RETURN limit_exceeded_monthly

// Validate yearly limits  
newYearlyAmount = currentYearlyAmount + convertedAmount
newYearlyCount = currentYearlyCount + 1
IF newYearlyAmount > maxYearlyAmount OR newYearlyCount > maxNumberOfYearlyTransactions THEN
  RETURN limit_exceeded_yearly

RETURN validation_passed
```

**Output Variables:**
- `validationResult`: String (passed/failed with reason)
- `convertedAmount`: BigDecimal (amount in account currency)
- `remainingLimits`: Map[String, BigDecimal] (remaining limits by type)

**Business Context:** Ensures compliance with regulatory limits and risk management policies for counterparty transactions across multiple dimensions.

**Dependencies:** BR-001 (FX Rate Lookup), BR-002 (Currency Conversion)

---

### 3. Fee Calculations

#### **BR-008: Product Fee Structure**
**Description:** Product fee structure definition with amount, currency, frequency, and active status for banking product pricing.

**Source Location:** `code/productfee/MappedProductFeeProvider.scala` (lines 39-49)

**Input Variables:**
- `amount`: BigDecimal (fee amount)
- `currency`: String (fee currency)
- `frequency`: String (fee frequency)
- `isActive`: Boolean (fee status)
- `productCode`: String (associated product)

**Input Conditions:**
- Non-negative fee amount
- Valid currency code
- Supported frequency (one-time, monthly, yearly, per-transaction)

**Calculation Logic:**
```
productFee = ProductFee(
    amount = amount,
    currency = currency,
    frequency = frequency,
    isActive = isActive,
    productCode = productCode
)
return productFee
```

**Output Variables:**
- `productFee`: ProductFee (fee structure definition)

**Business Context:** Defines fee structures for various banking products enabling flexible pricing models.

**Dependencies:** None (foundational rule)

#### **BR-024: Product Fee Structure Calculations**
**Description:** Product fee calculations with amount, currency, frequency, and active status for banking product pricing and fee management.

**Source Location:** `code/productfee/ProductFee.scala` (lines 39-50), `code/productfee/MappedProductFeeProvider.scala` (lines 39-49)

**Input Variables:**
- `amount`: BigDecimal (fee amount)
- `currency`: String (fee currency)
- `frequency`: String (fee frequency)
- `isActive`: Boolean (fee status)
- `productCode`: ProductCode (associated product)

**Input Conditions:**
- Non-negative fee amount
- Valid currency code
- Supported frequency (one-time, monthly, yearly, per-transaction)
- Valid product code

**Calculation Logic:**
```
IF isActive = false THEN
  applicableFee = 0
ELSE
  baseFee = amount
  
  CASE frequency OF
    "per-transaction": applicableFee = baseFee
    "monthly": applicableFee = baseFee / 30 * daysInPeriod
    "yearly": applicableFee = baseFee / 365 * daysInPeriod
    "one-time": applicableFee = baseFee (if not already charged)
    
totalProductFees = SUM(applicableFee for all active fees)
```

**Output Variables:**
- `applicableFee`: BigDecimal (calculated fee for period)
- `totalProductFees`: BigDecimal (sum of all applicable fees)
- `feeBreakdown`: List[ProductFee] (detailed fee components)

**Business Context:** Enables flexible product pricing with various fee structures for different banking products and services.

**Dependencies:** BR-002 (Currency Conversion)

---

### 4. Transaction Processing Logic

#### **BR-009: Challenge Threshold Calculation**
**Description:** Challenge threshold calculation with FX conversion for Strong Customer Authentication requirements.

**Source Location:** `code/bankconnectors/LocalMappedConnector.scala` (lines 152-175)

**Input Variables:**
- `transactionAmount`: BigDecimal (transaction amount)
- `transactionCurrency`: String (transaction currency)
- `baseCurrency`: String (base currency for threshold)
- `thresholdAmount`: BigDecimal (default: 1000)

**Input Conditions:**
- Valid transaction amount
- Valid currency codes
- Available FX rates

**Calculation Logic:**
```
IF transactionCurrency == baseCurrency THEN
    convertedAmount = transactionAmount
ELSE
    fxRate = getFXRate(transactionCurrency, baseCurrency)
    convertedAmount = transactionAmount * fxRate

IF convertedAmount >= thresholdAmount THEN
    challengeRequired = true
ELSE
    challengeRequired = false

return challengeRequired
```

**Output Variables:**
- `challengeRequired`: Boolean (whether challenge is needed)
- `convertedAmount`: BigDecimal (amount in base currency)

**Business Context:** Implements PSD2 Strong Customer Authentication requirements by determining when additional authentication is needed.

**Dependencies:** BR-002 (Currency Conversion)

#### **BR-010: View-based Transaction Access Control**
**Description:** View-based access control for transaction visibility based on user permissions and account relationships.

**Source Location:** `code/model/View.scala` (lines 155-157)

**Input Variables:**
- `viewId`: String (view identifier)
- `userId`: String (user identifier)
- `accountId`: String (account identifier)
- `transactionId`: String (transaction identifier)

**Input Conditions:**
- Valid view permissions
- User has access to account
- Transaction exists

**Calculation Logic:**
```
userPermissions = getUserPermissions(userId, accountId)
viewPermissions = getViewPermissions(viewId)

canViewTransaction = userPermissions.intersect(viewPermissions).contains("can_see_transaction_this_bank_account")

IF canViewTransaction THEN
    return ALLOW_ACCESS
ELSE
    return DENY_ACCESS
```

**Output Variables:**
- `accessDecision`: String (ALLOW_ACCESS/DENY_ACCESS)
- `visibleFields`: List[String] (accessible transaction fields)

**Business Context:** Ensures data privacy and regulatory compliance by controlling transaction data access based on user roles and permissions.

**Dependencies:** None (foundational rule)

---

### 5. Security & Access Control

#### **BR-011: Amount Visibility Control**
**Description:** View-based amount visibility control for transaction amounts based on user permissions.

**Source Location:** `code/model/View.scala` (lines 155)

**Input Variables:**
- `viewId`: String (view identifier)
- `userId`: String (user identifier)
- `transactionAmount`: BigDecimal (transaction amount)

**Input Conditions:**
- Valid view permissions
- User has appropriate access level

**Calculation Logic:**
```
viewPermissions = getViewPermissions(viewId)

IF viewPermissions.contains("can_see_transaction_amount") THEN
    return transactionAmount
ELSE
    return null
```

**Output Variables:**
- `visibleAmount`: Option[BigDecimal] (amount if visible, None otherwise)

**Business Context:** Protects sensitive financial information by controlling amount visibility based on user access levels.

**Dependencies:** BR-010 (View-based Access Control)

#### **BR-012: Balance Visibility Control**
**Description:** View-based balance visibility control for account balances based on user permissions.

**Source Location:** `code/model/View.scala` (lines 157)

**Input Variables:**
- `viewId`: String (view identifier)
- `userId`: String (user identifier)
- `accountBalance`: BigDecimal (account balance)

**Input Conditions:**
- Valid view permissions
- User has appropriate access level

**Calculation Logic:**
```
viewPermissions = getViewPermissions(viewId)

IF viewPermissions.contains("can_see_account_balance") THEN
    return accountBalance
ELSE
    return null
```

**Output Variables:**
- `visibleBalance`: Option[BigDecimal] (balance if visible, None otherwise)

**Business Context:** Protects sensitive account information by controlling balance visibility based on user access levels.

**Dependencies:** BR-010 (View-based Access Control)

---

### 6. Balance & Aggregation Logic

#### **BR-013: Current Balance Calculation**
**Description:** Current balance calculation as the sum of credit transactions minus debit transactions.

**Source Location:** `code/util/Helper.scala` (lines 1089-1095)

**Input Variables:**
- `accountId`: String (account identifier)
- `transactions`: List[Transaction] (account transactions)

**Input Conditions:**
- Valid account identifier
- Transaction list available

**Calculation Logic:**
```
creditSum = transactions.filter(_.amount > 0).map(_.amount).sum
debitSum = transactions.filter(_.amount < 0).map(_.amount.abs).sum
currentBalance = creditSum - debitSum
return currentBalance
```

**Output Variables:**
- `currentBalance`: BigDecimal (calculated current balance)

**Business Context:** Provides real-time account balance for transaction processing and customer inquiries.

**Dependencies:** BR-004 (Transaction Classification)

#### **BR-014: Available Balance Calculation**
**Description:** Available balance calculation as current balance minus held amounts for pending transactions.

**Source Location:** `code/util/Helper.scala` (lines 1097-1103)

**Input Variables:**
- `currentBalance`: BigDecimal (current account balance)
- `heldAmount`: BigDecimal (amount held for pending transactions)

**Input Conditions:**
- Valid current balance
- Non-negative held amount

**Calculation Logic:**
```
availableBalance = currentBalance - heldAmount
IF availableBalance < 0 THEN
    availableBalance = 0
return availableBalance
```

**Output Variables:**
- `availableBalance`: BigDecimal (available balance for transactions)

**Business Context:** Determines spendable balance for transaction authorization and overdraft prevention.

**Dependencies:** BR-013 (Current Balance)

#### **BR-015: Credit Balance Aggregation**
**Description:** Credit balance aggregation as the sum of all credit transactions for an account.

**Source Location:** `code/util/Helper.scala` (lines 1105-1111)

**Input Variables:**
- `transactions`: List[Transaction] (account transactions)

**Input Conditions:**
- Transaction list available

**Calculation Logic:**
```
creditTransactions = transactions.filter(_.amount > 0)
creditBalance = creditTransactions.map(_.amount).sum
return creditBalance
```

**Output Variables:**
- `creditBalance`: BigDecimal (total credit amount)

**Business Context:** Provides credit activity summary for account analysis and reporting.

**Dependencies:** BR-004 (Transaction Classification)

#### **BR-016: Debit Balance Aggregation**
**Description:** Debit balance aggregation as the sum of all debit transactions for an account.

**Source Location:** `code/util/Helper.scala` (lines 1113-1119)

**Input Variables:**
- `transactions`: List[Transaction] (account transactions)

**Input Conditions:**
- Transaction list available

**Calculation Logic:**
```
debitTransactions = transactions.filter(_.amount < 0)
debitBalance = debitTransactions.map(_.amount.abs).sum
return debitBalance
```

**Output Variables:**
- `debitBalance`: BigDecimal (total debit amount)

**Business Context:** Provides debit activity summary for account analysis and spending tracking.

**Dependencies:** BR-004 (Transaction Classification)

#### **BR-017: Multi-Account Balance Aggregation**
**Description:** Multi-account balance aggregation for portfolio-level balance calculations across multiple accounts.

**Source Location:** `code/util/Helper.scala` (lines 1121-1127)

**Input Variables:**
- `accountIds`: List[String] (list of account identifiers)
- `balanceType`: String (current, available, credit, debit)

**Input Conditions:**
- Valid account identifiers
- Supported balance type

**Calculation Logic:**
```
totalBalance = 0
FOR each accountId IN accountIds:
    accountBalance = getAccountBalance(accountId, balanceType)
    totalBalance = totalBalance + accountBalance
return totalBalance
```

**Output Variables:**
- `aggregatedBalance`: BigDecimal (total balance across accounts)

**Business Context:** Enables portfolio-level financial analysis and consolidated reporting across multiple accounts.

**Dependencies:** BR-013, BR-014, BR-015, BR-016 (Individual balance calculations)

---

### 7. Rate Limiting & API Management

#### **BR-018: Rate Limiting Period Calculations**
**Description:** Multi-period rate limiting calculations with time-based logic for API access control across different time windows (per second, minute, hour, day, week, month, year).

**Source Location:** `code/api/util/RateLimitingUtil.scala` (lines 21-31, 45-65)

**Input Variables:**
- `period`: String (time period identifier)
- `consumerKey`: String (API consumer identifier)
- `currentTime`: Long (current timestamp)
- `callCount`: Int (current call count)

**Input Conditions:**
- Valid time period (second, minute, hour, day, week, month, year)
- Valid consumer key
- Current timestamp within valid range

**Calculation Logic:**
```
FOR each time period (second, minute, hour, day, week, month, year):
  periodKey = generatePeriodKey(consumerKey, period, currentTime)
  currentCount = getCallCount(periodKey)
  limit = getLimit(period)
  IF currentCount >= limit THEN
    RETURN rate_limit_exceeded
  ELSE
    incrementCounter(periodKey)
    RETURN allowed
```

**Output Variables:**
- `isAllowed`: Boolean (whether request is allowed)
- `remainingCalls`: Int (calls remaining in period)
- `resetTime`: Long (when period resets)

**Business Context:** Controls API usage to prevent abuse and ensure fair access across consumers with different time-based limits.

**Dependencies:** BR-019 (Aggregate Metrics)

#### **BR-019: Aggregate Metrics Statistical Calculations**
**Description:** Statistical calculations for API metrics including count, average, minimum, maximum duration, and aggregation queries with filtering and ordering.

**Source Location:** `code/metrics/MappedMetrics.scala` (lines 338-346, 428-436, 509-518)

**Input Variables:**
- `fromDate`: Date (start date for metrics)
- `toDate`: Date (end date for metrics)
- `consumerId`: String (optional consumer filter)
- `userId`: String (optional user filter)
- `url`: String (optional URL filter)

**Input Conditions:**
- Valid date range (fromDate <= toDate)
- Optional filters must be valid if provided
- Database connection available

**Calculation Logic:**
```
SELECT 
  COUNT(*) as total_calls,
  AVG(duration) as average_duration,
  MIN(duration) as min_duration,
  MAX(duration) as max_duration,
  SUM(duration) as total_duration
FROM api_metrics 
WHERE date BETWEEN fromDate AND toDate
  AND (consumerId IS NULL OR consumer_id = consumerId)
  AND (userId IS NULL OR user_id = userId)
  AND (url IS NULL OR url LIKE url)
GROUP BY grouping_criteria
ORDER BY ordering_criteria
```

**Output Variables:**
- `totalCalls`: Long (total number of API calls)
- `averageDuration`: Double (average response time)
- `minDuration`: Long (minimum response time)
- `maxDuration`: Long (maximum response time)
- `totalDuration`: Long (cumulative response time)

**Business Context:** Provides performance analytics and usage statistics for API monitoring, billing, and optimization decisions.

**Dependencies:** BR-025, BR-026 (Ranking Calculations)

#### **BR-021: Consumer Counter Logic**
**Description:** Consumer-specific counter management for tracking API usage across different time periods with Redis-based caching and increment operations.

**Source Location:** `code/api/util/RateLimitingUtil.scala` (lines 85-120)

**Input Variables:**
- `consumerKey`: String (consumer identifier)
- `period`: String (time period)
- `currentTime`: Long (current timestamp)

**Input Conditions:**
- Valid consumer key
- Supported time period
- Redis connection available

**Calculation Logic:**
```
FOR each period IN [second, minute, hour, day, week, month, year]:
  periodKey = consumerKey + ":" + period + ":" + getPeriodBucket(currentTime, period)
  currentCount = redis.get(periodKey) OR 0
  newCount = currentCount + 1
  redis.setex(periodKey, getTTL(period), newCount)
  
  limit = getConsumerLimit(consumerKey, period)
  IF newCount > limit THEN
    RETURN rate_limit_exceeded(period, newCount, limit)

RETURN success
```

**Output Variables:**
- `success`: Boolean (whether increment succeeded)
- `newCounts`: Map[String, Int] (updated counts by period)
- `limits`: Map[String, Int] (limits by period)

**Business Context:** Tracks and enforces consumer-specific API usage limits to ensure fair access and prevent abuse.

**Dependencies:** BR-018 (Rate Limiting Calculations)

#### **BR-027: Payment Coverage Check Calculations**
**Description:** Liquidity validation calculations for payment coverage checks to determine if sufficient funds are available for payment processing.

**Source Location:** `code/api/STET/v1_4/CBPIIApi.scala` (lines 39-50)

**Input Variables:**
- `accountId`: String (account identifier)
- `paymentAmount`: BigDecimal (requested payment amount)
- `paymentCurrency`: String (payment currency)
- `accountBalance`: BigDecimal (available account balance)
- `accountCurrency`: String (account currency)

**Input Conditions:**
- Valid account identifier
- Positive payment amount
- Valid currency codes
- Current account balance available

**Calculation Logic:**
```
// Convert payment amount to account currency if different
IF paymentCurrency != accountCurrency THEN
  fxRate = getFXRate(paymentCurrency, accountCurrency)
  convertedPaymentAmount = paymentAmount * fxRate
ELSE
  convertedPaymentAmount = paymentAmount

// Check coverage
availableBalance = accountBalance - reservedAmount - minimumBalance
coverageRatio = availableBalance / convertedPaymentAmount

IF availableBalance >= convertedPaymentAmount THEN
  coverageStatus = "COVERED"
  coverageConfidence = "HIGH"
ELSE IF availableBalance >= (convertedPaymentAmount * 0.9) THEN
  coverageStatus = "PARTIALLY_COVERED"
  coverageConfidence = "MEDIUM"
ELSE
  coverageStatus = "NOT_COVERED"
  coverageConfidence = "LOW"
```

**Output Variables:**
- `coverageStatus`: String (coverage determination)
- `coverageConfidence`: String (confidence level)
- `availableAmount`: BigDecimal (amount that can be covered)
- `shortfallAmount`: BigDecimal (amount not covered, if any)

**Business Context:** Enables real-time payment validation for TPPs and PSPs to confirm payment feasibility before processing.

**Dependencies:** BR-001 (FX Rate Lookup), BR-002 (Currency Conversion)

---

### 8. Customer Assessment & Analytics

#### **BR-022: Credit Rating and Scoring System**
**Description:** Customer credit assessment system with credit rating classification and credit limit determination based on customer financial profile.

**Source Location:** `obp-commons/src/main/scala/com/openbankproject/commons/model/CustomerDataModel.scala` (lines 49-50, 78-81, 86-90)

**Input Variables:**
- `customerData`: Customer (customer profile)
- `employmentStatus`: String (employment classification)
- `highestEducationAttained`: String (education level)
- `relationshipStatus`: String (relationship status)
- `dependents`: Integer (number of dependents)

**Input Conditions:**
- Valid customer profile
- Complete employment and education data
- Valid relationship status

**Calculation Logic:**
```
creditScore = calculateBaseScore(employmentStatus, highestEducationAttained)
creditScore = adjustForRelationship(creditScore, relationshipStatus, dependents)
creditScore = adjustForHistory(creditScore, customerHistory)

IF creditScore >= 750 THEN
  creditRating = "EXCELLENT"
  creditLimit = calculateLimit(creditScore, "HIGH")
ELSE IF creditScore >= 650 THEN
  creditRating = "GOOD"  
  creditLimit = calculateLimit(creditScore, "MEDIUM")
ELSE IF creditScore >= 550 THEN
  creditRating = "FAIR"
  creditLimit = calculateLimit(creditScore, "LOW")
ELSE
  creditRating = "POOR"
  creditLimit = calculateLimit(creditScore, "MINIMAL")
```

**Output Variables:**
- `creditRating`: CreditRating (rating and source)
- `creditLimit`: AmountOfMoney (currency and amount)
- `creditScore`: Int (calculated score)

**Business Context:** Enables automated credit assessment for loan approvals, credit card limits, and risk management decisions.

**Dependencies:** None (foundational customer assessment)

#### **BR-023: Standing Order Amount Calculations**
**Description:** Currency conversion calculations for standing orders using smallest currency unit conversion with BigDecimal precision for recurring payment processing.

**Source Location:** `code/standingorders/MappedStandingOrder.scala` (lines 33, 88)

**Input Variables:**
- `amountValue`: BigDecimal (standing order amount)
- `amountCurrency`: String (currency code)
- `frequency`: String (payment frequency)

**Input Conditions:**
- Positive amount value
- Valid ISO currency code
- Supported frequency (daily, weekly, monthly, yearly)

**Calculation Logic:**
```
// Convert to smallest currency units for storage
smallestUnits = convertToSmallestCurrencyUnits(amountValue, amountCurrency)

// Store standing order with converted amount
standingOrder.AmountValue = smallestUnits
standingOrder.AmountCurrency = amountCurrency

// Convert back for display/processing
displayAmount = smallestCurrencyUnitToBigDecimal(smallestUnits, amountCurrency)
```

**Output Variables:**
- `smallestUnits`: Long (amount in smallest currency units)
- `displayAmount`: BigDecimal (amount for display)
- `standingOrderId`: String (created order identifier)

**Business Context:** Ensures precise handling of recurring payment amounts without floating-point precision errors in automated payment processing.

**Dependencies:** BR-002 (Currency Conversion)

#### **BR-025: Top API Ranking Calculations**
**Description:** Usage-based ranking algorithm for identifying most frequently used APIs with count-based sorting and statistical analysis.

**Source Location:** `code/metrics/MappedMetrics.scala` (lines 428-436)

**Input Variables:**
- `fromDate`: Date (analysis start date)
- `toDate`: Date (analysis end date)
- `limit`: Int (number of top APIs to return)

**Input Conditions:**
- Valid date range
- Positive limit value
- Available metrics data

**Calculation Logic:**
```
SELECT 
  url,
  COUNT(*) as call_count,
  AVG(duration) as avg_duration,
  SUM(duration) as total_duration
FROM api_metrics 
WHERE date BETWEEN fromDate AND toDate
GROUP BY url
ORDER BY call_count DESC, avg_duration ASC
LIMIT limit
```

**Output Variables:**
- `topAPIs`: List[APIRanking] (ranked API list)
- `callCount`: Long (number of calls per API)
- `averageDuration`: Double (average response time)
- `rank`: Int (API ranking position)

**Business Context:** Identifies popular APIs for optimization, capacity planning, and feature prioritization decisions.

**Dependencies:** BR-019 (Aggregate Metrics)

#### **BR-026: Top Consumer Ranking Calculations**
**Description:** Consumer usage ranking algorithm based on API call frequency and usage patterns for identifying high-value consumers.

**Source Location:** `code/metrics/MappedMetrics.scala` (lines 509-518)

**Input Variables:**
- `fromDate`: Date (analysis start date)
- `toDate`: Date (analysis end date)
- `limit`: Int (number of top consumers to return)

**Input Conditions:**
- Valid date range
- Positive limit value
- Available consumer metrics

**Calculation Logic:**
```
SELECT 
  consumer_id,
  COUNT(*) as total_calls,
  COUNT(DISTINCT url) as unique_apis_used,
  AVG(duration) as avg_response_time,
  SUM(duration) as total_duration
FROM api_metrics 
WHERE date BETWEEN fromDate AND toDate
  AND consumer_id IS NOT NULL
GROUP BY consumer_id
ORDER BY total_calls DESC, unique_apis_used DESC
LIMIT limit
```

**Output Variables:**
- `topConsumers`: List[ConsumerRanking] (ranked consumer list)
- `totalCalls`: Long (total API calls per consumer)
- `uniqueAPIs`: Int (number of different APIs used)
- `rank`: Int (consumer ranking position)

**Business Context:** Identifies high-value API consumers for account management, support prioritization, and business development.

**Dependencies:** BR-019 (Aggregate Metrics)

---

## Rule Dependencies and Relationships

The business rules exhibit several key dependency patterns:

### **Critical Dependencies**
- **BR-002** (Currency Conversion) depends on **BR-001** (FX Rate Lookup)
- **BR-006** (Counterparty Limit Validation) depends on **BR-001** and **BR-002**
- **BR-009** (Challenge Threshold) depends on **BR-002**
- **BR-020** (Advanced Counterparty Validation) depends on **BR-001** and **BR-002**
- **BR-027** (Payment Coverage Check) depends on **BR-001** and **BR-002**

### **Secondary Dependencies**  
- **BR-004** (Transaction Classification) feeds into **BR-017** (Balance Aggregation)
- **BR-010** (View-based Access) controls **BR-011** (Amount Visibility) and **BR-012** (Balance Visibility)
- **BR-013** through **BR-017** (Balance calculations) depend on **BR-004** (Transaction Classification)
- **BR-018** (Rate Limiting) depends on **BR-019** (Aggregate Metrics)
- **BR-019** (Aggregate Metrics) feeds into **BR-025** and **BR-026** (Ranking Calculations)
- **BR-021** (Consumer Counters) depends on **BR-018** (Rate Limiting)
- **BR-023** (Standing Orders) depends on **BR-002** (Currency Conversion)
- **BR-024** (Product Fees) depends on **BR-002** (Currency Conversion)

### **Foundational Rules**
- **BR-001** (FX Rate Lookup) - Core currency system
- **BR-003** (Rate Resolution) - Exchange rate hierarchy  
- **BR-004** (Transaction Classification) - Transaction processing foundation
- **BR-005** (Currency Unit Conversion) - Precision handling
- **BR-019** (Aggregate Metrics) - Analytics foundation
- **BR-022** (Credit Rating) - Customer assessment foundation

### **Analytics Chain**
- **BR-019** (Aggregate Metrics) → **BR-025** (API Rankings) → Business Intelligence
- **BR-019** (Aggregate Metrics) → **BR-026** (Consumer Rankings) → Account Management
- **BR-018** (Rate Limiting) → **BR-021** (Consumer Counters) → API Access Control

## Implementation Notes

### **Precision Requirements**
- All monetary calculations use `BigDecimal` with `HALF_UP` rounding mode
- Currency unit conversion maintains precision through smallest unit representation
- FX rate calculations preserve decimal precision for regulatory compliance
- Standing order amounts use smallest currency units to prevent precision loss

### **Performance Considerations**
- FX rates implement three-tier caching (bank-specific → cached → hardcoded)
- Balance calculations leverage transaction classification for efficient aggregation
- View-based access controls optimize data visibility without compromising security
- Rate limiting uses Redis caching for high-performance API access control
- Aggregate metrics employ SQL optimization for statistical calculations

### **Regulatory Compliance**
- Challenge thresholds support PSD2 Strong Customer Authentication requirements
- Counterparty limits enforce transaction monitoring and AML compliance
- Currency conversion follows standard financial calculation practices
- Payment coverage checks enable TPP compliance with PSD2 requirements
- Credit rating systems support responsible lending practices

### **Scalability Features**
- Multi-period rate limiting supports various API usage patterns
- Consumer ranking algorithms scale with usage volume
- Statistical calculations optimize for large datasets
- Credit assessment systems handle high-volume customer processing

---

**Analysis Completed:** 16-9-2025  
**Total Business Rules:** 27  
**Source Files Analyzed:** 18  
**Categories Covered:** 8
