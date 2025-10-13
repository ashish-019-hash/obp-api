# Business Rules Documentation - OBP-API

## Overview
This document captures the core business calculation rules and decision-making logic extracted from the Open Bank Project (OBP) API codebase. These rules govern financial calculations, eligibility criteria, and business thresholds used throughout the banking platform.

**Analysis Date**: September 14, 2025  
**Source Repository**: karunam2/OBP-API (00.phase-1-input folder)  
**Framework**: Lift Web Framework (Scala)

## Business Rule Categories

### 1. Currency Exchange Rate Calculations

#### 1.1 Fallback Exchange Rate Lookup
**Location**: `code/fx/fx.scala:40-57`  
**Purpose**: Provides default exchange rates when bank-specific rates are unavailable

**Supported Currencies**: 14 major currencies with hardcoded fallback rates
- EUR: 1.0 (base currency)
- USD: 1.12, GBP: 0.87, CHF: 1.08
- JPY: 121.0, CAD: 1.46, AUD: 1.61
- SEK: 10.8, NOK: 9.8, DKK: 7.44
- PLN: 4.31, CZK: 25.4, HUF: 325.0, RUB: 69.0

**Business Logic**:
```scala
def fallbackExchangeRates: Map[String, Double] = Map(
  "EUR" -> 1.0,    // Base currency
  "USD" -> 1.12,   // US Dollar
  "GBP" -> 0.87,   // British Pound
  // ... additional currencies
)
```

#### 1.2 Currency Conversion with Rounding
**Location**: `code/fx/fx.scala:127-130`  
**Purpose**: Standardized currency conversion with consistent rounding

**Rounding Rule**: HALF_UP to 2 decimal places
**Formula**: `amount * exchangeRate` rounded to 2 decimal places

#### 1.3 Three-Tier Rate Resolution Strategy
**Location**: `code/fx/fx.scala:151-162`  
**Purpose**: Hierarchical exchange rate lookup with fallback mechanism

**Resolution Order**:
1. **Bank-specific rates** (highest priority)
2. **Cached rates** (medium priority)  
3. **Hardcoded fallback rates** (lowest priority)

**Business Impact**: Ensures exchange rate availability even when external rate services are unavailable

### 2. Counterparty Transaction Limits

#### 2.1 Six-Dimensional Limit Enforcement
**Location**: `code/counterpartylimit/MappedCounterpartyLimit.scala:54-75`  
**Purpose**: Comprehensive transaction limit validation across multiple dimensions

**Limit Dimensions**:
- **Single Transaction Amount**: Maximum per-transaction limit
- **Monthly Transaction Amount**: Cumulative monthly spending limit
- **Yearly Transaction Amount**: Cumulative annual spending limit
- **Single Transaction Count**: Maximum transactions per operation
- **Monthly Transaction Count**: Maximum transactions per month
- **Yearly Transaction Count**: Maximum transactions per year

**Validation Logic**:
```scala
class MappedCounterpartyLimit {
  object mMaxSingleAmount extends MappedString(this, 32)
  object mMaxMonthlyAmount extends MappedString(this, 32)
  object mMaxYearlyAmount extends MappedString(this, 32)
  object mSingleTransactionCount extends MappedInt(this)
  object mMonthlyTransactionCount extends MappedInt(this)
  object mYearlyTransactionCount extends MappedInt(this)
}
```

### 3. Transaction Processing Rules

#### 3.1 Credit/Debit Classification
**Location**: `code/transaction/MappedTransaction.scala:285-296`  
**Purpose**: Automatic transaction type classification based on amount sign

**Classification Logic**:
- **Positive Amount**: Credit transaction (money received)
- **Negative Amount**: Debit transaction (money sent)
- **Zero Amount**: Neutral transaction

**Business Impact**: Triggers balance update events and account reconciliation

#### 3.2 Currency Unit Conversion
**Location**: `code/transaction/MappedTransaction.scala:111-112`  
**Purpose**: Convert from smallest currency units to decimal representation

**Conversion Rule**: Divide by 100 to convert from cents/pence to major currency units
**Example**: 12345 cents → 123.45 USD

### 4. Product Fee Calculations

#### 4.1 Fee Structure Management
**Location**: `code/productfee/MappedProductFeeProvider.scala:39-49`  
**Purpose**: Standardized fee calculation framework for banking products

**Fee Components**:
- **Amount**: Fee value in specified currency
- **Currency**: Fee currency code (3-letter ISO)
- **Frequency**: Fee application frequency (monthly, yearly, per-transaction)
- **Active Status**: Whether fee is currently applicable

**Fee Structure**:
```scala
class MappedProductFee {
  object mFeeValue extends MappedString(this, 32)
  object mFeeCurrency extends MappedString(this, 3)
  object mFeeFrequency extends MappedString(this, 16)
  object mIsActive extends MappedBoolean(this)
}
```

