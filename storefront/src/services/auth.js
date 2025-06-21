// src/services/auth.js
import api from './api'

export function loginCustomer(payload) {
  return api.post('/auth/customer/login', payload).then(res => res.data)
}

export function registerCustomer(payload) {
  return api.post('/auth/customer/register', payload).then(res => res.data)
}

// kicks off the Google flow
export function oauthRedirect(shopId) {
  const BACKEND = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
  window.location.href = `${BACKEND}/auth/customer/oauth/google?shopId=${shopId}`
}

// NEW: fetch the current logged‐in customer
export function fetchCurrentCustomer() {
  return api.get('/auth/customer/me').then(res => res.data)
}
