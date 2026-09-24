# Directory & Codebase Structure

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Directory Tree Layout

```
ProjectEcommerceGolang_BNCC/
├── .git/                                   # Git repository metadata
├── README.md                               # Project documentation & Postman link
└── Day10/                                  # Main Go application package / module root
    ├── .env                                # Environment variables (MySQL credentials)
    ├── go.mod                              # Go module definition (Day10) and dependencies
    ├── go.sum                              # Dependency checksums
    ├── main.go                             # Application entry point & server setup
    │
    ├── configs/                            # Configuration packages
    │   └── database.go                     # MySQL connection setup & global DB instance
    │
    ├── databases/                          # Database operations
    │   ├── automigrate.go                  # GORM AutoMigrate schema definition
    │   └── seeders/                        # Database seeders
    │       ├── seeder-products.go          # Product table truncation and seed rows
    │       └── seeder-user.go              # User table truncation and seed rows
    │
    ├── handlers/                           # HTTP controllers / route handlers
    │   ├── auth-handler.go                 # Login, authentication & token generation
    │   ├── dashboardadmin-handler.go       # Admin dashboard aggregations & stats
    │   ├── home-handler.go                 # Home / welcome template handlers
    │   ├── product-handler.go              # Product CRUD (web forms + JSON API)
    │   ├── public-handler.go               # Public catalog display & Excel export
    │   └── user-handler.go                 # User CRUD (web forms + JSON API)
    │
    ├── middlewares/                        # Gin HTTP middleware
    │   └── jwt.go                          # JWT extraction, verification & token minting
    │
    ├── models/                             # GORM domain models / entities
    │   ├── product.go                      # Product entity struct definition
    │   └── user.go                         # User entity struct definition
    │
    ├── routes/                             # Route configurations
    │   └── route.go                        # Route groups, endpoints & handler wiring
    │
    ├── static/                             # Static public assets
    │   ├── login.css                       # Login page custom styles
    │   └── *.jpg, *.jpeg, *.png, *.webp    # Static sample product images & screenshots
    │
    └── views/                              # HTML templates (server-side rendering)
        ├── create.html                     # Admin create user form
        ├── createproducts.html             # Admin create product form
        ├── editproduct.html                # Admin edit product form
        ├── edituser.html                   # Admin edit user form
        ├── home.html                       # Admin dashboard overview & summary cards
        ├── login.html                      # Authentication login view
        ├── products.html                   # Admin product management table
        ├── public.html                     # Public store catalog view
        └── users.html                      # Admin user management table
```

---

## 2. Component Responsibility Matrix

| Path | Primary Responsibility | Key Exported Symbols |
| :--- | :--- | :--- |
| `configs/database.go` | Database connection lifecycle | `DB`, `SetupMySQL()` |
| `databases/automigrate.go` | Database schema initialization | `AutoMigrate()` |
| `databases/seeders/` | Database fixture seeding | `SeedProducts()`, `SeedUsers()` |
| `handlers/auth-handler.go` | User authentication & sessions | `Login()`, `ShowLogin()` |
| `handlers/dashboardadmin-handler.go` | Dashboard metrics calculation | `AdminDashboard()` |
| `handlers/product-handler.go` | Product business logic | `GetProducts()`, `CreateProducts()`, `UpdateProduct()`, `DeleteProduct()`, `GetProductsAPI()` |
| `handlers/user-handler.go` | User management logic | `GetUsers()`, `CreateUser()`, `UpdateUser()`, `DeleteUser()`, `GetUsersAPI()`, `DeleteUserAPI()`, `UpdateUserAPI()` |
| `handlers/public-handler.go` | Public views & reports | `ShowProductsPage()`, `ExportProducts()` |
| `middlewares/jwt.go` | Auth gatekeeper | `AuthMiddleware()`, `GenerateToken()` |
| `models/product.go` | Product data model | `Product` |
| `models/user.go` | User data model | `User` |
| `routes/route.go` | URL routing & grouping | `SetupRoutes()` |

---

## 3. Entry Point Flow

When starting the application via `go run main.go`:
1. `main()` registers a deferred `recover()` block.
2. `godotenv.Load()` reads `.env` variables.
3. `configs.SetupMySQL()` opens the GORM MySQL pool.
4. `databases.AutoMigrate()` ensures DB tables exist.
5. `seeders.SeedProducts()` and `seeders.SeedUsers()` truncate and re-populate records.
6. `gin.Default()` initializes the router with logger & recovery middlewares.
7. Templates (`views/*`) and static files (`static/`) are mapped.
8. `routes.SetupRoutes(r)` maps all endpoints.
9. `r.Run()` binds and starts the server on port `:8080`.
