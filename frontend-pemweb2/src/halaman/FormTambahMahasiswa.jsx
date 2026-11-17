import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { simpanMahasiswa, ambilDaftarProgramStudi } from '../api/layananMahasiswa'
import { aktif, galatPerField } from '../api/konfigurasi'

const ISIAN_KOSONG = {
  nim: '',
  nama: '',
  email: '',
  angkatan: '',
  ipk: '',
  program_studi: '',
  program_studi_id: '',
}

const GAYA_INPUT = { width: '100%', padding: 8, marginTop: 4 }

export default function FormTambahMahasiswa() {
  const [isian, setIsian] = useState(ISIAN_KOSONG)
  const [galatLokasi, setGalatLokasi] = useState({})
  const [pesanGalat, setPesanGalat] = useState('')
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const [daftarProgramStudi, setDaftarProgramStudi] = useState([])

  const navigasi = useNavigate()

  // Backend Fiber memerlukan program_studi_id; Laravel memakai teks program_studi.
  const perluDaftarProgramStudi = aktif.kolomForm.includes('program_studi_id')

  useEffect(() => {
    if (!perluDaftarProgramStudi) {
      return
    }
    ambilDaftarProgramStudi()
      .then((daftar) => setDaftarProgramStudi(daftar))
      .catch(() => setDaftarProgramStudi([]))
  }, [perluDaftarProgramStudi])

  function isiField(kunci, nilai) {
    setIsian((lama) => ({ ...lama, [kunci]: nilai }))
  }

  async function tanganiKirim(peristiwa) {
    peristiwa.preventDefault()
    setPesanGalat('')
    setGalatLokasi({})
    setSedangMemuat(true)

    try {
      const hasil = await simpanMahasiswa(isian)
      if (hasil.sukses) {
        navigasi('/mahasiswa')
      }
    } catch (galat) {
      const isi = galat.response?.data
      if (isi) {
        // Pesan galat validasi dari server:
        // Laravel 422 memuat galat per-field; Fiber hanya pesan umum.
        setPesanGalat(isi.pesan || 'Data yang dikirim tidak valid')
        setGalatLokasi(galatPerField(isi))
      } else {
        setPesanGalat('Tidak dapat menghubungi server')
      }
    } finally {
      setSedangMemuat(false)
    }
  }

  function tampil(kunci) {
    return aktif.kolomForm.includes(kunci)
  }

  function galat(kunci) {
    return galatLokasi[kunci] ? (
      <p style={{ color: 'crimson', margin: '4px 0 0' }}>{galatLokasi[kunci]}</p>
    ) : null
  }

  return (
    <div style={{ maxWidth: 480, margin: '40px auto' }}>
      <h1>Tambah Mahasiswa</h1>
      <p style={{ color: '#555' }}>Backend aktif: {aktif.nama}</p>

      {pesanGalat && (
        <p
          role="alert"
          style={{ color: 'crimson', border: '1px solid crimson', padding: 8 }}
        >
          {pesanGalat}
        </p>
      )}

      <form onSubmit={tanganiKirim}>
        {tampil('nim') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="nim">NIM</label>
            <input
              id="nim"
              value={isian.nim}
              onChange={(peristiwa) => isiField('nim', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            />
            {galat('nim')}
          </div>
        )}

        {tampil('nama') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="nama">Nama</label>
            <input
              id="nama"
              value={isian.nama}
              onChange={(peristiwa) => isiField('nama', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            />
            {galat('nama')}
          </div>
        )}

        {tampil('email') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="email">Email</label>
            <input
              id="email"
              type="email"
              value={isian.email}
              onChange={(peristiwa) => isiField('email', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            />
            {galat('email')}
          </div>
        )}

        {tampil('angkatan') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="angkatan">Angkatan</label>
            <input
              id="angkatan"
              type="number"
              value={isian.angkatan}
              onChange={(peristiwa) => isiField('angkatan', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            />
            {galat('angkatan')}
          </div>
        )}

        {tampil('program_studi') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="program-studi">Program Studi</label>
            <input
              id="program-studi"
              value={isian.program_studi}
              onChange={(peristiwa) => isiField('program_studi', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            />
            {galat('program_studi')}
          </div>
        )}

        {tampil('program_studi_id') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="program-studi-id">Program Studi</label>
            <select
              id="program-studi-id"
              value={isian.program_studi_id}
              onChange={(peristiwa) => isiField('program_studi_id', peristiwa.target.value)}
              required
              style={GAYA_INPUT}
            >
              <option value="">— Pilih program studi —</option>
              {daftarProgramStudi.map((ps) => (
                <option key={ps.id} value={ps.id}>
                  {ps.kode} — {ps.nama}
                </option>
              ))}
            </select>
            {galat('program_studi_id')}
          </div>
        )}

        {tampil('ipk') && (
          <div style={{ marginBottom: 12 }}>
            <label htmlFor="ipk">IPK (opsional)</label>
            <input
              id="ipk"
              type="number"
              step="0.01"
              min="0"
              max="4"
              value={isian.ipk}
              onChange={(peristiwa) => isiField('ipk', peristiwa.target.value)}
              style={GAYA_INPUT}
            />
            {galat('ipk')}
          </div>
        )}

        <button type="submit" disabled={sedangMemuat}>
          {sedangMemuat ? 'Menyimpan' : 'Simpan'}
        </button>{' '}
        <button type="button" onClick={() => navigasi('/mahasiswa')}>
          Batal
        </button>
      </form>
    </div>
  )
}
