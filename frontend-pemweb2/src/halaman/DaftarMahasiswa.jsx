import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { ambilDaftarMahasiswa, hapusMahasiswa } from '../api/layananMahasiswa'
import { normalisasiMeta } from '../api/konfigurasi'

export default function DaftarMahasiswa() {
  const [daftar, setDaftar] = useState([])
  const [meta, setMeta] = useState(null)
  const [halaman, setHalaman] = useState(1)
  const [cari, setCari] = useState('')
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const [pesanGalat, setPesanGalat] = useState('')
  const [pesanSukses, setPesanSukses] = useState('')
  const [segar, setSegar] = useState(0)
  const [permintaanHapus, setPermintaanHapus] = useState(null)
  const [sedangMenghapus, setSedangMenghapus] = useState(false)

  const navigasi = useNavigate()

  useEffect(() => {
    let dibatalkan = false

    async function muatData() {
      setSedangMemuat(true)
      setPesanGalat('')

      try {
        const hasil = await ambilDaftarMahasiswa({
          halaman: halaman,
          cari: cari,
          perHalaman: 10,
        })

        if (dibatalkan === false) {
          setDaftar(hasil.data)
          setMeta(normalisasiMeta(hasil.meta))
        }
      } catch (galat) {
        if (dibatalkan === false) {
          setPesanGalat(
            galat.response?.data?.pesan || 'Gagal memuat data mahasiswa',
          )
        }
      } finally {
        if (dibatalkan === false) {
          setSedangMemuat(false)
        }
      }
    }

    const penunda = setTimeout(muatData, 400)

    return () => {
      dibatalkan = true
      clearTimeout(penunda)
    }
  }, [halaman, cari, segar])

  function mintaKonfirmasi(mahasiswa) {
    setPesanGalat('')
    setPesanSukses('')
    setPermintaanHapus(mahasiswa)
  }

  async function konfirmasiHapus() {
    setSedangMenghapus(true)
    setPesanGalat('')

    try {
      const hasil = await hapusMahasiswa(permintaanHapus.id)
      setPesanSukses(hasil.pesan || 'Data mahasiswa berhasil dihapus')
      setPermintaanHapus(null)
      // Muat ulang daftar dari server setelah penghapusan berhasil.
      setSegar((nilai) => nilai + 1)
    } catch (galat) {
      setPesanGalat(galat.response?.data?.pesan || 'Gagal menghapus data mahasiswa')
      setPermintaanHapus(null)
    } finally {
      setSedangMenghapus(false)
    }
  }

  return (
    <div style={{ maxWidth: 900, margin: '40px auto' }}>
      <div
        style={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
        }}
      >
        <h1>Daftar Mahasiswa</h1>
        <button onClick={() => navigasi('/mahasiswa/tambah')}>
          Tambah Mahasiswa
        </button>
      </div>

      <input
        type="search"
        placeholder="Cari nama atau NIM"
        value={cari}
        onChange={(peristiwa) => {
          setHalaman(1)
          setCari(peristiwa.target.value)
        }}
        style={{ padding: 8, width: '100%' }}
      />

      {sedangMemuat && <p>Memuat data</p>}

      {pesanSukses && (
        <p style={{ color: 'green', border: '1px solid green', padding: 8 }}>
          {pesanSukses}
        </p>
      )}

      {pesanGalat && <p style={{ color: 'crimson' }}>{pesanGalat}</p>}

      {sedangMemuat === false && daftar.length === 0 && (
        <p>Data tidak ditemukan</p>
      )}

      <table
        border="1"
        cellPadding="8"
        style={{ width: '100%', marginTop: 16, borderCollapse: 'collapse' }}
      >
        <thead>
          <tr>
            <th>NIM</th>
            <th>Nama</th>
            <th>Angkatan</th>
            <th>IPK</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          {daftar.map((mahasiswa) => (
            <tr key={mahasiswa.id}>
              <td>{mahasiswa.nim}</td>
              <td>{mahasiswa.nama}</td>
              <td>{mahasiswa.angkatan}</td>
              <td>{mahasiswa.ipk ?? '-'}</td>
              <td>
                <button
                  onClick={() => mintaKonfirmasi(mahasiswa)}
                  style={{ color: 'crimson' }}
                >
                  Hapus
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {meta && (
        <div style={{ marginTop: 16 }}>
          <button
            onClick={() => setHalaman(halaman - 1)}
            disabled={meta.halaman <= 1}
          >
            Sebelumnya
          </button>
          <span style={{ margin: '0 12px' }}>
            Halaman {meta.halaman} dari {meta.totalHalaman}
          </span>
          <button
            onClick={() => setHalaman(halaman + 1)}
            disabled={meta.halaman >= meta.totalHalaman}
          >
            Berikutnya
          </button>
        </div>
      )}

      {permintaanHapus && (
        <div
          role="dialog"
          aria-modal="true"
          style={{
            position: 'fixed',
            inset: 0,
            background: 'rgba(0, 0, 0, 0.45)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            zIndex: 10,
          }}
        >
          <div
            style={{
              background: '#fff',
              padding: 24,
              borderRadius: 8,
              maxWidth: 380,
            }}
          >
            <h2 style={{ marginTop: 0 }}>Konfirmasi Hapus</h2>
            <p>
              Hapus mahasiswa <b>{permintaanHapus.nama}</b> (
              {permintaanHapus.nim})?
            </p>
            <button
              onClick={konfirmasiHapus}
              disabled={sedangMenghapus}
              style={{ color: 'crimson' }}
            >
              {sedangMenghapus ? 'Menghapus' : 'Hapus'}
            </button>{' '}
            <button onClick={() => setPermintaanHapus(null)} disabled={sedangMenghapus}>
              Batal
            </button>
          </div>
        </div>
      )}
    </div>
  )
}
