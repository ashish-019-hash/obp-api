# OBP-API Golang Migration — Stepwise Implementation Guide

This document defines the **role-based, step-by-step** process for migrating the OBP-API backend into Go.  

> **Key Principle:**  
> All reference `.md` files in `04.phase-2-input` are the **single source of truth**. Follow specifications exactly without adding or skipping any fields, screens, or logic.

> **Critical Instruction:**  
> You will be provided with **ONE STEP AT A TIME**. Implement **ONLY** the current step that is provided to you.  
> - Do **not** proceed to subsequent steps.  
> - After completing the current step, **STOP** and wait for the next step to be explicitly provided.  
> - Do **not** assume or implement future steps.

---

## Shared Context

- **Repo:** `OBP-API`  
- **Branch:** `main`  
- **Input folder:** `04.phase-2-input`

### Reference Files (Single Source of Truth)

- **Business Rules:** `business_rules.md`  
- **User Requirements:** `user_stories.md`  
- **Data Structure:** `business_entities.md`  
- **Field Rules:** `validation-rules.md`  

Always read these files for domain, API, and business logic. Do not invent behavior beyond what is described.

---

## Step 1 — Project Skeleton Initialization

**Role:** Golang Project Bootstrapper (Senior Backend Engineer)

You are responsible for creating the **project skeleton only**.  
Use the `.md` files in `04.phase-2-input` as your only source of truth for *later* steps, but **do not implement any business logic now**.

### Tasks

1. **Initialize Go module**
   - Initialize a Go module under `Migrated_Code/backend`.
   - Run `go mod init` with an appropriate module path.
   - If using Gin, add it with:
     - `go get github.com/gin-gonic/gin`.

2. **Create folder structure (exactly as listed)**

   - `cmd/main.go` — application entry point  
   - `internal/config/`  
   - `internal/controllers/`  
   - `internal/middleware/`  
   - `internal/models/`  
   - `internal/repositories/`  
   - `internal/routes/`  
   - `internal/services/`  
   - `internal/utils/`  
   - `pkg/db/` — DB helpers  
   - `test/` — tests  
   - `go.mod` and `go.sum`

3. **Bootstrap HTTP server**

   - Use either:
     - Standard `net/http`, **or**
     - `github.com/gin-gonic/gin`.
   - The server **must compile and start**, but must **not** include any domain logic.
   - Keep everything minimal and clean:
     - No DB queries
     - No service implementations
     - No controllers with logic
   - Implement only a minimal `"hello"` or `/health` endpoint to confirm the server starts.

### Deliverables (after Step 1)

- A short report listing **created files and folders**.
- **Commands** to build and run the server  
  - Example: `go run ./cmd`.
- A one-line example **curl** command to verify the server  
  - Example: `curl http://localhost:8080/health`  
  - Include the expected minimal response.
- Suggested **branch name** and **commit message**, e.g.:
  - Branch: `devin/step-1-skeleton`  
  - Commit: `chore: init go module + project skeleton`

> After completing Step 1, **wait for input before proceeding to Step 2.**

---

## Step 2 — Translate .md Entities to Go Structs

**Role:** Golang Domain Modeler

Read the entity definitions **only** from the `.md` files inside `04.phase-2-input` (for example `business_entities.md`, `business_entities_analysis.md`). Implement **exact 1:1 translations** to Go structs.

### Tasks

1. **Create model files**

   - For **every entity** in the `.md` files, create one `.go` file under `internal/models/` named after the entity.  
     - Example: `user.go`, `order.go`.

2. **Map fields exactly**

   - Preserve **field order** from the documentation.
   - Use **PascalCase** Go field names.
   - Choose Go types that match the described types:  
     - `string`, `int`, `float64`, `bool`, `time.Time`, etc.
   - If fields are **optional** in the docs, add `omitempty` in the `json` tag.
   - Use `json:"exactApiName"` tags that match the **exact API field names** in the `.md` files.
   - Convert constraints and metadata from the `.md` into struct tags **where explicitly specified**, e.g.:
     - `validate:"required,max=100"`  
     - or other custom tags.

3. **Defaults & constructors**

   - If default values or initialization rules are provided in the `.md`, implement a **constructor function**:
     - `NewEntityName(...)`  
     - Set defaults exactly as documented.

4. **Documentation comments**

   - Add Go comments above structs and fields to reflect descriptions from the `.md`.

5. **No extra behavior**

   - Do **not** modify names.  
   - Do **not** add fields.  
   - Do **not** infer behavior that is not in the docs.  
   - Keep each file focused: **one entity per file**.

### Deliverables (after Step 2)

- A **list of structs** implemented and their corresponding **file names**.
- An **example snippet** of one generated struct (just one small sample to confirm style).
- A **unit test file skeleton** under `test/` that:
  - Imports the models.
  - Confirms they compile (no business logic tests yet).
- Suggested commit message, e.g.:
  - `feat(models): add entity structs from .md files`

> After completing Step 2, **wait for input before proceeding to Step 3.**

---

## Step 3 — Implement Services and Repository Layer

**Role:** Golang Service & Data Architect

Use only the `.md` files that describe **business rules, workflows, and queries** (for example `COBOL_Business_Rules_Analysis.md`, `story_points.md`, `entity_queries.md`). Implement all **business logic** and **data access** exactly as described.

### Tasks — Services

- Implement service-layer code in `internal/services/`.
- Each service method must follow:
  - Validations
  - Data transformations
  - Workflows and orchestration rules
- Services may call other services or repositories as documented.
- When the `.md` lists step-by-step workflows (status transitions, chained steps), implement them directly inside service functions.

### Tasks — Repositories

- Implement repository **interfaces** and **implementations** in `internal/repositories/`.
- Use a **single shared SQLite in-memory connection** provided by `pkg/db/`.
  - Create a connection helper there.
