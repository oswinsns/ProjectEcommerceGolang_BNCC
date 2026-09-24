# Codebase Concerns & Technical Debt

**Repository Root:** `ProjectEcommerceGolang_BNCC/`  
**Application Root:** `ProjectEcommerceGolang_BNCC/Day10/`  
**Mapped Date:** September 2026

---

## 1. Critical Functional Bugs

### ✅ 1. [FIXED] `DeleteProduct` Deletes Users Instead of Products
- **Location:** [`handlers/product-handler.go:150-160`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/product-handler.go#L150-L160)
- **Status:** **TERATASI (RESOLVED)**

#### 🔍 Kondisi Awal (Sebelum Fix):
- **Severity:** **CRITICAL**
- **Impact:** Administrator yang menghapus produk justru menghapus record User yang memiliki ID angka yang sama! Produk tetap berada di database, sedangkan data akun pengguna hilang/soft-deleted.
- **Kode Asli (Bermasalah):**
  ```go
  func DeleteProduct(c *gin.Context) {
      id := c.Param("id")

      // ❌ BUG: Memanggil Delete pada struct models.User{} bukan models.Product{}
      if err := configs.DB.Delete(&models.User{}, id).Error; err != nil {
          c.String(http.StatusInternalServerError, "Failed to delete user")
          return
      }

      // After delete, go back to users list
      c.Redirect(http.StatusFound, "/admin/products")
  }
  ```

#### 🛠️ Solusi & Perbaikan yang Diterapkan:
1. Mengubah target delete dari `&models.User{}` menjadi `&models.Product{}`.
2. Memperbaiki pesan error dari `"Failed to delete user"` menjadi `"Failed to delete product"`.
3. Memperbaiki `routes/route.go` agar mendukung method `POST` (`admin.POST("/products/delete/:id", ...)`) selain `DELETE`, karena form HTML di browser hanya mengirimkan `POST`.
4. Memperbaiki label dan konfirmasi tombol di [`views/products.html`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/views/products.html) dari "delete this user" menjadi "delete this product".
- **Kode Sesudah Fix:**
  ```go
  func DeleteProduct(c *gin.Context) {
      id := c.Param("id")

      // ✅ FIX: Menggunakan models.Product{}
      if err := configs.DB.Delete(&models.Product{}, id).Error; err != nil {
          c.String(http.StatusInternalServerError, "Failed to delete product")
          return
      }

      // After delete, go back to products list
      c.Redirect(http.StatusFound, "/admin/products")
  }
  ```

### ✅ 2. [FIXED] Plaintext Password Saved on User Creation
- **Location:** [`handlers/user-handler.go:39-64`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/handlers/user-handler.go#L39-L64)
- **Status:** **TERATASI (RESOLVED)**

#### 🔍 Kondisi Awal (Sebelum Fix):
- **Severity:** **HIGH**
- **Impact:** Meskipun `seeder-user.go` meng-hash password menggunakan Bcrypt, fungsi `CreateUser` pada handler form web menyimpan string password mentah langsung ke database. Akibatnya, user yang dibuat via web admin tidak akan pernah bisa login (karena fungsi `Login` di `auth-handler.go` memverifikasi menggunakan `bcrypt.CompareHashAndPassword` yang membutuhkan hash Bcrypt valid) dan password tersimpan tanpa enkripsi.
- **Kode Asli (Bermasalah):**
  ```go
  func CreateUser(c *gin.Context) {
      username := c.PostForm("username")
      password := c.PostForm("password")
      email := c.PostForm("email")
      isActiveStr := c.PostForm("is_active")

      isActive, _ := strconv.ParseBool(isActiveStr)

      // ❌ BUG: Menyimpan password mentah (plaintext) tanpa hash Bcrypt
      user := models.User{
          Username: username,
          Password: password, // ⚠️ should hash before saving!
          Email:    email,
          IsActive: isActive,
      }

      result := configs.DB.Create(&user)
      // ...
  }
  ```

#### 🛠️ Solusi & Perbaikan yang Diterapkan:
1. Mengimpor library `golang.org/x/crypto/bcrypt` ke dalam `handlers/user-handler.go`.
2. Menambahkan enkripsi password menggunakan `bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)` sebelum struct `models.User` disimpan ke database.
3. Memberikan penanganan error jika proses hashing gagal (`HTTP 500`).
4. Menyimpan kodingan lama dalam bentuk komentar (commented-out) beranotasi di source code untuk keperluan belajar dan ulasan.
- **Kode Sesudah Fix:**
  ```go
  func CreateUser(c *gin.Context) {
      username := c.PostForm("username")
      password := c.PostForm("password")
      email := c.PostForm("email")
      isActiveStr := c.PostForm("is_active")

      isActive, _ := strconv.ParseBool(isActiveStr)

      // ❌ SEBELUMNYA (BUG): Password disimpan langsung secara polos (plaintext)
      // user := models.User{
      // 	Username: username,
      // 	Password: password,
      // 	Email:    email,
      // 	IsActive: isActive,
      // }

      // ✅ PERBAIKAN: Hash password dengan Bcrypt sebelum disimpan ke database
      hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
      if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
          return
      }

      user := models.User{
          Username: username,
          Password: string(hashedPassword),
          Email:    email,
          IsActive: isActive,
      }

      result := configs.DB.Create(&user)
      if result.Error != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
          return
      }

      c.Redirect(http.StatusSeeOther, "/admin/users")
  }
  ```

### ✅ 3. [FIXED] Automatic Database Truncation on Every Server Start
- **Location:** [`main.go:42-43`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/main.go#L42-L43), [`databases/seeders/seeder-products.go:19`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/seeders/seeder-products.go#L19), [`databases/seeders/seeder-user.go:25`](file:///c:/Users/LOQ/Documents/Project_listrik/ProjectEcommerceGolang_BNCC/Day10/databases/seeders/seeder-user.go#L25)
- **Status:** **TERATASI (RESOLVED)**

#### 🔍 Kondisi Awal (Sebelum Fix):
- **Severity:** **HIGH**
- **Impact:** Setiap kali server Go dimatikan dan dinyalakan ulang (`go run main.go`), perintah `TRUNCATE TABLE` dieksekusi secara membabi buta. Akibatnya, seluruh produk baru atau pengguna baru yang ditambahkan oleh admin melalui aplikasi web akan langsung hilang terhapus, ter-reset kembali ke 5 produk awal dan 3 user seeder.
- **Kode Asli (Bermasalah):**
  ```go
  // seeder-products.go & seeder-user.go
  configs.DB.Exec("TRUNCATE TABLE products")
  if err := configs.DB.Create(&products).Error; err != nil {
      log.Fatalf("failed seeding products: %v", err)
  }
  ```

#### 🛠️ Solusi & Perbaikan yang Diterapkan:
1. Menghapus eksekusi perintah destruktif `TRUNCATE TABLE`.
2. Menerapkan pola seeder **Idempoten**: Menghitung terlebih dahulu jumlah data di tabel menggunakan `configs.DB.Model(&models.Entity{}).Count(&count)`.
3. Hanya melakukan seeding awal jika tabel benar-benar masih kosong (`count == 0`).
4. Jika tabel sudah memiliki data (`count > 0`), proses seeding dilewati dan menampilkan log informatif `"Products/Users already exist, skipping seeding."`.
5. Menyimpan kodingan lama dalam bentuk komentar (commented-out) beranotasi di source code untuk keperluan belajar dan ulasan.
- **Kode Sesudah Fix:**
  ```go
  // ❌ SEBELUMNYA (BUG): TRUNCATE menghapus 100% isi tabel setiap kali server restart
  // configs.DB.Exec("TRUNCATE TABLE products")
  // if err := configs.DB.Create(&products).Error; err != nil {
  // 	log.Fatalf("failed seeding products: %v", err)
  // }
  // log.Println("Re-seeded products successfully 🚀")

  // ✅ PERBAIKAN: Hanya isi data seeder jika tabel masih kosong (Count == 0)
  var count int64
  configs.DB.Model(&models.Product{}).Count(&count)
  if count == 0 {
      if err := configs.DB.Create(&products).Error; err != nil {
          log.Fatalf("failed seeding products: %v", err)
      }
      log.Println("Seeded products successfully 🚀")
  } else {
      log.Println("Products already exist, skipping seeding.")
  }
  ```

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
