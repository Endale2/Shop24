import { api, shopUrl } from './api'

export interface Shop {
  id: string
  name: string
  slug: string
  image: string
  description: string
}

export function fetchShop(): Promise<Shop> {
  return api.get(shopUrl('')).then(r => r.data)
}
