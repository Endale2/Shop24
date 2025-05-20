// src/services/auth.js
import api from './api';

export default {
  // Registers then auto-logs in (cookies set by backend)
  async register({ username, email, password }) {
    const payload = { username, email, password };
    // Backend will set access and refresh cookies
    const res = await api.post('/auth/seller/register', payload);
    return res.data; // { message, seller: { … } }
  },

  // Logs in, sets cookies
  async login({ email, password }) {
    const res = await api.post('/auth/seller/login', { email, password });
    return res.data; // { message, seller: { … } }
  },

  // Clears HTTP-only cookies
  async logout() {
    await api.post('/auth/seller/logout');
  },

  // Checks /auth/seller/me for current seller profile
  async me() {
    const res = await api.get('/auth/seller/me');
    return res.data; // { id, name, … }
  },

  // Returns true if the user is logged in (refreshing if needed)
  async isAuthenticated() {
    try {
      await this.me();
      return true;
    } catch {
      return false;
    }
  },
};
