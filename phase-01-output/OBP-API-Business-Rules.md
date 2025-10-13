# Open Bank Project API - Comprehensive Business Rules Documentation


This document provides an in-depth analysis of business rules governing the Open Bank Project (OBP) API, a sophisticated open-source banking platform enabling secure, standardized access to account information and banking services. The analysis covers seven role-based perspectives ensuring complete business coverage.

**System Overview:**
- Multi-standard banking API supporting OBP, Berlin Group, and UK Open Banking protocols
- Built on Scala using Lift web framework with advanced security and compliance features
- Over 100 granular view permissions for precise access control
- Six-level rate limiting (second/minute/hour/day/week/month) for API protection
- Dynamic entity creation, webhooks, and real-time event notifications
- Extensive product/account attribute management for customization
- Physical infrastructure management (branches, ATMs) with accessibility features
- Comprehensive card lifecycle management with security controls
- Multi-party meeting scheduling and customer communication tracking
- JSON schema validation and authentication type restrictions per operation
- Bad login tracking, user locking, and signing basket workflows
- FX rate conversion with multi-tier fallback mechanisms
- Transaction types with fee structures and comprehensive attribute management
- Authentication context tracking for users and consents
- CRM event management and API metrics for comprehensive audit trails

---



**What it does:** Separates customer information updates into distinct categories requiring different approval levels

**When it applies:** When updating any customer information field in the system

**Who it affects:** Bank staff updating records, compliance officers reviewing changes, customers whose data is modified

**Example:** Updating a customer's address requires basic permissions, but updating credit rating or credit limit requires special "CanUpdateCustomerCreditRating" or "CanUpdateCustomerCreditLimit" permissions. General information like phone numbers follows different rules than sensitive financial data.

---

**What it does:** Enforces strict state transitions for customer consent with complete audit trail

**When it applies:** Throughout consent lifecycle from initiation to revocation

**Who it affects:** Customers granting consent, third-party providers, compliance teams auditing consent usage

**Example:** A consent must progress INITIATED → ACCEPTED → REVOKED. Cannot jump states. Each transition recorded with timestamps. System tracks time_of_first_use and time_of_last_use for compliance reporting. PSD2 and open banking compliance depends on this state machine.

---

**What it does:** Maintains complete audit trail of all permission grants, revocations, and actual usage

**When it applies:** Whenever any user permission is granted, revoked, or used

**Who it affects:** Compliance officers auditing access, security teams investigating incidents, administrators managing permissions

**Example:** Bank administrator grants "CanCreateCustomer" entitlement to staff member on January 1st. System records: grant timestamp, granter ID, bank scope. On January 15th, staff member uses permission to create customer. System records usage timestamp. On February 1st, administrator revokes entitlement. System records: revocation timestamp, revoker ID. Complete trail enables: "Show me all permissions user X had on date Y" and "Show me everyone who used permission Z in time period".

---

**What it does:** Tracks Know Your Customer verification checks with date, status, and staff documentation

**When it applies:** During customer onboarding and ongoing compliance reviews

**Who it affects:** Compliance officers performing KYC, customers being verified, regulators auditing compliance

**Example:** New customer submits identity documents. Compliance officer creates KYC check record: customer_id, date "2024-01-15", status "PASSED", customer_name, social_media_check "LinkedIn verified", news "No adverse media", comments "Passport verified, utility bill confirmed address". If status changes to "FAILED" or "REVIEW", customer onboarding halts. Regulators can audit all KYC checks performed.

---

**What it does:** Captures comprehensive metadata for every transaction request including timestamps, initiator, status changes

**When it applies:** Throughout transaction request lifecycle from initiation to completion

**Who it affects:** Customers initiating payments, banks processing payments, regulators auditing transactions

**Example:** Customer initiates €1,000 payment at 14:00. System records: start_date "2024-01-15T14:00:00", initiator_id, transaction_request_type "SEPA", status "INITIATED". At 14:01, bank system validates and status becomes "PENDING". At 14:05, final authorization received, status "COMPLETED", end_date "2024-01-15T14:05:00". Complete timeline enables SLA monitoring and dispute resolution.

---

**What it does:** Enforces separate rate limits for authenticated API consumers and anonymous traffic across six time windows

**When it applies:** On every API call to prevent abuse and ensure fair resource distribution

**Who it affects:** API consumers managing call volume, anonymous users, operations teams preventing DoS

**Example:** Consumer "FinTechApp" has limits: 10 calls/second, 100/minute, 1000/hour, 10000/day, 50000/week, 200000/month. Anonymous traffic (no consumer key) limited to: 1/second, 10/minute, 100/hour, 500/day, 2000/week, 5000/month. At 14:00:00, FinTechApp makes 11th call in same second. Rejected: "Rate limit exceeded - 10 per second". System tracks each time window independently. Consumer can burst to 100/minute even if hitting second limit, as long as minute limit not exceeded.

---

**What it does:** Automatically expires consents after configured time-to-live period for regulatory compliance

**When it applies:** System checks consent expiry on every access attempt using the consent

**Who it affects:** Customers whose consents expire, third-party providers losing access, compliance teams ensuring time-limited access

**Example:** Customer grants 90-day consent to PFM app on January 1st. System calculates expiry: time_of_expiration = "2024-04-01T00:00:00" (start + 90 days). On March 15th, app accesses data successfully. On April 2nd, app attempts access. System checks: current_time > time_of_expiration. Access denied: "Consent expired". App must request fresh consent. Ensures regulatory compliance with time-limited data access requirements.

---

**What it does:** Maintains list of countries where customer is tax resident for CRS/FATCA reporting

**When it applies:** During customer onboarding and when customer tax status changes

**Who it affects:** Customers with multi-country tax obligations, compliance teams filing CRS/FATCA reports, tax authorities

**Example:** Customer is US citizen living in Germany doing business in Switzerland. Tax residence records: [{"country_code":"US", "tax_id":"123-45-6789"}, {"country_code":"DE", "tax_id":"12345678901"}, {"country_code":"CH", "tax_id":"CHE-123.456.789"}]. Bank automatically includes customer in FATCA reporting to US, CRS reporting to Germany and Switzerland. When customer changes residence, records updated and reporting adjusted accordingly.

---

**What it does:** Manages account opening applications through defined workflow states from submission to approval/rejection

**When it applies:** When customers apply for new accounts

**Who it affects:** Customers applying for accounts, bank staff reviewing applications, operations tracking application pipeline

**Example:** Customer submits checking account application online. Status "PENDING". Application includes customer_id, product_id "standard_checking", date_of_application "2024-01-15". Compliance reviews KYC. Status "UNDER_REVIEW". After approval, status "APPROVED", account created. Or if declined: status "REJECTED", reason documented. Customer can check application status. Bank tracks conversion rates and processing times by status transitions.

---

**What it does:** Tracks user invitations through lifecycle states with timestamps and secret links

**When it applies:** When inviting new users to access the banking system

**Who it affects:** System administrators sending invitations, invited users, security teams tracking access grants

**Example:** Admin invites new business user. System creates invitation: status "CREATED", secret_link (unique UUID URL), purpose "Business Account Access". Email sent to user. User clicks link, registers. Status "ACCEPTED". System records accepted_date. If user never registers, status remains "CREATED", admin can resend or cancel. If admin revokes before acceptance, status "CANCELLED". Complete lifecycle tracking enables auditing who invited whom and when access was granted.

---

**What it does:** Provides "forget me" functionality to scramble personal data while preserving system referential integrity

**When it applies:** When customers exercise GDPR right to be forgotten

**Who it affects:** Customers requesting data deletion, compliance teams ensuring GDPR compliance, IT teams maintaining referential integrity

**Example:** Customer exercises right to be forgotten. System scrambles fields: username→"SCRUBBED_[timestamp]", email→"scrubbed@example.com", legal_name→"[REDACTED]", password_hash→random, etc. Preserves user_id for referential integrity (transactions, consents reference this). Customer identity erased while system relationships maintained. Enables GDPR compliance without breaking database constraints.

---

**What it does:** Retrieves web UI configuration properties using hierarchical lookup: brand+language → language → requested → default

**When it applies:** When loading UI customization properties for different brands and languages

**Who it affects:** Multi-brand deployments, international customers, operations teams managing branding

**Example:** Requested property "welcome_message" with brand "ACME" and language "de_DE". System looks for: (1) "welcome_message_FOR_BRAND_ACME_de_DE", (2) "welcome_message_de_DE", (3) "welcome_message", (4) default value. This enables brand-specific German welcome message, falling back to generic German, then generic message, then default.

---

**What it does:** Each API endpoint can have custom JSON schema validation configured to enforce specific data format requirements beyond standard API validation

**When it applies:** When a bank needs to enforce additional data structure requirements for specific API operations, administrators can define custom JSON schemas that validate request payloads before processing

**Who it affects:** API consumers submitting requests, banks defining custom validation requirements, compliance officers ensuring data quality standards

**Example:** A bank requires that all transaction requests to corporate accounts include a mandatory "project_code" field with a specific format (e.g., "PROJ-2024-0001"). They configure a JSON schema for the create transaction endpoint that validates this field exists and matches the pattern. When a corporate client attempts to create a transaction without the proper project code, the request is rejected with a clear validation error before any processing occurs.

---

**What it does:** API operations can restrict which authentication methods are allowed (OAuth, Direct Login, Gateway Login, etc.) providing granular authentication policy control per endpoint

**When it applies:** When different operations require different levels of authentication assurance, administrators can configure comma-separated lists of allowed authentication types for each endpoint

**Who it affects:** API consumers choosing authentication methods, security teams defining authentication policies, operations teams managing endpoint access

**Example:** A bank determines that customer data modification endpoints should only allow OAuth 2.0 authentication (not Direct Login) to ensure proper consent management. They configure the update customer endpoint with "allowed_auth_types=OAuth". When a third-party application tries to update customer information using Direct Login credentials, the system rejects the request with "Authentication type not allowed for this operation" even though the credentials are valid.

---

**What it does:** System captures 17 detailed fields per API call including userId, url, date, duration, userName, appName, developerEmail, consumerId, implementedByPartialFunction, implementedInVersion, verb, httpCode, correlationId, responseBody, sourceIp, and targetIp for complete audit capability

**When it applies:** Every API call automatically generates comprehensive metrics that can be queried for audit, compliance, performance analysis, and security investigation purposes

**Who it affects:** Compliance officers conducting audits, security teams investigating incidents, operations teams analyzing performance, developers debugging issues

**Example:** During a regulatory audit, examiners request complete information about all API calls that accessed a specific customer's data during Q1 2024. The compliance team queries the API metrics using the customer's userId and retrieves detailed records showing: which applications accessed the data (appName, developerEmail, consumerId), when they accessed it (date), what operations they performed (url, verb), how long each call took (duration), the network path (sourceIp, targetIp), what data was returned (responseBody if captured), and correlation IDs for tracing across systems. This comprehensive audit trail satisfies the regulatory requirement without requiring custom logging.

---


**What it does:** Controls account information visibility and actions through configurable "views" with 100+ granular permission flags

