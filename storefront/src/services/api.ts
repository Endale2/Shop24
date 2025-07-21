import axios from 'axios'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import router from '../router'
import { getCurrentShopSlug } from './shop';

// Utility: Redirect to login or shop selection
function redirectToLogin() {
  const authStore = useAuthStore()
  authStore.user = null
  const shopSlug = getCurrentShopSlug()
  if (shopSlug) {
    router.push(`/${shopSlug}/login`)
  } else {
    router.push('/select-shop')
  }
}

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

// --- REQUEST INTERCEPTOR: Attach access token if present ---
api.interceptors.request.use(
  (config) => {
    // Try to get access token from cookies
    const match = document.cookie.match(/(?:^|; )access_token=([^;]*)/)
    if (match && match[1]) {
      config.headers = config.headers || {}
      config.headers['Authorization'] = `Bearer ${decodeURIComponent(match[1])}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// --- RESPONSE INTERCEPTOR: Handle 401, refresh token, queue requests ---
let isRefreshing = false
let failedQueue: any[] = []

function processQueue(error: any, token: string | null = null) {
  failedQueue.forEach(prom => {
    if (error) {
      prom.reject(error)
    } else {
      prom.resolve(token)
    }
  })
  failedQueue = []
}

api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config
    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        // Queue up requests while refresh is in progress
        return new Promise(function(resolve, reject) {
          failedQueue.push({resolve, reject})
        })
        .then(() => api(originalRequest))
        .catch(err => Promise.reject(err))
      }
      originalRequest._retry = true
      isRefreshing = true
      const authStore = useAuthStore()
      // Check if refresh token exists in cookies before attempting refresh
      const hasRefreshToken = document.cookie.includes('refresh_token')
      // Only redirect to login for protected routes
      const protectedRoutes = ['/account', '/orders', '/checkout'];
      const currentPath = router.currentRoute.value.path;
      if (!hasRefreshToken) {
        // Only redirect if on a protected route
        if (protectedRoutes.some(route => currentPath.startsWith(route))) {
          redirectToLogin()
        } else {
          authStore.user = null
        }
        processQueue(error, null)
        isRefreshing = false
        return Promise.reject(error)
      }
      try {
        await authStore.refreshToken()
        processQueue(null, null)
        isRefreshing = false
        return api(originalRequest)
      } catch (refreshError) {
        // Only redirect if on a protected route
        if (protectedRoutes.some(route => currentPath.startsWith(route))) {
          redirectToLogin()
        } else {
          authStore.user = null
        }
        processQueue(refreshError, null)
        isRefreshing = false
        return Promise.reject(refreshError)
      }
    }
    return Promise.reject(error)
  }
)

export default api

export function shopUrl(path: string) {
  const route = useRoute()
  const shopSlug = route.params.shopSlug as string
  return `/shops/${shopSlug}${path}`
}

// Alternative function that takes shopSlug as parameter
export function getShopUrl(shopSlug: string, path: string) {
  return `/shops/${shopSlug}${path}`
}

// Utility to extract shop slug from subdomain
export function getShopSlugFromSubdomain(): string | null {
  // e.g., sophia.localhost, sophia.127.0.0.1, sophia.example.com
  const host = window.location.hostname;
  // Remove port if present
  const [hostname] = host.split(':');
  // For localhost or 127.0.0.1, treat as no subdomain
  if (hostname === 'localhost' || hostname === '127.0.0.1') return null;
  const parts = hostname.split('.')
  if (parts.length >= 2) {
    return parts[0]
  }
  return null
}
