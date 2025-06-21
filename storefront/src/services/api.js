// src/services/api.js
import axios from 'axios'

const BACKEND = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

const api = axios.create({
  baseURL: BACKEND,
  withCredentials: true,
})

export default api
