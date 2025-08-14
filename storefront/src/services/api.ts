import axios from 'axios'
import { useAuthStore } from '../stores/auth'
import router from '../router'
import { getCurrentShopSlug } from './shop';

const host = window.location.hostname
const isLocalHost = host === 'localhost' || host === '127.0.0.1' || host.endsWith('.localhost')
const resolvedBaseURL = isLocalHost ? 'http://localhost:8080' : 'http://api.shop24.sbs'

const api = axios.create({
  baseURL: resolvedBaseURL,
  withCredentials: true,
})

let isRefreshing = false
let failedQueue: any[] = []

function processQueue(error: any) {
  failedQueue.forEach(prom => {
    if (error) prom.reject(error)
    else prom.resolve()
  })
  failedQueue = []
}

api.interceptors.response.use(
  (response) => response,
  async (error: any) => {
    const originalRequest = error.config as any

    // Handle 401 errors
    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      console.log('[API Interceptor] 401 error detected, attempting refresh...')

      // Don't retry auth endpoints to avoid infinite loops
      if (originalRequest.url && (
        String(originalRequest.url).includes('/auth/customer/me') ||
        String(originalRequest.url).includes('/auth/customer/refresh')
      )) {
        console.log('[API Interceptor] Auth endpoint failed, not retrying')
        const authStore = useAuthStore()
        authStore.user = null
        authStore.sessionLoading = false
        return Promise.reject(error)
      }

      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        })
          .then(() => api(originalRequest))
          .catch((err) => Promise.reject(err))
      }

      ;(originalRequest as any)._retry = true
      isRefreshing = true
      const authStore = useAuthStore()

      try {
        await authStore.refreshSession()
        processQueue(null)
        isRefreshing = false
        console.log('[API Interceptor] Token refresh successful, retrying original request')
        return api(originalRequest)
      } catch (refreshError) {
        console.log('[API Interceptor] Token refresh failed, setting user as unauthenticated')
        // Set user as unauthenticated and clear session loading
        authStore.user = null
        authStore.sessionLoading = false
        processQueue(refreshError)
        isRefreshing = false

        // Redirect to login if on protected route
        const protectedRoutes = ['/account', '/orders', '/checkout', '/cart', '/wishlist']
        const currentPath = router.currentRoute.value.path
        if (protectedRoutes.some(route => currentPath.includes(route))) {
          console.log('[API Interceptor] Redirecting to login')
          router.push('/login')
        }
        return Promise.reject(refreshError)
      }
    }
    return Promise.reject(error)
  }
)

export default api

export function shopUrl(path: string) {
  const slug = getShopSlugFromSubdomain() || getCurrentShopSlug()
  return `/shops/${slug}${path}`
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
