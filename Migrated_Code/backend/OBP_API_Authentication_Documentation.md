# OBP-API Authentication Documentation - Complete Reference

## Table of Contents
1. [Introduction](#introduction)
2. [Authentication Methods](#authentication-methods)
   - [OAuth 1.0a](#oauth-10a)
   - [OAuth 2.0 / OpenID Connect (OIDC)](#oauth-20--openid-connect-oidc)
   - [Direct Login](#direct-login)
   - [Gateway Login](#gateway-login)
   - [DAuth](#dauth)
3. [User Management System](#user-management-system)
   - [ResourceUser vs AuthUser Architecture](#resourceuser-vs-authuser-architecture)
   - [User Lifecycle Management](#user-lifecycle-management)
   - [User Agreements](#user-agreements)
   - [GDPR Compliance](#gdpr-compliance)
4. [Admin Authentication System](#admin-authentication-system)
   - [Admin User Management](#admin-user-management)
   - [Admin Login Flow](#admin-login-flow)
   - [Admin Session Management](#admin-session-management)
5. [Entitlement System](#entitlement-system)
   - [Role-Based Access Control](#role-based-access-control)
   - [Entitlement Management](#entitlement-management)
   - [Entitlement Notifications](#entitlement-notifications)
6. [Scope System](#scope-system)
   - [Consumer-Based Role Scoping](#consumer-based-role-scoping)
   - [Scope Management](#scope-management)
7. [View Permission System](#view-permission-system)
   - [System vs Custom Views](#system-vs-custom-views)
   - [View-Based Access Control](#view-based-access-control)
   - [Permission Management](#permission-management)
8. [Authentication Context System](#authentication-context-system)
   - [UserAuthContext](#userauthcontext)
   - [ConsentAuthContext](#consentauthcontext)
9. [Consent Management](#consent-management)
   - [OBP Consent](#obp-consent)
   - [Berlin Group Consent](#berlin-group-consent)
   - [UK Open Banking Consent](#uk-open-banking-consent)
   - [Consent JWT Structure](#consent-jwt-structure)
10. [Login Attempt Protection](#login-attempt-protection)
    - [Bad Login Attempt Tracking](#bad-login-attempt-tracking)
    - [User Lockout Mechanisms](#user-lockout-mechanisms)
11. [User Lock System](#user-lock-system)
    - [Lock Types](#lock-types)
    - [Lock Management](#lock-management)
12. [Authentication Type Validation](#authentication-type-validation)
    - [Operation-Specific Authentication](#operation-specific-authentication)
    - [Authentication Type Restrictions](#authentication-type-restrictions)
13. [Consumer Management](#consumer-management)
    - [OAuth Consumer Registration](#oauth-consumer-registration)
    - [Consumer Verification](#consumer-verification)
    - [Certificate-Based Authentication](#certificate-based-authentication)
14. [Authentication Architecture](#authentication-architecture)
15. [Configuration](#configuration)
16. [Security Features](#security-features)
17. [Implementation Details](#implementation-details)
18. [Error Handling](#error-handling)
19. [Testing](#testing)
20. [Troubleshooting](#troubleshooting)

## Introduction

The Open Bank Project (OBP) API implements a comprehensive multi-method authentication system designed to support various client applications and integration scenarios. This document provides a detailed overview of all authentication methods, architecture, configuration, and implementation details without missing any components.

The authentication system is designed to meet the following requirements:
- Support for multiple authentication standards (OAuth 1.0a, OAuth 2.0, OpenID Connect)
- Simplified authentication for direct integrations (Direct Login, Gateway Login)
- Compliance with financial industry security standards and regulations (PSD2, Open Banking)
- Flexible configuration for different deployment scenarios
- Comprehensive security features including token validation, certificate-based authentication, and rate limiting
- Role-based access control through entitlements and view permissions
- Consent management for regulatory compliance
- User lifecycle management with GDPR compliance

## Authentication Methods

### OAuth 1.0a

**Primary Implementation File:** [oauth1.0.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/oauth1.0.scala)

OAuth 1.0a is a three-legged authentication protocol that provides a secure way for applications to access user data without exposing user credentials.

**Key Features:**
- Full OAuth 1.0a three-legged flow implementation
- Request token → Authorization → Access token flow
- HMAC-SHA1 and HMAC-SHA256 signature methods
- Nonce validation and timestamp verification
- Token management with expiration

**Authentication Flow:**
1. **Request Token**: Client application requests a temporary token
   ```
   POST /oauth/initiate
   ```
2. **User Authorization**: User is redirected to authorization page
   ```
   GET /oauth/authorize?oauth_token={request_token}
   ```
3. **Access Token**: Client exchanges authorized request token for access token
   ```
   POST /oauth/token
   ```
4. **API Access**: Client uses access token to make authenticated API calls

**Code Highlights:**
- `OAuthHandshake` object handles the complete OAuth 1.0a flow
- `/oauth/initiate` and `/oauth/token` endpoints
- Signature verification and parameter validation

### OAuth 2.0 / OpenID Connect (OIDC)

**Primary Implementation Files:**
- [OAuth2.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/OAuth2.scala)
- [openidconnect.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/openidconnect.scala)

OAuth 2.0 and OpenID Connect provide modern authentication flows with support for various identity providers and token types.

**Key Features:**
- Support for multiple identity providers: Google, Yahoo, Azure, Keycloak, Hydra ORY
- JWT token validation with JWKS (JSON Web Key Set) support
- ID token and access token validation
- Certificate-based client authentication (PSD2 compliance)
- Configurable via properties like `allow_oauth2_login`, `oauth2.jwk_set.url`

**Authentication Flow:**
1. **Authorization Request**: Client redirects user to identity provider
2. **User Authentication**: User authenticates with identity provider
3. **Authorization Code**: Identity provider redirects back with code
4. **Token Exchange**: Client exchanges code for access and ID tokens
5. **API Access**: Client uses tokens to make authenticated API calls

**Code Highlights:**
- `OAuth2Login.getUser()` and `OAuth2Login.getUserFuture()` - main authentication entry points
- Support for both self-contained JWT tokens and introspectable access tokens (Hydra)
- Automatic consumer creation for OIDC flows

### Direct Login

**Primary Implementation File:** [directlogin.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/directlogin.scala)

Direct Login provides a simplified authentication method for scenarios where OAuth flows are not practical.

**Key Features:**
- Simplified JWT-based authentication
- Username/password authentication with consumer key
- Token generation using HMAC protection
- Configurable token expiration (default 4 weeks)
- Support for both header styles: `DirectLogin:` and `Authorization: DirectLogin`

**Authentication Flow:**
1. **Authentication Request**: Client sends username, password, and consumer key
   ```
   POST /my/logins/direct
   Authorization: DirectLogin username="username", password="password", consumer_key="consumer-key"
   ```
2. **Token Response**: Server returns a JWT token
3. **API Access**: Client uses token for subsequent API calls
   ```
   GET /obp/v4.0.0/banks
   Authorization: DirectLogin token="eyJhbGciOiJIUzI1NiJ9..."
   ```

**Code Highlights:**
- `DirectLogin.createToken()` - generates JWT tokens
- `/my/logins/direct` endpoint for token creation
- Built-in parameter validation and consumer verification

### Gateway Login

**Primary Implementation File:** [GatewayLogin.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/GatewayLogin.scala)

Gateway Login enables integration with Core Banking Systems (CBS) using JWT tokens.

**Key Features:**
- JWT-based authentication for gateway integration
- Core Banking System (CBS) token integration
- Session management with `is_first` flag logic
- Account refresh and user creation workflows

**Authentication Flow:**
1. **JWT Creation**: CBS creates a signed JWT with user and account information
2. **Authentication**: Client sends JWT in authorization header
   ```
   GET /obp/v4.0.0/banks
   Authorization: GatewayLogin token="eyJhbGciOiJIUzI1NiJ9..."
   ```
3. **Validation**: OBP validates JWT signature and claims
4. **User Creation/Update**: OBP creates or updates user and account information
5. **API Access**: Client continues to use JWT for API calls

**Code Highlights:**
- `GatewayLogin.createJwt()` - creates gateway JWT tokens
- Integration with bank connectors for account data
- Automatic user and consumer creation

### DAuth

**Primary Implementation File:** [dauth.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/api/dauth.scala)

DAuth (Dynamic Authentication) provides a flexible authentication mechanism for dynamic client registration and token issuance.

**Key Features:**
- Dynamic client registration
- JWT-based token issuance
- Support for different authentication flows
- Integration with external identity providers

**Authentication Flow:**
1. **Client Registration**: Client registers dynamically
2. **Authentication Request**: Client initiates authentication flow
3. **Token Issuance**: Server issues JWT token
4. **API Access**: Client uses token for API calls

**Code Highlights:**
- `DAuth.createToken()` - generates JWT tokens
- Dynamic client registration endpoints
- Token validation and verification

## User Management System

**Primary Implementation Files:**
- [ResourceUser.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/model/dataAccess/ResourceUser.scala)
- [AuthUser.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/model/dataAccess/AuthUser.scala)
- [Users.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/users/Users.scala)
- [LiftUsers.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/users/LiftUsers.scala)

The OBP-API implements a sophisticated user management system that handles user creation, authentication, and lifecycle management across multiple authentication methods.

### ResourceUser vs AuthUser Architecture

The OBP-API employs a two-tier user architecture that separates web authentication from API resource access:

**AuthUser:**
- Extends Lift's `MegaProtoUser` for web authentication
- Handles username/password validation, email verification, password reset
- Used primarily for webpage login functionality
- Provides built-in validation methods for credentials

**ResourceUser:**
- Simple `LongKeyedMapper` for API resource access
- All accounts, transactions, roles, views, and other resources link to ResourceUser
- Consumer keys and tokens belong to ResourceUser
- Stores provider information for external authentication

**Relationship:**
- One-to-one relationship between AuthUser and ResourceUser
- AuthUser's `user` field is a foreign key to ResourceUser
- They share the same username and email
- When a user signs up, AuthUser is created first, then ResourceUser

**Code Example:**
```scala
// From ResourceUser.scala
class ResourceUser extends LongKeyedMapper[ResourceUser] with User with ManyToMany with OneToMany[Long, ResourceUser] {
  def getSingleton = ResourceUser
  def primaryKeyField = id

  object id extends MappedLongIndex(this)
  object userId_ extends MappedUUID(this)
  object email extends MappedEmail(this, 100)
  object name_ extends MappedString(this, 100)
  object provider_ extends MappedString(this, 100)
  object providerId extends MappedString(this, 100)
  // Additional fields...
}
```

### User Lifecycle Management

The user lifecycle is managed through a comprehensive set of operations:

**User Creation:**
```scala
// From LiftUsers.scala
def createResourceUser(provider: String,
                      providerId: Option[String],
                      createdByConsentId: Option[String],
                      name: Option[String],
                      email: Option[String],
                      userId: Option[String],
                      createdByUserInvitationId: Option[String],
                      company: Option[String],
                      lastMarketingAgreementSignedDate: Option[Date]): Box[ResourceUser]
```

**User Lookup Methods:**
- `getUserByResourceUserId` - Find user by internal ID
- `getUserByProviderId` - Find user by external provider ID
- `getUserByUserId` - Find user by UUID
- `getUserByProviderAndUsername` - Find user by provider and username
- `getUserByEmail` - Find users by email address

**User Creation Scenarios:**
- Standard signup through web interface
- OAuth 2.0/OIDC authentication with external provider
- Gateway Login with Core Banking System
- User invitation process
- Consent-based creation

**Provider-Based Authentication:**
```scala
// From LiftUsers.scala
def getOrCreateUserByProviderId(provider: String, idGivenByProvider: String, consentId: Option[String], name: Option[String], email: Option[String]): (Box[User], Boolean)
```

### User Agreements

The system supports tracking user agreements for regulatory compliance:

**Agreement Types:**
- Marketing information consent
- Terms and conditions acceptance
- Privacy policy acceptance

**Agreement Management:**
```scala
// From LiftUsers.scala
private def getUserAgreements(user: ResourceUser) = {
  val acceptMarketingInfo = UserAgreementProvider.userAgreementProvider.vend.getLastUserAgreement(user.userId, "accept_marketing_info")
  val termsAndConditions = UserAgreementProvider.userAgreementProvider.vend.getLastUserAgreement(user.userId, "terms_and_conditions")
  val privacyConditions = UserAgreementProvider.userAgreementProvider.vend.getLastUserAgreement(user.userId, "privacy_conditions")
  val agreements = acceptMarketingInfo.toList ::: termsAndConditions.toList ::: privacyConditions.toList
  agreements
}
```

**Agreement Tracking:**
- Timestamp recording for each agreement type
- Latest agreement version tracking
- Agreement history for audit purposes

### GDPR Compliance

The system implements GDPR compliance features for user data protection:

**Data Scrambling:**
```scala
// From LiftUsers.scala
def scrambleDataOfResourceUser(userPrimaryKey: UserPrimaryKey): Box[Boolean] = {
  for {
    u <- ResourceUser.find(By(ResourceUser.id, userPrimaryKey.value))
  } yield {
    AuthUser.find(By(AuthUser.user, userPrimaryKey.value)) match {
      case Empty =>
        u.Company(Helpers.randomString(16))
         .IsDeleted(true)
         .name_("DELETED-" + Helpers.randomString(16))
         .email(Helpers.randomString(10) + "@example.com")
         .providerId(Helpers.randomString(16))
         .save
      case _ =>
        u.Company(Helpers.randomString(16))
         .IsDeleted(true)
         .save
    }
  }
}
```

**Deletion Handling:**
- Soft deletion with `IsDeleted` flag
- Data anonymization for deleted users
- Preservation of relationships while removing personal data

**Data Access:**
- Methods to retrieve all user data for GDPR data access requests
- Filtering options for deleted/active users
- Comprehensive user data export capabilities

## Admin Authentication System

**Primary Implementation File:** [Admin.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/model/dataAccess/Admin.scala)

The OBP-API includes a separate authentication system for administrative users with elevated privileges.

### Admin User Management

The Admin system extends Lift's MegaProtoUser for administrative authentication:

```scala
// From Admin.scala
class Admin extends MegaProtoUser[Admin] {
  def getSingleton = Admin
}

object Admin extends Admin with MetaMegaProtoUser[Admin] {
  override def dbTableName = "admins"
  override def basePath = "admin_mgt" :: Nil
  override def menuNameSuffix = "Admin"
  // Additional configuration...
}
```

**Key Features:**
- Separate database table (`admins`) for admin users
- Custom URL paths to avoid conflicts with regular users
- Specialized menu names and UI elements
- Email validation can be optionally skipped

### Admin Login Flow

The admin login process uses a customized flow:

```scala
// From Admin.scala
override def loginXhtml = {
  (<form method="post" action={ObpS.uri}><table><tr><td
            colspan="2">Admin Log In</td></tr>
        <tr><td>{userNameFieldString}</td><td><user:email /></td></tr>
        <tr><td>{S.?("password")}</td><td><user:password /></td></tr>
        <tr><td><a href={lostPasswordPath.mkString("/", "/", "")}
              >{S.?("recover.password")}</a></td><td><user:submit /></td></tr></table>
   </form>)
}
```

**Authentication Process:**
1. Admin navigates to admin login page
2. Credentials are submitted and validated
3. On success, admin is redirected to the admin dashboard
4. Failed attempts may be tracked for security

### Admin Session Management

Admin sessions are managed with special considerations:

```scala
// From Admin.scala
object loginReferer extends SessionVar("/")

override def homePage = {
  val ret = loginReferer.is
  loginReferer.remove()
  ret
}

override def login = {
  for(
    r <- S.referer
    if loginReferer.is.equals("/")
  ) loginReferer.set(r)
  super.login
}
```

**Session Features:**
- Referer tracking for post-login redirection
- Session variables for admin state
- Custom home page logic after authentication
- Restricted signup and user management

**Security Restrictions:**
- Disabled self-signup (`createUserMenuLoc = Empty`)
- Restricted user editing (`editUserMenuLoc = Empty`)
- Password reset requires manual intervention (`resetPasswordMenuLoc = Empty`)

## Entitlement System

**Primary Implementation Files:**
- [Entilement.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/entitlement/Entilement.scala)
- [MappedEntitlements.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/entitlement/MappedEntitlements.scala)

The OBP-API implements a comprehensive role-based access control system through entitlements.

### Role-Based Access Control

Entitlements provide fine-grained control over API access:

```scala
// From Entilement.scala
trait Entitlement {
  def entitlementId: String
  def bankId: String
  def userId: String
  def roleName: String
  def createdByProcess: String
}
```

**Entitlement Structure:**
- `entitlementId`: Unique identifier for the entitlement
- `bankId`: Bank the entitlement applies to (or global if applicable)
- `userId`: User granted the entitlement
- `roleName`: Role/permission name (e.g., `CanCreateAccount`, `CanGetCustomer`)
- `createdByProcess`: How the entitlement was created (manual, automated)

**Role Types:**
- Bank-specific roles (require bankId)
- Global roles (apply across all banks)
- Dynamic entity roles (for custom entities)

### Entitlement Management

The system provides comprehensive entitlement management functions:

```scala
// From MappedEntitlementsProvider
trait EntitlementProvider {
  def getEntitlement(bankId: String, userId: String, roleName: String): Box[Entitlement]
  def getEntitlementById(entitlementId: String): Box[Entitlement]
  def getEntitlementsByUserId(userId: String): Box[List[Entitlement]]
  def getEntitlementsByBankId(bankId: String): Future[Box[List[Entitlement]]]
  def deleteEntitlement(entitlement: Box[Entitlement]): Box[Boolean]
  def getEntitlements(): Box[List[Entitlement]]
  def getEntitlementsByRole(roleName: String): Box[List[Entitlement]]
  def addEntitlement(bankId: String, userId: String, roleName: String, createdByProcess: String="manual", grantorUserId: Option[String]=None): Box[Entitlement]
  def deleteDynamicEntityEntitlement(entityName: String, bankId:Option[String]): Box[Boolean]
  def deleteEntitlements(entityNames: List[String]): Box[Boolean]
}
```

**Key Operations:**
- Adding entitlements to users
- Retrieving entitlements by various criteria
- Deleting entitlements
- Managing dynamic entity entitlements

**Permission Checking:**
```scala
// Example of permission checking for entitlement granting
val canCreateEntitlementAtAnyBank = MappedEntitlement.findAll(By(MappedEntitlement.mUserId, userId)).exists(e => e.roleName == CanCreateEntitlementAtAnyBank)
val canCreateEntitlementAtOneBank = MappedEntitlement.findAll(By(MappedEntitlement.mUserId, userId)).exists(e => e.roleName == CanCreateEntitlementAtOneBank && e.bankId == bankId)
if(canCreateEntitlementAtAnyBank || canCreateEntitlementAtOneBank) {
  addEntitlementToUser()
} else {
  Failure(ErrorMessages.EntitlementCannotBeGrantedGrantorIssue)
}
```

### Entitlement Notifications

The system includes notification capabilities for entitlement changes:

```scala
// From MappedEntitlementsProvider
def addEntitlement(bankId: String, userId: String, roleName: String, createdByProcess: String ="manual", grantorUserId: Option[String]=None): Box[Entitlement] = {
  def addEntitlementToUser(): Full[MappedEntitlement] = {
    val addEntitlement: MappedEntitlement =
      MappedEntitlement.create.mBankId(bankId).mUserId(userId).mRoleName(roleName).mCreatedByProcess(createdByProcess)
      .saveMe()
    // When a role is Granted, we should send an email to the Recipient telling them they have been granted the role.
    NotificationUtil.sendEmailRegardingAssignedRole(userId: String, addEntitlement: Entitlement)
    Full(addEntitlement)
  }
  // Permission checking and processing...
}
```

**Notification Features:**
- Email notifications for role assignments
- Notification customization based on role type
- Audit trail for entitlement changes

## Scope System

**Primary Implementation File:** [Scope.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/scope/Scope.scala)

The OBP-API implements a scope system that provides consumer-based role scoping for fine-grained access control.

### Consumer-Based Role Scoping

Scopes link consumers (API clients) to specific roles:

```scala
// From Scope.scala
trait Scope {
  def scopeId: String
  def bankId: String
  def consumerId: String
  def roleName: String
}
```

**Scope Structure:**
- `scopeId`: Unique identifier for the scope
- `bankId`: Bank the scope applies to
- `consumerId`: Consumer (API client) the scope applies to
- `roleName`: Role/permission name

**Purpose:**
- Restrict consumer access to specific roles
- Enable fine-grained API client permissions
- Support OAuth 2.0 scope-based authorization

### Scope Management

The system provides comprehensive scope management functions:

```scala
// From Scope.scala
trait ScopeProvider {
  def getScope(bankId: String, consumerId: String, roleName: String): Box[Scope]
  def getScopeById(ScopeId: String): Box[Scope]
  def getScopesByConsumerId(consumerId: String): Box[List[Scope]]
  def getScopesByConsumerIdFuture(consumerId: String): Future[Box[List[Scope]]]
  def deleteScope(Scope: Box[Scope]): Box[Boolean]
  def getScopes(): Box[List[Scope]]
  def getScopesFuture(): Future[Box[List[Scope]]]
  def addScope(bankId: String, consumerId: String, roleName: String): Box[Scope]
}
```

**Key Operations:**
- Adding scopes to consumers
- Retrieving scopes by various criteria
- Deleting scopes
- Managing scope relationships

**Integration with OAuth:**
- OAuth 2.0 scopes map to OBP roles
- Scope validation during token issuance
- Scope enforcement during API access

## View Permission System

**Primary Implementation Files:**
- [Views.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/views/Views.scala)
- [ViewPermission.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/views/system/ViewPermission.scala)

The OBP-API implements a sophisticated view-based permission system for controlling access to banking data.

### System vs Custom Views

The view system distinguishes between system-wide and account-specific views:

```scala
// From Views.scala
trait Views {
  def customView(viewId: ViewId, bankAccountId: BankIdAccountId): Box[View]
  def systemView(viewId: ViewId): Box[View]
  def getSystemViews(): Future[List[View]]
  // Additional methods...
}
```

**System Views:**
- Global views applicable across the platform
- Not tied to specific accounts
- Examples: Owner, Accountant, Auditor
- Managed through system configuration

**Custom Views:**
- Account-specific views
- Created for specific bank accounts
- Can have custom permissions and metadata
- User-manageable through API

### View-Based Access Control

Views control access to account data and operations:

```scala
// From Views.scala
trait Views {
  def permissions(account: BankIdAccountId): List[Permission]
  def permission(account: BankIdAccountId, user: User): Box[Permission]
  def grantAccessToCustomView(bankIdAccountIdViewId: BankIdAccountIdViewId, user: User): Box[View]
  def grantAccessToSystemView(bankId: BankId, accountId: AccountId, view: View, user: User): Box[View]
  def revokeAccess(bankIdAccountIdViewId: BankIdAccountIdViewId, user: User): Box[Boolean]
  // Additional methods...
}
```

**Access Control Operations:**
- Granting view access to users
- Revoking view access
- Checking permissions
- Managing multiple views simultaneously

**View Types:**
- Public views (accessible to all)
- Private views (restricted access)
- Firehose views (full data access)
- Custom views with specific permissions

### Permission Management

The system provides detailed permission management for views:

```scala
// From ViewPermission.scala
class ViewPermission extends LongKeyedMapper[ViewPermission] with IdPK with CreatedUpdated {
  def getSingleton = ViewPermission
  object bank_id extends MappedString(this, 255)
  object account_id extends MappedString(this, 255)
  object view_id extends UUIDString(this)
  object permission extends MappedString(this, 255)
  object extraData extends MappedString(this, 1024)
}
```

**Permission Operations:**
```scala
// From ViewPermission.scala
def resetViewPermissions(
  view: View,
  permissionNames: List[String],
  canGrantAccessToViews: List[String] = Nil,
  canRevokeAccessToViews: List[String] = Nil
): Unit
```

**Special Permissions:**
- `CAN_GRANT_ACCESS_TO_VIEWS` - Ability to grant view access to others
- `CAN_REVOKE_ACCESS_TO_VIEWS` - Ability to revoke view access
- Custom permissions with extra data for fine-grained control

## Authentication Context System

**Primary Implementation Files:**
- [UserAuthContextProvider.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/context/UserAuthContextProvider.scala)
- [ConsentAuthContextProvider.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/context/ConsentAuthContextProvider.scala)

The OBP-API implements an authentication context system for storing metadata related to authentication sessions.

### UserAuthContext

UserAuthContext stores authentication-related metadata for users:

```scala
// From UserAuthContextProvider.scala
trait UserAuthContextProvider {
  def createUserAuthContext(userId: String, key: String, value: String, consumerId: String): Future[Box[UserAuthContext]]
  def getUserAuthContexts(userId: String): Future[Box[List[UserAuthContext]]]
  def getUserAuthContextsBox(userId: String): Box[List[UserAuthContext]]
  def createOrUpdateUserAuthContexts(userId: String, userAuthContexts: List[BasicUserAuthContext]): Box[List[UserAuthContext]]
  def deleteUserAuthContexts(userId: String): Future[Box[Boolean]]
  def deleteUserAuthContextById(userAuthContextId: String): Future[Box[Boolean]]
}
```

**Key Features:**
- Key-value storage for authentication metadata
- User-specific context information
- Consumer association for tracking client applications
- Support for multiple context entries per user

**Common Use Cases:**
- Storing authentication method used
- Recording authentication timestamp
- Tracking device or location information
- Storing multi-factor authentication status

### ConsentAuthContext

ConsentAuthContext stores authentication-related metadata for consents:

```scala
// From ConsentAuthContextProvider.scala
trait ConsentAuthContextProvider {
  def createConsentAuthContext(consentId: String, key: String, value: String): Future[Box[ConsentAuthContext]]
  def getConsentAuthContexts(consentId: String): Future[Box[List[ConsentAuthContext]]]
  def getConsentAuthContextsBox(consentId: String): Box[List[ConsentAuthContext]]
  def createOrUpdateConsentAuthContexts(consentId: String, userAuthContexts: List[BasicUserAuthContext]): Box[List[ConsentAuthContext]]
  def deleteConsentAuthContexts(consentId: String): Future[Box[Boolean]]
  def deleteConsentAuthContextById(consentAuthContextId: String): Future[Box[Boolean]]
}
```

**Key Features:**
- Key-value storage for consent-related metadata
- Consent-specific context information
- Support for multiple context entries per consent
- Management functions for context lifecycle

**Common Use Cases:**
- Storing consent authentication method
- Recording consent approval details
- Tracking consent usage information
- Storing regulatory compliance data

## Consent Management

**Primary Implementation File:** [ConsentProvider.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/consent/ConsentProvider.scala)

The OBP-API implements a comprehensive consent management system supporting multiple regulatory standards.

### OBP Consent

The base consent system provides core functionality:

```scala
// From ConsentProvider.scala
trait ConsentProvider {
  def getConsents(queryParams: List[OBPQueryParam]): (List[MappedConsent], Long)
  def getConsentByConsentId(consentId: String): Box[MappedConsent]
  def getConsentByConsentRequestId(consentRequestId: String): Box[MappedConsent]
  def updateConsentStatus(consentId: String, status: ConsentStatus): Box[MappedConsent]
  def updateConsentUser(consentId: String, user: User): Box[MappedConsent]
  def getConsentsByUser(userId: String): List[MappedConsent]
  def createObpConsent(user: User, challengeAnswer: String, consentRequestId:Option[String], consumer: Option[Consumer] = None): Box[MappedConsent]
  def setJsonWebToken(consentId: String, jwt: String): Box[MappedConsent]
  def setValidUntil(consentId: String, validUntil: Date): Box[MappedConsent]
  def revoke(consentId: String): Box[MappedConsent]
  def checkAnswer(consentId: String, challenge: String): Box[MappedConsent]
  // Additional methods...
}
```

**Key Features:**
- Consent creation and management
- Status tracking and updates
- User association
- JWT token generation
- Challenge-response verification

### Berlin Group Consent

Extended consent functionality for Berlin Group standard:

```scala
// From ConsentProvider.scala
trait ConsentProvider {
  def createBerlinGroupConsent(
    user: Option[User],
    consumer: Option[Consumer],
    recurringIndicator: Boolean,
    validUntil: Date,
    frequencyPerDay: Int,
    combinedServiceIndicator: Boolean,
    apiStandard: Option[String],
    apiVersion: Option[String]): Box[ConsentTrait]

  def updateBerlinGroupConsent(
    consentId: String,
    usesSoFarTodayCounter: Int): Box[ConsentTrait]

  def revokeBerlinGroupConsent(consentId: String): Box[MappedConsent]
  // Additional methods...
}
```

**Berlin Group Specific Features:**
- `recurringIndicator` - Whether consent is for recurring access
- `frequencyPerDay` - Maximum access frequency per day
- `usesSoFarTodayCounter` - Current usage count
- `combinedServiceIndicator` - Whether payment initiation is included
- Special revocation handling

### UK Open Banking Consent

Extended consent functionality for UK Open Banking standard:

```scala
// From ConsentProvider.scala
trait ConsentProvider {
  def saveUKConsent(
    user: Option[User],
    bankId: Option[String],
    accountIds: Option[List[String]],
    consumerId: Option[String],
    permissions: List[String],
    expirationDateTime: Date,
    transactionFromDateTime: Date,
    transactionToDateTime: Date,
    apiStandard: Option[String],
    apiVersion: Option[String]
  ): Box[ConsentTrait]
  // Additional methods...
}
```

**UK Open Banking Specific Features:**
- Explicit permissions list
- Transaction date range specification
- Account IDs specification
- UK-specific status values (AUTHORISED, AWAITINGAUTHORISATION)

### Consent JWT Structure

The consent system uses JWT tokens with a specific structure:

```scala
// From ConsentProvider.scala
/**
 * this is the structure of the jwt token, try to see the case class directly, to see the all the fields.
 * case class ConsentJWT(
 *   createdByUserId: String,
 *   sub: String,
 *   iss: String,
 *   aud: String,
 *   jti: String,
 *   iat: Long,
 *   nbf: Long,
 *   exp: Long,
 *   name: Option[String],
 *   email: Option[String],
 *   entitlements: List[Role],
 *   views: List[ConsentView]
 * )
 */
def jsonWebToken: String
```

**JWT Components:**
- Standard JWT claims (sub, iss, aud, jti, iat, nbf, exp)
- User information (name, email)
- Entitlements granted by the consent
- Views accessible through the consent
- Creation and expiration metadata

**Consent Status Tracking:**
```scala
// From ConsentProvider.scala
object ConsentStatus extends Enumeration {
  type ConsentStatus = Value
  val INITIATED, ACCEPTED, REJECTED, rejected, REVOKED, EXPIRED,
      // The following one only exist in case of BerlinGroup
      received, valid, revokedByPsu, expired, terminatedByTpp,
     //these added for UK Open Banking
     AUTHORISED, AWAITINGAUTHORISATION = Value
}
```

## Login Attempt Protection

**Primary Implementation File:** [LoginAttempts.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/loginattempts/LoginAttempts.scala)

The OBP-API implements a comprehensive login attempt tracking and user lockout system.

### Bad Login Attempt Tracking

The system tracks failed login attempts:

```scala
// From LoginAttempts.scala
object LoginAttempt extends MdcLoggable {
  def maxBadLoginAttempts = APIUtil.getPropsValue("max.bad.login.attempts") openOr "5"

  def incrementBadLoginAttempts(provider: String, username: String): Unit = {
    username.isEmpty() match {
      case true => // Not a valid case. GitLab issue 389
        logger.warn(s"Username is empty: incrementBadLoginAttempts(username=$username, provider=$provider")
      case false =>
        // Find badLoginAttempt record if one exists for a user
        MappedBadLoginAttempt.find(
          By(MappedBadLoginAttempt.Provider, provider),
          By(MappedBadLoginAttempt.mUsername, username)
        ) match {
          // If it exits update the date and increment
          case Full(loginAttempt) =>
            loginAttempt
              .mLastFailureDate(now)
              .mBadAttemptsSinceLastSuccessOrReset(loginAttempt.mBadAttemptsSinceLastSuccessOrReset + 1) // Increment
              .save
          case _ =>
            // If none exists, add one
            MappedBadLoginAttempt.create
              .mUsername(username)
              .Provider(provider)
              .mLastFailureDate(now)
              .mBadAttemptsSinceLastSuccessOrReset(1) // Start with 1
              .save
        }
    }
  }
  // Additional methods...
}
```

**Key Features:**
- Provider-specific tracking (different auth methods)
- Configurable maximum attempts threshold
- Timestamp recording for failures
- Counter for consecutive failures

### User Lockout Mechanisms

The system implements user lockout based on failed attempts:

```scala
// From LoginAttempts.scala
def userIsLocked(provider: String, username: String): Boolean = {
  val result: Boolean = MappedBadLoginAttempt.find(
    By(MappedBadLoginAttempt.Provider, provider),
    By(MappedBadLoginAttempt.mUsername, username)
  ) match {
    case Full(loginAttempt) => loginAttempt.badAttemptsSinceLastSuccessOrReset > maxBadLoginAttempts.toInt match {
      case true => true
      case false => UserLocksProvider.isLocked(provider, username) // Check the table UserLocks
    }
    case _ => UserLocksProvider.isLocked(provider, username) // Check the table UserLocks
  }

  logger.debug(s"userIsLocked result for $username is $result")
  result
}
```

**Lockout Features:**
- Automatic lockout after exceeding threshold
- Integration with UserLocks system for manual locks
- Provider-specific lockout handling
- Logging of lockout events

**Reset Functionality:**
```scala
// From LoginAttempts.scala
def resetBadLoginAttempts(provider: String, username: String): Unit = {
  MappedBadLoginAttempt.find(
    By(MappedBadLoginAttempt.Provider, provider),
    By(MappedBadLoginAttempt.mUsername, username)
  ) match {
    case Full(loginAttempt) =>
      loginAttempt.mLastFailureDate(now).mBadAttemptsSinceLastSuccessOrReset(0).save
    case _ =>
      // don't need to create here
      Empty
  }
}
```

## User Lock System

**Primary Implementation File:** [UserLocks.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/userlocks/UserLocks.scala)

The OBP-API implements an additional user locking system beyond login attempts.

### Lock Types

The system supports different types of user locks:

```scala
// From UserLocks.scala
class UserLocks extends UserLocksTrait with LongKeyedMapper[UserLocks] with IdPK {
  def getSingleton = UserLocks

  object UserId extends MappedUUID(this)
  object TypeOfLock extends MappedString(this, 100)
  object LastLockDate extends MappedDateTime(this)

  override def userId: String = UserId.get
  override def typeOfLock: String = TypeOfLock.get
  override def lastLockDate: Date = LastLockDate.get
}
```

**Lock Structure:**
- `userId`: User being locked
- `typeOfLock`: Type/reason for the lock
- `lastLockDate`: When the lock was applied

**Common Lock Types:**
- Security-related locks
- Administrative locks
- Temporary locks
- Compliance-related locks

### Lock Management

The system provides functions for managing user locks:

**Lock Operations:**
- Creating locks for users
- Checking if a user is locked
- Removing locks
- Updating lock information

**Integration with Authentication:**
- Authentication methods check for locks before processing
- Locked users are prevented from authenticating
- Lock status is logged for audit purposes

## Authentication Type Validation

**Primary Implementation File:** [AuthenticationTypeValidationProvider.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/authtypevalidation/AuthenticationTypeValidationProvider.scala)

The OBP-API implements a system for validating and restricting authentication methods for specific operations.

### Operation-Specific Authentication

The system allows defining which authentication methods are allowed for specific operations:

```scala
// From AuthenticationTypeValidationProvider.scala
case class JsonAuthTypeValidation(operationId: String, authTypes: List[AuthenticationType]) extends JsonAble {
  override def toJValue(implicit format: Formats): JsonAST.JValue =
    ("operation_id", operationId) ~ ("allowed_authentication_types", json.Extraction.decompose(authTypes.map(_.toString)))
}

object JsonAuthTypeValidation {
  def apply(operationId: String, authTypes: String): JsonAuthTypeValidation = {
    val typeList = StringUtils.split(authTypes, ",").toList.map(AuthenticationType.withName)
    JsonAuthTypeValidation(operationId, typeList)
  }
}
```

**Key Components:**
- `operationId`: API operation identifier
- `authTypes`: List of allowed authentication types
- Conversion between string and enum representations

### Authentication Type Restrictions

The system provides management functions for authentication type restrictions:

```scala
// From AuthenticationTypeValidationProvider.scala
trait AuthenticationTypeValidationProvider {
  def getByOperationId(operationId: String): Box[JsonAuthTypeValidation]
  def getAll(): List[JsonAuthTypeValidation]
  def create(jsonValidation: JsonAuthTypeValidation): Box[JsonAuthTypeValidation]
  def update(jsonValidation: JsonAuthTypeValidation): Box[JsonAuthTypeValidation]
  def deleteByOperationId(operationId: String): Box[Boolean]
}
```

**Key Operations:**
- Creating authentication type restrictions
- Retrieving restrictions by operation
- Updating allowed authentication types
- Deleting restrictions

**Integration with API Layer:**
- API endpoints check authentication type restrictions
- Requests with disallowed authentication methods are rejected
- Provides regulatory compliance for sensitive operations

## Consumer Management

**Primary Implementation File:** [ConsumerProvider.scala](https://github.com/OpenBankProject/OBP-API/blob/develop/obp-api/src/main/scala/code/consumer/ConsumerProvider.scala)

The OBP-API implements comprehensive OAuth consumer (API client) management with support for certificates, call limits, and multi-provider authentication.

### OAuth Consumer Registration

The system provides extensive consumer management capabilities:

```scala
// From ConsumerProvider.scala
trait ConsumersProvider {
  def getConsumerByPrimaryIdFuture(id: Long): Future[Box[Consumer]]
  def getConsumerByPrimaryId(id: Long): Box[Consumer]
  def getConsumerByConsumerKey(consumerKey: String): Box[Consumer]
  def getConsumerByConsumerKeyFuture(consumerKey: String): Future[Box[Consumer]]
  def getConsumerByPemCertificate(pem: String): Box[Consumer]
  def getConsumerByConsumerId(consumerId: String): Box[Consumer]
  def getConsumerByConsumerIdFuture(consumerId: String): Future[Box[Consumer]]
  def getConsumersByUserIdFuture(userId: String): Future[List[Consumer]]
  def getConsumersFuture(httpParams: List[OBPQueryParam], callContext: Option[CallContext]): Future[List[Consumer]]
  // Additional methods...
}
```

**Consumer Lookup Methods:**
- By primary key (internal ID)
- By consumer key (OAuth identifier)
- By PEM certificate (for certificate-based auth)
- By consumer ID (external identifier)
- By user ID (all consumers for a user)

### Consumer Verification

The system supports multiple consumer verification methods:

**Consumer Creation:**
```scala
// From ConsumerProvider.scala
def createConsumer(
  key: Option[String],
  secret: Option[String],
  isActive: Option[Boolean],
  name: Option[String],
  appType: Option[AppType],
  description: Option[String],
  developerEmail: Option[String],
  redirectURL: Option[String],
  createdByUserId: Option[String],
  clientCertificate: Option[String],
  company: Option[String],
  logoURL: Option[String]
): Box[Consumer]
```

**Consumer Updates:**
```scala
// From ConsumerProvider.scala
def updateConsumer(id: Long,
                   key: Option[String] = None,
                   secret: Option[String] = None,
                   isActive: Option[Boolean] = None,
                   name: Option[String] = None,
                   appType: Option[AppType] = None,
                   description: Option[String] = None,
                   developerEmail: Option[String] = None,
                   redirectURL: Option[String] = None,
                   createdByUserId: Option[String] = None,
                   LogoURL: Option[String] = None,
                   certificate: Option[String] = None,
): Box[Consumer]
```

**Call Limit Management:**
```scala
// From ConsumerProvider.scala
def updateConsumerCallLimits(id: Long, perSecond: Option[String], perMinute: Option[String], perHour: Option[String], perDay: Option[String], perWeek: Option[String], perMonth: Option[String]): Future[Box[Consumer]]
```

### Certificate-Based Authentication

The system supports certificate-based consumer authentication for PSD2 compliance:

**Certificate Lookup:**
- `getConsumerByPemCertificate` - Find consumer by certificate
- Certificate validation during authentication
- Support for eIDAS certificates

**Dynamic Consumer Creation:**
```scala
// From ConsumerProvider.scala
def getOrCreateConsumer(consumerId: Option[String],
                        key: Option[String],
                        secret: Option[String],
                        aud: Option[String],
                        azp: Option[String],
                        iss: Option[String],
                        sub: Option[String],
                        isActive: Option[Boolean],
                        name: Option[String],
                        appType: Option[AppType],
                        description: Option[String],
                        developerEmail: Option[String],
                        redirectURL: Option[String],
                        createdByUserId: Option[String],
                        certificate: Option[String] = None,
                        logoUrl: Option[String] = None
                       ): Box[Consumer]
```

**Key Features:**
- Automatic consumer creation for OIDC flows
- JWT claim-based consumer creation
- Certificate storage and validation
- Rate limiting per consumer

## Authentication Architecture

The OBP-API authentication system is built on a modular architecture that separates concerns and allows for flexible configuration:

**Core Components:**
1. **Authentication Handlers**: Implement specific authentication methods (OAuth, Direct Login, etc.)
2. **Token Management**: Handles token creation, validation, and expiration
3. **User Management**: Manages user creation, lookup, and association with authentication methods
4. **Consumer Management**: Handles API consumer (client) registration and verification
5. **Security Utilities**: Provides cryptographic functions, certificate validation, and security checks
6. **Entitlement System**: Manages role-based access control
7. **View Permission System**: Controls data access through views
8. **Consent Management**: Handles user consent for data access
9. **Authentication Context**: Stores metadata for authentication sessions
10. **Login Protection**: Prevents brute force attacks and manages user locks

**Authentication Flow:**
1. **Request Interception**: Incoming requests are intercepted by authentication filters
2. **Header Parsing**: Authentication headers are parsed and routed to appropriate handlers
3. **Token Validation**: Tokens are validated for authenticity, expiration, and permissions
4. **User Resolution**: Authenticated user is resolved from token information
5. **Consumer Verification**: API consumer is verified and rate limits are checked
6. **Entitlement Checking**: User entitlements are verified for the requested operation
7. **View Permission Checking**: View permissions are checked for data access
8. **Request Processing**: Authenticated request is processed by API endpoints

**Key Classes and Objects:**
- `APIUtil`: Contains authentication utilities and helper functions
- `OAuthHandshake`: Handles OAuth 1.0a authentication
- `OAuth2Login`: Manages OAuth 2.0 and OIDC authentication
- `DirectLogin`: Processes Direct Login authentication
- `GatewayLogin`: Handles Gateway Login authentication
- `AuthUser`: Represents authenticated web users
- `ResourceUser`: Represents API resource users
- `Consumer`: Represents API consumers/applications
- `Entitlement`: Represents user roles and permissions
- `View`: Represents data access views
- `Consent`: Represents user consent for data access

## Secure Random Generation System

### Overview

The OBP-API implements a cryptographically strong pseudo-random number generator for all security-critical operations. This system ensures that generated values are unpredictable and suitable for authentication tokens, challenges, and other security-sensitive data.

### Implementation

**File**: `code/api/util/SecureRandomUtil.scala`

The system uses Java's `SecureRandom` with the `NativePRNG` algorithm, which obtains random numbers from the underlying native OS:

```scala
object SecureRandomUtil {
  // Obtains random numbers from the underlying native OS
  val csprng = SecureRandom.getInstance("NativePRNG")

  def alphanumeric(nrChars: Int = 24): String = {
    new BigInteger(nrChars * 5, csprng).toString(32)
  }

  def numeric(maxNumber: Int = 99999999): String = {
    csprng.nextInt(maxNumber).toString()
  }
}
```

### Key Features

- **Cryptographically Strong**: Uses `SecureRandom` instead of standard `Random` for unpredictable values
- **Native OS Integration**: Leverages OS-level entropy sources for maximum randomness
- **Multiple Formats**: Supports both alphanumeric and numeric random generation
- **Configurable Length**: Allows customization of generated string length

### Usage Examples

```scala
// Generate 24-character alphanumeric string (default)
val token = SecureRandomUtil.alphanumeric()

// Generate custom length alphanumeric string
val customToken = SecureRandomUtil.alphanumeric(32)

// Generate numeric challenge code
val challengeCode = SecureRandomUtil.numeric()

// Generate numeric code with custom max value
val customChallenge = SecureRandomUtil.numeric(999999)
```

### Security Considerations

- **No Predictability**: Unlike Linear Congruential Generator (LCG) algorithms, values are cryptographically unpredictable
- **Attack Resistance**: Prevents attackers from predicting future values based on observed patterns
- **Compliance**: Meets security requirements for banking and financial applications

---

## X.509 Certificate Processing System

### Overview

The OBP-API provides comprehensive X.509 certificate processing capabilities, including validation, parsing, PSD2 role extraction, and certificate information extraction. This system is essential for regulatory compliance and secure client authentication.

### Implementation

**File**: `code/api/util/X509.scala`

### Core Functionality

#### Certificate Validation

```scala
def validate(encodedCert: String): Box[Boolean] = {
  val cert: X509Certificate = X509CertUtils.parse(encodedCert)
  if (cert == null) {
    Failure(ErrorMessages.X509ParsingFailed)
  } else {
    try {
      cert.checkValidity()
      Full(true)
    } catch {
      case _: CertificateExpiredException =>
        Failure(ErrorMessages.X509CertificateExpired)
      case _: CertificateNotYetValidException =>
        Failure(ErrorMessages.X509CertificateNotYetValid)
    }
  }
}
```

#### PSD2 Role Extraction

The system can extract PSD2-specific roles from qualified certificates:

```scala
def extractPsd2Roles(pem: String): Box[List[String]] = {
  val cert: X509Certificate = X509CertUtils.parse(pem)
  if (cert == null) {
    Failure(ErrorMessages.X509ParsingFailed)
  } else {
    try {
      val qcstatements = extractQcStatements(cert)
      val asn1encodable = extractPsd2QcStatements(qcstatements)
      Full(getPsd2Roles(asn1encodable: Array[ASN1Encodable]))
    } catch {
      case _: Throwable => Failure(ErrorMessages.X509ThereAreNoPsd2Roles)
    }
  }
}
```

#### Public Key Extraction

```scala
def getRSAPublicKey(encodedCert: String): Box[PublicKey] = {
  val cert: X509Certificate = X509CertUtils.parse(encodedCert)
  if (cert == null) {
    Failure(ErrorMessages.X509ParsingFailed)
  } else {
    val pubKey: PublicKey = cert.getPublicKey()
    if (pubKey.isInstanceOf[RSAPublicKey]) {
      Full(pubKey)
    } else {
      Failure(ErrorMessages.X509CannotGetRSAPublicKey)
    }
  }
}
```

### Key Features

- **Certificate Validation**: Checks certificate validity periods and parsing
- **PSD2 Compliance**: Extracts PSD2-specific roles from qualified certificates
- **Multi-Algorithm Support**: Handles both RSA and EC public keys
- **JWK Integration**: Converts certificates to JSON Web Key format
- **Subject Information**: Extracts common name, organization, and email addresses
- **ASN.1 Processing**: Handles complex ASN.1 structures for qualified statements

---

## JWS (JSON Web Signature) System

### Overview

The OBP-API implements a comprehensive JSON Web Signature (JWS) system for request signing, verification, and integrity protection. This system supports both attached and detached payload signatures with multiple cryptographic algorithms.

### Implementation

**File**: `code/api/util/JwsUtil.scala`

### Core Functionality

#### JWS Verification

```scala
def verifyJws(jwsString: String, publicKey: PublicKey): Boolean = {
  try {
    val jwsObject = JWSObject.parse(jwsString)
    val verifier = new RSASSAVerifier(publicKey.asInstanceOf[RSAPublicKey])
    jwsObject.verify(verifier)
  } catch {
    case _: Exception => false
  }
}
```

#### Detached Payload Handling

```scala
def rebuildJwsWithDetachedPayload(jwsString: String, payload: String): String = {
  val parts = jwsString.split("\\.")
  if (parts.length == 3) {
    val encodedPayload = Base64.getUrlEncoder.withoutPadding().encodeToString(payload.getBytes("UTF-8"))
    s"${parts(0)}.$encodedPayload.${parts(2)}"
  } else {
    jwsString
  }
}
```

### Key Features

- **Multiple Algorithms**: Supports RS256, RS384, RS512, HS256, HS384, HS512
- **Detached Payloads**: Handles JWS with detached payloads for large requests
- **Request Signing**: Signs HTTP requests with critical headers
- **Certificate Integration**: Works with X.509 certificates for public key extraction
- **Digest Verification**: Validates request body digests

---

## Post-Authentication Processing

### Overview

The post-authentication processing system handles all actions that occur after successful user authentication, including user initialization, rate limiting checks, consumer validation, and security verifications.

### Implementation

**File**: `code/api/util/AfterApiAuth.scala`

### Core Processing Pipeline

```scala
def processAfterAuthentication(user: User, consumer: Consumer, callContext: CallContext): Box[CallContext] = {
  for {
    _ <- validateUserStatus(user)
    _ <- checkRateLimits(consumer, callContext)
    _ <- verifyConsumerStatus(consumer)
    _ <- checkUserLocks(user)
    updatedContext <- initializeUserSession(user, consumer, callContext)
  } yield updatedContext
}
```

### Key Processing Steps

- **User Status Validation**: Verifies user account is active and not suspended
- **Consumer Status**: Checks consumer application is valid and active
- **Rate Limits**: Enforces API call limits per consumer
- **User Locks**: Prevents access for locked user accounts
- **Session Security**: Initializes secure session context

---

## Certificate Management Utilities

### Overview

The certificate management utilities provide comprehensive cryptographic operations including RSA key management, X.509 certificate handling, PEM normalization, JWT encryption/decryption, and HMAC protection.

### Implementation

**File**: `code/api/util/CertificateUtil.scala`

### Core Functionality

#### RSA Key Pair Generation

```scala
def generateRSAKeyPair(keySize: Int = 2048): KeyPair = {
  val keyPairGenerator = KeyPairGenerator.getInstance("RSA")
  keyPairGenerator.initialize(keySize)
  keyPairGenerator.generateKeyPair()
}
```

#### JWT Signing with HMAC

```scala
def jwtWithHmacProtection(claims: JWTClaimsSet, secret: String): String = {
  val signer = new MACSigner(secret.getBytes("UTF-8"))
  val jwsObject = new JWSObject(
    new JWSHeader(JWSAlgorithm.HS256),
    new Payload(claims.toJSONObject)
  )
  jwsObject.sign(signer)
  jwsObject.serialize()
}
```

### Key Features

- **RSA Key Management**: Generation, loading, and manipulation of RSA key pairs
- **Certificate Operations**: Loading from keystores, self-signed generation, PEM normalization
- **JWT Protection**: Both HMAC and RSA signing/verification
- **JWT Encryption**: RSA-based encryption and decryption of JWTs

---

## JWT Processing Utilities

### Overview

The JWT processing utilities provide comprehensive JWT validation, claim extraction, algorithm support, access token validation, and ID token validation for various authentication flows.

### Implementation

**File**: `code/api/util/JwtUtil.scala`

### Core Functionality

#### JWT Parsing and Validation

```scala
def parseJWT(jwtString: String): Option[JWTClaimsSet] = {
  try {
    val jwt = SignedJWT.parse(jwtString)
    Some(jwt.getJWTClaimsSet)
  } catch {
    case _: Exception => None
  }
}
```

#### OAuth2 Access Token Validation

```scala
def validateOAuth2AccessToken(accessToken: String, jwkSetUrl: String): Future[Boolean] = {
  for {
    jwkSet <- fetchJWKSet(jwkSetUrl)
    isValid <- validateTokenAgainstJWKSet(accessToken, jwkSet)
  } yield isValid
}
```

### Key Features

- **Multi-Algorithm Support**: RSA, HMAC, and EC signature algorithms
- **Remote Validation**: Validates tokens against remote JWK sets
- **Comprehensive Claims**: Extracts standard and custom JWT claims
- **OIDC Compliance**: Full OpenID Connect ID token validation

---

## Hydra Integration System

### Overview

The Hydra integration system provides seamless integration with ORY Hydra OAuth2 server, including client management, JWK generation, and token endpoint configuration for enterprise-grade OAuth2 flows.

### Implementation

**File**: `code/util/HydraUtil.scala`

### Core Functionality

#### Hydra Client Management

```scala
def createHydraClient(clientId: String, clientSecret: String, redirectUris: List[String], grantTypes: List[String]): Future[Box[HydraClient]] = {
  val client = HydraClientRequest(
    client_id = Some(clientId),
    client_secret = Some(clientSecret),
    redirect_uris = redirectUris,
    grant_types = grantTypes,
    response_types = List("code"),
    scope = "openid profile email"
  )

  postToHydra("/clients", client.toJson) map { response =>
    response match {
      case Full(json) => parseHydraClientResponse(json)
      case failure => failure
    }
  }
}
```

### Key Features

- **Client Lifecycle**: Create, update, and delete OAuth2 clients
- **JWK Management**: Generate and manage JSON Web Key sets
- **Flow Control**: Handle login and consent flows
- **Token Management**: Introspect and validate tokens

---

## Consumer Registration System

### Overview

The consumer registration system provides a web-based interface for OAuth2 client registration with certificate validation and multi-provider integration support.

### Implementation

**File**: `code/snippet/ConsumerRegistration.scala`

### Key Features

- **Web Interface**: User-friendly registration form
- **Multi-Provider**: Supports multiple OAuth2 providers
- **Certificate Support**: Optional client certificate authentication
- **Email Integration**: Automated notification system
- **Validation**: Comprehensive input and certificate validation

---

## OpenID Connect Invocation System

### Overview

The OpenID Connect invocation system handles OIDC authentication flow initiation with state management, nonce generation, and response mode configuration for secure authentication flows.

### Implementation

**File**: `code/snippet/OpenidConnectInvoke.scala`

### Key Features

- **Multi-Provider Support**: Supports Google, Azure, Keycloak, Hydra, and custom providers
- **Secure Flow Management**: Proper state and nonce handling
- **Session Management**: Secure session storage with expiration
- **Error Handling**: Comprehensive error handling and user feedback

---

## OAuth Authorization Handling

### Overview

The OAuth authorization handling system manages the OAuth 1.0a authorization flow with token verification, user authentication, and secure redirection handling.

### Implementation

**File**: `code/snippet/OAuthAuthorisation.scala`

### Key Features

- **OAuth 1.0a Compliance**: Full OAuth 1.0a authorization flow
- **User Integration**: Seamless integration with user authentication
- **Secure Redirects**: Safe handling of callback URLs
- **Error Handling**: Comprehensive error handling and user feedback

---

## Session Management System

### Overview

The session management system provides comprehensive session handling, CSRF protection, authentication type detection, and call context management for secure API operations.

### Implementation

**File**: `code/api/util/ApiSession.scala`

### Core Data Structures

#### CallContext

```scala
case class CallContext(
  gatewayLoginRequestPayload: Option[PayloadOfJwtJSON] = None,
  gatewayLoginResponseHeader: Option[String] = None,
  dauthRequestPayload: Option[JSONFactoryDAuth.PayloadOfJwtJSON] = None,
  dauthResponseHeader: Option[String] = None,
  user: Box[User] = Empty,
  consumer: Box[Consumer] = Empty,
  sessionId: Option[String] = None,
  correlationId: String = "",
  // ... additional fields
)
```

### Authentication Type Detection

```scala
def authType: AuthenticationType = {
  if(hasGatewayHeader(authReqHeaderField)) {
    GatewayLogin
  } else if(requestHeaders.exists(_.name==DAuthHeaderKey)) {
    DAuth
  } else if(has2021DirectLoginHeader(requestHeaders)) {
    DirectLogin
  } else if(hasAnOAuthHeader(authReqHeaderField)) {
    AuthenticationType.`OAuth1.0a`
  } else if(hasAnOAuth2Header(authReqHeaderField)) {
    OAuth2_OIDC
  } else {
    Anonymous
  }
}
```

### Key Features

- **Multi-Authentication Support**: Handles all authentication methods
- **Context Management**: Rich context information for API calls
- **Session Security**: Secure session handling with expiration
- **Request Correlation**: Tracks requests across system boundaries

---

## Rate Limiting System

### Overview

The rate limiting system provides traffic control and quota management for authentication endpoints with consumer-based limits, supporting multiple time windows and granular control.

### Implementation

**Files**:
- `code/ratelimiting/RateLimiting.scala`
- `code/ratelimiting/MappedRateLimiting.scala`

### Core Data Structure

```scala
trait RateLimitingTrait {
  def rateLimitingId: String
  def consumerId: String
  def perSecondCallLimit: Long
  def perMinuteCallLimit: Long
  def perHourCallLimit: Long
  def perDayCallLimit: Long
  def perWeekCallLimit: Long
  def perMonthCallLimit: Long
  def fromDate: Date
  def toDate: Date
}
```

### Key Features

- **Multi-Window Limits**: Supports six different time windows
- **Hierarchical Resolution**: Intelligent fallback strategy for rate limit lookup
- **Consumer-Specific**: Individual limits per OAuth consumer
- **API-Specific**: Granular limits per API endpoint or version
- **Date-Based**: Time-bound rate limit configurations

---

## Webhook/Notification System

### Overview

The webhook/notification system provides event-driven notifications for authentication events with HTTP callbacks, supporting account-based webhooks and system notifications.

### Implementation

**Files**:
- `code/webhook/AccountWebhook.scala`
- `code/webhook/MappedAccountWebhook.scala`

### Core Data Structure

```scala
trait AccountWebhook {
  def accountWebhookId: String
  def bankId: String
  def accountId: String
  def triggerName: String
  def url: String
  def httpMethod: String
  def httpProtocol: String
  def createdByUserId: String
  def isActive(): Boolean
}
```

### Key Features

- **Event-Driven**: Triggers on authentication and account events
- **HTTP Callbacks**: Supports POST, PUT, PATCH methods
- **Flexible Triggers**: Multiple trigger types for different events
- **User-Specific**: Webhooks created and managed by users
- **Security**: HMAC signature verification for webhook authenticity

---

## Certificate Verification System

### Overview

The certificate verification system provides SSL/TLS certificate validation, trust store management, and PKIX validation for secure client authentication and regulatory compliance.

### Implementation

**File**: `code/api/util/CertificateVerifier.scala`

### Core Functionality

#### Trust Store Loading

```scala
private def loadTrustStore(): Option[KeyStore] = {
  val trustStorePath = APIUtil.getPropsValue("truststore.path.tpp_signature")
    .or(APIUtil.getPropsValue("truststore.path")).getOrElse("")
  val trustStorePassword = APIUtil.getPropsValue("truststore.password.tpp_signature", "").toCharArray

  Try {
    val trustStore = KeyStore.getInstance("PKCS12")
    val trustStoreInputStream = new FileInputStream(trustStorePath)
    try {
      trustStore.load(trustStoreInputStream, trustStorePassword)
    } finally {
      trustStoreInputStream.close()
    }
    trustStore
  } match {
    case Success(store) => Some(store)
    case Failure(e) => None
  }
}
```

#### Certificate Validation

```scala
def validateCertificate(certificate: X509Certificate): Boolean = {
  Try {
    val trustStore = loadTrustStore().getOrElse(throw new Exception("Trust store could not be loaded."))
    val trustManagerFactory = TrustManagerFactory.getInstance(TrustManagerFactory.getDefaultAlgorithm)
    trustManagerFactory.init(trustStore)

    val trustAnchors = trustStore.aliases().asScala
      .filter(trustStore.isCertificateEntry)
      .map(alias => trustStore.getCertificate(alias).asInstanceOf[X509Certificate])
      .map(cert => new TrustAnchor(cert, null))
      .toSet.asJava

    val pkixParams = new PKIXParameters(trustAnchors)
    pkixParams.setRevocationEnabled(APIUtil.getPropsAsBoolValue("use_tpp_signature_revocation_list", defaultValue = true))

    val certPath = CertificateFactory.getInstance("X.509").generateCertPath(Collections.singletonList(certificate))
    val validator = CertPathValidator.getInstance("PKIX")
    validator.validate(certPath, pkixParams)

    true
  } match {
    case Success(_) => true
    case Failure(_) => false
  }
}
```

### Key Features

- **PKCS12 Trust Store**: Supports PKCS12 trust store format
- **PKIX Validation**: Full PKIX certificate path validation
- **Revocation Checking**: Configurable CRL and OCSP checking
- **Trust Anchor Management**: Automatic trust anchor extraction
- **Chain Validation**: Validates complete certificate chains

---

## Configuration

The OBP-API authentication system is highly configurable through properties files. Key configuration properties include:

**OAuth 1.0a Configuration:**
```properties
# Enable/disable OAuth 1.0a
allow_oauth1_login=true

# OAuth token expiration time in seconds (default: 1 hour)
oauth1.token_expiration_seconds=3600
```

**OAuth 2.0 / OIDC Configuration:**
```properties
# Enable/disable OAuth 2.0
allow_oauth2_login=true

# JWKS URL for token validation
oauth2.jwk_set.url=https://identity-provider/.well-known/jwks.json

# OAuth 2.0 identity provider settings
oauth2.client_id=client-id
oauth2.client_secret=client-secret
oauth2.redirect_url=https://api.example.com/oauth2/callback

# Token validation settings
oauth2.use_introspection=false
oauth2.introspection_url=https://identity-provider/oauth2/introspect
```

**Direct Login Configuration:**
```properties
# Enable/disable Direct Login
allow_direct_login=true

# Direct Login token expiration time in seconds (default: 4 weeks)
direct_login_token_expiration_seconds=2419200
```

**Gateway Login Configuration:**
```properties
# Enable/disable Gateway Login
allow_gateway_login=true

# JWT signing secret
gateway.token_secret=your-secret-key

# Gateway host whitelist
gateway.host_whitelist=api.example.com,localhost
```

**DAuth Configuration:**
```properties
# Enable/disable DAuth
allow_dauth=true

# DAuth token expiration time in seconds
dauth.token_expiration_seconds=3600
```

**Login Protection Configuration:**
```properties
# Maximum bad login attempts before lockout
max.bad.login.attempts=5

# User lock duration in seconds
user.lock.duration.seconds=1800
```

**Entitlement Configuration:**
```properties
# Enable entitlement checking
entitlement.enabled=true

# Default entitlements for new users
default.entitlements=CanGetUser,CanGetCustomer
```

**View Permission Configuration:**
```properties
# Enable view permission checking
view.permission.enabled=true

# Default system views
system.views=owner,accountant,auditor,public
```

**Consent Configuration:**
```properties
# Enable consent management
consent.enabled=true

# Consent token expiration in seconds
consent.token.expiration.seconds=3600

# Berlin Group consent settings
berlin.group.consent.enabled=true
berlin.group.max.frequency.per.day=4

# UK Open Banking consent settings
uk.open.banking.consent.enabled=true
uk.open.banking.transaction.history.days=90
```

**Authentication Context Configuration:**
```properties
# Enable authentication context storage
auth.context.enabled=true

# Context retention period in days
auth.context.retention.days=30
```

**Security Configuration:**
```properties
# SSL certificate validation
require_client_certificate=false
client_certificate_keystore_path=/path/to/keystore.jks
client_certificate_keystore_password=password

# Rate limiting
rate_limiting.enabled=true
rate_limiting.anonymous_access_per_minute=100
rate_limiting.authenticated_access_per_minute=1000
```

For detailed OIDC configuration, refer to the [OBP OIDC Configuration Guide](https://github.com/OpenBankProject/OBP-API/blob/develop/OBP_OIDC_Configuration_Guide.md).

## Security Features

The OBP-API authentication system implements numerous security features to protect against common threats:

**Token Security:**
- JWT signature validation (HMAC and RSA)
- Token expiration and refresh mechanisms
- Nonce validation to prevent replay attacks
- Secure token storage recommendations

**Transport Security:**
- TLS/SSL encryption for all communications
- Certificate validation for client authentication
- Secure header handling and validation

**Access Control:**
- Role-based access control through entitlements
- Fine-grained API permissions
- View-based data access control
- Consumer-based scope restrictions

**Threat Protection:**
- Rate limiting to prevent brute force attacks
- Login attempt tracking and user lockout
- IP-based filtering options
- Request logging and audit trails

**User Security:**
- Password hashing and validation
- Multi-factor authentication support
- Session management and timeout
- Account lockout protection
- User agreement tracking
- Data scrambling for GDPR compliance

**Access Control Security:**
- Role-based access control through entitlements
- View-based data access permissions
- Consumer-based scope restrictions
- Operation-specific authentication type validation
- Dynamic permission management

**Compliance Features:**
- PSD2 compliant authentication flows
- Support for eIDAS certificates
- Strong customer authentication options
- Consent management for data access
- GDPR compliance for user data
- Berlin Group PSD2 compliance
- UK Open Banking compliance
- Regulatory audit trails

## Implementation Details

### OAuth 1.0a Implementation

The OAuth 1.0a implementation follows the standard three-legged OAuth flow:

**Request Token Endpoint:**
```scala
def requestToken = {
  // Validate consumer key and signature
  // Generate request token
  // Return token and secret
}
```

**Authorization Endpoint:**
```scala
def authorizeToken = {
  // Display authorization form to user
  // User approves or denies access
  // Redirect to callback URL with verifier
}
```

**Access Token Endpoint:**
```scala
def accessToken = {
  // Validate request token and verifier
  // Generate access token
  // Return token and secret
}
```

**Signature Verification:**
```scala
def verifySignature(request, consumerSecret, tokenSecret) = {
  // Calculate signature base string
  // Apply HMAC-SHA1 or HMAC-SHA256
  // Compare with provided signature
}
```

### OAuth 2.0 / OIDC Implementation

The OAuth 2.0 and OIDC implementation supports multiple identity providers and token types:

**Token Validation:**
```scala
def validateToken(token: String): Box[User] = {
  // Decode JWT token
  // Verify signature using JWKS
  // Validate claims (iss, aud, exp, etc.)
  // Return user information
}
```

**ID Token Processing:**
```scala
def processIdToken(idToken: String): Box[User] = {
  // Validate ID token signature
  // Extract user information
  // Create or update user
  // Return user object
}
```

**Access Token Introspection:**
```scala
def introspectToken(token: String): Box[TokenInfo] = {
  // Call introspection endpoint
  // Validate response
  // Return token information
}
```

### Direct Login Implementation

The Direct Login implementation provides a simplified authentication method:

**Token Generation:**
```scala
def createToken(username: String, password: String, consumerKey: String): Box[String] = {
  // Validate username and password
  // Verify consumer key
  // Generate JWT token
  // Return token
}
```

**Token Validation:**
```scala
def validateToken(token: String): Box[User] = {
  // Decode JWT token
  // Verify signature
  // Validate expiration
  // Return user information
}
```

### Gateway Login Implementation

The Gateway Login implementation enables integration with Core Banking Systems:

**JWT Validation:**
```scala
def validateJwt(token: String): Box[PayloadOfJwtJSON] = {
  // Decode JWT token
  // Verify signature
  // Validate claims
  // Return payload
}
```

**User Creation/Update:**
```scala
def getOrCreateResourceUser(payload: PayloadOfJwtJSON): Box[User] = {
  // Extract user information from payload
  // Create or update user
  // Update account information
  // Return user object
}
```

### DAuth Implementation

The DAuth implementation provides dynamic client registration and authentication:

**Client Registration:**
```scala
def registerClient(clientInfo: ClientInfo): Box[ClientRegistration] = {
  // Validate client information
  // Generate client ID and secret
  // Store client registration
  // Return registration information
}
```

**Token Issuance:**
```scala
def issueToken(clientId: String, clientSecret: String): Box[String] = {
  // Validate client credentials
  // Generate JWT token
  // Return token
}
```

### User Management Implementation

The user management system handles user creation and authentication:

**User Creation:**
```scala
def createResourceUser(provider: String, providerId: Option[String], ...): Box[ResourceUser] = {
  val ru = ResourceUser.create
  ru.provider_(provider)
  // Set additional fields
  Full(ru.saveMe())
}
```

**User Authentication:**
```scala
def getUserByProviderId(provider: String, idGivenByProvider: String): Box[User] = {
  ResourceUser.find(
    By(ResourceUser.provider_, provider),
    By(ResourceUser.providerId, idGivenByProvider)
  )
}
```

### Entitlement Implementation

The entitlement system manages role-based access control:

**Adding Entitlement:**
```scala
def addEntitlement(bankId: String, userId: String, roleName: String, ...): Box[Entitlement] = {
  // Check permissions
  // Create entitlement
  // Send notification
  // Return entitlement
}
```

**Checking Entitlement:**
```scala
def hasEntitlement(bankId: String, userId: String, roleName: String): Boolean = {
  getEntitlement(bankId, userId, roleName).isDefined
}
```

### View Permission Implementation

The view permission system controls data access:

**Granting View Access:**
```scala
def grantAccessToCustomView(bankIdAccountIdViewId: BankIdAccountIdViewId, user: User): Box[View] = {
  // Validate view exists
  // Check permissions
  // Grant access
  // Return view
}
```

**Checking View Permission:**
```scala
def hasViewAccess(user: User, view: View): Boolean = {
  permission(view.bankIdAccountId, user).isDefined
}
```

### Consent Implementation

The consent system manages user consent for data access:

**Creating Consent:**
```scala
def createObpConsent(user: User, challengeAnswer: String, consentRequestId: Option[String], consumer: Option[Consumer]): Box[MappedConsent] = {
  // Validate challenge
  // Create consent
  // Generate JWT
  // Return consent
}
```

**Validating Consent:**
```scala
def validateConsentJWT(jwt: String): Box[ConsentJWT] = {
  // Decode JWT
  // Verify signature
  // Check expiration
  // Return consent claims
}
```

### Login Attempt Protection Implementation

The login attempt protection system prevents brute force attacks:

**Tracking Failed Attempts:**
```scala
def incrementBadLoginAttempts(provider: String, username: String): Unit = {
  // Find existing record
  // Increment counter
  // Update timestamp
  // Save record
}
```

**Checking User Lock Status:**
```scala
def userIsLocked(provider: String, username: String): Boolean = {
  // Check bad login attempts
  // Check manual locks
  // Return lock status
}
```

### Authentication Context Implementation

The authentication context system stores session metadata:

**Creating Context:**
```scala
def createUserAuthContext(userId: String, key: String, value: String, consumerId: String): Future[Box[UserAuthContext]] = {
  // Create context entry
  // Associate with user and consumer
  // Save to database
  // Return context
}
```

**Retrieving Context:**
```scala
def getUserAuthContexts(userId: String): Future[Box[List[UserAuthContext]]] = {
  // Query by user ID
  // Return all context entries
}
```

## Error Handling

The OBP-API authentication system implements comprehensive error handling to provide clear feedback on authentication failures:

**Common Error Scenarios:**
- Invalid credentials (username/password)
- Expired or invalid tokens
- Missing or malformed authentication headers
- Unauthorized access attempts
- Rate limiting exceeded
- User lockout due to failed attempts
- Missing or invalid entitlements
- View permission denied
- Consent expired or revoked
- Authentication type not allowed for operation
- Consumer not found or inactive
- Scope restrictions violated

**Error Response Format:**
```json
{
  "error": "invalid_token",
  "error_description": "The access token expired"
}
```

**Error Logging:**
- Authentication errors are logged with appropriate severity
- Sensitive information is redacted from logs
- Error patterns are monitored for security threats
- Login attempts are tracked for security analysis

**Error Recovery:**
- Token refresh mechanisms
- Clear instructions for re-authentication
- Graceful degradation for partial authentication failures
- User lockout reset procedures

## Testing

The OBP-API includes comprehensive testing tools for authentication:

**Authentication Test Endpoints:**
- `/test/authentication` - Test current authentication status
- `/test/authentication/methods` - List available authentication methods

**Testing Utilities:**
- Test consumers and users for development
- Authentication bypass options for testing
- Mock identity providers for OAuth 2.0 testing
- Entitlement testing helpers

**Test Configuration:**
```properties
# Enable test authentication endpoints
allow_test_authentication_endpoints=true

# Test consumer key and secret
test.consumer_key=test-consumer-key
test.consumer_secret=test-consumer-secret
```

**Authentication Testing Methods:**
- Unit tests for authentication components
- Integration tests for authentication flows
- Performance tests for authentication endpoints
- Security tests for authentication vulnerabilities
- Entitlement testing with role assignments
- View permission testing with access controls
- Consent testing with multiple standards
- Login attempt protection testing
- User lifecycle testing with GDPR compliance

## Troubleshooting

Common authentication issues and their solutions:

**OAuth 1.0a Issues:**
- **Invalid signature**: Ensure correct signature method and parameters
- **Token expired**: Request a new token
- **Consumer not found**: Verify consumer key is registered

**OAuth 2.0 / OIDC Issues:**
- **Invalid token**: Check token expiration and signature
- **JWKS retrieval failed**: Verify JWKS URL and connectivity
- **Invalid claims**: Check issuer, audience, and scope configuration

**Direct Login Issues:**
- **Authentication failed**: Verify username, password, and consumer key
- **Token expired**: Request a new token
- **Invalid token format**: Ensure correct JWT format

**Gateway Login Issues:**
- **Invalid JWT**: Check signature and claims
- **Missing required fields**: Ensure all required claims are present
- **Host not whitelisted**: Add host to gateway whitelist

**User Management Issues:**
- **User not found**: Verify user exists and is not deleted
- **User locked**: Check login attempts and user locks
- **Provider mismatch**: Ensure correct provider is specified
- **User agreement missing**: Check required agreements are signed
- **GDPR compliance**: Verify data handling for deleted users

**Entitlement Issues:**
- **Missing entitlement**: Verify user has required entitlement
- **Entitlement creation failed**: Check permissions of grantor
- **Bank-specific entitlement**: Ensure correct bank ID is specified
- **Dynamic entity entitlement**: Check entity-specific roles
- **Entitlement notification failed**: Verify email configuration

**View Permission Issues:**
- **View not found**: Verify view exists for the account
- **Permission denied**: Check user has access to the view
- **System view access**: Verify system view permissions
- **Custom view creation**: Check view creation permissions
- **View permission inheritance**: Verify permission propagation

**Consent Issues:**
- **Consent expired**: Check consent validity period
- **Consent revoked**: Verify consent status
- **Invalid consent JWT**: Check JWT signature and claims
- **Berlin Group consent**: Verify frequency limits and usage
- **UK Open Banking consent**: Check permissions and account access

**Authentication Context Issues:**
- **Context not found**: Verify context exists for user/consent
- **Context creation failed**: Check storage permissions
- **Context retention**: Verify retention policy compliance

**Scope Issues:**
- **Scope restriction**: Check consumer has required scope
- **Scope creation failed**: Verify consumer permissions
- **OAuth scope mapping**: Check scope to role mapping

**Authentication Type Validation Issues:**
- **Authentication type not allowed**: Check operation restrictions
- **Validation rule not found**: Verify validation configuration
- **Type restriction update failed**: Check admin permissions

**General Troubleshooting Steps:**
1. Check authentication headers for correct format
2. Verify token expiration and validity
3. Confirm consumer registration and permissions
4. Check user status (locked, deleted)
5. Verify entitlements for the operation
6. Check view permissions for data access
7. Review server logs for detailed error information
8. Verify configuration properties for authentication methods

For detailed troubleshooting, enable debug logging:
```properties
logging.level.code.api.oauth=DEBUG
logging.level.code.api.directlogin=DEBUG
logging.level.code.api.OAuth2=DEBUG
logging.level.code.api.GatewayLogin=DEBUG
logging.level.code.users=DEBUG
logging.level.code.entitlement=DEBUG
logging.level.code.loginattempts=DEBUG
logging.level.code.views=DEBUG
logging.level.code.consent=DEBUG
logging.level.code.context=DEBUG
logging.level.code.scope=DEBUG
logging.level.code.consumer=DEBUG
logging.level.code.authtypevalidation=DEBUG
logging.level.code.userlocks=DEBUG
```

**Performance Monitoring:**
- Monitor authentication response times
- Track failed authentication rates
- Monitor user lockout frequency
- Track consent usage patterns
- Monitor entitlement assignment patterns

**Security Monitoring:**
- Monitor for brute force attacks
- Track suspicious authentication patterns
- Monitor certificate validation failures
- Track consent revocation patterns
- Monitor for privilege escalation attempts
