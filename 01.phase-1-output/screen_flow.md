# OBP-API Screen Flow Analysis - 16-9-2025

This document maps all user-facing navigation flows in the Open Bank Project API web application.

## Screen Flow Summary

The OBP-API uses the Lift web framework with a sitemap-based navigation system. The main navigation flows include authentication, OAuth authorization, consumer registration, user onboarding, consent management, and various administrative functions.

## Screen Flows

### 1. Home/Landing Screen
- **Screen ID**: HOME
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/index.html`
- **Trigger**: Direct URL access to `/` or `/index`
- **Navigates To**: 
  - LOGIN (via login links)
  - CONSUMER_REGISTRATION (via "Register your app" links)
  - API_EXPLORER (via API Explorer links)
  - INTRODUCTION (via introduction links)
- **Condition/Notes**: Main landing page with navigation to various sections

### 2. Login Screen
- **Screen ID**: LOGIN
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/templates-hidden/_login.html`
- **Trigger**: Click login links, authentication required redirects
- **Navigates To**:
  - HOME (on successful login via POST `/user_mgt/login`)
  - PASSWORD_RESET (via "Forgotten password?" link to `/user_mgt/lost_password`)
  - SIGNUP (via "Register" link to `/user_mgt/sign_up`)
  - OPENID_CONNECT (via OpenID Connect buttons)
- **Condition/Notes**: Form submission to `/user_mgt/login` endpoint

### 3. OAuth Authorization Screen
- **Screen ID**: OAUTH_AUTHORIZE
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/oauth/authorize.html`
- **Trigger**: OAuth flow initiation via `/oauth/authorize`
- **Navigates To**:
  - OAUTH_THANKS (on successful authorization with callback URL)
  - LOGIN (if user not authenticated)
  - VERIFIER_DISPLAY (for out-of-band OAuth flows)
- **Condition/Notes**: Handles OAuth 1.0a authorization with token validation

### 4. OAuth Thanks/Redirect Screen
- **Screen ID**: OAUTH_THANKS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/snippet/OAuthAuthorisation.scala` (OAuthWorkedThanks.menu)
- **Trigger**: Successful OAuth authorization
- **Navigates To**: External application (via redirect URL)
- **Condition/Notes**: Redirects to application callback URL with OAuth verifier

### 5. Consumer Registration Screen
- **Screen ID**: CONSUMER_REGISTRATION
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/consumer-registration.html`
- **Trigger**: Access to `/consumer-registration` (requires login)
- **Navigates To**:
  - CONSUMER_REGISTRATION_SUCCESS (on successful form submission)
  - CONSUMER_REGISTRATION (on validation errors)
- **Condition/Notes**: Form submission creates OAuth consumer credentials

### 6. Consumer Registration Success Screen
- **Screen ID**: CONSUMER_REGISTRATION_SUCCESS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/snippet/ConsumerRegistration.scala` (showResults method)
- **Trigger**: Successful consumer registration
- **Navigates To**: 
  - DUMMY_USER_TOKENS (via "Get dummy users' token" link)
  - External links (API documentation, etc.)
- **Condition/Notes**: Displays consumer key, secret, and OAuth endpoints

### 7. User Invitation Screen
- **Screen ID**: USER_INVITATION
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/user-invitation.html`
- **Trigger**: Access via invitation link with secret ID parameter
- **Navigates To**:
  - PASSWORD_RESET (on successful registration via redirect)
  - USER_INVITATION_INVALID (on invalid/expired invitation)
  - USER_INVITATION_WARNING (if already logged in)
- **Condition/Notes**: Validates invitation TTL and creates user account

### 8. User Invitation Invalid Screen
- **Screen ID**: USER_INVITATION_INVALID
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/user-invitation-invalid.html`
- **Trigger**: Invalid or expired invitation link
- **Navigates To**: HOME (via navigation)
- **Condition/Notes**: Error state for invalid invitations

