import api, { getShopUrl } from './api'
import { getCurrentShopSlug } from './shop'

export interface Product {
  id: string
  _id?: string // For backward compatibility
  name: string
  slug: string
  description: string
  main_image: string
  display_price?: number
  price?: number
  starting_price?: number
  stock?: number
  total_stock?: number
  category: string
  images: string[]
  variants?: Array<{
    id: string
    variant_id?: string
    options: Array<{
      name: string
      value: string
    }>
    price: number
    stock: number
    image?: string
  }>
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

export async function fetchAllProducts(shopSlug?: string): Promise<Product[]> {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return [];
  try {
    const response = await api.get(getShopUrl(slug, '/products'))
    return response.data
  } catch (error) {
    console.error('Failed to fetch products:', error)
    return []
  }
}

export function fetchProductDetail(shopSlug: string, handle: string): Promise<Product> {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return Promise.reject('No shop slug');
  return api
    .get(getShopUrl(slug, `/products/${encodeURIComponent(handle)}`))
    .then(r => r.data)
}

export async function fetchProductDetailById(shopSlug: string, productId: string) {
  const slug = shopSlug || getCurrentShopSlug();
  if (!slug) return Promise.reject('No shop slug');
  return api.get(`/shops/${slug}/products/id/${productId}`);
}