**When it applies:** Whenever anyone attempts to access account information or perform actions

**Who it affects:** Account holders, customer service representatives, third-party applications, delegates

**Example:** "owner" view allows full account number, balance, complete transaction details. "public" view shows only public alias and transaction dates without amounts. "accountant" view shows transactions with amounts but hides personal info. Each view has 100+ flags like canSeeTransactionAmount, canSeeOtherAccountNumber, canAddTransactionRequest, canCreateStandingOrder.

---

**What it does:** Validates transaction requests match supported types configured for specific account

**When it applies:** When customer or application initiates any transaction request

**Who it affects:** Customers initiating payments, third-party payment apps, bank staff, payment operations

**Example:** UK current account customer tries SEPA credit transfer. System checks account's supported types. If account only supports "FASTER_PAYMENTS", "CHAPS", "INTERNAL_TRANSFER", SEPA request rejected with error listing supported types. Prevents incompatible payment instructions.

---

**What it does:** Enforces seven different limit types simultaneously when sending money to counterparties

**When it applies:** Before authorizing any payment to a counterparty

**Who it affects:** Customers sending payments, recipients, fraud prevention teams, risk management

**Example:** Rent payment counterparty has limits: MAX_LIMIT_PER_TRANSACTION (€2,000), MAX_LIMIT_PER_DAY (€2,000), MAX_LIMIT_PER_WEEK (€2,000), MAX_LIMIT_PER_MONTH (€2,000), MIN_LIMIT_PER_TRANSACTION (€1,000), MAX_NUMBER_OF_TRANSACTIONS_PER_WEEK (1), MAX_NUMBER_OF_TRANSACTIONS_PER_MONTH (1). Ensures exactly one monthly rent payment. Second attempt rejected for violating monthly transaction count.

---

**What it does:** Automatically calculates and applies charges/fees based on type, amount, and charge policy

**When it applies:** When processing transaction requests with associated fees

**Who it affects:** Customers paying fees, receiving parties potentially sharing fees, bank revenue accounting

**Example:** €1,000 international wire with €25 charge policy "SHARED": deduct €1,025 from sender (€1,000 principal + €25 fee) and €25 from recipient, credit €975 to recipient. If "SENDER" policy, only sender pays €25. If "RECEIVER", only recipient charged. Automatic calculation ensures transparent fee handling.

---

**What it does:** Routes customer messages to appropriate staff based on message purpose classification

**When it applies:** When customers send messages through banking system

**Who it affects:** Customers sending inquiries, customer service staff receiving messages, message routing logic

**Example:** Customer sends message with purpose "COMPLAINT" about unauthorized transaction. System routes to fraud investigation team's queue with high priority. Message with purpose "INQUIRY" about balance goes to general customer service. Purpose "SUPPORT" for technical issues routes to IT helpdesk. Ensures messages reach right team for faster resolution.

---

**What it does:** Enforces frequency field must match one of predefined enum values when creating standing orders

**When it applies:** When customers or bank staff create standing orders for recurring payments

**Who it affects:** Customers setting up recurring payments, operations teams processing standing orders

**Example:** Customer attempts standing order with frequency "BIWEEKLY". System rejects: must be "DAILY", "WEEKLY", "MONTHLY", "QUARTERLY", "YEARLY", or other defined values. Prevents ambiguous frequencies that could cause payment processing errors. If customer wants biweekly, must use "WEEKLY" with when_detail specifying weeks.

---

**What it does:** Validates active_to date is after active_from date and prevents collection past expiry

**When it applies:** When creating direct debits and during collection processing

**Who it affects:** Customers authorizing direct debits, merchants collecting payments, operations preventing expired collections

**Example:** Customer authorizes gym membership direct debit: active_from "2024-01-01", active_to "2024-12-31". System validates active_to > active_from. On December 31st, final collection processes. On January 1st 2025, attempted collection rejected: "Direct debit expired". Gym must request new authorization. Prevents unauthorized collections past agreement term.

---

**What it does:** Sends real-time HTTP notifications when specified account events occur

**When it applies:** When configured triggers fire (transaction posted, balance change, etc.)

**Who it affects:** Third-party applications subscribing to events, customers wanting real-time notifications, integration developers

**Example:** Fraud detection service registers webhook for account "ACC123": trigger "TRANSACTION_OVER_5000", url "https://fraud.example.com/alert". When €6,000 transaction occurs, OBP immediately POSTs JSON payload to webhook URL containing transaction details. Fraud system analyzes in real-time and can block suspicious transaction before completion.

---

**What it does:** Tracks card through security states: ACTIVE, LOST, STOLEN, DESTROYED, CANCELLED

**When it applies:** Throughout card lifecycle from issuance to cancellation

**Who it affects:** Cardholders reporting issues, fraud teams blocking cards, operations managing card inventory

**Example:** Customer reports card stolen. Representative updates card_state to "STOLEN". All transactions immediately blocked. Card added to hot list. When attempting purchase, merchant receives "Card reported stolen". If card recovered and theft was false alarm, cannot revert to "ACTIVE" from "STOLEN" for security. Must issue new card. Permanent state prevents fraud.

---

**What it does:** Records when card PIN is reset with timestamp for security audit

**When it applies:** When customer or bank resets card PIN

**Who it affects:** Cardholders resetting PINs, security teams investigating fraud, audit teams tracking security events

**Example:** Customer forgets PIN, calls support. Representative initiates PIN reset, system records pin_reset timestamp "2024-01-15T10:30:00" and increments reset counter. Customer receives temporary PIN. Later, customer reports fraud. Investigator checks PIN reset history. Discovery: PIN reset occurred 1 hour before fraudulent transactions. Indicates possible social engineering. Complete reset history aids fraud investigation.

---

**What it does:** Records reason for card replacement: FIRST, RENEW, LOST, STOLEN, DAMAGED

**When it applies:** When issuing replacement cards

**Who it affects:** Operations tracking replacement patterns, fraud teams analyzing loss trends, cardholders receiving replacements

**Example:** Customer's card expires. System issues replacement: replacement_reason "RENEW", issue_date "2024-01-01". Three months later, customer loses renewed card. New replacement: replacement_reason "LOST". System tracks: customer has 2 replacements in 3 months, second due to loss. Risk system flags: multiple lost cards may indicate carelessness or fraud. Enables pattern analysis.

---

**What it does:** Tracks when card was collected by customer and when it was posted/sent

**When it applies:** During card delivery logistics

**Who it affects:** Operations managing card delivery, customers waiting for cards, audit teams tracking delivery issues

**Example:** Card produced and mailed: issue_date "2024-01-01", posting_date "2024-01-02". Customer picks up from branch: collection_date "2024-01-05". Gap between posting and collection (3 days) is normal. But if collection_date is weeks after posting_date, may indicate delivery issues or customer lost card. Null posting_date means card not yet sent. Enables delivery tracking and issue resolution.

---

**What it does:** Maintains detailed operating hours for branches separately for lobby and drive-up services per day of week

**When it applies:** When customers search for branch hours or check service availability

**Who it affects:** Customers planning visits, branch operations teams, mobile apps showing information, mapping applications

**Example:** Branch has lobby hours: Monday-Wednesday 9:00-17:00, Thursday 9:00-19:00 (extended), Friday 9:00-17:00, Saturday 10:00-14:00, Sunday CLOSED. Drive-up differs: Monday-Friday 8:00-18:00, Saturday 9:00-15:00, Sunday CLOSED. Separate opening/closing times per day (OpeningTimeOnMonday, ClosingTimeOnMonday, etc.) for both lobby and drive-up. Enables accurate "Open until 7pm today" displays.

---

**What it does:** Tracks whether ATM accepts deposits separately from withdrawal capability using hasDepositCapability flag

**When it applies:** When customers search for ATMs accepting deposits or displaying capabilities

**Who it affects:** Customers needing deposits, ATM placement teams, customer service, monitoring systems

**Example:** ATM "ATM001" has hasDepositCapability=true, ATM "ATM002" has hasDepositCapability=false (withdrawal-only). Customer mobile app filter "Accepts deposits" queries where hasDepositCapability=true. Deposit-capable ATMs require more frequent servicing and higher security, enabling differentiated operational procedures.

---

**What it does:** Maintains list of supported currencies per ATM enabling multi-currency cash withdrawal

**When it applies:** When customers search for ATMs supporting specific currencies, particularly in airports or tourist areas

**Who it affects:** International travelers, tourists, foreign currency exchange operations, ATM placement teams

**Example:** Airport ATM has supportedCurrencies ["USD", "EUR", "GBP", "JPY", "CHF"] enabling multi-currency withdrawals. Tourist area ATM has ["EUR"] only. US customer searching for USD withdrawals in Europe finds only multi-currency airport ATM. Customer selects ATM and chooses to withdraw in USD or EUR. Optimizes currency stocking at high-traffic ATMs.

---

**What it does:** Transaction attributes can be filtered by view permissions through AttributeDefinition integration, allowing different customer service levels to see different transaction metadata based on their assigned views

**When it applies:** When banks need to control which transaction attributes are visible to different user types or service tiers, they configure AttributeDefinitions with "canBeSeenOnViews" lists that specify which views can access each attribute type

**Who it affects:** Customer service representatives with different access levels, customers with different account views, product managers defining service tiers

**Example:** A bank offers premium and standard account tiers. Premium customers should see detailed merchant category codes and loyalty points on their transactions, while standard customers see only basic transaction information. The bank creates AttributeDefinitions for "merchant_category" and "loyalty_points" attributes with canBeSeenOnViews=["owner","premium_view"] but excludes "standard_view". When a standard customer views their transactions through the API, these attributes are automatically filtered out, while premium customers see the complete information.

---

**What it does:** System tracks both scheduled and actual dates for CRM events, enabling performance monitoring of customer service delivery and identification of delayed actions

**When it applies:** When customer service teams schedule follow-up actions, callbacks, or other customer interactions, the system records when the event was planned to occur and when it actually occurred

**Who it affects:** Customer service managers monitoring team performance, representatives managing their task schedules, customers expecting timely service

**Example:** A customer service representative schedules a callback to a customer for Monday at 2 PM to discuss a loan application. The system records scheduledDate="2024-01-15T14:00:00" and creates a CRM event. The representative actually makes the call on Tuesday at 10 AM due to high call volume. The system records actualDate="2024-01-16T10:00:00" when the representative logs the interaction. The customer service manager reviews weekly metrics and identifies that 15% of callbacks occurred later than scheduled, indicating a need for additional staffing during peak periods.

---

**What it does:** Events are categorized by type and communication channel with result tracking, providing comprehensive customer interaction history across all touchpoints

**When it applies:** Every customer interaction is logged with a category (e.g., "complaint", "inquiry", "follow-up"), channel (e.g., "phone", "email", "in-person"), and result (e.g., "resolved", "escalated"), enabling pattern analysis and service quality monitoring

**Who it affects:** Customer service teams logging interactions, managers analyzing customer engagement patterns, customers receiving consistent service across channels

