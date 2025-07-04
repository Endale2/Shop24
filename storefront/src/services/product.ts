import { api, shopUrl } from './api'

export interface Product {
  id: string
  name: string
  slug: string
  description: string
  main_image: string
  display_price: number
  // add other fields as needed
}

export function fetchAllProducts(): Promise<Product[]> {
  return api.get(shopUrl('/products')).then(r => r.data)
}
export function fetchProductDetail(handle: string): Promise<Product> {
  return api
    .get(shopUrl(`/products/${encodeURIComponent(handle)}`))
    .then(r => r.data)
}
