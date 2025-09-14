# Screen Flows Documentation

## Overview
This document maps the business-level screen/navigation flows for the Open Bank Project (OBP) API web application built with the Lift web framework (Scala). The application provides user authentication, OAuth authorization, consent management, and various banking API interfaces.

## Framework Notes
- **Framework**: Lift Web Framework (Scala) - not Spring as originally specified
- **Navigation Pattern**: Uses "snippets" instead of controllers for request handling
- **Template Binding**: Uses `data-lift` attributes for snippet-to-template binding
- **Routing**: Configured via sitemap in Boot.scala rather than annotations

## Screen Flow Mappings

### 1. Home Page / Index
- **Screen ID**: HOME_PAGE
- **Title**: Open Bank Project Home Page
- **Triggers**: 
  - GET / (root URL)
  - Navigation menu "Home" link
- **Navigates To**: Various sections based on user state
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/bootstrap/liftweb/Boot.scala`
  - Method: sitemap configuration
  - Lines: 565-566 (`Menu.i("Home") / "index"`)
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/index.html`
  - Template: Uses default-en.html layout
  - Lines: 1-72 (full template structure)
- **Conditions/Guards**: None (public access)
- **Post/Redirect Pattern**: Direct render
- **Notes**: Entry point for the application, shows different content based on authentication state

### 2. Login Page
- **Screen ID**: LOGIN_PAGE
- **Title**: User Authentication
- **Triggers**: 
  - Click "Log on" button in navigation
  - Redirect from protected pages
  - GET /user_mgt/login
- **Navigates To**: 
  - Dashboard/Home (success)
  - Same page with error (failure)
  - Registration page (new user)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/Login.scala`
  - Methods: `loggedIn`, `loggedOut`, `customiseLogin`
  - Lines: 45-77, 98-106
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/templates-hidden/default-en.html`
  - Binding: `data-lift="Login.loggedOut"`
  - Lines: 143-147 (login state section)
- **Conditions/Guards**: 
  - Show login form if `!AuthUser.loggedIn_?`
  - Show logout option if `AuthUser.loggedIn_?`
- **Post/Redirect Pattern**: POST-Redirect-GET after successful login
- **Notes**: Supports OpenID Connect integration, customizable login instructions

### 3. OAuth Authorization Page
- **Screen ID**: OAUTH_AUTHORIZE
- **Title**: OAuth Application Authorization
- **Triggers**: 
  - GET /oauth/authorize with token parameter
  - Third-party application authorization request
- **Navigates To**: 
  - Application callback URL (success)
  - Login form (not authenticated)
  - Error page (invalid token)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/OAuthAuthorisation.scala`
  - Method: `tokenCheck`
  - Lines: 77-191
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/oauth/authorize.html`
  - Binding: `data-lift="OAuthAuthorisation.tokenCheck"`
  - Lines: 1-72 (complete authorization form)
- **Conditions/Guards**: 
  - Valid OAuth token required
  - User must be authenticated
  - Token expiration check
- **Post/Redirect Pattern**: 
  - Redirect to callback URL with verifier
  - S.redirectTo() at line 122
- **Notes**: Handles OAuth 1.0a authorization flow, supports out-of-band (oob) callbacks

### 4. Consumer Registration
- **Screen ID**: CONSUMER_REGISTRATION
- **Title**: API Consumer Registration
- **Triggers**: 
  - Click "Consumer Registration" in navigation
  - GET /consumer-registration
- **Navigates To**: Registration success/failure pages
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/bootstrap/liftweb/Boot.scala`
  - Menu definition: `Menu("Consumer Registration", Helper.i18n("consumer.registration.nav.name"))`
  - Lines: 580
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/consumer-registration.html`
  - Navigation binding: `data-lift="Nav.item?name=Consumer%20Registration"`
  - Lines: 122-124 (navigation template)
- **Conditions/Guards**: `AuthUser.loginFirst` - requires authentication
- **Post/Redirect Pattern**: Direct render with form processing
- **Notes**: Allows developers to register for API access keys

