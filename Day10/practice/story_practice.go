package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
)

// Mencegah error "imported and not used" saat blok latihan masih di-comment
var (
	_ = os.Getwd
	_ = strconv.Itoa
	_ = excelize.NewFile
)

// ============================================================================
// 🎮 PETUNJUK PENGGUNAAN LAB PRAKTEK STORYTELLING:
// ============================================================================
// 1. Jalankan script ini langsung di terminal dengan:
//      go run practice/story_practice.go
// 2. Setiap "Episode" memiliki bagian kode yang bisa Anda "UNCOMMENT"
//    (Hapus tanda komentar // di blok yang sudah ditandai).
// 3. Setelah uncomment, jalankan lagi untuk melihat perubahan outputnya secara spontan!
// ============================================================================

func main() {
	cetakHeader()

	for {
		fmt.Println("\n========================================================")
		fmt.Println("📜 PILIH EPISODE CERITA UNTUK DIPELAJARI:")
		fmt.Println("========================================================")
		fmt.Println(" [1] 🔐 Episode 1: Misteri Brankas Token JWT & Pembobol Kunci")
		fmt.Println(" [2] 🧪 Episode 2: Sang Alkemis Password Bcrypt vs Plaintext")
		fmt.Println(" [3] 🕵️  Episode 3: Detektif & Kasus User Menghilang (Bug Delete)")
		fmt.Println(" [4] 📊 Episode 4: Pabrik Laporan Otomatis Excelize")
		fmt.Println(" [5] ⚡ Episode 5: Kutukan Amnesia Database (Seeder Truncate)")
		fmt.Println(" [0] 🚪 Keluar")
		fmt.Print("\n👉 Masukkan nomor pilihan (0-5): ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch strings.TrimSpace(pilihan) {
		case "1":
			episode1JWT()
		case "2":
			episode2Bcrypt()
		case "3":
			episode3DeleteBug()
		case "4":
			episode4Excelize()
		case "5":
			episode5SeederAmnesia()
		case "0":
			fmt.Println("\n👋 Terima kasih telah belajar! Terus eksplorasi kodingan Go-mu!")
			return
		default:
			fmt.Println("⚠️ Pilihan tidak valid, silakan ketik angka 1 - 5.")
		}
	}
}

