import axios from 'axios';

export default {
  user: null,

  login(credentials) {
    return axios.post('/auth/admin/login', credentials);
  },

  async isAuthenticated() {
    try {
      const res = await axios.get('/auth/admin/me');
      this.user = res.data;
      return true;
    } catch {
      this.user = null;
      return false;
    }
  },

  async getUser() {
    if (!this.user) {
      const res = await axios.get('/auth/admin/me');
      this.user = res.data;
    }
    return this.user;
  }
};