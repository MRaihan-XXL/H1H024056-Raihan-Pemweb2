# E. Tugas Praktikum

## Persiapan Sanctum

Project Laravel sudah menggunakan Laravel Sanctum untuk autentikasi API berbasis token. Jalankan dari folder `latihan-laravel`:

```powershell
php artisan migrate:fresh --seed
php artisan serve --host=127.0.0.1 --port=8000
```

Akun uji yang tersedia:

| Peran | Email | Kata sandi | Kemampuan token |
|---|---|---|---|
| Admin | `admin@unsoed.ac.id` | `rahasia123` | `mahasiswa:baca`, `mahasiswa:tulis` |
| Mahasiswa | `mahasiswa@unsoed.ac.id` | `rahasia123` | `mahasiswa:baca` |

File utama Sanctum:

- `latihan-laravel/app/Models/User.php`
- `latihan-laravel/app/Http/Controllers/Api/AuthController.php`
- `latihan-laravel/app/Http/Middleware/PeranAdmin.php`
- `latihan-laravel/routes/api.php`
- `latihan-laravel/bootstrap/app.php`
- `latihan-laravel/database/migrations/2026_09_29_000001_add_peran_and_terakhir_login_to_users_table.php`

## 1. Endpoint PUT `/api/auth/password`

Endpoint ini hanya dapat digunakan oleh pengguna yang sudah login. Kata sandi lama harus benar sebelum kata sandi baru disimpan.

```text
PUT http://127.0.0.1:8000/api/auth/password
```

Header:

```text
Accept: application/json
Authorization: Bearer <token-login>
Content-Type: application/json
```

Body:

```json
{
    "kata_sandi_lama": "rahasia123",
    "kata_sandi_baru": "rahasia456",
    "kata_sandi_baru_confirmation": "rahasia456"
}
```

Respons berhasil:

```json
{
    "sukses": true,
    "pesan": "Kata sandi berhasil diubah."
}
```

Implementasi berada di method `ubahPassword()` pada `AuthController`.

## 2. Middleware `PeranAdmin`

Middleware `PeranAdmin` menolak permintaan apabila pengguna yang login bukan admin. Middleware ini diterapkan pada route penghapusan mahasiswa bersama kemampuan token `mahasiswa:tulis`.

```text
DELETE /api/mahasiswa/{mahasiswa}
```

Aturan akses:

- Admin dengan token `mahasiswa:tulis`: berhasil, status `200`.
- Mahasiswa dengan token `mahasiswa:baca`: ditolak, status `403`.
- Tanpa token: ditolak, status `401`.

File yang digunakan:

- `latihan-laravel/app/Http/Middleware/PeranAdmin.php`
- `latihan-laravel/bootstrap/app.php`
- `latihan-laravel/routes/api.php`

## 3. Kolom `terakhir_login`

Kolom `terakhir_login` ditambahkan pada tabel `users` melalui migration:

```text
latihan-laravel/database/migrations/2026_09_29_000001_add_peran_and_terakhir_login_to_users_table.php
```

Setiap login berhasil akan mengisi kolom tersebut dengan waktu saat login. Nilai ini dapat dilihat melalui endpoint profil:

```text
GET http://127.0.0.1:8000/api/auth/profil
```

Contoh bagian respons:

```json
{
    "peran": "admin",
    "terakhir_login": "2026-09-29T10:00:00.000000Z",
    "kemampuan": [
        "mahasiswa:baca",
        "mahasiswa:tulis"
    ]
}
```

## 4. Skenario Pengujian Postman

Buat koleksi Postman bernama `Pemweb2 Sanctum`. Simpan token login pada environment dengan nama `token`, lalu gunakan `Bearer Token` bernilai `{{token}}` pada request yang membutuhkan autentikasi.

### Skenario A: Login berhasil

```text
POST http://127.0.0.1:8000/api/auth/login
```

Body:

```json
{
    "email": "admin@unsoed.ac.id",
    "password": "rahasia123"
}
```

Hasil yang diharapkan: status `200`, `sukses: true`, dan token.

Pada tab **Tests** Postman, simpan token:

```javascript
const data = pm.response.json();
pm.environment.set('token', data.data.token);
```

### Skenario B: Login gagal

Gunakan password yang salah:

```json
{
    "email": "admin@unsoed.ac.id",
    "password": "salah123"
}
```

Hasil yang diharapkan: status `401` dan pesan email atau kata sandi tidak sesuai.

### Skenario C: Akses tanpa token

```text
GET http://127.0.0.1:8000/api/auth/profil
```

Jangan kirim header `Authorization`. Hasil yang diharapkan: status `401` dan respons JSON token tidak valid atau belum dikirim.

### Skenario D: Token kadaluwarsa setelah logout

1. Login admin dan simpan token.
2. Kirim:

```text
POST http://127.0.0.1:8000/api/auth/logout
```

3. Kirim kembali:

```text
GET http://127.0.0.1:8000/api/auth/profil
```

Hasil yang diharapkan: status `401`, karena token sudah dihapus dari tabel `personal_access_tokens` di sisi server.

### Skenario E: Kemampuan tidak cukup

1. Login menggunakan akun:

```text
Email: mahasiswa@unsoed.ac.id
Kata sandi: rahasia123
```

2. Gunakan token tersebut untuk mengirim:

```text
POST http://127.0.0.1:8000/api/mahasiswa
```

Hasil yang diharapkan: status `403`, karena token hanya memiliki kemampuan `mahasiswa:baca`.

Untuk menguji middleware peran admin secara langsung, gunakan token mahasiswa untuk:

```text
DELETE http://127.0.0.1:8000/api/mahasiswa/1
```

