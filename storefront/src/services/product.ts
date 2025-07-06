import { api, getShopUrl } from './api'

export interface Product {
  id: string
  name: string
  slug: string
  description: string
  main_image: string
  display_price: number
  price: number
  stock: number
  category: string
  images: string[]
  discounts?: Array<{
    id: string
    name: string
    description: string
    category: string
    type: string
    value: number
    coupon_code?: string
    start_at: string
    end_at: string
    active: boolean
  }>
  // add other fields as needed
}

export function fetchAllProducts(shopSlug: string): Promise<Product[]> {
  return api.get(getShopUrl(shopSlug, '/products')).then(r => r.data)
}

export function fetchProductDetail(shopSlug: string, handle: string): Promise<Product> {
  return api
    .get(getShopUrl(shopSlug, `/products/${encodeURIComponent(handle)}`))
    .then(r => r.data)
}