### 5. Berlin Group Consent Screen
- **Screen ID**: BG_CONSENT_REQUEST
- **Title**: Berlin Group Consent Request
- **Triggers**: 
  - GET /confirm-bg-consent-request with consent parameters
  - Berlin Group PSD2 consent flow
- **Navigates To**: 
  - SCA confirmation page
  - Redirect URI (success/failure)
  - Error page (invalid consent)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/BerlinGroupConsent.scala`
  - Method: `confirmBgConsentRequest`
  - Lines: 448-451
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/confirm-bg-consent-request.html`
  - Binding: `data-lift="BerlinGroupConsent.confirmBgConsentRequest"`
- **Conditions/Guards**: 
  - `AuthUser.loginFirst` - requires authentication
  - Valid consent request ID
  - User must own requested accounts
- **Post/Redirect Pattern**: 
  - Redirect to SCA page or final redirect URI
  - Complex flow with multiple redirect steps
- **Notes**: Implements PSD2 consent management for account access

### 6. Berlin Group SCA (Strong Customer Authentication)
- **Screen ID**: BG_CONSENT_SCA
- **Title**: Strong Customer Authentication
- **Triggers**: 
  - POST from consent request page
  - GET /confirm-bg-consent-request-sca
- **Navigates To**: 
  - Final redirect URI (success)
  - Error page (invalid OTP)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/BerlinGroupConsent.scala`
  - Method: `confirmConsentRequestProcessSca`
  - Lines: 417-429
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/confirm-bg-consent-request-sca.html`
  - Binding: `data-lift="BerlinGroupConsent.confirmBgConsentRequest"`
  - Lines: 30-31
- **Conditions/Guards**: 
  - Valid OTP required
  - Active consent session
  - User authentication
- **Post/Redirect Pattern**: Final redirect to TPP application
- **Notes**: Handles two-factor authentication for PSD2 compliance

### 7. VRP Consent Request
- **Screen ID**: VRP_CONSENT_REQUEST
- **Title**: Variable Recurring Payment Consent
- **Triggers**: 
  - GET /confirm-vrp-consent-request with consent ID
  - VRP consent flow initiation
- **Navigates To**: 
  - VRP consent confirmation page
  - Error page (invalid consent)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/VrpConsentCreation.scala`
  - Method: `confirmVrpConsentRequest`
  - Lines: 157-162 (redirect logic)
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/confirm-vrp-consent-request.html`
  - Binding: `data-lift="VrpConsentCreation.confirmVrpConsentRequest"`
  - Lines: 30-31
- **Conditions/Guards**: 
  - Valid consent ID
  - User authentication required
- **Post/Redirect Pattern**: 
  - S.redirectTo() with consent parameters
  - Line 160-162
- **Notes**: Handles Variable Recurring Payment consent flows

### 8. User Registration/Sign Up
- **Screen ID**: USER_REGISTRATION
- **Title**: New User Registration
- **Triggers**: 
  - Click "Register" button
  - GET /user_mgt/sign_up
- **Navigates To**: 
  - Registration success page
  - Login page (existing user)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/Login.scala`
  - Method: `loggedOut` (signup link generation)
  - Lines: 72-75
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/templates-hidden/default-en.html`
  - Links: Multiple signup links
  - Lines: 146, 201, 238-240
- **Conditions/Guards**: Available to non-authenticated users
- **Post/Redirect Pattern**: Form submission with validation
- **Notes**: Multiple entry points for user registration

### 9. User Invitation Flow
- **Screen ID**: USER_INVITATION
- **Title**: User Invitation Processing
- **Triggers**: 
  - GET /user-invitation with invitation parameters
  - Email invitation links
- **Navigates To**: 
  - Password reset page (valid invitation)
  - Invalid invitation page
  - Warning page (already logged in)
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/UserInvitation.scala`
  - Redirects: Lines 139, 210, 213
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/user-invitation.html`
  - Related pages: user-invitation-invalid.html, user-invitation-warning.html
- **Conditions/Guards**: 
  - Valid invitation token
  - User not already logged in
- **Post/Redirect Pattern**: 
  - S.redirectTo() to appropriate destination
  - Multiple redirect paths based on validation
