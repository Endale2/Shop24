import axios from 'axios'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import router from '../router'
import { getCurrentShopSlug } from './shop';

const api = axios.create({
  baseURL: 'http://localhost:8080',
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
  response => response,
  async error => {
    const originalRequest = error.config
    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        })
          .then(() => api(originalRequest))
          .catch(err => Promise.reject(err))
      }
      originalRequest._retry = true
      isRefreshing = true
      const authStore = useAuthStore()
      try {
        await authStore.refreshSession()
        processQueue(null)
        isRefreshing = false
        return api(originalRequest)
      } catch (refreshError) {
        authStore.user = null
        processQueue(refreshError)
        isRefreshing = false
        // Optionally redirect to login if on protected route
        const protectedRoutes = ['/account', '/orders', '/checkout', '/cart', '/wishlist']
        const currentPath = router.currentRoute.value.path
        if (protectedRoutes.some(route => currentPath.startsWith(route))) {
          const shopSlug = getCurrentShopSlug()
          if (shopSlug) router.push(`/${shopSlug}/login`)
          else router.push('/select-shop')
        }
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