Hasil yang diharapkan juga status `403`.

# F. Pertanyaan Pembahasan

## 1. Mengapa kata sandi harus disimpan dalam bentuk hash dan bukan terenkripsi dua arah?

Kata sandi harus disimpan dalam bentuk hash karena aplikasi hanya perlu memeriksa apakah kata sandi yang dimasukkan cocok, bukan mengembalikan kata sandi asli. Hash bersifat satu arah sehingga hasilnya tidak dapat dibalik secara langsung menjadi kata sandi awal.

Enkripsi dua arah memiliki kunci untuk mengembalikan data asli. Jika kunci tersebut bocor, semua kata sandi dapat dibuka. Karena itu, kata sandi sebaiknya menggunakan hash yang kuat seperti bcrypt atau Argon2. Pada project ini Laravel menggunakan `Hash::make()` dan cast `password` bertipe `hashed`.

## 2. Apa perbedaan antara kemampuan token pada Sanctum dan peran pengguna pada tabel users?

Kemampuan token adalah izin khusus yang melekat pada satu token, misalnya `mahasiswa:baca` atau `mahasiswa:tulis`. Satu pengguna dapat memiliki beberapa token dengan kemampuan yang berbeda.

Peran pengguna adalah identitas atau kelompok akses yang disimpan pada tabel `users`, misalnya `admin` atau `mahasiswa`. Peran dapat digunakan oleh middleware `PeranAdmin` untuk menentukan apakah pengguna boleh melakukan tindakan tertentu.

Jadi, peran menjelaskan siapa pengguna tersebut, sedangkan kemampuan token menjelaskan tindakan yang diizinkan oleh token tertentu. Pada endpoint hapus mahasiswa, keduanya diperiksa: token harus memiliki `mahasiswa:tulis` dan pengguna harus memiliki peran `admin`.

## 3. Jelaskan risiko penyimpanan token pada localStorage peramban serta alternatif yang lebih aman.

Token pada `localStorage` dapat dibaca oleh JavaScript. Jika aplikasi terkena serangan XSS, script berbahaya dapat mencuri token dan menggunakannya untuk mengakses API sebagai pengguna tersebut.

Alternatif yang lebih aman adalah menggunakan cookie `HttpOnly`, `Secure`, dan `SameSite` sehingga token tidak dapat dibaca langsung oleh JavaScript. Untuk aplikasi SPA satu domain, Laravel Sanctum juga menyediakan autentikasi berbasis cookie dan perlindungan CSRF.

Jika API token tetap digunakan untuk klien eksternal, token harus disimpan menggunakan penyimpanan aman milik sistem operasi dan tidak boleh dimasukkan ke repository, log, atau query string.

## 4. Mengapa endpoint logout perlu menghapus token di sisi server padahal token disimpan pada klien?

Menghapus token di klien saja tidak cukup karena salinan token mungkin masih ada di perangkat lain, riwayat, atau tempat penyimpanan sementara. Selama token masih ada di tabel `personal_access_tokens`, token tersebut masih dapat diterima oleh server.

Endpoint logout menghapus token aktif menggunakan `currentAccessToken()->delete()`. Setelah itu, token lama tidak lagi berlaku meskipun seseorang masih memilikinya. Inilah yang membuat logout dapat mencabut akses secara nyata di sisi server.

# Lokasi Screenshot

## Screenshot file kode

Ambil screenshot bagian berikut dari VS Code:

1. `latihan-laravel/app/Http/Controllers/Api/AuthController.php`
   - method `login()` untuk `terakhir_login`
   - method `ubahPassword()`
   - method `logout()`
2. `latihan-laravel/app/Http/Middleware/PeranAdmin.php`
3. `latihan-laravel/routes/api.php`
   - route password
   - middleware `auth:sanctum`
   - middleware `ability:mahasiswa:tulis`
   - middleware `peran.admin`
4. `latihan-laravel/database/migrations/2026_09_29_000001_add_peran_and_terakhir_login_to_users_table.php`
5. `latihan-laravel/app/Models/User.php`
   - trait `HasApiTokens`
   - `peran` pada `$fillable`
   - cast `terakhir_login`
6. `latihan-laravel/bootstrap/app.php`
   - alias middleware `peran.admin`
   - respons error `401`

## Screenshot halaman API di browser

Jalankan server, lalu buka:

- http://127.0.0.1:8000/api/status
- http://127.0.0.1:8000/api/auth/profil setelah memasukkan token melalui Postman, bila browser mendukung header token

Browser cocok untuk dokumentasi endpoint GET publik. Untuk endpoint yang membutuhkan header Authorization, gunakan Postman.

## Screenshot Postman yang wajib

Ambil screenshot halaman Postman yang menampilkan URL, metode, header, body, status, dan respons:

1. Login berhasil: `POST /api/auth/login` status `200`
2. Login gagal: `POST /api/auth/login` status `401`
3. Akses tanpa token: `GET /api/auth/profil` status `401`
4. Ubah kata sandi: `PUT /api/auth/password` status `200`
5. Logout: `POST /api/auth/logout` status `200`
6. Token setelah logout: `GET /api/auth/profil` status `401`
7. Akses mahasiswa tanpa kemampuan cukup: `POST /api/mahasiswa` status `403`
8. Hapus mahasiswa dengan akun mahasiswa: `DELETE /api/mahasiswa/1` status `403`
9. Hapus mahasiswa dengan akun admin: `DELETE /api/mahasiswa/1` status `200`

Jangan tampilkan nilai token secara penuh pada screenshot laporan. Tutup atau samarkan sebagian token sebelum memasukkannya ke laporan.
