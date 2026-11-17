import klien from './klien'
import { aktif } from './konfigurasi'

export async function login(email, kataSandi) {
  const respons = await klien.post('/auth/login', {
    email: email,
    // Laravel memakai field 'password', Fiber memakai 'kata_sandi'.
    [aktif.kataSandi]: kataSandi,
  })
  return respons.data
}

export async function ambilProfil() {
  const respons = await klien.get('/auth/profil')
  return respons.data
}

export async function logout() {
  const respons = await klien.post('/auth/logout')
  return respons.data
}