- **Notes**: Handles user onboarding via email invitations

### 10. Terms and Conditions
- **Screen ID**: TERMS_CONDITIONS
- **Title**: Terms and Conditions Agreement
- **Triggers**: 
  - Click "Terms and Conditions" link
  - Mandatory agreement flow
- **Navigates To**: 
  - Home page (after agreement)
  - Skip to home page
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/TermsAndConditions.scala`
  - Methods: `skipButtonDefense`, agreement processing
  - Lines: 47, 76
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/terms-and-conditions.html`
  - Binding: `data-lift="TermsAndConditions.updateForm"`
  - Lines: 29
- **Conditions/Guards**: May be required for authenticated users
- **Post/Redirect Pattern**: 
  - S.redirectTo("/") after processing
  - Lines 47, 76
- **Notes**: Handles user agreement to updated terms

### 11. Privacy Policy
- **Screen ID**: PRIVACY_POLICY
- **Title**: Privacy Policy Agreement
- **Triggers**: 
  - Click "Privacy Policy" link
  - Mandatory privacy agreement flow
- **Navigates To**: 
  - Home page (after agreement)
  - Skip to home page
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/code/snippet/PrivacyPolicy.scala`
  - Methods: `skipButtonDefense`, agreement processing
  - Lines: 47, 73
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/privacy-policy.html`
- **Conditions/Guards**: May be required for authenticated users
- **Post/Redirect Pattern**: 
  - S.redirectTo("/") after processing
  - Lines 47, 73
- **Notes**: Handles user agreement to privacy policy updates

### 12. User Information Page
- **Screen ID**: USER_INFORMATION
- **Title**: User Account Information
- **Triggers**: 
  - Click username in navigation (when logged in)
  - GET /user-information
- **Navigates To**: Various account management functions
- **Controller Source**: 
  - File: `/obp-api/src/main/scala/bootstrap/liftweb/Boot.scala`
  - Menu: `Menu("User Information", "User Information") / "user-information"`
  - Lines: 585
- **Template Source**: 
  - File: `/obp-api/src/main/webapp/user-information.html`
  - Navigation link: Line 152 in default-en.html
- **Conditions/Guards**: No explicit authentication required in menu definition
- **Post/Redirect Pattern**: Direct render
- **Notes**: Displays user account details and management options

## Navigation Patterns Summary

### Authentication-Based Navigation
- **Logged Out State**: Shows login form, registration links, forgot password
- **Logged In State**: Shows username, logout button, user-specific content
- **Admin State**: Additional admin logout functionality

### OAuth Flow Navigation
1. Application requests authorization → OAuth authorize page
2. User authenticates → Token validation
3. User grants permission → Callback redirect with verifier
4. Application receives access → OAuth thanks page

### Consent Flow Navigation (Berlin Group PSD2)
1. TPP initiates consent → Consent request page
2. User reviews permissions → Account selection
3. Strong Customer Authentication → OTP verification
4. Final consent → Redirect to TPP application

### Error Handling Navigation
- Invalid tokens → Error messages with retry options
- Authentication failures → Login form with error display
- Invalid invitations → Dedicated error pages
- Expired sessions → Redirect to login with return URL

## Technical Implementation Notes

### Lift Framework Specifics
- **Snippets**: Replace traditional controllers, return CSS selectors for DOM manipulation
- **SiteMap**: Centralized routing configuration in Boot.scala
- **CSS Selectors**: Used for template binding and DOM updates
- **Session Variables**: Maintain state across requests (e.g., consent flow data)

### Security Patterns
- **AuthUser.loginFirst**: Enforces authentication for protected pages
- **Token Validation**: OAuth tokens checked for validity and expiration
- **CSRF Protection**: Built into Lift's form handling
- **Session Management**: Automatic session handling with configurable timeouts

### Redirect Patterns
- **S.redirectTo()**: Primary redirect mechanism used throughout
- **Post-Redirect-Get**: Implemented for form submissions
- **Callback URLs**: OAuth and consent flows use external redirects
- **Internal Navigation**: Menu-based navigation with current page highlighting
