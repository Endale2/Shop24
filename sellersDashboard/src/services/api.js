// src/services/api.js
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
});

// Automatic token refresh on 401
let isRefreshing = false;
let failedQueue = [];

function processQueue(error, result = null) {
  failedQueue.forEach(({ resolve, reject }) => {
    error ? reject(error) : resolve(result);
  });
  failedQueue = [];
}

api.interceptors.response.use(
  (res) => res,
  async (err) => {
    const original = err.config;
    if (
      err.response?.status === 401 &&
      !original._retry &&
      !original.url.includes('/auth/seller/refresh')
    ) {
      original._retry = true;
      if (isRefreshing) {
        return new Promise((resolve, reject) =>
          failedQueue.push({ resolve, reject })
        ).then(() => api(original));
      }
      isRefreshing = true;
      try {
        await api.post('/auth/seller/refresh');
        processQueue(null);
        return api(original);
      } catch (e) {
        processQueue(e, null);
        throw e;
      } finally {
        isRefreshing = false;
      }
    }
    throw err;
  }

  
);

export default api;
