/**
 * Lapisan abstraksi backend (Tugas 9.E3).
 *
 * Satu pengaturan VITE_BACKEND pada berkas .env menentukan backend yang
 * dipakai: 'laravel' atau 'fiber'. Seluruh perbedaan kontrak kedua backend
 * (baseURL, nama parameter, nama field kata sandi, kunci API, struktur meta,
 * payload tambah, dan bentuk galat validasi) dikumpulkan di berkas ini
 * sehingga halaman tidak perlu berubah saat backend berganti.
 */

const pilihan = (import.meta.env.VITE_BACKEND || 'laravel').toLowerCase()

const daftarBackend = {
  laravel: {
    nama: 'Laravel',
    baseURL: import.meta.env.VITE_API_LARAVEL,
    paramHalaman: 'page',
    kataSandi: 'password',
    kunciApi: '',
    // Form tambah: Laravel menerima program_studi berupa teks.
    kolomForm: ['nim', 'nama', 'angkatan', 'program_studi'],
    dataTambah(isian) {
      return {
        nim: isian.nim,
        nama: isian.nama,
        angkatan: Number(isian.angkatan),
        program_studi: isian.program_studi,
      }
    },
  },
  fiber: {
    nama: 'Fiber',
    baseURL: import.meta.env.VITE_API_FIBER,
    paramHalaman: 'halaman',
    kataSandi: 'kata_sandi',
    kunciApi: import.meta.env.VITE_API_KEY_FIBER || '',
    // Form tambah: Fiber mewajibkan email dan program_studi_id (pilihan ID).
    kolomForm: ['nim', 'nama', 'email', 'angkatan', 'program_studi_id', 'ipk'],
    dataTambah(isian) {
      return {
        nim: isian.nim,
        nama: isian.nama,
        email: isian.email,
        angkatan: Number(isian.angkatan),
        ipk: Number(isian.ipk || 0),
        program_studi_id: Number(isian.program_studi_id),
      }
    },
  },
}

/** Backend aktif sesuai satu pengaturan VITE_BACKEND. */
export const aktif = daftarBackend[pilihan] || daftarBackend.laravel

/**
 * Menormalkan meta pagination.
 * Laravel: { current_page, last_page, total, ... }
 * Fiber:   { halaman, per_halaman, total, total_halaman }
 */
export function normalisasiMeta(meta) {
  return {
    halaman: Number(meta?.current_page ?? meta?.halaman ?? 1),
    totalHalaman: Number(meta?.last_page ?? meta?.total_halaman ?? 1),
    total: Number(meta?.total ?? 0),
  }
}

/**
 * Mengambil galat validasi per-field dari respons server.
 * Laravel 422: { sukses, pesan, galat: { nim: ["..."] } }
 * Fiber (mahasiswa): { sukses, pesan } tanpa galat per-field,
 * sehingga hanya pesan umum yang ditampilkan.
 */
export function galatPerField(respons) {
  const galat = respons?.galat
  if (!galat || typeof galat !== 'object' || Array.isArray(galat)) {
    return {}
  }
  return Object.fromEntries(
    Object.entries(galat).map(([kunci, daftarPesan]) => [
      kunci,
      Array.isArray(daftarPesan) ? daftarPesan[0] : String(daftarPesan),
    ]),
  )
}
