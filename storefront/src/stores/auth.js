// src/stores/auth.js
import { defineStore } from 'pinia'
import { loginCustomer, registerCustomer, oauthRedirect } from '@/services/auth'
import { useShopStore } from './shop'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    loading: false,
    error: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.user,
  },
  actions: {
    async authenticate({ email, password }) {
      this.loading = true
      this.error = null
      const shopStore = useShopStore()
      try {
        const payload = { email, password, shopId: shopStore.current._id }
        const resp = await loginCustomer(payload)
        this.user = resp.customer
      } catch (e) {
        this.error = e.response?.data?.error || e.message
        throw e
      } finally {
        this.loading = false
      }
    },
    async register(profile) {
      this.loading = true
      this.error = null
      const shopStore = useShopStore()
      profile.shopId = shopStore.current._id
      try {
        await registerCustomer(profile)
        await this.authenticate({ email: profile.email, password: profile.password })
      } catch (e) {
        this.error = e.response?.data?.error || e.message
        throw e
      } finally {
        this.loading = false
      }
    },
    oauthLogin() {
      const shopStore = useShopStore()
      oauthRedirect(shopStore.current._id)
    },
    logout() {
      this.user = null
      this.error = null
    }
  },
  persist: true
})
