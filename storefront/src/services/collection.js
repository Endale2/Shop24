import api from './api'

export async function fetchCollections(shopSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/collections`)
  return data
}

export async function fetchCollectionDetail(shopSlug, collId) {
  const { data } = await api.get(
    `/shops/${shopSlug}/collections/${collId}`
  )
  return data
}
