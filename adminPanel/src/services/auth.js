import axios from 'axios';

const api = axios.create({
  withCredentials: true, // Include cookies for access and refresh tokens
});

let isRefreshing = false;
let failedQueue = [];

function processQueue(error, token = null) {
  failedQueue.forEach(prom => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  failedQueue = [];
}

api.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config;

    // If access token is expired and request hasn't already been retried
    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
        .then(() => api(originalRequest))
        .catch(err => Promise.reject(err));
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        await api.get('/auth/refresh'); // This sets new access_token cookie
        processQueue(null);
        return api(originalRequest); // Retry original request
      } catch (refreshErr) {
        processQueue(refreshErr, null);
        return Promise.reject(refreshErr);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  }
);

export default {
  user: null,

  login(credentials) {
    return api.post('/auth/admin/login', credentials);
  },

  async isAuthenticated() {
    try {
      const res = await api.get('/auth/admin/me');
      this.user = res.data;
      return true;
    } catch {
      this.user = null;
      return false;
    }
  },

  async getUser() {
    if (!this.user) {
      const res = await api.get('/auth/admin/me');
      this.user = res.data;
    }
    return this.user;
  },

  api, // expose configured instance in case you want to use it in other components
};
