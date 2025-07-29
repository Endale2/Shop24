import api, { getShopUrl } from './api'
import { getCurrentShopSlug } from './shop'

export interface Collection {
  handle: string
  id: string
  title: string
  image: string
}

export interface CollectionDetail extends Collection {
  description: string
  created_at: string
  filters: any
  products: Array<{
    id: string
    slug: string
    name: string
    main_image: string
    description: string
    price?: number
    display_price?: number
    variant_count?: number
    discounts?: Array<{
      id: string
      name: string
      type: string
      value: number
      active: boolean
    }>
    variants?: Array<any>
  }>
}

export async function fetchCollections(shopSlug?: string): Promise<Collection[]> {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return [];
  try {
    const response = await api.get(getShopUrl(slug, '/collections'))
    return response.data
  } catch (error) {
    console.error('Failed to fetch collections:', error)
    return []
  }
}

export async function fetchCollectionDetail(shopSlug: string, handle: string): Promise<CollectionDetail | null> {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return null;
  try {
    const response = await api.get(getShopUrl(slug, `/collections/${encodeURIComponent(handle)}`))
    return response.data
  } catch (error) {
    console.error('Failed to fetch collection detail:', error)
    return null
  }
}
