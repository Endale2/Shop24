import api from './api'

export interface RegisterPayload {
  firstName: string
  lastName: string
  username: string
  email: string
  password: string
  phone: string
  address: string
  city: string
  state: string
  postalCode: string
  country: string
  shopId: string
}
export interface LoginPayload {
  email: string
  password: string
  shopId: string
}

export async function register(data: { username: string; email: string; password: string; shopId: string; firstName?: string; lastName?: string }) {
  return api.post('/auth/customer/register', data)
}

export async function login(email: string, password: string, shopId: string) {
  return api.post('/auth/customer/login', { email, password, shopId })
}

export async function getProfile() {
  return api.get('/auth/customer/me')
}

export async function logout() {
  return api.post('/auth/customer/logout')
}

export async function refreshToken() {
  return api.post('/auth/customer/refresh')
}

/**
 * Initiate Google OAuth login for customers
 */
export function loginWithGoogle(shopId?: string) {
  // Build the redirect URL with shopId if provided
  const currentPath = window.location.pathname
  const returnTo = window.location.origin + currentPath + '?oauth=google'
  const redirectUri = encodeURIComponent(returnTo)
  
  // Add shopId to the OAuth URL if provided
  const shopIdParam = shopId ? `&shopId=${encodeURIComponent(shopId)}` : ''
  
  // Hit the customer OAuth redirect endpoint
  const apiBase = import.meta.env.VITE_API_BASE || 'http://localhost:8080'
  window.location.href = `${apiBase}/auth/customer/oauth/google?redirect_uri=${redirectUri}${shopIdParam}`
}
