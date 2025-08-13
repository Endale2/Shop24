// store/auth.js
import { defineStore } from 'pinia'
import { authService } from '@/services/auth'
import { useShopStore } from './shops'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    sessionLoading: false,
    lastAuthCheck: null,
    persistenceError: null
  }),
  getters: {
    isAuthenticated: (state) => !!state.user,
    isSessionExpired: (state) => {
      if (!state.lastAuthCheck) return false
      // Consider session expired after 30 minutes of no auth check
      return Date.now() - state.lastAuthCheck > 30 * 60 * 1000
    }
  },
  actions: {
    async fetchMe() {
      try {
        this.sessionLoading = true
        this.user = await authService.me()
        this.lastAuthCheck = Date.now()
        this.persistenceError = null
        return this.user
      } catch (error) {
        console.error('[AuthStore] Failed to fetch user:', error)
        // Only clear user if it's a 401/403 (unauthorized)
        if (error.response?.status === 401 || error.response?.status === 403) {
          this.user = null
          this.lastAuthCheck = null
        }
        throw error
      } finally {
        this.sessionLoading = false
      }
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
      try {
        await authService.logout()
      } catch (error) {
        console.error('[AuthStore] Logout error:', error)
        // Continue with local cleanup even if server logout fails
      } finally {
        this.$reset()
        const shopStore = useShopStore()
        shopStore.$reset()
      }
    },

    async updateProfile(payload) {
      const res = await authService.updateProfile(payload)
      // assume res.data is the updated profile
      this.user = res.data
      this.lastAuthCheck = Date.now()
    },

    /**
     * Verify session is still valid
     */
    async verifySession() {
      if (!this.user) return false

      try {
        await this.fetchMe()
        return true
      } catch (error) {
        console.error('[AuthStore] Session verification failed:', error)
        return false
      }
    },

    /**
     * Initiate Google OAuth login
     */
    loginWithGoogle() {
      authService.loginWithGoogle()
    },

    /**
     * Handle persistence errors gracefully
     */
    handlePersistenceError(error) {
      console.error('[AuthStore] Persistence error:', error)
      this.persistenceError = error.message || 'Failed to persist auth state'
      // Try to recover by re-fetching user data
      if (this.user) {
        this.fetchMe().catch(() => {
          // If recovery fails, clear the state
          this.user = null
          this.lastAuthCheck = null
        })
      }
    }
  },
  persist: {
    paths: ['user', 'lastAuthCheck'],
    storage: localStorage,
    serializer: {
      serialize: (value) => {
        try {
          return JSON.stringify(value)
        } catch (error) {
          console.error('[AuthStore] Serialization error:', error)
          return '{}'
        }
      },
      deserialize: (value) => {
        try {
          return JSON.parse(value)
        } catch (error) {
          console.error('[AuthStore] Deserialization error:', error)
          return {}
        }
      }
    },
    afterRestore: (context) => {
      // Check if session might be expired after restore
      if (context.store.isSessionExpired) {
        console.warn('[AuthStore] Session appears expired, will verify on next auth check')
      }
    }
  }
})
