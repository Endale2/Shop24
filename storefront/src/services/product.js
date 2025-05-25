import api from './api'


export async function fetchProducts(shopId) {
  const { data } = await api.get(`/shops/${shopId}/products`)
  return data
}

export async function fetchProductById(shopId, productId) {
  const { data } = await api.get(`/shops/${shopId}/products/${productId}`)
  return data
}
