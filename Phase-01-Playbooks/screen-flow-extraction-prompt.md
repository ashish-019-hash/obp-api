# Screen Flow Extraction Prompt - Open Bank Project API

## Purpose
This guide will help you understand and document all the user journeys and screen flows in the banking application. Think of it as mapping out every path a user can take through the system, from start to finish.

## What is a Screen Flow?
A screen flow is the sequence of pages or screens a user sees and interacts with to complete a specific task. For example, when someone wants to log in, they might see a login screen, then a verification screen, and finally a dashboard.

## Who Should Use This Prompt?

### For Business Analysts
You'll document how users interact with the banking platform at each step. Focus on understanding what information is collected, what choices users make, and what happens after each action.

### For Product Managers
You'll map out the complete user experience to identify gaps, improvements, and new feature opportunities. Pay attention to how different user journeys connect and where users might face friction.

### For QA Engineers
You'll create test scenarios based on every possible path through the application. Document each screen's purpose and the conditions that lead users there.

### For Designers
You'll understand the current user interface flow to identify design improvements and ensure consistency across all screens.

### For Project Managers
You'll get a complete overview of the system's functionality to better plan resources and understand dependencies between features.

### For Developers
You'll document the technical flow to onboard new team members and maintain system documentation.

## What to Document

### 1. Entry Points
- **Question to answer**: Where do users start their journey?
- Look for:
  - Home page or landing page
  - Direct links from emails or external systems
  - Login screens
  - Registration pages
  - API endpoints that trigger user interfaces

### 2. Main User Journeys
Document each major task users can accomplish:

#### Account Access Journey
- How does a user create an account?
- What information do they need to provide?
- Are there different registration types?
- What verification steps are required?
- How does login work?
- What happens if they forget their password?

#### Authorization and Permissions Journey
- How do third-party applications get permission to access user data?
- What consent screens do users see?
- Can users review and revoke permissions?
- How is two-factor authentication handled?

#### Banking Operations Journey
- How do users view their accounts?
- How do users check their balances?
- How do users view transaction history?
- How do users make payments or transfers?
- How do users manage beneficiaries?
- How do users set up standing orders or direct debits?

#### Customer Management Journey
- How do users update their personal information?
- How do users manage their profile?
- How do users contact support?
- How do users handle complaints?

#### Developer Journey
- How do developers register for API access?
- How do they get API keys?
- How do they explore available APIs?
- How do they test their integrations?

### 3. Decision Points
At each screen, identify:
- What choices does the user have?
- What buttons or links can they click?
- What happens for each choice?
- Where does each path lead?

### 4. Required Information
For each screen, document:
- What information is displayed to the user?
- What information does the user need to provide?
- Which fields are required vs optional?
- What validation rules apply?

### 5. Alternative Paths
Consider:
- What happens when something goes wrong?
- Where do error messages appear?
- How does the user recover from errors?
- What are the alternative ways to accomplish the same task?

### 6. Different User Types
Document flows for:
- Regular customers
- Business customers
- Bank administrators
- Third-party developers
- Customer service representatives
- Compliance officers

### 7. Integration Points
Identify where the screens connect to:
- External payment systems
- Identity verification services
- Third-party applications
- Partner banks
- Regulatory reporting systems

## How to Extract the Information

### Step 1: Start with the Home Page
- What does a new visitor see first?
- What are all the possible actions from this page?
- Document each link and where it goes

### Step 2: Follow Each Path Systematically
- Pick one journey (like user registration)
- Document every screen in order
- Note what information appears on each screen
- Record what actions are available
- Identify what happens after each action

### Step 3: Map Dependencies
- Which screens require the user to complete previous steps first?
- What information carries over from one screen to the next?
- Where does user data get saved or validated?

### Step 4: Document Variations
- How does the flow change for different user types?
- Are there different paths for mobile vs desktop?
- What changes based on user preferences or settings?

### Step 5: Capture Edge Cases
- What happens if a session expires?
- What if the user navigates away and comes back?
- How are incomplete processes handled?
- What security measures interrupt normal flow?

## Questions to Answer for Each Screen

1. **Screen Purpose**: What is this screen trying to accomplish?
2. **User Entry**: How did the user arrive at this screen?
3. **Information Display**: What does the user see?
4. **User Input**: What can the user enter or select?
5. **Actions Available**: What buttons, links, or options are available?
6. **Validation**: What rules apply to user input?
7. **Success Path**: Where does the user go if everything works correctly?
8. **Error Path**: What happens if something goes wrong?
9. **Exit Points**: How can the user leave this screen?
10. **Dependencies**: What must happen before this screen can be accessed?

## Documentation Format

For each screen flow, create a document that includes:

### Flow Name
Give it a clear, descriptive name (e.g., "New Customer Registration Flow")

### Flow Description
Briefly explain what this flow accomplishes from the user's perspective

### Starting Point
Describe how users begin this journey

### Step-by-Step Flow
List each screen in order with:
- Screen name/title
- What the user sees
- What the user does
- What happens next
- Any variations or branches

### Completion Point
Describe what happens when the user successfully completes this flow

### Alternative Paths
Document any variations, shortcuts, or error handling paths

### User Types
Note which types of users can access this flow

## Special Considerations

### Security and Privacy Screens
- Identify all authentication checkpoints
- Document consent and permission screens
- Note where sensitive information is displayed or collected
- Map out multi-factor authentication flows

### Regulatory Compliance Screens
- Document terms and conditions acceptance
- Identify privacy policy acknowledgments
- Note where compliance information is displayed
- Map consent management interfaces

### Error Handling
- Document all error messages
- Identify recovery options for each error
- Map out help and support access points

### Session Management
- Document login/logout flows
- Identify session timeout handling
- Note re-authentication requirements

## Deliverables

By the end of this exercise, you should have:

1. **Complete Journey Maps**: Visual or written documentation of each user journey from start to finish

2. **Screen Inventory**: A list of all unique screens in the system with descriptions

3. **Flow Diagrams**: Simple diagrams showing how screens connect (can be drawn by hand or using simple tools)

4. **User Scenarios**: Real-world examples of how different users would navigate the system

5. **Gap Analysis**: Identification of missing flows, unclear paths, or inconsistent experiences

6. **Recommendations**: Suggestions for improving user flows based on your findings

## Tips for Success

- Start simple - pick one flow and document it completely before moving to the next
- Take screenshots or notes as you explore each screen
- Ask "what if" questions to uncover alternative paths
- Think from the user's perspective, not the system's
- Document what you see, not what you think should be there
- Use simple language that anyone can understand
- Group related flows together
- Create a visual map to see the big picture
- Validate your documentation by walking through it with someone else

## Common Pitfalls to Avoid

- Don't skip screens that seem unimportant - every screen matters
- Don't assume users follow the happy path - document all paths
- Don't forget about system-generated screens (confirmations, errors, etc.)
- Don't overlook administrative or support user flows
- Don't document just the latest version - note if there are multiple versions of flows

## Getting Started

Begin by answering these questions:
1. What is the main purpose of this banking application?
2. Who are the primary users?
3. What are the top 5 tasks users need to accomplish?
4. Where do most users start their journey?
5. What are the most critical security or compliance requirements?

Then, pick the most important user journey and start documenting it step by step, screen by screen.

---

Remember: The goal is to create documentation that anyone in your organization can read and understand, regardless of their technical background. Focus on clarity, completeness, and practical usability.
