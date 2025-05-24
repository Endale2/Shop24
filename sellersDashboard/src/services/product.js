// src/services/product.js
import api from './api'

export const productService = {
  /**
   * Fetch all products for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchByShop(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/products`)
    return res.data.map(p => ({
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      userId:      p.user_id ?? p.UserID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images ?? [],
      category:    p.Category ?? p.category ?? null,
      price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
      variants:    p.Variants ?? p.variants ?? [],
      createdAt:   p.CreatedAt ?? p.createdAt,
      updatedAt:   p.UpdatedAt ?? p.updatedAt ?? null,
    }))
  },

  /**
   * Fetch a single product by ID.
   * @param {string} shopId
   * @param {string} productId
   * @returns {Promise<Object>}
   */
  async fetchById(shopId, productId) {
    const res = await api.get(`/seller/shops/${shopId}/products/${productId}`)
    const p = res.data
    return {
      id:          p.ID ?? p.id,
      shopId:      p.shop_id ?? p.ShopID,
      userId:      p.user_id ?? p.UserID,
      name:        p.Name ?? p.name,
      description: p.Description ?? p.description,
      images:      p.Images ?? p.images ?? [],
      category:    p.Category ?? p.category ?? null,
      price:       isFinite(p.Price) ? p.Price : isFinite(p.price) ? p.price : null,
      variants:    p.Variants ?? p.variants ?? [],
      createdAt:   p.CreatedAt ?? p.createdAt,
      updatedAt:   p.UpdatedAt ?? p.updatedAt ?? null,
    }
  },

  /**
   * Create a new product in a shop.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.name
   * @param {string} [data.description]
   * @param {number} [data.price]
   * @param {Array<string>} [data.images]
   * @param {string} [data.category]
   * @param {Array<Object>} [data.variants]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/products`, data)
    const newId = res.data.id ?? res.data.insertedId
    return this.fetchById(shopId, newId)
  },

  /**
   * Update an existing product.
   * @param {string} shopId
   * @param {string} productId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, productId, data) {
    await api.put(`/seller/shops/${shopId}/products/${productId}`, data)
    return this.fetchById(shopId, productId)
  },

  /**
   * Delete a product.
   * @param {string} shopId
   * @param {string} productId
   * @returns {Promise<void>}
   */
  async delete(shopId, productId) {
    await api.delete(`/seller/shops/${shopId}/products/${productId}`)
  }
}
