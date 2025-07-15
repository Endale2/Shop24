// src/services/auth.js
import api from './api'

export const authService = {
  /**
   * Step 1: Create a bare auth record (email+password).
   * The backend will issue tokens immediately.
   */
  register: async ({ email, password }) => {
    const res = await api.post('/auth/seller/register', { email, password })
    return res.data
  },

  /**
   * Email + password login.
   */
  login: async ({ email, password }) => {
    const res = await api.post('/auth/seller/login', { email, password })
    return res.data
  },

  /**
   * Logout (clears cookies server-side).
   */
  logout: async () => {
    await api.post('/auth/seller/logout')
  },

  /**
   * Fetch the current user (seller) profile.
   */
  me: async () => {
    const res = await api.get('/auth/seller/me')
    return res.data
  },

  /**
   * Refresh the access token.
   */
  refresh: async () => {
    const res = await api.post('/auth/seller/refresh')
    return res.data
  },

  /**
   * Step 2: Complete or update the seller profile.
   * Expects any subset of { firstName, lastName, phone, address, businessName, profileImage }.
   */
   updateProfile: async (payload) => {
    const res = await api.patch('/auth/seller/update-profile', payload)
    return res.data
  },

  /**
   * Kick off Google OAuth login.
   * Frontend will receive the final redirect after callback.
   */
  loginWithGoogle: () => {
    // Where in your SPA you want to land after OAuth:
    // Use the current page to maintain context (login or register)
    const currentPath = window.location.pathname
    const returnTo = window.location.origin + currentPath + '?oauth=google'
    const redirectUri = encodeURIComponent(returnTo)

    // Hit your backend's OAuth redirect endpoint:
    const apiBase = import.meta.env.VITE_API_BASE || 'http://localhost:8080'
    window.location.href =
      `${apiBase}/auth/seller/oauth/google?redirect_uri=${redirectUri}`
  },

  /**
   * Handle Telegram OAuth login.
   * This function is called by the Telegram widget with user data.
   */
  loginWithTelegram: async (telegramUser) => {
    const apiBase = import.meta.env.VITE_API_BASE || 'http://localhost:8080'
    
    // Send Telegram user data to backend for authentication
    const response = await fetch(`${apiBase}/auth/seller/oauth/telegram`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: telegramUser.id,
        first_name: telegramUser.first_name,
        last_name: telegramUser.last_name,
        username: telegramUser.username,
        photo_url: telegramUser.photo_url,
        auth_date: telegramUser.auth_date,
        hash: telegramUser.hash
      }),
      credentials: 'include' // Include cookies for session management
    })

    if (!response.ok) {
      throw new Error('Telegram authentication failed')
    }

    return response.json()
  }
}