**Example:** A customer initially emails (channel="email") about unauthorized transaction (category="fraud_inquiry"). The system logs the event with result="investigation_started". The case is escalated and a phone call is made (channel="phone", category="follow-up") where the issue is resolved (result="resolved"). When the customer contacts the bank again months later, any representative can view the complete interaction history showing the sequence: email inquiry → investigation → phone resolution, understanding the full context regardless of which channel the customer uses.

---

**What it does:** System records when users accept agreements with hash verification to detect agreement changes, ensuring users have accepted current versions of terms and conditions

**When it applies:** When banks update terms of service, privacy policies, or other customer agreements, the system tracks which users have accepted which versions by storing agreement hashes and acceptance timestamps

**Who it affects:** Compliance officers ensuring proper consent, legal teams managing agreement versions, customers who must acknowledge updated terms

**Example:** A bank updates its privacy policy on June 1st to comply with new regulations. The system calculates a new hash for the agreement text. When a customer logs in, the system compares the hash of the most recent agreement the customer accepted against the current hash. They don't match, so the system presents the updated privacy policy and requires acceptance before the customer can proceed. Upon acceptance, the system records acceptedDate="2024-06-03T09:15:00" and stores the current agreement hash, creating an audit trail showing the customer accepted the new terms.

---

**What it does:** Each transaction type has an associated fee structure with amount and currency that is automatically applied based on the transaction category

**When it applies:** When transactions are created, the system looks up the transaction type's configured fee and applies it to the transaction, ensuring consistent fee application across all transactions of the same type

**Who it affects:** Product managers defining fee structures, customers paying transaction fees, treasury teams reconciling fee revenue

**Example:** A bank configures transaction type "ATM_WITHDRAWAL_OTHER_BANK" with shortCode="ATM_OTHER", summary="ATM withdrawal at another bank", and charge={"currency":"USD","amount":"2.50"}. When a customer withdraws cash from another bank's ATM, the system creates a transaction with this transaction type. The $2.50 fee is automatically recorded and deducted from the customer's account. The customer sees on their statement: "ATM Withdrawal - Other Bank" with a separate line item "ATM Fee: $2.50", both linked to the same transaction type for consistent presentation and fee transparency.

---


**What it does:** Requires additional authentication challenges for high-risk transactions using secure BCrypt hashing

**When it applies:** When transaction exceeds risk thresholds or involves sensitive operations

**Who it affects:** Customers performing sensitive operations, fraud prevention systems, transaction processors

**Example:** Customer transfers €10,000 to new beneficiary. System identifies high-risk and generates challenge sent via SMS or push notification. Customer must provide correct answer within 5 minutes. System stores challenge answer as BCrypt hash. When customer submits answer, it's hashed with BCrypt and compared to stored hash. Only after successful verification does transaction proceed. Failed challenges logged for security analysis.

---

**What it does:** Controls what action types (CREDIT, DEBIT, CASH) are allowed on card based on allows list

**When it applies:** When processing card transactions at merchants, ATMs, or payment terminals

**Who it affects:** Card holders attempting transactions, merchants processing payments, ATM networks, card issuers

**Example:** Card has allows ["DEBIT", "CASH"] but not "CREDIT". Cardholder can make debit purchases and withdraw cash, but cannot use for credit transactions. When merchant attempts credit transaction, payment network checks allows list and declines with "transaction type not permitted". Empty allows list means no transactions permitted. Enables banks to issue cards with specific capabilities matching account types.

---

**What it does:** Stores card CVV numbers as SHA-256 hashes rather than plaintext for security

**When it applies:** When creating or updating physical card records in database

**Who it affects:** Security teams protecting card data, compliance teams ensuring PCI DSS compliance, cardholders whose data is protected

**Example:** New card issued with CVV "123". System immediately hashes: SHA-256("123") = "a665a45...". Only hash stored in database. When customer enters CVV for online purchase, merchant system hashes entered value and bank compares hashes. Match confirms CVV correct without bank ever storing plaintext. Even database breach doesn't expose CVVs. Meets PCI DSS requirement not to store plaintext CVV.

---

**What it does:** Tracks API usage across six independent time windows simultaneously with separate limits per window

**When it applies:** On every API call to enforce rate limits

**Who it affects:** API consumers managing usage patterns, abuse prevention systems, fair usage enforcement

**Example:** Consumer makes 9 calls at 14:00:00 (within 10/second limit). At 14:00:01, makes 2 more calls (11 total in minute). Second limit not violated (100/minute). At 14:00:02, makes 90 more calls rapidly. Total in minute: 101 calls. Rejected: "Minute limit exceeded". But second limit resets each second, so continuing calls at <10/second allowed. Each window (second, minute, hour, day, week, month) tracked independently, enabling burst handling while preventing sustained abuse.

---

**What it does:** Validates from_date is before to_date and prevents execution outside date range

**When it applies:** When creating standing orders and during execution scheduling

**Who it affects:** Customers setting up recurring payments, operations executing standing orders, systems scheduling payments

**Example:** Customer creates standing order: from_date "2024-01-01", to_date "2024-12-31", frequency "MONTHLY". System validates from_date < to_date. Standing order executes first time on or after January 1st. Executes monthly through December. On January 1st 2025, execution attempted but blocked: "Standing order ended". Customer must create new standing order for 2025 if recurring payment should continue.

---

**What it does:** Links direct debits to both customer (who authorized) and user (who created record) for accountability

**When it applies:** When creating direct debit mandates

**Who it affects:** Customers authorizing direct debits, bank staff creating records, audit teams tracking authorizations

**Example:** Customer authorizes direct debit at branch. Bank representative creates record: customer_id "CUST123" (who authorized), user_id "USER456" (representative who entered it). Later, unauthorized direct debit dispute. Investigation shows: customer_id matches legitimate customer, but user_id belongs to representative who had no authorization to create this mandate. Separate tracking enables identifying internal fraud or process violations.

---

**What it does:** Routes API operations to appropriate connector implementation based on pattern matching rules

**When it applies:** When API operation needs to interact with core banking or external systems

**Who it affects:** Integration teams configuring routing, operations teams managing connectors, system architects

**Example:** Bank has multiple core banking systems. Routing rule: if bank_id matches "uk.*" use "uk_connector", if bank_id matches "de.*" use "sepa_connector", otherwise use "default_connector". API call for bank "uk.hsbc.001" matches first pattern, routes to UK connector. Call for "de.bundesbank" routes to SEPA connector. Call for "us.bofa" uses default. Pattern-based routing enables multi-connector deployments without hardcoding.

---

**What it does:** Validates dynamic entity reference fields match one of 20+ predefined types (CUSTOMER, ACCOUNT, etc.)

**When it applies:** When creating custom dynamic entities with references to other entities

**Who it affects:** Product managers defining custom entities, developers using dynamic entities, data integrity enforcement

**Example:** Bank creates dynamic entity "LoanApplication" with reference to customer. System validates entity_reference_type is "CUSTOMER" (from allowed list). Later, developer tries creating entity with entity_reference_type "INVALID". Rejected: must be CUSTOMER, ACCOUNT, TRANSACTION, BANK, USER, etc. Prevents typos and ensures references resolve correctly. Type validation maintains referential integrity in dynamic entity system.

---

**What it does:** Controls which account attributes are visible based on user's view permissions

**When it applies:** When accessing account attributes through views with different permission levels

**Who it affects:** Users with different view access, privacy enforcement, service tier differentiation

**Example:** Account has attributes: "credit_score" and "account_nickname". AttributeDefinition for "credit_score" specifies canBeSeenOnViews=["owner"]. "account_nickname" specifies ["owner","accountant","auditor"]. User with "owner" view sees both. User with "accountant" view sees only nickname, not credit score. Privacy-sensitive attributes automatically filtered by view permissions without manual code.

---

**What it does:** Tracks ATM fee structure separately for national vs international cards with currency specification

**When it applies:** When customers check ATM fees before withdrawal or after transaction

**Who it affects:** Customers evaluating ATM fees, ATM operations setting pricing, transparency compliance

**Example:** ATM has fee_for_use_local {"currency":"EUR", "amount":"0"} (free for local cards) and fee_for_use_foreign {"currency":"EUR", "amount":"3.50"} (€3.50 for foreign cards). Customer with domestic card sees "No fee" before withdrawal. International tourist sees "€3.50 fee will apply" and can choose different ATM or accept fee. Transparent fee disclosure prevents surprise charges and meets regulatory requirements.

---

**What it does:** Marks branches as deleted using is_deleted flag rather than removing from database

**When it applies:** When branches close but historical data must be retained

**Who it affects:** Operations closing branches, customers with historical transactions at closed branches, reporting teams

**Example:** Branch closes on June 30th. System sets is_deleted=true, last_update=closing_date. Branch no longer appears in active branch searches or mobile apps. But historical transactions reference branch_id and still resolve to branch details including "Branch closed June 30, 2024". Annual reports can still show transaction volumes by branch including closed branches. Soft delete preserves referential integrity and historical accuracy.

---

**What it does:** Enables multi-party meetings with unique access tokens per participant

**When it applies:** When scheduling meetings requiring multiple stakeholders like loan approvals

**Who it affects:** Meeting organizers, participants joining meetings, compliance teams auditing decision meetings

**Example:** Loan officer schedules loan approval meeting for 3-member committee. System generates unique meeting_id and 3 tokens (one per invitee). Each committee member receives email with personal token. At meeting time, each logs in with their token. System tracks: who attended (by token used), when they joined, meeting start/end times. If all 3 don't attend, loan approval incomplete. Audit trail shows which committee members participated in which decisions.

---

**What it does:** Configurable maximum bad login attempts (default 5) before automatic user lockout prevents brute force attacks and unauthorized access attempts

**When it applies:** When a user fails to authenticate with correct credentials, the system increments their bad login attempt counter. Upon reaching the configured threshold, the user account is automatically locked

**Who it affects:** Users entering incorrect passwords, security teams monitoring suspicious login patterns, attackers attempting credential stuffing or brute force attacks

**Example:** A bank configures max_bad_login_attempts=5. A user forgets their password and tries to log in 5 times with incorrect credentials. On the 5th failed attempt, the system automatically locks their account and records the lock reason. The user receives a message: "Account locked due to multiple failed login attempts. Please contact customer service or use the password reset feature." The security team receives an alert about the account lock and can investigate whether this was a legitimate user or an attack attempt. This prevents unlimited password guessing while protecting legitimate users who simply forgot their credentials.

---

**What it does:** Successful authentication resets the bad login attempt counter to zero, enabling users to continue after failed attempts without permanent penalties

**When it applies:** When a user successfully authenticates after previous failed attempts (but before reaching the lockout threshold), the system clears their bad login attempt counter

**Who it affects:** Users who occasionally mistype passwords, security monitoring systems tracking authentication patterns

**Example:** A user attempts to log in but makes typos on their first two attempts due to rushing. On the third attempt, they carefully enter their correct password and authenticate successfully. The system resets their bad login counter from 2 to 0. Later that day, they log in again and accidentally mistype their password once before succeeding. Since the counter was reset to 0 after their successful morning login, this new failure only counts as attempt 1 of 5, not attempt 4 of 5. This approach balances security (preventing brute force) with usability (forgiving occasional mistakes).

