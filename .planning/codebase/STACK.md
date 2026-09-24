# Tech Stack

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Core Languages & Runtime

- **Language:** Go (Golang)
- **Version:** `go 1.24.6` (specified in [`Day10/go.mod`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/go.mod#L3))
- **Module Name:** `Day10`
- **Architecture / Target:** Cross-platform (tested on Windows x64)

---

## 2. Web Framework & HTTP Layer

- **Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin) (`v1.10.1`)
  - Lightweight, high-performance HTTP web framework with middleware pipeline
  - Used for routing, JSON serialization/deserialization, HTML template rendering, and cookie management
- **Template Engine:** Go standard library `html/template` integrated via Gin's `LoadHTMLGlob("views/*")` and custom `SetFuncMap` ([`main.go:51-76`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/main.go#L51-L76))

---

## 3. Database & ORM

- **Database:** MySQL (Relational Database)
- **ORM:** [GORM](https://gorm.io/) (`v1.30.1`)
  - Driver: `gorm.io/driver/mysql` (`v1.6.0`)
  - Underlying Driver: `github.com/go-sql-driver/mysql` (`v1.9.3`)
- **Schema Management:** Auto-migration via `gorm.DB.AutoMigrate()` for `models.Product` and `models.User` ([`databases/automigrate.go:10-14`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/automigrate.go#L10-L14))

---

## 4. Authentication & Security

- **JWT Tokens:** `github.com/golang-jwt/jwt/v5` (`v5.3.0`)
  - HMAC-SHA256 (`HS256`) symmetric token generation and parsing
  - Stored in HTTP cookies (`token`) and accepted via `Authorization: Bearer <token>`
- **Password Hashing:** `golang.org/x/crypto/bcrypt` (`v0.41.0`)
  - Used in auth handlers and user seeders for secure hashing and password comparison

---

## 5. File Processing & Utilities

- **Excel Processing:** `github.com/xuri/excelize/v2` (`v2.9.1`)
  - Dynamic generation and streaming of `.xlsx` spreadsheets for product exports
- **Environment Management:** `github.com/joho/godotenv` (`v1.5.1`)
  - Loads key-value configuration from `.env` on startup
- **Validation:** `github.com/go-playground/validator/v10` (`v10.27.0`)
  - Model tag validation for struct fields (`form:"stock,min=0"`)

---

## 6. Frontend & Static Assets

- **CSS Framework:** Twitter Bootstrap 4.0.0 (loaded via CDN in [`views/home.html`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/views/home.html#L7))
- **Custom Styling:** [`Day10/static/login.css`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/static/login.css)
- **Static Assets:** Static files served via Gin router at `/static` pointing to `./static` directory

---

## 7. Configuration & Environment Variables

Loaded from [`Day10/.env`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/.env):

| Variable | Description | Example / Default |
| :--- | :--- | :--- |
| `DB_USER` | MySQL database username | `root` |
| `DB_PASSWORD` | MySQL database password | `""` |
| `DB_HOST` | MySQL server hostname/address | `localhost` |
| `DB_PORT` | MySQL server port | `3306` |
| `DB_NAME` | Database schema name | `ecommerce_db` |

---

## 8. Build & Execution Commands

- **Run Dev Server:**
  ```powershell
  cd ProjectEcommerceGolang_BNCC/Day10
  go run main.go
  ```
- **Tidy Dependencies:**
  ```powershell
  go mod tidy
  ```
- **Compile Binary:**
  ```powershell
  go build -o app.exe main.go
  ```
