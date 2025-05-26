import api from './api'

export async function fetchProducts(shopSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/products`)
  return data
}

export async function fetchProductBySlug(shopSlug, productSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/products/${productSlug}`)
  return data
}