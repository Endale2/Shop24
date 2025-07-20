import axios from 'axios'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import router from '../router'
import { getCurrentShopSlug } from './shop';

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

// Add a response interceptor to handle 401 errors and refresh token
api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config
    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true
      try {
        const authStore = useAuthStore()
        // Check if refresh token exists in cookies before attempting refresh
        const hasRefreshToken = document.cookie.includes('refresh_token')
        if (!hasRefreshToken) {
          // No refresh token, clear user and redirect to login
          authStore.user = null
          // Get current shop slug from path
          const shopSlug = getCurrentShopSlug()
          if (shopSlug) {
            router.push(`/${shopSlug}/login`)
          } else {
            router.push('/select-shop')
          }
          return Promise.reject(error)
        }
        await authStore.refreshToken()
        return api(originalRequest)
      } catch (refreshError) {
        const authStore = useAuthStore()
        authStore.user = null
        // Get current shop slug from path
        const shopSlug = getCurrentShopSlug()
        if (shopSlug) {
          router.push(`/${shopSlug}/login`)
        } else {
          router.push('/select-shop')
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
  const parts = hostname.split('.');
  // e.g., [sophiaboutique, localhost]
  if (parts.length >= 2) {
    return parts[0];
  }
  return null;
}
