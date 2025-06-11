import api from './api'

export const authService = {
  register: async ({ username, email, password }) => {
    const res = await api.post('/auth/seller/register', { username, email, password })
    return res.data
  },
  login: async ({ email, password }) => {
    const res = await api.post('/auth/seller/login', { email, password })
    return res.data
  },
  logout: async () => {
    await api.post('/auth/seller/logout')
  },
  me: async () => {
    const res = await api.get('/auth/seller/me')
    return res.data
  },
  refresh: async () => {
    const res = await api.post('/auth/seller/refresh')
    return res.data
  },
  loginWithGoogle() {
    window.location.href = 'http://localhost:8080/auth/seller/oauth/google'
  }
}
