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
      await authService.logout()
      this.$reset()
      const shopStore = useShopStore()
      shopStore.$reset()
    },
    async updateProfile(payload) {
      const res = await authService.updateProfile(payload)
      // assume res.data is the updated profile
      this.user = res.data
    },
    /**
     * Initiate Google OAuth login
     */
    loginWithGoogle() {
      authService.loginWithGoogle()
    },


  },
  persist: {
    paths: ['user']
  }
})