// ============================================================================
// 🔐 EPISODE 1: MISTERI BRANKAS TOKEN JWT
// Lokasi Kode Terkait di Repo: Day10/middlewares/jwt.go & Day10/handlers/auth-handler.go
// ============================================================================
func episode1JWT() {
	fmt.Println("\n-------------------------------------------------------------")
	fmt.Println("📖 EPISODE 1: MISTERI BRANKAS TOKEN JWT (JSON WEB TOKEN)")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("CERITA:")
	fmt.Println("Toko 'Maju Jaya' baru saja membuat sistem login. Admin bernama 'Budi' berhasil")
	fmt.Println("login. Sistem tidak memakai session di server, melainkan memberikan selembar")
	fmt.Println("'Tiket Digital Bertanda Tangan' (JWT Token) yang disimpan di Cookie browser.")
	fmt.Println("Namun, seorang hacker mencoba memalsukan tiket tersebut! Apa yang terjadi?")

	secretKeyAsli := []byte("rahasia") // Ini kunci yang ada di middlewares/jwt.go

	// 1. Membuaat Token untuk Budi
	claims := &jwt.RegisteredClaims{
		Subject:   "budi_admin",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStringAsli, err := token.SignedString(secretKeyAsli)
	if err != nil {
		fmt.Println("❌ Gagal membuat token:", err)
		return
	}

	fmt.Println("✅ [SISTEM TOKO]: Budi berhasil login! Token JWT diterbitkan:")
	fmt.Println("🔑 Token Asli:", tokenStringAsli)

	// Membedah 3 bagian JWT (Header.Payload.Signature)
	bagian := strings.Split(tokenStringAsli, ".")
	if len(bagian) == 3 {
		fmt.Println("\n🔍 BEDAH STRUKTUR TOKEN:")
		fmt.Println("   1. Header    :", bagian[0], "(Algoritma HS256)")
		fmt.Println("   2. Payload   :", bagian[1], "(Data: username 'budi_admin')")
		fmt.Println("   3. Signature :", bagian[2], "(Tanda tangan rahasia server)")
	}

	// ------------------------------------------------------------------------
	// 🛠️ TANTANGAN PRAKTEK:
	// Hacker mencoba memverifikasi token dengan kunci palsu ("kunci_palsu_hacker")
	// vs Server Toko yang memverifikasi dengan kunci asli ("rahasia").
	//
	// 👉 COBA UNCOMMENT BLOK DI BAWAH INI UNTUK MELIHAT BAGAIMANA SERVER MENANGKAP PEMBOBOL!
	// ------------------------------------------------------------------------

	/*
		fmt.Println("\n🚨 [SKENARIO HACKER]: Hacker mencoba memalsukan token dengan rahasia sendiri...")
		secretKeyHacker := []byte("kunci_palsu_hacker")

		// Verifikasi dengan kunci palsu
		_, errPalsu := jwt.ParseWithClaims(tokenStringAsli, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return secretKeyHacker, nil
		})

		if errPalsu != nil {
			fmt.Printf("🛡️ [PENJAGA AUTH SERVER]: 'Akses Ditolak! Tanda tangan token tidak cocok! Error: %v'\n", errPalsu)
		} else {
			fmt.Println("😱 Gawat! Hacker berhasil lolos!")
		}

		// Verifikasi dengan kunci asli server (seperti di middlewares/jwt.go)
		tokenValid, errValid := jwt.ParseWithClaims(tokenStringAsli, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return secretKeyAsli, nil
		})

		if errValid == nil && tokenValid.Valid {
			fmt.Println("🎉 [PENJAGA AUTH SERVER]: 'Token valid! Selamat datang Admin Budi!'")
		}
	*/

	fmt.Println("\n💡 TIPS: Buka file practice/story_practice.go di Episode 1, lalu hapus tanda /* dan */ untuk mengaktifkan skenario hacker!")
}

// ============================================================================
// 🧪 EPISODE 2: SANG ALKEMIS PASSWORD BCRYPT
// Lokasi Kode Terkait di Repo: Day10/handlers/user-handler.go vs Day10/databases/seeders/seeder-user.go
// ============================================================================
func episode2Bcrypt() {
	fmt.Println("\n-------------------------------------------------------------")
	fmt.Println("📖 EPISODE 2: SANG ALKEMIS PASSWORD (PLAINTEXT VS BCRYPT)")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("CERITA:")
	fmt.Println("Di Toko Maju Jaya, ada 2 cara user dibuat:")
	fmt.Println("1. Melalui Seeder (seeder-user.go) -> Password diubah jadi hash Bcrypt.")
	fmt.Println("2. Melalui Form Web Admin (user-handler.go) -> Password disimpan POLOS (Plaintext)!")
	fmt.Println("Akibatnya, user yang daftar via web admin TIDAK BISA LOGIN sama sekali!")
	fmt.Println("Mari kita buktikan kenapa hal itu terjadi.")

	passwordInput := "kopi_susu_123"

	// 1. Password disimpan mentah (Seperti bug di user-handler.go baris 51)
	passwordDiDatabasePolos := passwordInput
	fmt.Println("⚠️ [STATUS SEKARANG]: User dibuat dengan password polos:", passwordDiDatabasePolos)

	// Saat login, sistem mengeksekusi bcrypt.CompareHashAndPassword (auth-handler.go baris 24)
	errLoginPolos := bcrypt.CompareHashAndPassword([]byte(passwordDiDatabasePolos), []byte(passwordInput))
	if errLoginPolos != nil {
		fmt.Printf("❌ [LOGIN GAGAL]: bcrypt.CompareHashAndPassword menolak karena password di database bukan format hash valid!\n   Detail Error: %v\n", errLoginPolos)
	}

	// ------------------------------------------------------------------------
	// 🛠️ TANTANGAN PRAKTEK:
	// Sekarang kita gunakan Alkimia Bcrypt seperti di seeder-user.go!
	//
	// 👉 COBA UNCOMMENT BLOK DI BAWAH INI UNTUK MENERAPKAN SOLUSI HASH BCRYPT!
	// ------------------------------------------------------------------------

	/*
		fmt.Println("\n✨ [PERBAIKAN DENGAN BCRYPT]:")
		// Meng-hash password dengan salt otomatis
		hashHasil, err := bcrypt.GenerateFromPassword([]byte(passwordInput), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Gagal hash:", err)
			return
		}
		passwordDiDatabaseHash := string(hashHasil)
		fmt.Println("🔒 Password setelah di-hash:", passwordDiDatabaseHash)
		fmt.Println("   (Lihat: karakter acak $2a$10$... yang aman bahkan jika database bocor!)")

		// Coba login kembali
		errLoginSukses := bcrypt.CompareHashAndPassword([]byte(passwordDiDatabaseHash), []byte(passwordInput))
		if errLoginSukses == nil {
			fmt.Println("🎉 [LOGIN BERHASIL]: Bcrypt mencocokkan password dengan aman! User bisa login!")
		}
	*/

	fmt.Println("\n💡 TIPS: Buka file practice/story_practice.go di Episode 2, lalu hapus tanda /* dan */ untuk mencoba Bcrypt Hash!")
}

// ============================================================================
// 🕵️ EPISODE 3: KASUS USER MENGHILANG (THE DELETE PRODUCT BUG)
// Lokasi Kode Terkait di Repo: Day10/handlers/product-handler.go baris 150-160
// ============================================================================
func episode3DeleteBug() {
	fmt.Println("\n-------------------------------------------------------------")
	fmt.Println("📖 EPISODE 3: KASUS USER MENGHILANG (MISTERI BUG DELETE)")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("CERITA:")
	fmt.Println("Admin ingin menghapus produk 'Keyboard Rusak' dengan ID = 2.")
	fmt.Println("Namun setelah tombol Hapus Produk diklik, produknya MASIH ADA, sedangkan")
	fmt.Println("akun user bernama 'John Doe' (yang kebetulan memiliki ID = 2) TIBA-TIBA TERHAPUS!")
	fmt.Println("Bagaimana mungkin?! Mari kita lihat kodingan di product-handler.go baris 153:")

	fmt.Println("Kode Asli di product-handler.go:")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("func DeleteProduct(c *gin.Context) {")
	fmt.Println("    id := c.Param(\"id\")")
	fmt.Println("    // 🚨 LIHAT INI: Tipe yang di-delete adalah models.User{}!")
	fmt.Println("    if err := configs.DB.Delete(&models.User{}, id).Error; err != nil { ... }")
	fmt.Println("}")
	fmt.Println("-------------------------------------------------------------")

	type SimulasiEntitas struct {
		ID   int
		Nama string
		Tipe string
	}

	daftarUser := map[int]SimulasiEntitas{
		1: {ID: 1, Nama: "Admin Toko", Tipe: "User"},
		2: {ID: 2, Nama: "John Doe (Customer)", Tipe: "User"},
	}

	daftarProduk := map[int]SimulasiEntitas{
		1: {ID: 1, Nama: "Laptop Gaming", Tipe: "Product"},
		2: {ID: 2, Nama: "Keyboard Rusak", Tipe: "Product"},
	}

	idYangMauDihapus := 2
	fmt.Printf("📦 Admin ingin menghapus PRODUK ID: %d ('%s')\n", idYangMauDihapus, daftarProduk[idYangMauDihapus].Nama)

	// Simulasi bug: Menghapus dari daftar User!
	delete(daftarUser, idYangMauDihapus)
	fmt.Printf("💥 AKIBAT BUG: Malah menghapus User ID: %d!\n", idYangMauDihapus)
	fmt.Printf("   Daftar Produk saat ini : %+v (Produk tidak terhapus!)\n", daftarProduk)
	fmt.Printf("   Daftar User saat ini   : %+v (John Doe menghilang!)\n", daftarUser)

	// ------------------------------------------------------------------------
	// 🛠️ TANTANGAN PRAKTEK:
	// Mari kita perbaiki kodingan tersebut menjadi yang benar!
	//
	// 👉 COBA UNCOMMENT BLOK DI BAWAH INI UNTUK MENJALANKAN KODE PERBAIKAN!
	// ------------------------------------------------------------------------

	/*
		fmt.Println("\n🛠️ [PERBAIKAN KODINGAN]:")
		fmt.Println("Mengubah target delete menjadi &models.Product{}!")

		// Kembalikan John Doe
		daftarUser[2] = SimulasiEntitas{ID: 2, Nama: "John Doe (Customer)", Tipe: "User"}

		// Eksekusi hapus yang benar
		delete(daftarProduk, idYangMauDihapus)

		fmt.Println("✅ HASIL SETELAH PERBAIKAN:")
		fmt.Printf("   Daftar Produk : %+v (Keyboard Rusak terhapus sempurna!)\n", daftarProduk)
		fmt.Printf("   Daftar User   : %+v (Akun John Doe aman!)\n", daftarUser)
	*/

	fmt.Println("\n💡 TIPS: Buka file practice/story_practice.go di Episode 3, lalu hapus tanda /* dan */ untuk mencoba perbaikan bug!")
}

// ============================================================================
// 📊 EPISODE 4: PABRIK LAPORAN OTOMATIS DENGAN EXCELIZE
// Lokasi Kode Terkait di Repo: Day10/handlers/public-handler.go baris 36-76
// ============================================================================
func episode4Excelize() {
	fmt.Println("\n-------------------------------------------------------------")
	fmt.Println("📖 EPISODE 4: PABRIK LAPORAN OTOMATIS EXCELIZE")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("CERITA:")
	fmt.Println("Manajer Toko meminta laporan inventaris produk dalam bentuk file Excel (.xlsx).")
	fmt.Println("Di Day10/handlers/public-handler.go, kita menggunakan library 'excelize/v2'.")
	fmt.Println("Mari kita coba buat file Excel biner secara nyata sekarang juga!")

	// ------------------------------------------------------------------------
	// 🛠️ TANTANGAN PRAKTEK:
	// Kode ini akan membuat file Excel sungguhan di disk Anda: "laporan_latihan.xlsx"
	//
	// 👉 COBA UNCOMMENT BLOK DI BAWAH INI UNTUK MENGHASILKAN FILE EXCEL SUNGGUHAN!
	// ------------------------------------------------------------------------

	/*
		f := excelize.NewFile()
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		sheetName := "Produk Toko"
		index, _ := f.NewSheet(sheetName)

		// 1. Tulis Header Kolom
		headers := []string{"ID", "Nama Produk", "Harga (Rp)", "Stok", "Status"}
		for i, h := range headers {
			kolom := string(rune('A' + i))
			f.SetCellValue(sheetName, kolom+"1", h)
		}

		// 2. Data Dummy Produk
		type Item struct {
			ID    int
			Nama  string
			Harga float64
			Stok  int
		}

		items := []Item{
			{ID: 1, Nama: "Laptop Asus ROG", Harga: 18500000, Stok: 5},
			{ID: 2, Nama: "Wireless Mouse Logitech", Harga: 250000, Stok: 20},
			{ID: 3, Nama: "Mechanical Keyboard", Harga: 650000, Stok: 0},
			{ID: 4, Nama: "Monitor 24 inch IPS", Harga: 1750000, Stok: 8},
		}

		// 3. Masukkan Baris Data
		for i, item := range items {
			baris := strconv.Itoa(i + 2)
			f.SetCellValue(sheetName, "A"+baris, item.ID)
			f.SetCellValue(sheetName, "B"+baris, item.Nama)
			f.SetCellValue(sheetName, "C"+baris, item.Harga)
			f.SetCellValue(sheetName, "D"+baris, item.Stok)

			status := "Tersedia"
			if item.Stok == 0 {
				status = "Habis"
			}
			f.SetCellValue(sheetName, "E"+baris, status)
		}

		f.SetActiveSheet(index)

		namaFile := "laporan_latihan.xlsx"
		if err := f.SaveAs(namaFile); err != nil {
			fmt.Println("❌ Gagal menyimpan file excel:", err)
			return
		}

		cwd, _ := os.Getwd()
		fmt.Printf("🎉 BERHASIL! File Excel telah dibuat di:\n   📁 %s\\%s\n", cwd, namaFile)
		fmt.Println("👉 Anda bisa membuka file tersebut langsung di Microsoft Excel atau Google Sheets!")
	*/

	fmt.Println("\n💡 TIPS: Buka file practice/story_practice.go di Episode 4, lalu hapus tanda /* dan */ untuk generate file Excel!")
}

// ============================================================================
// ⚡ EPISODE 5: KUTUKAN AMNESIA DATABASE (SEEDER TRUNCATE)
// Lokasi Kode Terkait di Repo: Day10/databases/seeders/seeder-products.go baris 19
// ============================================================================
func episode5SeederAmnesia() {
	fmt.Println("\n-------------------------------------------------------------")
	fmt.Println("📖 EPISODE 5: KUTUKAN AMNESIA DATABASE (SEEDER TRUNCATE)")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("CERITA:")
	fmt.Println("Setiap kali Anda mematikan server (Ctrl+C) lalu menyalakannya lagi (go run main.go),")
	fmt.Println("seluruh data transaksi atau produk yang baru saja Anda input HILANG!")
	fmt.Println("Kenapa? Karena di seeder-products.go dan seeder-user.go ada perintah:")
	fmt.Println("   configs.DB.Exec(\"TRUNCATE TABLE products\")")
	fmt.Println("Perintah TRUNCATE menghapus 100% isi tabel tanpa ampun!")

	fmt.Println("Solusi yang benar:")
	fmt.Println("Cek dulu apakah tabel sudah ada isinya (Count).")
	fmt.Println("Jika count == 0, barulah jalankan seeder. Jika count > 0, lewati!")

	// ------------------------------------------------------------------------
	// 🛠️ TANTANGAN PRAKTEK:
	// Coba amati perbedaan logika TRUNCATE vs LOGIKA COUNT di bawah ini.
	//
	// 👉 COBA UNCOMMENT BLOK DI BAWAH INI UNTUK MELIHAT SIMULASINYA!
	// ------------------------------------------------------------------------

	/*
		type MockDatabase struct {
			TotalData int
		}

		db := MockDatabase{TotalData: 10} // Katakanlah sudah ada 10 produk dari admin
		fmt.Printf("📊 Keadaan database sebelum restart server: Ada %d produk.\n", db.TotalData)

		fmt.Println("\n--- Skenario Lama (Penyebab Amnesia) ---")
		// Menjalankan truncate
		db.TotalData = 0 // TRUNCATE TABLE
		fmt.Println("💥 [TRUNCATE TABLE]: Seluruh data terhapus bersih!")
		db.TotalData = 5 // Isi 5 data seeder default
		fmt.Printf("   Jumlah data sekarang: %d (Data buatan admin lenyap!)\n", db.TotalData)

		fmt.Println("\n--- Skenario Baru (Perbaikan Idempoten) ---")
		db.TotalData = 10 // Kembalikan kondisi ada 10 produk buatan admin
		jumlahDataSaatIni := db.TotalData

		if jumlahDataSaatIni == 0 {
			fmt.Println("Database masih kosong! Jalankan seeder...")
			db.TotalData = 5
		} else {
			fmt.Printf("🛡️ Database sudah memiliki %d produk! Seeder dilewati (Data admin aman!)\n", jumlahDataSaatIni)
		}
	*/

	fmt.Println("\n💡 TIPS: Buka file practice/story_practice.go di Episode 5, lalu hapus tanda /* dan */ untuk mencoba logika seeder yang aman!")
}

func cetakHeader() {
	fmt.Print(`
   ╔═══════════════════════════════════════════════════════════════╗
   ║        🛍️  GO E-COMMERCE BNCC - LAB PRAKTIK INTERAKTIF        ║
   ║          Belajar Santai, Storytelling & Hands-on Coding       ║
   ╚═══════════════════════════════════════════════════════════════╝
`)
}
