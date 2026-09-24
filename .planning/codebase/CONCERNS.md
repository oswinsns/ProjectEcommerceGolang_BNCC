# Codebase Concerns & Technical Debt

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Critical Functional Bugs

### 🚨 1. `DeleteProduct` Deletes Users Instead of Products
- **Location:** [`handlers/product-handler.go:150-160`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/product-handler.go#L150-L160)
- **Code:**
  ```go
  func DeleteProduct(c *gin.Context) {
      id := c.Param("id")

      if err := configs.DB.Delete(&models.User{}, id).Error; err != nil {
          c.String(http.StatusInternalServerError, "Failed to delete user")
          return
      }

      // After delete, go back to users list
      c.Redirect(http.StatusFound, "/admin/products")
  }
  ```
- **Severity:** **CRITICAL**.
- **Impact:** An administrator deleting a product actually deletes the User record sharing that numeric ID! The product remains untouched in the database, while user data is deleted or soft-deleted.

### 🚨 2. Plaintext Password Saved on User Creation
- **Location:** [`handlers/user-handler.go:39-64`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/user-handler.go#L39-L64)
- **Code:**
  ```go
  user := models.User{
      Username: username,
      Password: password, // ⚠️ should hash before saving!
      Email:    email,
      IsActive: isActive,
  }
  ```
- **Severity:** **HIGH**.
- **Impact:** While `seeder-user.go` hashes passwords with bcrypt, the admin `CreateUser` handler saves the raw password string directly. Any user created via this form will never be able to log in (because `bcrypt.CompareHashAndPassword` in `auth-handler.go` expects a valid bcrypt hash) and passwords are saved unencrypted.

### 🚨 3. Automatic Database Truncation on Every Server Start
- **Location:** [`main.go:42-43`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/main.go#L42-L43), [`databases/seeders/seeder-products.go:19`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/seeders/seeder-products.go#L19), [`databases/seeders/seeder-user.go:25`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/seeders/seeder-user.go#L25)
- **Code:**
  ```go
  configs.DB.Exec("TRUNCATE TABLE products")
  configs.DB.Exec("TRUNCATE TABLE users")
  ```
- **Severity:** **HIGH**.
- **Impact:** Any new products or users created by administrators are immediately wiped out the next time the Go server restarts. Seeding should check if records exist before inserting rather than executing `TRUNCATE TABLE`.

---

## 2. Security & Authentication Concerns

### 🔒 1. Hardcoded Secret Key in Source Code
- **Location:** [`middlewares/jwt.go:12`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/middlewares/jwt.go#L12)
- **Code:** `var jwtKey = []byte("rahasia")`
- **Impact:** Anyone with read access to the repository can forge arbitrary administrator JWT tokens. The secret key must be extracted to `.env` (e.g. `JWT_SECRET`).

### 🔒 2. Committed `.env` File & Missing `.gitignore`
- **Location:** [`Day10/.env`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/.env)
- **Impact:** The `.env` file containing database credentials is fully tracked in Git. There is no `.gitignore` file in either the repository root or the `Day10/` folder, risking accidental leakage of production secrets and binary files.

### 🔒 3. Lack of CSRF Protection on Admin Forms
- **Location:** All admin forms in [`Day10/views/`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/views/)
- **Impact:** Because the authentication token is stored in a browser cookie (`token`) with `SameSite` not strictly configured, authenticated admin sessions may be susceptible to Cross-Site Request Forgery attacks.

---

## 3. Structural & Maintenance Debt

### ⚠️ 1. Inconsistent Endpoint Naming & HTTP Methods
- **Location:** [`routes/route.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/routes/route.go)
- **Issues:**
  - Route duplication: `/admin/login` and `/login` are both registered to the same handlers.
  - Inconsistent delete methods: `POST /users/delete/:id` vs `DELETE /products/delete/:id` vs `DELETE /api/users/delete/:id`.
  - Non-RESTful paths: `/add-users`, `/add-products` instead of RESTful `POST /admin/users` and `POST /admin/products`.

### ⚠️ 2. Cluttered Codebase with Extensive Commented-out Code
- Multiple handlers (`product-handler.go`, `home-handler.go`, `user-handler.go`, `seeder-user.go`) contain large blocks of dead / commented-out code, obscuring active logic.

### ⚠️ 3. Subfolder Root Architecture
- `go.mod` is located inside `ProjectEcommerceGolang_BNCC/Day10/` rather than the root of the repository. Running Go commands from the repository root fails unless the developer navigates into `Day10/`.