---

**What it does:** Administrators can manually lock user accounts via API with type "lock_via_api" and timestamp tracking for emergency security actions or policy enforcement

**When it applies:** When administrators need to immediately prevent a user from accessing the system due to security concerns, policy violations, or other administrative reasons

**Who it affects:** Security administrators responding to incidents, compliance officers enforcing policies, users whose accounts are locked, audit teams reviewing lock history

**Example:** The fraud detection team identifies suspicious activity on a customer's account suggesting the credentials may have been compromised. A security administrator immediately calls the lockUser API with the customer's provider and username. The system creates a UserLock record with typeOfLock="lock_via_api" and lastLockDate=current timestamp. The customer's next login attempt is immediately denied with "Account locked" message. Meanwhile, the security team investigates the suspicious transactions and contacts the customer to verify their identity. Once the situation is resolved, the administrator calls unlockUser API to restore access.

---

**What it does:** System checks both BadLoginAttempt table (automatic lockout) and UserLocks table (manual locks) for comprehensive user access control

**When it applies:** Every authentication attempt checks both locking mechanisms to determine if the user should be allowed to log in

**Who it affects:** All users attempting to authenticate, security teams managing access control

**Example:** A user account has 3 failed login attempts in the BadLoginAttempt table (below the threshold of 5) but a security administrator has manually locked the account via API due to a fraud investigation. When the user tries to log in with correct credentials, the system checks both tables: BadLoginAttempt (3 attempts - would allow login) and UserLocks (lock exists - denies login). The manual lock takes precedence and the login is denied, even though the automatic lockout threshold wasn't reached. This dual-check ensures administrators can override automatic systems when necessary.

---

**What it does:** Signing baskets follow a defined status lifecycle from RCVD (received) to CANC (cancelled) managing multi-party approval workflows for payments and consents

**When it applies:** When payments or consents require approval from multiple parties, they are placed in signing baskets that track approval status and manage the workflow

**Who it affects:** Corporate customers requiring multi-party approval, compliance officers enforcing maker-checker controls, treasury teams managing payment approvals

**Example:** A corporation configures dual approval for payments over $100,000. An employee creates a payment for $150,000, and the system automatically creates a signing basket with status="RCVD". The basket contains the payment ID and requires approvals from two authorized signers. The first signer reviews and approves. The second signer is unavailable and the payment deadline passes, so a treasurer cancels the signing basket, updating status="CANC". The system notifies all parties that the payment was cancelled and will need to be resubmitted. The signing basket maintains complete audit trail of: creation time, who created it, who approved/rejected, and who cancelled it.

---


**What it does:** Enforces type validation on product attributes supporting STRING, INTEGER, DOUBLE, BOOLEAN, DATE_WITH_DAY

**When it applies:** When banks configure product attributes during product definition or updates

**Who it affects:** Product managers defining products, system validation, customers affected by product rules

**Example:** Product "Premium Savings" has attribute "minimum_balance" type=INTEGER, value="10000". Another attribute "interest_rate" type=DOUBLE, value="2.5". Attempting to set minimum_balance="high" rejected: "Expected INTEGER". Type enforcement prevents data quality issues that could cause incorrect product behavior or customer confusion.

---

**What it does:** Account attributes can inherit from product attribute definitions creating consistent attribute sets

**When it applies:** When creating accounts based on product templates

**Who it affects:** Product managers defining templates, operations creating accounts, customers with consistent product features

**Example:** Product "Business Checking" defines attributes: "overdraft_limit" (INTEGER, "5000"), "monthly_fee" (DOUBLE, "25.00"), "free_transactions" (INTEGER, "100"). When account created from this product, account attributes automatically populated with these values. Product manager updates product template's monthly_fee to "20.00". Existing accounts unchanged (attributes copied at creation), new accounts get updated value. Enables product standardization with account-level customization.

---

**What it does:** Groups products into collections enabling bundle offerings and themed product sets

**When it applies:** When banks want to market related products together or offer bundles

**Who it affects:** Product managers creating bundles, marketing teams promoting offerings, customers browsing product catalogs

**Example:** Bank creates "Student Banking Package" collection containing: "Student Checking" product, "Student Savings" product, "Student Credit Card" product. Collection has collection_code "STUDENT2024", description, metadata. Mobile app "Student Offers" section queries products by collection. Customer sees all 3 products together. Signing up for bundle applies to all 3 products simultaneously. Enables package deals and targeted marketing.

---

**What it does:** Specifies fee frequency (ONCE, MONTHLY, ANNUALLY) and active status for each product fee

**When it applies:** When configuring product pricing and fee schedules

**Who it affects:** Product managers setting pricing, billing systems charging fees, customers paying fees

**Example:** Product has 3 fees: account_opening (frequency=ONCE, active=true, amount="50"), monthly_maintenance (frequency=MONTHLY, active=true, amount="15"), annual_review (frequency=ANNUALLY, active=false, amount="100"). At account opening, customer charged $50 once. Each month, $15 maintenance fee. Annual review fee inactive (active=false) so not charged even though defined. Enables flexible fee structures with easy activation/deactivation.

---

**What it does:** Controls product visibility based on bank license with superadmin override capability

**When it applies:** When displaying available products to customers or staff

**Who it affects:** Product managers restricting product availability, banks with licenses, superadmin users

**Example:** Premium investment product requires license "INVESTMENT_LICENSE". Bank A has license, Bank B doesn't. Customer of Bank A sees product in catalog. Customer of Bank B doesn't see it. Bank B product manager (non-superadmin) also doesn't see it. Superadmin user logs in and sees all products regardless of licenses for system administration. Enables license-based product distribution while allowing system oversight.

---

**What it does:** Automatically generates Create, Read, Update, Delete API endpoints when dynamic entity defined

**When it applies:** When administrators create new dynamic entities for custom business objects

**Who it affects:** Administrators defining entities, developers consuming APIs, applications using custom entities

**Example:** Bank defines dynamic entity "LoanApplication" with fields: applicant_id, loan_amount, property_address, employment_status. System automatically generates: POST /dynamic-entities/LoanApplication (create), GET /dynamic-entities/LoanApplication/:id (read), PUT /dynamic-entities/LoanApplication/:id (update), DELETE /dynamic-entities/LoanApplication/:id (delete). No custom code needed. APIs immediately available for use. Enables rapid custom entity deployment.

---

**What it does:** Allows product attributes to be marked active or inactive without deletion

**When it applies:** When product features need to be temporarily disabled or phased out

**Who it affects:** Product managers managing feature rollouts, customers affected by feature availability

**Example:** Product "Premium Checking" has attribute "concierge_service" active=true. During cost cutting, product manager sets active=false. Attribute still exists in system (historical data preserved) but new accounts don't receive concierge service. Later, when budget allows, set active=true to restore feature. Enables feature lifecycle management without data loss.

---

**What it does:** Enables hierarchical product collections with nested sub-collections and full product details at each level

**When it applies:** When displaying product catalogs with category hierarchies

**Who it affects:** Product managers organizing catalogs, customers browsing product hierarchies, mobile apps displaying structures

**Example:** Top-level collection "Personal Banking" contains sub-collections: "Deposit Accounts" and "Lending Products". "Deposit Accounts" contains: "Checking Accounts" collection and "Savings Accounts" collection. Each collection includes full product details. Customer browses: Personal Banking → Deposit Accounts → Savings Accounts → sees list of specific savings products with full details. Enables intuitive navigation of complex product portfolios.

---

**What it does:** Stores bank-level configuration as type-safe attributes supporting STRING, INTEGER, DOUBLE, BOOLEAN, DATE_WITH_DAY

**When it applies:** When configuring bank-specific settings, features, regulatory requirements, or operational parameters

**Who it affects:** Operations teams configuring banks, compliance teams setting requirements, feature flag management

**Example:** Bank configuration includes attributes: "swift_code" (STRING, "DEUTDEFF"), "max_transaction_amount" (INTEGER, "1000000" meaning €1M), "psd2_compliant" (BOOLEAN, "true"), "license_expiry" (DATE_WITH_DAY, "2025-12-31"), "overdraft_interest_rate" (DOUBLE, "12.5"). Each attribute has defined type. When system components need configuration, they query bank attributes by name and get type-safe values. Transaction processing queries "max_transaction_amount" and gets integer 1000000, then compares transaction amounts against this limit.

---

**What it does:** Transforms API requests and responses between internal and external formats using JSON mapping rules stored in database

**When it applies:** When routing API calls through transformation layers, supporting multiple API standards, or adapting to different core banking systems

**Who it affects:** API gateway operators, system integrators, developers using transformed endpoints, multi-standard API support teams

**Example:** Endpoint with operation_id "getAccount" serves UK Open Banking clients. request_mapping JSON transforms incoming UK Open Banking format to internal OBP format: maps UK field "Account.AccountId" to OBP "account_id", "Account.Currency" to "currency". response_mapping transforms OBP response back to UK format: maps "account_id" to "Data.Account[0].AccountId", adds required UK fields like "Status" and "StatusUpdateDateTime". Mapping stored as JSON strings in database and can be updated without code changes.

---

**What it does:** Allows endpoint mapping rules to be configured globally (all banks) or for specific banks (bank_id specified)

**When it applies:** When different banks need different API transformations or providing bank-specific API customization

**Who it affects:** Multi-bank deployment administrators, bank-specific integration teams, API operations teams

**Example:** Global mapping rule (bank_id=null) for operation_id "createTransaction" applies to all banks by default. Bank "gh.29.uk" needs special transformation for regulatory reasons, so bank-specific mapping created with bank_id "gh.29.uk" and operation_id "createTransaction". When processing createTransaction for "gh.29.uk", system finds bank-specific mapping and uses it. For all other banks, global mapping used. Method getByOperationId(bank_id, operation_id) looks for bank-specific mapping first, then falls back to global.

---

**What it does:** Each bank defines unique short codes for transaction types, enabling custom categorization while maintaining uniqueness per bank to prevent conflicts

**When it applies:** When banks create or configure transaction types, the system enforces that short codes are unique within each bank's scope but allows different banks to use the same short codes for different purposes

**Who it affects:** Product managers defining transaction categorization schemes, operations teams processing transactions, customers viewing transaction categories

**Example:** Bank A defines transaction type short code "INT" for "International Wire Transfer" while Bank B uses "INT" for "Interest Payment". Both configurations are valid because short codes are unique within each bank, not globally. When Bank A processes an international wire, it uses their "INT" type showing "International Wire Transfer" to customers. Bank B's interest payments use their "INT" type showing "Interest Payment". This bank-specific scoping allows each institution to customize transaction categorization to match their products and customer base without coordination with other banks.

---

**What it does:** Card attributes support typed values (STRING, INTEGER, DOUBLE, BOOLEAN, DATE_WITH_DAY) with automatic validation ensuring data integrity for card metadata

**When it applies:** When banks store additional information about physical or virtual cards, they define attributes with specific types that the system validates on creation and update

**Who it affects:** Product managers defining card attributes, operations teams managing card data, customers whose card features depend on attribute values

