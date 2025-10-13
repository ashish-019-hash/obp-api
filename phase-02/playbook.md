# Golang Migration Project - Step-by-Step Guide

## Core Rules (Apply to Every Step)

### Work from Documentation Only
All your work must come from the markdown files in the `04.phase-2-input` folder. These files are your single source of truth:
- business_rules.md - How the business logic works
- user_stories.md - What users need to do
- business_entities.md - What data exists and what fields it has
- validation-rules.md - What rules apply to each field

**Important**: Only implement what's written in these files. Don't add features, don't skip things, don't make assumptions. If something is unclear, note it in your report and ask.

### One Step at a Time
You'll get instructions for one step at a time. Complete that step fully, then stop and wait for the next step. Never jump ahead or work on multiple steps together.

### Always Use Branches
Never work on the main branch. Always create a new branch for your work with a name like `devin/step-1-project-skeleton`. When done, create a Pull Request.

### Be Exact
If the documentation says an entity has 12 fields, create exactly 12 fields. Keep the same names, same order, same types. Don't reorganize, don't optimize, don't improve unless it's documented.

### Communicate Clearly
Write your reports in simple language. Explain what you did and why. If you hit problems, describe them clearly. List any documentation that was confusing or incomplete.

---

## Step 1: Create Project Structure

### Your Role: Project Setup Specialist
You're setting up the foundation for the entire project. Your job is to create a clean, organized folder structure and a basic server that starts up successfully. You're NOT adding business logic or features yet - just the skeleton.

