import api, { getShopUrl } from './api'

export interface Shop {
  id: string
  name: string
  slug: string
  image: string
  banner?: string
  description: string
}

export async function fetchShop(shopSlug: string): Promise<Shop | null> {
  try {
    const response = await api.get(getShopUrl(shopSlug, ''))
    return response.data
  } catch (error) {
    console.error('Failed to fetch shop:', error)
    return null
  }
}
