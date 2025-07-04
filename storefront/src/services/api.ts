import axios from 'axios'

// Derive shop slug from subdomain (e.g. sophiaboutique.localhost)
const host = window.location.hostname
const slug = host.split('.')[0]

export const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

export function shopUrl(path: string) {
  return `/shops/${slug}${path}`
}