**Example:** A bank offers premium credit cards with higher reward points multipliers. They create a card attribute "rewards_multiplier" with type=DOUBLE. When issuing a premium card, they set rewards_multiplier=2.5. When a customer makes a purchase, the transaction system reads this attribute and applies 2.5x points. If someone tries to set rewards_multiplier="high" (a string), the system rejects it with "Invalid type: expected DOUBLE, got STRING". This type safety prevents data quality issues that could cause incorrect rewards calculations.

---

**What it does:** Card attributes can be scoped to specific banks or globally available across all banks, providing flexibility in attribute definition and management

**When it applies:** When defining card attributes, banks can choose whether the attribute is specific to their bank's cards or should be available for cards across the entire system

**Who it affects:** Product managers defining card programs, system administrators managing shared attributes, fintech developers building multi-bank applications

**Example:** A payment network wants to add a "network_token_status" attribute applicable to all cards in the system regardless of issuing bank. They create this attribute with bankId=null (global scope). Meanwhile, a specific bank wants to track their proprietary "vip_concierge_tier" attribute only for their premium cardholders. They create this attribute with their specific bankId. When querying card attributes, the VIP tier attribute only appears for that bank's cards, while network_token_status appears for all cards system-wide.

---

**What it does:** Transaction attributes are scoped to specific bank and transaction combinations, ensuring proper data isolation and preventing cross-contamination of transaction metadata

**When it applies:** When creating or querying transaction attributes, the system enforces that attributes belong to specific bank-transaction pairs and prevents access to attributes outside this scope

**Who it affects:** Banks storing transaction metadata, compliance teams ensuring data isolation, developers building transaction enrichment features

**Example:** Bank A stores a transaction attribute "merchant_category_code"="5411" on transaction TX123 at their bank. Bank B also has a transaction with ID TX123 (IDs can coincide across banks) with merchant_category_code="5812". When Bank A queries attributes for their transaction TX123, they only receive "5411", not Bank B's "5812". This scoping prevents banks from seeing each other's transaction metadata even when transaction IDs happen to match, maintaining proper multi-tenancy and data privacy.

---


**What it does:** Generates unique secret links for user invitations serving as secure, non-guessable invitation tokens

**When it applies:** When inviting new users to access banking system

**Who it affects:** System administrators sending invitations, invited users receiving access links, security teams

**Example:** Admin creates invitation for new contractor. System generates secret_link with format "https://api.example.com/invite/{UUID}" where UUID is random 128-bit UUID like "550e8400-e29b-41d4-a716-446655440000". Link emailed to contractor. secret_link must be globally unique (database constraint) to prevent collision attacks. Link contains no predictable information (sequential IDs, timestamps, user data) preventing link guessing. After registration, status→ACCEPTED and link can no longer be used, preventing replay attacks.

---

**What it does:** Enables real-time HTTP/HTTPS notifications when account events occur with per-webhook enable/disable toggle

**When it applies:** When configuring event-driven integrations, real-time monitoring, or third-party service notifications

**Who it affects:** Integration developers, external monitoring systems, third-party applications, fraud detection systems, webhook management tools

**Example:** Fraud system registers webhook for account "ACC123": trigger "TRANSACTION_OVER_5000", http_protocol "HTTPS", http_method "POST", url "https://fraud.example.com/alert", is_enabled=true. When €6,000 transaction occurs, OBP makes HTTPS POST to URL with JSON payload. Receiving system can request webhook disablement temporarily during maintenance by calling API to set is_enabled=false, without deleting configuration. Later, is_enabled=true to resume notifications.

---

**What it does:** Enforces unique combinations of api_collection_id + operation_id preventing duplicate endpoint entries

**When it applies:** When adding endpoints to API collections

**Who it affects:** API collection administrators, documentation teams, endpoint management

**Example:** API collection "Customer_Operations_v1" already contains endpoint operation_id "createCustomer". Administrator attempts to add same operation_id again to collection. System rejects: "Combination already exists". Prevents duplicate entries that would confuse documentation or cause collection queries to return duplicates. Ensures each endpoint appears exactly once per collection.

---

**What it does:** Tracks detailed accessibility features per branch including wheelchair access, audio support, braille signage

**When it applies:** When customers with accessibility needs search for suitable branches

**Who it affects:** Customers with disabilities, branch operations, ADA compliance teams, mobile apps

**Example:** Branch has accessibility features: has_wheelchair_accessibility=true, has_audio_loop=true (hearing assistance), has_braille_signage=false. Customer using wheelchair searches for accessible branches. Mobile app filters where has_wheelchair_accessibility=true. Customer finds suitable branch. Deaf customer searches for audio loop support. Missing braille signage noted for compliance improvement. Enables targeted accessibility information.

---

**What it does:** Lists languages supported by ATM interface enabling localized user experience

**When it applies:** When customers interact with ATM in their preferred language

**Who it affects:** International customers, tourists, multilingual communities, ATM interface design

**Example:** Airport ATM supports ["EN", "DE", "FR", "ES", "IT", "ZH", "JA", "AR"]. Suburban ATM supports ["EN", "ES"]. Chinese tourist searches for ATMs with Chinese language support. Finds only airport ATM. Tourist navigates to airport location. Selects "ZH" from language menu and completes transaction in Chinese. Supported languages stored and searchable enables language-based ATM location finding.

---

**What it does:** Tracks detailed service capabilities per ATM beyond just deposit/withdrawal (bill pay, check deposit, account opening, etc.)

**When it applies:** When customers search for ATMs with specific capabilities

**Who it affects:** Customers needing specific services, ATM operations teams, service capability planning

**Example:** ATM "ATM001" supports: "WITHDRAWAL", "BALANCE_INQUIRY", "MINI_STATEMENT". ATM "ATM002" supports: "WITHDRAWAL", "BALANCE_INQUIRY", "MINI_STATEMENT", "DEPOSIT", "BILL_PAYMENT", "CHECK_DEPOSIT". Customer wants to pay utility bill. Searches for ATMs with "BILL_PAYMENT" capability. Finds only ATM002. Navigates there and completes bill payment. Supported services as list enables precise capability-based ATM discovery.

---

**What it does:** Maintains completely separate operating hours for drive-up and lobby services

**When it applies:** When branches offer drive-up services with different hours than walk-in lobby

**Who it affects:** Customers choosing service method, branch operations, mobile apps displaying hours

**Example:** Branch lobby: Monday-Friday 9:00-17:00, Saturday CLOSED. Drive-up: Monday-Friday 7:00-19:00, Saturday 8:00-14:00. Working parent can't visit during business day, uses drive-up at 18:00 (lobby closed). Saturday shopper needs in-person service but lobby closed; must use drive-up window. Separate hour tracking enables accurate "Drive-up open until 7pm" messages when lobby closed.

---

**What it does:** Supports property name aliases enabling property renames without breaking existing deployments

**When it applies:** When API properties need to be renamed but backward compatibility must be maintained

**Who it affects:** Operations teams managing property files, developers using properties, deployment upgrades

**Example:** Property originally named "max_tx_limit" needs to be renamed to "maximum_transaction_limit" for clarity. System configuration adds alias: "max_tx_limit" → "maximum_transaction_limit". Old deployments still using "max_tx_limit" continue working (alias resolution). New deployments use "maximum_transaction_limit". During upgrade transition period, both names work. Eventually old name deprecated. Enables gradual property migration without service disruption.

---

**What it does:** Links user records to original consent or invitation that authorized their creation

**When it applies:** When creating users via consent flows or invitation processes

**Who it affects:** Audit teams tracking user origins, compliance verifying proper authorization, security investigating access

**Example:** Customer grants consent to third-party app. App creates user via API. User record includes consent_id linking to original consent. Later, security questions: "Why does this third-party have user account?" Audit trail shows: user created via consent_id "CNS123", granted by customer on specific date with specific scopes. Proves legitimate authorization. Without link, user creation origin unclear. Tracking enables proving proper authorization chain.

---

**What it does:** Enables finding users by provider+username OR provider+user_id OR email, with email potentially returning multiple results

**When it applies:** When systems need to locate users using different identifiers

**Who it affects:** Integration systems doing lookups, support staff finding users, multi-provider environments

**Example:** System knows user's email "john@example.com" but not provider details. Calls getUsersByEmail. Returns array of user objects (could be multiple providers with same email). System picks user matching provider. Another system knows user_id "12345" from external system. Calls getUserByProviderId(provider, user_id). Gets exact user. Support staff has username "jsmith". Calls getUserByProviderAndUsername. Finds user. Multiple lookup methods enable flexible user resolution.

---

**What it does:** Signing baskets can contain multiple payment IDs and/or consent IDs, enabling batch approval workflows where multiple related items are approved together

**When it applies:** When organizations need to approve multiple related payments or consents as a group, they can be placed in a single signing basket for collective approval rather than individual approval workflows

**Who it affects:** Corporate treasurers managing batch payments, compliance officers approving multiple consents, authorized signers reviewing grouped transactions

**Example:** A corporation's payroll department processes bi-weekly salary payments for 500 employees. Rather than creating 500 separate signing baskets requiring individual approval, they create one signing basket containing all 500 payment IDs. The signing basket requires approval from two signers per company policy. The first signer reviews the batch total and employee list, then approves. The second signer independently verifies and approves. Once both approvals are recorded, all 500 payments are released for processing simultaneously. This batch approach reduces approval workload from 1000 individual approvals (500 payments × 2 approvers) to just 2 approvals for the entire batch.

---

**What it does:** System archives metrics with primary key tracking for long-term compliance and performance analysis, enabling historical data retention beyond active database periods

**When it applies:** Metrics older than a configured retention period are moved to archive storage while maintaining queryability for compliance audits and historical analysis

**Who it affects:** Compliance teams conducting historical audits, operations teams analyzing long-term trends, database administrators managing storage capacity

**Example:** A bank configures metrics retention policy: active database keeps 90 days, archive keeps 7 years per regulatory requirements. After 90 days, metrics are automatically moved to archive storage. During a regulatory audit covering the past 3 years, auditors request all API calls that accessed a specific customer account. The system queries both active and archive metrics, returning complete results spanning the full 3-year period. The archive includes all original metrics fields (userId, timestamp, duration, responseBody, etc.) enabling full audit capability even for old data, while the active database stays optimized for recent high-volume queries.

---

**What it does:** Configurable interval controls how often user data is refreshed from external authentication sources, balancing data freshness with system performance

**When it applies:** When users authenticate through external identity providers (LDAP, Active Directory, OAuth providers), the system checks if sufficient time has passed since last refresh before querying the external source

**Who it affects:** Users whose profile data may change in external systems, operations teams managing external API call volumes, security teams ensuring current user data

**Example:** A bank integrates with corporate clients' Active Directory for employee authentication. They configure user_refresh_interval=3600 (1 hour). When an employee logs in at 9:00 AM, the system queries Active Directory for their current user details and caches them. The employee logs out and back in at 9:30 AM. The system checks the last refresh time (9:00 AM), sees only 30 minutes have passed (less than 1 hour interval), and uses cached data without querying Active Directory again. At 10:15 AM, they log in again. Now 75 minutes have passed since last refresh, exceeding the 1-hour interval, so the system queries Active Directory for fresh data. This approach reduces external API calls by 67% while ensuring user data is never more than 1 hour stale.

