# Validation Rules Extraction Prompt

## Purpose
This guide helps you understand and document all the rules that check whether banking information is correct and acceptable in the Open Bank Project API system.

---

## For Business Analysts

### Your Role
You need to understand what business rules are being checked when someone uses the banking system.

### What to Look For
- **Account Requirements**: What information must be provided when creating or accessing bank accounts?
- **Transaction Limits**: What are the minimum and maximum amounts allowed for different types of transactions?
- **Customer Information**: What details about customers must be collected and what format should they be in?
- **Permission Rules**: Who is allowed to see or do what in the system?
- **Time Restrictions**: Are there specific times when certain operations can or cannot be performed?
- **Geographic Restrictions**: Are there rules about which countries or regions can access certain features?

### Questions to Answer
1. What makes a valid bank account number in this system?
2. What personal information is required versus optional for customers?
3. What are the rules around money transfers between accounts?
4. What conditions must be met before someone can perform sensitive operations?
5. What business scenarios are prevented by the validation rules?

---

## For Compliance Officers

### Your Role
You need to ensure all regulatory and legal requirements are being enforced through validation rules.

### What to Look For
- **Identity Verification**: What checks are in place to confirm customer identities?
- **Anti-Money Laundering**: What rules prevent suspicious transactions or patterns?
- **Data Privacy**: How is sensitive customer information protected and validated?
- **Regulatory Requirements**: Which validations enforce banking regulations and standards?
- **Audit Trail Requirements**: What information must be captured and validated for compliance reporting?
- **Age Restrictions**: Are there age-related validations for certain banking products or services?

### Questions to Answer
1. Which validation rules enforce Know Your Customer requirements?
2. How does the system ensure transactions comply with local banking regulations?
3. What validations protect customer privacy and data security?
4. Which rules help detect and prevent fraudulent activities?
5. How are regulatory limits and thresholds enforced in the system?

---

## For Quality Assurance Teams

### Your Role
You need to understand what gets checked so you can test whether the system correctly accepts good data and rejects bad data.

### What to Look For
- **Input Validation**: What formats and patterns are expected for different data fields?
- **Boundary Conditions**: What are the minimum and maximum acceptable values?
- **Required Fields**: Which information must always be provided?
- **Relationship Rules**: How do different pieces of data relate to and validate against each other?
- **Error Messages**: What feedback is given when validation fails?
- **Edge Cases**: What unusual but valid scenarios are handled?

### Questions to Answer
1. What happens when someone provides invalid account information?
2. What are all the ways a transaction can be rejected?
3. Which combinations of data are not allowed together?
4. What error messages guide users when they make mistakes?
5. How does the system handle unusual but legitimate cases?

---

## For Product Managers

### Your Role
You need to understand what limitations and requirements exist in the current system to plan features and improvements.

### What to Look For
- **User Experience Impact**: How do validation rules affect the customer journey?
- **Feature Limitations**: What can and cannot be done due to current validation rules?
- **Business Logic**: What business decisions are embedded in the validation rules?
- **Customer Friction Points**: Where might validations create obstacles for users?
- **Competitive Constraints**: How do these rules compare to what competitors allow?
- **Innovation Opportunities**: What rules might need to change for new features?

### Questions to Answer
1. What customer requests are prevented by current validation rules?
2. Which validations create the most user complaints or support tickets?
3. What new banking products or features are limited by existing rules?
4. How do validation rules align with customer expectations and needs?
5. What rules should be relaxed or tightened based on business goals?

---

## For Customer Support Teams

### Your Role
You need to understand validation rules to help customers when they encounter errors or restrictions.

### What to Look For
- **Common Error Scenarios**: What mistakes do customers frequently make?
- **Helpful Error Messages**: What information is provided when something goes wrong?
- **Resolution Paths**: How can customers fix validation errors?
- **System Limitations**: What requests are impossible due to validation rules?
- **Workarounds**: Are there alternative approaches when direct methods fail?
- **Clear Explanations**: How can technical rules be explained in customer-friendly language?

### Questions to Answer
1. What are the most common reasons customer actions get rejected?
2. How can you help customers provide information in the correct format?
3. What validation errors cannot be resolved by customers themselves?
4. Which rules require special permissions or administrator intervention?
5. How do you explain technical validation failures in simple terms?

---

## For System Administrators

### Your Role
You need to understand validation rules to configure the system correctly and troubleshoot issues.

