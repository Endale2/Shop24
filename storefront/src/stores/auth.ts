import { defineStore } from 'pinia'
import { requestOTP, verifyOTP, getProfile, logout, refreshToken } from '@/services/auth'
import { useCartStore } from './cart'
import piniaPersist from 'pinia-plugin-persistedstate'
import api from '@/services/api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as null | { email: string; id: string; firstName?: string; lastName?: string; profile_image?: string },
    loading: false,
    error: null as null | string,
    sessionLoading: true,
    otpRequested: false,
    email: '',
    profileComplete: false,
  }),
  actions: {
    async requestOTP(email: string) {
      this.loading = true
      this.error = null
      try {
        await requestOTP(email)
        this.otpRequested = true
        this.email = email
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Failed to send OTP'
        setTimeout(() => { this.error = null }, 5000)
      } finally {
        this.loading = false
      }
    },
    async verifyOTP(otp: string) {
      this.loading = true
      this.error = null
      try {
        const { data } = await verifyOTP(this.email, otp)
        this.user = data.profile
        this.profileComplete = data.profileComplete
        await this.fetchProfile()
        this.otpRequested = false
        this.email = ''
      } catch (e: any) {
        this.error = e.response?.data?.error || 'OTP verification failed'
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
        this.otpRequested = false
        this.email = ''
      }
    },
    async refreshToken() {
      this.loading = true
      this.error = null
      try {
        await refreshToken()
        const { data } = await getProfile()
        this.user = data
      } catch (e: any) {
        this.error = e.response?.data?.error || 'Token refresh failed'
        this.user = null
        setTimeout(() => { this.error = null }, 5000)
      } finally {
        this.loading = false
      }
    },
    async verifySession() {
      this.sessionLoading = true
      // Only try to refresh if a token exists in cookies
      if (!document.cookie.includes('refresh_token')) {
        this.user = null
        this.sessionLoading = false
        return
      }
      try {
        await this.refreshToken()
        await this.fetchProfile()
      } catch {
        this.user = null
      } finally {
        this.sessionLoading = false
      }
    },
    async logout() {
      await logout()
      this.user = null
      this.otpRequested = false
      this.email = ''
      this.profileComplete = false
    },
  },
  persist: {
    paths: ['user'],
  },
})
