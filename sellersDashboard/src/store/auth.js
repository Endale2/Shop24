import { defineStore } from 'pinia'
import { authService } from '@/services/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({ user: null }),
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
      await authService.logout()
      this.user = null
    }
  }
})
