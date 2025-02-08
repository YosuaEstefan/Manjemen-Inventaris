# Manajemen Inventaris

Proyek ini adalah sistem manajemen inventaris berbasis REST API yang dikembangkan menggunakan **Golang** dengan framework **Gin**, serta menggunakan **MySQL** sebagai database utama.

## 📌 Fitur Utama

- CRUD Produk
- CRUD Inventaris
- CRUD Pesanan
- Penyimpanan dan pengambilan gambar produk
- Middleware untuk autentikasi

---

## 🛠 Persyaratan

Sebelum menjalankan proyek ini, pastikan Anda telah menginstal:

- **Golang** (versi terbaru) - [Download Golang](https://go.dev/dl/)
- **MySQL** (versi terbaru) - [Download MySQL](https://www.mysql.com/downloads/)
- **Git** (untuk clone repository) - [Download Git](https://git-scm.com/downloads)

---

## 📂 Struktur Folder

```
inventory-management/
│── main.go                  // Entry point aplikasi
│── config/
│   ├── database.go          // Koneksi ke MySQL
│── models/
│   ├── product.go           // Model Produk
│   ├── inventory.go         // Model Inventaris
│   ├── order.go             // Model Pesanan
│── controllers/
│   ├── productController.go // Controller Produk
│   ├── inventoryController.go // Controller Inventaris
│   ├── orderController.go   // Controller Pesanan
│── routes/
│   ├── routes.go            // Routing API
│── utils/
│   ├── fileHandler.go       // Penanganan file (upload/download gambar)
│── storage/                 // Menyimpan gambar produk
│── .env                     // Konfigurasi environment
│── go.mod
│── go.sum
│── README.md                // Dokumentasi proyek
```

---

## ⚙️ Konfigurasi Database

1. Buat database baru di MySQL:
   ```sql
   CREATE DATABASE manajemen_inventaris;
   ```
2. Atur kredensial database di file `.env`:
   ```ini
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=Yosua
   DB_NAME=manajemen_inventaris
   ```

---

## 🚀 Cara Menjalankan Proyek

### 1️⃣ Clone Repository

```sh
git clone https://github.com/username/manajemen-inventaris.git
cd manajemen-inventaris
```

### 2️⃣ Install Dependencies

```sh
go mod tidy
```

### 3️⃣ Jalankan Server

```sh
go run main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 🛠 API Endpoint

| Method | Endpoint        | Deskripsi          |
| ------ | --------------- | ------------------ |
| POST   | `/products`     | Tambah produk      |
| GET    | `/products`     | Ambil semua produk |
| PUT    | `/products/:id` | Perbarui produk    |
| DELETE | `/products/:id` | Hapus produk       |

**Gunakan Postman atau Curl untuk menguji API ini.**

---

🚀 **Selamat coding!** Jika ada masalah, silakan ajukan issue atau pull request. 😊

