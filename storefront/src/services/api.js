// src/services/api.js
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

let isRefreshing = false
let failedQueue = []
function processQueue(error, token = null) {
  failedQueue.forEach(p => error ? p.reject(error) : p.resolve(token))
  failedQueue = []
}

// (your interceptor logic…)

export default api   // <<— default export
