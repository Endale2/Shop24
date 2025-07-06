import axios from 'axios'
import { useRoute } from 'vue-router'

export const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
})

export function shopUrl(path: string) {
  const route = useRoute()
  const shopSlug = route.params.shopSlug as string
  return `/shops/${shopSlug}${path}`
}

// Alternative function that takes shopSlug as parameter
export function getShopUrl(shopSlug: string, path: string) {
  return `/shops/${shopSlug}${path}`
}
