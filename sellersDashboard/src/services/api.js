// src/services/api.js
import axios from 'axios';

const api = axios.create({
  baseURL:  'http://localhost:8080',
  withCredentials: true,
});

let isRefreshing = false;
let failedQueue = [];

function processQueue(error, token = null) {
  failedQueue.forEach(prom => {
    error ? prom.reject(error) : prom.resolve(token);
  });
  failedQueue = [];
}

api.interceptors.response.use(
  response => response,
  async error => {
    const { config, response } = error;
    const originalRequest = config;

    if (response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        }).then(() => api(originalRequest));
      }

      isRefreshing = true;
      try {
        await api.post('/auth/seller/refresh');
        processQueue(null);
        return api(originalRequest);
      } catch (err) {
        processQueue(err);
        return Promise.reject(err);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  }
);

export default api;
