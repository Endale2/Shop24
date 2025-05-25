import { api } from './api'

export async function fetchShop(shopId) {
  const { data } = await api.get(`/shops/${shopId}`)
  return data
}
