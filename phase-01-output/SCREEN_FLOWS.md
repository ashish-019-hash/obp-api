# Open Bank Project API - Screen Flows and User Journey Documentation

## Document Overview

This document provides comprehensive documentation of all screen flows and user journeys in the Open Bank Project (OBP) API web application. It serves as a guide for business analysts, product managers, QA engineers, designers, developers, and project managers to understand how users interact with the system.

**Document Purpose:**
- Map all user entry points and navigation paths
- Document authentication and authorization flows
- Detail screen-by-screen user interactions
- Identify decision points and alternative paths
- Provide technical implementation references
- Support testing, design, and development activities

**Version Information:**
- Based on: OBP-API develop branch
- Web Application Location: `obp-api/src/main/webapp/`
- Backend Implementation: `obp-api/src/main/scala/`

---

## Table of Contents

1. [Entry Points](#entry-points)
2. [Authentication Methods](#authentication-methods)
3. [User Types and Personas](#user-types-and-personas)
4. [Main User Journeys](#main-user-journeys)
   - [Account Access Journey](#account-access-journey)
   - [Authorization and Permissions Journey](#authorization-and-permissions-journey)
   - [Developer Journey](#developer-journey)
   - [Consent Management Journey](#consent-management-journey)
   - [Banking Operations Journey](#banking-operations-journey)
   - [Compliance and Legal Journey](#compliance-and-legal-journey)
   - [User Authentication Context Journey](#user-authentication-context-journey)
   - [Customer Management Journey](#customer-management-journey)
5. [Complete Screen Inventory](#complete-screen-inventory)
6. [Integration Points](#integration-points)
7. [Flow Diagrams](#flow-diagrams)
8. [Alternative Paths and Error Handling](#alternative-paths-and-error-handling)

---

## Entry Points

### Primary Entry Points

Users can access the OBP API system through multiple entry points:

1. **Home Page / Landing Page**
   - URL: `/` (root of the application)
   - File: `webapp/index.html`
   - What users see: Welcome page with links to API Explorer, registration, login
   - Actions available:
     - Navigate to API Explorer
     - Register for developer account
     - Log in to existing account
     - View documentation links
     - Access sandbox environments

2. **Direct Login URL**
   - URL: `/user_mgt/login`
   - File: `webapp/templates-hidden/_login.html`
   - What users see: Login form with username/password fields
   - Actions available:
     - Login with credentials
     - OpenID Connect login
     - Password recovery
     - Sign up for new account

3. **OAuth Authorization URL**
   - URL: `/oauth/authorize?oauth_token=...`
   - File: `webapp/oauth/authorize.html`
   - What users see: Authorization request from third-party application
   - Actions available:
     - Authorize application
     - Deny authorization
     - View application details

4. **Consent Screen URL**
   - URL: `/consumer_consent/...`
   - File: `webapp/consent-screen.html`
   - What users see: OAuth 2.0/OIDC consent request
   - Actions available:
     - Grant consent
     - Deny consent
     - View requested permissions

5. **API Explorer (External)**
   - Configured via: `webui_api_explorer_url` property
   - Default: External web application (separate repository)
   - What users see: Interactive API documentation and testing interface
   - Actions available:
     - Browse API endpoints
     - Test API calls
     - View request/response examples
     - Access glossary and documentation

6. **Email Links**
   - User invitation emails
   - Password reset emails
   - Email verification links
   - OAuth notification emails

---

## Authentication Methods

The OBP API supports five primary authentication methods:

### 1. OAuth 1.0a

**Purpose:** Industry-standard authentication for third-party applications

**Implementation:**
- Backend: `code/snippet/OAuthAuthorisation.scala`
- Screen: `webapp/oauth/authorize.html`
- Success screen: `webapp/oauth/thanks.html`

**Flow:**
1. Third-party app redirects user to authorization URL with request token
2. User sees authorization screen showing app details
3. User logs in (if not already logged in)
4. User authorizes or denies the application
5. System generates verifier code
6. User redirected back to app with verifier
7. App exchanges verifier for access token

**User Experience:**
- Clear display of requesting application name and details
- Explanation of what access is being requested
- Login form if user not authenticated
- Authorize/Deny buttons
- Redirect to thank you page on success

**Key Decision Points:**
- Is user logged in? → If not, show login form
- Does user authorize? → If yes, generate verifier; if no, redirect with error

### 2. OAuth 2.0 / OpenID Connect (OIDC) with Hydra

**Purpose:** Modern authentication with consent management via ORY Hydra

**Implementation:**
- Backend: `code/snippet/ConsentScreen.scala`
- Screen: `webapp/consent-screen.html`
- Integration: ORY Hydra identity provider

**Flow:**
1. User redirected to consent screen with challenge parameter
2. System validates user session
3. User sees requested scopes and permissions
4. User accepts or rejects consent
5. System communicates with Hydra to finalize consent
6. User redirected back to application

**User Experience:**
- Display of application name and description
- List of requested permissions/scopes
- Accept/Reject buttons
- Integration with external identity providers

**Key Decision Points:**
- Is consent challenge valid? → If not, show error
- Does user accept? → If yes, grant consent and redirect; if no, deny and redirect
- Are all required scopes available? → Validate before proceeding

### 3. Direct Login

**Purpose:** Simplified header-based authentication for direct API access

**Implementation:**
- Header-based: Uses `DirectLogin` headers
- No dedicated UI screen (API-based)
- Token display: `webapp/dummy-user-tokens.html` (for sandbox testing)

**Flow:**
1. User obtains Direct Login token via API call
2. User includes token in `Authorization` header
3. System validates token on each request
4. No interactive screens involved

**User Experience:**
- Developers use API calls to obtain tokens
- Tokens displayed in sandbox environment for testing
- No browser-based interaction required

**Key Decision Points:**
- Is token valid? → Validate signature and expiration
- Does user have required entitlements? → Check permissions

### 4. Gateway Login

**Purpose:** Integration with API gateway for enterprise deployments

**Implementation:**
- Gateway-based authentication
- No dedicated UI screens (handled by gateway)
- Backend validation of gateway headers

**Flow:**
1. API Gateway authenticates user
2. Gateway adds authentication headers to request
3. OBP API validates gateway headers
4. Request processed if valid

**User Experience:**
- Transparent to end users
- Handled entirely by gateway infrastructure
- No OBP-specific UI screens

### 5. OpenID Connect (OIDC) Direct

**Purpose:** Integration with external OIDC identity providers

**Implementation:**
- Backend: `code/snippet/OpenIDConnectSnippet.scala`
- Uses standard OIDC protocol
- Integrates with providers like Keycloak, Auth0, etc.

**Flow:**
1. User clicks "Login with OIDC" button
2. Redirected to identity provider
3. User authenticates at provider
4. Provider redirects back with authorization code
5. OBP exchanges code for tokens
6. User logged in to OBP API

**User Experience:**
- Button on login page for OIDC providers
- Redirect to external login page
- Seamless return to OBP after authentication
- User information display after login

**Key Decision Points:**
- Is OIDC provider configured? → Check props settings
- Is authorization code valid? → Validate with provider
- Can user be mapped to OBP user? → Create or link user account

---

## User Types and Personas

### 1. Developer / API Consumer

**Description:** Software developers building applications that integrate with bank APIs

**Goals:**
- Register for API access
- Obtain API keys (consumer key/secret)
- Test API endpoints in sandbox
- Integrate banking functionality into applications
- Manage API consumers and applications

**Primary Journeys:**
- Developer registration
- Consumer key generation
- Sandbox account creation
- API testing via API Explorer
- OAuth application setup

**Access Levels:**
- Public registration available
- Sandbox environment access
- Production access (with approval)
- Rate limiting applies

**Key Screens:**
- Consumer registration (`consumer-registration.html`)
- Dummy user tokens (`dummy-user-tokens.html`)
- Sandbox account creation (`create-sandbox-account.html`)
- API Explorer (external)

### 2. Bank Customer / End User

**Description:** Individual or business customers accessing their banking information through third-party apps

**Goals:**
- Authorize third-party applications
- View account information
- Manage consents and permissions
- Perform banking operations via authorized apps
- Revoke access when needed

**Primary Journeys:**
- Login and authentication
- OAuth authorization
- Consent management
- Account access via third-party apps

**Access Levels:**
- Personal account access
- Own data only (privacy protected)
- Consent-based third-party access
- Can revoke permissions

**Key Screens:**
- Login (`_login.html`)
- OAuth authorization (`oauth/authorize.html`)
- Consent screen (`consent-screen.html`)
- Consent management (`consents.html`)
- User information (`user-information.html`)

### 3. Third-Party Application

**Description:** Software applications built by developers that access bank APIs on behalf of users

**Goals:**
- Obtain user authorization
- Access customer data (with consent)
- Perform banking operations
- Maintain valid access tokens
- Handle consent expiration

**Primary Journeys:**
- OAuth authorization flow
- Token refresh
- API calls with valid credentials
- Consent renewal

**Access Levels:**
- Limited to consented permissions
- Rate limited
- Time-limited access (token expiration)
- Revocable by user

**Key Screens:**
- OAuth authorization (user-facing)
- Consent screens (user-facing)
- Error pages for invalid/expired tokens

### 4. Bank Administrator

**Description:** Bank staff managing the API platform, users, and permissions

**Goals:**
- Manage user accounts
- Grant/revoke entitlements
- Configure system settings
- Monitor API usage
- Manage consumers and applications

**Primary Journeys:**
- Admin login
- Entitlement management (via API)
- User management (via API)
- System configuration

**Access Levels:**
- Full system access
- User management capabilities
- Entitlement granting authority
- System configuration access

**Key Screens:**
- Admin login (uses same login screen)
- Management operations via API calls
- No dedicated admin UI (API-based management)

### 5. Invited User

**Description:** Users who receive invitation links to join the platform

**Goals:**
- Complete registration via invitation
- Accept terms and conditions
- Set up account
- Begin using platform

**Primary Journeys:**
- Invitation link click
- Registration completion
- Initial login

**Access Levels:**
- As configured in invitation
- May include pre-granted entitlements
- Invitation-specific permissions

**Key Screens:**
- User invitation (`user-invitation.html`)

### 6. Compliance Officer / Auditor

**Description:** Staff responsible for regulatory compliance and audit trails

**Goals:**
- Access audit logs
- Review consent history
- Monitor regulatory compliance
- Generate compliance reports

**Primary Journeys:**
- Login with audit credentials
- Access compliance APIs
- Review consent records

**Access Levels:**
- Read-only access to audit data
- Compliance-specific entitlements
- Historical data access

**Key Screens:**
- Uses API-based access
- No dedicated UI screens (API Explorer for queries)

---

## Main User Journeys

## Account Access Journey

### Journey 1: New User Registration (Sign Up)

**Journey Name:** New User Registration

**Description:** A new user creates an account to access the OBP API platform, either as a developer or end user.

**Starting Point:**
- URL: Home page `/` or direct to sign-up page
- Entry method: Click "Sign Up" link from home page or login page
- Screen: `webapp/index.html` → Sign up form

**User Type:** Any new user (developers, customers, etc.)

**Step-by-Step Flow:**

1. **Home Page**
   - User lands on index.html
   - Sees "Sign Up" link
   - Clicks to begin registration
   - **Backend:** `code/snippet/WebUI.scala` renders home page

2. **Sign Up Form**
   - User sees registration form with fields:
     - First Name (required)
     - Last Name (required)
     - Email (required)
     - Username (required)
     - Password (required)
     - Password Repeat (required)
     - Privacy Policy checkbox (required)
     - Terms and Conditions checkbox (required)
   - **Backend:** `code/model/dataAccess/AuthUser.scala` (signupFields method)
   - **Validation:** Client-side JavaScript validates matching passwords
   - **Submit button:** Disabled until all required fields filled

3. **Form Submission**
   - User clicks "Sign Up" button
   - System validates:
     - All required fields present
     - Email format valid
     - Username not already taken
     - Password meets requirements
     - Passwords match
     - Checkboxes accepted
   - **Backend:** `code/model/dataAccess/AuthUser.scala` (actionsAfterSignup method)

4. **Email Verification (if enabled)**
   - System creates AuthUser and ResourceUser records
   - Sends verification email (if skipEmailValidation = false)
   - User sees message: "Please check your email to verify your account"
   - **Backend:** Sends email via `sendValidationEmail` method
   - **Alternative:** If email verification disabled, user logged in immediately

5. **Email Verification Link**
   - User checks email
   - Clicks verification link
   - Redirected to OBP API
   - Account activated
   - User can now log in

6. **First Login**
   - User redirected to login page
   - Enters username and password
   - Successfully authenticated
   - Sees home page or redirected to intended destination

**Completion Point:** User has verified account and can log in

**Alternative Paths:**

- **Username already exists:**
  - Error message displayed
  - User must choose different username
  - Remains on sign-up form

- **Email already registered:**
  - Error message displayed
  - Suggests password recovery
  - Cannot proceed with registration

- **Password validation fails:**
  - Error message shows requirements
  - User must correct password
  - Remains on sign-up form

- **Terms not accepted:**
  - Submit button remains disabled
  - User must check boxes to proceed

- **Email verification skipped:**
  - User immediately logged in
  - Can begin using platform
  - No verification email sent

**Required Information:**
- First Name
- Last Name
- Email address (unique)
- Username (unique)
- Password (meeting requirements)
- Agreement to Privacy Policy
- Agreement to Terms and Conditions

**Decision Points:**
- Email verification enabled? → Determines if verification email sent
- Username available? → Validates uniqueness
- Email available? → Validates uniqueness
- Password meets requirements? → Validates complexity

**Backend Implementation:**
- `code/model/dataAccess/AuthUser.scala` - Main user authentication model
- `code/model/dataAccess/ResourceUser.scala` - User resource management
- User creation creates both AuthUser and linked ResourceUser

**Integration Points:**
- Email service for verification
- Database (creates AuthUser and ResourceUser records)
- Props configuration for email settings
- Entitlement system (default entitlements may be granted)

---

### Journey 2: User Login

**Journey Name:** User Login

**Description:** Existing user authenticates to access the OBP API platform.

**Starting Point:**
- URL: `/user_mgt/login`
- Entry method: Click "Login" from home page or direct URL access
- Screen: `webapp/templates-hidden/_login.html`

**User Type:** Any registered user

**Step-by-Step Flow:**

1. **Navigate to Login Page**
   - User clicks "Login" link from home page
   - Redirected to login screen
   - **URL:** `/user_mgt/login`
   - **File:** `templates-hidden/_login.html`
   - **Backend:** `code/snippet/Login.scala`

2. **Login Screen Display**
   - User sees form with:
     - Username field
     - Password field
     - "Login" button
     - "Forgot password?" link
     - "Sign up" link
     - Optional: "Login with OpenID Connect" button
   - **Customization:** Text and instructions from props (webui_login_page_text)

3. **Enter Credentials**
   - User enters username
   - User enters password
   - Clicks "Login" button

4. **Authentication Validation**
   - System validates credentials:
     - Username exists?
     - Password correct?
     - Account verified?
     - Account not locked?
   - **Backend:** `code/model/dataAccess/AuthUser.scala` validates credentials
   - **Security:** Bad login attempts tracked in `code/loginattempts/MappedBadLoginAttempt`

5. **Success Path**
   - Credentials valid
   - User authenticated
   - Session created
   - **Backend:** `code/snippet/Login.scala` manages session
   - Redirect to:
     - Originally requested page (if deep link)
     - Home page (default)
     - User information page

6. **Post-Login State**
   - User sees logged-in interface
   - Navigation shows "Logout" option
   - User name displayed
   - Access to protected resources enabled

**Completion Point:** User successfully logged in and can access protected resources

**Alternative Paths:**

- **Invalid Credentials:**
  - Error message: "Invalid username or password"
  - Remains on login page
  - Login attempt recorded
  - **Backend:** `MappedBadLoginAttempt` incremented

- **Account Locked:**
  - Error message: "Account locked due to too many failed login attempts"
  - Cannot proceed with login
  - User must contact support or wait for unlock
  - **Backend:** `code/userlocks/UserLocks` checks lock status

- **Account Not Verified:**
  - Error message: "Please verify your email address"
  - Link to resend verification email
  - Cannot proceed with login

- **Forgot Password:**
  - User clicks "Forgot password?" link
  - Redirected to password recovery page
  - Enters email address
  - Receives password reset link

- **OpenID Connect Login:**
  - User clicks "Login with OIDC" button
  - Redirected to OIDC provider
  - Authenticates at provider
  - Returned to OBP API authenticated
  - **Backend:** `code/snippet/OpenIDConnectSnippet.scala`

- **Session Timeout:**
  - If returning after session timeout
  - User must log in again
  - Redirected to login page with message

**Required Information:**
- Username (or email)
- Password

**Decision Points:**
- Credentials valid? → Proceed or show error
- Account verified? → Proceed or require verification
- Account locked? → Block login or allow
- OIDC enabled? → Show OIDC button
- SSO enabled? → Redirect to SSO provider

**Backend Implementation:**
- `code/snippet/Login.scala` - Login UI logic
- `code/model/dataAccess/AuthUser.scala` - Authentication validation
- `code/loginattempts/MappedBadLoginAttempt` - Failed attempt tracking
- `code/userlocks/UserLocks` - Account locking
- `code/snippet/OpenIDConnectSnippet.scala` - OIDC integration

**Security Features:**
- Password hashing (not stored in plain text)
- Failed login attempt tracking
- Account locking after repeated failures
- Session management
- CSRF protection

**Integration Points:**
- Database (user lookup and validation)
- Session management (Lift framework)
- OIDC provider (if enabled)
- Email service (for password reset)

---

### Journey 3: User Invitation Completion

**Journey Name:** Complete User Invitation

**Description:** User who received an invitation email completes registration to join the platform.

**Starting Point:**
- URL: Invitation link from email (contains unique token)
- Entry method: Click link in invitation email
- Screen: `webapp/user-invitation.html`

**User Type:** Invited user (developer, partner, customer)

**Step-by-Step Flow:**

1. **Receive Invitation Email**
   - User receives email with invitation link
   - Email contains unique invitation token
   - Link includes: `/user/invitation?token=...`

2. **Click Invitation Link**
   - User clicks link in email
   - Browser opens invitation page
   - **File:** `webapp/user-invitation.html`
   - **Backend:** `code/snippet/UserInvitation.scala`

3. **Invitation Form Display**
   - System validates invitation token
   - If valid, shows form with:
     - First Name (required)
     - Last Name (required)
     - Developer Email (required)
     - Company (required)
     - Country (dropdown, required)
     - Username (required)
     - Privacy Policy checkbox (required)
     - Terms and Conditions checkbox (required)
   - Some fields may be pre-filled from invitation
   - Submit button initially disabled

4. **Fill Out Form**
   - User enters all required information
   - Selects country from dropdown
   - Checks both required checkboxes
   - **JavaScript validation:** Submit button enables when all fields complete

5. **Form Submission**
   - User clicks "Submit" button
   - System validates:
     - Token still valid (not expired)
     - All required fields present
     - Username not already taken
     - Email format valid
     - Checkboxes checked

6. **Account Creation**
   - System creates user account
   - Links to invitation
   - Assigns any pre-configured entitlements from invitation
   - Records privacy policy and T&C acceptance
   - **Backend:** Creates AuthUser and ResourceUser

7. **Success Response**
   - User sees success message
   - May auto-login user
   - Or redirect to login page
   - Email confirmation may be sent

8. **First Login**
   - User logs in with created username
   - Can begin using platform
   - Pre-granted permissions active

**Completion Point:** User account created from invitation and ready to use

**Alternative Paths:**

- **Invalid/Expired Token:**
  - Error message: "Invalid or expired invitation"
  - Cannot proceed
  - User must request new invitation

- **Username Already Taken:**
  - Error message displayed
  - User must choose different username
  - Remains on form

- **Missing Required Fields:**
  - Submit button remains disabled
  - User must complete all fields
  - Validation messages shown

- **Terms Not Accepted:**
  - Submit button remains disabled
  - User must check boxes
  - Cannot submit form

- **Invitation Already Used:**
  - Error message: "This invitation has already been used"
  - Cannot proceed
  - User directed to login page

**Required Information:**
- First Name
- Last Name
- Developer Email
- Company name
- Country
- Username (unique)
- Privacy Policy acceptance
- Terms and Conditions acceptance

**Decision Points:**
- Is invitation token valid? → Proceed or show error
- Is token expired? → Check timestamp
- Has invitation been used? → Check status
- Username available? → Validate uniqueness
- All required fields complete? → Enable submit

**Backend Implementation:**
- `code/snippet/UserInvitation.scala` - Invitation handling logic
- `code/model/dataAccess/AuthUser.scala` - User creation
- Invitation table - Stores invitation details and status

**JavaScript Functionality:**
- File: `webapp/user-invitation.html`
- Monitors form inputs
- Validates required fields
- Enables/disables submit button
- Shows/hides error messages

**Integration Points:**
- Email service (invitation email)
- Database (user creation, invitation tracking)
- Entitlement system (pre-granted permissions)
- User agreement storage (privacy policy, T&C)

---

### Journey 4: User Information Display

**Journey Name:** View User Information

**Description:** Authenticated user views their account information, tokens, and consent settings.

**Starting Point:**
- URL: `/user-information`
- Entry method: Navigation after login or from consent management
- Screen: `webapp/user-information.html`

**User Type:** Authenticated end user

**Step-by-Step Flow:**

1. **Navigate to User Information**
   - User already logged in
   - Navigates to user information page
   - **URL:** `/user-information`
   - **File:** `webapp/user-information.html`
   - **Backend:** `code/snippet/UserInformation.scala`

2. **Information Display**
   - User sees read-only information:
     - Email address
     - Username
     - Provider (authentication method)
     - ID Token (if using OIDC)
     - Access Token (if using OIDC)
     - User ID (internal)
   - All fields are read-only (not editable)
   - Clean, white background for read-only fields

3. **Available Actions**
   - "My Consents" button
     - Redirects to `/consents` page
     - View and manage granted consents
   - Logout option (in navigation)

4. **View Tokens**
   - Developer can view authentication tokens
   - Useful for debugging
   - Tokens displayed in full (no masking for user viewing own tokens)

**Completion Point:** User has viewed their information

**Alternative Paths:**

- **Not Logged In:**
  - Redirected to login page
  - Must authenticate first
  - Return to this page after login

- **Navigate to Consents:**
  - Click "My Consents" button
  - Redirected to consent management page
  - Can view/revoke consents

**Required Information:**
- User must be authenticated
- No input required (display only)

**Decision Points:**
- Is user authenticated? → Show information or redirect to login
- Which provider? → Display appropriate token information

**Backend Implementation:**
- `code/snippet/UserInformation.scala` - Information retrieval logic
- Session management - Validates authentication
- User lookup - Retrieves user details

**Display Fields:**
- Email (from user account)
- Username (from user account)
- Provider (authentication method used)
- ID Token (OIDC only)
- Access Token (OIDC only)
- User ID (internal identifier)

**Integration Points:**
- Session management
- User database
- OIDC token storage
- Consent management system (via button link)

---

## Authorization and Permissions Journey

### Journey 5: OAuth 1.0a Authorization Flow

**Journey Name:** OAuth 1.0a Third-Party Authorization

**Description:** User authorizes a third-party application to access their banking data using OAuth 1.0a protocol.

**Starting Point:**
- URL: `/oauth/authorize?oauth_token=REQUEST_TOKEN`
- Entry method: Redirect from third-party application
- Screen: `webapp/oauth/authorize.html`

**User Type:** Bank customer authorizing third-party app

**Step-by-Step Flow:**

1. **Third-Party App Initiates Flow**
   - External application requests access
   - App obtains request token from OBP API
   - App redirects user to authorization URL
   - **URL format:** `/oauth/authorize?oauth_token={REQUEST_TOKEN}`

2. **Authorization Page Load**
   - Browser navigates to authorization URL
   - **File:** `webapp/oauth/authorize.html`
   - **Backend:** `code/snippet/OAuthAuthorisation.scala`
   - System validates request token

3. **Check User Authentication**
   - **Decision Point:** Is user logged in?
   - **If NOT logged in:**
     - Show login form within authorization page
     - User must enter username and password
     - Authenticate before proceeding
   - **If logged in:**
     - Skip to authorization display

4. **Authorization Request Display**
   - User sees:
     - Name of requesting application
     - Application description
     - What access is requested
     - "Authorize" button
     - "Deny" button
   - Clear explanation of what authorization means

5. **User Decision**
   - **Option A: Authorize**
     - User clicks "Authorize" button
     - System generates OAuth verifier code
     - **Backend:** Creates token mapping
   - **Option B: Deny**
     - User clicks "Deny" button
     - Authorization rejected
     - User redirected back to app with error

6. **Generate Verifier (if authorized)**
   - System creates verifier code
   - Links verifier to request token
   - Stores authorization in database
   - Prepares callback URL

7. **Redirect to Application**
   - **If authorized:**
     - Redirect to app callback URL
     - Includes: `oauth_token` and `oauth_verifier`
     - URL format: `{callback}?oauth_token={TOKEN}&oauth_verifier={VERIFIER}`
   - **If denied:**
     - Redirect to callback URL with error
     - URL format: `{callback}?oauth_token={TOKEN}&error=access_denied`

8. **Thank You Page (optional)**
   - **File:** `webapp/oauth/thanks.html`
   - Shows confirmation message
   - May auto-redirect to application
   - User sees success feedback

9. **Application Completes Flow**
   - App receives verifier code
   - App exchanges request token + verifier for access token
   - App can now make authenticated API calls
   - **Note:** This step happens in app backend (not visible to user)

**Completion Point:** User has authorized app and returned to application

**Alternative Paths:**

- **Invalid Request Token:**
  - Error message: "Invalid or expired token"
  - Cannot proceed with authorization
  - User may see error page or redirect

- **User Denies Authorization:**
  - Authorization rejected
  - No verifier generated
  - User redirected to app with error
  - App cannot access data

- **User Not Logged In:**
  - Login form displayed within authorization page
  - User must authenticate
  - After login, authorization request shown
  - Can then authorize or deny

- **Callback URL Missing:**
  - Error: Cannot redirect user
  - Manual copy of verifier may be provided
  - User manually returns to app

- **Token Expired:**
  - Error message shown
  - User cannot proceed
  - App must restart OAuth flow

**Required Information:**
- Valid OAuth request token (in URL)
- User authentication (username/password if not logged in)
- User decision (authorize or deny)

**Decision Points:**
- Is request token valid? → Proceed or show error
- Is user logged in? → Show auth request or show login form
- Does user authorize? → Generate verifier or redirect with error
- Is callback URL present? → Redirect or show manual verifier

**Backend Implementation:**
- `code/snippet/OAuthAuthorisation.scala` - Main OAuth logic
  - `tokenCheck` method - Validates token and manages flow
  - `loggedIn` method - Checks authentication
  - `loggedOut` method - Handles unauthenticated state
- OAuth token storage - Request tokens, verifiers, access tokens
- `code/snippet/OAuthWorkedThanks.scala` - Thank you page logic

**Security Features:**
- Request token validation
- Verifier code generation (one-time use)
- Token expiration
- User consent required
- Callback URL validation

**Integration Points:**
- Third-party application (initiates and completes flow)
- OAuth token database
- User authentication system
- Callback URL validation

**User Experience Notes:**
- Clear explanation of what app is requesting
- Simple authorize/deny choice
- Login form integrated if needed
- Quick redirect back to app
- Optional confirmation page

---

### Journey 6: OAuth 2.0 / OIDC Consent Flow

**Journey Name:** OAuth 2.0 / OpenID Connect Consent

**Description:** User grants consent for OAuth 2.0 or OIDC application access, managed through ORY Hydra.

**Starting Point:**
- URL: `/consumer_consent?consent_challenge=...`
- Entry method: Redirect from Hydra during OAuth 2.0/OIDC flow
- Screen: `webapp/consent-screen.html`

**User Type:** Bank customer authorizing OAuth 2.0/OIDC application

**Step-by-Step Flow:**

1. **OAuth 2.0 Flow Initiated**
   - Third-party application initiates OAuth 2.0 flow
   - User redirected through ORY Hydra
   - Hydra generates consent challenge
   - User redirected to OBP consent screen
   - **URL format:** `/consumer_consent?consent_challenge={CHALLENGE}`

2. **Consent Screen Load**
   - Browser navigates to consent URL
   - **File:** `webapp/consent-screen.html`
   - **Backend:** `code/snippet/ConsentScreen.scala`
   - System validates consent challenge with Hydra

3. **Fetch Consent Request Details**
   - **Backend:** Calls Hydra API to get consent details
   - Retrieves:
     - Client application name
     - Requested OAuth scopes
     - User information
     - Challenge validity
   - **Method:** `submitConsentApprovalRequest` or `submitConsentDenyRequest`

4. **Display Consent Request**
   - User sees:
     - Name of requesting application
     - List of requested permissions (scopes)
     - User identity being shared
     - "Accept" button (green)
     - "Reject" button (red)
   - Clear explanation of each scope

5. **User Decision**
   - **Option A: Accept Consent**
     - User clicks "Accept" button
     - Form submitted with action "accept"
     - **Backend:** Processes acceptance
   - **Option B: Reject Consent**
     - User clicks "Reject" button
     - Form submitted with action "reject"
     - **Backend:** Processes rejection

6. **Process Acceptance (if accepted)**
   - System calls Hydra consent acceptance API
   - Hydra generates authorization code
   - OBP receives redirect URL from Hydra
   - **Backend method:** `submitConsentApprovalRequest`

7. **Process Rejection (if rejected)**
   - System calls Hydra consent rejection API
   - Hydra generates error response
   - OBP receives redirect URL with error
   - **Backend method:** `submitConsentDenyRequest`

8. **Redirect to Application**
   - User redirected to URL provided by Hydra
   - **If accepted:**
     - Contains authorization code
     - App can exchange for access token
   - **If rejected:**
     - Contains error information
     - App knows consent denied

9. **Application Completes Flow**
   - App exchanges authorization code for tokens
   - Receives access token and optionally refresh token
   - Can make authenticated API calls
   - **Note:** Happens in app backend

**Completion Point:** User returned to application with authorization result

**Alternative Paths:**

- **Invalid Consent Challenge:**
  - Error message displayed
  - Cannot proceed
  - User may see error page
  - Must restart OAuth flow

- **Consent Challenge Expired:**
  - Error: "Consent request expired"
  - Cannot proceed
  - Application must reinitiate flow

- **User Rejects Consent:**
  - Rejection processed
  - User redirected to app with error
  - App cannot access data
  - No tokens issued

- **Hydra Communication Error:**
  - Error message: "Unable to process consent"
  - Technical error displayed
  - User may retry or contact support

- **User Already Consented:**
  - Hydra may skip consent screen
  - User directly redirected to app
  - Previous consent used (if "remember consent" was selected)

**Required Information:**
- Valid consent challenge (in URL parameter)
- User must be authenticated
- User decision (accept or reject)

**Decision Points:**
- Is consent challenge valid? → Proceed or show error
- Is user authenticated? → Continue or redirect to login
- Does user accept? → Approve or reject in Hydra
- Is Hydra reachable? → Process or show error

**Backend Implementation:**
- `code/snippet/ConsentScreen.scala` - Main consent logic
  - `submitConsentApprovalRequest` - Handles acceptance
  - `submitConsentDenyRequest` - Handles rejection
  - `render` - Displays consent form
- Hydra integration - External HTTP calls
- `code/util/HydraUtil.scala` - Hydra communication utilities

**Security Features:**
- Consent challenge validation
- One-time use challenges
- Hydra integration for identity management
- User consent required
- Scope validation

**Integration Points:**
- ORY Hydra (identity provider)
- Third-party OAuth 2.0 application
- OBP user authentication
- Consent storage

**Hydra Interaction:**
1. GET consent request details from Hydra
2. Display consent UI to user
3. POST acceptance or rejection to Hydra
4. Receive redirect URL from Hydra
5. Redirect user to application

**User Experience Notes:**
- Modern, clean consent UI
- Clear explanation of permissions
- Simple accept/reject choice
- Fast redirect back to application
- Integration with enterprise identity systems

---

### Journey 7: Berlin Group Consent (PSD2)

**Journey Name:** Berlin Group PSD2 Consent Request

**Description:** User confirms consent for Berlin Group (NextGenPSD2) compliant third-party access to banking data.

**Starting Point:**
- URL: `/berlin-group/consent-request?consentId=...`
- Entry method: Redirect during Berlin Group API consent flow
- Screen: `webapp/confirm-bg-consent-request.html`

**User Type:** Bank customer (EU/EEA jurisdiction with PSD2)

**Step-by-Step Flow:**

1. **Berlin Group API Request**
   - Third-party provider (TPP) requests consent via Berlin Group API
   - System creates consent request
   - User redirected to consent confirmation screen
   - **URL format:** `/berlin-group/consent-request?consentId={ID}`

2. **Consent Request Screen Load**
   - Browser navigates to confirmation URL
   - **File:** `webapp/confirm-bg-consent-request.html`
   - **Backend:** `code/snippet/BerlinGroupConsent.scala`
   - System retrieves consent details

3. **Display Consent Details**
   - User sees:
     - Berlin Group Consent Request title
     - Detailed consent information (JSON formatted)
     - Account access details
     - Validity period
     - TPP information
     - "Confirm" button (red/warning style)
     - "Deny" link/button
   - **Format:** Pre-formatted text with proper word wrapping

4. **User Reviews Consent**
   - User reads consent details
   - Understands what data will be shared
   - Sees which accounts are involved
   - Notes expiration date

5. **User Decision**
   - **Option A: Confirm**
     - User clicks "Confirm" button
     - Form submitted via POST
     - **Backend:** Processes confirmation
   - **Option B: Deny**
     - User clicks "Deny" link
     - Redirected to home page (`/`)
     - Consent not granted

6. **Process Confirmation**
   - System updates consent status to confirmed
   - Records user acceptance
   - Updates Berlin Group consent database
   - **Backend:** `confirmBerlinGroupConsentRequest` method

7. **Strong Customer Authentication (SCA)**
   - If SCA required:
     - User may be redirected to SCA screen
     - **File:** `webapp/confirm-bg-consent-request-sca.html`
     - Additional authentication required
     - See Journey 8 for SCA details

8. **Redirect to TPP**
   - After confirmation (and SCA if needed)
   - User redirected back to TPP application
   - TPP can now use consent to access data
   - Consent ID valid for specified period

**Completion Point:** Berlin Group consent confirmed and TPP authorized

**Alternative Paths:**

- **Invalid Consent ID:**
  - Error message displayed
  - Cannot proceed
  - User may contact support

- **User Denies Consent:**
  - Clicked "Deny" button
  - Redirected to home page
  - Consent not granted
  - TPP cannot access data

- **Consent Already Confirmed:**
  - Error or redirect
  - Cannot confirm twice
  - May show status page

- **SCA Required:**
  - User redirected to SCA screen
  - Must complete additional authentication
  - Then consent finalized

- **Consent Expired:**
  - Error message
  - Cannot confirm expired consent
  - TPP must create new consent

**Required Information:**
- Valid consent ID (in URL parameter)
- User must be authenticated
- User decision (confirm or deny)

**Decision Points:**
- Is consent ID valid? → Proceed or show error
- Is user authenticated? → Continue or require login
- Does user confirm? → Update status or deny
- Is SCA required? → Redirect to SCA or complete

**Backend Implementation:**
- `code/snippet/BerlinGroupConsent.scala` - Consent confirmation logic
  - `confirmBerlinGroupConsentRequest` - Processes confirmation
  - Renders consent details
- Berlin Group consent storage - Consent records
- SCA integration - For authentication

**Regulatory Context:**
- PSD2 (Payment Services Directive 2) compliance
- Berlin Group NextGenPSD2 framework
- Strong Customer Authentication (SCA)
- Consent management requirements
- Audit trail requirements

**Consent Details Displayed:**
- Consent ID
- Accounts to be accessed
- Permissions granted (transactions, balances, etc.)
- Validity period
- TPP identification
- Purpose of access

**Integration Points:**
- Berlin Group API compliance layer
- TPP application
- SCA mechanism
- Consent database
- Audit logging

**User Experience Notes:**
- Clear consent details in readable format
- PSD2-compliant consent display
- Simple confirm/deny choice
- May include SCA step
- Regulatory information visible

---

### Journey 8: Berlin Group SCA (Strong Customer Authentication)

**Journey Name:** Berlin Group Strong Customer Authentication

**Description:** User completes Strong Customer Authentication as required by PSD2 for Berlin Group consent.

**Starting Point:**
- URL: `/berlin-group/consent-request-sca?consentId=...`
- Entry method: Redirect after consent confirmation requiring SCA
- Screen: `webapp/confirm-bg-consent-request-sca.html`

**User Type:** Bank customer (EU/EEA with PSD2 SCA requirement)

**Step-by-Step Flow:**

1. **SCA Requirement Triggered**
   - User confirmed Berlin Group consent
   - System determines SCA required
   - User redirected to SCA screen
   - **URL format:** `/berlin-group/consent-request-sca?consentId={ID}`

2. **SCA Screen Load**
   - Browser navigates to SCA URL
   - **File:** `webapp/confirm-bg-consent-request-sca.html`
   - **Backend:** `code/snippet/BerlinGroupConsent.scala`
   - System prepares SCA challenge

3. **Display SCA Request**
   - User sees:
     - SCA requirement explanation
     - Authentication method options
     - Instructions for completing SCA
     - Form for SCA input
     - "Confirm" button

4. **SCA Methods**
   - **Possible methods:**
     - SMS OTP (One-Time Password)
     - Mobile app authentication
     - Hardware token
     - Biometric authentication
   - **Depends on:** Bank's SCA implementation

5. **User Completes SCA**
   - User receives OTP or opens authenticator app
   - Enters authentication code
   - Clicks "Confirm" button
   - Form submitted

6. **Validate SCA**
   - System validates authentication
   - Checks code/token validity
   - Verifies timing (not expired)
   - **Backend:** SCA validation logic

7. **SCA Success**
   - Authentication confirmed
   - Consent finalized
   - User authenticated per PSD2 requirements
   - Redirect to completion page or TPP

8. **Complete Consent Flow**
   - Consent now fully confirmed with SCA
   - TPP notified
   - User can return to TPP application

**Completion Point:** SCA completed and consent fully authorized

**Alternative Paths:**

- **Invalid SCA Code:**
  - Error message: "Invalid authentication code"
  - User can retry
  - Limited number of attempts
  - After max attempts, must restart

- **SCA Timeout:**
  - Error: "Authentication timed out"
  - User must request new code
  - Restart SCA process

- **SCA Not Available:**
  - Error: "Cannot complete authentication"
  - User may need to contact bank
  - Alternative authentication may be offered

- **User Cancels:**
  - User can go back
  - Consent not finalized
  - TPP not authorized

**Required Information:**
- Valid consent ID
- SCA authentication code/token
- Timing within validity window

**Decision Points:**
- Which SCA method available? → Display appropriate form
- Is SCA code valid? → Approve or reject
- Is SCA timeout? → Allow or require new code
- Max attempts exceeded? → Block or allow retry

**Backend Implementation:**
- `code/snippet/BerlinGroupConsent.scala` - SCA handling
- SCA validation logic
- OTP generation/validation (if SMS method)
- Integration with authentication provider

**Regulatory Requirements:**
- PSD2 Strong Customer Authentication
- Two-factor authentication minimum
- Time-bound authentication
- Secure code generation
- Audit trail

**SCA Factors:**
1. **Knowledge:** Something user knows (password, PIN)
2. **Possession:** Something user has (phone, token)
3. **Inherence:** Something user is (biometric)

**Minimum:** Two of three factors required

**Integration Points:**
- SMS gateway (for OTP)
- Mobile app (for push notifications)
- Biometric systems (if supported)
- Hardware token systems
- Consent database

**User Experience Notes:**
- Clear explanation of SCA requirement
- Multiple authentication options if available
- Retry capability for errors
- Time pressure (limited validity)
- Secure, compliant process

---

### Journey 9: VRP Consent Creation

**Journey Name:** Variable Recurring Payments Consent

**Description:** User creates consent for Variable Recurring Payments (UK Open Banking).

**Starting Point:**
- URL: `/vrp-consent-request?...`
- Entry method: Redirect during VRP setup
- Screen: `webapp/confirm-vrp-consent-request.html` or `webapp/confirm-vrp-consent.html`

**User Type:** Bank customer (UK Open Banking jurisdiction)

**Step-by-Step Flow:**

1. **VRP Request Initiated**
   - Third-party provider requests VRP consent
   - User redirected to VRP consent screen
   - **Files:**
     - `webapp/confirm-vrp-consent-request.html` - Request confirmation
     - `webapp/confirm-vrp-consent.html` - Final consent

2. **Display VRP Details**
   - User sees:
     - VRP consent details
     - Payment limits
     - Frequency limits
     - Valid from/to dates
     - Merchant information
     - "Confirm" button
     - "Deny" option

3. **VRP Consent Parameters**
   - **Displayed information:**
     - Maximum amount per payment
     - Maximum total amount (period)
     - Payment frequency limits
     - Validity period
     - Account to be debited
     - Merchant/payee details

4. **User Reviews VRP Terms**
   - User reads VRP conditions
   - Understands payment limits
   - Notes validity period
   - Checks account details

5. **User Decision**
   - **Option A: Confirm**
     - User clicks "Confirm" button
     - VRP consent created
     - **Backend:** Processes consent
   - **Option B: Deny**
     - User denies VRP
     - No consent created
     - Redirect back to merchant

6. **Process Confirmation**
   - System creates VRP consent record
   - Stores payment limits
   - Records validity period
   - **Backend:** `code/snippet/VrpConsentCreation.scala`

7. **Redirect to Merchant**
   - User redirected back to merchant
   - Merchant notified of consent
   - VRP now active for recurring payments

**Completion Point:** VRP consent created and active

**Alternative Paths:**

- **User Denies VRP:**
  - No consent created
  - Redirect to merchant with error
  - Merchant cannot make VRP

- **Invalid Parameters:**
  - Error displayed
  - User cannot proceed
  - Must contact merchant

- **Consent Limits Exceeded:**
  - Warning displayed
  - User must approve explicitly
  - Clear indication of unusual amounts

**Required Information:**
- Valid VRP request
- User authenticated
- User confirmation

**Decision Points:**
- Are VRP limits acceptable? → User reviews and decides
- Is user authenticated? → Require login
- Does user confirm? → Create consent or deny

**Backend Implementation:**
- `code/snippet/VrpConsentCreation.scala` - VRP consent logic
- VRP consent storage
- Payment limit validation
- UK Open Banking compliance

**UK Open Banking Context:**
- Variable Recurring Payments standard
- CMA9 compliance
- OBIE (Open Banking Implementation Entity) standards
- Payment limit requirements
- Consumer protection requirements

**VRP Use Cases:**
- Subscription payments
- Regular bills (variable amounts)
- Savings transfers
- Loan repayments

**Integration Points:**
- Merchant/TPP application
- Payment initiation systems
- VRP consent database
- Payment processing systems

**User Experience Notes:**
- Clear display of payment limits
- Transparency on recurring nature
- Easy revocation (via consent management)
- UK Open Banking compliant UI

---

### Journey 10: Consent Management (View and Revoke)

**Journey Name:** Manage My Consents

**Description:** Authenticated user views all granted consents and can revoke them.

**Starting Point:**
- URL: `/consents`
- Entry method: Link from user information page or direct navigation
- Screen: `webapp/consents.html`

**User Type:** Any authenticated user who has granted consents

**Step-by-Step Flow:**

1. **Navigate to Consents Page**
   - User clicks "My Consents" button from user information page
   - Or directly navigates to `/consents`
   - **File:** `webapp/consents.html`
   - **Backend:** `code/snippet/ConsentScreen.scala`

2. **Consents Page Load**
   - System retrieves all consents for logged-in user
   - **Backend method:** `renderExistingConsents`
   - Fetches from API: `/obp/v5.1.0/my/consents`

3. **Display Consents Table**
   - User sees table with columns:
     - Consent Reference ID
     - Consumer ID (application)
     - Status (active, revoked, expired)
     - View ID (if applicable)
     - Actions (Revoke button)
   - **Dynamic loading:** Table populated via JavaScript/AJAX

4. **Review Consents**
   - User browses list of consents
   - Sees which applications have access
   - Checks consent status
   - Identifies consents they want to keep/revoke

5. **Revoke Consent (Optional)**
   - User clicks "Revoke" button for specific consent
   - **JavaScript confirmation:** May show confirm dialog
   - **Action triggers:** AJAX call to revoke API

6. **Revocation Processing**
   - System calls revoke consent API
   - **Backend method:** `callRevokeMyConsent`
   - **API endpoint:** `DELETE /obp/v5.1.0/my/consents/{consent_id}`
   - Consent status updated to revoked

7. **Update Display**
   - Consent table refreshed
   - Revoked consent removed or marked as revoked
   - Success message displayed
   - **UI update:** Via JavaScript DOM manipulation

8. **Application Impact**
   - Third-party application loses access
   - Cannot make API calls with revoked consent
   - App may prompt user to re-authorize

**Completion Point:** User has viewed consents and optionally revoked some

**Alternative Paths:**

- **No Consents Exist:**
  - Empty table displayed
  - Message: "No consents found"
  - User can navigate away

- **Revocation Fails:**
  - Error message displayed
  - Consent remains active
  - User can retry
  - **Error handling:** JavaScript shows error alert

- **API Error:**
  - Message: "Error loading consents"
  - User can refresh page
  - May need to re-login

- **Consent Already Revoked:**
  - Marked as revoked in table
  - Cannot revoke again
  - No action button shown

**Required Information:**
- User must be authenticated
- Consent ID for revocation

**Decision Points:**
- User has consents? → Show table or empty message
- User wants to revoke? → Confirm and process
- Revocation successful? → Update UI or show error

**Backend Implementation:**
- `code/snippet/ConsentScreen.scala` - Consent management logic
  - `renderExistingConsents` - Loads consent list
  - `callRevokeMyConsent` - Processes revocation
  - `updateConsentTable` - Refreshes display
- API endpoints - Consent retrieval and revocation
- Consent database - Storage and updates

**JavaScript Functionality:**
- File: `webapp/consents.html`
- AJAX calls for loading consents
- AJAX calls for revoking consents
- Dynamic table updates
- Error handling and display

**Consent Information Displayed:**
- Consent Reference ID - Unique identifier
- Consumer ID - Application that has access
- Status - Active, revoked, expired
- View ID - Specific view granted (if applicable)
- Creation date - When consent was granted
- Expiration date - When consent expires

**Integration Points:**
- OBP API consent endpoints
- Third-party applications (lose access when revoked)
- Consent database
- User authentication system

**Security Notes:**
- User can only view own consents
- User can only revoke own consents
- Revocation is immediate
- Audit trail maintained

**User Experience Notes:**
- Clear table format
- Easy identification of applications
- One-click revocation
- Immediate feedback
- No accidental revocations (confirmation)

---

## Developer Journey

### Journey 11: Consumer Registration (API Key Generation)

**Journey Name:** Developer Consumer Registration

**Description:** Developer registers for API access and obtains consumer key and secret.

**Starting Point:**
- URL: `/consumer-registration`
- Entry method: Link from home page or developer documentation
- Screen: `webapp/consumer-registration.html`

**User Type:** Developer / API consumer

**Step-by-Step Flow:**

1. **Navigate to Registration**
   - Developer clicks "Register" or "Get API Key" link
   - Navigates to consumer registration page
   - **URL:** `/consumer-registration`
   - **File:** `webapp/consumer-registration.html`
   - **Backend:** `code/snippet/ConsumerRegistration.scala`

2. **Registration Form Display**
   - Developer sees form with fields:
     - Application Name (required)
     - Application Type (dropdown, required)
     - Description (required)
     - Developer Email (required)
     - Redirect URL (optional, for OAuth)
   - **Form validation:** Client-side JavaScript

3. **Fill Application Details**
   - Developer enters application name
   - Selects application type:
     - Web Application
     - Mobile Application
     - Desktop Application
     - Other
   - Provides description of what app will do
   - Enters email for notifications

4. **OAuth Configuration (Optional)**
   - If planning to use OAuth:
     - Enter redirect URL (callback URL)
     - URL will be validated during OAuth flow
   - If using Direct Login:
     - Can skip redirect URL

5. **Submit Registration**
   - Developer clicks "Register" or "Submit" button
   - **Validation checks:**
     - All required fields present
     - Email format valid
     - Application name unique (if required)
   - Form submitted to backend

6. **Consumer Creation**
   - **Backend:** `code/snippet/ConsumerRegistration.scala`
   - System generates:
     - Consumer Key (public identifier)
     - Consumer Secret (private secret)
     - Consumer ID (internal)
   - Stores consumer record in database
   - Links to developer's user account

7. **Display Credentials**
   - **Success screen shows:**
     - Consumer Key (visible)
     - Consumer Secret (visible, one-time display)
     - Application Name
     - Status: Active
     - Warning: "Save your secret now - it won't be shown again"
   - **Email sent:** Credentials sent to developer email

8. **Save Credentials**
   - Developer must copy and save:
     - Consumer Key
     - Consumer Secret
   - **Security warning:** Secret cannot be retrieved later
   - If lost, must regenerate or create new consumer

9. **Begin Development**
   - Developer can now use credentials
   - For OAuth: Use consumer key/secret in OAuth flow
   - For Direct Login: Use to sign requests
   - Test in sandbox environment

**Completion Point:** Developer has consumer key/secret and can begin API integration

**Alternative Paths:**

- **Registration Fails:**
  - Error message displayed
  - Reasons:
     - Duplicate application name
     - Invalid email format
     - Missing required fields
     - Server error
  - Developer must correct and resubmit

- **Email Delivery Fails:**
  - Credentials still shown on screen
  - Developer can copy manually
  - May need to check spam folder
  - Can contact support for re-send

- **User Not Logged In:**
  - Redirected to login page first
  - Must authenticate before registration
  - Return to registration after login

- **Consumer Already Exists:**
  - If user already has consumer for same app name
  - Error or warning displayed
  - May need to use different name

**Required Information:**
- Application Name (required)
- Application Type (required, from dropdown)
- Description (required)
- Developer Email (required)
- Redirect URL (optional, but recommended for OAuth)

**Decision Points:**
- User logged in? → Continue or redirect to login
- All required fields filled? → Submit or show validation
- Application name unique? → Create or show error
- Email delivery successful? → Show success or warning

**Backend Implementation:**
- `code/snippet/ConsumerRegistration.scala` - Registration logic
  - Form rendering
  - Validation
  - Consumer creation
  - Credential generation
  - Email sending
- `code/consumer/Consumers` - Consumer storage
- Email service - Credential delivery

**Consumer Credentials:**
- **Consumer Key:** Public identifier, can be shared
- **Consumer Secret:** Private secret, must be kept secure
- **Consumer ID:** Internal database ID

**Security Considerations:**
- Consumer secret displayed only once
- Sent to registered email
- Must be stored securely by developer
- Cannot be retrieved later (must regenerate)
- Used for signing OAuth requests or Direct Login

**Integration Points:**
- User authentication (must be logged in)
- Consumer database (stores credentials)
- Email service (delivers credentials)
- Rate limiting (applied to consumer key)

**Email Content:**
- Contains consumer key
- Contains consumer secret
- Instructions for use
- Links to documentation
- API Explorer link

**Post-Registration Actions:**
1. Save credentials securely
2. Test in sandbox environment
3. Review API documentation
4. Test OAuth flow (if applicable)
5. Begin application development
6. Monitor usage and rate limits

**User Experience Notes:**
- Simple, clear registration form
- Immediate credential generation
- Clear security warnings
- Email backup of credentials
- Links to next steps and documentation

---

### Journey 12: Sandbox Account Creation

**Journey Name:** Create Sandbox Bank Account

**Description:** Developer creates test bank accounts in sandbox environment for testing API integrations.

**Starting Point:**
- URL: `/create-sandbox-account`
- Entry method: Link from developer documentation or API Explorer
- Screen: `webapp/create-sandbox-account.html`

**User Type:** Developer with consumer credentials

**Step-by-Step Flow:**

1. **Navigate to Sandbox Creation**
   - Developer navigates to sandbox account creation page
   - **URL:** `/create-sandbox-account`
   - **File:** `webapp/create-sandbox-account.html`
   - **Backend:** Sandbox account creation snippet

2. **Sandbox Creation Form**
   - Developer sees form with fields:
     - Bank ID (dropdown or input)
     - Account ID (generated or custom)
     - Account Type (e.g., savings, checking)
     - Currency (e.g., EUR, USD, GBP)
     - Initial Balance (optional)
     - Account Label (optional)
   - Instructions for creating test accounts

3. **Select Bank**
   - Choose bank ID from dropdown
   - Default: Sandbox test bank
   - Or use custom bank ID if configured

4. **Configure Account**
   - Select account type
   - Choose currency
   - Set initial balance (if supported)
   - Add label/description

5. **Submit Creation**
   - Developer clicks "Create Account" button
   - Form submitted to API endpoint
   - **API call:** POST to sandbox account creation endpoint

6. **Account Created**
   - System creates sandbox account
   - Generates account ID (if not provided)
   - Sets initial balance
   - Links to developer's user

7. **Display Account Details**
   - Success message shown
   - Account details displayed:
     - Account ID
     - Bank ID
     - Account Type
     - Currency
     - Balance
     - View IDs (system views created)

8. **Use in Testing**
   - Developer can now use account ID in API calls
   - Test transactions
   - Test balance queries
   - Test account operations

**Completion Point:** Sandbox account created and ready for testing

**Alternative Paths:**

- **Account Creation Fails:**
  - Error message displayed
  - Possible reasons:
     - Bank ID not found
     - Invalid parameters
     - Permission denied
     - Server error
  - Developer must correct and retry

- **Account Already Exists:**
  - If account ID already used
  - Error message shown
  - Must use different ID or delete existing

- **Invalid Bank ID:**
  - Error: Bank not found
  - Developer must use valid sandbox bank
  - May need to create bank first

- **Multiple Accounts:**
  - Developer can create multiple accounts
  - Each gets unique account ID
  - All linked to same user for testing

**Required Information:**
- Bank ID (required)
- Account type (required)
- Currency (required)
- Account ID (auto-generated or custom)

**Decision Points:**
- Bank ID valid? → Create or error
- Account ID unique? → Create or error
- User authenticated? → Continue or login
- Has required permissions? → Create or deny

**Backend Implementation:**
- Sandbox account creation logic
- Account ID generation
- Initial balance setting
- View creation (system views)
- API endpoint for account creation

**Sandbox Features:**
- Test accounts (not real money)
- Reset capability
- Multiple accounts per user
- All account types supported
- Multiple currencies
- Test transactions can be created

**Created Resources:**
- Bank account record
- System views for account
- Initial balance (if specified)
- Account holder link

**Integration Points:**
- Sandbox bank infrastructure
- Account creation API
- View creation system
- User authentication

**Testing Use Cases:**
- Test account queries
- Test transaction creation
- Test balance updates
- Test OAuth scopes
- Test permissions
- Test multi-account scenarios

**User Experience Notes:**
- Quick account creation
- Immediate availability
- Clear account details
- Ready for testing
- Can create multiple accounts

---

### Journey 13: Dummy User Tokens (Sandbox Testing)

**Journey Name:** View Dummy User Direct Login Tokens

**Description:** Developer views Direct Login tokens for dummy/sandbox users for testing purposes.

**Starting Point:**
- URL: `/dummy-user-tokens`
- Entry method: Link from sandbox documentation
- Screen: `webapp/dummy-user-tokens.html`

**User Type:** Developer in sandbox environment

**Step-by-Step Flow:**

1. **Navigate to Dummy Tokens**
   - Developer navigates to dummy tokens page
   - **URL:** `/dummy-user-tokens`
   - **File:** `webapp/dummy-user-tokens.html`
   - **Backend:** `code/snippet/ConsumerRegistration.scala` - `showDummyCustomerTokens`

2. **Dummy Tokens Display**
   - Page shows Direct Login tokens for sandbox users
   - **Information displayed:**
     - Consumer username
     - Direct Login token
     - Token components for manual signing
   - **Purpose:** Testing Direct Login without full OAuth flow

3. **Token Information**
   - Tokens are pre-generated for dummy users
   - Valid in sandbox environment only
   - Can be used immediately for testing
   - No need for token generation flow

4. **Copy Token**
   - Developer copies Direct Login token
   - Uses in API requests
   - Format: Add to `Authorization` header

5. **Test API Calls**
   - Developer makes API calls with dummy token
   - Tests Direct Login authentication
   - Validates API integration
   - No impact on production data

**Completion Point:** Developer has dummy tokens for sandbox testing

**Alternative Paths:**

- **No Dummy Tokens Available:**
  - Page shows empty state
  - May need sandbox configuration
  - Contact support or check documentation

- **Tokens Not Working:**
  - Check sandbox environment
  - Verify correct endpoint
  - Check token format
  - Review Direct Login documentation

**Required Information:**
- Access to sandbox environment
- Understanding of Direct Login authentication

**Decision Points:**
- Sandbox environment? → Show dummy tokens
- Production environment? → Do not show (security)

**Backend Implementation:**
- `code/snippet/ConsumerRegistration.scala`
- `showDummyCustomerTokens` method
- Retrieves pre-configured dummy tokens
- Displays for developer convenience

**Sandbox Only:**
- Feature not available in production
- Security measure
- Testing convenience
- Pre-generated test credentials

**Direct Login Format:**
- Token includes:
  - Consumer key
  - User token
  - Timestamp
  - Signature
- Ready to use in Authorization header

**Integration Points:**
- Sandbox user database
- Direct Login authentication
- Consumer registration system

**Testing Scenarios:**
- Quick API testing
- Direct Login validation
- Integration development
- API Explorer usage

**User Experience Notes:**
- Quick access to test tokens
- No setup required
- Immediate testing capability
- Sandbox safety
- Clear documentation

---

## Consent Management Journey

### Journey 14: Update User Auth Context

**Journey Name:** Update User Authentication Context

**Description:** User updates their authentication context details for consent and authorization purposes.

**Starting Point:**
- URL: `/user-auth-context-update`
- Entry method: Link from consent or authentication flow
- Screen: `webapp/user-auth-context-update.html`

**User Type:** Authenticated user

**Step-by-Step Flow:**

1. **Navigate to Auth Context Update**
   - User navigates to context update page
   - **URL:** `/user-auth-context-update`
   - **File:** `webapp/user-auth-context-update.html`
   - **Backend:** `code/snippet/UserAuthContextUpdate.scala`

2. **Display Current Context**
   - User sees current authentication context
   - Includes:
     - Auth method used
     - Context data
     - Timestamp
     - Related consents

3. **Update Form**
   - User can update context information
   - Fields depend on auth method
   - Submit changes

4. **Process Update**
   - System validates changes
   - Updates context record
   - Affects related consents

5. **Confirmation**
   - Success message shown
   - Updated context displayed
   - User can proceed

**Completion Point:** User auth context updated

**Backend Implementation:**
- `code/snippet/UserAuthContextUpdate.scala`
- `code/context/MappedUserAuthContext`
- `code/context/MappedUserAuthContextUpdate`

**Integration Points:**
- User authentication system
- Consent management
- Authorization tracking

---

## Banking Operations Journey

### Journey 15: OTP Validation for Payments

**Journey Name:** One-Time Password Validation

**Description:** User validates a payment transaction using OTP (One-Time Password) sent via SMS.

**Starting Point:**
- URL: `/otp`
- Entry method: Redirect during payment authorization
- Screen: `webapp/otp.html`

**User Type:** Bank customer making a payment

**Step-by-Step Flow:**

1. **Payment Initiated**
   - User initiates payment transaction via API
   - System determines OTP required
   - User redirected to OTP validation page
   - **URL:** `/otp`
   - **File:** `webapp/otp.html`

2. **OTP Page Load**
   - User sees OTP validation form
   - **File:** `webapp/otp.html`
   - **Backend:** `code/snippet/PaymentOTP.scala` - `validateOTP` method
   - Instructions displayed

3. **OTP Sent**
   - System sends OTP to user's mobile phone via SMS
   - **Backend:** SMS gateway integration
   - Message: "Your OTP is: XXXXXX"
   - Valid for limited time (e.g., 5 minutes)

4. **Display OTP Form**
   - User sees:
     - Message: "We sent an OTP to your mobile phone SMS, please check and send back it to do validate"
     - OTP input field (text box)
     - "Send OTP" button
     - "Reset" button
   - **Form ID:** `form_otp`

5. **User Receives SMS**
   - User checks mobile phone
   - Reads OTP code
   - Returns to browser

6. **Enter OTP**
   - User types OTP code into input field
   - **Field ID:** `otp_input`
   - Auto-focus on field for convenience
   - No spaces or special formatting

7. **Submit OTP**
   - User clicks "Send OTP" button
   - Form submitted via POST
   - **Backend validation:** PaymentOTP.validateOTP

8. **OTP Validation**
   - System checks:
     - OTP matches generated code
     - OTP not expired
     - OTP not already used
     - Correct format
   - **Backend:** Validates against stored OTP

9. **Validation Success**
   - OTP valid
   - Success message displayed:
     - "OTP validate success" (green background)
   - Payment transaction authorized
   - **Backend:** Marks transaction as authenticated
   - Proceed with payment processing

10. **Payment Completion**
    - Transaction processed
    - User may be redirected to confirmation page
    - Or shown next step in payment flow

**Completion Point:** OTP validated and payment authorized

**Alternative Paths:**

- **Invalid OTP:**
  - Error message: "Invalid OTP"
  - **Element:** `#otp-validation-errors` shown (red alert)
  - User can retry
  - Limited number of attempts (e.g., 3)
  - **Backend:** Tracks failed attempts

- **OTP Expired:**
  - Error message: "OTP has expired"
  - User must request new OTP
  - Payment not authorized
  - Must restart validation

- **Too Many Attempts:**
  - Error: "Too many failed attempts"
  - OTP invalidated
  - User must request new OTP
  - May need to restart payment

- **SMS Not Received:**
  - User didn't receive SMS
  - May need to:
     - Check phone signal
     - Wait a few moments
     - Request new OTP
     - Contact support
  - **Alternative:** May offer other auth methods

- **User Clicks Reset:**
  - Form cleared
  - OTP field empty
  - User can re-enter

**Required Information:**
- Valid OTP code (6 digits typically)
- Must be entered within validity period
- One-time use only

**Decision Points:**
- OTP valid? → Authorize payment or reject
- OTP expired? → Reject and request new
- Too many attempts? → Block validation
- OTP already used? → Reject

**Backend Implementation:**
- `code/snippet/PaymentOTP.scala` - OTP validation logic
  - `validateOTP` method - Validates submitted OTP
  - OTP generation - Creates random code
  - OTP storage - Temporarily stores code
- SMS gateway integration - Sends OTP
- Payment transaction tracking - Links OTP to payment

**OTP Characteristics:**
- Random numeric code (e.g., 6 digits)
- Time-limited validity (5-10 minutes typical)
- One-time use (invalidated after use)
- Linked to specific transaction
- Sent to registered mobile number

**Security Features:**
- OTP sent to verified mobile number only
- Time-limited validity
- One-time use (cannot reuse)
- Limited retry attempts
- Secure generation algorithm
- Transaction-specific (not generic)

**SMS Content Example:**
```
Your OBP API payment OTP is: 123456
Valid for 5 minutes. Do not share this code.
```

**Integration Points:**
- SMS gateway (Twilio, AWS SNS, etc.)
- Payment transaction system
- Mobile phone registration
- OTP storage (temporary)

**User Experience Notes:**
- Clear instructions
- Auto-focus on input field
- Immediate feedback on success/error
- Reset option available
- Time pressure indicated (if shown)
- Mobile-friendly interface

**Regulatory Context:**
- PSD2 SCA (Strong Customer Authentication)
- Two-factor authentication
- Secure transaction authorization
- Fraud prevention

---

## Compliance and Legal Journey

### Journey 16: Terms and Conditions Acceptance

**Journey Name:** Terms and Conditions Acceptance

**Description:** Users accept updated terms and conditions to continue using the platform.

**Starting Point:**
- Triggered during login when terms have been updated
- User may also access directly via link
- URL: `/terms-and-conditions`
- Screen: `webapp/terms-and-conditions.html`

**User Type:** All registered users

**Step-by-Step Flow:**

1. **Terms Update Detection**
   - User logs in successfully
   - System checks terms acceptance status
   - If terms updated since last acceptance, redirect to terms page
   - **Backend:** `code/snippet/TermsAndConditions.scala`

2. **Terms and Conditions Display**
   - User sees terms and conditions page
   - Full text displayed
   - Version number shown
   - Last updated date visible
   - **File:** `webapp/terms-and-conditions.html`
   - **URL:** `/terms-and-conditions`

3. **User Review**
   - User reads terms and conditions
   - Scrollable text area
   - Clear formatting for readability
   - Links to related policies (if applicable)

4. **User Decision**
   - Two options available:
     - **Accept Button:** Agree to terms
     - **Skip/Decline Button:** Postpone or decline
   - Decision recorded in system

5. **Accept Path**
   - User clicks "Accept" button
   - System records:
     - User ID
     - Terms version accepted
     - Acceptance timestamp
   - User redirected to originally requested page or home
   - Session continues normally

6. **Skip/Decline Path**
   - User clicks "Skip" or "Decline"
   - System may:
     - Allow limited access
     - Require acceptance for certain features
     - Log out user (depending on configuration)
   - User informed of consequences

**Completion Point:** Terms acceptance recorded and user can continue using platform

**Alternative Paths:**

- **Mandatory Acceptance:**
  - Some deployments require acceptance
  - Skip button may not be available
  - User cannot proceed without acceptance
  - Must accept or log out

- **Direct Access:**
  - User can access terms page anytime
  - Available from footer links
  - Available from settings/profile
  - No acceptance action if already current

**Required Information:**
- Terms version number
- User authentication (must be logged in)
- Acceptance timestamp

**Decision Points:**
- Terms updated since last login? → Triggers display
- Acceptance mandatory? → Determines if skip available
- User accepts? → Records acceptance, continues session
- User declines? → Handles according to configuration

**Backend Implementation:**
- `code/snippet/TermsAndConditions.scala` - Terms display and acceptance
- Terms version tracking
- User acceptance history

**Integration Points:**
- User authentication system
- Database (acceptance records)
- Props configuration for mandatory/optional setting

---

### Journey 17: Privacy Policy Acceptance

**Journey Name:** Privacy Policy Acceptance

**Description:** Users review and accept privacy policy to comply with data protection regulations.

**Starting Point:**
- Triggered during login when policy updated
- May be accessed directly via link
- URL: `/privacy-policy`
- Screen: `webapp/privacy-policy.html`

**User Type:** All registered users

**Step-by-Step Flow:**

1. **Policy Update Detection**
   - User logs in successfully
   - System checks privacy policy acceptance status
   - If policy updated since last acceptance, redirect to policy page
   - **Backend:** `code/snippet/PrivacyPolicy.scala`

2. **Privacy Policy Display**
   - User sees privacy policy page
   - Full policy text displayed
   - Version number shown
   - Last updated date visible
   - **File:** `webapp/privacy-policy.html`
   - **URL:** `/privacy-policy`

3. **Policy Content Review**
   - User reads privacy policy
   - Sections clearly organized:
     - Data collection practices
     - Data usage and sharing
     - User rights (GDPR, CCPA compliance)
     - Cookie usage
     - Contact information
   - Scrollable text area with clear formatting

4. **User Decision**
   - Two options available:
     - **Accept Button:** Agree to privacy policy
     - **Skip/Decline Button:** Postpone or decline
   - Decision impacts data processing

5. **Accept Path**
   - User clicks "Accept" button
   - System records:
     - User ID
     - Policy version accepted
     - Acceptance timestamp
     - Consent for data processing
   - User redirected to originally requested page or home
   - Full platform access granted

6. **Skip/Decline Path**
   - User clicks "Skip" or "Decline"
   - System may:
     - Restrict data collection
     - Limit platform features
     - Require acceptance for compliance
   - User informed of limitations

**Completion Point:** Privacy policy acceptance recorded and user can continue with appropriate data processing consent

**Alternative Paths:**

- **GDPR/Regulatory Compliance:**
  - European users may require explicit consent
  - Cannot process personal data without acceptance
  - More restrictive than other regions
  - Detailed rights information provided

- **Withdrawal of Consent:**
  - User can access policy page anytime
  - May withdraw consent (GDPR right)
  - System adjusts data processing accordingly
  - Account functionality may be affected

**Required Information:**
- Privacy policy version number
- User authentication
- Acceptance timestamp
- Geographic location (for regulatory compliance)

**Decision Points:**
- Policy updated since last login? → Triggers display
- User in regulated jurisdiction? → May require acceptance
- User accepts? → Records consent, enables full processing
- User declines? → Restricts data processing, limits features

**Backend Implementation:**
- `code/snippet/PrivacyPolicy.scala` - Policy display and acceptance
- Policy version tracking
- Consent management system
- GDPR compliance features

**Integration Points:**
- User authentication system
- Database (consent records)
- Data processing systems
- Geographic location detection
- Regulatory compliance framework

**Regulatory Context:**
- GDPR (General Data Protection Regulation)
- CCPA (California Consumer Privacy Act)
- Other regional privacy laws
- Data protection requirements
- User consent tracking
- Right to withdraw consent

---

## User Authentication Context Journey

### Journey 18: Add User Authentication Context Update Request

**Journey Name:** Initiate User Authentication Context Update

**Description:** User initiates a request to update their authentication context by providing customer identification information.

**Starting Point:**
- User navigates to auth context update page
- URL: `/user-auth-context-update/add`
- Screen: `webapp/add-user-auth-context-update-request.html`

**User Type:** Registered users requiring context update

**Step-by-Step Flow:**

1. **Navigate to Auth Context Update**
   - User accesses the auth context update page
   - **File:** `webapp/add-user-auth-context-update-request.html`
   - **URL:** `/user-auth-context-update/add`
   - **Backend:** `code/snippet/UserOnBoarding.scala`

2. **Update Request Form Display**
   - User sees form requesting:
     - Customer number
     - Bank selection (dropdown)
   - Clear instructions provided

3. **Enter Customer Information**
   - User enters customer number
   - User selects bank from dropdown
   - Form validation performed

4. **Submit Request**
   - User clicks "Submit" button
   - System validates customer number
   - **Backend:** `code/snippet/UserOnBoarding.scala`

5. **OTP Generation**
   - System generates one-time password
   - OTP sent via SMS or email
   - Request ID created

6. **Redirect to Confirmation**
   - User redirected to confirmation page
   - **Next Screen:** `confirm-user-auth-context-update-request.html`

**Completion Point:** OTP sent and user ready to confirm

**Alternative Paths:**
- Invalid customer number → Error, retry
- Customer already linked → Error, contact support
- SMS/Email delivery failure → Retry or alternative method

**Backend Implementation:**
- `code/snippet/UserOnBoarding.scala` - Context update logic
- Customer number validation
- OTP generation and delivery

**Integration Points:**
- Core banking system
- SMS gateway
- Email service

---

### Journey 19: Confirm User Authentication Context Update Request

**Journey Name:** Confirm Authentication Context Update with OTP

**Description:** User confirms the authentication context update by entering the OTP received via SMS or email.

**Starting Point:**
- Redirected from add context update request page
- URL: `/user-auth-context-update/confirm`
- Screen: `webapp/confirm-user-auth-context-update-request.html`

**User Type:** User who initiated context update request

**Step-by-Step Flow:**

1. **Arrive at Confirmation Page**
   - **File:** `webapp/confirm-user-auth-context-update-request.html`
   - **URL:** `/user-auth-context-update/confirm`
   - **Backend:** `code/snippet/UserOnBoarding.scala`

2. **Confirmation Page Display**
   - OTP input field
   - Instructions for entering code
   - Resend OTP option

3. **Receive OTP**
   - User receives SMS or email with 6-digit code

4. **Enter OTP**
   - User enters 6-digit code
   - Real-time format validation

5. **Submit for Validation**
   - User clicks "Confirm" button
   - System validates OTP
   - **Backend:** `code/snippet/UserOnBoarding.scala`

6. **Success Path**
   - OTP valid and not expired
   - System updates authentication context
   - Success message displayed

**Completion Point:** Authentication context successfully updated

**Alternative Paths:**
- Invalid OTP → Error, retry
- Expired OTP → Resend option
- Max attempts exceeded → Lock request
- OTP not received → Resend button

**Backend Implementation:**
- `code/snippet/UserOnBoarding.scala` - OTP validation
- Context linking
- Attempt limiting

**Integration Points:**
- SMS/Email services
- Core banking system
- Database (context records)

**Security Features:**
- Time-limited OTP (5-10 minutes)
- Attempt limiting (3-5 attempts)
- Request locking after max attempts

---

## Customer Management Journey

### Journey 20: API-Based Customer Operations

**Journey Name:** Customer Management via API

**Description:** Bank staff and automated systems manage customer data through API endpoints.

**Starting Point:**
- API Explorer or direct API calls
- No dedicated web UI screens
- API endpoints accessed programmatically

**User Type:** Bank administrators, customer service representatives

**Available Operations:**

1. **Create Customer**
   - API: POST `/obp/v{version}/banks/{BANK_ID}/customers`
   - Creates customer record
   - Requires customer information
   - Entitlement: CanCreateCustomer

2. **Get Customer**
   - API: GET `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}`
   - Retrieves customer details
   - Requires appropriate view permissions
   - Entitlement: CanGetCustomer

3. **Update Customer**
   - API: PUT `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}`
   - Updates customer information
   - Requires customer update entitlement
   - Entitlement: CanUpdateCustomer

4. **KYC Checks**
   - Add KYC check: POST `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_checks`
   - Get KYC checks: GET `/obp/v{version}/customers/{CUSTOMER_ID}/kyc_checks`
   - Update KYC status
   - Entitlement: CanCreateKycCheck

5. **KYC Documents**
   - Add document: POST `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}/kyc_documents`
   - Get documents: GET `/obp/v{version}/customers/{CUSTOMER_ID}/kyc_documents`
   - Manage compliance documents
   - Entitlement: CanCreateKycDocument

6. **Customer Attributes**
   - Create attribute: POST `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attributes`
   - Get attributes: GET `/obp/v{version}/banks/{BANK_ID}/customers/{CUSTOMER_ID}/attributes`
   - Manage custom customer data
   - Entitlement: CanCreateCustomerAttribute

**Access Method:**
- API Explorer (external application)
- Direct API integration
- Admin tools
- Customer service systems

**No Dedicated UI:**
- These operations are API-based only
- No HTML screens in webapp directory
- Accessed through API Explorer or integrated systems
- Requires appropriate entitlements

**Integration Points:**
- Core banking systems
- CRM systems
- Compliance systems
- KYC providers
- Document management systems

---

## Complete Screen Inventory

### Web UI Screens (HTML Files)

Total: 28 screens identified

#### 1. Home and Navigation

**1.1 Home Page**
- File: `webapp/index.html`
- URL: `/`
- Purpose: Landing page, main entry point
- Features:
  - API Explorer links
  - Login/Register links
  - Documentation links
  - Branding customization
  - Language selection
  - Cookie consent
- Backend: `code/snippet/WebUI.scala`

**1.2 Login Page**
- File: `webapp/templates-hidden/_login.html`
- URL: `/user_mgt/login`
- Purpose: User authentication
- Features:
  - Username/password form
  - OpenID Connect option
  - Forgot password link
  - Sign up link
  - Custom login instructions
- Backend: `code/snippet/Login.scala`

#### 2. OAuth and Authorization

**2.1 OAuth 1.0a Authorization**
- File: `webapp/oauth/authorize.html`
- URL: `/oauth/authorize?oauth_token=...`
- Purpose: OAuth 1.0a authorization request
- Features:
  - App name display
  - Authorization request details
  - Login form (if not authenticated)
  - Authorize/Deny buttons
- Backend: `code/snippet/OAuthAuthorisation.scala`

**2.2 OAuth Success Page**
- File: `webapp/oauth/thanks.html`
- URL: Redirect after OAuth authorization
- Purpose: OAuth authorization confirmation
- Features:
  - Success message
  - App name display
  - Redirect information
- Backend: `code/snippet/OAuthWorkedThanks.scala`

**2.3 OAuth 2.0/OIDC Consent Screen**
- File: `webapp/consent-screen.html`
- URL: `/consumer_consent?consent_challenge=...`
- Purpose: OAuth 2.0/OIDC consent management
- Features:
  - App name and description
  - Requested scopes display
  - Accept/Reject buttons
  - Hydra integration
- Backend: `code/snippet/ConsentScreen.scala`

#### 3. Consent Management

**3.1 My Consents**
- File: `webapp/consents.html`
- URL: `/consents`
- Purpose: View and manage user consents
- Features:
  - Consent list table
  - Consent details display
  - Revoke buttons
  - Dynamic loading via AJAX
- Backend: `code/snippet/ConsentScreen.scala`

**3.2 Berlin Group Consent Request**
- File: `webapp/confirm-bg-consent-request.html`
- URL: `/berlin-group/consent-request?consentId=...`
- Purpose: Berlin Group (PSD2) consent confirmation
- Features:
  - Consent details display
  - Confirm/Deny buttons
  - JSON formatted details
- Backend: `code/snippet/BerlinGroupConsent.scala`

**3.3 Berlin Group SCA**
- File: `webapp/confirm-bg-consent-request-sca.html`
- URL: `/berlin-group/consent-request-sca?consentId=...`
- Purpose: Strong Customer Authentication for Berlin Group
- Features:
  - SCA challenge display
  - Authentication method form
  - Confirm button
- Backend: `code/snippet/BerlinGroupConsent.scala`

**3.4 VRP Consent Request**
- File: `webapp/confirm-vrp-consent-request.html`
- URL: `/vrp-consent-request?...`
- Purpose: Variable Recurring Payments consent
- Features:
  - VRP details display
  - Payment limits
  - Confirm/Deny buttons
- Backend: `code/snippet/VrpConsentCreation.scala`

**3.5 VRP Consent Confirmation (OTP)**
- File: `webapp/confirm-vrp-consent.html`
- URL: `/vrp-consent?...`
- Purpose: OTP validation for VRP consent
- Features:
  - OTP input form
  - VRP consent summary
  - SMS code validation
  - Confirm button
- Backend: `code/snippet/VrpConsentCreation.scala`

**3.6 Berlin Group Consent Redirect URI**
- File: `webapp/confirm-bg-consent-request-redirect-uri.html`
- URL: `/berlin-group/consent-redirect?...`
- Purpose: TPP redirection after Berlin Group consent
- Features:
  - Consent completion message
  - App deep linking support
  - Automatic redirect to TPP app
  - Fallback manual link if auto-redirect fails
- Backend: `code/snippet/BerlinGroupConsent.scala`

#### 4. Developer Tools

**4.1 Consumer Registration**
- File: `webapp/consumer-registration.html`
- URL: `/consumer-registration`
- Purpose: Developer API key registration
- Features:
  - Application details form
  - Consumer key/secret generation
  - Email delivery of credentials
  - OAuth redirect URL configuration
- Backend: `code/snippet/ConsumerRegistration.scala`

**4.2 Create Sandbox Account**
- File: `webapp/create-sandbox-account.html`
- URL: `/create-sandbox-account`
- Purpose: Create test bank accounts for API testing
- Features:
  - Bank selection
  - Account configuration
  - Initial balance setting
  - Account creation
- Backend: Sandbox account creation logic

**4.3 Dummy User Tokens**
- File: `webapp/dummy-user-tokens.html`
- URL: `/dummy-user-tokens`
- Purpose: Display Direct Login tokens for sandbox testing
- Features:
  - Pre-generated tokens display
  - Consumer username
  - Direct Login credentials
- Backend: `code/snippet/ConsumerRegistration.scala`

#### 5. User Management

**5.1 User Invitation**
- File: `webapp/user-invitation.html`
- URL: `/user/invitation?token=...`
- Purpose: Complete user invitation registration
- Features:
  - Invitation form
  - Required field validation
  - Privacy policy acceptance
  - Terms and conditions acceptance
- Backend: `code/snippet/UserInvitation.scala`

**5.2 User Invitation Invalid**
- File: `webapp/user-invitation-invalid.html`
- URL: `/user/invitation/invalid`
- Purpose: Error page for invalid/expired invitation links
- Features:
  - Error message display
  - Explanation of invalid token
  - Contact support link
- Backend: `code/snippet/UserInvitation.scala`

**5.3 User Invitation Info**
- File: `webapp/user-invitation-info.html`
- URL: `/user/invitation/info`
- Purpose: Notice page for invitation-only access
- Features:
  - Explanation of invitation requirement
  - Instructions for obtaining invitation
  - Contact information
- Backend: `code/snippet/UserInvitation.scala`

**5.4 User Invitation Warning**
- File: `webapp/user-invitation-warning.html`
- URL: `/user/invitation/warning`
- Purpose: Warning when accessing invitation while logged in
- Features:
  - Already logged in message
  - Option to logout and use invitation
  - Continue with current session button
- Backend: `code/snippet/UserInvitation.scala`

**5.5 User Information**
- File: `webapp/user-information.html`
- URL: `/user-information`
- Purpose: Display user account information
- Features:
  - Email, username display
  - Provider information
  - Token display (OIDC)
  - Link to consent management
- Backend: `code/snippet/UserInformation.scala`

**5.6 Add User Auth Context Update Request**
- File: `webapp/add-user-auth-context-update-request.html`
- URL: `/user-auth-context-update/add`
- Purpose: Initiate user authentication context update
- Features:
  - Customer number input form
  - Bank selection
  - Request submission
  - OTP generation trigger
- Backend: `code/snippet/UserOnBoarding.scala`

**5.7 Confirm User Auth Context Update Request**
- File: `webapp/confirm-user-auth-context-update-request.html`
- URL: `/user-auth-context-update/confirm`
- Purpose: Confirm auth context update with OTP
- Features:
  - OTP input form
  - SMS/Email code validation
  - Context update confirmation
  - Success/error messages
- Backend: `code/snippet/UserOnBoarding.scala`

#### 6. Payment Operations

**6.1 OTP Validation**
- File: `webapp/otp.html`
- URL: `/otp`
- Purpose: One-Time Password validation for payments
- Features:
  - OTP input form
  - SMS code validation
  - Success/error messages
  - Reset option
- Backend: `code/snippet/PaymentOTP.scala`

#### 7. Compliance and Legal

**7.1 Terms and Conditions**
- File: `webapp/terms-and-conditions.html`
- URL: `/terms-and-conditions`
- Purpose: Display and accept terms and conditions
- Features:
  - Full terms text display
  - Accept/Skip buttons
  - Version tracking
  - Acceptance record update
- Backend: `code/snippet/TermsAndConditions.scala`

**7.2 Privacy Policy**
- File: `webapp/privacy-policy.html`
- URL: `/privacy-policy`
- Purpose: Display and accept privacy policy
- Features:
  - Privacy policy text display
  - Accept/Skip buttons
  - Version tracking
  - Acceptance record update
- Backend: `code/snippet/PrivacyPolicy.scala`

#### 8. Error and Status Pages

**8.1 Already Logged In**
- File: `webapp/already-logged-in.html`
- URL: `/user_mgt/already_logged_in`
- Purpose: Error page when attempting login while authenticated
- Features:
  - Already logged in message
  - Current user display
  - Logout button
  - Continue to home page link
- Backend: `code/snippet/WebUI.scala`

#### 9. Informational and Resource Pages

**9.1 SDKs Showcase**
- File: `webapp/sdks.html`
- URL: `/sdks` (embedded in home page)
- Purpose: Display available SDK examples and resources
- Features:
  - SDK listings (Python, Java, Node.js, etc.)
  - GitHub repository links
  - Code examples
  - Integration instructions
  - Embedded via `data-lift="embed?what=sdks"`
- Backend: `code/snippet/WebUI.scala`

**9.2 Main FAQ**
- File: `webapp/main-faq.html`
- URL: `/faq` (embedded in home page)
- Purpose: Frequently Asked Questions about API usage
- Features:
  - Collapsible Q&A sections
  - Getting started information
  - Authentication help
  - API usage guidance
  - Embedded via `data-lift="embed?what=main-faq"`
- Backend: `code/snippet/WebUI.scala`

**9.3 Introduction/API Documentation**
- File: `webapp/introduction.html`
- URL: `/introduction`
- Purpose: API introduction and getting started guide
- Features:
  - Dynamic content loading
  - API overview
  - Quick start guide
  - Documentation links
  - Loaded via `WebUI.apiDocumentation`
- Backend: `code/snippet/WebUI.scala`

---

### Backend Scala Snippets

#### Snippet Classes

1. **WebUI.scala** (`code/snippet/WebUI.scala`)
   - Purpose: Main web UI rendering and customization
   - Features:
     - Home page rendering
     - Branding customization
     - Language selection
     - Cookie consent
     - API Explorer links
     - Documentation links

2. **Login.scala** (`code/snippet/Login.scala`)
   - Purpose: Login/logout functionality
   - Features:
     - Login state checking
     - Logout handling
     - SSO integration
     - Custom login instructions

3. **OAuthAuthorisation.scala** (`code/snippet/OAuthAuthorisation.scala`)
   - Purpose: OAuth 1.0a authorization flow
   - Features:
     - Token validation
     - Authorization request handling
     - Verifier generation
     - Callback URL management

4. **ConsentScreen.scala** (`code/snippet/ConsentScreen.scala`)
   - Purpose: Consent management (OAuth 2.0/OIDC)
   - Features:
     - Consent request handling
     - Hydra integration
     - Consent approval/denial
     - Consent list display
     - Consent revocation

5. **OpenIDConnectSnippet.scala** (`code/snippet/OpenIDConnectSnippet.scala`)
   - Purpose: OpenID Connect integration
   - Features:
     - OIDC authentication
     - Token exchange
     - User mapping

6. **ConsumerRegistration.scala** (`code/snippet/ConsumerRegistration.scala`)
   - Purpose: Developer consumer registration
   - Features:
     - Consumer key generation
     - Application registration
     - Email delivery
     - Dummy token display

7. **UserInvitation.scala** (`code/snippet/UserInvitation.scala`)
   - Purpose: User invitation handling
   - Features:
     - Invitation token validation
     - User account creation
     - Entitlement assignment

8. **BerlinGroupConsent.scala** (`code/snippet/BerlinGroupConsent.scala`)
   - Purpose: Berlin Group consent handling
   - Features:
     - Consent confirmation
     - SCA integration
     - PSD2 compliance

9. **VrpConsentCreation.scala** (`code/snippet/VrpConsentCreation.scala`)
   - Purpose: VRP consent creation
   - Features:
     - VRP consent handling
     - Payment limit validation
     - UK Open Banking compliance

10. **PaymentOTP.scala** (`code/snippet/PaymentOTP.scala`)
    - Purpose: OTP validation
    - Features:
      - OTP generation
      - SMS sending
      - OTP validation
      - Payment authorization

11. **UserInformation.scala** (`code/snippet/UserInformation.scala`)
    - Purpose: User information display
    - Features:
      - User data retrieval
      - Token display
      - Account information

12. **UserAuthContextUpdate.scala** (`code/snippet/UserAuthContextUpdate.scala`)
    - Purpose: Auth context management
    - Features:
      - Context update
      - Consent linking

---

## Integration Points

### External Systems

#### 1. ORY Hydra
- **Purpose:** Identity and consent management for OAuth 2.0/OIDC
- **Integration:**
  - Consent challenge validation
  - Consent approval/rejection
  - Token generation
- **Files:**
  - `code/snippet/ConsentScreen.scala`
  - `code/util/HydraUtil.scala`
- **Endpoints:**
  - GET consent request details
  - POST consent acceptance
  - POST consent rejection

#### 2. API Explorer
- **Purpose:** Interactive API documentation and testing
- **Type:** External web application (separate repository)
- **Integration:**
  - Configured via `webui_api_explorer_url` property
  - Links from home page
  - API documentation reference
- **Features:**
  - Browse endpoints
  - Test API calls
  - View examples
  - Glossary access

#### 3. Email Service
- **Purpose:** Send transactional emails
- **Use Cases:**
  - User verification emails
  - Password reset emails
  - Consumer credential delivery
  - User invitation emails
  - OAuth notifications
- **Configuration:** Via props (SMTP settings)

#### 4. SMS Gateway
- **Purpose:** Send OTP codes for payment authentication
- **Use Cases:**
  - Payment OTP delivery
  - SCA authentication codes
- **Integration:** Via SMS provider (Twilio, AWS SNS, etc.)

#### 5. OpenID Connect Providers
- **Purpose:** External authentication
- **Providers:**
  - Keycloak
  - Auth0
  - Generic OIDC providers
- **Integration:**
  - `code/snippet/OpenIDConnectSnippet.scala`
  - Standard OIDC protocol

#### 6. Core Banking Systems
- **Purpose:** Actual banking operations
- **Integration:**
  - Via connector layer
  - Configured via `connector` property
  - Multiple connector types supported

---

### Internal Systems

#### 1. Database (PostgreSQL/MySQL/H2)
- **Purpose:** Data persistence
- **Storage:**
  - User accounts (AuthUser, ResourceUser)
  - Consumers (API keys)
  - Tokens (OAuth, Direct Login)
  - Consents
  - Transactions
  - Entitlements
  - Audit logs

#### 2. Session Management
- **Purpose:** User authentication state
- **Framework:** Lift web framework
- **Features:**
  - Session creation
  - Session validation
  - Session timeout
  - Logout handling

#### 3. Entitlement System
- **Purpose:** Role-based access control
- **Features:**
  - Permission checking
  - Role assignment
  - API endpoint protection
- **Storage:** `code/entitlement/MappedEntitlement`

#### 4. Rate Limiting
- **Purpose:** API usage throttling
- **Features:**
  - Per-consumer limits
  - Per-user limits
  - Anonymous limits
- **Implementation:** `code/ratelimiting/RateLimiting`

#### 5. Webhook System
- **Purpose:** Event notifications
- **Features:**
  - Account balance changes
  - Transaction notifications
  - Consent events
- **Storage:** `code/webhook/MappedAccountWebhook`

---

## Flow Diagrams

### High-Level User Journey Map

```
┌─────────────────────────────────────────────────────────────────┐
│                        Entry Points                              │
├─────────────────────────────────────────────────────────────────┤
│  Home Page  │  Direct URL  │  Email Link  │  API Explorer      │
└──────┬──────┴──────┬───────┴──────┬───────┴─────────┬───────────┘
       │             │              │                 │
       └─────────────┴──────────────┴─────────────────┘
                            │
                    ┌───────▼────────┐
                    │  User Action   │
                    └───────┬────────┘
                            │
       ┌────────────────────┼────────────────────┐
       │                    │                    │
┌──────▼───────┐   ┌────────▼────────┐   ┌──────▼──────────┐
│   Sign Up    │   │     Login       │   │   Guest Browse  │
│              │   │                 │   │                 │
└──────┬───────┘   └────────┬────────┘   └─────────────────┘
       │                    │
       │                    │
       └──────────┬─────────┘
                  │
          ┌───────▼────────┐
          │  Authenticated │
          │      User      │
          └───────┬────────┘
                  │
    ┌─────────────┼─────────────┐
    │             │             │
┌───▼────┐  ┌─────▼─────┐  ┌───▼─────┐
│Developer│  │  End User │  │  Admin  │
│ Journey │  │  Journey  │  │ Journey │
└────┬────┘  └─────┬─────┘  └────┬────┘
     │             │              │
     │             │              │
┌────▼──────────┐  │         ┌────▼──────────┐
│ 1. Register   │  │         │ Manage System │
│    Consumer   │  │         │ via API calls │
│ 2. Get API    │  │         └───────────────┘
│    Keys       │  │
│ 3. Create     │  │
│    Sandbox    │  │
│ 4. Test APIs  │  │
└───────────────┘  │
                   │
              ┌────▼──────────┐
              │ 1. Authorize  │
              │    OAuth Apps │
              │ 2. Manage     │
              │    Consents   │
              │ 3. Banking    │
              │    via Apps   │
              └───────────────┘
```

### OAuth 1.0a Flow Diagram

```
Third-Party App          User Browser         OBP API
     │                        │                  │
     │  1. Get Request Token  │                  │
     ├───────────────────────────────────────────►│
     │◄──────────────────────────────────────────┤
     │        Request Token                      │
     │                        │                  │
     │  2. Redirect to Auth   │                  │
     ├───────────────────────►│                  │
     │                        │  3. Load Auth    │
     │                        │     Page         │
     │                        ├─────────────────►│
     │                        │◄─────────────────┤
     │                        │  Auth Page HTML  │
     │                        │                  │
     │                        │  4. User Login   │
     │                        │     (if needed)  │
     │                        ├─────────────────►│
     │                        │◄─────────────────┤
     │                        │                  │
     │                        │  5. User         │
     │                        │     Authorizes   │
     │                        ├─────────────────►│
     │                        │◄─────────────────┤
     │                        │  Verifier Code   │
     │                        │                  │
     │  6. Callback with      │                  │
     │     Verifier           │                  │
     │◄───────────────────────┤                  │
     │                        │                  │
     │  7. Exchange for       │                  │
     │     Access Token       │                  │
     ├───────────────────────────────────────────►│
     │◄──────────────────────────────────────────┤
     │        Access Token                       │
     │                        │                  │
     │  8. Make API Calls     │                  │
     ├───────────────────────────────────────────►│
     │◄──────────────────────────────────────────┤
     │        API Response                       │
```

### OAuth 2.0/OIDC with Hydra Flow

```
Third-Party App    User Browser    OBP API       Hydra
     │                  │             │            │
     │  1. Auth Request │             │            │
     ├─────────────────►│             │            │
     │                  │  2. Forward │            │
     │                  ├────────────►│            │
     │                  │             │  3. Create │
     │                  │             │    Consent │
     │                  │             │    Challenge│
     │                  │             ├───────────►│
     │                  │             │◄───────────┤
     │                  │             │  Challenge │
     │                  │             │            │
     │                  │  4. Redirect│            │
     │                  │     to OBP  │            │
     │                  │◄────────────┤            │
     │                  │  Consent    │            │
     │                  │  Screen     │            │
     │                  │             │            │
     │                  │  5. User    │            │
     │                  │     Reviews │            │
     │                  │     & Accepts│           │
     │                  ├────────────►│            │
     │                  │             │  6. Accept │
     │                  │             │    Consent │
     │                  │             ├───────────►│
     │                  │             │◄───────────┤
     │                  │             │  Redirect  │
     │                  │             │  URL       │
     │                  │  7. Redirect│            │
     │                  │◄────────────┤            │
     │                  │  to Hydra   │            │
     │                  ├────────────────────────►│
     │                  │◄────────────────────────┤
     │  8. Callback     │  Auth Code              │
     │◄─────────────────┤             │            │
     │                  │             │            │
     │  9. Exchange     │             │            │
     │     Code for     ├────────────────────────►│
     │     Tokens       │             │            │
     │◄────────────────────────────────────────────┤
     │  Access Token,   │             │            │
     │  ID Token        │             │            │
     │                  │             │            │
     │  10. API Calls   │             │            │
     ├──────────────────────────────►│            │
     │◄────────────────────────────────            │
     │  API Response    │             │            │
```

### Developer Registration Flow

```
Developer Browser         OBP API
      │                      │
      │  1. Navigate to      │
      │     /consumer-       │
      │     registration     │
      ├─────────────────────►│
      │◄─────────────────────┤
      │  Registration Form   │
      │                      │
      │  2. Fill Form:       │
      │     - App Name       │
      │     - App Type       │
      │     - Description    │
      │     - Email          │
      │     - Redirect URL   │
      ├─────────────────────►│
      │                      │
      │  3. System Creates:  │
      │     - Consumer Key   │
      │     - Consumer Secret│
      │◄─────────────────────┤
      │  Credentials Display │
      │                      │
      │  4. Email Sent:      │
      │     - Consumer Key   │
      │     - Consumer Secret│
      │◄─────────────────────┤
      │                      │
      │  5. Save Credentials │
      │     (One-time view)  │
      │                      │
      │  6. Begin API        │
      │     Development      │
      │     - Test in        │
      │       Sandbox        │
      │     - Integrate      │
      │     - Production     │
```

### Consent Management Flow

```
User Browser             OBP API
     │                      │
     │  1. Login            │
     ├─────────────────────►│
     │◄─────────────────────┤
     │  Authenticated       │
     │                      │
     │  2. Navigate to      │
     │     /consents        │
     ├─────────────────────►│
     │◄─────────────────────┤
     │  Consents List       │
     │                      │
     │  3. View Consents:   │
     │     ┌──────────────┐ │
     │     │ Consent 1    │ │
     │     │ App: FinApp  │ │
     │     │ Status: Active│ │
     │     │ [Revoke]     │ │
     │     ├──────────────┤ │
     │     │ Consent 2    │ │
     │     │ App: BudgetX │ │
     │     │ Status: Active│ │
     │     │ [Revoke]     │ │
     │     └──────────────┘ │
     │                      │
     │  4. Click [Revoke]   │
     │     on Consent 1     │
     ├─────────────────────►│
     │                      │
     │  5. Confirm Revoke   │
     ├─────────────────────►│
     │                      │
     │  6. Consent Revoked  │
     │◄─────────────────────┤
     │  Updated List:       │
     │     ┌──────────────┐ │
     │     │ Consent 2    │ │
     │     │ App: BudgetX │ │
     │     │ Status: Active│ │
     │     │ [Revoke]     │ │
     │     └──────────────┘ │
     │                      │
     │  7. App "FinApp"     │
     │     Loses Access     │
```

### Payment OTP Flow

```
User Browser         OBP API         SMS Gateway
     │                  │                 │
     │  1. Initiate     │                 │
     │     Payment      │                 │
     ├─────────────────►│                 │
     │                  │  2. Generate    │
     │                  │     OTP         │
     │                  │  3. Send SMS    │
     │                  ├────────────────►│
     │                  │                 │
     │  4. Redirect     │                 │  5. SMS Delivered
     │     to /otp      │                 │     "Your OTP:
     │◄─────────────────┤                 │      123456"
     │  OTP Form        │                 ├────────►
     │                  │                 │  User's Phone
     │  6. User Receives│                 │
     │     SMS & Enters │                 │
     │     OTP: 123456  │                 │
     ├─────────────────►│                 │
     │                  │  7. Validate    │
     │                  │     OTP         │
     │  8. Success or   │                 │
     │     Error        │                 │
     │◄─────────────────┤                 │
     │                  │                 │
     │  If Success:     │                 │
     │  9. Process      │                 │
     │     Payment      │                 │
```

---

## Alternative Paths and Error Handling

### Common Error Scenarios

#### 1. Authentication Errors

**Scenario: Invalid Credentials**
- **Where:** Login page
- **Trigger:** Wrong username or password
- **User sees:** "Invalid username or password"
- **System action:**
  - Increment bad login attempt counter
  - Show error message
  - Keep user on login page
- **Recovery:**
  - User can retry with correct credentials
  - After max attempts, account locked
  - User can use "Forgot Password" link

**Scenario: Account Locked**
- **Where:** Login page
- **Trigger:** Too many failed login attempts
- **User sees:** "Account locked due to too many failed attempts"
- **System action:**
  - Block login attempts
  - Require unlock (time-based or manual)
- **Recovery:**
  - Wait for auto-unlock period
  - Contact support for manual unlock
  - Use alternative authentication method

**Scenario: Account Not Verified**
- **Where:** Login page
- **Trigger:** User hasn't verified email
- **User sees:** "Please verify your email address"
- **System action:**
  - Block login
  - Offer to resend verification email
- **Recovery:**
  - Check email and verify
  - Request new verification email
  - Contact support if email not received

#### 2. Authorization Errors

**Scenario: Invalid OAuth Token**
- **Where:** OAuth authorization page
- **Trigger:** Expired or invalid request token
- **User sees:** "Invalid or expired token"
- **System action:**
  - Show error message
  - Cannot proceed with authorization
- **Recovery:**
  - Application must restart OAuth flow
  - User returns to application
  - User initiates authorization again

**Scenario: Consent Challenge Invalid**
- **Where:** OAuth 2.0 consent screen
- **Trigger:** Invalid or expired Hydra challenge
- **User sees:** Error message about invalid consent request
- **System action:**
  - Cannot display consent screen
  - No consent granted
- **Recovery:**
  - Application must restart OAuth flow
  - User redirected back to application

**Scenario: User Denies Authorization**
- **Where:** OAuth authorization or consent screen
- **Trigger:** User clicks "Deny" button
- **User sees:** Redirect back to application
- **System action:**
  - No authorization granted
  - Redirect with error parameter
- **Application sees:** Error indicating access denied
- **Recovery:**
  - Application handles denial gracefully
  - User can retry authorization if desired

#### 3. Registration Errors

**Scenario: Username Already Taken**
- **Where:** Sign-up or invitation form
- **Trigger:** Username exists in database
- **User sees:** "Username already taken"
- **System action:**
  - Form submission rejected
  - User remains on form
- **Recovery:**
  - Choose different username
  - Resubmit form

**Scenario: Email Already Registered**
- **Where:** Sign-up form
- **Trigger:** Email already in use
- **User sees:** "Email already registered"
- **System action:**
  - Form submission rejected
  - Suggest password recovery
- **Recovery:**
  - Use password recovery if forgot
  - Use different email
  - Contact support if error

**Scenario: Password Requirements Not Met**
- **Where:** Sign-up or password change
- **Trigger:** Weak password
- **User sees:** Password requirements message
- **System action:**
  - Form submission rejected
  - Display requirements
- **Recovery:**
  - Use stronger password
  - Meet all requirements
  - Resubmit

#### 4. Consumer Registration Errors

**Scenario: Application Name Duplicate**
- **Where:** Consumer registration form
- **Trigger:** Name already used
- **User sees:** "Application name already exists"
- **System action:**
  - Registration rejected
- **Recovery:**
  - Choose different name
  - Resubmit registration

**Scenario: Invalid Email Format**
- **Where:** Consumer registration
- **Trigger:** Malformed email
- **User sees:** "Invalid email format"
- **System action:**
  - Form validation error
- **Recovery:**
  - Correct email format
  - Resubmit

#### 5. Payment/OTP Errors

**Scenario: Invalid OTP**
- **Where:** OTP validation page
- **Trigger:** Wrong OTP code entered
- **User sees:** "Invalid OTP"
- **System action:**
  - Increment failed attempt counter
  - Payment not authorized
- **Recovery:**
  - Re-enter correct OTP
  - Limited retry attempts
  - Request new OTP if needed

**Scenario: OTP Expired**
- **Where:** OTP validation page
- **Trigger:** Time limit exceeded
- **User sees:** "OTP has expired"
- **System action:**
  - Reject OTP
  - Payment not authorized
- **Recovery:**
  - Request new OTP
  - Restart payment flow

**Scenario: SMS Not Received**
- **Where:** OTP validation page
- **Trigger:** SMS delivery failure
- **User sees:** Waiting for OTP that never arrives
- **System action:**
  - OTP generated but not delivered
- **Recovery:**
  - Check phone signal
  - Wait a few minutes
  - Request new OTP
  - Contact support
  - Try alternative authentication

#### 6. Consent Management Errors

**Scenario: Consent Revocation Fails**
- **Where:** Consent management page
- **Trigger:** API error during revocation
- **User sees:** "Error revoking consent"
- **System action:**
  - Consent remains active
  - Error logged
- **Recovery:**
  - Retry revocation
  - Refresh page
  - Contact support if persists

**Scenario: Cannot Load Consents**
- **Where:** Consent management page
- **Trigger:** API error or network issue
- **User sees:** "Error loading consents"
- **System action:**
  - Empty table or error message
- **Recovery:**
  - Refresh page
  - Check network connection
  - Re-login if session expired

#### 7. Session Errors

**Scenario: Session Timeout**
- **Where:** Any authenticated page
- **Trigger:** Inactivity timeout
- **User sees:** Redirect to login with message
- **System action:**
  - Session invalidated
  - Redirect to login page
- **Recovery:**
  - Log in again
  - Continue from where left off (if state preserved)

**Scenario: Concurrent Session Conflict**
- **Where:** Any page while logged in
- **Trigger:** Login from another device/browser
- **User sees:** "Your session is no longer valid"
- **System action:**
  - Current session invalidated
- **Recovery:**
  - Log in again on current device
  - Or use other active session

#### 8. Integration Errors

**Scenario: Hydra Communication Failure**
- **Where:** OAuth 2.0 consent screen
- **Trigger:** Hydra server unavailable
- **User sees:** "Unable to process consent"
- **System action:**
  - Cannot complete consent flow
  - Error logged
- **Recovery:**
  - Retry later
  - Contact support
  - Use alternative auth method

**Scenario: Email Service Unavailable**
- **Where:** Sign-up, password recovery
- **Trigger:** SMTP server down
- **User sees:** Warning about email delivery
- **System action:**
  - Account created but no email sent
  - Credentials shown on screen
- **Recovery:**
  - Copy credentials from screen
  - Contact support for email resend
  - Check spam folder

**Scenario: SMS Gateway Failure**
- **Where:** OTP validation
- **Trigger:** SMS service unavailable
- **User sees:** OTP not received
- **System action:**
  - OTP generated but not delivered
- **Recovery:**
  - Try alternative auth method
  - Contact support
  - Retry later

### Error Recovery Patterns

#### Pattern 1: Retry with Correction
- User sees error
- Error message explains issue
- User corrects input
- User resubmits
- Success

#### Pattern 2: Alternative Path
- User encounters error
- System offers alternative method
- User chooses alternative
- Completes task via different route

#### Pattern 3: Support Escalation
- User encounters persistent error
- Multiple retry attempts fail
- System provides support contact
- User contacts support
- Support resolves issue

#### Pattern 4: Session Reset
- User encounters session-related error
- System requires re-authentication
- User logs in again
- User continues task

#### Pattern 5: Flow Restart
- Error prevents continuation
- System requires starting over
- User initiates flow from beginning
- Completes successfully

---

## Summary and Key Takeaways

### Screen Flow Categories

1. **Entry and Landing (2 screens)**
   - Home page
   - Login page

2. **OAuth and Authorization (5 screens)**
   - OAuth 1.0a authorization
   - OAuth success page
   - OAuth 2.0/OIDC consent
   - Berlin Group consent (2 screens)

3. **Consent Management (4 screens)**
   - My consents
   - VRP consent (2 screens)
   - User auth context update

4. **Developer Tools (3 screens)**
   - Consumer registration
   - Sandbox account creation
   - Dummy user tokens

5. **User Management (2 screens)**
   - User invitation
   - User information

6. **Payment Operations (1 screen)**
   - OTP validation

7. **Compliance and Legal (2 screens)**
   - Terms and Conditions acceptance
   - Privacy Policy acceptance

8. **User Authentication Context (2 screens)**
   - Add auth context update request
   - Confirm auth context update request

9. **Error and Status Pages (4 screens)**
   - User invitation invalid
   - User invitation info
   - User invitation warning
   - Already logged in

10. **Informational/Resource Pages (3 screens)**
    - SDKs showcase
    - FAQ page
    - Introduction/API documentation

**Total: 28 Web UI Screens**

### Authentication Methods Summary

1. **OAuth 1.0a** - Traditional three-legged OAuth
2. **OAuth 2.0/OIDC** - Modern OAuth with Hydra
3. **Direct Login** - Header-based authentication
4. **Gateway Login** - API gateway integration
5. **OpenID Connect** - External identity providers

### User Type Summary

1. **Developers** - API integration and testing
2. **Bank Customers** - Banking operations via apps
3. **Third-Party Apps** - OAuth consumers
4. **Bank Administrators** - System management (API-based)
5. **Invited Users** - Onboarding via invitation
6. **Compliance Officers** - Audit and compliance (API-based)

### Key Integration Points

1. **ORY Hydra** - Identity and consent management
2. **API Explorer** - Interactive documentation (external)
3. **Email Service** - Transactional emails
4. **SMS Gateway** - OTP delivery
5. **OIDC Providers** - External authentication
6. **Core Banking** - Backend systems

### Important Notes

#### API Explorer Limitation
The API Explorer is an **external web application** configured via the `webui_api_explorer_url` property. Its internal screens and flows are not part of this repository and therefore not documented in detail here. The API Explorer provides:
- Interactive API documentation
- API endpoint testing interface
- Request/response examples
- Glossary and documentation
- OAuth testing tools

#### API-Based Operations
Several administrative and customer management operations do not have dedicated web UI screens but are instead accessed via API endpoints:
- Customer management (CRUD operations)
- KYC checks and documents
- Customer attributes
- Entitlement management
- System administration
- Account management (beyond sandbox creation)

These operations are performed through:
- API Explorer (external application)
- Direct API integration
- Admin tools and scripts
- Customer service systems

#### Regulatory Compliance
Multiple flows are designed for regulatory compliance:
- **PSD2 (EU):** Berlin Group consent flows with SCA
- **UK Open Banking:** VRP consent creation
- **GDPR:** Privacy policy acceptance, consent management, data access
- **AML/KYC:** KYC document and check management (API-based)

#### Security Features
All flows incorporate security best practices:
- Session management and timeout
- CSRF protection
- Rate limiting
- Audit logging
- Secure token storage
- Password hashing
- Account locking after failed attempts
- Two-factor authentication (OTP, SCA)

---

## Document Maintenance

### When to Update This Document

- New web UI screens added
- Authentication methods changed or added
- User journey modifications
- New regulatory requirements
- Integration point changes
- Backend implementation updates

### Version History

- **Version 1.0** - October 12, 2025 - Initial comprehensive documentation
  - Documented 28 web UI screens
  - 5 authentication methods
  - 6 user types
  - Multiple consent flows
  - Integration points
  - Flow diagrams
  - Compliance and legal journeys
  - User authentication context flows
  - Complete error handling screens

### Contact and Support

For questions about screen flows or user journeys:
- Review API documentation
- Check API Explorer (external)
- Contact: OBP community forums

---

## Appendix: Technical Implementation Details

### Backend Snippet Classes Location

All Scala snippet classes located at:
```
obp-api/src/main/scala/code/snippet/
```

Key files:
- `WebUI.scala` - Main UI customization
- `Login.scala` - Authentication
- `OAuthAuthorisation.scala` - OAuth 1.0a
- `ConsentScreen.scala` - Consent management
- `OpenIDConnectSnippet.scala` - OIDC
- `ConsumerRegistration.scala` - Developer registration
- `UserInvitation.scala` - User invitations
- `BerlinGroupConsent.scala` - PSD2 compliance
- `VrpConsentCreation.scala` - UK Open Banking
- `PaymentOTP.scala` - Payment authentication
- `UserInformation.scala` - User data display
- `UserAuthContextUpdate.scala` - Auth context management

### Configuration Properties

Key properties that control screen flows:
- `webui_api_explorer_url` - API Explorer URL
- `webui_login_page_text` - Login page customization
- `webui_signup_body_password_repeat_text` - Signup text
- `webui_privacy_policy` - Privacy policy link
- `webui_terms_and_conditions` - Terms and conditions link
- `authUser.skipEmailValidation` - Skip email verification
- `sso.enabled` - SSO integration
- `integrate_with_keycloak` - Keycloak integration
- `allow_pre_filled_password` - Development convenience
- `create_system_views_at_boot` - System view creation
- Various `webui_*` properties for branding and customization

### Database Tables

Key tables for screen flows:
- `authusers` - User authentication
- `resourceusers` - User resources
- `consumers` - API consumers
- `tokens` - OAuth tokens
- `nonces` - OAuth nonces
- `consents` - User consents
- `accountholders` - Account holders
- `viewprivileges` - View permissions
- `entitlements` - User entitlements
- `badloginattempts` - Failed logins
- `userlocks` - Account locks

### API Endpoints Referenced

Key API endpoints used by web UI:
- `/obp/v5.1.0/my/consents` - Get user consents
- `/obp/v5.1.0/my/consents/{CONSENT_ID}` - Revoke consent
- `/obp/v{VERSION}/banks/{BANK_ID}/customers` - Customer operations
- `/obp/v{VERSION}/users/current` - Current user info
- Various OAuth endpoints (`/oauth/initiate`, `/oauth/authorize`, `/oauth/token`)
- Hydra integration endpoints (external)

---

**End of Screen Flow Documentation**

This document provides comprehensive coverage of all web-based user journeys and screen flows in the Open Bank Project API. For API-based operations, please refer to the API documentation available through the API Explorer.
