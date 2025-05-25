import api from './api'

// Products by shop slug (used on Home and ProductsPage)
export async function fetchProducts(shopSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/products`)
  return data
}

// Single product by shop and product slug
export async function fetchProductBySlug(shopSlug, productSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/products/${productSlug}`)
  return data
}
