import axios from 'axios'
import { aktif } from './konfigurasi'

// baseURL, header kunci API, dan seluruh perilaku backend datang dari
// lapisan abstraksi konfigurasi.js — diganti hanya lewat VITE_BACKEND.
const klien = axios.create({
  baseURL: aktif.baseURL,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
  timeout: 10000,
})

klien.interceptors.request.use((konfigurasi) => {
  const token = sessionStorage.getItem('token')
  if (token) {
    konfigurasi.headers.Authorization = `Bearer ${token}`
  }
  if (aktif.kunciApi) {
    konfigurasi.headers['X-Api-Key'] = aktif.kunciApi
  }
  return konfigurasi
})

klien.interceptors.response.use(
  (respons) => respons,
  (galat) => {
    if (galat.response && galat.response.status === 401) {
      sessionStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(galat)
  },
)

export default klien