### 5. Security Challenge Thresholds

#### 5.1 Challenge Threshold Calculation
**Location**: `code/bankconnectors/LocalMappedConnector.scala:152-175`  
**Purpose**: Determine when additional security challenges are required

**Default Threshold**: 1000 (in account currency)
**FX Conversion**: Automatic currency conversion when account currency differs from threshold currency

**Challenge Logic**:
```scala
def getChallengeThreshold(bankId: String, accountId: String, 
                         viewId: String, transactionRequestType: String, 
                         currency: String, userId: String, userName: String): AmountOfMoney = {
  // Default threshold with FX conversion support
  AmountOfMoney(currency, "1000")
}
```

### 6. View-Based Access Control Rules

#### 6.1 Transaction Amount Visibility
**Location**: `code/model/View.scala:155-157`  
**Purpose**: Control transaction amount visibility based on user permissions

**Visibility Rules**:
- **Full Access**: Show actual transaction amounts
- **Restricted Access**: Hide or mask transaction amounts
- **No Access**: Completely hide transaction data

#### 6.2 Account Balance Visibility
**Location**: `code/model/View.scala:175-177`  
**Purpose**: Control account balance visibility with permission checks

**Balance Access Levels**:
- **Owner View**: Full balance visibility
- **Public View**: Limited or no balance information
- **Fiduciary View**: Conditional balance access

## Business Rule Dependencies

### Primary Dependencies
1. **Currency Exchange** → **Transaction Processing** → **Balance Calculation**
2. **Counterparty Limits** → **Transaction Validation** → **Payment Authorization**
3. **Security Thresholds** → **Challenge Requirements** → **Transaction Approval**
4. **View Permissions** → **Data Visibility** → **User Experience**

### Secondary Dependencies
1. **Product Fees** → **Account Charges** → **Revenue Calculation**
2. **FX Rates** → **Multi-currency Operations** → **Financial Reporting**

## Implementation Patterns

### 1. Lift Mapper Pattern
Most business rules are implemented using Lift's Mapper ORM pattern:
```scala
class BusinessRuleEntity extends LongKeyedMapper[BusinessRuleEntity] with IdPK {
  object ruleField extends MappedString(this, length)
  // Business logic methods
}
```

### 2. Provider Pattern
Business rule access follows the Provider pattern:
```scala
object BusinessRuleProvider extends BusinessRuleProvider {
  override def calculateRule(params: Parameters): Result = {
    // Implementation
  }
}
```

### 3. Box Pattern
Error handling and optional results use Lift's Box pattern:
```scala
def businessCalculation(input: String): Box[Result] = {
  // Safe calculation with error handling
}
```

## Configuration and Customization

### 1. Property-Based Configuration
Many thresholds and limits are configurable via properties files:
- Challenge thresholds
- Cache TTL values
- Rate limits

### 2. Bank-Specific Overrides
Rules can be customized per bank:
- Exchange rates
- Fee structures
- Security thresholds

### 3. Currency-Specific Rules
Different rules may apply based on currency:
- Rounding precision
- Minimum/maximum amounts
- Regulatory compliance

## Business Impact Analysis

### 1. Revenue Impact
- **Product Fees**: Direct revenue from banking products
- **FX Spreads**: Revenue from currency conversion margins
- **Transaction Fees**: Per-transaction revenue streams

### 2. Risk Management
- **Transaction Limits**: Fraud prevention and risk mitigation
- **Security Challenges**: Additional authentication for high-risk transactions
- **Balance Visibility**: Information security and privacy protection

### 3. Regulatory Compliance
- **PSD2 Compliance**: Strong customer authentication thresholds
- **AML Requirements**: Transaction monitoring and reporting
- **Data Protection**: View-based access controls

## Technical Implementation Notes

### 1. Database Storage
- Business rules stored in Lift Mapper entities
- Indexed for performance (UniqueIndex, Index patterns)
- Audit trail via CreatedUpdated trait

### 2. Caching Strategy
- Exchange rates cached with configurable TTL
- Metrics aggregation cached for performance
- Challenge thresholds cached per session

### 3. Error Handling
- Box pattern for safe calculations
- Fallback mechanisms for critical operations
- Logging for audit and debugging

## Future Considerations

### 1. Rule Engine Integration
Consider implementing a dedicated rule engine for:
- Dynamic rule modification
- A/B testing of business rules
- Real-time rule updates

### 2. Machine Learning Integration
Potential areas for ML enhancement:
- Dynamic fraud thresholds
- Personalized transaction limits
- Predictive fee optimization

### 3. Regulatory Adaptability
Design for changing regulations:
- Configurable compliance rules
- Regional rule variations
- Automated regulatory reporting

---

**Document Version**: 1.0  
**Last Updated**: September 14, 2025  
**Extracted From**: OBP-API Scala Codebase  
**Total Rules Documented**: 12 core business calculation rules
