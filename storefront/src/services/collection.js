// src/services/collection.js
import api from './api'

/**
 * Fetch all collections for a shop.
 * GET /shops/:shopSlug/collections
 */
export async function fetchCollections(shopSlug) {
  const { data } = await api.get(`/shops/${shopSlug}/collections`)
  return data
}

/**
 * Fetch a single collection by handle.
 * GET /shops/:shopSlug/collections/:handle
 */
export async function fetchCollectionDetail(shopSlug, handle) {
  const { data } = await api.get(`/shops/${shopSlug}/collections/${handle}`)
  return data
}
