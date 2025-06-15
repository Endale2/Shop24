// store/auth.js
import { defineStore } from 'pinia'
import { authService } from '@/services/auth'
import { useShopStore } from './shops'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null
  }),
  getters: {
    isAuthenticated: (state) => !!state.user,
  },
  actions: {
    async fetchMe() {
      this.user = await authService.me()
      return this.user
    },
    async register(payload) {
      await authService.register(payload)
      return this.fetchMe()
    },
    async login(payload) {
      await authService.login(payload)
      return this.fetchMe()
    },
     async logout() {
      // Perform server-side logout
      await authService.logout()
      // Reset auth store state
      this.$reset()
      // Completely clear shop store state (list and active)
      const shopStore = useShopStore()
      shopStore.$reset()
    }
  },
 persist: {
   paths: ['user']
 }
})
