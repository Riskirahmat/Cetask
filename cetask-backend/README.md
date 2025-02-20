# Cetask Backend - Capstone Project

Cetask adalah aplikasi Kanban yang dikembangkan sebagai bagian dari capstone project. Bagian backend aplikasi ini dibangun menggunakan **Golang** dengan **Gin framework** serta menggunakan **MongoDB** sebagai database utama.

## 🛠️ Teknologi yang Digunakan

Backend Cetask dikembangkan menggunakan teknologi berikut:
- **Golang** - Bahasa pemrograman utama untuk backend.
- **Gin** - Framework untuk membangun REST API dengan performa tinggi.
- **MongoDB** - Database utama untuk menyimpan data aplikasi.
- **JWT (JSON Web Token)** - Untuk otentikasi dan otorisasi pengguna.

## 🎨 Setup & Instalasi

### 📌 Prerequisites

Pastikan Anda telah menginstal:
- [Go](https://go.dev/dl/) (minimal versi 1.18)
- [MongoDB](https://www.mongodb.com/try/download/community)

### 🔧 Instalasi

1. Clone repository ini:
   ```sh
   git clone https://github.com/username/cetask-be.git
   cd cetask-be
   ```
2. Buat file `.env` berdasarkan contoh di `.env.example` dan isi dengan konfigurasi yang sesuai:
   ```sh
   cp .env.example .env
   ```
3. Jalankan aplikasi backend:
   ```sh
   go run main.go
   ```
4. Backend akan berjalan di `http://localhost:8080`.

## 📜 Environment Variables (BE)

```env
PORT=8080
MONGODB_URI=mongodb://username:password@localhost:27017/cetask_db
JWT_SECRET=your_jwt_secret
```

## 🚀 Fitur Saat Ini

Saat ini, backend Cetask memiliki fitur berikut:
- **User Authentication** (Register & Login dengan JWT)
- **CRUD Proyek** (Membuat, membaca, memperbarui, dan menghapus proyek)
- **Middleware Autentikasi** (Melindungi endpoint yang memerlukan login)
- **Validasi Request** (Menjaga integritas data yang diterima API)

## 🔧 Pengembangan Berkelanjutan

Fitur yang sedang dikembangkan dan direncanakan untuk ditambahkan:
- **Manajemen Task dalam proyek**
- **Fitur Drag & Drop** untuk mengubah status task
- **Peningkatan Error Handling & Logging**
- **Unit Testing & Integration Testing**

## ❗ Issues Saat Ini

- API error handling masih memerlukan peningkatan.
- Beberapa fitur CRUD tambahan masih dalam tahap pengembangan.
- 
**Cetask** 🚀