### Which Files to Use
For this step, you're just creating the structure. You don't need to implement from the .md files yet, but you should review them to understand what's coming:
- **All files in 04.phase-2-input/** - Familiarize yourself with what exists, but don't implement yet

### What You're Building
A basic Golang project that:
- Has the correct folder organization
- Can compile without errors
- Starts a simple web server
- Responds to a basic health check

### Your Tasks

1. **Initialize the Go Module**
   - Set up the project as a Go module in the `Migrated_Code/backend` folder
   - Choose an appropriate module name based on the repository

2. **Create the Folder Structure**
   Create these folders exactly as listed:
   - `cmd/` - Where the main application starts
   - `internal/config/` - Configuration settings
   - `internal/controllers/` - Request handlers
   - `internal/middleware/` - HTTP middleware
   - `internal/models/` - Data structures
   - `internal/repositories/` - Database operations
   - `internal/routes/` - URL routing
   - `internal/services/` - Business logic
   - `internal/utils/` - Helper functions
   - `pkg/db/` - Database connection helpers
   - `test/` - Test files

3. **Set Up Basic Server**
   - Pick either the standard Go web library or Gin framework
   - Create a simple main file that starts the server
   - Add one health check endpoint that returns "ok"
   - Make sure the server starts on port 8080

4. **Verify Everything Works**
   - Build the project and confirm no errors
   - Run the server and confirm it starts
   - Test the health check endpoint and confirm it responds

### What to Deliver

Provide a report with:
1. List of all folders and files you created
2. Commands to build and run the server
3. A test command to check the health endpoint and what response to expect
4. The branch name you used and your commit message
5. Confirmation that everything compiles and runs
6. Any questions about unclear documentation

Then stop and wait for Step 2.

---

## Step 2: Create Data Models

### Your Role: Data Structure Designer
You're translating the business entities from the documentation into Go code. Your job is to create accurate representations of every piece of data the system needs to track. Think of yourself as a careful translator - converting documentation into code without adding your own interpretation.

### Which Files to Use
- **business_entities.md** - Primary source for all entity definitions, field names, and types
- **validation-rules.md** - Field-level constraints and requirements

### What You're Building
Go struct definitions for every entity in the documentation. Each struct will represent one business entity with all its fields, types, and rules exactly as documented.

### Your Tasks

1. **Study the Entity Documentation**
   - Read all entity definitions thoroughly
   - Make a list of every entity you need to create
   - For each entity, note the field names, types, and any special rules

2. **Map Documentation Types to Go Types**
   - Text becomes string
   - Numbers become int or float64
   - Yes/No becomes bool
   - Dates become time.Time
   - If a field is optional, use a pointer type

3. **Create One File Per Entity**
   - Put each entity in its own file in the `internal/models/` folder
   - Name files after the entity (like user.go, account.go)
   - Keep fields in the same order as the documentation
   - Add tags for JSON field names that match the documentation exactly

4. **Add Documentation**
   - Write a brief comment above each struct describing what it represents
   - Add comments for fields when the documentation explains them

5. **Create Constructors If Needed**
   - If the documentation specifies default values or initialization logic, create a function to build the struct with those defaults

6. **Create a Simple Test**
   - Make a test file that creates one instance of each struct
   - This just confirms everything compiles correctly

### What to Deliver

Provide a report with:
1. A table listing all entities you created (entity name, filename, number of fields)
2. One complete example showing how you structured a struct
3. Notes on any type decisions you made
4. Confirmation that your compilation test passes
5. List of all files you created
6. The branch name and commit message
7. Any unclear parts of the documentation

Then stop and wait for Step 3.

---

## Step 3: Build Business Logic and Data Access

### Your Role: Backend Logic Developer
You're implementing how the system actually works. You'll create two layers: services that contain business rules and workflows, and repositories that handle database operations. Think of services as the "brain" that makes decisions, and repositories as the "hands" that save and retrieve data.

### Which Files to Use
- **business_rules.md** - All business logic, calculations, and decision rules

### What You're Building
- **Service Layer**: Functions that execute business rules, validate data, and orchestrate workflows
- **Repository Layer**: Functions that save, retrieve, update, and delete data from the database

These layers work together but stay separate. Services call repositories, but repositories never call services.

### Your Tasks

1. **Set Up Database**
   - Create a helper to connect to an in-memory SQLite database
   - This is just for development and testing

2. **Design Repository Interfaces**
   - For each entity that needs database operations, create an interface
   - List all the operations needed (create, read, update, delete, search, etc.)
   - Base this on what the documentation says the system needs to do

3. **Implement Repositories**
   - Write the actual database code for each operation
   - Handle basic errors (like "not found")
   - Keep repositories simple - no business logic here

4. **Implement Services**
   - For each business workflow in the documentation, create a service function
   - Follow any step-by-step processes exactly as documented
   - Apply all validation rules from the documentation
   - Implement calculations exactly as specified
   - Call repository functions to save or retrieve data

5. **Write Tests**
   - Test repository operations with the in-memory database
   - Test service functions by mocking the repositories
   - Test all the business rules mentioned in the documentation

### What to Deliver

Provide a report with:
1. A table listing all services and repositories (name, purpose, what they depend on)
2. List of all service methods and which business rules they implement
3. List of all repository methods and what data they access
4. Example of one complete service method showing how it works
5. Test results - all tests should pass
6. How the database is set up
7. The branch name and commit message
8. Any unclear business rules or workflows

Then stop and wait for Step 4.

---

## Step 4: Create API Endpoints

### Your Role: API Developer
You're exposing the business logic through HTTP endpoints that external applications can call. Your job is to create clean, well-defined REST APIs that match the documentation exactly - every URL, every field name, every status code.

### Which Files to Use
- **user_stories.md** - User actions that map to API endpoints
- **story_points.md** - Additional endpoint specifications if available

### What You're Building
A complete REST API with:
- HTTP endpoints for all documented operations
- Request and response handling
- Proper error responses
- Clean URL structure

### Your Tasks

1. **Plan Your Endpoints**
   - Read all API documentation
   - Make a list of every endpoint (method, path, what it does)
   - Group related endpoints together

2. **Create Request and Response Formats**
   - For endpoints with custom formats, create data transfer objects
   - Match field names exactly to the API documentation
   - Keep these separate from your internal models

3. **Implement Handlers**
   - For each endpoint, create a handler function
   - Parse incoming requests
   - Validate the data
   - Call the appropriate service method
   - Format the response
   - Return the correct HTTP status code

4. **Set Up Routing**
   - Organize routes by resource type
   - Use a clear URL structure
   - Version your API if needed
   - Keep the health check endpoint

5. **Wire Everything Together**
   - Update the main file to connect everything
   - Initialize the database
   - Create repositories
   - Create services
   - Create controllers
   - Register routes
   - Start the server

6. **Test the Endpoints**
   - Write integration tests that call the endpoints
   - Test both success and error cases
   - Verify the responses match the documentation

### What to Deliver

Provide a report with:
1. A complete table of all endpoints (method, path, handler, service called, status codes)
2. Example commands to test 2-3 key endpoints with expected responses
3. Summary of request and response formats created
4. List of controller files and what they handle
5. The routing structure
6. How everything is connected in the main file
7. Integration test results - all should pass
8. Manual testing confirmation
9. The branch name and commit message
10. Any unclear API specifications

Then stop and wait for Step 5.

---

## Step 5: Add Supporting Features

### Your Role: Platform Engineer
You're adding the infrastructure features that support the entire application - things like authentication, logging, documentation, and error handling. Your job is to implement only what's documented, making the application more robust and production-ready.

### Which Files to Use
- **auth_rules.md** - Authentication and authorization requirements if specified
- **validation-rules.md** - Additional validation specifications if needed

**Important**: Only use the files that actually exist. If a file doesn't exist, that feature is not required.

### What You're Building
Infrastructure features that might include:
- API documentation (if specified)
- Logging system (if specified)
- Error handling (if specified)
- Authentication (if specified)
- Configuration management (if specified)
- Other cross-cutting features as documented

### Important Note
Only implement features that are explicitly documented. Don't add "nice to have" features on your own.

### Your Tasks

1. **Review Documentation**
   - Check all infrastructure-related documentation
   - Make a checklist of what's required vs what's not mentioned
   - Focus only on required items

2. **API Documentation (If Required)**
   - Set up Swagger or OpenAPI as specified
   - Document each endpoint
   - Make the documentation accessible at a URL

3. **Logging (If Required)**
   - Set up the logging system as specified
   - Use the exact format documented
   - Log at the specified locations
   - Add logging middleware if needed

4. **Error Handling (If Required)**
   - Create a standardized error response format
   - Use the exact format from documentation
   - Add error handling middleware
   - Define error codes as documented

5. **Authentication (If Required)**
   - Implement the authentication method specified (like JWT)
   - Use the exact token format and lifetime documented
   - Create authentication middleware
   - Protect the endpoints that need it
   - If roles are specified, implement role-based access control

6. **Configuration (If Required)**
   - Set up configuration using the specified method
   - Load settings from the documented source
   - Make configuration available throughout the app

7. **Apply All Middleware**
   - Wire middleware in the correct order
   - Apply to appropriate routes

### What to Deliver

Provide a report with:
1. A checklist showing what was required and what you implemented
2. If API docs were added, where to access them
3. If logging was added, example log format and what gets logged
4. If error handling was added, example error response
5. If authentication was added, how to get and use tokens, which endpoints are protected
6. If configuration was added, how it's structured and what values are needed
7. List of all middleware and their order
8. Any new tests added
9. The branch name and commit message
10. Summary of what was implemented vs not specified

This completes all 5 steps.

---

## Important Reminders

### Always Remember
- The documentation files are your only source - don't deviate from them
- Do one step at a time and wait for approval between steps
- Work on branches, never on main
- Implement exactly what's documented - no more, no less
- Test everything you build
- Write clear reports explaining what you did
- If anything is unclear, document it and ask

### When Uncertain
If documentation is unclear:
- Implement the simplest version that works
- Clearly note what was ambiguous
- Explain your choice
- Wait for clarification

### Success Checklist
Before saying you're done with a step:
- Everything compiles without errors
- Everything runs without crashing
- All tests pass
- You can demonstrate it works (with commands or tests)
- All deliverables are provided
- Branch is created, code is committed and pushed
- Pull request is created

Follow these steps carefully and you'll successfully migrate the application from COBOL to Golang.
