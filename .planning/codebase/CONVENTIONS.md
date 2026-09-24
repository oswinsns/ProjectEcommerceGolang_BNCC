# Code Conventions & Standards

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Naming & Case Conventions

- **Packages:**
  - Lowercase, single-word names: `configs`, `databases`, `seeders`, `handlers`, `middlewares`, `models`, `routes`.
- **Files:**
  - Mixed naming convention: kebab-case with descriptive prefixes/suffixes (e.g. `product-handler.go`, `dashboardadmin-handler.go`, `seeder-products.go`, `auth-handler.go`).
- **Structs & Types:**
  - Standard Go PascalCase: `Product`, `User`.
  - JSON and Form tags use snake_case: `json:"is_active" form:"is_active"`.
- **Functions:**
  - Exported functions are PascalCase: `GetProducts`, `CreateUser`, `SetupMySQL`.
  - Private helpers are camelCase: `hashPassword`.
- **Variables & Fields:**
  - PascalCase for exported struct fields and global variables (`DB`).
  - camelCase for local variables (`tokenString`, `authHeader`, `productCount`).

---

## 2. Architectural Patterns & Data Access

- **Global ORM Instance:**
  - Database access relies on a shared global variable `configs.DB *gorm.DB` rather than dependency injection.
- **ORM Patterns:**
  - Uses GORM built-in querying methods: `.Find()`, `.Where()`, `.Order()`, `.Limit()`, `.Count()`, `.Save()`, `.Delete()`.
  - Embedded `gorm.Model` provides `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` (enabling soft deletes).
- **Binding & Parsing:**
  - Handlers parse form inputs using `c.PostForm("name")` or `c.ShouldBind(&form)`.
  - REST API handlers parse JSON payloads using `c.ShouldBindJSON(&input)`.
  - String-to-number conversions explicitly use `strconv.Atoi()` and `strconv.ParseFloat()`.

---

## 3. Error Handling & Logging

- **Panic in Bootstrapping:**
  - Fatal bootstrapping failures in `.env` loading, database connections, and migrations trigger `panic()` (caught by a deferred `recover()` block in [`main.go:24-28`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/main.go#L24-L28)).
- **HTTP Error Responses:**
  - Web UI routes return error strings: `c.String(http.StatusInternalServerError, "Error fetching products")` or `c.String(http.StatusNotFound, "Product not found")`.
  - REST API routes return JSON dictionaries: `c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})`.
- **Logging:**
  - Standard Go `log` package used in seeders (`log.Println`, `log.Fatalf`).
  - Gin built-in logger middleware handles HTTP request logging.

---

## 4. Routing Conventions

- **Hybrid Routing Schema:**
  - Web page forms submit via standard HTTP `POST` methods (e.g. `/add-products`, `/users/update/:id`, `/admin/users/delete/:id`).
  - REST APIs follow standard REST verbs with prefix `/api/*` (e.g. `PUT /api/users/:id`, `GET /api/users`, `DELETE /api/users/delete/:id`).
  - Notice slight inconsistency: Some delete routes use POST, some DELETE, and some contain `/delete/` inside the path.

---

## 5. Security & Authentication Conventions

- **JWT Tokens:**
  - Standard claims via `jwt.RegisteredClaims`.
  - Issued with 24-hour expiration (`ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour))`).
  - Transmitted via HTTP-only cookie for browser navigation, and Bearer token for API tools.
- **Password Security:**
  - Bcrypt hashing using `bcrypt.DefaultCost`.
