# Integrations & External Services

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Database Integration (MySQL)

- **Engine:** MySQL Server (typically local MySQL instance or XAMPP / MariaDB on port `3306`)
- **Connection Configuration:**
  - Configured in [`Day10/configs/database.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/configs/database.go#L13-L29)
  - Connection DSN: `user:password@tcp(host:port)/dbname?parseTime=true`
  - Global connection instance: `configs.DB *gorm.DB`
- **Schema Management:**
  - Automated tables creation via `configs.DB.AutoMigrate(&models.Product{}, &models.User{})` in [`Day10/databases/automigrate.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/automigrate.go#L9-L18).
- **Tables Created:**
  - `products`: Managed by `models.Product` (contains `id`, `name`, `stock`, `price`, `created_at`, `updated_at`, `deleted_at`)
  - `users`: Managed by `models.User` (contains `id`, `username`, `password`, `is_active`, `email`, `created_at`, `updated_at`, `deleted_at`)
- **Lifecycle Integration:**
  - Seeders run on every launch via `seeders.SeedProducts()` and `seeders.SeedUsers()` in [`Day10/main.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/main.go#L42-L43).
  - Note: Current seeders truncate tables upon startup.

---

## 2. Authentication & Authorization Integration

- **Standard:** JSON Web Tokens (JWT) according to RFC 7519
- **Implementation:** [`Day10/middlewares/jwt.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/middlewares/jwt.go)
- **Token Delivery Mechanisms:**
  1. **HTTP Cookie:** Read from `token` cookie (set as `HttpOnly`, 24 hours expiry in [`handlers/auth-handler.go:36`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/auth-handler.go#L36)).
  2. **Authorization Header:** Read from `Authorization: Bearer <token>`.
- **Protected Group:**
  - Router group `/admin/*` protected by `middlewares.AuthMiddleware()`.

---

## 3. Spreadsheet / File Export Integration

- **Library:** `excelize/v2`
- **Endpoints:**
  - `GET /products/export` (Public export in [`handlers/public-handler.go:36-76`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/public-handler.go#L36-L76))
  - `GET /admin/products/export` (Admin export in [`routes/route.go:34`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/routes/route.go#L34))
- **MIME Type:** `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- **Output Filename:** `products.xlsx`
- **Export Columns:** `ID`, `Name`, `Price`, `Stock`, `Created At`

---

## 4. External Services & APIs

- **Postman API Collection:**
  - Documented in [`README.md:28`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/README.md#L28):
  - URL: `https://www.postman.com/aerospace-geoscientist-45054066/projehk-goecommerce`
- **External CDNs:**
  - Bootstrap 4 CSS: `https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css` used in HTML views.

---

## 5. Webhooks & Message Queues

- **Current Status:** None configured. All operations are synchronous request-response.
