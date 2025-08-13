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

// Enhanced token refresh queue logic with better error handling
let isRefreshing = false
let failedQueue = []
let refreshAttempts = 0
const MAX_REFRESH_ATTEMPTS = 3

function processQueue(err, token = null) {
  failedQueue.forEach(p => err ? p.reject(err) : p.resolve(token))
  failedQueue = []
}

// Expanded list of paths that should trigger refresh
const refreshablePaths = ['/auth/seller/me', '/seller/', '/shop/', '/dashboard/']
const nonRefreshablePaths = ['/auth/seller/login', '/auth/seller/register', '/auth/seller/refresh']

api.interceptors.response.use(
  response => response,
  async error => {
    const { config, response } = error
    const originalRequest = config
    const is401 = response?.status === 401
    const isRefreshableRequest = refreshablePaths.some(path => originalRequest.url.includes(path))
    const isNonRefreshable = nonRefreshablePaths.some(path => originalRequest.url.includes(path))

    // Don't attempt refresh for auth endpoints or if already retried
    if (is401 && isRefreshableRequest && !isNonRefreshable && !originalRequest._retry) {
      originalRequest._retry = true

      // If already refreshing, queue this request
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        })
          .then(() => api(originalRequest))
          .catch(err => Promise.reject(err))
      }

      // Start refresh process
      isRefreshing = true

      try {
        // Check if we've exceeded max refresh attempts
        if (refreshAttempts >= MAX_REFRESH_ATTEMPTS) {
          console.error('[API] Max refresh attempts exceeded, clearing auth state')
          // Import auth store dynamically to avoid circular dependency
          const { useAuthStore } = await import('@/store/auth')
          const authStore = useAuthStore()
          authStore.user = null
          authStore.lastAuthCheck = null
          processQueue(new Error('Max refresh attempts exceeded'))
          return Promise.reject(error)
        }

        refreshAttempts++
        console.log(`[API] Attempting token refresh (attempt ${refreshAttempts}/${MAX_REFRESH_ATTEMPTS})`)

        await api.post('/auth/seller/refresh')

        // Reset attempts on successful refresh
        refreshAttempts = 0
        console.log('[API] Token refresh successful')

        processQueue(null)
        return api(originalRequest)

      } catch (refreshError) {
        console.error('[API] Token refresh failed:', refreshError)

        // If refresh fails, clear auth state
        try {
          const { useAuthStore } = await import('@/store/auth')
          const authStore = useAuthStore()
          authStore.user = null
          authStore.lastAuthCheck = null
        } catch (importError) {
          console.error('[API] Failed to import auth store:', importError)
        }

        processQueue(refreshError)
        return Promise.reject(refreshError)

      } finally {
        isRefreshing = false
      }
    }

    // For non-401 errors or non-refreshable requests, just reject
    return Promise.reject(error)
  }
)

export default api