---

**What it does:** User and consent authentication contexts stored as flexible key-value pairs with consumer tracking, enabling extensible authentication metadata without schema changes

**When it applies:** When authentication flows need to store additional context information beyond standard user credentials, such as device fingerprints, IP addresses, or authentication factors used

**Who it affects:** Security teams implementing adaptive authentication, developers integrating authentication flows, compliance teams auditing authentication methods

**Example:** A bank implements step-up authentication for sensitive operations. When a user logs in normally with password, the system creates UserAuthContext entries: key="auth_method", value="password"; key="device_fingerprint", value="fp123456"; key="ip_address", value="203.0.113.45". Later, the user attempts to transfer a large amount. The system requires step-up authentication via SMS OTP. After successful OTP verification, it adds UserAuthContext: key="step_up_method", value="sms_otp"; key="step_up_timestamp", value="2024-01-15T14:30:00". When reviewing this transaction later, auditors can see the complete authentication context: normal login used password from a known device, then transaction was additionally protected by SMS OTP, creating strong audit trail without modifying the core user authentication schema.

---

**What it does:** System supports bulk creation of transaction attributes for efficiency in batch processing scenarios where multiple attributes need to be added to transactions simultaneously

**When it applies:** When processing large batches of transactions that need to be enriched with multiple attributes, such as nightly batch imports or data migration operations

**Who it affects:** Operations teams running batch processes, integration developers importing transaction data, performance engineers optimizing data loads

**Example:** A bank imports daily transaction data from their core banking system into OBP API each night. The import includes 50,000 transactions, and each transaction needs 5 attributes added (merchant_category, merchant_name, merchant_city, card_type, rewards_points). Using bulk creation, the import process calls createTransactionAttributes once per transaction with all 5 attributes in a list, resulting in 50,000 API calls. Without bulk creation, it would require 250,000 individual calls (50,000 transactions × 5 attributes). The bulk approach reduces API overhead by 80%, completing the nightly import in 15 minutes instead of 75 minutes while maintaining the same data quality.

---


**What it does:** Determines who pays transaction fees using policy: SHARED (split), SENDER (originator pays), RECEIVER (beneficiary pays)

**When it applies:** When processing transaction requests with associated charges

**Who it affects:** Senders paying fees, receivers potentially paying fees, treasury calculating fee revenue

**Example:** €1,000 payment with €10 fee. Policy "SHARED": deduct €1,010 from sender, deduct €10 from receiver, credit €990 to receiver. Policy "SENDER": deduct €1,010 from sender, credit €1,000 to receiver. Policy "RECEIVER": deduct €1,000 from sender, deduct €10 from receiver, credit €990 to receiver. Clear policy enforcement ensures transparent, consistent fee application per agreement.

---

**What it does:** Enforces seven simultaneous limit types on counterparty payments: max/min per transaction, max per day/week/month, max count per week/month

**When it applies:** Every payment to counterparty checked against all seven limit types

**Who it affects:** Customers with counterparty limits, risk management, fraud prevention

**Example:** Utility company counterparty: MAX_LIMIT_PER_TRANSACTION €500, MAX_LIMIT_PER_MONTH €1500, MAX_NUMBER_OF_TRANSACTIONS_PER_MONTH 3. Customer already paid €600 twice this month (€1200 total, 2 payments). Attempts third payment of €300. Transaction amount check passes (€300 < €500). Count check passes (3 ≤ 3). Monthly amount check fails: €1200 + €300 = €1500 equals limit but policy is <, not ≤. Payment blocked. All seven types checked independently.

---

**What it does:** Validates frequency matches predefined enum preventing ambiguous recurrence patterns

**When it applies:** When creating standing orders

**Who it affects:** Customers setting up recurring payments, operations processing standing orders

**Example:** Customer attempts frequency "FORTNIGHTLY". Rejected: must be "DAILY", "WEEKLY", "MONTHLY", "QUARTERLY", "YEARLY". For fortnightly, use "WEEKLY" with when_detail="every 2 weeks". Enum validation prevents processing ambiguity and ensures consistent frequency interpretation across system.

---

**What it does:** Provides flexible scheduling using when_detail field for complex recurrence patterns

**When it applies:** When standing order frequency alone doesn't specify exact schedule

**Who it affects:** Customers needing specific payment schedules, payment processing systems

**Example:** Standing order frequency "MONTHLY", when_detail "15th of each month". System executes on 15th. Another: frequency "WEEKLY", when_detail "every other Friday". System tracks last execution and schedules next Friday two weeks later. when_detail provides natural language schedule detail supplementing frequency enum.

---

**What it does:** Validates active_from < active_to and enforces collection only within date range

**When it applies:** When creating direct debits and during collection execution

**Who it affects:** Customers authorizing direct debits, merchants collecting payments

**Example:** Gym membership: active_from "2024-01-01", active_to "2024-12-31". System validates dates. Monthly collection processes January through December. On January 1, 2025, collection attempt rejected: outside active date range. Gym must request new mandate for 2025. Prevents unauthorized collections beyond agreed period.

---

**What it does:** Tracks which payment networks card can use: Visa, Mastercard, AMEX, etc. (multiple allowed)

**When it applies:** When processing card transactions through payment networks

**Who it affects:** Cardholders using cards, payment networks routing transactions, merchants accepting cards

**Example:** Card has networks ["VISA", "MAESTRO"]. Card accepted at merchants supporting either network. Transaction can route through either. Card with only ["AMEX"] limited to AMEX-accepting merchants. Multi-network support increases acceptance while single-network cards may have lower fees.

---

**What it does:** Specifies card technology: CHIP, CONTACTLESS, CHIP_AND_CONTACTLESS, MAGNETIC_STRIPE

**When it applies:** When determining transaction authorization methods

**Who it affects:** Cardholders using different transaction types, merchants with different terminals, security policies

**Example:** Card technology "CHIP_AND_CONTACTLESS" allows: chip insertion OR contactless tap. Card with only "CHIP" requires insertion, no tap. Card with "MAGNETIC_STRIPE" only for legacy compatibility, higher fraud risk. Terminal checks card technology and presents appropriate interaction method. Security policies may restrict MAGNETIC_STRIPE for new cards.

---

**What it does:** Enforces minimum and maximum cash withdrawal limits per transaction

**When it applies:** When customer attempts cash withdrawal at ATM or branch

**Who it affects:** Customers withdrawing cash, ATMs enforcing limits, fraud prevention

**Example:** Card has minimum_withdrawal "20", maximum_withdrawal "500" (currency from card). Customer attempts €10 withdrawal. Rejected: below minimum. Attempts €600. Rejected: above maximum. Withdraws €300: approved. Minimum prevents excessive small transactions. Maximum limits fraud exposure if card compromised.

---

**What it does:** Stores latitude/longitude GPS coordinates for precise branch location

**When it applies:** When displaying branches on maps, calculating distances, providing directions

**Who it affects:** Customers finding nearby branches, mobile apps showing maps, location-based services

**Example:** Branch has location: latitude "51.5074", longitude "-0.1278" (London). Mobile app uses device GPS, calculates distances to all branches, sorts by proximity. Shows: "0.3 miles away". Tapping branch opens map with pin at exact coordinates. GPS coordinates enable precise mapping beyond street addresses.

---

**What it does:** Tracks ATM site name and location details beyond address for identification

**When it applies:** When customers search for specific ATMs or report issues

**Who it affects:** Customers finding ATMs, operations managing ATMs, support handling reports

**Example:** ATM has name "Main Street Branch Lobby", site_name "MainSt", location_name "Branch Entrance". Customer reports: "ATM at Main Street not working". Support searches by site_name "MainSt", identifies specific ATM, dispatches technician. Multiple ATMs at one location (lobby, drive-up) distinguished by name and location_name fields. Enables precise ATM identification.

---

**What it does:** Exchange rates retrieved via three-tier workflow: (1) Connector from core banking system, (2) Fallback JSON resource files with reference rates, (3) Hard-coded map for common currency pairs, ensuring rates are always available

**When it applies:** Whenever currency conversion is needed, the system attempts each lookup method in sequence until a rate is found, providing resilience against external system failures

**Who it affects:** Treasury teams managing multi-currency operations, customers making cross-border payments, operations teams ensuring service availability

**Example:** A customer initiates a GBP to USD payment. The system first queries the connector which calls the bank's real-time FX pricing system (tier 1). This returns GBP/USD=1.2650 with a 2-minute cache TTL. The payment is processed at this rate. An hour later, the connector's FX system goes down for emergency maintenance. Another customer initiates a EUR to USD payment. The connector fails, so the system checks tier 2 fallback JSON files containing yesterday's closing rates and finds EUR/USD=1.0850. The payment proceeds with this rate. Later, someone attempts an unusual currency pair (ISK to THB) not in the connector or JSON files. The system checks tier 3 hard-coded map, which doesn't have this direct pair but has USD rates for both. It calculates the cross-rate via USD and completes the conversion. This three-tier approach provides 99.9%+ availability for currency conversion.

---

**What it does:** System automatically returns 1.0 exchange rate for same currency conversions without external lookup, optimizing performance and ensuring consistency

**When it applies:** When source and target currencies are identical, the system skips all external rate lookups and immediately returns 1.0

**Who it affects:** Developers integrating currency conversion, operations teams optimizing performance, customers transferring within same currency

**Example:** A customer transfers $500 from their USD checking account to their USD savings account at the same bank. The transfer processing logic requests currency conversion from USD to USD. The system immediately recognizes source_currency == target_currency and returns exchange rate 1.0 without querying the connector, JSON files, or hard-coded map. The transfer completes instantly with amount_converted=500.00. This optimization prevents unnecessary external system calls for same-currency operations, which represent approximately 70% of all transfer volume in typical multi-currency banks, significantly reducing load on FX pricing systems.

---

**What it does:** If direct exchange rate not found, system attempts inverse rate lookup and calculates reciprocal, maximizing rate availability from limited data sets

**When it applies:** When requesting a currency pair (e.g., EUR/GBP) that isn't directly available, the system checks if the inverse pair (GBP/EUR) exists and calculates 1/inverse_rate

**Who it affects:** Treasury teams managing currency data, customers converting between less common currency pairs, operations teams minimizing data maintenance

**Example:** A bank's FX rate file contains GBP/EUR=1.1650 but not the inverse EUR/GBP. A customer requests a EUR to GBP conversion. The system first looks for EUR/GBP (not found), then looks for GBP/EUR (found=1.1650), calculates inverse 1/1.1650=0.8583, and returns EUR/GBP=0.8583. This bidirectional lookup effectively doubles the available currency pairs from the same data set. A file with 50 currency pairs provides 100 conversion paths, reducing data maintenance effort by 50% while maintaining full conversion capability.

---

