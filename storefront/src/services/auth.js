// src/services/auth.js
import api from './api'

export function loginCustomer(payload) {
  return api.post('/auth/customers/login', payload).then(res => res.data)
}

export function registerCustomer(payload) {
  return api.post('/auth/customers/register', payload).then(res => res.data)
}

export function oauthRedirect(shopId) {
  window.location.href = `${import.meta.env.VITE_API_BASE_URL}/auth/customer/oauth/redirect?shopId=${shopId}`
}
