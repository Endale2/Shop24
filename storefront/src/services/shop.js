import api from './api'

export async function getShop(slug) {
  const { data } = await api.get(`/shops/${slug}`)
  return data
}
export async function getProducts(slug) {
  const { data } = await api.get(`/shops/${slug}/products`)
  return data
}
