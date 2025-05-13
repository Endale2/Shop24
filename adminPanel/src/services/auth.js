import axios from 'axios';

let isRefreshing = false;
let refreshSubscribers = [];

function onRrefreshed(token) {
  refreshSubscribers.forEach((callback) => callback(token));
  refreshSubscribers = [];
}

function addRefreshSubscriber(callback) {
  refreshSubscribers.push(callback);
}

// Axios response interceptor for token refresh
axios.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      if (!isRefreshing) {
        isRefreshing = true;
        try {
          const res = await axios.post('/auth/admin/refresh');
          isRefreshing = false;
          onRrefreshed(); // No token to pass since you're using http-only cookies
        } catch (refreshError) {
          isRefreshing = false;
          return Promise.reject(refreshError);
        }
      }

      return new Promise((resolve) => {
        addRefreshSubscriber(() => {
          resolve(axios(originalRequest)); // retry original request
        });
      });
    }

    return Promise.reject(error);
  }
);

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