- Implement all CRUD and custom queries exactly as described in the `.md`.
- Repositories must **not contain business logic** — only data access.
- Use **interface-based repositories** so they can be mocked in tests.

### Database

- Configure **SQLite in-memory** mode (`:memory:`) in `pkg/db/`.
- Provide a helper to access it across repositories.
- Implement migration helpers **only if** the `.md` requires them.

### Testing

- Add unit tests for:
  - Service boundaries
  - Repository boundaries  
- Use the **in-memory DB**.
- Mock interfaces where appropriate.

### Deliverables (after Step 3)

- Short summary of **implemented services and repositories**:
  - File names
  - Brief responsibilities
- Instructions on **how to run tests** and expected outcomes.
- Example of:
  - One **service method signature**  
  - Its corresponding **repository interface** method.
- Suggested commit message, e.g.:
  - `feat(service/repo): implement business logic & repositories from .md`

> After completing Step 3, **wait for input before proceeding to Step 4.**

---

## Step 4 — Implement API Endpoints

**Role:** Golang API Engineer (Handlers & Routing)

Implement REST endpoints exactly according to the **API contract** `.md` files (for example `api_contracts.md`, `workflow_endpoints.md`, `story_points.md`). Do **not** deviate from documented URLs, methods, fields, or status codes.

### Tasks

1. **Controllers / Handlers**

   - For each documented endpoint:
     - Create a controller/handler in `internal/controllers/`.
   - Implement handlers using:
     - Gin **or** `net/http` (whichever was chosen in Step 1).
   - For each handler:
     - Validate and parse requests exactly as the `.md` requires.
     - Call the appropriate **service layer** methods from Step 3.
     - Return responses with:
       - Matching JSON structure
       - Correct status codes

2. **Routing**

   - Add route registration in `internal/routes/`, organized by version and resource.  
     - Example: `/api/v1/users`
   - Match the routes defined in the `.md`.
   - Wire routes into `cmd/main.go`.
   - Start the server on the documented port (or default from `.md` if provided).

3. **Integration Tests**

   - Add minimal integration tests that exercise endpoints **end-to-end** using the in-memory DB.

### Deliverables (after Step 4)

- A mapping table of implemented endpoints:

  | HTTP METHOD | PATH            | Controller         | Service Called      | Expected Status Codes |
  |------------|-----------------|--------------------|---------------------|-----------------------|

- Example `curl` commands for **two representative endpoints** and the expected response shape.
- List of tests added and **how to run them**.
- Suggested commit message, e.g.:
  - `feat(api): implement endpoints from api_contracts.md`

> After completing Step 4, **wait for input before proceeding to Step 5.**

---

## Step 5 — Cross-Cutting Concerns

**Role:** Golang Platform Engineer (Cross-Cutting Concerns)

Only implement cross-cutting features that are **explicitly documented** in the `.md` files (for example `api_documentation.md`, `logging_spec.md`, `error_handling.md`, `middleware_specs.md`, `auth_rules.md`). Do **not** add extra features beyond what the docs call for.

### Tasks (Apply Only If Present in .md)

- **API Documentation**
  - If required, integrate OpenAPI/Swagger:
    - Annotate handlers as described.
    - Integrate `swaggo` + `gin-swagger` (or as specified).

- **Logging**
  - If the `.md` specifies:
    - Logging format
    - Components to log
  - Use the exact library and format (e.g., `log`, `zap`, `zerolog`).
  - Add logs only where the docs specify.

- **Error Handling**
  - Implement centralized error format/middleware if described.
  - Match:
    - Response shape
    - Keys
    - Fields
    - Status codes

- **Middleware**
  - Implement only documented middleware:
    - Request logging
    - Auth
    - Header validation
    - etc.
  - Wire into routes exactly as described.

- **Auth**
  - If JWT or RBAC rules are defined:
    - Implement token lifetimes, claims, roles exactly as documented.

- **Validation**
  - Use `go-playground/validator` only where specified.
  - Match validation rules precisely.

- **Utilities & Mappers**
  - Implement helper functions listed in `.md`:
    - Date formats
    - DTO ↔ Entity mappers
  - Place them under:
    - `internal/utils/` or
    - `internal/mappers/`.

- **Configuration**
  - Store configuration as specified:
    - `.env`, `.yaml`, etc.
  - Use tools like `viper` or `godotenv` only if the `.md` requires them.

- **Data Mappers / Converters**
  - Implement DTO ↔ entity mapping logic exactly as described.
  - Place in `internal/mappers/` or as helper methods.

### Deliverables (after Step 5)

- A **checklist** of which cross-cutting items existed in the `.md` files and which were implemented.
- If Swagger/OpenAPI was added, include the **path** where the UI is served.
- Example:
  - One sample **log line** showing format.
  - One example **error response body**.
- Suggested commit message, e.g.:
  - `chore(platform): add error handling + logging per docs`

> After finishing Step 5, **wait for further instructions** before making any additional changes.

---

## Generic Instructions (Apply to Every Step)

- Always work from files in `04.phase-2-input` **only** for domain, API, and business rules.
- Do **not** add features not present in the `.md` files — **no guessing, no assumptions**.
- Keep each commit focused on the current step (**one logical change per commit**).
- Use clear branch naming:
  - `devin/step-<n>-<short-desc>`
- Use commit messages starting with:
  - `feat:` for new features.
  - `chore:` for setup and non-feature changes.
- After finishing a step:
  - Provide the **deliverables** listed for that step.
  - **Wait for go-ahead** before continuing.
- If anything in the `.md` is **ambiguous**:
  - Implement the **minimal behavior** needed to compile and run.
  - Document the ambiguity in your step report.
  - Wait for instructions before changing behavior.

