import { defineStore } from 'pinia'
import { login, register, getProfile, logout, refreshToken } from '@/services/auth'
import { useCartStore } from './cart'
import piniaPersist from 'pinia-plugin-persistedstate'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as null | { username: string; email: string; _id: string; firstName?: string; lastName?: string; profilePic?: string },
    loading: false,
    error: null as null | string,
    sessionLoading: true,
  }),
  actions: {
    async login(email: string, password: string, shopId: string) {
      this.loading = true
      this.error = null
      try {
        const loginRes = await login(email, password, shopId)
        console.debug('Login response:', loginRes)
        const { data } = await getProfile()
        console.debug('Profile fetch success:', data)
        this.user = data
        
        // Refresh cart after successful login
        const cartStore = useCartStore()
        await cartStore.refreshCart()
      } catch (e: any) {
        console.debug('Login or profile fetch error:', e)
        this.error = e.response?.data?.error || 'Login failed'
        setTimeout(() => { this.error = null }, 5000)
      } finally {
        this.loading = false
      }
    },
    async register(username: string, email: string, password: string, shopId: string, firstName?: string, lastName?: string) {
      this.loading = true
      this.error = null
      try {
        await register({ username, email, password, shopId, firstName, lastName })
        const { data } = await getProfile()
        this.user = data
        
        // Refresh cart after successful registration
        const cartStore = useCartStore()
        await cartStore.refreshCart()
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Registration failed'
        setTimeout(() => { this.error = null }, 5000)
      } finally {
        this.loading = false
      }
    },
    async fetchProfile() {
      try {
        const { data } = await getProfile()
        this.user = data
      } catch {
        this.user = null
      }
    },
    async logout() {
      await logout()
      this.user = null
      
      // Clear cart state when user logs out
      const cartStore = useCartStore()
      cartStore.clearCartState()
    },
    async refreshToken() {
      this.loading = true
      this.error = null
      try {
        console.debug('Attempting to refresh token...')
        console.debug('Current cookies:', document.cookie)
        await refreshToken()
        console.debug('Token refresh successful')
        const { data } = await getProfile()
        this.user = data
        
        // Refresh cart after successful token refresh
        const cartStore = useCartStore()
        await cartStore.refreshCart()
      } catch (e: any) {
        console.debug('Token refresh failed:', e.response?.status, e.response?.data)
        this.error = e.response?.data?.error || 'Token refresh failed'
        this.user = null
        setTimeout(() => { this.error = null }, 5000)
      } finally {
        this.loading = false
      }
    },
    async verifySession() {
      this.sessionLoading = true
      try {
        await this.refreshToken()
        await this.fetchProfile()
      } catch {
        this.user = null
      } finally {
        this.sessionLoading = false
      }
    },
  },
  persist: {
    paths: ['user'],
  },
})
