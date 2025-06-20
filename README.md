# ✨ BAKULOS - Belanja Jadi Serius & Santai! 🛒

**Bakulos** adalah aplikasi e-commerce modern berbasis Flutter + GraphQL + Golang yang dirancang untuk mempertemukan pembeli dan penjual dengan pengalaman belanja yang **cepat**, **interaktif**, dan **aman**.

Temukan produk favoritmu, chat langsung dengan penjual, checkout dalam sekali klik — semua dalam satu aplikasi 📱

---

## 🎯 Tujuan Proyek

Menyediakan platform e-commerce dengan UI/UX yang ramah pengguna, sistem autentikasi aman, serta komunikasi real-time antara pembeli dan penjual.

---

## 💼 Fitur Unggulan

### 👤 Untuk Pembeli:
- 🔐 **Registrasi & Login aman** (JWT Authentication)
- 🏠 **Beranda dinamis** menampilkan berbagai produk
- 🔍 **Pencarian & filter kategori** untuk produk impianmu
- ❤️ **Tambah ke Favorit** dan simpan untuk nanti
- 🛒 **Keranjang belanja cerdas** + total harga otomatis
- 💳 **Checkout cepat** dengan pilihan pengiriman & pembayaran
- 🧾 **Riwayat pembelian** yang terorganisir
- 📍 **Manajemen alamat** utama & alternatif
- 💬 **Chat real-time** dengan penjual
- 🧑‍💼 **Edit profil** & ubah foto pengguna

### 🧾 Untuk Penjual:
- ✍️ **Registrasi khusus penjual** (JWT Authentication)
- 📦 **Manajemen produk** (tambah/edit/hapus)
- 💬 **Chat dengan pembeli** tanpa delay
- 📈 **Edit profil** & ubah foto penjual**

---

## ⚙️ Teknologi yang Digunakan

| Layer       | Stack                                                                 |
|-------------|-----------------------------------------------------------------------|
| **Frontend**| Flutter, GraphQL Flutter, SharedPreferences, WebSocket               |
| **Backend** | Golang, Gin, GORM, GraphQL (gqlgen), JWT Auth, WebSocket             |
| **Database**| MySQL / PostgreSQL (konfigurasi opsional)                            |

---

## 🛠️ Instalasi & Jalankan Proyek

### 🔧 Backend - GraphQL API (Golang)

```bash
git clone https://github.com/username/bakulos-backend.git
cd bakulos-backend

# Install dependensi & jalankan server
go mod tidy
go run main.go
