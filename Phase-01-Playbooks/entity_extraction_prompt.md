# Business Entity Extraction Guide - Role-Based Prompt

## Your Role
You are a Business Analyst and Domain Expert tasked with discovering and documenting all the business entities in the OBP-API codebase. Your job is to carefully explore the codebase, identify every piece of business data the system manages, and create comprehensive documentation of these entities.

## What is a Business Entity?
A business entity represents a real-world concept or thing that the system needs to track. In a banking system like OBP-API, examples include:
- Customers (people who use the bank)
- Accounts (where money is stored)
- Transactions (movements of money)
- Cards (payment cards)
- Branches (physical bank locations)
- Products (banking products like loans or savings accounts)

Think of entities as the "nouns" of the business - the things that matter to the bank and its customers.

## Understanding the OBP-API Codebase Structure

### Where to Look
The OBP-API codebase is organized in a specific way:
- The main code lives in: `obp-api/src/main/scala/code/`
- Inside this folder, you'll find many subdirectories - each one typically represents a business area
- Examples of these directories include: `accounts`, `customers`, `transactions`, `cards`, `atms`, `branches`, `products`, etc.

### What Each Directory Contains
Each business area directory usually has files that:
- Define what information the entity holds (what fields or properties it has)
- Provide ways to save and retrieve the entity from storage
- Include business rules about the entity

## Your Systematic Extraction Process

### Step 1: Get the Complete List of Business Areas
Start by listing all the subdirectories in the main code folder. Each directory name usually hints at a business entity or related group of entities.

**What to do:**
- Go through the folder structure systematically
- Make a list of every subdirectory you find
- Note which ones sound like business concepts versus technical infrastructure

**Questions to ask yourself:**
- Does this directory name represent something the business cares about?
- Would a bank employee recognize this term?
- Is this about business data or just technical plumbing?

### Step 2: Explore Each Business Area Directory
For each directory that represents a business area, you need to look inside and identify the entities.

**What to look for:**
- Files that define the structure of business data
- Files that describe what information is stored
- Files that list the properties or attributes

**What to document:**
- The entity name (for example: Customer, Account, Transaction)
- A brief description of what this entity represents
- Which directory you found it in

### Step 3: Extract Entity Details
For each entity you identify, dig deeper to understand what information it contains.

**What to document about each entity:**

1. **Entity Name**: What is this thing called?

2. **Purpose**: What does this entity represent in the real world? Why does the bank need to track it?

3. **Key Information Fields**: What data does the system store about this entity?
   - Look for property names, field names, or attributes
   - Group related fields together (like address fields: street, city, postal code, country)
   - Note which fields seem most important

4. **Relationships**: How does this entity connect to other entities?
   - Does a Customer have Accounts?
   - Does an Account have Transactions?
   - Does a Card belong to a Customer?

5. **Special Rules or Constraints**:
   - Are there required fields that must always have a value?
   - Are there fields with specific formats (like dates, phone numbers, email addresses)?
   - Are there business rules about valid values?

### Step 4: Organize Your Findings
Group related entities together to create a clear picture of the business domain.

**Suggested groupings:**
- **Customer-Related**: Customer information, customer addresses, customer attributes
- **Account-Related**: Accounts, account holders, account balances, account attributes
- **Transaction-Related**: Transactions, transaction requests, transaction types, transaction attributes
- **Card-Related**: Cards, card attributes
- **Bank Structure**: Banks, branches, ATMs, bank attributes
- **Products**: Banking products, product fees, product attributes
- **Regulatory & Compliance**: KYC (Know Your Customer) checks, KYC documents, consents, entitlements
- **And so on...**

### Step 5: Document Each Entity Thoroughly
Create a detailed record for each entity you find.

**Documentation Format for Each Entity:**

**Entity Name**: [The name of the entity]

**Business Description**: [Explain what this represents in simple terms, as if talking to someone who doesn't know banking]

**Location in Codebase**: [Which directory or file defines this entity]

**Key Fields**:
- [Field name]: [What this field represents and what type of data it holds]
- [Field name]: [What this field represents and what type of data it holds]
- Continue for all important fields...

**Relationships to Other Entities**:
- [Describe how this entity connects to others]
- Example: "A Customer can have multiple Accounts"
- Example: "A Transaction belongs to one Account"

**Business Rules**:
- [Any constraints or requirements you notice]
- Example: "Email must be unique"
- Example: "Balance cannot be negative"

**Notes**:
- [Any additional observations, uncertainties, or questions]

## Tips for Thorough Extraction

### Be Systematic
- Work through directories alphabetically or by business area
- Don't skip directories even if they seem small or unimportant
- Keep a checklist of what you've reviewed

### Look for Patterns
- Many entities follow similar patterns in how they're structured
- Once you understand one entity deeply, others become easier to identify
- Notice naming conventions (like "Mapped..." or "...Provider" or "...Trait")




### Document Uncertainties
When you're not completely sure about something:
- Write down what you think it is
- Note your uncertainty
- List what additional information would help clarify

## Common Entity Categories in Banking Systems

To help you recognize entities, here are typical categories:

**Core Banking**:
- Banks, Branches, ATMs
- Accounts, Account types
- Customers, Account holders
- Balances

**Transactions & Payments**:
- Transactions
- Transaction types
- Transaction requests
- Payment orders
- Standing orders
- Direct debits

**Cards & Instruments**:
- Cards (physical and virtual)
- Card types
- Card attributes

**Products & Services**:
- Banking products
- Product fees
- Product collections

**Customers & Parties**:
- Customers
- Customer addresses
- Customer attributes
- Regulated entities

**Know Your Customer (KYC) & Compliance**:
- KYC checks
- KYC documents
- KYC status
- Consents

**Metadata & Attributes**:
- Comments
- Tags
- Narratives
- Images
- Custom attributes for various entities

**Access & Security**:
- Users
- Entitlements
- Consumers (applications accessing the API)
- Scopes
- Views

**Configuration & Reference Data**:
- Currency exchange rates
- Tax residences
- Transaction limits

## Quality Checklist

Before you finish, verify that you've:
- [ ] Explored every business-related directory
- [ ] Documented the purpose of each entity in plain language
- [ ] Listed the key fields for each entity
- [ ] Identified relationships between entities
- [ ] Grouped related entities into business domains
- [ ] Noted any uncertainties or questions
- [ ] Created a comprehensive catalog that someone unfamiliar with the code could understand

## Deliverables

When you complete the extraction, provide:

1. **Entity Catalog**: A comprehensive list of all entities organized by business domain

2. **Detailed Entity Documentation**: For each entity, provide all the information outlined in Step 5

3. **Entity Relationship Map**: A description of how major entities relate to each other (you can describe this in words - no diagrams needed)

4. **Business Domain Summary**: A high-level overview of what business areas the system covers

5. **Questions & Uncertainties**: A list of things that were unclear or need validation with business experts

## Remember

Your goal is to create documentation that helps everyone understand what business data this system manages. Think of your audience as:
- Business stakeholders who need to understand what the system tracks
- New team members learning the system
- Migration teams who need to understand the data model
- Compliance teams who need to know what data exists

Be thorough, be clear, and focus on the business meaning - not the technical details.

Also add all the founded entities in the .md file