### What to Look For
- **Configuration Options**: Which validation rules can be customized or adjusted?
- **System Settings**: What parameters control validation behavior?
- **Integration Points**: How do validations affect connections with other systems?
- **Performance Impact**: Do any validations slow down system operations?
- **Override Capabilities**: When and how can validation rules be bypassed for special cases?
- **Monitoring Requirements**: What validation failures should trigger alerts?

### Questions to Answer
1. Which validation rules can be modified through configuration?
2. How do validation settings affect system performance?
3. What validation failures indicate system problems versus user errors?
4. When is it appropriate to override or bypass validation rules?
5. How should validation-related errors be logged and monitored?

---

## For Security Teams

### Your Role
You need to ensure validation rules protect the system from malicious activities and unauthorized access.

### What to Look For
- **Authentication Checks**: What validates user identities and credentials?
- **Authorization Rules**: How is permission to perform actions verified?
- **Input Sanitization**: What protections exist against malicious input?
- **Rate Limiting**: Are there rules preventing abuse through excessive requests?
- **Data Integrity**: How is the accuracy and consistency of data ensured?
- **Security Boundaries**: What validations enforce security policies?

### Questions to Answer
1. Which validation rules prevent unauthorized access to accounts or data?
2. How does the system protect against injection attacks through validation?
3. What rules detect and prevent suspicious or fraudulent patterns?
4. How are security credentials validated before granting access?
5. What validations ensure data cannot be tampered with?

---

## General Extraction Guidelines

### Document Structure
For each validation rule you find, capture:
1. **What it checks**: A plain English description of what is being validated
2. **Why it exists**: The business or technical reason for this rule
3. **When it applies**: Under what circumstances this validation is performed
4. **Who it affects**: Which users or roles are impacted by this rule
5. **What happens when it fails**: The outcome or error when validation is not met
6. **Where it is enforced**: At what point in the user journey or system process

### Organization Approach
Group validation rules by:
- **Type of Operation**: Account creation, transactions, user management, etc.
- **Data Category**: Personal information, financial data, system settings, etc.
- **Severity Level**: Critical rules versus helpful warnings
- **User Impact**: Customer-facing versus internal system validations
- **Regulatory Category**: Compliance requirements versus business preferences

### Documentation Format
Present each rule in a way that:
- Uses everyday language, not technical terms
- Explains the purpose before the details
- Provides examples of both valid and invalid scenarios
- Clarifies any exceptions or special cases
- Links related rules that work together

---

## Key Areas to Investigate

### Account and Customer Information
- What information is required to open a bank account?
- How are names, addresses, and contact details validated?
- What identification documents are needed and how are they verified?
- What are the rules for different types of accounts (savings, checking, business)?

### Financial Transactions
- What limits exist on transaction amounts?
- How are payment details verified before processing?
- What rules govern transfers between different account types?
- When can transactions be canceled or reversed?

### Authentication and Access
- What makes a strong password or credential?
- How are user permissions and roles validated?
- What rules control who can see what information?
- How is multi-factor authentication validated?

### Data Format and Quality
- What formats are required for dates, currencies, and numbers?
- How are email addresses and phone numbers validated?
- What character sets and lengths are acceptable for text fields?
- How are special characters and symbols handled?

### Business Logic and Workflows
- What steps must be completed in what order?
- Which operations depend on others being completed first?
- What states or statuses prevent certain actions?
- How are conflicting operations prevented?

---

## Success Criteria

Your validation rules extraction is complete when you can:
- Explain every situation where the system says "no" to a user request
- Describe all the requirements that must be met for operations to succeed
- Identify which rules protect security and compliance versus improving user experience
- Document how different rules work together to enforce business policies
- Provide clear examples that anyone can understand without technical knowledge

---

## Final Checklist

Before concluding your extraction, ensure you have:
- [ ] Identified all required fields and their acceptable formats
- [ ] Documented all numerical limits and ranges
- [ ] Listed all permission and authorization checks
- [ ] Captured all relationships between different data elements
- [ ] Noted all time-based or schedule-related restrictions
- [ ] Recorded all error messages and what triggers them
- [ ] Explained the business purpose behind each rule
- [ ] Organized rules in a logical, easy-to-navigate structure
- [ ] Used plain language throughout without technical jargon
- [ ] Provided real-world examples for clarity
