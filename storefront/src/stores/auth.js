// src/stores/auth.js
import { defineStore } from 'pinia'
import {
  loginCustomer,
  registerCustomer,
  oauthRedirect,
  fetchCurrentCustomer
} from '@/services/auth'
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
    // call on app startup or on route‐guard
    async fetchCurrent() {
      this.loading = true
      try {
        const customer = await fetchCurrentCustomer()
        this.user = {
          id:          customer.id,
          username:    customer.username,
          email:       customer.email,
          avatar:      customer.avatarUrl || null,
          displayName: (customer.firstName && customer.lastName)
                        ? `${customer.firstName} ${customer.lastName}`
                        : (customer.username || customer.email),
        }
      } catch {
        this.user = null
      } finally {
        this.loading = false
      }
    },

    async authenticate({ email, password }) {
      this.loading = true
      this.error = null
      try {
        const shopId = useShopStore().current?.id
        if (!shopId) throw new Error('No shopId available')

        const { customer } = await loginCustomer({ email, password, shopId })
        // backend sets cookies automatically
        this.user = {
          id:          customer.id,
          username:    customer.username,
          email:       customer.email,
          avatar:      customer.avatarUrl || null,
          displayName: (customer.firstName && customer.lastName)
                        ? `${customer.firstName} ${customer.lastName}`
                        : (customer.username || customer.email),
        }
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
      try {
        const shopId = useShopStore().current?.id
        if (!shopId) throw new Error('No shopId available')
        await registerCustomer({ ...profile, shopId })
        await this.authenticate({ email: profile.email, password: profile.password })
      } catch (e) {
        this.error = e.response?.data?.error || e.message
        throw e
      } finally {
        this.loading = false
      }
    },

    oauthLogin() {
      const shopId = useShopStore().current?.id
      if (!shopId) return console.error('No shopId for OAuth')
      oauthRedirect(shopId)
    },

    logout() {
      this.user = null
    },
  },
  persist: true,
})