**What it does:** All currency conversions rounded to 2 decimal places using HALF_UP rounding mode for consistency with standard financial practices and preventing micro-amount discrepancies

**When it applies:** After applying exchange rate to convert an amount between currencies, the system always rounds the result to exactly 2 decimal places

**Who it affects:** Customers receiving converted amounts, treasury teams reconciling conversions, compliance teams ensuring calculation consistency

**Example:** A customer converts 1000.00 EUR to USD at rate 1.08355. Raw calculation: 1000.00 × 1.08355 = 1083.55000. The system applies HALF_UP rounding to 2 decimals: 1083.55 USD. Another customer converts 333.33 EUR at the same rate. Raw calculation: 333.33 × 1.08355 = 361.18331. HALF_UP rounding: 361.18 USD (not 361.19 or 361.183). This consistent rounding prevents accumulation of micro-amounts that could cause reconciliation issues and ensures all monetary amounts comply with standard currency precision of 2 decimal places, matching physical currency denominations and accounting standards.

---


**What it does:** Provides 100+ independent boolean permission flags per view for fine-grained access control

**When it applies:** When configuring view permissions for accounts

**Who it affects:** View administrators, users with different access levels, security policy enforcement

**Example:** View has: canSeeTransactionAmount=true, canSeeOtherAccountNumber=false, canAddTransactionRequest=true, canCreateStandingOrder=false, canSeeOwnerComment=true (100+ more flags). User with this view sees transaction amounts and can initiate payments, but cannot see recipient account numbers or create standing orders. Each permission independently controllable enables precise access policies.

---

**What it does:** Controls transaction field visibility (amount, other account, description, etc.) independently per view

**When it applies:** When users access transaction information through different views

**Who it affects:** Users seeing filtered transaction data, privacy enforcement, view designers

**Example:** "accountant" view: canSeeTransactionAmount=true, canSeeTransactionDescription=true, canSeeOtherAccountNumber=false. Accountant sees amounts and descriptions for bookkeeping but not account numbers (privacy). "auditor" view: all fields visible. "public" view: only dates visible, no amounts or details. Transaction visibility precisely controlled per view requirements.

---

**What it does:** Separates permissions for viewing vs managing metadata (tags, comments, images) per view

**When it applies:** When users need to see metadata but not modify it, or vice versa

**Who it affects:** View users managing metadata, read-only users, metadata governance

**Example:** View permissions: canSeeTags=true, canAddTag=false, canDeleteTag=false, canSeeComments=true, canAddComment=true, canDeleteComment=false. User can see all tags and comments, add new comments, but cannot add/delete tags or delete comments. Enables contribution while preventing removal of others' metadata. Precise metadata permission control.

---

**What it does:** Controls counterparty viewing and creation permissions independently per view

**When it applies:** When managing who can see and create counterparties (payment recipients)

**Who it affects:** Users managing counterparties, payment security, counterparty data governance

**Example:** View: canSeeAvailableViews=true (can see counterparty list), canAddCounterparty=false (cannot create new). User can pay existing counterparties but cannot add new recipients without approval. Another view: both permissions true for power users. Prevents unauthorized recipient creation while allowing legitimate payments.

---

**What it does:** Controls which transaction request types user can initiate per view

**When it applies:** When restricting payment types based on user authorization level

**Who it affects:** Users initiating payments, transaction type security, payment fraud prevention

**Example:** View allows: canAddTransactionRequestToAnyAccount=false (only own accounts), but canAddTransactionRequestType specific types: "INTERNAL_TRANSFER"=true, "SEPA"=true, "SWIFT"=false. User can initiate internal and SEPA transfers but not SWIFT (requires approval). Different views enable different payment types per authorization level.

---

**What it does:** Controls recurring payment setup permissions separately for standing orders and direct debits

**When it applies:** When configuring who can set up automated recurring payments

**Who it affects:** Users setting up recurring payments, operations managing automated payments

**Example:** View: canCreateStandingOrder=true, canCreateDirectDebit=false. User can set up own standing orders (push payments initiated by user) but not direct debits (pull payments initiated by merchant). Reflects different trust levels and risk profiles of payment types.

---

**What it does:** Controls who can create, update, or delete custom views

**When it applies:** When managing view administration permissions

**Who it affects:** View administrators, account owners delegating access, view lifecycle management

**Example:** View: canCreateCustomView=true, canUpdateCustomView=true, canDeleteCustomView=false. Account owner can create and update custom views for delegates but cannot delete system views. Administrator view: all permissions true. Prevents accidental deletion while enabling view customization.

---

**What it does:** Controls who can grant or revoke view access to other users

**When it applies:** When managing view delegation and access control

**Who it affects:** Account owners granting access, administrators managing permissions, delegates receiving access

**Example:** View: canGrantAccessToView=true, canRevokeAccessToView=true for account owner. Owner can give accountant access to "accountant" view and later revoke it. Standard user view: both permissions false. Cannot delegate access. Enables controlled delegation while preventing unauthorized access grants.

---

**What it does:** Special "firehose" permission enables seeing all transactions across all accounts for system monitoring

**When it applies:** When system monitoring or fraud detection needs comprehensive transaction visibility

**Who it affects:** Fraud detection systems, system monitoring, compliance oversight, security operations

**Example:** System creates special view with canSeeTransactionFirehose=true for fraud detection service. This view sees ALL transactions across ALL accounts in real-time for pattern analysis. Regular views see only specific accounts. Firehose view enables system-wide monitoring without granting account-specific access. Strict access control on firehose views prevents abuse.

---

**What it does:** Distinguishes between system-defined views (owner, public, accountant) and bank-custom views

**When it applies:** When managing view lifecycle and determining which views can be modified

**Who it affects:** View administrators, system upgrades, custom view creators

**Example:** System views: "owner", "public", "accountant" have is_system=true. Cannot be deleted, core permissions locked (can add but not remove). Custom view "regional_manager" has is_system=false. Can be fully modified or deleted. During system upgrade, system views updated automatically. Custom views preserved unchanged. Type distinction enables safe system evolution.

---

**What it does:** Entitlements can be bank-specific or global (all banks) using bank_id field

**When it applies:** When granting permissions that should apply to specific bank or across all banks

**Who it affects:** Multi-bank administrators, bank-specific staff, global administrators

**Example:** User granted "CanCreateCustomer" with bank_id="bank.uk.001" (specific bank). Can create customers only at this bank. Another user granted same entitlement with bank_id="" (empty/null = global). Can create customers at any bank. Global entitlements for platform administrators, bank-specific for bank staff. Scoping prevents cross-bank privilege escalation.

---

**What it does:** Records timestamp every time entitlement is actually used (not just granted)

**When it applies:** Whenever user exercises a permission they hold

**Who it affects:** Audit teams tracking permission usage, security investigating incidents, usage analytics

**Example:** User granted "CanCreateTransaction" on January 1st. Permission used: January 5th (first use), January 7th, January 12th. Audit shows: permission granted but not used for 4 days (possible forgotten grant). Another permission granted January 1st, never used (over-provisioned?). Usage timestamps enable: identifying unused permissions for cleanup, proving permission was actually exercised at specific times for investigations, measuring actual permission utilization.

---

**What it does:** Prevents duplicate entitlements: unique combination of user_id + role_name + bank_id

**When it applies:** When granting entitlements

**Who it affects:** Administrators granting permissions, entitlement data integrity

**Example:** User already has "CanCreateCustomer" for "bank.001". Administrator attempts to grant same role for same bank again. Rejected: "Entitlement already exists". Prevents duplicate grants that would confuse audit trails and complicate revocation. Each user has each role at most once per bank scope. Ensures clean entitlement data.

---

**What it does:** Users identified by (provider, provider_id) tuple enabling multi-tenancy with multiple identity providers

**When it applies:** When users authenticate from different identity sources (OAuth, LDAP, local, etc.)

**Who it affects:** Multi-tenant deployments, users from different identity providers, SSO integrations

**Example:** System has users: ("local", "alice123"), ("ldap_corpA", "alice.smith"), ("oauth_google", "alice@gmail.com"). All different users despite similar names. Authentication checks provider + provider_id. User logs in via Google OAuth. System looks up ("oauth_google", google_id). User from Corp A logs in via LDAP. System looks up ("ldap_corpA", ldap_id). Multi-provider support enables diverse authentication without conflicts.

---

**What it does:** Controls whether view is publicly visible or private using is_public flag

**When it applies:** When determining view discoverability and access patterns

**Who it affects:** View visibility in listings, public data access, privacy enforcement

**Example:** Account has views: "owner" (is_public=false), "public" (is_public=true), "accountant_alice" (is_public=false). Public API query lists only "public" view. Accessing "owner" or "accountant_alice" requires direct authorization. Public view enables open data sharing while private views enforce access control. Public views discoverable, private views require explicit grant.

---

**What it does:** Configures what information is shown as "alias" when true account number is hidden

**When it applies:** When views hide account numbers but need to display alternative identifier

**Who it affects:** Privacy enforcement, users seeing account aliases instead of numbers

**Example:** View hides account number (canSeeAccountNumber=false) but shows label (canSeeLabel=true). View configuration: use_alias="label". User sees account label "Travel Fund" instead of "GB29NWBK60161331926819". Another view: use_alias="public_alias", shows bank-generated public identifier. Alias policy balances privacy with usability.

---

**What it does:** Links views to metadata views enabling metadata visibility delegation

**When it applies:** When transaction metadata needs different access control than transaction itself

**Who it affects:** Metadata governance, delegated metadata access, privacy separation

**Example:** Transaction view allows seeing amounts but not adding comments. Metadata_view_id references separate "metadata_viewer" view with canSeeComments=true, canAddComment=false. User can see amounts (transaction view) AND see comments (metadata view) but not add comments. Separate metadata view reference enables independent metadata permission management.

---

**What it does:** Email lookup returns array of users because email may exist for multiple providers

**When it applies:** When searching users by email in multi-provider environment

**Who it affects:** Support staff finding users, systems integrating by email, multi-provider scenarios

**Example:** Email "alice@example.com" exists for: ("local", "alice123"), ("ldap_corp", "alice.smith"), ("oauth_google", google_id). Query getUsersByEmail("alice@example.com") returns array with all 3 users. Caller must disambiguate using provider context. Single user lookup would fail; multi-result array accommodates multi-provider reality.

---

**What it does:** Enables bulk user deletion (multiple users in one call) for test environment cleanup

**When it applies:** When cleaning test environments between test runs

**Who it affects:** Test automation, environment management, test data cleanup

**Example:** Integration test creates 100 test users. After test, calls bulkDeleteUsers([user_id1, user_id2, ..., user_id100]). All users deleted in one transaction. Alternative: 100 individual delete calls. Bulk operation reduces cleanup time from minutes to seconds. Typically restricted to non-production environments for safety.

---

**What it does:** Enables atomic view permission updates by deleting all existing permissions and creating new set

**When it applies:** When updating view permission templates or migrating permissions

**Who it affects:** View administrators, system migration scripts, view lifecycle management

