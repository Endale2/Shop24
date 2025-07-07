import axios from 'axios'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import router from '../router'

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
          // Get current route to extract shop slug for login redirect
          const currentRoute = router.currentRoute.value
          const shopSlug = currentRoute.params.shopSlug
          if (shopSlug) {
            router.push(`/shops/${shopSlug}/login`)
          } else {
            router.push('/')
          }
          return Promise.reject(error)
        }
        await authStore.refreshToken()
        return api(originalRequest)
      } catch (refreshError) {
        const authStore = useAuthStore()
        authStore.user = null
        // Get current route to extract shop slug for login redirect
        const currentRoute = router.currentRoute.value
        const shopSlug = currentRoute.params.shopSlug
        if (shopSlug) {
          router.push(`/shops/${shopSlug}/login`)
        } else {
          router.push('/')
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
