import api from './api'
import { getCurrentShopSlug } from './shop';

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

export async function requestOTP(email: string) {
  return api.post('/auth/customer/request-otp', { email });
}

export async function verifyOTP(email: string, otp: string) {
  return api.post('/auth/customer/verify-otp', { email, otp });
}

export async function getProfile() {
  return api.get('/auth/customer/me');
}

export async function logout() {
  return api.post('/auth/customer/logout');
}

export async function refreshToken() {
  return api.post('/auth/customer/refresh');
}

export async function updateCustomerProfile(profile: any) {
  // PATCH /auth/customer/me
  return api.patch('/auth/customer/me', profile)
}