**Example:** View "customer_service_view" needs permission updates. Administrator calls setFromViewData() or createViewAndPermissions() with new ViewSpecification containing updated permission set. Method ViewPermission.resetViewPermissions() deletes all existing permissions for view and creates new ones according to specification. This enables: (1) Atomic permission updates (all old permissions removed, all new permissions added), (2) View permission migrations during system upgrades, (3) Template-based view creation with predefined permission sets. Ensures consistent permission state without manual manipulation of individual permission flags.

---

**What it does:** Every user lock operation records lastLockDate timestamp for audit trail and lock age monitoring, enabling time-based analysis of security incidents

**When it applies:** When a user account is locked (either automatically via bad login attempts or manually via API), the system records or updates the lastLockDate field with the current timestamp

**Who it affects:** Security teams investigating account locks, compliance officers auditing security controls, operations teams monitoring lock patterns

**Example:** A security analyst investigates a series of account locks to determine if they're related to a coordinated attack. They query all UserLocks records and sort by lastLockDate. They discover 50 accounts were locked within a 2-minute window on March 15th at 2:30 AM, all with similar usernames following a pattern. This temporal clustering indicates an automated attack targeting specific accounts. The analyst identifies the attack source and implements IP blocking. Without timestamp tracking, these locks would appear as isolated incidents rather than revealing the coordinated attack pattern.

---

**What it does:** Unlocking user removes lock record entirely rather than updating status flag, providing clean state management and simplifying queries for currently locked users

**When it applies:** When an administrator or automated process unlocks a user account, the system deletes the UserLock record rather than setting a status field to "unlocked"

**Who it affects:** Administrators managing user access, developers querying lock status, database administrators managing data volume

**Example:** A user's account was locked due to bad login attempts. After verifying their identity via phone, a customer service representative calls unlockUser(provider="LDAP", username="jsmith"). The system finds the UserLock record for user ID "12345" and deletes it entirely. Subsequently, when the user tries to log in, the isLocked() function queries for UserLock records matching their user ID and finds none, allowing authentication to proceed. If the user's account is locked again later, a fresh UserLock record is created. This delete-based approach means "locked users" can always be found with a simple query "SELECT * FROM UserLocks" rather than needing "WHERE status='locked'" filters, and historical lock data doesn't accumulate in the active table.

---

**What it does:** User authentication context tracks which consumer (API client) created the context for accountability, enabling per-application authentication context management

**When it applies:** When creating user authentication context entries, the system records the consumerId of the API application that created the context, linking authentication metadata to specific applications

**Who it affects:** Security teams tracking authentication per application, developers managing app-specific authentication flows, compliance officers auditing authentication methods

**Example:** A bank has two fintech partners integrated via API: App A (mobile banking) and App B (investment platform). Both use the same user account but have different authentication requirements. When a user authenticates to App A using fingerprint biometrics, App A creates UserAuthContext entries with consumerId="consumer_app_a": key="biometric_type", value="fingerprint"; key="device_id", value="device_123". When the same user later authenticates to App B using SMS OTP, App B creates separate contexts with consumerId="consumer_app_b": key="otp_method", value="sms"; key="phone_verified", value="true". When investigating a security incident in App A, analysts can query contexts for consumer_app_a only, seeing fingerprint authentication history without the noise of App B's SMS OTP records. This consumer scoping enables precise audit trails per application.

---

**What it does:** Consent authentication contexts stored as flexible key-value pairs similar to user contexts, enabling extensible consent metadata tracking without schema modifications

**When it applies:** When managing consent lifecycle and authentication methods used to grant consent, the system stores arbitrary context as key-value pairs associated with specific consent IDs

**Who it affects:** Compliance teams tracking consent authenticity, security teams verifying consent authorization methods, audit teams reviewing consent history

**Example:** A customer grants account access consent to a third-party PFM (Personal Finance Management) app. During the consent flow, the system creates ConsentAuthContext records: consentId="consent_789", key="auth_method", value="redirect_oauth"; key="scopes_requested", value="accounts,transactions"; key="consent_screen_shown", value="true"; key="ip_address", value="198.51.100.42". Later, the customer claims they never granted this consent. The compliance team retrieves the consent's authentication context and sees the complete audit trail: consent was granted via OAuth redirect flow, customer viewed consent screen, specific scopes were requested and shown, and originating IP address. This detailed context proves the consent was properly authorized, protecting the bank from liability.

---

**What it does:** System supports createOrUpdate batch operations for efficient context management, allowing multiple context entries to be created or updated in a single API call

**When it applies:** When authentication flows need to record multiple context fields simultaneously, batch operations reduce API overhead and ensure atomic updates

**Who it affects:** Developers implementing authentication flows, operations teams optimizing API performance, security teams ensuring complete context capture

**Example:** A bank's authentication service implements a risk-based authentication flow that evaluates 8 different factors (device fingerprint, IP reputation, velocity checks, location, time of day, transaction amount, customer behavior pattern, blacklist status). Rather than making 8 separate createUserAuthContext calls, the service builds a list of BasicUserAuthContext objects with all 8 key-value pairs and calls createOrUpdateUserAuthContexts once. The system processes all 8 contexts in a single transaction, ensuring either all are recorded or none are (atomic operation). This reduces authentication flow latency from ~800ms (8 calls × 100ms each) to ~150ms (1 batch call), improving user experience while ensuring complete context capture for security analysis.

---

**What it does:** Transactions can be queried by attribute name-value combinations for flexible filtering, enabling sophisticated transaction searches beyond standard fields

**When it applies:** When users need to find transactions based on custom attributes like merchant categories, project codes, or other metadata that varies by business needs

**Who it affects:** Corporate customers tracking project expenses, compliance teams investigating specific transaction types, developers building custom reporting tools

**Example:** A construction company tags all project expenses with transaction attributes: "project_id", "vendor_category", and "approval_status". Their accounting team needs to find all pending lumber purchases for Project Alpha. They query: getTransactionIdsByAttributeNameValues(bankId="bank123", params={"project_id":["PROJ_ALPHA"], "vendor_category":["lumber"], "approval_status":["pending"]}). The system returns transaction IDs matching ALL three criteria. The team reviews these specific transactions and approves them, updating the "approval_status" attribute to "approved". This flexible query capability enables custom workflows and reporting without modifying the core transaction schema.

---

**What it does:** System optionally captures full response bodies in metrics for detailed audit and debugging, providing complete API interaction history when enabled

**When it applies:** When detailed audit requirements or debugging needs require preserving the actual data returned by API calls, metrics can be configured to capture response bodies

**Who it affects:** Compliance teams conducting detailed audits, developers debugging production issues, security teams investigating data exfiltration, privacy teams managing sensitive data

**Example:** During a regulatory audit, examiners ask for evidence showing what customer data was returned when a specific third-party application accessed accounts during Q1 2024. The bank has responseBody capture enabled for production environment. They query API metrics for that consumer ID and date range, retrieving the actual JSON response bodies from the metrics archive. The audit reveals that the application only received basic account balances (as consented) and never accessed transaction details, proving compliance with consent scope. However, this comprehensive logging also creates privacy considerations, so the bank implements metrics retention policies ensuring responseBody data is purged after the regulatory retention period to balance audit capability with privacy requirements.

---



The Open Bank Project API implements a sophisticated, multi-layered approach to banking operations with **123 comprehensive business rules** covering:

**Regulatory Compliance (15 rules):** Customer data segregation for regulatory updates, consent lifecycle state machines, entitlement audit trails, KYC check documentation, transaction request metadata tracking, multi-level rate limiting, consent time-to-live expiry, tax residence tracking, account application workflows, user invitation lifecycle, GDPR-compliant data scrambling, web UI property management, JSON schema validation per operation, authentication type restrictions, comprehensive API metrics

**Customer Service (20 rules):** View-based granular access control with 100+ permission flags, transaction request type validation, seven-type counterparty limits, automatic charge calculation, customer message routing, standing order frequency validation, direct debit expiry enforcement, account webhook notifications, physical card security states, PIN reset tracking, card replacement management, card collection/posting tracking, branch operating hours, ATM deposit capabilities, ATM multi-currency support, transaction attribute view-based visibility, CRM event tracking (scheduled vs actual dates), CRM event categorization, user agreement acceptance, transaction type fees

**Risk Management (17 rules):** Challenge-response authentication with BCrypt, physical card action type restrictions, CVV secure hash storage, multi-period rate limiting with independent windows, standing order date range validation, direct debit customer/user association, method routing pattern matching, dynamic entity reference validation, account attribute view-based visibility, ATM fee structure transparency, branch soft delete with historical retention, meeting multi-party participation, bad login threshold enforcement, bad login reset on success, manual user locks via API, user lock dual-check system, signing basket status lifecycle

**Product Management (15 rules):** Product attribute type-safe value management, account attribute inheritance from product templates, product collection bundle management, product fee frequency and active status, product license-based visibility with admin override, dynamic entity automatic CRUD API generation, product attribute active/inactive status toggle, product collection tree structure with full details, bank attribute type-safe configuration, endpoint mapping request/response transformation, endpoint mapping bank-specific vs global configuration, transaction type bank-specific short codes, card attribute type-safe management, card attribute scoping, transaction attribute bank scoping

**Operations (15 rules):** User invitation secret link uniqueness, account webhook HTTP protocol support, API collection endpoint composite unique constraints, branch accessibility feature tracking, ATM supported languages, ATM supported services capability tracking, branch drive-up hours separate from lobby hours, API property backward compatibility with alias support, user creation with consent/invitation tracking, user multi-lookup methods for flexible access, signing basket multi-entity support, API metrics archive functionality, user refresh interval control, authentication context key-value storage, transaction attribute bulk creation

**Treasury & Payment (14 rules):** Transaction request charge policy three-way split, counterparty limit seven independent type enforcement, standing order frequency enum validation, standing order when detail schedule specification, direct debit date range validation, physical card network multi-network support, physical card technology type specification, card minimum and maximum withdrawal amounts, branch location GPS coordinates tracking, ATM location and site identification, FX three-tier exchange rate lookup, FX same currency optimization, FX bidirectional rate lookup, FX currency conversion rounding

**Security & Access Control (27 rules):** View permission 100+ granular boolean flags, view-based transaction visibility granularity, view-based metadata management permissions separation, view-based counterparty management permissions, view-based transaction request type permissions, view-based standing order and direct debit creation permissions, view-based custom view management permissions, view permission grant and revoke capabilities, view firehose capability for system monitoring, view system vs custom type distinction, role-based entitlement bank scoping with global override, entitlement usage timestamp tracking, entitlement composite unique constraint, user provider-based multi-tenant authentication, view public vs private visibility control, view alias selection policy configuration, view metadata view reference for delegation, user email-based lookup with multiple results, user bulk delete for testing environments, view permission reset and bulk update, user lock timestamp tracking, user lock delete-based unlock, consumer-scoped user auth context, consent authentication context storage, authentication context batch operations, transaction attribute query by name-value pairs, API metrics response body capture

The system successfully balances **security, compliance, flexibility, and user experience** through these well-defined business rules, supporting traditional banking operations while enabling innovative fintech applications across multiple API standards (OBP, Berlin Group, UK Open Banking).
