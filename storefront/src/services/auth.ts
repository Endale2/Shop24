import api from './api'
import { getCurrentShopSlug } from './shop';

export interface RegisterPayload {
  firstName?: string;
  lastName?: string;
  username: string;
  email: string;
  password: string;
  shopId: string;
}
export interface LoginPayload {
  email: string;
  password: string;
  shopId: string;
}

export async function register(data: RegisterPayload) {
  const shopId = data.shopId || getCurrentShopSlug();
  return api.post('/auth/customer/register', { ...data, shopId });
}

export async function login(data: LoginPayload) {
  const shopId = data.shopId || getCurrentShopSlug();
  return api.post('/auth/customer/login', { ...data, shopId });
}

export async function requestOTP(email: string) {
  const shopId = getCurrentShopSlug();
  return api.post('/auth/customer/request-otp', { email, shopId });
}

export async function verifyOTP(email: string, otp: string) {
  const shopId = getCurrentShopSlug();
  return api.post('/auth/customer/verify-otp', { email, otp, shopId });
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
  return api.patch('/auth/customer/me', profile);
}