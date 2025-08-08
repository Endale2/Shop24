import axios from 'axios'

// Determine API base dynamically for local vs production
function resolveApiBase() {
  const envBase = import.meta?.env?.VITE_API_BASE
  if (envBase) return envBase
  const host = window.location.hostname
  if (host === 'localhost' || host === '127.0.0.1') {
    // mirror the current dev host to avoid cookie host mismatches
    return `http://${host}:8080`
  }
  return 'http://api.shop24.sbs'
}

const api = axios.create({
  baseURL: resolveApiBase(),
  withCredentials: true,
})

// token‐refresh queue logic, only for seller routes:
let isRefreshing = false
let failedQueue = []
function processQueue(err, token = null) {
  failedQueue.forEach(p => err ? p.reject(err) : p.resolve(token))
  failedQueue = []
}
const refreshablePaths = ['/auth/seller/me', '/seller/']

api.interceptors.response.use(
  r => r,
  async err => {
    const { config, response } = err
    const orig = config
    const is401 = response?.status === 401
    const needsRefresh = refreshablePaths.some(p => orig.url.includes(p))

    if (is401 && needsRefresh && !orig._retry) {
      orig._retry = true
      if (isRefreshing) {
        return new Promise((res, rej) => {
          failedQueue.push({ resolve: res, reject: rej })
        }).then(() => api(orig))
      }
      isRefreshing = true
      try {
        await api.post('/auth/seller/refresh')
        processQueue(null)
        return api(orig)
      } catch (refreshErr) {
        processQueue(refreshErr)
        return Promise.reject(refreshErr)
      } finally {
        isRefreshing = false
      }
    }

    return Promise.reject(err)
  }
)

export default api
