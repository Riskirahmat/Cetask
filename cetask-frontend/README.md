# Cetask Frontend - Capstone Project

Cetask adalah aplikasi Kanban yang dikembangkan sebagai bagian dari capstone project. Aplikasi ini membantu pengguna dalam mengelola tugas dengan tampilan papan (board) seperti Trello.

## 🛠️ Teknologi yang Digunakan

Frontend Cetask dikembangkan menggunakan teknologi berikut:
- **React.js** - Library utama untuk membangun UI.
- **Chakra UI** - Framework styling berbasis komponen untuk desain yang responsif dan mudah dikustomisasi.
- **React Router** - Digunakan untuk menangani navigasi antar halaman.
- **Axios** - Digunakan untuk melakukan request API ke backend.
- **Dnd-Kit** - Library untuk implementasi drag and drop yang lebih modern.
- **React Beautiful DnD** - Digunakan untuk fitur drag and drop pada board.
- **Framer Motion** - Library untuk animasi pada React.
- **React Hook Form** - Digunakan untuk manajemen formulir yang lebih efisien.
- **React Quill** - Editor teks berbasis WYSIWYG.

### 📦 Dependensi Utama
Berikut adalah beberapa dependensi utama yang digunakan dalam frontend Cetask:

```json
{
  "@chakra-ui/icons": "^2.2.4",
  "@chakra-ui/react": "^2.8.0",
  "@dnd-kit/accessibility": "^3.1.1",
  "@dnd-kit/core": "^6.3.1",
  "@dnd-kit/sortable": "^10.0.0",
  "@emotion/react": "^11.11.0",
  "@emotion/styled": "^11.11.0",
  "axios": "^1.7.9",
  "framer-motion": "^10.12.16",
  "jwt-decode": "^4.0.0",
  "react": "^18.2.0",
  "react-beautiful-dnd": "^13.1.1",
  "react-dom": "^18.2.0",
  "react-hook-form": "^7.54.2",
  "react-quill": "^2.0.0",
  "react-router-dom": "^7.1.5"
}
```

## 📌 Struktur Proyek

Struktur utama dari frontend Cetask menggunakan React.js dengan Chakra UI untuk styling. Struktur proyeknya sebagai berikut:

```
cetask-fe/
│── public/           # Static files (favicon, logo, dll.)
│── src/
│   ├── api/          # API calls (Axios)
│   ├── assets/       # Gambar dan ikon
│   ├── components/   # Komponen UI reusable
│   ├── hooks/        # Custom hooks
│   ├── layouts/      # Struktur layout aplikasi
│   ├── pages/        # Halaman utama aplikasi
│   ├── routes/       # Konfigurasi React Router
│   ├── store/        # State management (jika diperlukan)
│   ├── theme/        # Konfigurasi tema Chakra UI
│   ├── utils/        # Fungsi helper
│   ├── App.js        # Root component
│   ├── main.js       # Entry point aplikasi
│── .env.example      # Contoh konfigurasi environment variables
│── package.json      # Dependencies dan script
│── README.md         # Dokumentasi proyek
```

## 🎨 Setup & Instalasi

### 📌 Prerequisites

Pastikan Anda telah menginstal:

- [Node.js](https://nodejs.org/en/download/) (minimal versi 16)
- [npm](https://www.npmjs.com/) atau [yarn](https://yarnpkg.com/)

### 🔧 Instalasi

1. Clone repository ini:
   ```sh
   git clone https://github.com/username/cetask-fe.git
   cd cetask-fe
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
4. Jalankan aplikasi frontend:
   ```sh
   npm run dev
   ```
5. Frontend akan berjalan di `http://localhost:3000`.

```

## 🚀 Fitur Saat Ini

Saat ini, frontend Cetask memiliki fitur berikut:

- Input proyek baru ke dalam sistem.
- Tampilan daftar proyek yang sudah dibuat.
- Navigasi dasar antar halaman.

## 🔧 Pengembangan Berkelanjutan

Fitur yang sedang dikembangkan dan direncanakan untuk ditambahkan:

- Manajemen task dalam proyek.
- Drag & drop untuk mengatur task.
- Peningkatan error handling.
- UI/UX improvements.

## ❗ Issues Saat Ini

- Frontend hanya bisa sampai tahap input project.
- Beberapa fitur masih belum berfungsi sepenuhnya.

---

**Cetask** 🚀

