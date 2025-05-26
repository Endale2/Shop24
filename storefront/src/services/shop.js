import api from './api'

export async function fetchShop(shopSlug) {
  const { data } = await api.get(`/shops/${shopSlug}`)
  return data
}