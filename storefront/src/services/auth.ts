import { api } from './api'

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

export function register(payload: RegisterPayload) {
  return api.post('/auth/customer/register', payload)
}
export function login(payload: LoginPayload) {
  return api.post('/auth/customer/login', payload)
}
export function fetchMe() {
  return api.get('/auth/customer/me').then(r => r.data)
}
export function logout() {
  return api.post('/auth/customer/logout')
}
