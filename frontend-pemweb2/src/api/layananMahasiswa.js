import klien from './klien'
import { aktif } from './konfigurasi'

export async function ambilDaftarMahasiswa(parameter = {}) {
  const halaman = parameter.halaman ?? 1
  const cari = parameter.cari ?? ''
  const perHalaman = parameter.perHalaman ?? 10

  // Laravel: ?page=1 — Fiber: ?halaman=1 (disingkirkan oleh konfigurasi).
  const respons = await klien.get('/mahasiswa', {
    params: {
      [aktif.paramHalaman]: halaman,
      cari: cari,
      per_halaman: perHalaman,
    },
  })
  return respons.data
}

export async function ambilDetailMahasiswa(id) {
  const respons = await klien.get(`/mahasiswa/${id}`)
  return respons.data
}

export async function simpanMahasiswa(isian) {
  // Bentuk payload dibentuk oleh backend aktif (lihat konfigurasi.dataTambah).
  const respons = await klien.post('/mahasiswa', aktif.dataTambah(isian))
  return respons.data
}

export async function hapusMahasiswa(id) {
  const respons = await klien.delete(`/mahasiswa/${id}`)
  return respons.data
}

export async function ambilDaftarProgramStudi() {
  // Hanya backend Fiber yang menyediakan GET /api/program-studi.
  const respons = await klien.get('/program-studi')
  return respons.data?.data ?? []
}
