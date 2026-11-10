# E. Tugas Praktikum

>Catatan screenshot: fitur autentikasi Fiber ini berupa REST API, bukan halaman website, jadi tidak ada tampilan web login/refresh yang bisa dibuka langsung. Untuk bukti hasil request, buka [Postman Web](https://web.postman.co/) dan kirim request ke API lokal. Untuk bukti kode, klik tautan file yang dicantumkan lalu ambil screenshot di editor. Screenshot basis data diambil dari klien MySQL. Cukup siapkan gambar `01` sampai `09` di bawah; bagian analisis nomor 4 tidak memerlukan screenshot tersendiri.

### 1. Tambahkan endpoint `POST /api/auth/segarkan` yang menerbitkan token baru berdasarkan token yang masih berlaku.

  >Endpoint refresh hanya menerima JWT yang masih berlaku. Respons berisi token baru; token lama langsung dimasukkan ke daftar hitam agar tidak dapat digunakan kembali.
  >
  >File yang digunakan: [`auth_handler.go`](../latihan-fiber-api/internal/handler/auth_handler.go) (`Segarkan`), [`router.go`](../latihan-fiber-api/internal/router/router.go), [`token.go`](../latihan-fiber-api/internal/auth/token.go), dan [`autentikasi.go`](../latihan-fiber-api/internal/middleware/autentikasi.go).
  >
  >**01 — kode:** klik [`auth_handler.go`](../latihan-fiber-api/internal/handler/auth_handler.go), buka fungsi `Segarkan`, lalu ambil screenshot kode.<br>
  >**02 — berhasil:** di [Postman Web](https://web.postman.co/), pilih metode `POST` dan alamat `http://localhost:3000/api/auth/segarkan`. Pada tab Authorization atur tipe Bearer dan masukkan token hasil login, lalu klik **Send**. Screenshot respons berhasil.<br>
  >**03 — token lama:** kirim ulang request yang sama dengan token sebelum refresh. Screenshot respons `401`. Samarkan token pada seluruh gambar.

![Cuplikan handler dan route refresh token](01-segarkan-kode.png)

![Respons refresh token berhasil](02-segarkan-berhasil.png)

![Token lama ditolak setelah refresh](03-token-lama-setelah-refresh.png)

### 2. Tambahkan aturan validasi kustom untuk memastikan NIM diawali huruf H diikuti angka.

  >Aturan validator kustom `nim` menerima format `H` diikuti satu atau lebih angka, misalnya `H12345678`. Aturan diterapkan saat membuat maupun memperbarui data mahasiswa. NIM yang tidak sesuai menghasilkan status `400`.
  >
  >File yang digunakan: [`validasi.go`](../latihan-fiber-api/internal/validasi/validasi.go) dan [`mahasiswa_handler.go`](../latihan-fiber-api/internal/handler/mahasiswa_handler.go).
  >
  >**04 — kode:** klik [`validasi.go`](../latihan-fiber-api/internal/validasi/validasi.go), lalu tampilkan aturan `nim`.<br>
  >**05 — hasil:** di [Postman Web](https://web.postman.co/), pilih `POST` dan alamat `http://localhost:3000/api/mahasiswa`. Pada tab Headers isi `X-Api-Key` dengan API key lokal; pada tab Body pilih raw JSON. Kirim data sekali dengan NIM valid (misalnya `H12345678`) dan sekali dengan NIM tidak valid. Screenshot kedua hasil/status respons; jangan tampilkan API key.

![Aturan validasi kustom NIM](04-validasi-nim-kode.png)

![Hasil pengujian NIM valid dan tidak valid](05-validasi-nim-api.png)

### 3. Buat tabel daftar hitam token sederhana agar endpoint logout dapat mencabut token sebelum kedaluwarsa.

  >Tabel `token_blacklists` menyimpan ID unik token (JTI) dan waktu kedaluwarsanya. `POST /api/auth/logout` memasukkan token aktif ke tabel tersebut; middleware memeriksa daftar hitam pada setiap endpoint autentikasi. Token yang dicabut ditolak dengan status `401`. Baris yang sudah kedaluwarsa dibersihkan saat token dicabut.
  >
  >File yang digunakan: [`pengguna.go`](../latihan-fiber-api/internal/model/pengguna.go), [`migrasi.go`](../latihan-fiber-api/internal/config/migrasi.go), [`token_blacklist_repository.go`](../latihan-fiber-api/internal/repository/token_blacklist_repository.go), [`auth_handler.go`](../latihan-fiber-api/internal/handler/auth_handler.go), dan [`autentikasi.go`](../latihan-fiber-api/internal/middleware/autentikasi.go).
  >
  >**06 — kode:** klik [`pengguna.go`](../latihan-fiber-api/internal/model/pengguna.go) untuk model blacklist, lalu tampilkan fungsi `Logout` pada [`auth_handler.go`](../latihan-fiber-api/internal/handler/auth_handler.go).<br>
  >**07 — logout:** di [Postman Web](https://web.postman.co/), kirim `POST http://localhost:3000/api/auth/logout` dengan otorisasi Bearer yang masih berlaku. Screenshot respons berhasil.<br>
  >**08 — database:** buka database `pemweb2_fiber` di klien MySQL, pilih tabel `token_blacklists`, lalu screenshot baris JTI token yang dicabut.
  >**09 — token dicabut:** di Postman kirim `GET http://localhost:3000/api/auth/profil` dengan token yang sama. Screenshot respons `401`. Samarkan token dan kredensial.

![Model dan mekanisme daftar hitam token](06-blacklist-kode.png)

![Logout berhasil](07-logout-berhasil.png)

![JTI token tersimpan di tabel daftar hitam](08-token-blacklist-di-database.png)

![Token ditolak setelah logout](09-token-ditolak-setelah-logout.png)

### 4. Bandingkan implementasi autentikasi Modul 5 dan Modul 8 dari jumlah kode, kemudahan, dan fleksibilitas.

  >Penghitungan LOC berikut menggunakan baris kode produksi yang tidak kosong dan bukan komentar. Cakupan Laravel mencakup controller autentikasi, model pengguna, middleware peran admin, migration kolom peran/waktu login, route autentikasi, dan pendaftaran middleware. Cakupan Fiber mencakup handler autentikasi, utilitas JWT, middleware autentikasi, repository blacklist, model pengguna/blacklist, route autentikasi, konfigurasi JWT, migration, dan pemeriksaan konfigurasi saat aplikasi dimulai. Tes, dependensi, validasi NIM, dan kode route yang tidak terkait autentikasi tidak dihitung.
  >
  >| Modul | Implementasi | Rincian LOC | Total |
  >|---|---|---|---:|
  >| 5 | Laravel + Sanctum | Controller 119 + model 21 + middleware 18 + migration 20 + route 8 + pendaftaran middleware 2 | **188** |
  >| 8 | Fiber + JWT | Handler 198 + utilitas token 80 + middleware 39 + repository blacklist 29 + model 21 + route 8 + konfigurasi JWT 9 + migration 2 + pemeriksaan konfigurasi 4 | **390** |
  >
  >Pada cakupan tersebut, implementasi Fiber membutuhkan 202 LOC lebih banyak. Angka ini merupakan pengukuran berkas dan bagian kode autentikasi pada repositori ini, bukan ukuran universal; fitur kedua modul juga tidak sepenuhnya sama. Laravel lebih ringkas karena Sanctum, validasi, hashing, dan middleware disediakan framework. Fiber memberi kendali eksplisit atas klaim serta proses tanda tangan dan verifikasi JWT, tetapi membutuhkan lebih banyak kode untuk middleware dan pencabutan token. Sanctum memudahkan pencabutan dengan menghapus token tersimpan; JWT lebih fleksibel untuk membawa klaim dan digunakan lintas layanan, tetapi logout memerlukan blacklist dan pemeriksaan tambahan.

# F. Pertanyaan Pembahasan

### 1. Mengapa bcrypt lebih sesuai untuk kata sandi dibandingkan fungsi hash cepat seperti SHA-256?

  >bcrypt dirancang untuk hashing kata sandi dan sengaja dibuat lambat. Faktor biayanya dapat dinaikkan sehingga tebakan kata sandi secara massal menjadi semakin mahal. bcrypt juga menggunakan salt sehingga kata sandi yang sama tidak menghasilkan hash yang sama. SHA-256 sangat cepat dan dirancang sebagai fungsi hash umum; kecepatan itu memudahkan penyerang mencoba banyak tebakan per detik jika hash kata sandi bocor.

### 2. Apa konsekuensi apabila nilai `JWT_SECRET` bocor?

  >Pihak yang memperoleh `JWT_SECRET` dapat membuat token bertanda tangan yang terlihat sah, termasuk mengaku sebagai pengguna lain atau memalsukan klaim peran. Secret harus kuat, acak, dan disimpan di environment, bukan di repositori. Jika bocor, secret harus segera dirotasi; rotasi akan membuat token lama yang ditandatangani dengan secret sebelumnya tidak lagi dapat diverifikasi.

### 3. Mengapa penerapan logout pada JWT memerlukan usaha tambahan dibandingkan token tersimpan?

  >JWT umumnya diverifikasi dari tanda tangan dan masa berlaku tanpa mencari token ke basis data. Karena itu, server tidak memiliki catatan untuk langsung menghapus sebuah token saat logout; token yang sudah diterbitkan tetap berlaku sampai `exp`. Implementasi ini menambahkan tabel blacklist berisi JTI dan masa kedaluwarsa, lalu middleware memeriksanya sebelum menerima token. Token tersimpan seperti Sanctum lebih mudah dicabut karena baris tokennya dapat langsung dihapus.

### 4. Jelaskan alasan pesan login gagal sebaiknya tidak menyebutkan apakah email atau kata sandi yang salah.

  >Pesan umum seperti “Email atau kata sandi tidak sesuai” mencegah enumerasi akun. Jika respons menyatakan email tidak terdaftar, penyerang dapat mengumpulkan alamat yang memiliki akun dan menargetkannya dengan brute force atau phishing. Respons yang seragam untuk kedua kondisi mengurangi informasi yang dibocorkan kepada pihak yang mencoba login.

## Menjalankan Project dan Mengambil Screenshot

Atur `JWT_SECRET` dan `JWT_TTL_MENIT` pada berkas `.env` di folder `latihan-fiber-api`. Gunakan secret acak minimal 32 byte dan jangan tampilkan nilainya pada screenshot. Pastikan MySQL berjalan dan database yang dikonfigurasi dapat diakses, lalu jalankan API dari folder tersebut:

```powershell
go run ./cmd/api
```

Registrasikan pengguna melalui `POST /api/auth/registrasi`, lalu login melalui `POST /api/auth/login` untuk memperoleh token pengujian. Pada request mahasiswa, sertakan `X-Api-Key` dari konfigurasi lokal. Simpan screenshot di folder `Tugas-8` dengan nama yang dicantumkan pada setiap tugas; samarkan token, secret, API key, dan kata sandi sebelum menyimpan atau membagikannya.