### 9. User Invitation Warning Screen
- **Screen ID**: USER_INVITATION_WARNING
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/user-invitation-warning.html`
- **Trigger**: Accessing invitation while already logged in
- **Navigates To**: HOME (via navigation)
- **Condition/Notes**: Warning for already authenticated users

### 10. Consent Screen
- **Screen ID**: CONSENT_SCREEN
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/consent-screen.html`
- **Trigger**: OAuth consent flow via `/consent-screen` (requires login)
- **Navigates To**:
  - OAUTH_REDIRECT (on "Allow access" button)
  - HOME (on "Deny access" button)
- **Condition/Notes**: Hydra integration for OAuth2/OpenID Connect consent

### 11. Consents Management Screen
- **Screen ID**: CONSENTS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/consents.html`
- **Trigger**: Access to `/consents`
- **Navigates To**: CONSENTS (AJAX refresh on consent revocation)
- **Condition/Notes**: Table view of user consents with revoke functionality

### 12. Berlin Group Consent Request Screen
- **Screen ID**: BG_CONSENT_REQUEST
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/confirm-bg-consent-request.html`
- **Trigger**: Access to `/confirm-bg-consent-request` with CONSENT_ID parameter (requires login)
- **Navigates To**:
  - EXTERNAL_REDIRECT (on "Confirm" - redirects to TPP)
  - HOME (on "Deny")
- **Condition/Notes**: Berlin Group PSD2 consent confirmation with account selection

### 13. Berlin Group Consent SCA Screen
- **Screen ID**: BG_CONSENT_SCA
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/confirm-bg-consent-request-sca.html`
- **Trigger**: Strong Customer Authentication required for Berlin Group consent
- **Navigates To**: EXTERNAL_REDIRECT (on successful SCA)
- **Condition/Notes**: Additional authentication step for PSD2 compliance

### 14. OTP Validation Screen
- **Screen ID**: OTP_VALIDATION
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/otp.html`
- **Trigger**: Access to `/otp` with flow parameter (requires login)
- **Navigates To**:
  - OTP_SUCCESS (on successful validation)
  - OTP_VALIDATION (on validation failure)
- **Condition/Notes**: Supports both payment and transaction_request flows

### 15. Terms and Conditions Screen
- **Screen ID**: TERMS_CONDITIONS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/terms-and-conditions.html`
- **Trigger**: Access to `/terms-and-conditions`
- **Navigates To**:
  - HOME (on "Accept" or "Skip" button)
- **Condition/Notes**: Updates user agreement and redirects to home

### 16. Privacy Policy Screen
- **Screen ID**: PRIVACY_POLICY
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/code/snippet/PrivacyPolicy.scala`
- **Trigger**: Access to `/privacy-policy`
- **Navigates To**:
  - HOME (on "Accept" or "Skip" button)
- **Condition/Notes**: Updates user agreement and redirects to home

### 17. Create Sandbox Account Screen
- **Screen ID**: CREATE_ACCOUNT
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/create-sandbox-account.html`
- **Trigger**: Access to create account functionality
- **Navigates To**:
  - CREATE_ACCOUNT_SUCCESS (on successful account creation)
  - CREATE_ACCOUNT (on validation errors)
- **Condition/Notes**: AJAX form for creating test bank accounts

### 18. User Authentication Context Update Screen
- **Screen ID**: AUTH_CONTEXT_UPDATE
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/add-user-auth-context-update-request.html`
- **Trigger**: Access to `/add-user-auth-context-update-request`
- **Navigates To**: CONFIRM_AUTH_CONTEXT_UPDATE (on form submission)
- **Condition/Notes**: First step of user authentication context update

