# E. Tugas Praktikum

## 1. CRUD API Mata Kuliah

  >Saya membangun endpoint CRUD lengkap untuk resource `matakuliah`. Implementasi memakai Form Request untuk validasi dan API Resource untuk menyeragamkan respons JSON.
  >
  >File yang digunakan:

  >- `latihan-laravel/app/Http/Controllers/Api/MatakuliahController.php`
  >- `latihan-laravel/app/Http/Requests/StoreMatakuliahRequest.php`
  >- `latihan-laravel/app/Http/Requests/UpdateMatakuliahRequest.php`
  >- `latihan-laravel/app/Http/Resources/MatakuliahResource.php`
  >- `latihan-laravel/routes/api.php`

| Metode | URI | Parameter | Contoh body | Respons berhasil |
|---|---|---|---|---|
| GET | `/api/matakuliah` | `cari`, `semester`, `urut`, `arah`, `per_halaman` | Tidak ada | Daftar data dengan `data`, `links`, dan `meta` |
| GET | `/api/matakuliah/{id}` | `id` mata kuliah | Tidak ada | Satu data mata kuliah |
| POST | `/api/matakuliah` | Tidak ada | `kode`, `nama`, `sks`, `semester` | Status `201` dan data baru |
| PUT/PATCH | `/api/matakuliah/{id}` | `id` mata kuliah | Kolom yang ingin diubah | Status `200` dan data terbaru |
| DELETE | `/api/matakuliah/{id}` | `id` mata kuliah | Tidak ada | Pesan berhasil dihapus |

  >Contoh body POST:
  >
```json
{
    "kode": "IF601",
    "nama": "Pemrograman API",
    "sks": 3,
    "semester": 6
}
```

  >Contoh respons POST:
  >
```json
{
    "sukses": true,
    "pesan": "Mata kuliah berhasil dibuat",
    "data": {
        "id": 11,
        "kode": "IF601",
        "nama": "Pemrograman API",
        "sks": 3,
        "semester": 6
    }
}
```

## 2. Daftar Mahasiswa Berdasarkan Program Studi

  >Endpoint ini mengembalikan daftar mahasiswa dari satu program studi dengan pagination. Pada project ini, program studi disimpan sebagai teks pada kolom `program_studi`, sehingga parameter `{id}` pada modul disesuaikan menjadi nama program studi.

```text
GET /api/program-studi/Informatika/mahasiswa?per_halaman=5
```

  >Respons berisi `data`, `links`, dan `meta` karena endpoint menggunakan pagination.

## 3. Parameter `fields` pada Daftar Mahasiswa

  >Parameter `fields` digunakan agar klien dapat memilih kolom yang ingin ditampilkan.

```text
GET /api/mahasiswa?fields=nim,nama,angkatan
```

  >Kolom yang diizinkan adalah `nim`, `nama`, `angkatan`, dan `program_studi`. Kolom yang tidak terdaftar akan diabaikan.

  >Parameter lain yang tersedia:
  >
```text
GET /api/mahasiswa?cari=Rina
GET /api/mahasiswa?angkatan=2024
GET /api/mahasiswa?urut=nama&arah=desc
GET /api/mahasiswa?per_halaman=5
```

## 4. Daftar URL untuk Pengujian dan Screenshot

  >Jalankan server dari folder `latihan-laravel`:
  >
```powershell
cd "C:\Users\WORKPLUS\Downloads\H1H024056-Raihan-Pemweb2\latihan-laravel"
php artisan serve --host=127.0.0.1 --port=8000
```

  >Jika terminal tetap berada di folder utama project, gunakan:

```powershell
php "C:\Users\WORKPLUS\Downloads\H1H024056-Raihan-Pemweb2\latihan-laravel\artisan" serve --host=127.0.0.1 --port=8000
```

  >Buka URL berikut di browser dan jadikan screenshot:

  >1. http://127.0.0.1:8000/api/status
  >2. http://127.0.0.1:8000/api/mahasiswa
  >3. http://127.0.0.1:8000/api/mahasiswa?fields=nim,nama,angkatan
  >4. http://127.0.0.1:8000/api/matakuliah
  >5. http://127.0.0.1:8000/api/program-studi/Informatika/mahasiswa?per_halaman=5
  >6. http://127.0.0.1:8000/api/matakuliah/1
  >
  >Untuk POST, PUT, PATCH, dan DELETE gunakan Postman atau Insomnia karena browser biasa paling mudah digunakan untuk GET.

# F. Pertanyaan Pembahasan

## 1. Mengapa penamaan URI sebaiknya menggunakan kata benda jamak dan bukan kata kerja?

  >URI sebaiknya menggunakan kata benda jamak karena URI mewakili resource atau kumpulan data, bukan tindakan. Contohnya `/api/mahasiswa` untuk kumpulan data mahasiswa dan `/api/matakuliah` untuk kumpulan data mata kuliah.

  >Tindakan sudah diwakili oleh metode HTTP. `GET` digunakan untuk membaca, `POST` untuk menambah, `PUT` atau `PATCH` untuk memperbarui, dan `DELETE` untuk menghapus data. Dengan demikian, URI menjadi konsisten dan mudah dipahami.

  >URI seperti `/api/getMahasiswa` kurang tepat karena kata `get` sudah diwakili oleh metode HTTP `GET`.

## 2. Jelaskan perbedaan status 401 dan 403 beserta contoh kasusnya.

  >Status `401 Unauthorized` berarti permintaan belum memiliki autentikasi yang valid. Contohnya pengguna mengakses endpoint yang membutuhkan token, tetapi tidak mengirimkan token atau tokennya sudah tidak berlaku.

  >Status `403 Forbidden` berarti pengguna sudah berhasil dikenali, tetapi tidak memiliki izin untuk melakukan tindakan tersebut. Contohnya pengguna biasa mencoba menghapus data, sedangkan hanya administrator yang memiliki izin menghapus.

  >Jadi, `401` berkaitan dengan identitas yang belum valid, sedangkan `403` berkaitan dengan izin yang tidak mencukupi.

## 3. Apa risiko keamanan apabila parameter pengurutan pada Langkah 5 diterima tanpa pemeriksaan daftar kolom yang diizinkan?

  >Jika parameter pengurutan langsung dimasukkan ke `orderBy()` tanpa pemeriksaan, klien dapat mengirim nama kolom atau ekspresi SQL yang tidak seharusnya digunakan. Hal ini dapat menyebabkan manipulasi query, kebocoran informasi, atau kesalahan pada basis data.

  >Pada project ini, parameter `urut` diperiksa menggunakan daftar kolom yang diizinkan:

```php
$kolomDiizinkan = ['nim', 'nama', 'angkatan', 'program_studi'];

if (in_array($urut, $kolomDiizinkan, true)) {
    $query->orderBy($urut, $arah);
}
```

  >Arah pengurutan juga dibatasi hanya menjadi `asc` atau `desc`.

## 4. Mengapa API Resource lebih baik daripada mengembalikan model secara langsung?

  >API Resource memberikan lapisan untuk menentukan data apa saja yang boleh dikirim ke klien. Dengan begitu, data internal atau kolom sensitif tidak ikut terbuka secara tidak sengaja.

  >API Resource juga membuat format respons konsisten, memudahkan perubahan struktur database tanpa langsung mengubah kontrak API, dan memudahkan penambahan data turunan atau relasi.

  >Pada project ini, `MahasiswaResource` dan `MatakuliahResource` menentukan format JSON yang dikirim melalui API, sehingga respons lebih rapi dan terkontrol.
