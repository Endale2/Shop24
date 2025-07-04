import { defineStore } from 'pinia'
import * as authService from '@/services/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as null | any,
    loading: false,
  }),
  actions: {
    async loadUser() {
      this.loading = true
      try {
        this.user = await authService.fetchMe()
      } finally {
        this.loading = false
      }
    },
  },
})
