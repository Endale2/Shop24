// src/services/collection.js
import api from './api'
import { productService } from './product'

export const collectionService = {
  /**
   * Fetch all collections for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAllByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/collections`)
    if (!Array.isArray(res.data)) return [];
    return res.data.map(c => ({
      id:          c.id,
      shopId:      c.shop_id,
      title:       c.title,
      description: c.description,
      handle:      c.handle,
      image:       c.image,
      productIds:  c.product_ids ?? [],
      products:    Array.isArray(c.products) ? c.products : [],
      createdAt:   c.created_at || c.createdAt,
      updatedAt:   c.updated_at || c.updatedAt,
    }))
  },

  /**
   * Fetch a single collection by ID.
   * @param {string} shopId
   * @param {string} collectionId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, collectionId) {
    const res = await api.get(
      `/seller/shops/${shopId}/collections/${collectionId}`
    )
    const c = res.data
    return {
      id:          c.id,
      shopId:      shopId,
      title:       c.title,
      description: c.description,
      handle:      c.handle,
      image:       c.image,
      // this endpoint returns full product objects
      products:    (c.products ?? []).map(p => ({
        id:    p.id,
        name:  p.name,
        image: p.main_image || null,
        category: p.category || '',
        price: p.price ?? null,
        stock: p.stock ?? null,
        starting_price: p.starting_price ?? null,
      })),
      createdAt:   c.created_at,
      updatedAt:   c.updated_at,
    }
  },

  /**
   * Create a new collection.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.title
   * @param {string} [data.description]
   * @param {string} [data.handle]
   * @param {string} [data.image]
   * @param {Array<string>} [data.productIds]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const payload = {
      title:       data.title,
      description: data.description,
      handle:      data.handle,
      image:       data.image,
      product_ids: data.productIds
    }
    const res = await api.post(`/seller/shops/${shopId}/collections`, payload)
    const newId = res.data.id ?? res.data.insertedId
    return this.fetchById(shopId, newId)
  },

  /**
   * Update an existing collection.
   * @param {string} shopId
   * @param {string} collectionId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, collectionId, data) {
    const payload = {}
    if (data.title       !== undefined) payload.title       = data.title
    if (data.description !== undefined) payload.description = data.description
    if (data.handle      !== undefined) payload.handle      = data.handle
    if (data.image       !== undefined) payload.image       = data.image
    if (data.productIds  !== undefined) payload.product_ids = data.productIds

    await api.patch(
      `/seller/shops/${shopId}/collections/${collectionId}`,
      payload
    )
    return this.fetchById(shopId, collectionId)
  },

  /**
   * Delete a collection.
   * @param {string} shopId
   * @param {string} collectionId
   * @returns {Promise<void>}
   */
  async delete(shopId, collectionId) {
    await api.delete(
      `/seller/shops/${shopId}/collections/${collectionId}`
    )
  },

  /**
   * Add a product to a collection.
   * @param {string} shopId
   * @param {string} collectionId
   * @param {string} productId
   * @returns {Promise<void>}
   */
  async addProduct(shopId, collectionId, productId) {
    // Update the product's collection_id to the new collection
    await productService.patch(shopId, productId, { collection_id: collectionId })
  },

  /**
   * Remove a product from a collection.
   * @param {string} shopId
   * @param {string} collectionId
   * @param {string} productId
   * @returns {Promise<void>}
   */
  async removeProduct(shopId, collectionId, productId) {
    // Update the product's collection_id to null (remove from collection)
    await productService.patch(shopId, productId, { collection_id: null })
  }
}
