// src/services/shop.js
import api from './api'

export const shopService = {
  /**
   * Fetch all shops for the current seller.
   * @returns {Promise<Array<Object>>}
   */
  async fetchShops() {
    const res = await api.get('/seller/shops')
    return res.data.map(s => ({
      id:          s._id ?? s.ID,
      name:        s.Name ?? s.name,
      ownerId:     s.OwnerID ?? s.ownerId,
      description: s.Description ?? s.description,
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }))
  },

  /**
   * Create a new shop.
   * @param {Object} payload
   * @param {string} payload.name
   * @param {string} payload.description
   * @returns {Promise<Object>} The newly created shop
   */
  async createShop({ name, description }) {
    const res = await api.post('/seller/shops', { name, description })
    // Some APIs return insertedId or InsertedID.$oid
    const newId = res.data.insertedId ?? res.data.InsertedID?.$oid
    return {
      id:          newId,
      name,
      description,
      ownerId:     res.data.ownerId ?? null,
      createdAt:   res.data.createdAt ?? new Date().toISOString(),
      updatedAt:   res.data.updatedAt ?? new Date().toISOString(),
    }
  },

  /**
   * Update an existing shop.
   * @param {string} shopId
   * @param {Object} data — fields to update (e.g. { name, description })
   * @returns {Promise<Object>}
   */
  async updateShop(shopId, data) {
    await api.put(`/seller/shops/${shopId}`, data)
    // Return the fresh object
    const res = await api.get(`/seller/shops/${shopId}`)
    const s = res.data
    return {
      id:          s._id ?? s.ID,
      name:        s.Name ?? s.name,
      ownerId:     s.OwnerID ?? s.ownerId,
      description: s.Description ?? s.description,
      createdAt:   s.CreatedAt ?? s.createdAt,
      updatedAt:   s.UpdatedAt ?? s.updatedAt,
    }
  },

  /**
   * Delete a shop.
   * @param {string} shopId
   * @returns {Promise<void>}
   */
  async deleteShop(shopId) {
    await api.delete(`/seller/shops/${shopId}`)
  }
}
