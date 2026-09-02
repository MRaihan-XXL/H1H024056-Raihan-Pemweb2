# Tabel Perbandingan Laravel vs Fiber

| Aspek | Laravel 13 (PHP) | Fiber v3 (Go) |
|-------|------------------|---------------|
| **Bahasa Pemrograman** | PHP (interpreted) | Go (compiled ke binary) |
| **Ukuran Proyek Awal** | ~70 MB (vendor/ + dependencies) | ~8 MB (single binary + go.mod) |
| **Kecepatan Install** | ~3-5 menit (composer install) | ~30 detik (go get) |
| **Port Default** | 8000 (`php artisan serve`) | 3000 (`go run main.go`) |
| **Cara Jalankan** | `php artisan serve` (butuh PHP runtime) | `go run main.go` atau jalankan binary langsung |
| **Konfigurasi Awal** | Banyak file konfigurasi (config/, .env, bootstrap/) | Minimal (hanya main.go + go.mod) |
| **Dependency Manager** | Composer (composer.json) | Go Modules (go.mod) |
| **Output** | Membutuhkan web server (PHP built-in/Apache) | Binary tunggal, bisa langsung dijalankan |
| **File yang Di-commit** | Kecuali vendor/ dan .env | Kecuali binary *.exe dan .env |
| **Kesesuaian untuk** | Aplikasi bisnis, CRUD kompleks, admin panel | Microservices, API ringan, performa tinggi |

## Kesimpulan
- **Laravel** lebih cocok untuk pengembangan cepat dengan banyak fitur bawaan (ORM, autentikasi, migrasi)
- **Fiber** lebih cocok untuk layanan yang menuntut latensi rendah dan deployment sederhana (single binary)
- Keduanya bisa membangun RESTful API yang sesuai kaidah
