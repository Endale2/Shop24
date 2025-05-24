// src/services/api.js
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

// Helpers for queueing refresh
let isRefreshing = false
let failedQueue = []
function processQueue(error, token = null) {
  failedQueue.forEach(p => error ? p.reject(error) : p.resolve(token))
  failedQueue = []
}

// Only try to refresh on these protected paths:
const refreshablePaths = ['/auth/seller/me', '/seller/']  // substring match

api.interceptors.response.use(
  r => r,
  async err => {
    const { config, response } = err
    const originalRequest = config

    // If 401 *and* this is a refreshable endpoint, attempt token refresh
    const is401 = response?.status === 401
    const needsRefresh = refreshablePaths.some(p => originalRequest.url.includes(p))

    if (is401 && needsRefresh && !originalRequest._retry) {
      originalRequest._retry = true

      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        }).then(() => api(originalRequest))
      }

      isRefreshing = true
      try {
        await api.post('/auth/seller/refresh')   // your refresh endpoint
        processQueue(null)
        return api(originalRequest)
      } catch (refreshErr) {
        processQueue(refreshErr)
        return Promise.reject(refreshErr)
      } finally {
        isRefreshing = false
      }
    }

    // Otherwise just reject immediately (so public storefront .catch() can handle it)
    return Promise.reject(err)
  }
)

export default api
