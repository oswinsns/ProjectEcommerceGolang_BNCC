# System Architecture

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Architectural Overview

The application follows a **Monolithic Layered MVC (Model-View-Controller) / Hybrid REST API** architecture implemented with Go and the Gin web framework.

It simultaneously serves:
1. **Server-Side Rendered (SSR) HTML Pages:** Admin dashboard, login views, and management forms using Go's `html/template` engine.
2. **Public / REST API Endpoints:** JSON-based data endpoints under `/api/*` and public endpoints for integration with external clients (e.g. Postman, mobile clients, SPAs).

```mermaid
graph TD
    Client[Web Browser / Postman] --> Router[Gin Engine (routes/route.go)]
    
    subgraph Routing & Middleware
        Router --> PublicRoutes[Public Routes (/products, /login)]
        Router --> APIRoutes[API Routes (/api/*)]
        Router --> AdminGroup[Admin Group (/admin/*)]
        AdminGroup --> AuthMW[middlewares.AuthMiddleware()]
    end

    subgraph Handlers Layer
        PublicRoutes --> PublicH[handlers.public-handler / product-handler]
        APIRoutes --> APIH[handlers.user-handler / product-handler]
        AuthMW --> AdminH[handlers.dashboardadmin-handler / product-handler / user-handler]
    end

    subgraph Data & Storage Layer
        PublicH --> Models[models.Product, models.User]
        APIH --> Models
        AdminH --> Models
        Models --> GORM[GORM ORM (configs.DB)]
        GORM --> MySQL[(MySQL Database)]
    end

    subgraph Presentation
        PublicH --> PublicViews[views/public.html]
        AdminH --> AdminViews[views/home.html, views/products.html, etc.]
        PublicH --> Excel[Excelize XLSX Export]
    end
```

---

## 2. Key Components & Layers

### A. Presentation & Routing Layer (`Day10/routes/`, `Day10/views/`)
- **Route Registrar:** [`routes/route.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/routes/route.go) sets up route mappings to handlers.
- **HTML Views:** Template files inside [`views/`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/views/) loaded using `r.LoadHTMLGlob("views/*")`.
- **Static Assets:** Served via `/static` mapping to [`Day10/static/`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/static/).

### B. Middleware Layer (`Day10/middlewares/`)
- **JWT Authentication:** [`middlewares/jwt.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/middlewares/jwt.go).
- Verifies HMAC token from Authorization Bearer header or HTTP-only `token` cookie.
- Attaches the parsed claims subject (`username`) to Gin context.

### C. Handler / Controller Layer (`Day10/handlers/`)
- Encapsulates HTTP request decoding, business validation, data retrieval/mutation, and response generation (HTML render or JSON payload).
- Handler modules:
  - `auth-handler.go`: User login, cookie setting, token generation.
  - `dashboardadmin-handler.go`: Aggregates statistics (user counts, active users, product counts) for the admin dashboard.
  - `product-handler.go`: CRUD operations on products for both admin and JSON APIs.
  - `user-handler.go`: CRUD operations on users for both admin and JSON APIs.
  - `public-handler.go`: Public catalog display and Excel export.
  - `home-handler.go`: Standalone / alternative welcome handlers.

### D. Model & Persistence Layer (`Day10/models/`, `Day10/configs/`, `Day10/databases/`)
- **Database Connection:** [`configs/database.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/configs/database.go) creates a global singleton `configs.DB`.
- **Data Models:** [`models/product.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/models/product.go) and [`models/user.go`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/models/user.go) embed `gorm.Model` for soft delete and timestamp management.
- **Migration & Seeders:** Automated schema synchronization and initial table seeding on startup.

---

## 3. Data Flow & Request Lifecycle

1. **HTTP Request Arrival:** Incoming request is matched against Gin routes.
2. **Authentication Gate:** If matching `/admin/*`, `AuthMiddleware` inspects token in cookie or header. If missing/invalid, aborts with `401 Unauthorized`.
3. **Handler Processing:**
   - Parses path parameters (`c.Param("id")`), query parameters, or form data (`c.PostForm` / `c.ShouldBind`).
   - Executes database queries via `configs.DB`.
4. **Response Delivery:**
   - For web views: `c.HTML(http.StatusOK, "view.html", gin.H{...})`
   - For REST APIs: `c.JSON(http.StatusOK, payload)`
   - For file downloads: streams binary bytes directly to `c.Writer` via `excelize`.

---

## 4. State Management

- **Database State:** Persisted in MySQL tables `products` and `users`.
- **User Authentication State:** Stateless JWT token. Stored client-side in an HTTP-only cookie `token` or passed as a Bearer token.
- **In-Memory Global State:** `configs.DB` global pointer and `middlewares.jwtKey` byte slice.
