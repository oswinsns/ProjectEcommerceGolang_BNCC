# Testing Strategy & Coverage

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Current Test State

- **Unit Test Files (`*_test.go`):** 0 files found.
- **Integration Tests:** None automated.
- **End-to-End Tests:** None automated.
- **CI / CD Pipeline:** No GitHub Actions or CI pipeline configured (`.github/workflows/` does not exist).
- **Test Coverage:** **0% automated coverage**.

---

## 2. Existing Verification Methods

While no automated tests are checked into the codebase, the project relies on the following manual testing approaches:

1. **Postman API Collection:**
   - Mentioned in [`README.md:28`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/README.md#L28) with public link: `https://www.postman.com/aerospace-geoscientist-45054066/projehk-goecommerce`.
   - Used for manual verification of authentication, product CRUD, and user CRUD endpoints.
2. **Manual Browser Smoke Testing:**
   - Navigating `/admin/login`, `/admin/dashboard`, `/admin/users`, `/admin/products`, and `/products`.
3. **Database Seeding Verification:**
   - Automatic execution of `seeders.SeedProducts()` and `seeders.SeedUsers()` resets tables on every launch, ensuring a known initial data state for testing.

---

## 3. Recommended Testing Roadmap

To establish a solid quality baseline, the following testing infrastructure is recommended:

### A. Unit Testing
- Add standard Go unit tests (`*_test.go`) utilizing `net/http/httptest` and `github.com/stretchr/testify`.
- Test token generation and validation in `middlewares/jwt.go`.
- Test password hashing logic with bcrypt.

### B. Handler / Controller Tests
- Use `httptest.NewRecorder()` and Gin's `gin.CreateTestContext()` to test endpoints without spinning up a live HTTP port.
- Mock database operations using `go-sqlmock` or an in-memory SQLite database (`gorm.io/driver/sqlite`).

### C. Continuous Integration
- Add `.github/workflows/go-test.yml` running `go vet ./...` and `go test -v ./...`.
