// src/services/auth.js
import api from './api'

export const authService = {
  /**
   * Register a new seller account.
   * @param {Object} payload
   * @param {string} payload.username
   * @param {string} payload.email
   * @param {string} payload.password
   * @returns {Promise<Object>} The newly created user data
   */
  register: async ({ username, email, password }) => {
    const res = await api.post('/auth/seller/register', { username, email, password })
    return res.data
  },

  /**
   * Log in an existing seller.
   * @param {Object} payload
   * @param {string} payload.email
   * @param {string} payload.password
   * @returns {Promise<Object>} The login response, typically including expiresIn
   */
  login: async ({ email, password }) => {
    const res = await api.post('/auth/seller/login', { email, password })
    return res.data
  },

  /**
   * Log out the current seller.
   * @returns {Promise<void>}
   */
  logout: async () => {
    await api.post('/auth/seller/logout')
  },

  /**
   * Fetch the currently authenticated seller's profile.
   * @returns {Promise<Object>} The user data
   */
  me: async () => {
    const res = await api.get('/auth/seller/me')
    return res.data
  },

  /**
   * Refresh the authentication token using the refresh endpoint.
   * @returns {Promise<Object>} The refresh response, typically including a new expiresIn
   */
  refresh: async () => {
    const res = await api.post('/auth/seller/refresh')
    return res.data
  },

  /**
   * Convenience to check if the seller is authenticated.
   * @returns {Promise<boolean>}
   */
  isAuthenticated: async () => {
    try {
      await authService.me()
      return true
    } catch {
      return false
    }
  },
}
