# Business Rules Extraction Prompt for Open Bank Project API

## Purpose
This prompt will help you understand and document the business rules that govern how the banking system works. Think of business rules as the "policies and procedures" that the system follows when handling banking operations.

---

## Role-Based Questions to Extract Business Rules

### Role 1: As a Bank Compliance Officer
**Your goal:** Understand what rules ensure the bank follows regulations and protects customers

Ask yourself these questions when reviewing the code:
- What information must we collect before allowing someone to become a customer?
- What verification steps are required before we can approve a customer?
- Who needs to review and approve sensitive customer information changes?
- What documentation must we keep for regulatory purposes?
- Are there limits on who can access customer data and when?
- What approvals are needed before processing large transactions?
- How does the system ensure we meet Know Your Customer requirements?

### Role 2: As a Customer Service Manager  
**Your goal:** Understand what rules govern how customers interact with the bank

Ask yourself these questions when reviewing the code:
- What can customers do on their own versus what requires bank staff assistance?
- What restrictions exist on customer transactions?
- When are customers required to provide additional verification?
- What information can customers see about their own accounts?
- How does the system handle customer address or contact information updates?
- What rules determine if a customer can link multiple accounts?
- Are there different service levels or permissions for different customer types?

### Role 3: As a Risk Management Specialist
**Your goal:** Understand what rules protect the bank from fraud and financial risk

Ask yourself these questions when reviewing the code:
- What limits are placed on transaction amounts?
- How does the system determine if a transaction is suspicious?
- What approvals are needed for different types of transactions?
- Are there daily or monthly limits on customer activities?
- How does the system track and prevent unauthorized access?
- What happens when someone tries to perform an action they're not allowed to do?
- How are transaction fees and charges determined and applied?

### Role 4: As a Product Manager
**Your goal:** Understand what rules define different banking products and services

Ask yourself these questions when reviewing the code:
- What types of accounts can be created?
- What features and capabilities does each account type have?
- How are account fees and charges structured?
- What products can be bundled together?
- Are there eligibility requirements for certain products?
- How do different account types differ in their permissions and limits?
- What services are available through different channels (online, branch, mobile)?

### Role 5: As an Operations Director
**Your goal:** Understand what rules govern internal bank processes and workflows

Ask yourself these questions when reviewing the code:
- What approval chains exist for different operations?
- Which staff members can perform which actions?
- How are responsibilities divided between different bank roles?
- What audit trails and records must be maintained?
- How does information flow between different departments?
- What automated processes run and when?
- How are errors and exceptions handled?

### Role 6: As a Treasury and Payment Specialist
**Your goal:** Understand what rules control money movement and payment processing

Ask yourself these questions when reviewing the code:
- What types of payments and transfers are supported?
- How are payment fees calculated and who pays them (sender, receiver, or shared)?
- What validations occur before a payment is processed?
- What statuses can a transaction have and when do they change?
- Are there different processing rules for domestic versus international payments?
- How does the system handle payment failures or reversals?
- What time-based rules affect payment processing?

### Role 7: As a Security and Access Control Manager
**Your goal:** Understand what rules protect the system and control access

Ask yourself these questions when reviewing the code:
- Who can create, view, update, or delete different types of information?
- Are there different permission levels for the same action at different banks?
- How are user privileges granted and revoked?
- What authentication methods are required for different operations?
- How does the system prevent unauthorized data access?
- What logging and monitoring rules are in place?
- How are sensitive operations challenged or verified?

---

## How to Document What You Find

For each business rule you discover, write it down in simple language using this format:

**Rule Name:** [A short, descriptive name]

**What it does:** [Explain the rule in one sentence]

**When it applies:** [Describe the situation or condition when this rule is used]

**Who it affects:** [Which people or parts of the system this impacts]

**Example:** [Give a real-world example of this rule in action]

---

## Example of a Well-Documented Business Rule

**Rule Name:** Customer Credit Limit Update Approval

**What it does:** Only authorized bank staff can change a customer's credit limit

**When it applies:** Whenever someone tries to modify the credit limit field for any customer

**Who it affects:** Bank staff members, customers whose credit limits are being changed

**Example:** If a customer service representative wants to increase a customer's credit limit from $5,000 to $10,000, they must have the "CanUpdateCustomerCreditLimit" permission for that specific bank. Without this permission, the system will reject their request.

---

## Tips for Extracting Business Rules

1. **Look for permission checks** - These tell you who can do what
2. **Look for validation steps** - These tell you what conditions must be met
3. **Look for status changes** - These tell you the lifecycle of operations
4. **Look for calculations** - These tell you how amounts and fees are determined
5. **Look for limits and thresholds** - These tell you the boundaries of operations
6. **Look for approval workflows** - These tell you the process for important decisions
7. **Look for different treatment** - These tell you when rules vary based on conditions

---

## Categories of Business Rules to Look For

### Access and Permission Rules
- Who can perform which actions
- What information different roles can see
- How permissions are granted or restricted

### Validation and Verification Rules
- What information must be provided
- What formats and values are acceptable
- What checks must pass before proceeding

### Processing and Workflow Rules  
- What steps must occur in what order
- What approvals are needed
- How different statuses flow

### Financial and Calculation Rules
- How fees and charges are determined
- What limits apply to amounts
- How costs are shared or allocated

### Compliance and Audit Rules
- What must be recorded for regulatory purposes
- What documentation is required
- What retention periods apply

### Customer and Account Rules
- What types of customers and accounts exist
- What features each type has
- How they can be created, modified, or closed

### Transaction and Payment Rules
- What types of transactions are allowed
- What validations apply to payments
- How transaction processing works

### Security and Authentication Rules
- How users prove their identity
- What additional verification is needed
- How sensitive operations are protected