### 19. Confirm Authentication Context Update Screen
- **Screen ID**: CONFIRM_AUTH_CONTEXT_UPDATE
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/confirm-user-auth-context-update-request.html`
- **Trigger**: Redirect from AUTH_CONTEXT_UPDATE with AUTH_CONTEXT_UPDATE_ID
- **Navigates To**:
  - HOME (on successful OTP validation)
  - CONFIRM_AUTH_CONTEXT_UPDATE (on validation failure)
- **Condition/Notes**: OTP validation for authentication context update

### 20. Dummy User Tokens Screen
- **Screen ID**: DUMMY_USER_TOKENS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/dummy-user-tokens.html`
- **Trigger**: Access to `/dummy-user-tokens` with consumer_key parameter (requires login)
- **Navigates To**: External API testing tools
- **Condition/Notes**: Displays Direct Login tokens for testing

### 21. User Information Screen
- **Screen ID**: USER_INFORMATION
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/user-information.html`
- **Trigger**: Access to `/user-information`
- **Navigates To**: Various navigation based on content
- **Condition/Notes**: Displays user information and related links

### 22. Introduction Screen
- **Screen ID**: INTRODUCTION
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/introduction.html`
- **Trigger**: Access to `/introduction`
- **Navigates To**: Various API documentation and external links
- **Condition/Notes**: API introduction and documentation links

### 23. SDKs Screen
- **Screen ID**: SDKS
- **Source**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/webapp/sdks.html`
- **Trigger**: Access to `/sdks`
- **Navigates To**: External SDK repositories and documentation
- **Condition/Notes**: SDK showcase and download links

## Navigation Patterns

### Authentication Flow
1. Unauthenticated user → LOGIN → HOME (on success)
2. Protected resource access → LOGIN → Original resource (on success)

### OAuth Flow
1. External app → OAUTH_AUTHORIZE → LOGIN (if needed) → OAUTH_AUTHORIZE → OAUTH_THANKS → External app

### User Registration Flow
1. Invitation email → USER_INVITATION → PASSWORD_RESET → LOGIN → HOME

### Consumer Registration Flow
1. HOME → CONSUMER_REGISTRATION → CONSUMER_REGISTRATION_SUCCESS → DUMMY_USER_TOKENS

### Consent Management Flow
1. External app → BG_CONSENT_REQUEST → BG_CONSENT_SCA (if needed) → External app
2. User → CONSENTS → CONSENTS (AJAX updates)

### Administrative Flows
1. User → TERMS_CONDITIONS → HOME
2. User → PRIVACY_POLICY → HOME
3. User → CREATE_ACCOUNT → CREATE_ACCOUNT_SUCCESS

## Key Navigation Components

### Sitemap Configuration
- **File**: `/home/ubuntu/repos/obp-api-ashish/00.phase-1-input/OBP-API-develop/obp-api/src/main/scala/bootstrap/liftweb/Boot.scala`
- **Lines**: 565-605
- **Function**: Defines all available routes and access controls

### Authentication Checks
- **AuthUser.loginFirst**: Requires user authentication
- **Admin.loginFirst**: Requires admin authentication
- **AuthUser.loggedIn_?**: Checks if user is logged in

### Redirect Mechanisms
- **S.redirectTo()**: Server-side redirects
- **OAuth callback URLs**: External application redirects
- **Form submissions**: POST-redirect-GET pattern

### Conditional Navigation
- **User role checks**: Admin vs regular user flows
- **Authentication state**: Logged in vs anonymous flows
- **Validation results**: Success vs error flows
- **Feature flags**: Hydra integration, Keycloak integration

## Error Handling

### Invalid States
- Expired invitations → USER_INVITATION_INVALID
- Already logged in during invitation → USER_INVITATION_WARNING
- Invalid OAuth tokens → Error messages on OAUTH_AUTHORIZE
- Validation failures → Return to form with error messages

### Security Checks
- CSRF protection on forms
- Token validation for OAuth flows
- Session management for authenticated users
- Account ownership validation for consent flows

---
*Analysis completed on 16-9-2025 for ashish-019-hash/obp-api repository*
