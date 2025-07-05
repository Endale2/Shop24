import api from './api'

export const customerSegmentService = {
  /**
   * Fetch all customer segments for a given shop.
   * @param {string} shopId
   * @returns {Promise<Array<Object>>}
   */
  async fetchAll(shopId) {
    const res = await api.get(`/seller/shops/${shopId}/customer-segments`)
    return res.data
  },

  /**
   * Create a new customer segment.
   * @param {string} shopId
   * @param {Object} data
   * @param {string} data.name
   * @param {string} [data.description]
   * @param {string} [data.color]
   * @returns {Promise<Object>}
   */
  async create(shopId, data) {
    const res = await api.post(`/seller/shops/${shopId}/customer-segments`, data)
    return res.data
  },

  /**
   * Update an existing customer segment.
   * @param {string} shopId
   * @param {string} segmentId
   * @param {Object} data — fields to update
   * @returns {Promise<Object>}
   */
  async update(shopId, segmentId, data) {
    const res = await api.patch(`/seller/shops/${shopId}/customer-segments/${segmentId}`, data)
    return res.data
  },

  /**
   * Delete a customer segment.
   * @param {string} shopId
   * @param {string} segmentId
   * @returns {Promise<void>}
   */
  async delete(shopId, segmentId) {
    await api.delete(`/seller/shops/${shopId}/customer-segments/${segmentId}`)
  },

  /**
   * Add a customer to a segment.
   * @param {string} shopId
   * @param {string} segmentId
   * @param {string} customerId
   * @returns {Promise<Object>}
   */
  async addCustomer(shopId, segmentId, customerId) {
    const res = await api.post(`/seller/shops/${shopId}/customer-segments/${segmentId}/customers`, {
      customerId
    })
    return res.data
  },

  /**
   * Remove a customer from a segment.
   * @param {string} shopId
   * @param {string} segmentId
   * @param {string} customerId
   * @returns {Promise<Object>}
   */
  async removeCustomer(shopId, segmentId, customerId) {
    const res = await api.delete(`/seller/shops/${shopId}/customer-segments/${segmentId}/customers/${customerId}`)
    return res.data
  }
} 