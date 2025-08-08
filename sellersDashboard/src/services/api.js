import axios from 'axios'

const api = axios.create({
  baseURL: 'http://api.shop24.sbs',
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
