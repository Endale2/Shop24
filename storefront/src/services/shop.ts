import { api, getShopUrl } from './api'

export interface Shop {
  id: string
  name: string
  slug: string
  image: string
  description: string
}

export function fetchShop(shopSlug: string): Promise<Shop> {
  return api.get(getShopUrl(shopSlug, '')).then(r => r.data)
}
